package progressbar

import (
	"github.com/fatih/color"
)

func getWhite() func(...interface{}) string {
	c := color.New(color.FgWhite).SprintFunc()
	return c
}
func getBlack() func(...interface{}) string {
	c := color.New(color.FgBlack).SprintFunc()
	return c
}
func getRed() func(...interface{}) string {
	c := color.New(color.FgRed).SprintFunc()
	return c
}
func getBlue() func(...interface{}) string {
	c := color.New(color.FgBlue).SprintFunc()
	return c
}
func getGreen() func(...interface{}) string {
	c := color.New(color.FgGreen).SprintFunc()
	return c
}
func getYellow() func(...interface{}) string {
	c := color.New(color.FgYellow).SprintFunc()
	return c
}

func getColor(rec string) func(...interface{}) string {
	switch rec {
	case "White":
		return getWhite()
	case "Black":
		return getBlack()
	case "Red":
		return getRed()
	case "Blue":
		return getBlue()
	case "Green":
		return getGreen()
	case "Yellow":
		return getYellow()
	default:
		return getWhite()
	}
}
