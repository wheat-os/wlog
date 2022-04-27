package wlog

import (
	"testing"
)

// 普通输出和日志输出
func TestColor(t *testing.T) {
	debugColor.Println("1111")
	infoColor.Println("111")
	Info("dwdsadw")
}

func TestChangeLogColor(t *testing.T) {
	Info("11111")
	SetStdOptions(SetLogLevelColor(InfoLevel, FgHiYellow))

	Info("1111")
}
