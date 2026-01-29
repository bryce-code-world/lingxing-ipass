package dsco_lingxing

import (
	"net/http"
	"time"

	"lingxingipass/infra/config"
	"lingxingipass/infra/store"
)

type Domain struct {
	env config.EnvConfig

	orderStore     *store.DSCOOrderSyncStore
	warehouseStore *store.DSCOWarehouseSyncStore

	httpClient *http.Client
}

func NewDomain(env config.EnvConfig, orderStore *store.DSCOOrderSyncStore, warehouseStore *store.DSCOWarehouseSyncStore) *Domain {
	return &Domain{
		env:            env,
		orderStore:     orderStore,
		warehouseStore: warehouseStore,
		httpClient:     &http.Client{Timeout: 30 * time.Second},
	}
}
