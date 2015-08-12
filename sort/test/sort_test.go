package test

import (
	"fmt"
	"github.com/dtromb/gogsl/permutation"
	"github.com/dtromb/gogsl/rng"
	"github.com/dtromb/gogsl/sort"
	"github.com/dtromb/gogsl/vector"
	"testing"
)

func TestSort(t *testing.T) {
	var n int = 100000
	var k int = 5
	x := make([]float64, n)
	small := make([]float64, k)
	rng.EnvSetup()
	T := rng.DefaultRngType()
	r := rng.RngAlloc(T)
	for i := 0; i < n; i++ {
		x[i] = rng.Uniform(r)
	}
	sort.SortSmallest(small, k, x, 1, n)
	fmt.Printf("%d smallest values from %d\n", k, n)
	for i := 0; i < k; i++ {
		fmt.Printf("%d: %.18f\n", i, small[i])
	}
}

func TestSortVectorIndex(t *testing.T) {
	var n int = 10000
	var k int = 5
	v := vector.VectorAlloc(n)
	p := permutation.PermutationAlloc(n)
	rng.EnvSetup()
	T := rng.DefaultRngType()
	r := rng.RngAlloc(T)
	for i := 0; i < n; i++ {
		vector.Set(v, i, rng.Uniform(r))
	}
	sort.SortVectorIndex(p, v)
	pData := p.Slice_().([]int)
	for i := 0; i < k; i++ {
		vpi := vector.Get(v, pData[i])
		fmt.Printf("order = %d, value = %g\n", i, vpi)
	}
}
