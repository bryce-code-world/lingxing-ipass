package integration

// PullDSCOOrdersOverride defines parameters for the manual "pull DSCO orders" operation.
//
// Time range is defined as [Start, End) in UTC, in Unix seconds.
type PullDSCOOrdersOverride struct {
	Start int64 `json:"start"`
	End   int64 `json:"end"`
}
