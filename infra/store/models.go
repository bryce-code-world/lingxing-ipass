package store

import "encoding/json"

type RuntimeConfigRow struct {
	ID        int64           `gorm:"column:id;primaryKey"`
	Domain    string          `gorm:"column:domain"`
	Config    json.RawMessage `gorm:"column:config"`
	UpdatedAt int64           `gorm:"column:updated_at"`
}

func (RuntimeConfigRow) TableName() string { return "runtime_config" }

type DSCOOrderSyncRow struct {
	ID                int64           `gorm:"column:id;primaryKey"`
	PONumber          string          `gorm:"column:po_number"`
	DSCOCreateTime    int64           `gorm:"column:dsco_create_time"`
	DSCOREtailerID    string          `gorm:"column:dsco_retailer_id"`
	DSCOStatus        string          `gorm:"column:dsco_status"`
	Status            int16           `gorm:"column:status"`
	Payload           json.RawMessage `gorm:"column:payload"`
	MSKUs             PGTextArray     `gorm:"column:mskus;type:text[]"`
	WarehouseID       string          `gorm:"column:warehouse_id"`
	Shipment          string          `gorm:"column:shipment"`
	ShippedTrackingNo string          `gorm:"column:shipped_tracking_no"`
	DSCOInvoiceID     string          `gorm:"column:dsco_invoice_id"`
	CreatedAt         int64           `gorm:"column:created_at"`
	UpdatedAt         int64           `gorm:"column:updated_at"`
}

func (DSCOOrderSyncRow) TableName() string { return "dsco_order_sync" }

type DSCOWarehouseSyncRow struct {
	ID                   int64  `gorm:"column:id;primaryKey"`
	SyncTime             int64  `gorm:"column:sync_time"`
	DSCOWarehouseID      string `gorm:"column:dsco_warehouse_id"`
	DSCOWarehouseSKU     string `gorm:"column:dsco_warehouse_sku"`
	DSCOWarehouseNum     int    `gorm:"column:dsco_warehouse_num"`
	LingXingWarehouseID  string `gorm:"column:lingxing_warehouse_id"`
	LingXingWarehouseSKU string `gorm:"column:lingxing_warehouse_sku"`
	LingXingWarehouseNum int    `gorm:"column:lingxing_warehouse_num"`
	Status               int16  `gorm:"column:status"`
	Reason               string `gorm:"column:reason"`
	CreatedAt            int64  `gorm:"column:created_at"`
	UpdatedAt            int64  `gorm:"column:updated_at"`
}

func (DSCOWarehouseSyncRow) TableName() string { return "dsco_warehouse_sync" }
