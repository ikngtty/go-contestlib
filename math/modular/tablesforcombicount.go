package modular

import "fmt"

// TablesForCombiCount holds tables to calculate combination counts.
type TablesForCombiCount struct {
	mod           Mod
	length        int
	factorials    []int
	invs          []int
	factorialInvs []int
}

// NewTablesForCombiCount returns a new TablesForCombiCount.
func NewTablesForCombiCount(mod Mod, length int) *TablesForCombiCount {
	if length < 0 {
		panic(fmt.Sprintf("invalid length: %d", length))
	}
	return &TablesForCombiCount{mod, length, nil, nil, nil}
}

// Factorials returns a factorial table.
func (t *TablesForCombiCount) Factorials() []int {
	if t.factorials == nil {
		t.factorials = t.mod.Factorials(t.length)
	}
	return t.factorials
}

// SetFactorials sets a factorial table.
func (t *TablesForCombiCount) SetFactorials(factorials []int) {
	t.factorials = factorials
}

// Invs returns an inverse table.
func (t *TablesForCombiCount) Invs() []int {
	if t.invs == nil {
		t.invs = t.mod.Invs(t.length)
	}
	return t.invs
}

// SetInvs sets an inverse table.
func (t *TablesForCombiCount) SetInvs(invs []int) {
	t.invs = invs
}

// FactorialInvs returns a factorial inverse table.
func (t *TablesForCombiCount) FactorialInvs() []int {
	if t.factorialInvs == nil {
		t.factorialInvs = make([]int, t.length)
		if t.length == 0 {
			return t.factorialInvs
		}
		t.factorialInvs[0] = 1
		invs := t.Invs()
		for i := 1; i < t.length; i++ {
			t.factorialInvs[i] = t.mod.Reg(t.factorialInvs[i-1] * invs[i])
		}
	}
	return t.factorialInvs
}

// SetFactorialInvs sets a factorial inverse table.
func (t *TablesForCombiCount) SetFactorialInvs(factorialInvs []int) {
	t.factorialInvs = factorialInvs
}

// Calc returns a combination count.
func (t *TablesForCombiCount) Calc(n, k int) int {
	if n < 0 || k < 0 || n < k {
		panic(fmt.Sprintf("invalid arguments (n, k): (%d, %d)", n, k))
	}
	if n > t.length-1 {
		panic(fmt.Sprintf("n should be less than length: %d", n))
	}

	facts := t.Factorials()
	factInvs := t.FactorialInvs()
	return t.mod.Reg(t.mod.Reg(facts[n]*factInvs[k]) * factInvs[n-k])
}
