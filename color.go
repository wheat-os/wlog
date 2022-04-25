package wlog

import "github.com/fatih/color"

type logColor struct {
	level Level
	*color.Color
}

var (
	DebugColor = newColor(color.BgBlack, color.FgCyan, DebugLevel)
	InfoColor  = newColor(color.BgBlack, color.FgGreen, InfoLevel)
	WarnColor  = newColor(color.BgBlack, color.FgYellow, WarnLevel)
	ErrorColor = newColor(color.BgBlack, color.FgRed, ErrorLevel)
	PanicColor = newColor(color.BgBlack, color.FgMagenta, PanicLevel)
	FatalColor = newColor(color.BgBlack, color.FgHiBlack, FatalLevel)

	logColors = map[Level]*logColor{
		DebugLevel: DebugColor,
		InfoLevel:  InfoColor,
		WarnLevel:  WarnColor,
		ErrorLevel: ErrorColor,
		PanicLevel: PanicColor,
		FatalLevel: FatalColor,
	}
)

func newColor(background color.Attribute, font color.Attribute, leve Level) *logColor {
	c := new(logColor)
	c.level = leve
	c.Color = color.New(background, font)
	c.EnableColor()
	return c
}

func (l *logColor) SetBackgroundColor(background color.Attribute) {
	l.Add(background)
}

func (l *logColor) SetFontColor(font color.Attribute) {
	l.Add(font)
}

func (l *logColor) SetOtherAttribute(attributes ...color.Attribute) {
	l.Add(attributes...)
}
