package arrayutil

// MakeInts returns a slice of the int array.
func MakeInts(length int, initVal int) []int {
	a := make([]int, length)

	if initVal != 0 {
		for i := 0; i < length; i++ {
			a[i] = initVal
		}
	}

	return a
}

// Make2DInts returns a slice of the two-dimensional int array.
func Make2DInts(xLen, yLen int, initVal int) [][]int {
	a := make([][]int, xLen)
	for x := 0; x < xLen; x++ {
		a[x] = make([]int, yLen)
	}

	if initVal != 0 {
		for x := 0; x < xLen; x++ {
			for y := 0; y < yLen; y++ {
				a[x][y] = initVal
			}
		}
	}

	return a
}

// Copy2DInts copy the two-dimensional int array.
func Copy2DInts(dst *[][]int, src [][]int) {
	*dst = make([][]int, len(src))
	for i := 0; i < len(src); i++ {
		(*dst)[i] = make([]int, len(src[i]))
		copy((*dst)[i], src[i])
	}
}
