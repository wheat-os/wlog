package wlog

import (
	"fmt"
	"path"
)

// 格式化编码器
type Formatter interface {
	Format(entry *Entry) error
}

type TextFormatter struct {
	// 忽略基本文件名输出
	ignoreBasicFields bool
}

const (
	formatTime = "2006-01-02 15:04:05"
)

func (t *TextFormatter) Format(entry *Entry) error {
	if !t.ignoreBasicFields {
		// 写入时间 level 信息
		entry.Buffer.WriteString(fmt.Sprintf("%s %s", entry.Time.Format(formatTime), LevelNameMapping[entry.Level]))

		if entry.File != "" {
			// 获取文件名
			short := path.Base(entry.File)
			entry.Buffer.WriteString(fmt.Sprintf(" %s:%d", short, entry.Line))
		}

		entry.Buffer.WriteString("\n")
	}

	switch entry.Format {
	// 无特殊输出，采用 %v
	case FmtEmptySeparate:
		entry.Buffer.WriteString(fmt.Sprint(entry.Args...))
	default:
		entry.Buffer.WriteString(fmt.Sprintf(entry.Format, entry.Args...))
	}
	return nil
}
