package api

import (
	"github.com/gin-gonic/gin"

	"lingxingipass/infra/config"
	"lingxingipass/infra/store"
	"lingxingipass/integration"
)

func Register(engine *gin.Engine, env config.EnvConfig, runner *integration.Runner, warehouseStore *store.DSCOWarehouseSyncStore) {
	if engine == nil {
		return
	}
	if !env.API.Enable {
		return
	}

	v1 := engine.Group("/api/v1")
	v1.Use(bearerAuth(env.API.Token))

	inv := &inventoryAPI{
		runner:         runner,
		warehouseStore: warehouseStore,
	}
	v1.GET("/inventory/diff", inv.getDiff)
	v1.POST("/inventory/sync", inv.postSync)
}
