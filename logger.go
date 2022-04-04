package wlog

import (
	"fmt"
	"io"
	"os"
	"sync"
	"unsafe"
)

// 默认 std
var std = New()

type logger struct {
	opt       *options
	mu        sync.Mutex
	entryPool *sync.Pool
}

// 设置 option
func (l *logger) SetOptions(opts ...OptionFunc) {
	l.mu.Lock()
	defer l.mu.Unlock()

	for _, optFunc := range opts {
		optFunc(l.opt)
	}
}

func New(opts ...OptionFunc) *logger {
	logger := &logger{opt: initOptions(opts...)}

	logger.entryPool = &sync.Pool{New: func() interface{} {
		return entry(logger)
	}}

	return logger
}

func StdLogger() *logger {
	return std
}

// 设置默认的 std 的日志输出器
func SetStdOptions(opts ...OptionFunc) {
	std.SetOptions(opts...)
}

func Writer() io.Writer {
	return std
}

//  获取对对象池里的对象
func (l *logger) entry() *Entry {
	return l.entryPool.Get().(*Entry)
}

// 输出编码后字节内存的 string  形式
func (l *logger) Write(data []byte) (int, error) {
	l.entry().Write(l.opt.level, FmtEmptySeparate, *(*string)(unsafe.Pointer(&data)))
	return 0, nil
}

func (l *logger) Debug(args ...interface{}) {
	l.entry().Write(DebugLevel, FmtEmptySeparate, args...)
}

func (l *logger) Info(args ...interface{}) {
	l.entry().Write(InfoLevel, FmtEmptySeparate, args...)
}

func (l *logger) Warn(args ...interface{}) {
	l.entry().Write(WarnLevel, FmtEmptySeparate, args...)
}

func (l *logger) Error(args ...interface{}) {
	l.entry().Write(ErrorLevel, FmtEmptySeparate, args...)
}

func (l *logger) Panic(args ...interface{}) {
	l.entry().Write(PanicLevel, FmtEmptySeparate, args...)
	panic(fmt.Sprint(args...))
}

func (l *logger) Fatal(args ...interface{}) {
	l.entry().Write(FatalLevel, FmtEmptySeparate, args...)
	os.Exit(1)
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.entry().Write(DebugLevel, format, args...)
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.entry().Write(InfoLevel, format, args...)
}

func (l *logger) Warnf(format string, args ...interface{}) {
	l.entry().Write(WarnLevel, format, args...)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.entry().Write(ErrorLevel, format, args...)
}

func (l *logger) Panicf(format string, args ...interface{}) {
	l.entry().Write(PanicLevel, format, args...)
	panic(fmt.Sprintf(format, args...))
}

func (l *logger) Fatalf(format string, args ...interface{}) {
	l.entry().Write(FatalLevel, format, args...)
	os.Exit(1)
}

// std logger
func Debug(args ...interface{}) {
	std.entry().Write(DebugLevel, FmtEmptySeparate, args...)
}

func Info(args ...interface{}) {
	std.entry().Write(InfoLevel, FmtEmptySeparate, args...)
}

func Warn(args ...interface{}) {
	std.entry().Write(WarnLevel, FmtEmptySeparate, args...)
}

func Error(args ...interface{}) {
	std.entry().Write(ErrorLevel, FmtEmptySeparate, args...)
}

func Panic(args ...interface{}) {
	std.entry().Write(PanicLevel, FmtEmptySeparate, args...)
	panic(fmt.Sprint(args...))
}

func Fatal(args ...interface{}) {
	std.entry().Write(FatalLevel, FmtEmptySeparate, args...)
	os.Exit(1)
}

func Debugf(format string, args ...interface{}) {
	std.entry().Write(DebugLevel, format, args...)
}

func Infof(format string, args ...interface{}) {
	std.entry().Write(InfoLevel, format, args...)
}

func Warnf(format string, args ...interface{}) {
	std.entry().Write(WarnLevel, format, args...)
}

func Errorf(format string, args ...interface{}) {
	std.entry().Write(ErrorLevel, format, args...)
}

func Panicf(format string, args ...interface{}) {
	std.entry().Write(PanicLevel, format, args...)
	panic(fmt.Sprintf(format, args...))
}

func Fatalf(format string, args ...interface{}) {
	std.entry().Write(FatalLevel, format, args...)
	os.Exit(1)
}
