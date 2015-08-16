package test

import (
	"fmt"
	"github.com/dtromb/gogsl/sum"
	"math"
	"testing"
)

func TestSum(t *testing.T) {

	var N int = 20
	T := make([]float64, N)

	w := sum.LevinUAlloc(N)

	var zeta2 float64 = math.Pi * math.Pi / 6.0
	var s float64

	for n := 0; n < N; n++ {
		np1 := float64(n) + 1.0
		T[n] = 1.0 / (np1 * np1)
		s += T[n]
	}

	_, sumAccel, err := sum.LevinUAccel(T, N, w)

	fmt.Printf("term-by-term sum = % .16f using %d terms\n", s, N)
	fmt.Printf("term-by-term sum = % .16f using %d terms\n", w.SumPlain(), w.TermsUsed())
	fmt.Printf("exact value      = % .16f\n", zeta2)
	fmt.Printf("accelerated sum  = % .16f using %d terms\n", sumAccel, w.TermsUsed())
	fmt.Printf("estimated error  = % .16f\n", err)
	fmt.Printf("actual error     = % .16f\n", sumAccel-zeta2)
}
