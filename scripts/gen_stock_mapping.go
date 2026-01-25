package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
)

// 作用：
// - 读取脚本同目录下的两个 JSON 文件：
//   - wid_to_warehouse_code.json（WID -> DSCO warehouseCode）
//   - sku_to_dsco_sku.json（领星 SKU -> DSCO SKU）
//
// - 输出两行可直接粘贴到根目录 .env 的配置（使用单引号包裹，避免转义问题）。
//
// 运行（PowerShell）：
//
//	go run .\scripts\gen_stock_mapping.go
//
// 运行（Git Bash）：
//
//	go run ./scripts/gen_stock_mapping.go
func main() {
	scriptDir := mustScriptDir()

	widPath := flag.String("wid", filepath.Join(scriptDir, "wid_to_warehouse_code.json"), "WID->warehouseCode JSON 文件路径（对象：string->string）")
	skuPath := flag.String("sku", filepath.Join(scriptDir, "sku_to_dsco_sku.json"), "SKU->DSCO SKU JSON 文件路径（对象：string->string）")
	flag.Parse()

	widMap, widExists, err := readStringMap(*widPath)
	if err != nil {
		fatalf("读取 wid 映射失败: %v", err)
	}
	skuMap, skuExists, err := readStringMap(*skuPath)
	if err != nil {
		fatalf("读取 sku 映射失败: %v", err)
	}

	fmt.Println()
	fmt.Println("# 复制下面两行到 .env（推荐用单引号，避免转义）")

	if widExists && len(widMap) > 0 {
		fmt.Printf("IPASS_STOCK_WID_TO_DSCO_WAREHOUSE_CODE_JSON='%s'\n", mustCompactJSON(widMap))
	} else {
		if !widExists {
			fmt.Println("# IPASS_STOCK_WID_TO_DSCO_WAREHOUSE_CODE_JSON='{}'  # skipped (file missing)")
		} else {
			fmt.Println("# IPASS_STOCK_WID_TO_DSCO_WAREHOUSE_CODE_JSON='{}'  # skipped (empty)")
		}
	}

	if skuExists && len(skuMap) > 0 {
		fmt.Printf("IPASS_STOCK_SKU_TO_DSCO_SKU_JSON='%s'\n", mustCompactJSON(skuMap))
	} else {
		if !skuExists {
			fmt.Println("# IPASS_STOCK_SKU_TO_DSCO_SKU_JSON='{}'  # skipped (file missing)")
		} else {
			fmt.Println("# IPASS_STOCK_SKU_TO_DSCO_SKU_JSON='{}'  # skipped (empty)")
		}
	}
}

func mustScriptDir() string {
	// 使用 runtime.Caller 定位源码文件路径，确保“同目录 json 文件”的约定在 go run 时也成立。
	_, file, _, ok := runtime.Caller(0)
	if !ok || strings.TrimSpace(file) == "" {
		fatalf("无法定位脚本路径（runtime.Caller 失败）")
	}
	return filepath.Dir(file)
}

func readStringMap(path string) (m map[string]string, exists bool, err error) {
	path = strings.TrimSpace(path)
	if path == "" {
		return nil, false, fmt.Errorf("path 不能为空")
	}

	b, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return map[string]string{}, false, nil
		}
		return nil, false, err
	}
	exists = true

	raw := strings.TrimSpace(string(b))
	if raw == "" {
		return map[string]string{}, true, nil
	}

	var tmp map[string]any
	if err := json.Unmarshal([]byte(raw), &tmp); err != nil {
		return nil, true, fmt.Errorf("%s 不是合法 JSON: %w", path, err)
	}

	m = make(map[string]string, len(tmp))
	for k, v := range tmp {
		key := strings.TrimSpace(k)
		if key == "" {
			return nil, true, fmt.Errorf("%s 存在空 key", path)
		}

		val, ok := v.(string)
		if !ok {
			return nil, true, fmt.Errorf("%s 的 value 必须是 string（key=%s）", path, key)
		}
		val = strings.TrimSpace(val)
		if val == "" {
			return nil, true, fmt.Errorf("%s 存在空 value（key=%s）", path, key)
		}

		m[key] = val
	}
	return m, true, nil
}

func mustCompactJSON(m map[string]string) string {
	// 为了输出稳定，可读性更好：对 key 排序后再输出。
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	ordered := make(map[string]string, len(m))
	for _, k := range keys {
		ordered[k] = m[k]
	}

	b, err := json.Marshal(ordered)
	if err != nil {
		fatalf("JSON 编码失败: %v", err)
	}
	return string(b)
}

func fatalf(format string, args ...any) {
	_, _ = fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}
