package main

import (
	"fmt"
	"strings"
	"time"
)

type theBar struct {
	whatPercent   int
	curPercent    int
	rePrint       string
	loading       string
	loadingLenght int
	max           int
}

func (theBar *theBar) createBar(current, max int) {
	theBar.curPercent = current
	theBar.loading = "#"
	theBar.loadingLenght = 40
	theBar.max = max
	theBar.whatPercent = theBar.getPercent()
}

func (theBar *theBar) getPercent() int {
	return int(float32(theBar.curPercent) / float32(theBar.max) * 100)
}

func (theBar *theBar) update(current int) {
	theBar.curPercent = current
	last := theBar.whatPercent
	theBar.whatPercent = theBar.getPercent()
	if theBar.whatPercent != last {
		theBar.rePrint = strings.Repeat(theBar.loading, int(theBar.curPercent*theBar.loadingLenght/100))
	}
	fmt.Printf("[%-50s]%3d%% %4d/%d\n", theBar.rePrint, theBar.whatPercent, theBar.curPercent, theBar.max)
}

func main() {

	var start int = 0
	var goal int = 140
	var progressBar theBar
	progressBar.createBar(start, goal)

	for {
		time.Sleep(15 * time.Millisecond)

		if progressBar.whatPercent == 100 {
			break
		}
		start = start + 1
		progressBar.update(start)
	}
}
