package wlog

import (
	"fmt"
	"github.com/fatih/color"
	"testing"
)

func TestColor(t *testing.T) {
	DebugColor.Println("1111")
	InfoColor.Println("111")
	WarnColor.Println("111")
	ErrorColor.Println("111")
	PanicColor.Println("111")
	FatalColor.Println("111")
}

func TestColorWrite(t *testing.T) {
	sprintf := DebugColor.Color.Sprintf("%s\n", "xxx")
	fmt.Println(sprintf)
	sprintln := DebugColor.Color.Sprintln("111")
	fmt.Println(sprintln)
	sprint := InfoColor.Color.Sprint("111\n")
	fmt.Println(sprint)

	colorer := logColors[InfoLevel]
	colorer.Println(111)
	InfoColor.SetFontColor(color.FgRed)
	//colorer.SetFontColor(color.FgRed)
	colorer.Println(111)
}
