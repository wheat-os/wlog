package wlog

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
)

const (
	FmtEmptySeparate = ""
)

type Level uint8

const (
	DebugLevel = 1 << 0
	InfoLevel  = 1 << 1
	WarnLevel  = 1 << 2
	ErrorLevel = 1 << 3
	PanicLevel = 1 << 4
	FatalLevel = 1 << 5
)

var LevelNameMapping = map[Level]string{
	DebugLevel: "DEBUG",
	InfoLevel:  "INFO",
	WarnLevel:  "WARN",
	ErrorLevel: "ERROR",
	PanicLevel: "PANIC",
	FatalLevel: "FATAL",
}

var errUnmarshalNilLevel = errors.New("can't unmarshal a nil *Level")

func (l *Level) unmarshalText(text []byte) bool {
	switch string(text) {
	case "debug", "DEBUG":
		*l = DebugLevel
	case "info", "INFO", "": // make the zero value useful
		*l = InfoLevel
	case "warn", "WARN":
		*l = WarnLevel
	case "error", "ERROR":
		*l = ErrorLevel
	case "panic", "PANIC":
		*l = PanicLevel
	case "fatal", "FATAL":
		*l = FatalLevel
	default:
		return false
	}
	return true
}

func (l *Level) UnmarshalText(text []byte) error {
	if l == nil {
		return errUnmarshalNilLevel
	}
	if !l.unmarshalText(text) && !l.unmarshalText(bytes.ToLower(text)) {
		return fmt.Errorf("unrecognized level: %q", text)
	}
	return nil
}

type options struct {
	output    io.Writer
	level     Level
	stdLevel  Level
	formatter Formatter

	// 打开堆栈打印
	disableCaller bool

	// 执行钩子
	hook Hook

	// 日志颜色等级
	logColors map[Level]*logColor
}

type OptionFunc func(*options)

func initOptions(opts ...OptionFunc) *options {
	o := &options{}

	for _, opt := range opts {
		// set log option
		opt(o)
	}

	if o.output == nil {
		o.output = os.Stderr
	}

	if o.formatter == nil {
		o.formatter = &TextFormatter{}
	}

	if o.logColors == nil {
		o.logColors = logColors
	}

	return o
}

func WithDisPlayLevel(level Level) OptionFunc {
	return func(o *options) {
		o.level = level
	}
}

func WithOutput(op io.Writer) OptionFunc {
	return func(o *options) {
		o.output = op
	}
}

func WithStdDisPlayLevel(level Level) OptionFunc {
	return func(o *options) {
		o.stdLevel = level
	}
}

func WithFormatter(formatter Formatter) OptionFunc {
	return func(o *options) {
		o.formatter = formatter
	}
}

// 关闭堆栈调用
func WithDisableCaller(caller bool) OptionFunc {
	return func(o *options) {
		o.disableCaller = caller
	}
}
