package int64

func Abs(x int64) int64 {
	if x >= 0 {
		return x
	}
	return -x
}

func Min(values ...int64) int64 {
	if len(values) == 0 {
		panic("no values")
	}
	min := values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
	}
	return min
}

func Max(values ...int64) int64 {
	if len(values) == 0 {
		panic("no values")
	}
	max := values[0]
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return max
}

func Pow(base int64, exponent uint) int64 {
	answer := int64(1)
	for i := uint(0); i < exponent; i++ {
		answer *= base
	}
	return answer
}
