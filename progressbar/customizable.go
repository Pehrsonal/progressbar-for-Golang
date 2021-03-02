package progressbar

import "errors"

// Change is the type to pass in when changing the default values in the progressbar
type Change func(bar *theBar) error

// SetWidth let you change the width of the progressbar
func SetWidth(setW int) Change {
	return func(bar *theBar) error {
		if setW <= 0 {
			return errors.New("Negativ widht")
		}
		bar.width = setW
		return nil
	}
}

//SetStyle to change the style of the progressbar
func SetStyle(style Style) Change {
	return func(bar *theBar) error {
		if bar.theme.StartChar == "0" || bar.theme.EndChar == "0" || bar.theme.ProgressChar == "0" {
			return errors.New("invalid characters")
		}
		bar.theme = style
		return nil
	}
}

// ShowPercent returns a function for setting whether the Bar displays a percentage.
func ShowPercent(show bool) Change {
	return func(b *theBar) error {
		b.showPercentage = show
		return nil
	}
}

// ShowTime returns a function for setting whether the Bar displays elapsed time.
func ShowTime(showTime bool) Change {
	return func(b *theBar) error {
		b.showTime = showTime
		return nil
	}
}

// NewDescription change the description in front of the bar
func NewDescription(dec string) Change {
	return func(b *theBar) error {
		if dec == "" {
			return errors.New("Empty description not possible")
		}
		b.description = dec
		return nil
	}
}
