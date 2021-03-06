package progressbar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Increment(t *testing.T) {

	assert := assert.New(t)
	bar := StartNew(50)
	for i := 1; i <= bar.GetMaxvalue(); i++ {
		bar.Increment()
		assert.Equal(i, bar.value, "Bar value after increment is not valid")
	}
}

func Test_Finish(t *testing.T) {

	assert := assert.New(t)
	bar := StartNew(50)
	bar.Finish()
	assert.Equal(50, bar.value, "Bar value should be the max value")

	bar.set(0)
	bar.Finish()
	assert.Equal(0, bar.value, "Bar value should be update to max as already finished")
}
