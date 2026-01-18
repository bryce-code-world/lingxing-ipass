package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadDotEnv_Behavior(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, ".env")

	content := `
# 注释行
IPASS_A="a1"
IPASS_B=b1
IPASS_C='c1'

# 已存在的变量不应被覆盖
IPASS_KEEP="new"
`
	if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
		t.Fatalf("write .env err=%v", err)
	}

	_ = os.Setenv("IPASS_KEEP", "old")
	t.Cleanup(func() { _ = os.Unsetenv("IPASS_KEEP") })

	if err := LoadDotEnv(p); err != nil {
		t.Fatalf("LoadDotEnv err=%v", err)
	}

	if got, want := os.Getenv("IPASS_A"), "a1"; got != want {
		t.Fatalf("IPASS_A=%q want %q", got, want)
	}
	if got, want := os.Getenv("IPASS_B"), "b1"; got != want {
		t.Fatalf("IPASS_B=%q want %q", got, want)
	}
	if got, want := os.Getenv("IPASS_C"), "c1"; got != want {
		t.Fatalf("IPASS_C=%q want %q", got, want)
	}
	if got, want := os.Getenv("IPASS_KEEP"), "old"; got != want {
		t.Fatalf("IPASS_KEEP=%q want %q", got, want)
	}
}

func TestLoadDotEnv_FileNotExists(t *testing.T) {
	if err := LoadDotEnv(filepath.Join(t.TempDir(), ".env")); err != nil {
		t.Fatalf("LoadDotEnv err=%v", err)
	}
}
