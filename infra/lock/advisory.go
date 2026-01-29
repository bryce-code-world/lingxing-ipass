package lock

import (
	"context"
	"database/sql"
	"errors"
)

var ErrLocked = errors.New("locked")

type Advisory struct {
	db *sql.DB
}

func NewAdvisory(db *sql.DB) *Advisory {
	return &Advisory{db: db}
}

// TryLock holds the lock on a dedicated session (sql.Conn). Caller must Release.
func (a *Advisory) TryLock(ctx context.Context, key int64) (*sql.Conn, error) {
	if a == nil || a.db == nil {
		return nil, errors.New("db 不能为空")
	}
	conn, err := a.db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	var ok bool
	if err := conn.QueryRowContext(ctx, `SELECT pg_try_advisory_lock($1)`, key).Scan(&ok); err != nil {
		_ = conn.Close()
		return nil, err
	}
	if !ok {
		_ = conn.Close()
		return nil, ErrLocked
	}
	return conn, nil
}

func (a *Advisory) Unlock(ctx context.Context, conn *sql.Conn, key int64) error {
	if conn == nil {
		return nil
	}
	defer conn.Close()
	_, _ = conn.ExecContext(ctx, `SELECT pg_advisory_unlock($1)`, key)
	return nil
}
