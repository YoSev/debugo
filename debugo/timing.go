package debugo

import "time"

func (l *Logger) elapsed() time.Duration {
	currentTime := time.Now()
	defer func() { l.lastLog = currentTime }()

	elapsed := time.Duration(0)
	if !l.lastLog.IsZero() {
		elapsed = currentTime.Sub(l.lastLog)
	}

	return elapsed
}
