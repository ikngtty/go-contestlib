package arrayutil

// Make2DBytes returns a slice of the two-dimensional byte array.
func Make2DBytes(xLen, yLen int) [][]byte {
	a := make([][]byte, xLen)
	for x := 0; x < xLen; x++ {
		a[x] = make([]byte, yLen)
	}
	return a
}

// Copy2DBytes copy the two-dimensional byte array.
func Copy2DBytes(dst *[][]byte, src [][]byte) {
	*dst = make([][]byte, len(src))
	for i := 0; i < len(src); i++ {
		(*dst)[i] = make([]byte, len(src[i]))
		copy((*dst)[i], src[i])
	}
}
