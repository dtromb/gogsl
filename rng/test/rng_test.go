package test

import (
	"fmt"
	"testing"
	"github.com/dtromb/gogsl/rng"
)

func TestRng(t *testing.T) {
	var n int = 10
	rng.EnvSetup()
	T := rng.DefaultRngType()
	r := rng.RngAlloc(T)
	for i := 0; i < n; i++ {
      u := rng.Uniform(r)
      fmt.Printf("%.5f\n", u);
	}
	fmt.Println()
}

