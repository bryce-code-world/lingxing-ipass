package dsco_lingxing

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"example.com/lingxing/golib/v2/sdk/lingxing"
	"example.com/lingxing/golib/v2/tool/logger"

	"lingxingipass/integration"
)

type skuMappingDiff struct {
	Added   map[string]string `json:"added"`
	Deleted map[string]string `json:"deleted"`
	Changed []skuMappingChange `json:"changed"`
}

type skuMappingChange struct {
	MSKU   string `json:"msku"`
	OldSKU string `json:"old_sku"`
	NewSKU string `json:"new_sku"`
}

type skuPairPullMeta struct {
	RunID       string `json:"run_id"`
	Domain      string `json:"domain"`
	Job         string `json:"job"`
	StartedAt   string `json:"started_at"`
	FinishedAt  string `json:"finished_at"`
	DurationMS  int64  `json:"duration_ms"`
	PlatformCode string `json:"platform_code"`
	Size        int    `json:"size"`

	StoreIDs []string `json:"store_ids"`

	OldCount int    `json:"old_count"`
	NewCount int    `json:"new_count"`
	Added    int    `json:"added"`
	Deleted  int    `json:"deleted"`
	Changed  int    `json:"changed"`
	OldHash  string `json:"old_hash"`
	NewHash  string `json:"new_hash"`

	Aborted bool   `json:"aborted"`
	Reason  string `json:"reason,omitempty"`
	Error   string `json:"error,omitempty"`

	Exports struct {
		BaseDir       string `json:"base_dir,omitempty"`
		OldMapping    string `json:"old_mapping,omitempty"`
		NewMapping    string `json:"new_mapping,omitempty"`
		Diff          string `json:"diff,omitempty"`
		Partial       string `json:"partial,omitempty"`
		Meta          string `json:"meta,omitempty"`
	} `json:"exports"`
}

// PullSKUPair 自动拉取领星“MSKU ↔ 本地 SKU”配对关系，并更新 runtime_config.mapping.sku。
//
// 规则（已确认）：
//  1. 拉取范围：runtime_config.mapping.shop 的 value（领星 store_id）去重后逐店铺拉取。
//  2. 平台过滤：请求携带 platform_codes=[env.yaml integration.lingxing.platform_code]。
//  3. 冲突：同一 store_id 下 MSKU->多个SKU 或 SKU->多个MSKU；以及跨 store 合并后的 MSKU 冲突，均视为配置错误：打 error 日志并退出，不写 runtime_config。
//  4. 并发保护：更新前后对 mapping.sku 做 hash 校验；如检测到被人工改动则退出不写。
//  5. 审计落盘：old/new/diff/partial/meta 写入 exports/sku-mapping-pull/{yyyy-MM-dd}/，由 cleanup_exports 统一清理。
func (d *Domain) PullSKUPair(ctx integration.TaskContext) (retErr error) {
	taskCtx := ctx.Ctx
	if taskCtx == nil {
		taskCtx = context.Background()
	}

	startedAt := time.Now().UTC()
	base := ctx.BaseLogFields()
	logger.Info(taskCtx, "task begin", append(base, "task", "pull_sku_pair")...)

	meta := skuPairPullMeta{
		RunID:      ctx.RunID,
		Domain:     ctx.Domain,
		Job:        string(ctx.Job),
		StartedAt:  startedAt.Format(time.RFC3339),
		PlatformCode: strconv.Itoa(d.env.Integration.LingXing.PlatformCode),
		Size:       ctx.Size,
	}

	exportBaseDir := strings.TrimSpace(d.env.Admin.Export.Dir)
	dateDir := startedAt.Format("2006-01-02")
	exportDir := ""
	if exportBaseDir != "" {
		exportDir = filepath.Join(exportBaseDir, "sku-mapping-pull", dateDir)
		meta.Exports.BaseDir = exportDir
	}

	defer func() {
		finishedAt := time.Now().UTC()
		meta.FinishedAt = finishedAt.Format(time.RFC3339)
		meta.DurationMS = time.Since(startedAt).Milliseconds()
		if exportDir != "" && meta.Exports.Meta != "" {
			_ = writeJSONFileAtomic(meta.Exports.Meta, meta)
		}

		fields := append(base,
			"task", "pull_sku_pair",
			"duration_ms", meta.DurationMS,
			"old_count", meta.OldCount,
			"new_count", meta.NewCount,
			"added", meta.Added,
			"deleted", meta.Deleted,
			"changed", meta.Changed,
			"old_hash", meta.OldHash,
			"new_hash", meta.NewHash,
			"aborted", meta.Aborted,
			"reason", meta.Reason,
			"exports_dir", meta.Exports.BaseDir,
			"exports_old", meta.Exports.OldMapping,
			"exports_new", meta.Exports.NewMapping,
			"exports_diff", meta.Exports.Diff,
			"exports_partial", meta.Exports.Partial,
			"exports_meta", meta.Exports.Meta,
		)
		if retErr != nil {
			logger.Error(taskCtx, "task end", append(fields, "result", "failed", "err", retErr)...)
			return
		}
		logger.Info(taskCtx, "task end", append(fields, "result", "ok")...)
	}()

	oldMapping := copyStringMap(ctx.Config.Mapping.SKU)
	meta.OldCount = len(oldMapping)
	meta.OldHash = hashStringMap(oldMapping)

	storeIDs := uniqueNonEmptyStrings(valuesOfStringMap(ctx.Config.Mapping.Shop))
	meta.StoreIDs = storeIDs
	if len(storeIDs) == 0 {
		meta.Aborted = true
		meta.Reason = "no_store_ids_from_mapping_shop"
		return nil
	}

	if strings.TrimSpace(meta.PlatformCode) == "" || meta.PlatformCode == "0" {
		meta.Aborted = true
		meta.Reason = "missing_platform_code"
		return errors.New("integration.lingxing.platform_code is empty/0")
	}
	platformCodes := []string{meta.PlatformCode}

	if ctx.SnapshotRuntimeConfig == nil || ctx.UpdateRuntimeConfig == nil {
		return errors.New("runner did not inject runtime config functions")
	}

	lx, err := d.lingxingClient(taskCtx)
	if err != nil {
		meta.Aborted = true
		meta.Reason = "init_lingxing_client_failed"
		retErr = err
		return retErr
	}

	newMapping := make(map[string]string, 1024)
	conflicts := make([]string, 0, 8)

	type storeState struct {
		mskuToSKU map[string]string
		skuToMSKU map[string]string
	}
	perStore := make(map[string]*storeState, len(storeIDs))

	var pulledItems int

	for _, storeID := range storeIDs {
		state := &storeState{
			mskuToSKU: make(map[string]string, 256),
			skuToMSKU: make(map[string]string, 256),
		}
		perStore[storeID] = state

		offset := 0
		for {
			out, raw, err := lx.Config.GetPairListV2WithRawBody(taskCtx, lingxing.PairListV2Request{
				Offset:        offset,
				Length:        ctx.Size,
				PlatformCodes: platformCodes,
				StoreIDs:      []string{storeID},
			})
			if err != nil {
				meta.Aborted = true
				meta.Reason = "pull_failed"
				meta.Error = err.Error()
				_ = writeSKUPairPartial(exportDir, ctx.RunID, oldMapping, newMapping, storeID, offset, raw, meta, err)
				retErr = err
				return retErr
			}

			if len(out.List) == 0 {
				break
			}

			for _, it := range out.List {
				itemStoreID := strings.TrimSpace(it.StoreID.String())
				if itemStoreID != "" && itemStoreID != storeID {
					conflicts = append(conflicts, fmt.Sprintf("store_id_mismatch: want=%s got=%s msku=%s sku=%s", storeID, itemStoreID, it.MSKU, it.SKU))
					continue
				}

				msku := strings.TrimSpace(it.MSKU)
				sku := strings.TrimSpace(it.SKU)
				if msku == "" || sku == "" {
					conflicts = append(conflicts, fmt.Sprintf("empty_msku_or_sku: store_id=%s msku=%q sku=%q", storeID, msku, sku))
					continue
				}

				if old, ok := state.mskuToSKU[msku]; ok && old != sku {
					conflicts = append(conflicts, fmt.Sprintf("store_msku_conflict: store_id=%s msku=%s sku1=%s sku2=%s", storeID, msku, old, sku))
					continue
				}
				state.mskuToSKU[msku] = sku

				if old, ok := state.skuToMSKU[sku]; ok && old != msku {
					conflicts = append(conflicts, fmt.Sprintf("store_sku_conflict: store_id=%s sku=%s msku1=%s msku2=%s", storeID, sku, old, msku))
					continue
				}
				state.skuToMSKU[sku] = msku

				if old, ok := newMapping[msku]; ok && old != sku {
					conflicts = append(conflicts, fmt.Sprintf("global_msku_conflict: msku=%s sku1=%s sku2=%s", msku, old, sku))
					continue
				}
				newMapping[msku] = sku
				pulledItems++
			}

			offset += len(out.List)
			if len(out.List) < ctx.Size {
				break
			}
		}
	}

	if len(conflicts) > 0 {
		meta.Aborted = true
		meta.Reason = "mapping_conflict"
		logger.Error(taskCtx, "pull_sku_pair conflict",
			append(base,
				"task", "pull_sku_pair",
				"conflicts", integration.JSONForLog(conflicts),
				"pulled_items", pulledItems,
			)...,
		)
		_ = writeSKUPairPartial(exportDir, ctx.RunID, oldMapping, newMapping, "", 0, "", meta, errors.New("mapping conflict"))
		return nil
	}

	meta.NewCount = len(newMapping)
	meta.NewHash = hashStringMap(newMapping)

	added, deleted, changed := diffStringMap(oldMapping, newMapping)
	meta.Added = len(added)
	meta.Deleted = len(deleted)
	meta.Changed = len(changed)

	diff := skuMappingDiff{
		Added:   added,
		Deleted: deleted,
		Changed: changed,
	}

	// 写审计文件（即使后续并发保护导致不更新，也保留这次拉取结果用于排查）。
	if exportDir != "" {
		if err := os.MkdirAll(exportDir, 0o755); err == nil {
			oldPath := filepath.Join(exportDir, fmt.Sprintf("old_mapping_sku_%s.json", ctx.RunID))
			newPath := filepath.Join(exportDir, fmt.Sprintf("new_mapping_sku_%s.json", ctx.RunID))
			diffPath := filepath.Join(exportDir, fmt.Sprintf("diff_%s.json", ctx.RunID))
			metaPath := filepath.Join(exportDir, fmt.Sprintf("meta_%s.json", ctx.RunID))

			meta.Exports.OldMapping = oldPath
			meta.Exports.NewMapping = newPath
			meta.Exports.Diff = diffPath
			meta.Exports.Meta = metaPath

			_ = writeJSONFileAtomic(oldPath, oldMapping)
			_ = writeJSONFileAtomic(newPath, newMapping)
			_ = writeJSONFileAtomic(diffPath, diff)
			_ = writeJSONFileAtomic(metaPath, meta)
		}
	}

	// 轻量并发保护：避免“人刚改完又被定时任务覆盖”。
	curRC, ok := ctx.SnapshotRuntimeConfig(ctx.Domain)
	if !ok {
		meta.Aborted = true
		meta.Reason = "runtime_config_not_loaded"
		return errors.New("runtime config not loaded")
	}
	curHash := hashStringMap(curRC.Config.Mapping.SKU)
	if curHash != meta.OldHash {
		meta.Aborted = true
		meta.Reason = "mapping_sku_changed_concurrently"
		logger.Warn(taskCtx, "pull_sku_pair aborted by concurrent change",
			append(base,
				"task", "pull_sku_pair",
				"expected_old_hash", meta.OldHash,
				"current_hash", curHash,
				"old_count", meta.OldCount,
				"current_count", len(curRC.Config.Mapping.SKU),
			)...,
		)
		if exportDir != "" && meta.Exports.Meta != "" {
			_ = writeJSONFileAtomic(meta.Exports.Meta, meta)
		}
		return nil
	}

	updated := curRC.Config
	updated.Mapping.SKU = newMapping
	if err := ctx.UpdateRuntimeConfig(taskCtx, ctx.Domain, updated); err != nil {
		meta.Aborted = true
		meta.Reason = "update_runtime_config_failed"
		meta.Error = err.Error()
		if exportDir != "" && meta.Exports.Meta != "" {
			_ = writeJSONFileAtomic(meta.Exports.Meta, meta)
		}
		retErr = err
		return retErr
	}

	logger.Info(taskCtx, "pull_sku_pair updated runtime_config.mapping.sku",
		append(base,
			"task", "pull_sku_pair",
			"store_count", len(storeIDs),
			"pulled_items", pulledItems,
			"old_hash", meta.OldHash,
			"new_hash", meta.NewHash,
			"old_count", meta.OldCount,
			"new_count", meta.NewCount,
			"diff_added", meta.Added,
			"diff_deleted", meta.Deleted,
			"diff_changed", meta.Changed,
		)...,
	)
	return nil
}

func writeSKUPairPartial(exportDir, runID string, oldMapping, newMapping map[string]string, storeID string, offset int, raw string, meta skuPairPullMeta, err error) error {
	if exportDir == "" {
		return nil
	}
	if mkErr := os.MkdirAll(exportDir, 0o755); mkErr != nil {
		return mkErr
	}
	partialPath := filepath.Join(exportDir, fmt.Sprintf("partial_%s.json", runID))
	metaPath := filepath.Join(exportDir, fmt.Sprintf("meta_%s.json", runID))
	meta.Exports.Partial = partialPath
	meta.Exports.Meta = metaPath
	if err != nil && meta.Error == "" {
		meta.Error = err.Error()
	}

	type partial struct {
		Meta      skuPairPullMeta    `json:"meta"`
		StoreID   string            `json:"store_id,omitempty"`
		Offset    int               `json:"offset,omitempty"`
		Raw       string            `json:"raw,omitempty"`
		Old       map[string]string `json:"old_mapping_sku"`
		Partial   map[string]string `json:"partial_new_mapping_sku"`
	}
	p := partial{
		Meta:    meta,
		StoreID: storeID,
		Offset:  offset,
		Raw:     raw,
		Old:     oldMapping,
		Partial: newMapping,
	}
	_ = writeJSONFileAtomic(partialPath, p)
	_ = writeJSONFileAtomic(metaPath, meta)
	return nil
}

func writeJSONFileAtomic(path string, v any) error {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, b, 0o644); err != nil {
		return err
	}
	if err := os.Rename(tmp, path); err == nil {
		return nil
	}
	// Windows 上 Rename 不允许覆盖已存在的文件；这里做一次“删除再替换”兜底。
	_ = os.Remove(path)
	return os.Rename(tmp, path)
}

func copyStringMap(m map[string]string) map[string]string {
	if len(m) == 0 {
		return map[string]string{}
	}
	out := make(map[string]string, len(m))
	for k, v := range m {
		out[k] = v
	}
	return out
}

func valuesOfStringMap(m map[string]string) []string {
	if len(m) == 0 {
		return nil
	}
	out := make([]string, 0, len(m))
	for _, v := range m {
		out = append(out, v)
	}
	return out
}

func hashStringMap(m map[string]string) string {
	if len(m) == 0 {
		return ""
	}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	h := sha256.New()
	for _, k := range keys {
		v := m[k]
		_, _ = h.Write([]byte(k))
		_, _ = h.Write([]byte{0})
		_, _ = h.Write([]byte(v))
		_, _ = h.Write([]byte{'\n'})
	}
	return hex.EncodeToString(h.Sum(nil))
}

func diffStringMap(oldM, newM map[string]string) (added map[string]string, deleted map[string]string, changed []skuMappingChange) {
	added = make(map[string]string)
	deleted = make(map[string]string)

	for k, newV := range newM {
		oldV, ok := oldM[k]
		if !ok {
			added[k] = newV
			continue
		}
		if oldV != newV {
			changed = append(changed, skuMappingChange{MSKU: k, OldSKU: oldV, NewSKU: newV})
		}
	}
	for k, oldV := range oldM {
		if _, ok := newM[k]; !ok {
			deleted[k] = oldV
		}
	}

	sort.Slice(changed, func(i, j int) bool { return changed[i].MSKU < changed[j].MSKU })
	return added, deleted, changed
}
