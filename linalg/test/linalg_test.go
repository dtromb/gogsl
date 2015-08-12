package test

import (
	"fmt"
	"github.com/dtromb/gogsl/linalg"
	"github.com/dtromb/gogsl/matrix"
	"github.com/dtromb/gogsl/permutation"
	"github.com/dtromb/gogsl/vector"
	"os"
	"testing"
)

func TestLinalg(t *testing.T) {

	aData := []float64{0.18, 0.60, 0.57, 0.96,
		0.41, 0.24, 0.99, 0.58,
		0.14, 0.30, 0.97, 0.66,
		0.51, 0.13, 0.19, 0.85}
	bData := []float64{1.0, 2.0, 3.0, 4.0}
	m := matrix.ViewArray(aData, 4, 4)
	b := vector.ViewArray(bData, 4)
	x := vector.VectorAlloc(4)
	p := permutation.PermutationAlloc(4)
	linalg.LUDecomp(m.Matrix(), p)
	linalg.LUSolve(m.Matrix(), p, b.Vector(), x)
	fmt.Printf("x = \n")
	vector.Fprintf(os.Stdout, x, "%g")
}
