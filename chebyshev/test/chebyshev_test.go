package test

import (
	"fmt"
	"github.com/dtromb/gogsl"
	"github.com/dtromb/gogsl/chebyshev"
	"testing"
)

func F(x float64, p interface{}) float64 {
	if x < 0.5 {
		return 0.25
	} else {
		return 0.75
	}
}

func TestChebyshev(t *testing.T) {

	var n int = 10000

	cs := chebyshev.Alloc(40)

	f := &gogsl.GslFunction{
		Function: F,
	}
	gogsl.InitializeGslFunction(f)
	chebyshev.Init(cs, f, 0.0, 1.0)

	for i := 0; i < n; i++ {
		x := float64(i) / float64(n)
		r10 := chebyshev.EvalN(cs, 10, x)
		r40 := chebyshev.Eval(cs, x)
		fmt.Printf("%g %g %g %g\n", x, f.Eval(x), r10, r40)
	}
}
