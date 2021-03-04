//Package progressbar makes it possible to create a progressbar to see the progress in
// any kind of loops and how long time the work in it takes!
package progressbar

import (
	"github.com/fatih/color"
)

//
// Made functions for the "normal colors" this for easier use for customers
//
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

//getColor is a switch that only checks what color they want if none the bar is White
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
