package debugo

import (
	"sync"
	"time"

	"github.com/fatih/color"
	"golang.org/x/exp/rand"
)

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

var fgColorMap = sync.Map{}

func getRandomColor(namespace string) *color.Color {
	if c, ok := fgColorMap.Load(namespace); ok {
		return c.(*color.Color)
	}

	src := rand.NewSource(uint64(time.Now().UnixNano()))
	r := rand.New(src)
	c := color.New(fgColors[r.Intn(len(fgColors))])
	fgColorMap.Store(namespace, c)
	return c
}
