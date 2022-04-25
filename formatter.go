package wlog

import (
	"fmt"
	"os"
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
		s := ""
		if entry.logger.opt.output == os.Stderr || entry.logger.opt.output == os.Stdout {
			s = logColors[entry.Level].Sprintf("%s %s", entry.Time.Format(formatTime), LevelNameMapping[entry.Level])
		} else {
			s = fmt.Sprintf("%s %s", entry.Time.Format(formatTime), LevelNameMapping[entry.Level])
		}
		entry.Buffer.WriteString(s)
		//entry.Buffer.WriteString()

		if entry.File != "" {
			// 获取文件名
			short := path.Base(entry.File)
			s := ""
			if entry.logger.opt.output == os.Stderr || entry.logger.opt.output == os.Stdout {
				s = logColors[entry.Level].Sprintf("%s %d", short, entry.Line)
			} else {
				s = fmt.Sprintf(" %s:%d", short, entry.Line)
			}
			//entry.Buffer.WriteString(fmt.Sprintf(" %s:%d", short, entry.Line))
			entry.Buffer.WriteString(s)
		}

		entry.Buffer.WriteString("\n")
	}

	switch entry.Format {
	// 无特殊输出，采用 %v
	case FmtEmptySeparate:
		s := ""
		if entry.logger.opt.output == os.Stderr || entry.logger.opt.output == os.Stdout {
			s = logColors[entry.Level].Sprint(entry.Args...)
		} else {
			s = fmt.Sprint(entry.Args...)
		}
		//entry.Buffer.WriteString(fmt.Sprint(entry.Args...))
		entry.Buffer.WriteString(s)
	default:
		//entry.Buffer.WriteString(fmt.Sprintf(entry.Format, entry.Args...))
		s := ""
		if entry.logger.opt.output == os.Stderr || entry.logger.opt.output == os.Stdout {
			s = logColors[entry.Level].Sprintf(entry.Format, entry.Args...)
		} else {
			s = fmt.Sprintf(entry.Format, entry.Args...)
		}
		entry.Buffer.WriteString(s)
	}
	return nil
}
