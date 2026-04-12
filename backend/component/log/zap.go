package log

import (
	"io"
	"os"
	"sync/atomic"
	"time"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type runtime struct {
	v uint32
}

func (r *runtime) ZapV() zapcore.Level {
	return -zapcore.Level(atomic.LoadUint32(&r.v))
}

func (r *runtime) Enabled(level zapcore.Level) bool {
	return level >= r.ZapV()
}

func (r *runtime) SetVerbosityLevel(v uint32) error {
	atomic.StoreUint32(&r.v, v)
	return nil
}

type RuntimeControl struct {
	SetVerbosityLevel func(uint32) error
	Flush             func()
}

var _ zapcore.LevelEnabler = &runtime{}

func NewZapJSONLogger(v uint32) (logr.Logger, RuntimeControl) {
	r := &runtime{
		v: v,
	}

	encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey: "msg",
		CallerKey:  "caller",
		NameKey:    "logger",
		TimeKey:    "ts",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})

	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel && r.Enabled(lvl)
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel && r.Enabled(lvl)
	})
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.Lock(AddNopSync(os.Stderr)), highPriority),
		zapcore.NewCore(encoder, zapcore.Lock(AddNopSync(os.Stdout)), lowPriority),
	)

	l := zap.New(core, zap.WithCaller(true))

	return zapr.NewLoggerWithOptions(l, zapr.LogInfoLevel("v"), zapr.ErrorKey("err")),
		RuntimeControl{
			SetVerbosityLevel: r.SetVerbosityLevel,
			Flush: func() {
				_ = l.Sync()
			},
		}
}

func AddNopSync(writer io.Writer) zapcore.WriteSyncer {
	return nopSync{
		Writer: writer,
	}
}

type nopSync struct {
	io.Writer
}

func (f nopSync) Sync() error {
	return nil
}
