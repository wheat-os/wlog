package example

import (
	"os"
	"testing"

	"github.com/wheat-os/wlog"
)

func TestExample(t *testing.T) {
	log := wlog.New()

	file, err := os.OpenFile("log.text", os.O_WRONLY|os.O_CREATE, 0775)
	if err != nil {
		return
	}

	log.SetOptions(wlog.WithOutput(file))

	log.Info("123", 123, map[string]int{"qqq": 111})
	log.Debug("123", 123, map[string]int{"qqq": 111})

}
