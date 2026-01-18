package sync

import (
	"os"
	"testing"

	"example.com/lingxing/golib/v2/tool/logger"
)

func TestMain(m *testing.M) {
	dir, err := os.MkdirTemp("", "lingxingipass-logs-*")
	if err == nil {
		_ = logger.Init(logger.Config{Dir: dir})
	}

	code := m.Run()

	logger.Close()
	if dir != "" {
		_ = os.RemoveAll(dir)
	}
	os.Exit(code)
}
