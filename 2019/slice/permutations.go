package slice

// PermutationsInt64 returns all permutations of elements from the given array
func PermutationsInt64(a []int64) [][]int64 {
	n := len(a)
	if n == 0 {
		return nil
	}

	g := newPermutationsGenerator(a)
	g.generate(n)
	return g.p
}

type permutationsGenerator struct {
	a []int64
	p [][]int64
}

func newPermutationsGenerator(a []int64) *permutationsGenerator {
	return &permutationsGenerator{
		a: CopyInt64(a),
	}
}

func (g *permutationsGenerator) generate(k int) {
	if k == 1 {
		g.p = append(g.p, CopyInt64(g.a))
		return
	}

	g.generate(k - 1)

	for i := 0; i < k-1; i++ {
		if k%2 == 0 {
			g.swap(i, k-1)
		} else {
			g.swap(0, k-1)
		}
		g.generate(k - 1)
	}
}

func (g *permutationsGenerator) swap(i int, j int) {
	g.a[i], g.a[j] = g.a[j], g.a[i]
}
