package modular

import "fmt"

// Mod is the struct that holds the modulus and provides modulo operations.
type Mod struct {
	modulus int
}

// NewMod returns a new Mod.
func NewMod(modulus int) Mod {
	if modulus <= 0 {
		panic(fmt.Sprintf("not positive modulus: %d", modulus))
	}
	return Mod{modulus}
}

// Modulus returns the modulus.
func (m Mod) Modulus() int {
	return m.modulus
}

// Reg returns the regularized number R, which satisfies 0 <= R < M,
// where M is the modulus. It is also the remainder of Euclidean division.
func (m Mod) Reg(n int) int {
	rem := n % m.modulus
	if rem < 0 {
		rem += m.modulus
	}
	return rem
}

// Inv returns the inverse.
func (m Mod) Inv(n int) int {
	if n%m.modulus == 0 {
		panic(fmt.Sprintf("congruent with 0 (mod %d): %d", m.modulus, n))
	}

	gcd, x, _ := ExtendedEuclidean(n, m.modulus)
	if gcd > 1 {
		panic(fmt.Sprintf("not prime to the modulus %d: %d", m.modulus, n))
	}

	return m.Reg(x)
}

// Invs returns the inverses of [0, 1, ..., n-1].
func (mod Mod) Invs(n int) []int {
	m := mod.modulus

	if n < 0 {
		panic(fmt.Sprintf("invalid length: %d", n))
	}
	if n > m {
		panic(fmt.Sprintf("inverses more than %d are redundant", m))
	}

	invs := make([]int, n)
	if n > 1 {
		invs[1] = 1
	}
	for i := 2; i < n; i++ {
		invs[i] = mod.Reg(-1 * invs[m%i] * (m / i))
	}
	return invs
}
