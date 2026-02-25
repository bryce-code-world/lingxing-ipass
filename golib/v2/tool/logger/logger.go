package logger

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Config 表示日志配置（一期开箱即用）。
type Config struct {
	// Dir 日志目录；为空时默认 ./logs
	Dir string

	// MaxSizeMB 单个日志文件最大大小（MB）
	MaxSizeMB int
	// MaxBackups 备份文件数量
	MaxBackups int
	// MaxAgeDays 保留天数
	MaxAgeDays int
	// Compress 是否压缩
	Compress bool

	// Stdout 是否同时输出到 stdout（默认 true）
	Stdout bool
}

var (
	mu          sync.Mutex
	initialized bool

	infoLogger  *zap.Logger
	debugLogger *zap.Logger
	warnLogger  *zap.Logger
	errorLogger *zap.Logger

	infoHook  *lumberjack.Logger
	debugHook *lumberjack.Logger
	warnHook  *lumberjack.Logger
	errorHook *lumberjack.Logger
)

// Init 初始化日志（重复调用会覆盖旧实例）。
func Init(cfg Config) error {
	mu.Lock()
	defer mu.Unlock()

	dir := strings.TrimSpace(cfg.Dir)
	if dir == "" {
		dir = "./logs"
	}
	if cfg.MaxSizeMB <= 0 {
		cfg.MaxSizeMB = 128
	}
	if cfg.MaxBackups <= 0 {
		cfg.MaxBackups = 300
	}
	if cfg.MaxAgeDays <= 0 {
		cfg.MaxAgeDays = 7
	}
	// 一期默认同时输出到 stdout，便于容器采集与排障。
	cfg.Stdout = true

	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}

	closeLocked()

	infoLogger, infoHook = newZap(filepath.Join(dir, "info.log"), zapcore.InfoLevel, cfg)
	debugLogger, debugHook = newZap(filepath.Join(dir, "debug.log"), zapcore.DebugLevel, cfg)
	warnLogger, warnHook = newZap(filepath.Join(dir, "warn.log"), zapcore.WarnLevel, cfg)
	errorLogger, errorHook = newZap(filepath.Join(dir, "error.log"), zapcore.ErrorLevel, cfg)

	initialized = true
	return nil
}

// Sync 主动刷盘（程序退出前调用）。
func Sync() {
	mu.Lock()
	defer mu.Unlock()
	if infoLogger != nil {
		_ = infoLogger.Sync()
	}
	if debugLogger != nil {
		_ = debugLogger.Sync()
	}
	if warnLogger != nil {
		_ = warnLogger.Sync()
	}
	if errorLogger != nil {
		_ = errorLogger.Sync()
	}
}

// Close 关闭日志文件句柄（主要用于测试或进程优雅退出）。
func Close() {
	mu.Lock()
	defer mu.Unlock()
	closeLocked()
}

func closeLocked() {
	if infoLogger != nil {
		_ = infoLogger.Sync()
	}
	if debugLogger != nil {
		_ = debugLogger.Sync()
	}
	if warnLogger != nil {
		_ = warnLogger.Sync()
	}
	if errorLogger != nil {
		_ = errorLogger.Sync()
	}

	if infoHook != nil {
		_ = infoHook.Close()
	}
	if debugHook != nil {
		_ = debugHook.Close()
	}
	if warnHook != nil {
		_ = warnHook.Close()
	}
	if errorHook != nil {
		_ = errorHook.Close()
	}

	infoLogger, debugLogger, warnLogger, errorLogger = nil, nil, nil, nil
	infoHook, debugHook, warnHook, errorHook = nil, nil, nil, nil
}

func ensureInitialized() {
	if initialized {
		return
	}
	_ = Init(Config{Dir: "./logs", Stdout: true})
}

// Info 记录 info 级别日志（自动追加 trace_id）。
func Info(ctx context.Context, msg string, kv ...any) {
	ensureInitialized()
	logWith(ctx, infoLogger.Info, msg, kv...)
}

// Debug 记录 debug 级别日志（自动追加 trace_id）。
func Debug(ctx context.Context, msg string, kv ...any) {
	ensureInitialized()
	logWith(ctx, debugLogger.Debug, msg, kv...)
}

// Warn 记录 warn 级别日志（自动追加 trace_id）。
func Warn(ctx context.Context, msg string, kv ...any) {
	ensureInitialized()
	logWith(ctx, warnLogger.Warn, msg, kv...)
}

// Error 记录 error 级别日志（自动追加 trace_id）。
func Error(ctx context.Context, msg string, kv ...any) {
	ensureInitialized()
	logWith(ctx, errorLogger.Error, msg, kv...)
}

type zapLogFunc func(msg string, fields ...zap.Field)

func logWith(ctx context.Context, fn zapLogFunc, msg string, kv ...any) {
	if fn == nil {
		return
	}

	if tid, ok := TraceIDFromContext(ctx); ok && strings.TrimSpace(tid) != "" {
		kv = append(kv, "trace_id", tid)
	}

	fields, err := kvToFields(kv...)
	if err != nil {
		fields = append(fields, zap.String("kv_error", err.Error()))
	}
	fn(msg, fields...)
}

func kvToFields(kv ...any) ([]zap.Field, error) {
	if len(kv) == 0 {
		return nil, nil
	}
	if len(kv)%2 != 0 {
		return nil, errors.New("kv 参数必须成对出现")
	}

	out := make([]zap.Field, 0, len(kv)/2)
	for i := 0; i < len(kv); i += 2 {
		k, ok := kv[i].(string)
		if !ok || strings.TrimSpace(k) == "" {
			out = append(out, zap.Any("bad_key", kv[i]))
			continue
		}
		key := strings.TrimSpace(k)
		val := kv[i+1]
		switch v := val.(type) {
		case string:
			out = append(out, zap.String(key, v))
		case int:
			out = append(out, zap.Int(key, v))
		case int64:
			out = append(out, zap.Int64(key, v))
		case float64:
			out = append(out, zap.Float64(key, v))
		case bool:
			out = append(out, zap.Bool(key, v))
		case error:
			out = append(out, zap.String(key, v.Error()))
		case time.Time:
			out = append(out, zap.String(key, v.UTC().Format(time.RFC3339Nano)))
		default:
			out = append(out, zap.Any(key, v))
		}
	}
	return out, nil
}

func newZap(filePath string, level zapcore.Level, cfg Config) (*zap.Logger, *lumberjack.Logger) {
	hook := &lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    cfg.MaxSizeMB,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAgeDays,
		Compress:   cfg.Compress,
	}

	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	var ws []zapcore.WriteSyncer
	if cfg.Stdout {
		ws = append(ws, zapcore.AddSync(os.Stdout))
	}
	ws = append(ws, zapcore.AddSync(hook))

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(ws...),
		atomicLevel,
	)

	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.WarnLevel), zap.Fields(zap.Int("pid", os.Getpid()))), hook
}

// MarshalAny 用于把任意对象安全转为 JSON 字符串（便于写入 payload 字段）。
func MarshalAny(v any) string {
	b, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprintf(`{"marshal_error":%q}`, err.Error())
	}
	return string(b)
}
