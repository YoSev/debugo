package debug

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

var (
	fgColorMap = make(map[string]*color.Color)
	fgColorMux sync.RWMutex
)

func getRandomColor(namespace string) *color.Color {
	fgColorMux.Lock()
	defer fgColorMux.Unlock()

	if c, ok := fgColorMap[namespace]; ok {
		return c
	}

	src := rand.NewSource(uint64(time.Now().UnixNano()))
	r := rand.New(src)
	c := color.New(fgColors[r.Intn(len(fgColors))])
	fgColorMap[namespace] = c
	return c
}
