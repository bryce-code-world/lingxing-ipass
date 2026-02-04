package admin

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestWriteCSVExportFile_CreatesFinalCSV(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	got, err := writeCSVExportFile(dir, "dsco_order_sync", func(w *csv.Writer) error {
		if err := w.Write([]string{"id", "po_number"}); err != nil {
			return err
		}
		return w.Write([]string{"1", "PO123"})
	})
	if err != nil {
		t.Fatalf("writeCSVExportFile() err = %v", err)
	}
	if !strings.HasSuffix(strings.ToLower(got), ".csv") {
		t.Fatalf("expected .csv file, got %q", got)
	}
	if filepath.Dir(got) != dir {
		t.Fatalf("expected file under %q, got %q", dir, got)
	}
	if _, err := os.Stat(got); err != nil {
		t.Fatalf("stat %q err = %v", got, err)
	}
}

func TestFormatUnixSecToDisplay_UsesLocation(t *testing.T) {
	t.Parallel()

	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		t.Fatalf("LoadLocation err = %v", err)
	}

	// 2024-01-01 00:00:00 UTC
	got := formatUnixSecToDisplay(1704067200, loc)
	if got != "2024-01-01 08:00:00.000" {
		t.Fatalf("unexpected formatted time: %q", got)
	}
}
