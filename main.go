package main

import (
	"time"

	"github.com/Pehrsonal/Progressbar-For-Golang/progressbar"
)

func main() {
	look := progressbar.Style{
		StartChar:    '!',
		EndChar:      '!',
		ProgressChar: 'C',
	}
	test := progressbar.StartNew(100, progressbar.SetWidth(50), progressbar.SetStyle(look))
	for i := 0; i < test.GetMaxvalue(); i++ {
		test.Increment()
		time.Sleep(100 * time.Millisecond)

	}
	test.Finish()
}
