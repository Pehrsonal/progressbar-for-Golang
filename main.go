package main

import (
	"time"

	"github.com/Pehrsonal/Progressbar-For-Golang/progressbar"
)

//Main test function
func main() {
	look := progressbar.Style{
		StartChar:    '!',
		EndChar:      '!',
		ProgressChar: 'C',
	}
	test := progressbar.StartNew(100, progressbar.SetWidth(50), progressbar.SetStyle(look), progressbar.BarShowPercent(false),
		progressbar.BarShowTime(false))
	for i := 0; i < test.GetMaxvalue(); i++ {
		test.Increment()
		time.Sleep(100 * time.Millisecond)

	}
	test.Finish()
}
