package debugo

import (
	"time"

	"github.com/fatih/color"
	"golang.org/x/exp/rand"
)

var noColor = color.New(color.Reset)

var fgColors = []color.Attribute{
	color.FgRed,
	color.FgGreen,
	color.FgYellow,
	color.FgBlue,
	color.FgMagenta,
	color.FgCyan,

	color.FgHiBlack,
	color.FgHiRed,
	color.FgHiGreen,
	color.FgHiYellow,
	color.FgHiBlue,
	color.FgHiMagenta,
	color.FgHiCyan,
	color.FgHiWhite,
}

var bgColors = []color.Attribute{
	color.BgBlack,
	color.BgRed,
	color.BgGreen,
	color.BgYellow,
	color.BgBlue,
	color.BgMagenta,
	color.BgCyan,
	color.BgWhite,

	color.BgHiBlack,
	color.BgHiRed,
	color.BgHiGreen,
	color.BgHiYellow,
	color.BgHiBlue,
	color.BgHiMagenta,
	color.BgHiCyan,
	color.BgHiWhite,
}

// cache bgColors by namespace
var fgColorMap = make(map[string]*color.Color)

// cache fgColors by namespace
var bgColorMap = make(map[string]*color.Color)

func (l *Logger) setRandomColor(useBg bool) {
	rand.Seed(uint64(time.Now().UnixNano()))

	if useBg {
		if color, ok := bgColorMap[l.namespace]; ok {
			l.color = color
			return
		}

		l.color = color.New(bgColors[rand.Intn(len(bgColors))])
		bgColorMap[l.namespace] = l.color
	} else {
		if color, ok := fgColorMap[l.namespace]; ok {
			l.color = color
			return
		}

		l.color = color.New(fgColors[rand.Intn(len(fgColors))])
		fgColorMap[l.namespace] = l.color
	}
}
