package bar

import {
	"strings"
	"errors"
	"time"

}

type style struct {
	startCharachter    byte
	progressChatachter byte
	endCharachter      byte
}

type theBar struct {
	barWitdh 	int
	look     	style
	val         int
	maxVal      int
	startTime   time.Time
	isDone		bool
}

//function for setting the width of the bar
func barWidth(width int) func(*Bar) error {
	return func(b *Bar) error {
		if width <= 0 {
			return errors.New("width must be positive")
		}
		b.width = width
		return nil
	}
}
	

func New(maxVal int, kwargs ...func(*theBar) error) (*theBar, error) {

	if maxVal <= 0 {
		return nil, errors.New("maxVal must be positive")
	}

	//Default values
	theme := style{
		startCharachter:    '[',
		progressChatachter: '#',
		endCharachter:      ']',
	}
	progressBar := &bar{
		width:          50,
		val:            0,
		maxVal:         maxVal,
		theme:          theme,
		isDone:     false,
	}

	// Apply optional arguments such as special width... will look into more
	for _, arg := range kwargs {
		err := arg(bar)
		if err != nil {
			return nil, err
		}
	}

	return progressBar, nil
}

func (b *theBar) update(i int) {

	if b.isFinished {
		return
	}

	// Generate characters to indicate progress
	level := b.width * i / b.maxVal
	progress := strings.Repeat(string(b.theme.ProgressChar), level)
	blanks := strings.Repeat(" ", b.width-level)

	fmt.Printf("\rProgress: %s%s%s%s", string(b.theme.StartChar), progress, blanks, string(b.theme.EndChar))

	b.val = i
}

func (b *theBar) end() {

	if b.isDone {
		return
	}

	b.update(b.maxVal)
	b.isDone = true

	elapsed := time.Since(b.startTime)
	fmt.Printf("\nWall time: %f\n", elapsed.Seconds())
}

// Set sets a new value of the Bar
func (b *Bar) Set(i int) {
	b.lock.Lock()
	defer b.lock.Unlock()

	if b.startTime.IsZero() {
		b.Start()
	}

	if i >= b.maxVal {
		b.end()
	} else if i < 0 {
		b.update(0)
	} else {
		b.update(i)
	}
}

func (b *theBar) Add(i int) {
	b.Set(b.val + i)
}

func (b *theBar) Increment() {
	b.Add(1)
}


func (b *theBar) Start() {
	b.startTime = time.Now()
	fmt.Printf("\n")
	b.Set(0)
}

func (b *theBar) Finish() {
	b.Set(b.maxVal)
}