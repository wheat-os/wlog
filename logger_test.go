package wlog

import (
	"errors"
	"fmt"
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

func Test_Color(t *testing.T) {

	userFormat := "%s %s %d"
	userArgs := []interface{}{"不好", "失败", 30}

	colorFormat := "\033[1;31;40m%s\033[0m\n"
	colorFormat = fmt.Sprintf(colorFormat, userFormat)

	fmt.Printf(colorFormat, userArgs...)
}
