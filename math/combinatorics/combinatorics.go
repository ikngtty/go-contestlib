package combinatorics

import "fmt"

// Permutations generates k-permutations of n and applies f for each permutation.
func Permutations(n, k int, f func([]int)) {
	if n < 0 || k < 0 || n < k {
		panic(fmt.Sprintf("invalid arguments (n, k): (%d, %d)", n, k))
	}

	checklist := make([]bool, n)
	pattern := make([]int, k)

	var body func(pos int)
	body = func(pos int) {
		if pos == k {
			f(pattern)
			return
		}

		for num := range checklist {
			if checklist[num] {
				continue
			}

			pattern[pos] = num
			checklist[num] = true
			body(pos + 1)
			checklist[num] = false
		}
	}
	body(0)
}

// Combinations generates k-combinations of n and applies f for each combination.
func Combinations(n, k int, f func([]int)) {
	if n < 0 || k < 0 || n < k {
		panic(fmt.Sprintf("invalid arguments (n, k): (%d, %d)", n, k))
	}

	pattern := make([]int, k)

	var body func(pos, begin int)
	body = func(pos, begin int) {
		if pos == k {
			f(pattern)
			return
		}

		for num := begin; num < n+pos-k+1; num++ {
			pattern[pos] = num
			body(pos+1, num+1)
		}
	}
	body(0, 0)
}

// DupPermutations generates duplicated k-permutations of n and applies f for each permutation.
func DupPermutations(n, k int, f func([]int)) {
	if n < 0 || k < 0 || n < k {
		panic(fmt.Sprintf("invalid arguments (n, k): (%d, %d)", n, k))
	}

	pattern := make([]int, k)

	var body func(pos int)
	body = func(pos int) {
		if pos == k {
			f(pattern)
			return
		}

		for num := 0; num < n; num++ {
			pattern[pos] = num
			body(pos + 1)
		}
	}
	body(0)
}

// DupCombinations generates duplicated k-combinations of n and applies f for each combination.
func DupCombinations(n, k int, f func([]int)) {
	if n < 0 || k < 0 || n < k {
		panic(fmt.Sprintf("invalid arguments (n, k): (%d, %d)", n, k))
	}

	pattern := make([]int, k)

	var body func(pos, begin int)
	body = func(pos, begin int) {
		if pos == k {
			f(pattern)
			return
		}

		for num := begin; num < n; num++ {
			pattern[pos] = num
			body(pos+1, num)
		}
	}
	body(0, 0)
}

// BitPatterns generates all bit patterns which size is bitsLen and applies f for each pattern.
func BitPatterns(bitsLen int, f func([]bool)) {
	if bitsLen < 0 {
		panic(fmt.Sprintf("invalid bitsLen: %d", bitsLen))
	}

	pattern := make([]bool, bitsLen)

	var body func(pos int)
	body = func(pos int) {
		if pos == bitsLen {
			f(pattern)
			return
		}

		pattern[pos] = false
		body(pos + 1)
		pattern[pos] = true
		body(pos + 1)
	}
	body(0)
}
