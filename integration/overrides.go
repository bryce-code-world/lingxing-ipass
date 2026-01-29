package integration

// PullDSCOOrdersOverride defines parameters for the manual "pull DSCO orders" operation.
//
// Time range is defined as [Start, End) in UTC, in Unix seconds.
// Status is the storage status to set when upserting into dsco_order_sync (1~5).
type PullDSCOOrdersOverride struct {
	Start  int64 `json:"start"`
	End    int64 `json:"end"`
	Status int16 `json:"status"`
}
