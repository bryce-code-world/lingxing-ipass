module lingxingipass

go 1.22.8

require (
	example.com/lingxing/golib/v2 v2.0.0
	github.com/DATA-DOG/go-sqlmock v1.5.2
	github.com/go-sql-driver/mysql v1.9.0
)

replace example.com/lingxing/golib/v2 => ./golib/v2

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
)
