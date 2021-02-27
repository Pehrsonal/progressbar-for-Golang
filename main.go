package main

import (
	"time"

	"github.com/Pehrsonal/Progressbar-For-Golang/progressbar"
)

func main() {

	test := progressbar.StartNew(500, progressbar.SetWidth(100))
	for i := 0; i < test.GetMaxvalue(); i++ {
		test.Increment()
		time.Sleep(100 * time.Millisecond)

	}
	test.Finish()
}
