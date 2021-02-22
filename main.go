package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type TheBar struct {
	barWidth       int
	value          int
	maxValue       int
	startTime      time.Time
	theme          Style
	showPercentage bool
	showTime       bool
	isFinished     bool
	lock           sync.Mutex
}

type Style struct {
	StartChar    rune
	EndChar      rune
	ProgressChar rune
}

func New(maxValue int) *TheBar {

	theme := Style{
		StartChar:    '{',
		EndChar:      '}',
		ProgressChar: '#',
	}

	bar := &TheBar{
		barWidth:       50,
		value:          0,
		maxValue:       maxValue,
		theme:          theme,
		showPercentage: true,
		showTime:       true,
		isFinished:     false,
	}

	return bar
}

func (b *TheBar) update(i int) {

	if b.isFinished {
		return
	}

	level := b.barWidth * i / b.maxValue
	progress := strings.Repeat(string(b.theme.ProgressChar), level)
	blanks := strings.Repeat(" ", b.barWidth-level)

	fmt.Printf("\rProgress: %s%s%s%s", string(b.theme.StartChar), progress, blanks, string(b.theme.EndChar))

	if b.showPercentage {
		percentage := 100 * float32(i) / float32(b.maxValue)
		fmt.Printf(" %.2f%%", percentage)
	}
	if b.showTime {
		elapsed := time.Since(b.startTime).Seconds()
		fmt.Printf(" - %.2fs ", elapsed)
	}

	b.value = i
}

func (b *TheBar) end() {

	if b.isFinished {
		return
	}

	b.update(b.maxValue)
	b.isFinished = true

	elapsed := time.Since(b.startTime)
	fmt.Printf("\nTime it took: %fs\n", elapsed.Seconds())
}

func (b *TheBar) Set(i int) {
	b.lock.Lock()
	defer b.lock.Unlock()

	if b.startTime.IsZero() {
		b.Start()
	}

	if i >= b.maxValue {
		b.end()
	} else if i < 0 {
		b.update(0)
	} else {
		b.update(i)
	}
}

func (b *TheBar) Add(i int) {
	b.Set(b.value + i)
}

func (b *TheBar) Increment() {
	b.Add(1)
}

func (b *TheBar) Start() {
	b.startTime = time.Now()
	fmt.Printf("\n")
	b.Set(0)
}

func StartNew(maxValue int) *TheBar {
	bar := New(maxValue)
	bar.Start()
	return bar
}

func (b *TheBar) Finish() {
	b.Set(b.maxValue)
}

func main() {
	b := StartNew(50)
	for i := 0; i < 50; i++ {
		b.Increment()
		time.Sleep(20 * time.Millisecond)
	}
	b.Finish()
}
