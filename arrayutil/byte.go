package arrayutil

// Make2DByteArray returns a slice of the two-dimensional byte array.
func Make2DByteArray(xLen, yLen int) [][]byte {
	a := make([][]byte, xLen)
	for x := 0; x < xLen; x++ {
		a[x] = make([]byte, yLen)
	}
	return a
}
