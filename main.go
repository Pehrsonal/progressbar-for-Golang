package main

import (
	"time"

	"github.com/Pehrsonal/Progressbar-For-Golang/progressbar"
)

//Main test function
func main() {

	look := progressbar.Style{
		StartChar:     "!",
		EndChar:       "!",
		ProgressChar:  "C",
		StartEndColor: "Red",
		ProgressColor: "Black",
	}
	test := progressbar.StartNew(50, progressbar.SetWidth(50), progressbar.ShowPercent(false),
		progressbar.ShowTime(false), progressbar.SetStyle(look), progressbar.NewDescription("HAHA"))

	for i := 0; i < test.GetMaxvalue(); i++ {
		test.Increment()
		time.Sleep(80 * time.Millisecond)
	}
	test.Finish()
}
