package store

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func newMockGormDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	t.Helper()

	rawDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New err=%v", err)
	}
	t.Cleanup(func() { _ = rawDB.Close() })

	gdb, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      rawDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("gorm.Open err=%v", err)
	}

	return gdb, mock
}

