package test

import (
	"fmt"
	"github.com/dtromb/gogsl/blas"
	"github.com/dtromb/gogsl/matrix"
	"testing"
)

func TestBlas(t *testing.T) {
	a := []float64{0.11, 0.12, 0.13,
		0.21, 0.22, 0.23}
	b := []float64{1011, 1012,
		1021, 1022,
		1031, 1032}
	c := []float64{0.00, 0.00,
		0.00, 0.00}

	va := matrix.ViewArray(a, 2, 3)
	vb := matrix.ViewArray(b, 3, 2)
	vc := matrix.ViewArray(c, 2, 2)

	blas.Dgemm(blas.NoTrans, blas.NoTrans, 1.0, va.Matrix(), vb.Matrix(), 0.0, vc.Matrix())
	fmt.Printf("[ %g, %g\n", c[0], c[1])
	fmt.Printf("  %g, %g ]\n", c[2], c[3])
}
