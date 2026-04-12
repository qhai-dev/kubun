package log

import (
	"github.com/qhai-dev/kaka/component/log/internal/setverbositylevel"
	"k8s.io/klog/v2"
)

func init() {
	log, control := NewZapJSONLogger(0)
	if control.SetVerbosityLevel != nil {
		setverbositylevel.Mutex.Lock()
		defer setverbositylevel.Mutex.Unlock()
		setverbositylevel.Callbacks = append(setverbositylevel.Callbacks, control.SetVerbosityLevel)
	}
	opts := []klog.LoggerOption{
		klog.FlushLogger(control.Flush),
	}
	// if writer, ok := log.GetSink().(textlogger.KlogBufferWriter); ok {
	// 	opts = append(opts, klog.WriteKlogBuffer(writer.WriteKlogBuffer))
	// }
	klog.SetLoggerWithOptions(log, opts...)
}

func InitLog() {
	klog.EnableContextualLogging(true)
}

func FlushLog() {
	klog.Flush()
}
