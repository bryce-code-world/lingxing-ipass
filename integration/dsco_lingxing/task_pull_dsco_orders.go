package dsco_lingxing

import (
	"context"
	"encoding/json"
	"time"

	"example.com/lingxing/golib/v2/sdk/dsco"
	"example.com/lingxing/golib/v2/tool/logger"

	"lingxingipass/infra/store"
	"lingxingipass/integration"
)

func (d *Domain) PullDSCOOrders(ctx integration.TaskContext) error {
	taskCtx := ctx.Ctx
	if taskCtx == nil {
		taskCtx = context.Background()
	}

	cli, err := d.dscoClient()
	if err != nil {
		return err
	}

	var (
		since  time.Time
		until  time.Time
		status int16 = 1
	)

	if ctx.Trigger == integration.TriggerManual && ctx.Override != nil {
		ov, ok := ctx.Override.(integration.PullDSCOOrdersOverride)
		if !ok {
			if p, ok2 := ctx.Override.(*integration.PullDSCOOrdersOverride); ok2 && p != nil {
				ov = *p
				ok = true
			}
		}
		if ok {
			since = time.Unix(ov.Start, 0).UTC()
			until = time.Unix(ov.End, 0).UTC()
			status = ov.Status
		}
	}

	// Default: incremental cursor from dsco_order_sync.dsco_create_time.
	if since.IsZero() || until.IsZero() {
		maxTime, ok, err := d.orderStore.GetMaxDSCOCreateTime(taskCtx)
		if err != nil {
			return err
		}
		if ok {
			since = time.Unix(maxTime, 0).UTC()
		} else {
			since = time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
		}

		until = time.Now().UTC().Add(-10 * time.Second)
	}
	q := dsco.OrderPageQuery{
		OrdersCreatedSince: since.Format(time.RFC3339),
		Until:              until.Format(time.RFC3339),
		OrdersPerPage:      ctx.Size,
	}

	var pulled int
	var scroll string
	for {
		if scroll != "" {
			q.ScrollID = scroll
		}
		resp, err := cli.Order.GetPageRaw(taskCtx, q)
		if err != nil {
			return err
		}
		scroll = resp.ScrollID
		if len(resp.Orders) == 0 {
			break
		}
		for _, raw := range resp.Orders {
			order, err := decodeDSCOOrder(raw)
			if err != nil {
				logger.Warn(taskCtx, "decode dsco order failed", "err", err)
				continue
			}
			createStr := ""
			if order.DscoCreateDate != nil {
				createStr = *order.DscoCreateDate
			} else if order.RetailerCreateDate != nil {
				createStr = *order.RetailerCreateDate
			}
			createUnix, err := parseRFC3339ToUnixSec(createStr)
			if err != nil {
				logger.Warn(taskCtx, "parse dsco create time failed", "err", err)
				continue
			}

			// Enforce [start, end) boundary for manual pull (in UTC seconds).
			if ctx.Trigger == integration.TriggerManual && !since.IsZero() && !until.IsZero() {
				if createUnix < since.Unix() || createUnix >= until.Unix() {
					continue
				}
			}

			mskus := make([]string, 0, len(order.LineItems))
			for _, li := range order.LineItems {
				if li.PartnerSKU != nil && *li.PartnerSKU != "" {
					mskus = append(mskus, *li.PartnerSKU)
				} else if li.SKU != nil && *li.SKU != "" {
					mskus = append(mskus, *li.SKU)
				}
			}

			row := store.DSCOOrderSyncRow{
				PONumber:            order.PoNumber,
				DSCOOrderID:         order.DscoOrderID,
				ConsumerOrderNumber: derefString(order.ConsumerOrderNumber),
				Channel:             derefString(order.Channel),
				DSCOCreateTime:      createUnix,
				Status:              status,
				Payload:             json.RawMessage(raw),
				MSKUs:               mskus,
				WarehouseID:         getDSCOWarehouseCode(order),
				Shipment:            getDSCOShipMethod(order),
			}
			if err := d.orderStore.Upsert(taskCtx, row); err != nil {
				logger.Warn(taskCtx, "upsert dsco_order_sync failed", "err", err)
				continue
			}
			pulled++
		}

		// Safety break: if scrollId empty, DSCO has no more pages.
		if scroll == "" {
			break
		}
	}

	logger.Info(taskCtx, "pull dsco orders done", "count", pulled)
	return nil
}

func derefString(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}
