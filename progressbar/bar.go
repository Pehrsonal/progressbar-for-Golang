package progressbar

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

//theBar struct for declaring whats inside
type theBar struct {
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

//Style is how to progressbar will look
type Style struct {
	StartChar    byte
	EndChar      byte
	ProgressChar byte
}

// Change is the type to pass in when changing the default values in the progressbar
type Change func(bar *theBar)

// SetWidth let you change the width of the progressbar
func SetWidth(setW int) Change {
	return func(bar *theBar) {
		bar.width = setW
	}
}

//New Creates a new bar with default values
func New(maxValue int, arg ...Change) *theBar {

	//Default values of the progressbar
	theme := Style{
		StartChar:    '{',
		EndChar:      '}',
		ProgressChar: '#',
	}
	bar := theBar{
		width:          50,
		value:          0,
		maxValue:       maxValue,
		theme:          theme,
		showPercentage: true,
		showTime:       true,
		isFinished:     false,
	}

	for _, o := range arg {
		o(&bar)
	}

	return &bar
}

func (b *theBar) update(i int) {

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

func (b *theBar) end() {

	if b.isFinished {
		return
	}

	b.update(b.maxValue)
	b.isFinished = true

	elapsed := time.Since(b.startTime)
	fmt.Printf("\nTime it took: %fs\n", elapsed.Seconds())
}

func (b *theBar) Set(i int) {
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

func (b *theBar) Add(i int) {
	b.Set(b.value + i)
}

func (b *theBar) Increment() {
	b.Add(1)
}

func (b *theBar) Start() {
	b.startTime = time.Now()
	fmt.Printf("\n")
	b.Set(0)
}
//StartNew creates a new bar with default values and takes in Change values if wanted + starts the counter
func StartNew(maxValue int, arg ...Change) *theBar {
	bar := New(maxValue, arg...)

	bar.Start()
	return bar
}

func (b *theBar) Finish() {
	b.Set(b.maxValue)
}

//GetMaxvalue returns the maxvalue from the bar
func (b *theBar) GetMaxvalue() int {
	return b.maxValue
}
