package sync

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"example.com/lingxing/golib/v2/sdk/dsco"
	"example.com/lingxing/golib/v2/sdk/lingxing"
	"lingxingipass/internal/store"
)

func TestStockPipeline_SyncStock_Behavior(t *testing.T) {
	t.Parallel()

	// DSCO mock：POST /inventory/singleItem
	var upsertCalls int
	dscoSrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("method=%s want=%s", r.Method, http.MethodPost)
		}
		if r.URL.Path != "/api/v3/inventory/singleItem" {
			t.Fatalf("path=%s want=%s", r.URL.Path, "/api/v3/inventory/singleItem")
		}
		upsertCalls++

		var body dsco.ItemInventory
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode body err=%v", err)
		}
		if body.SKU != "SKU-1" || len(body.Warehouses) != 1 || body.Warehouses[0].Code != "WH1" || body.Warehouses[0].Quantity == nil || *body.Warehouses[0].Quantity != 681 {
			t.Fatalf("body=%+v", body)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"success":true}`)
	}))
	t.Cleanup(dscoSrv.Close)

	dscoCli, err := dsco.New(dsco.Config{
		BaseURL:    dscoSrv.URL + "/api/v3",
		HTTPClient: dscoSrv.Client(),
		Token:      "t",
	})
	if err != nil {
		t.Fatalf("dsco.New err=%v", err)
	}

	// 领星 mock：POST /inventoryDetails
	now := func() time.Time { return time.Unix(1720429074, 0) }
	lxSrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("method=%s want=%s", r.Method, http.MethodPost)
		}
		if r.URL.Path != "/erp/sc/routing/data/local_inventory/inventoryDetails" {
			t.Fatalf("path=%s want=%s", r.URL.Path, "/erp/sc/routing/data/local_inventory/inventoryDetails")
		}

		var req lingxing.InventoryDetailsRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("decode req err=%v", err)
		}
		if req.WID != "26" || req.Offset != 0 || req.Length != 200 {
			t.Fatalf("req=%+v", req)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"code":0,"message":"success","data":[{"wid":26,"product_id":10001,"sku":"SKU-1","product_total":1943,"product_valid_num":681}],"total":1}`)
	}))
	t.Cleanup(lxSrv.Close)

	lxCli, err := lingxing.New(lingxing.Config{
		BaseURL:     lxSrv.URL,
		AppID:       "1234567890abcdef",
		AccessToken: "tok",
		Now:         now,
	})
	if err != nil {
		t.Fatalf("lingxing.New err=%v", err)
	}

	gdb, _ := newMockGormDB(t)
	manual, err := store.NewManualTaskStore(gdb)
	if err != nil {
		t.Fatalf("NewManualTaskStore err=%v", err)
	}

	p, err := NewStockPipeline(dscoCli, lxCli, manual, map[string]string{"26": "WH1"}, nil)
	if err != nil {
		t.Fatalf("NewStockPipeline err=%v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	t.Cleanup(cancel)

	if err := p.SyncStock(ctx, 200); err != nil {
		t.Fatalf("SyncStock err=%v", err)
	}
	if upsertCalls != 1 {
		t.Fatalf("upsertCalls=%d want=1", upsertCalls)
	}
}
