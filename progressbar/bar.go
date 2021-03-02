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
	description    string
}

//Style is how to progressbar will look
type Style struct {
	StartChar     string
	EndChar       string
	ProgressChar  string
	StartEndColor string
	ProgressColor string
}

//New Creates a new bar with default values
func New(maxValue int, arg ...Change) *theBar {
	//Default values of the progressbar
	theme := Style{
		StartChar:     "{",
		EndChar:       "}",
		ProgressChar:  "#",
		StartEndColor: "White",
		ProgressColor: "White",
	}
	bar := theBar{
		width:          50,
		value:          0,
		maxValue:       maxValue,
		theme:          theme,
		showPercentage: true,
		showTime:       true,
		isFinished:     false,
		description:    "Progress",
	}

	for _, o := range arg {
		o(&bar)
	}

	return &bar
}

func (b *theBar) update(i int) {
	if b.isFinished {
	}
	level := b.width * i / b.maxValue
	progress := strings.Repeat(string(b.theme.ProgressChar), level)
	blanks := strings.Repeat(" ", b.width-level)

	whatColorStartEnd := getColor(b.theme.StartEndColor)
	whatColorProgress := getColor(b.theme.ProgressColor)

	fmt.Printf("\r%s: %s%s%s%s", b.description, whatColorStartEnd((string(b.theme.StartChar))), whatColorProgress(progress), blanks, whatColorStartEnd(string(b.theme.EndChar)))

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

func (b *theBar) set(i int) {
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

func (b *theBar) add(i int) {
	b.set(b.value + i)
}

func (b *theBar) Increment() {
	b.add(1)
}

//Start the timer for the bar
func (b *theBar) Start() {
	b.startTime = time.Now()
	fmt.Printf("\n")
	b.set(0)
}

//GetMaxvalue returns the maxvalue from the bar
func (b *theBar) GetMaxvalue() int {
	return b.maxValue
}

//StartNew creates a new bar with default values and takes in Change values if wanted + starts the counter
func StartNew(maxValue int, arg ...Change) *theBar {
	bar := New(maxValue, arg...)

	bar.Start()
	return bar
}

func (b *theBar) Finish() {
	b.set(b.maxValue)
}
