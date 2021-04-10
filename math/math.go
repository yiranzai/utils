package math

//MinInt get the min int vars
func MinInt(vars ...int) int {
	min := vars[0]

	for _, i := range vars {
		if min > i {
			min = i
		}
	}

	return min
}

//MinInt64 get the min int64 vars
func MinInt64(vars ...int64) int64 {
	min := vars[0]

	for _, i := range vars {
		if min > i {
			min = i
		}
	}

	return min
}

//MaxInt get the max int vars
func MaxInt(vars ...int) int {
	max := vars[0]

	for _, i := range vars {
		if max < i {
			max = i
		}
	}

	return max
}

//MaxInt64 get the max int64 vars
func MaxInt64(vars ...int64) int64 {
	max := vars[0]

	for _, i := range vars {
		if max < i {
			max = i
		}
	}

	return max
}

//AbsInt abs int
func AbsInt(abs int) int {
	if abs < 0 {
		return 0 - abs
	}
	return abs
}

//AbsInt64 abs int64
func AbsInt64(abs int64) int64 {
	if abs < 0 {
		return 0 - abs
	}
	return abs
}
