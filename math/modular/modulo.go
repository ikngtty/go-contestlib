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
