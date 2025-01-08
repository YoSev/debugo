package debugo

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

var colorCodes = map[string]string{
	// Standard Colors
	"red":    "\033[31m",
	"green":  "\033[32m",
	"yellow": "\033[33m",
	"blue":   "\033[34m",
	"purple": "\033[35m",
	"cyan":   "\033[36m",

	// Bright Colors
	"brightBlack":  "\033[90m",
	"brightRed":    "\033[91m",
	"brightGreen":  "\033[92m",
	"brightYellow": "\033[93m",
	"brightBlue":   "\033[94m",
	"brightPurple": "\033[95m",
	"brightCyan":   "\033[96m",

	// 256-Color Palette (Foreground Colors)
	"color8":  "\033[38;5;8m",  // Gray
	"color9":  "\033[38;5;9m",  // Bright Red
	"color10": "\033[38;5;10m", // Bright Green
	"color11": "\033[38;5;11m", // Bright Yellow
	"color12": "\033[38;5;12m", // Bright Blue
	"color13": "\033[38;5;13m", // Bright Magenta
	"color14": "\033[38;5;14m", // Bright Cyan

	// Additional Colors (256-Color Codes)
	"color17": "\033[38;5;17m", // Dark Blue
	"color18": "\033[38;5;18m", // Dark Green
	"color19": "\033[38;5;19m", // Teal
	"color20": "\033[38;5;20m", // Navy Blue
	"color21": "\033[38;5;21m", // Deep Blue
	"color22": "\033[38;5;22m", // Olive
	"color23": "\033[38;5;23m", // Sea Green
	"color24": "\033[38;5;24m", // Turquoise
	"color25": "\033[38;5;25m", // Sky Blue

	// Extended Colors
	"color26": "\033[38;5;26m", // Indigo
	"color27": "\033[38;5;27m", // Light Green
	"color28": "\033[38;5;28m", // Lime
	"color29": "\033[38;5;29m", // Forest Green
	"color30": "\033[38;5;30m", // Spring Green
	"color31": "\033[38;5;31m", // Aqua
	"color32": "\033[38;5;32m", // Cyan Green
	"color33": "\033[38;5;33m", // Aqua Blue
	"color34": "\033[38;5;34m", // Steel Blue
	"color35": "\033[38;5;35m", // Dodger Blue

	// Vibrant Colors
	"color36": "\033[38;5;36m", // Royal Blue
	"color37": "\033[38;5;37m", // Violet
	"color38": "\033[38;5;38m", // Plum
	"color39": "\033[38;5;39m", // Orchid
	"color40": "\033[38;5;40m", // Lavender
	"color41": "\033[38;5;41m", // Pink
	"color42": "\033[38;5;42m", // Salmon
	"color43": "\033[38;5;43m", // Coral
	"color44": "\033[38;5;44m", // Orange
	"color45": "\033[38;5;45m", // Gold

	// Final Set
	"color46": "\033[38;5;46m", // Yellow Green
	"color47": "\033[38;5;47m", // Olive Drab
	"color48": "\033[38;5;48m", // Light Goldenrod
	"color49": "\033[38;5;49m", // Golden Yellow
	"color50": "\033[38;5;50m", // Light Pink
}

var usedColors = []string{}

func (l *Logger) applyColor(text string) string {
	colorCode, ok := colorCodes[l.color]
	if !ok {
		colorCode = colorCodes["reset"]
	}

	return fmt.Sprintf("%s%s%s", colorCode, text, colorCodes["reset"])
}

func resetColor(str string) string {
	return fmt.Sprintf("\033[0m%s", str)
}

func (l *Logger) getNextColor() string {
	rand.Seed(uint64(time.Now().UnixNano()))

	colorNames := []string{}
	for color := range colorCodes {
		colorNames = append(colorNames, color)
	}

	if len(usedColors) == len(colorNames) {
		usedColors = nil
	}

	var randomColor string
	for {
		randomColor = colorNames[rand.Intn(len(colorNames))]
		if !containsColor(usedColors, randomColor) {
			break
		}
	}

	usedColors = append(usedColors, randomColor)

	return randomColor
}

func containsColor(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}
