package main

import (
	"time"

	"github.com/Pehrsonal/Progressbar-For-Golang/progressbar"
)

func main() {

	test := progressbar.StartNew(30)
	for i := 0; i < 30; i++ {
		test.Increment()
		time.Sleep(100 * time.Millisecond)

	}
	test.Finish()
}
