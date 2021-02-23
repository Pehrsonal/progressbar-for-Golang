package progressbar

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
)

type Bar struct {
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

func New(maxValue int) *Bar {

	theme := Style{
		StartChar:    '{',
		EndChar:      '}',
		ProgressChar: '#',
	}

	bar := &Bar{
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

func (b *Bar) update(i int) {

	if b.isFinished {
		return
	}

	level := b.barWidth * i / b.maxValue
	progress := strings.Repeat(string(b.theme.ProgressChar), level)
	blanks := strings.Repeat(" ", b.barWidth-level)

	color.Red("We have red") // testing color import

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

func (b *Bar) end() {

	if b.isFinished {
		return
	}

	b.update(b.maxValue)
	b.isFinished = true

	elapsed := time.Since(b.startTime)
	fmt.Printf("\nWall time: %f\n", elapsed.Seconds())
}

func (b *Bar) Set(i int) {
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

func (b *Bar) Add(i int) {
	b.Set(b.value + i)
}

func (b *Bar) Increment() {
	b.Add(1)
}

func (b *Bar) Start() {
	b.startTime = time.Now()
	fmt.Printf("\n")
	b.Set(0)
}

func StartNew(maxValue int) *Bar {
	bar := New(maxValue)
	bar.Start()
	return bar
}

func (b *Bar) Finish() {
	b.Set(b.maxValue)
}
