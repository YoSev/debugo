package debugo

func (l *Debugger) applyOptions(options *Options) {
	if options != nil {
		if options.ForceEnable {
			l.forced = true
		}
		if options.Color != nil {
			l.color = options.Color
		}
		if options.Output != nil {
			l.output = options.Output
		}
		if options.Threaded {
			l.channel = make(chan string)
			go l.listen()
		}
		if options.Timestamp != nil {
			l.timestamp = options.Timestamp
		}
	}

	if l.color == nil {
		if options != nil {
			l.setRandomColor(options.UseBackgroundColors)
		} else {
			l.setRandomColor(false)
		}
	}
}
