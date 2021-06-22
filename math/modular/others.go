package modular

import "github.com/ikngtty/go-contestlib/math/simple"

// ExtendedEuclidean does the extended Euclidean algorithm.
func ExtendedEuclidean(a, b int) (gcd, x, y int) {
	if b == 0 {
		gcd = a
		x = 1
		y = 0
		return
	}
	q, r := simple.EucDiv(a, b)
	gcd, s, t := ExtendedEuclidean(b, r)
	x = t
	y = s - q*t
	return
}
