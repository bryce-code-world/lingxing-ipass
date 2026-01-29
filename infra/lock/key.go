package lock

import (
	"hash/fnv"
)

// KeyFromString returns a stable int64 key (PostgreSQL advisory lock key).
// It must be stable across processes and machines.
func KeyFromString(s string) int64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(s))
	return int64(h.Sum64())
}
