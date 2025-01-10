package debugo

func (l *Debugger) applyOptions() {
	if l.options != nil {
		if l.options.ForceEnable {
			l.forced = true
		}
		if l.options.Color != nil {
			l.color = l.options.Color
		}
		if l.options.Output != nil {
			l.output = l.options.Output
		}
		if l.options.Threaded {
			l.channel = make(chan string)
			go l.listen()
		}
		if l.options.Timestamp != nil {
			l.timestamp = l.options.Timestamp
		}
	}

	if l.color == nil {
		if l.options != nil {
			l.setRandomColor(l.options.UseBackgroundColors)
		} else {
			l.setRandomColor(false)
		}
	}
}
