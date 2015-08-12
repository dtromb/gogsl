package poly

import (
	"fmt"
	"testing"
	//"runtime"
	//"time"
)

func TestPoly(t *testing.T) {
	res := Eval([]float64{1, 0, 2}, 3, 3)
	fmt.Printf("%f\n", res)

	w := ComplexWorkspaceAlloc(3)
	ret := make([]complex128, 2)
	r := ComplexSolve([]float64{1, 0, 2}, 3, w, ret)
	fmt.Printf("ret=%d, (%f+%fi),(%f+%fi)\n", r, real(ret[0]), imag(ret[0]), real(ret[0]), imag(ret[1]))
	//w = nil
	//runtime.GC()
	//<-time.NewTimer(time.Millisecond*time.Duration(2000)).C
}

func TestPoly2(t *testing.T) {
	/* coefficients of P(x) =  -1 + x^5  */
	a := []float64{-1, 0, 0, 0, 0, 1}
	z := make([]complex128, 5)
	w := ComplexWorkspaceAlloc(6)
	ComplexSolve(a, 6, w, z)
	for i := 0; i < 5; i++ {
		fmt.Printf("z%d = %+.18f %+.18f\n", i, real(z[i]), imag(z[i]))
	}
}
