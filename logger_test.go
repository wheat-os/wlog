package wlog

import (
	"errors"
	"testing"
)

func Test_logger_Debug(t *testing.T) {
	Debug("111", 222, []int{1, 2})
	Info(111.22)

	log := New()
	log.Info(1111)
}

func Test_logger(t *testing.T) {
	logger := New(WithFormatter(JsonFormatter(false)))
	logger.Debug("111", "222", errors.New("err is open"))
	logger.SetOptions(WithFormatter(JsonFormatter(true)))
	logger.SetOptions(WithDisPlayLevel(ErrorLevel))

	logger.Info("错误")
	logger.Error("发生错误")
	logger.Error("发生错误2")
}
