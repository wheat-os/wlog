package wlog

import (
	"github.com/fatih/color"
)

type Attribute color.Attribute

// Base attributes
const (
	Reset Attribute = iota
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkRapid
	ReverseVideo
	Concealed
	CrossedOut
)

// Foreground text colors
const (
	FgBlack Attribute = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

// Foreground Hi-Intensity text colors
const (
	FgHiBlack Attribute = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

// Background text colors
const (
	BgBlack Attribute = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

// Background Hi-Intensity text colors
const (
	BgHiBlack Attribute = iota + 100
	BgHiRed
	BgHiGreen
	BgHiYellow
	BgHiBlue
	BgHiMagenta
	BgHiCyan
	BgHiWhite
)

type logColor struct {
	level Level
	*color.Color
}

var (
	debugColor = newColor(FgCyan, DebugLevel)
	infoColor  = newColor(FgGreen, InfoLevel)
	warnColor  = newColor(FgYellow, WarnLevel)
	errorColor = newColor(FgRed, ErrorLevel)
	panicColor = newColor(FgMagenta, PanicLevel)
	fatalColor = newColor(FgHiBlack, FatalLevel)

	logColors = map[Level]*logColor{
		DebugLevel: debugColor,
		InfoLevel:  infoColor,
		WarnLevel:  warnColor,
		ErrorLevel: errorColor,
		PanicLevel: panicColor,
		FatalLevel: fatalColor,
	}
)

func newColor(font Attribute, leve Level) *logColor {
	c := new(logColor)
	c.level = leve
	c.Color = color.New(color.Attribute(font))
	c.EnableColor()
	return c
}

func WithLogLevelColor(level Level, font Attribute) OptionFunc {
	return func(o *options) {
		o.logColors[level].Add(color.Attribute(font))
	}
}
