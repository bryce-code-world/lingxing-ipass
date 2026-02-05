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
		{name: "shipment_pending", in: "shipment_pending", want: 1, wantOK: true},
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

func TestShouldSkipInvoiceBecauseAlreadyHasInvoiceID(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name            string
		onlyPONumber    string
		existingInvoice string
		want            bool
	}{
		{name: "empty_all", onlyPONumber: "", existingInvoice: "", want: false},
		{name: "existing_invoice_default_skip", onlyPONumber: "", existingInvoice: "inv-1", want: true},
		{name: "existing_invoice_trim_default_skip", onlyPONumber: "  ", existingInvoice: "  inv-1  ", want: true},
		{name: "manual_force_no_skip", onlyPONumber: "PO123", existingInvoice: "inv-1", want: false},
		{name: "manual_force_trim_no_skip", onlyPONumber: "  PO123  ", existingInvoice: "  inv-1  ", want: false},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := shouldSkipInvoiceBecauseAlreadyHasInvoiceID(tc.onlyPONumber, tc.existingInvoice)
			if got != tc.want {
				t.Fatalf("shouldSkipInvoiceBecauseAlreadyHasInvoiceID(%q,%q) = %v, want %v", tc.onlyPONumber, tc.existingInvoice, got, tc.want)
			}
		})
	}
}
