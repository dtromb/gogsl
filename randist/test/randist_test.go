package test

import (
	"fmt"
	"github.com/dtromb/gogsl/randist"
	"github.com/dtromb/gogsl/rng"
	"testing"
)

func TestRandist(t *testing.T) {
	var n int = 10
	var mu float64 = 3.0
	rng.EnvSetup()
	T := rng.DefaultRngType()
	r := rng.RngAlloc(T)
	for i := 0; i < n; i++ {
		k := randist.Poisson(r, mu)
		fmt.Printf(" %d", k)
	}
	fmt.Println()
}
