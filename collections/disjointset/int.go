package disjointset

type DisjointSetForest struct {
	n      int
	parent []int
	size   []int
}

func NewDisjointSetForest(n int) *DisjointSetForest {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}

	size := make([]int, n)
	for i := range size {
		size[i] = 1
	}

	return &DisjointSetForest{n, parent, size}
}

func (f *DisjointSetForest) Root(x int) int {
	if f.parent[x] == x {
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
