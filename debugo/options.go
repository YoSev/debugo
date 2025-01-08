package debugo

func (l *Logger) applyOptions(options *Options) {
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
	}

	if l.color == nil {
		if options != nil {
			l.setRandomColor(options.UseBackgroundColors)
		} else {
			l.setRandomColor(false)
		}
	}
}
