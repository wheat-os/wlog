package wlog

import (
	"testing"
)

func Test_logger_Debug(t *testing.T) {
	Debug("111", 222, []int{1, 2})
	Info(111.22)

	log := New()
	log.Info(1111)
}
