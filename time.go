package debug

import "time"

func (d *Debugger) elapsed() time.Duration {
	currentTime := time.Now()
	defer func() { d.lastLog = currentTime }()

	elapsed := time.Duration(0)
	if !d.lastLog.IsZero() {
		elapsed = currentTime.Sub(d.lastLog)
	}

	return elapsed
}
