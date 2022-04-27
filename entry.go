package wlog

import (
	"bytes"
	"os"
	"runtime"
	"strings"
	"time"
)

// entry 输出模型

type Entry struct {
	logger *logger
	Buffer *bytes.Buffer
	Level  Level
	Time   time.Time
	File   string
	Line   int
	Func   string
	Format string
	Args   []interface{}
}

func entry(logger *logger) *Entry {
	return &Entry{logger: logger, Buffer: bytes.NewBuffer(nil)}
}

func (e *Entry) Write(level Level, format string, args ...interface{}) {
	// 不到等级
	if e.logger.opt.level > level {
		return
	}
	e.Time = time.Now()
	e.Level = level
	e.Format = format
	e.Args = args

	if !e.logger.opt.disableCaller {
		if pc, file, line, ok := runtime.Caller(2); !ok {
			e.File = "None"
			e.Func = "None"
		} else {
			// 获取调用堆栈的名字
			e.File, e.Line, e.Func = file, line, runtime.FuncForPC(pc).Name()
			// 获取调用堆栈的名字
			e.Func = e.Func[strings.LastIndex(e.Func, "/")+1:]
		}
	}

	e.format()
	e.writer()
	e.hooks()
	e.release()
}

func (e *Entry) format() {
	e.logger.opt.formatter.Format(e)
}

func (e *Entry) writer() {

	// 标准输出优化
	if e.logger.opt.output == os.Stderr || e.logger.opt.output == os.Stdout {

		l := e.logger.opt.logColors[e.Level].Sprintf("%s%s", e.Buffer.String(), "\n")
		e.Buffer.Reset()
		e.Buffer.WriteString(l)
	}
	e.logger.mu.Lock()
	e.logger.opt.output.Write(e.Buffer.Bytes())
	e.logger.mu.Unlock()
}

func (e *Entry) release() {
	// 回收初始化
	e.Args, e.Line, e.File, e.Format, e.Func = nil, 0, "", "", ""
	e.Buffer.Reset()
	e.logger.entryPool.Put(e)
}

func (e *Entry) hooks() {
	if e.logger.opt.hook != nil {
		e.logger.opt.hook.Handler(e.Level, e.Buffer.Bytes())
	}
}
