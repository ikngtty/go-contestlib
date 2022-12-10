package arrayutil

// MakeBools returns a slice of the bool array.
func MakeBools(length int, initVal bool) []bool {
	a := make([]bool, length)

	if initVal {
		for i := 0; i < length; i++ {
			a[i] = initVal
		}
	}

	return a
}

// Make2DBools returns a slice of the two-dimensional bool array.
func Make2DBools(xLen, yLen int, initVal bool) [][]bool {
	a := make([][]bool, xLen)
	for x := 0; x < xLen; x++ {
		a[x] = make([]bool, yLen)
	}

	if initVal {
		for x := 0; x < xLen; x++ {
			for y := 0; y < yLen; y++ {
				a[x][y] = initVal
			}
		}
	}

	return a
}

// Copy2DBools copy the two-dimensional bool array.
func Copy2DBools(dst *[][]bool, src [][]bool) {
	*dst = make([][]bool, len(src))
	for i := 0; i < len(src); i++ {
		(*dst)[i] = make([]bool, len(src[i]))
		copy((*dst)[i], src[i])
	}
}
