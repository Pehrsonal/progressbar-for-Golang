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
		ProgressColor: "Yellow",
	}

	look1 := progressbar.Style{
		StartChar:     "{",
		EndChar:       "}",
		ProgressChar:  "%",
		StartEndColor: "Blue",
		ProgressColor: "Red",
	}

	test := progressbar.StartNew(50, progressbar.SetWidth(50), progressbar.ShowPercent(false),
		progressbar.ShowTime(false), progressbar.SetStyle(look), progressbar.NewDescription("HAHA"))

	hest := progressbar.New(30, progressbar.SetWidth(10), progressbar.ShowPercent(true),
		progressbar.ShowTime(true))
	for i := 0; i < test.GetMaxvalue(); i++ {
		test.Increment()
		time.Sleep(60 * time.Millisecond)
	}
	test.Finish()
	hest.Start()
	for i := 0; i < hest.GetMaxvalue(); i++ {
		hest.Increment()
		time.Sleep(60 * time.Millisecond)
	}
	hest.Finish()

	ba := progressbar.StartNew(50, progressbar.SetWidth(10), progressbar.SetStyle(look1),
		progressbar.ShowPercent(false))
	for i := 0; i < ba.GetMaxvalue(); i++ {
		ba.Increment()
		time.Sleep(60 * time.Millisecond)
	}
	ba.Finish()
}
