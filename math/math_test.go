package math

import (
	"testing"

	"gotest.tools/v3/assert"
)

//Test_MinInt Test The Min Int
func Test_MinInt(t *testing.T) {
	assert.Equal(t, MinInt(6, 5, 7, 9, 2, 1, 9), 1)
}

//Test_MinInt64 Test The Min Int64
func Test_MinInt64(t *testing.T) {
	assert.Equal(t, MinInt64(6, 5, 7, 9, 2, 1, 9), int64(1))
}

//Test_MaxInt Test The Max Int
func Test_MaxInt(t *testing.T) {
	assert.Equal(t, MaxInt(6, 5, 7, 9, 2, 1, 9), 9)
}

//Test_MaxInt64 Test The Max Int64
func Test_MaxInt64(t *testing.T) {
	assert.Equal(t, MaxInt64(6, 5, 7, 9, 2, 1, 9), int64(9))
}

//Test_MaxInt64 Test Abs Int
func Test_AbsInt(t *testing.T) {
	assert.Equal(t, AbsInt(9), 9)
	assert.Equal(t, AbsInt(0), 0)
	assert.Equal(t, AbsInt(-9), 9)
}

//Test_AbsInt64 Test Abs Int64
func Test_AbsInt64(t *testing.T) {
	assert.Equal(t, AbsInt64(int64(-9)), int64(9))
	assert.Equal(t, AbsInt64(int64(0)), int64(0))
	assert.Equal(t, AbsInt64(int64(9)), int64(9))
}
