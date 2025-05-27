package debugo

import (
	"fmt"
	"time"
)

func (d *Debugger) elapsed() time.Duration {
	currentTime := time.Now()
	defer func() { d.lastLog = currentTime }()

	elapsed := time.Duration(0)
	if !d.lastLog.IsZero() {
		elapsed = currentTime.Sub(d.lastLog)
	}

	return elapsed
}

func prettyPrintDuration(d time.Duration) string {
	hours := d / time.Hour
	d %= time.Hour
	minutes := d / time.Minute
	d %= time.Minute
	seconds := d / time.Second
	d %= time.Second
	milliseconds := d / time.Millisecond

	result := ""
	if hours > 0 {
		result += fmt.Sprintf("%dh", hours)
	}
	if minutes > 0 || hours > 0 {
		result += fmt.Sprintf("%dm", minutes)
	}
	if seconds > 0 || minutes > 0 || hours > 0 {
		result += fmt.Sprintf("%ds", seconds)
	}

	result += fmt.Sprintf("%dms", milliseconds) // Always include milliseconds

	return result
}
