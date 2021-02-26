package progressbar

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

//TheBar struct for declaring whats inside
type TheBar struct {
	width          int
	value          int
	maxValue       int
	startTime      time.Time
	theme          Style
	showPercentage bool
	showTime       bool
	isFinished     bool
	lock           sync.Mutex
}

//Style struct to be able to change how the bar will look at
type Style struct {
	StartChar    byte
	EndChar      byte
	ProgressChar byte
}

// BarWidth returns a function for setting the width of a Bar.
func BarWidth(width int) func(*TheBar) {
	return func(b *TheBar) {
		if width <= 0 {
		}
		b.width = width
	}
}

func New(maxValue int, kwargs ...func(*TheBar)) *TheBar {

	theme := Style{
		StartChar:    '{',
		EndChar:      '}',
		ProgressChar: '#',
	}

	bar := &TheBar{
		width:          50,
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

	level := b.width * i / b.maxValue
	progress := strings.Repeat(string(b.theme.ProgressChar), level)
	blanks := strings.Repeat(" ", b.width-level)

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

func StartNew(maxValue int, arg ...func(*TheBar)) *TheBar {
	bar := New(maxValue, arg...)

	bar.Start()
	return bar
}

func (b *TheBar) Finish() {
	b.Set(b.maxValue)

}