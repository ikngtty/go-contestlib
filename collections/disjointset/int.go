package disjointset

type DisjointSetForest struct {
	n      int
	parent []int
	size   []int
}

func NewDisjointSetForest(n int) *DisjointSetForest {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = -1
	}

	size := make([]int, n)
	for i := range size {
		size[i] = 1
	}

	return &DisjointSetForest{n, parent, size}
}

func (f *DisjointSetForest) Add(x int) {
	f.parent[x] = x
}

func (f *DisjointSetForest) Root(x int) int {
	if f.parent[x] == -1 {
		return -1
	} else if f.parent[x] == x {
		return x
	}
	f.parent[x] = f.Root(f.parent[x])
	return f.parent[x]
}

func (f *DisjointSetForest) Merge(x, y int) {
	xRoot := f.Root(x)
	yRoot := f.Root(y)
	if xRoot == yRoot {
		return
	}

	var lower, higher int
	if f.size[xRoot] < f.size[yRoot] {
		lower = xRoot
		higher = yRoot
	} else {
		lower = yRoot
		higher = xRoot
	}

	f.size[higher] += f.size[lower]
	f.parent[lower] = higher
}

func (f *DisjointSetForest) Same(x, y int) bool {
	return f.Root(x) == f.Root(y)
}

func (f *DisjointSetForest) Count() int {
	roots := make([]bool, f.n)
	for i := 0; i < f.n; i++ {
		root := f.Root(i)
		if root == -1 {
			continue
		}

		roots[root] = true
	}

	count := 0
	for _, exist := range roots {
		if exist {
			count++
		}
	}
	return count
}
