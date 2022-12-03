package simple

import "fmt"

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

// Min returns the minimum value of the specified values.
func Min(values ...int) int {
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

// Max returns the maximum value of the specified values.
func Max(values ...int) int {
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

// Pow returns base^exponent.
func Pow(base, exponent int) int {
	if exponent < 0 {
		panic(fmt.Sprintf("invalid exponent: %d", exponent))
	}

	if exponent == 0 {
		return 1
	}

	if exponent%2 == 0 {
		half := Pow(base, exponent/2)
		return half * half
	} else {
		return base * Pow(base, exponent-1)
	}
}

// Ceil returns ceil(divident/dividor).
func Ceil(divident, dividor int) int {
	if dividor == 0 {
		panic("divide by zero")
	}

	quo := divident / dividor
	rem := divident % dividor

	if rem != 0 {
		if (divident > 0 && dividor > 0) ||
			(divident < 0 && dividor < 0) {
			return quo + 1
		}
	}
	return quo
}

// EucDiv does Euclidean divison.
func EucDiv(divident, dividor int) (quo, rem int) {
	// // duplicated with EucRem
	// if dividor == 0 {
	// 	panic("divide by zero")
	// }

	rem = EucRem(divident, dividor)
	quo = (divident - rem) / dividor
	return
}

// EucRem returns the remainder of Euclidean divison.
func EucRem(divident, dividor int) int {
	if dividor == 0 {
		panic("divide by zero")
	}

	rem := divident % dividor
	if rem < 0 {
		rem += Abs(dividor)
	}
	return rem
}

// GCD returns the Greatest Common Divisor.
func GCD(a, b int) int {
	if a <= 0 {
		panic(fmt.Sprintf("invalid value: %d", a))
	}
	if b <= 0 {
		panic(fmt.Sprintf("invalid value: %d", b))
	}

	var body func(a, b int) int
	body = func(a, b int) int {
		r := a % b
		if r == 0 {
			return b
		}
		return body(b, r)
	}
	return body(a, b)
}
