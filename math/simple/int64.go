package simple

// Abs64 returns the absolute value of x.
func Abs64(x int64) int64 {
	if x >= 0 {
		return x
	}
	return -x
}

// Min64 returns the minimum value of the specified values.
func Min64(values ...int64) int64 {
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

// Max64 returns the maximum value of the specified values.
func Max64(values ...int64) int64 {
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

// Pow64 returns base^exponent.
func Pow64(base int64, exponent uint) int64 {
	answer := int64(1)
	for i := uint(0); i < exponent; i++ {
		answer *= base
	}
	return answer
}
