package pinyin

import (
	"reflect"
	"testing"
)

// todo 拼音占坑
func TestGeneratePinyin(t *testing.T) {
	str := "你好star不好"
	pinyin, first := GeneratePinyin(str)

	equal(t, pinyin, "nihaostarbuhao")
	equal(t, first, "nhsbh")
}

// todo 拼音占坑
func TestResetPinyin(t *testing.T) {
	equal(t, resetString("你好star不好"), "你好s不好")
	equal(t, resetString("你好1234star不好"), "你好s不好")
	equal(t, resetString("a你好1234star不好"), "a你好s不好")
	equal(t, resetString("a你好1234-=/';';star不好"), "a你好s不好")
}

// todo 拼音占坑
func TestIsLetter(t *testing.T) {
	equal(t, IsLetter('a'), true)
	equal(t, IsLetter(1), false)
	equal(t, IsLetter('Z'), true)

	equal(t, IsLetter('A'), true)
	equal(t, IsLetter('哈'), false)
	equal(t, IsLetter('z'), true)
}

// Expected to be equal.
func equal(t *testing.T, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf(
			"Expected %v (type %v) - Got %v (type %v)",
			expected,
			reflect.TypeOf(expected),
			actual,
			reflect.TypeOf(actual),
		)
	}
}

// Expected to be unequal.
func unequal(t *testing.T, expected, actual interface{}) {
	if reflect.DeepEqual(expected, actual) {
		t.Errorf(
			"Did not expect %v (type %v) - Got %v (type %v)",
			expected,
			reflect.TypeOf(expected),
			actual,
			reflect.TypeOf(actual),
		)
	}
}

// Expect a greater than b.
func gt(t *testing.T, a, b float64) {
	if a <= b {
		t.Errorf("Expected %v (type %v) > Got %v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
	}
}

// Expect a greater than or equal to b.
func gte(t *testing.T, a, b float64) {
	if a < b {
		t.Errorf("Expected %v (type %v) > Got %v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
	}
}

// Expected value needs to be within range.
func rangeValue(t *testing.T, min, max, actual float64) {
	if actual < min || actual > max {
		t.Errorf(
			"Expected range of %v-%v (type %v) > Got %v (type %v)",
			min,
			max,
			reflect.TypeOf(min),
			actual,
			reflect.TypeOf(actual),
		)
	}
}
