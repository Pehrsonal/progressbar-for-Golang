package progressbar

// Change is the type to pass in when changing the default values in the progressbar
type Change func(bar *theBar)

// SetWidth let you change the width of the progressbar
func SetWidth(setW int) Change {
	return func(bar *theBar) {
		bar.width = setW
	}
}

//SetStyle to change the style of the progressbar
func SetStyle(style Style) Change {
	return func(bar *theBar) {
		bar.theme = style
	}
}

// ShowPercent returns a function for setting whether the Bar displays a percentage.
func ShowPercent(show bool) Change {
	return func(b *theBar) {
		b.showPercentage = show
	}
}

// ShowTime returns a function for setting whether the Bar displays elapsed time.
func ShowTime(showTime bool) Change {
	return func(b *theBar) {
		b.showTime = showTime
	}
}

// NewDescription change the description in front of the bar
func NewDescription(dec string) Change {
	return func(b *theBar) {
		b.description = dec
	}
}
