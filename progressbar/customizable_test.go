package progressbar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SetWidth(t *testing.T) {
	assert := assert.New(t)

	bar := StartNew(50)
	assert.NoError(SetWidth(10)(bar), "Should not throw error")
	assert.Equal(10, bar.width, "Bar width not updated")
	assert.Error(SetWidth(-1)(bar), "Should throw error")
}

func Test_SetStyle(t *testing.T) {
	assert := assert.New(t)

	bar := StartNew(50)
	look1 := Style{
		StartChar:     "{",
		EndChar:       "}",
		ProgressChar:  "%",
		StartEndColor: "Blue",
		ProgressColor: "Red",
	}
	assert.NoError(SetStyle(look1)(bar), "Should not throw error")
}

func Test_ShowPercent(t *testing.T) {
	assert := assert.New(t)

	bar := StartNew(50)
	assert.NoError(ShowPercent(true)(bar), "Should not throw error")
	assert.True(true, bar.showPercentage, "show percentage should be enabled")
}

func Test_ShowTime(t *testing.T) {
	assert := assert.New(t)

	bar := StartNew(50)
	assert.NoError(ShowTime(true)(bar), "Should not throw error")
	assert.True(true, bar.showTime, "show time should be enabled")
}

func Test_NewDescription(t *testing.T) {
	assert := assert.New(t)

	bar := StartNew(50)
	assert.NoError(NewDescription("test")(bar), "Should not throw error")
	assert.Equal("test", bar.description, "description is not matching")
	assert.Error(NewDescription("")(bar), "Should throw error")
}
