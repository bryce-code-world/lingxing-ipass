package dsco_lingxing

import "testing"

func TestDscoStatusToSyncStatus(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name   string
		in     string
		want   int16
		wantOK bool
	}{
		{name: "created", in: "created", want: 1, wantOK: true},
		{name: "shipment_pending", in: "shipment_pending", want: 3, wantOK: true},
		{name: "shipped", in: "shipped", want: 5, wantOK: true},
		{name: "cancelled", in: "cancelled", want: 6, wantOK: true},
		{name: "trim+case", in: "  ShIpPeD  ", want: 5, wantOK: true},
		{name: "empty", in: "", want: 0, wantOK: false},
		{name: "unknown", in: "completed", want: 0, wantOK: false},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, ok := dscoStatusToSyncStatus(tc.in)
			if ok != tc.wantOK || got != tc.want {
				t.Fatalf("dscoStatusToSyncStatus(%q) = (%d,%v), want (%d,%v)", tc.in, got, ok, tc.want, tc.wantOK)
			}
		})
	}
}
