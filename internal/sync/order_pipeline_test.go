package sync

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"

	"example.com/lingxing/golib/v2/sdk/dsco"
	"example.com/lingxing/golib/v2/sdk/lingxing"
	"lingxingipass/internal/store"
)

func TestOrderPipeline_PullOrders_PaginatesAndAdvancesWatermark(t *testing.T) {
	t.Parallel()

	// DSCO mock server：3 次：第一页、第二页、空页结束
	var calls int
	dscoSrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("method=%s want=%s", r.Method, http.MethodGet)
		}
		if r.URL.Path != "/api/v3/order/page" {
			t.Fatalf("path=%s want=%s", r.URL.Path, "/api/v3/order/page")
		}
		calls++

		q := r.URL.Query()
		switch calls {
		case 1:
			if q.Get("ordersUpdatedSince") == "" || q.Get("until") == "" {
				t.Fatalf("missing ordersUpdatedSince/until: %v", q)
			}
			_, _ = io.WriteString(w, `{"scrollId":"s1","orders":[{"dscoOrderId":"d1"}]}`)
		case 2:
			if q.Get("scrollId") != "s1" {
				t.Fatalf("scrollId=%q want %q", q.Get("scrollId"), "s1")
			}
			_, _ = io.WriteString(w, `{"scrollId":"s1","orders":[{"dscoOrderId":"d2"}]}`)
		default:
			if q.Get("scrollId") != "s1" {
				t.Fatalf("scrollId=%q want %q", q.Get("scrollId"), "s1")
			}
			_, _ = io.WriteString(w, `{"orders":[]}`)
		}
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

	// store 使用 sqlmock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New err=%v", err)
	}
	t.Cleanup(func() { _ = db.Close() })

	orderState, err := store.NewOrderStateStore(db)
	if err != nil {
		t.Fatalf("NewOrderStateStore err=%v", err)
	}
	watermark, err := store.NewWatermarkStore(db)
	if err != nil {
		t.Fatalf("NewWatermarkStore err=%v", err)
	}
	manual, err := store.NewManualTaskStore(db)
	if err != nil {
		t.Fatalf("NewManualTaskStore err=%v", err)
	}
	orderRaw, err := store.NewDscoOrderRawStore(db)
	if err != nil {
		t.Fatalf("NewDscoOrderRawStore err=%v", err)
	}

	mock.ExpectQuery("SELECT watermark FROM job_watermark").
		WithArgs("pull_dsco_orders").
		WillReturnRows(sqlmock.NewRows([]string{"watermark"}).AddRow([]byte(`{"mode":"updatedSince","since":"2025-01-01T00:00:00Z"}`)))

	mock.ExpectExec("INSERT INTO dsco_order_raw").
		WithArgs("d1", []byte(`{"dscoOrderId":"d1"}`), sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO sync_order_state").
		WithArgs("d1", sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec("INSERT INTO dsco_order_raw").
		WithArgs("d2", []byte(`{"dscoOrderId":"d2"}`), sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO sync_order_state").
		WithArgs("d2", sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// 水位推进：since=until（now-10s）
	mock.ExpectExec("INSERT INTO job_watermark").
		WithArgs("pull_dsco_orders", []byte(`{"mode":"updatedSince","since":"2025-01-01T00:10:10Z"}`)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	now := func() time.Time { return time.Date(2025, 1, 1, 0, 10, 20, 0, time.UTC) }

	p, err := NewOrderPipeline(dscoCli, nil, orderState, watermark, manual, orderRaw, now, 5, "delivered_at")
	if err != nil {
		t.Fatalf("NewOrderPipeline err=%v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	t.Cleanup(cancel)

	if err := p.PullOrders(ctx); err != nil {
		t.Fatalf("PullOrders err=%v", err)
	}
	if calls != 3 {
		t.Fatalf("calls=%d want=3", calls)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations err=%v", err)
	}
}

func TestOrderPipeline_PushOrdersToLingXing_Behavior(t *testing.T) {
	t.Parallel()

	// DSCO mock：回源单笔订单
	dscoSrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("method=%s want=%s", r.Method, http.MethodGet)
		}
		if r.URL.Path != "/api/v3/order/" {
			t.Fatalf("path=%s want=%s", r.URL.Path, "/api/v3/order/")
		}
		q := r.URL.Query()
		if q.Get("orderKey") != "dscoOrderId" || q.Get("value") != "d1" {
			t.Fatalf("query=%v want orderKey=dscoOrderId value=d1", q)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{
			"dscoOrderId":"d1",
			"currencyCode":"USD",
			"shipping":{"country":"US","city":"New York","name":"Tom","address1":"Street 1"},
			"lineItems":[{"quantity":1,"sku":"SKU-1","consumerPrice":9.99}]
		}`)
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

	// 领星 mock：创建订单成功
	now := func() time.Time { return time.Unix(1720429074, 0) }
	lxSrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("method=%s want=%s", r.Method, http.MethodPost)
		}
		if r.URL.Path != "/pb/mp/order/v2/create" {
			t.Fatalf("path=%s want=%s", r.URL.Path, "/pb/mp/order/v2/create")
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"code":0,"message":"success","data":{"error_details":[],"success_details":[{"platform_order_no":"d1","global_order_no":"g1"}]}}`)
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

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New err=%v", err)
	}
	t.Cleanup(func() { _ = db.Close() })

	orderState, err := store.NewOrderStateStore(db)
	if err != nil {
		t.Fatalf("NewOrderStateStore err=%v", err)
	}
	watermark, err := store.NewWatermarkStore(db)
	if err != nil {
		t.Fatalf("NewWatermarkStore err=%v", err)
	}
	manual, err := store.NewManualTaskStore(db)
	if err != nil {
		t.Fatalf("NewManualTaskStore err=%v", err)
	}
	orderRaw, err := store.NewDscoOrderRawStore(db)
	if err != nil {
		t.Fatalf("NewDscoOrderRawStore err=%v", err)
	}

	// ClaimForPush：BEGIN -> SELECT -> UPDATE -> COMMIT
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT dsco_order_id").
		WithArgs(sqlmock.AnyArg(), 10).
		WillReturnRows(sqlmock.NewRows([]string{"dsco_order_id"}).AddRow("d1"))
	mock.ExpectExec("UPDATE sync_order_state").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), "d1").
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	// MarkPushSuccess
	mock.ExpectExec("UPDATE sync_order_state").
		WithArgs(sqlmock.AnyArg(), "g1", sqlmock.AnyArg(), "d1").
		WillReturnResult(sqlmock.NewResult(0, 1))

	p, err := NewOrderPipeline(dscoCli, lxCli, orderState, watermark, manual, orderRaw, now, 5, "delivered_at")
	if err != nil {
		t.Fatalf("NewOrderPipeline err=%v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	t.Cleanup(cancel)

	if err := p.PushOrdersToLingXing(ctx, 10009, "s1", 10); err != nil {
		t.Fatalf("PushOrdersToLingXing err=%v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations err=%v", err)
	}
}

func TestOrderPipeline_PushOrdersToLingXing_RetryExceededGoesManual(t *testing.T) {
	t.Parallel()

	// DSCO mock：回源订单直接失败（触发失败+重试上限转人工）
	dscoSrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("method=%s want=%s", r.Method, http.MethodGet)
		}
		if r.URL.Path != "/api/v3/order/" {
			t.Fatalf("path=%s want=%s", r.URL.Path, "/api/v3/order/")
		}
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = io.WriteString(w, `{"message":"server error"}`)
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

	// 领星 client（本用例不应被调用，但 PushOrdersToLingXing 需要非空 client）
	now := func() time.Time { return time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC) }
	lxSrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatalf("unexpected lingxing call path=%s", r.URL.Path)
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

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New err=%v", err)
	}
	t.Cleanup(func() { _ = db.Close() })

	orderState, _ := store.NewOrderStateStore(db)
	watermark, _ := store.NewWatermarkStore(db)
	manual, _ := store.NewManualTaskStore(db)
	orderRaw, _ := store.NewDscoOrderRawStore(db)

	// ClaimForPush：BEGIN -> SELECT -> UPDATE -> COMMIT
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT dsco_order_id").
		WithArgs(sqlmock.AnyArg(), 1).
		WillReturnRows(sqlmock.NewRows([]string{"dsco_order_id"}).AddRow("d1"))
	mock.ExpectExec("UPDATE sync_order_state").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), "d1").
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	// MarkPushFailure
	mock.ExpectExec("UPDATE sync_order_state").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), "d1").
		WillReturnResult(sqlmock.NewResult(0, 1))

	// retry_count 已到上限：触发转人工
	mock.ExpectQuery("SELECT retry_count FROM sync_order_state").
		WithArgs("d1").
		WillReturnRows(sqlmock.NewRows([]string{"retry_count"}).AddRow(5))
	mock.ExpectExec("INSERT INTO manual_task").
		WithArgs("max_retry_exceeded", sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("UPDATE sync_order_state").
		WithArgs("max_retry_exceeded", sqlmock.AnyArg(), "d1").
		WillReturnResult(sqlmock.NewResult(0, 1))

	p, err := NewOrderPipeline(dscoCli, lxCli, orderState, watermark, manual, orderRaw, now, 5, "delivered_at")
	if err != nil {
		t.Fatalf("NewOrderPipeline err=%v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	t.Cleanup(cancel)

	// batchSize=1，确保只处理一条
	_ = p.PushOrdersToLingXing(ctx, 10009, "s1", 1)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations err=%v", err)
	}
}

func TestOrderPipeline_AckToDSCO_Behavior(t *testing.T) {
	t.Parallel()

	// DSCO mock：acknowledge
	var ackCalls int
	dscoSrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("method=%s want=%s", r.Method, http.MethodPost)
		}
		if r.URL.Path != "/api/v3/order/acknowledge" {
			t.Fatalf("path=%s want=%s", r.URL.Path, "/api/v3/order/acknowledge")
		}
		ackCalls++
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"status":"success","requestId":"r1","eventDate":"2025-01-01T00:00:00Z"}`)
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

	// 领星 mock：list 返回待发货订单（status=5），平台单号=d1
	now := func() time.Time { return time.Unix(1720429074, 0) }
	lxSrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("method=%s want=%s", r.Method, http.MethodPost)
		}
		if r.URL.Path != "/pb/mp/order/v2/list" {
			t.Fatalf("path=%s want=%s", r.URL.Path, "/pb/mp/order/v2/list")
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"code":0,"message":"success","data":{"total":1,"list":[{"platform_info":[{"platform_order_no":"d1"}]}]}}`)
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

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New err=%v", err)
	}
	t.Cleanup(func() { _ = db.Close() })

	orderState, err := store.NewOrderStateStore(db)
	if err != nil {
		t.Fatalf("NewOrderStateStore err=%v", err)
	}
	watermark, err := store.NewWatermarkStore(db)
	if err != nil {
		t.Fatalf("NewWatermarkStore err=%v", err)
	}
	manual, err := store.NewManualTaskStore(db)
	if err != nil {
		t.Fatalf("NewManualTaskStore err=%v", err)
	}
	orderRaw, err := store.NewDscoOrderRawStore(db)
	if err != nil {
		t.Fatalf("NewDscoOrderRawStore err=%v", err)
	}

	mock.ExpectQuery("SELECT watermark FROM job_watermark").
		WithArgs("ack_to_dsco").
		WillReturnRows(sqlmock.NewRows([]string{"watermark"}).AddRow([]byte(`{"mode":"update_time","since":1710925191}`)))
	mock.ExpectExec("INSERT INTO sync_order_state").
		WithArgs("d1", sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("UPDATE sync_order_state").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), "d1").
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec("UPDATE sync_order_state").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), "d1").
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec("INSERT INTO job_watermark").
		WithArgs("ack_to_dsco", []byte(`{"mode":"update_time","since":1720429064}`)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	p, err := NewOrderPipeline(dscoCli, lxCli, orderState, watermark, manual, orderRaw, now, 5, "delivered_at")
	if err != nil {
		t.Fatalf("NewOrderPipeline err=%v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	t.Cleanup(cancel)

	if err := p.AckToDSCO(ctx, 10009, "s1"); err != nil {
		t.Fatalf("AckToDSCO err=%v", err)
	}
	if ackCalls != 1 {
		t.Fatalf("ackCalls=%d want=1", ackCalls)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations err=%v", err)
	}
}

func TestOrderPipeline_ShipToDSCO_Behavior(t *testing.T) {
	t.Parallel()

	// DSCO mock：singleShipment
	var shipCalls int
	dscoSrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("method=%s want=%s", r.Method, http.MethodPost)
		}
		if r.URL.Path != "/api/v3/order/singleShipment" {
			t.Fatalf("path=%s want=%s", r.URL.Path, "/api/v3/order/singleShipment")
		}
		shipCalls++
		var body dsco.ShipmentsForUpdate
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode body err=%v", err)
		}
		if body.DscoOrderID != "d1" || len(body.Shipments) != 1 || body.Shipments[0].TrackingNumber != "TN1" {
			t.Fatalf("body=%+v", body)
		}
		if body.Shipments[0].ShipDate != "2021-01-01T00:00:00Z" {
			t.Fatalf("ShipDate=%q want %q", body.Shipments[0].ShipDate, "2021-01-01T00:00:00Z")
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

	// 领星 mock：list(status=6) + wmsOrderList(status=3,tracking_no)
	now := func() time.Time { return time.Unix(1720429074, 0) }
	lxSrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/pb/mp/order/v2/list":
			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, `{"code":0,"message":"success","data":{"total":1,"list":[{"platform_info":[{"platform_order_no":"d1"}]}]}}`)
		case "/erp/sc/routing/wms/order/wmsOrderList":
			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, `{"code":0,"message":"success","data":[{"platform_order_no":["d1"],"tracking_no":"TN1","delivered_at":"2021-01-01 00:00:00","product_info":[{"sku":"SKU-1","count":1}]}],"total":1}`)
		default:
			t.Fatalf("unexpected path=%s", r.URL.Path)
		}
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

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New err=%v", err)
	}
	t.Cleanup(func() { _ = db.Close() })

	orderState, err := store.NewOrderStateStore(db)
	if err != nil {
		t.Fatalf("NewOrderStateStore err=%v", err)
	}
	watermark, err := store.NewWatermarkStore(db)
	if err != nil {
		t.Fatalf("NewWatermarkStore err=%v", err)
	}
	manual, err := store.NewManualTaskStore(db)
	if err != nil {
		t.Fatalf("NewManualTaskStore err=%v", err)
	}
	orderRaw, err := store.NewDscoOrderRawStore(db)
	if err != nil {
		t.Fatalf("NewDscoOrderRawStore err=%v", err)
	}

	mock.ExpectQuery("SELECT watermark FROM job_watermark").
		WithArgs("ship_to_dsco").
		WillReturnRows(sqlmock.NewRows([]string{"watermark"}).AddRow([]byte(`{"mode":"update_time","since":1710925191}`)))
	mock.ExpectExec("INSERT INTO sync_order_state").
		WithArgs("d1", sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("UPDATE sync_order_state").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), "d1").
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec("UPDATE sync_order_state").
		WithArgs(sqlmock.AnyArg(), "TN1", sqlmock.AnyArg(), "d1").
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec("INSERT INTO job_watermark").
		WithArgs("ship_to_dsco", []byte(`{"mode":"update_time","since":1720429064}`)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	p, err := NewOrderPipeline(dscoCli, lxCli, orderState, watermark, manual, orderRaw, now, 5, "delivered_at")
	if err != nil {
		t.Fatalf("NewOrderPipeline err=%v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	t.Cleanup(cancel)

	if err := p.ShipToDSCO(ctx, 10009, "s1", 26); err != nil {
		t.Fatalf("ShipToDSCO err=%v", err)
	}
	if shipCalls != 1 {
		t.Fatalf("shipCalls=%d want=1", shipCalls)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations err=%v", err)
	}
}

func TestOrderPipeline_InvoiceToDSCO_Behavior(t *testing.T) {
	t.Parallel()

	// DSCO mock：GET /invoice -> empty, GET /order/ -> order, POST /invoice -> success
	var getInvoiceCalls, getOrderCalls, postInvoiceCalls int
	dscoSrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/v3/invoice":
			switch r.Method {
			case http.MethodGet:
				getInvoiceCalls++
				q := r.URL.Query()
				if q.Get("key") != "invoiceId" || q.Get("value") != "INV-d1" {
					t.Fatalf("query=%v want key=invoiceId value=INV-d1", q)
				}
				w.Header().Set("Content-Type", "application/json")
				_, _ = io.WriteString(w, `{"invoices":[]}`)
			case http.MethodPost:
				postInvoiceCalls++
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)
				_, _ = io.WriteString(w, `{"success":true}`)
			default:
				t.Fatalf("method=%s unexpected", r.Method)
			}
		case "/api/v3/order/":
			if r.Method != http.MethodGet {
				t.Fatalf("method=%s want=%s", r.Method, http.MethodGet)
			}
			getOrderCalls++
			q := r.URL.Query()
			if q.Get("orderKey") != "dscoOrderId" || q.Get("value") != "d1" {
				t.Fatalf("query=%v want orderKey=dscoOrderId value=d1", q)
			}
			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, `{
				"dscoOrderId":"d1",
				"currencyCode":"USD",
				"lineItems":[{"quantity":2,"sku":"SKU-1","consumerPrice":10.0}]
			}`)
		default:
			t.Fatalf("unexpected path=%s", r.URL.Path)
		}
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

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New err=%v", err)
	}
	t.Cleanup(func() { _ = db.Close() })

	orderState, err := store.NewOrderStateStore(db)
	if err != nil {
		t.Fatalf("NewOrderStateStore err=%v", err)
	}
	watermark, err := store.NewWatermarkStore(db)
	if err != nil {
		t.Fatalf("NewWatermarkStore err=%v", err)
	}
	manual, err := store.NewManualTaskStore(db)
	if err != nil {
		t.Fatalf("NewManualTaskStore err=%v", err)
	}
	orderRaw, err := store.NewDscoOrderRawStore(db)
	if err != nil {
		t.Fatalf("NewDscoOrderRawStore err=%v", err)
	}

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT dsco_order_id").
		WithArgs(sqlmock.AnyArg(), 1).
		WillReturnRows(sqlmock.NewRows([]string{"dsco_order_id"}).AddRow("d1"))
	mock.ExpectExec("UPDATE sync_order_state").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), "d1").
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	mock.ExpectExec("UPDATE sync_order_state").
		WithArgs(sqlmock.AnyArg(), "INV-d1", sqlmock.AnyArg(), "d1").
		WillReturnResult(sqlmock.NewResult(0, 1))

	now := func() time.Time { return time.Unix(1720429074, 0) }
	p, err := NewOrderPipeline(dscoCli, nil, orderState, watermark, manual, orderRaw, now, 5, "delivered_at")
	if err != nil {
		t.Fatalf("NewOrderPipeline err=%v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	t.Cleanup(cancel)

	if err := p.InvoiceToDSCO(ctx, 1); err != nil {
		t.Fatalf("InvoiceToDSCO err=%v", err)
	}
	if getInvoiceCalls != 1 || getOrderCalls != 1 || postInvoiceCalls != 1 {
		t.Fatalf("calls invoice_get=%d order_get=%d invoice_post=%d", getInvoiceCalls, getOrderCalls, postInvoiceCalls)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations err=%v", err)
	}
}

func TestOrderPipeline_InvoiceToDSCO_UsesRetailPriceFallback(t *testing.T) {
	t.Parallel()

	// DSCO mock：GET /invoice -> empty, GET /order/ -> order(retailPrice), POST /invoice -> success
	var getInvoiceCalls, getOrderCalls, postInvoiceCalls int
	dscoSrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/v3/invoice":
			switch r.Method {
			case http.MethodGet:
				getInvoiceCalls++
				w.Header().Set("Content-Type", "application/json")
				_, _ = io.WriteString(w, `{"invoices":[]}`)
			case http.MethodPost:
				postInvoiceCalls++
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)
				_, _ = io.WriteString(w, `{"success":true}`)
			default:
				t.Fatalf("method=%s unexpected", r.Method)
			}
		case "/api/v3/order/":
			if r.Method != http.MethodGet {
				t.Fatalf("method=%s want=%s", r.Method, http.MethodGet)
			}
			getOrderCalls++
			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, `{
				"dscoOrderId":"d1",
				"currencyCode":"USD",
				"lineItems":[{"quantity":2,"sku":"SKU-1","retailPrice":10.0}]
			}`)
		default:
			t.Fatalf("unexpected path=%s", r.URL.Path)
		}
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

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New err=%v", err)
	}
	t.Cleanup(func() { _ = db.Close() })

	orderState, _ := store.NewOrderStateStore(db)
	watermark, _ := store.NewWatermarkStore(db)
	manual, _ := store.NewManualTaskStore(db)
	orderRaw, _ := store.NewDscoOrderRawStore(db)

	// ClaimForInvoice：BEGIN -> SELECT -> UPDATE -> COMMIT
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT dsco_order_id").
		WithArgs(sqlmock.AnyArg(), 1).
		WillReturnRows(sqlmock.NewRows([]string{"dsco_order_id"}).AddRow("d1"))
	mock.ExpectExec("UPDATE sync_order_state").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), "d1").
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	// MarkInvoiceSuccess
	mock.ExpectExec("UPDATE sync_order_state").
		WithArgs(sqlmock.AnyArg(), "INV-d1", sqlmock.AnyArg(), "d1").
		WillReturnResult(sqlmock.NewResult(0, 1))

	now := func() time.Time { return time.Date(2025, 1, 1, 0, 10, 20, 0, time.UTC) }
	p, err := NewOrderPipeline(dscoCli, nil, orderState, watermark, manual, orderRaw, now, 5, "delivered_at")
	if err != nil {
		t.Fatalf("NewOrderPipeline err=%v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	t.Cleanup(cancel)

	if err := p.InvoiceToDSCO(ctx, 1); err != nil {
		t.Fatalf("InvoiceToDSCO err=%v", err)
	}
	if getInvoiceCalls != 1 || getOrderCalls != 1 || postInvoiceCalls != 1 {
		t.Fatalf("calls getInvoice=%d getOrder=%d postInvoice=%d", getInvoiceCalls, getOrderCalls, postInvoiceCalls)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations err=%v", err)
	}
}
