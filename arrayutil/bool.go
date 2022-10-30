package arrayutil

// MakeBoolArray returns a slice of the bool array.
func MakeBoolArray(length int, initVal bool) []bool {
	a := make([]bool, length)

	if initVal {
		for i := 0; i < length; i++ {
			a[i] = initVal
		}
	}

	return a
}

// Make2DBoolArray returns a slice of the two-dimensional bool array.
func Make2DBoolArray(xLen, yLen int, initVal bool) [][]bool {
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
