package combinatorics

import "fmt"

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

func BitPatterns(bitsLen int, f func([]bool)) {
	if bitsLen < 0 {
		panic(fmt.Sprintf("bitsLen (%d) should not be a minus", bitsLen))
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
