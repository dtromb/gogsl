package test

import (
	"fmt"
	"github.com/dtromb/gogsl/eigen"
	"github.com/dtromb/gogsl/matrix"
	"github.com/dtromb/gogsl/vector"
	"os"
	"testing"
)

func TestEigen(t *testing.T) {

	data := []float64{1.0, 1 / 2.0, 1 / 3.0, 1 / 4.0,
		1 / 2.0, 1 / 3.0, 1 / 4.0, 1 / 5.0,
		1 / 3.0, 1 / 4.0, 1 / 5.0, 1 / 6.0,
		1 / 4.0, 1 / 5.0, 1 / 6.0, 1 / 7.0}
	m := matrix.ViewArray(data, 4, 4)
	eval := vector.VectorAlloc(4)
	evec := matrix.MatrixAlloc(4, 4)
	w := eigen.SymmvAlloc(4)
	eigen.Symmv(m.Matrix(), eval, evec, w)
	eigen.SymmvSort(eval, evec, eigen.EIGEN_SORT_ABS_ASC)
	for i := 0; i < 4; i++ {
		evalI := vector.Get(eval, i)
		evecI := matrix.Column(evec, i)
		fmt.Printf("eigenvalue = %g\n", evalI)
		fmt.Printf("eigenvector = \n")
		vector.Fprintf(os.Stdout, evecI.Vector(), "%g")
	}
}

func TestEigenNonsymm(t *testing.T) {
	data := []float64{-1.0, 1.0, -1.0, 1.0,
		-8.0, 4.0, -2.0, 1.0,
		27.0, 9.0, 3.0, 1.0,
		64.0, 16.0, 4.0, 1.0}

	m := matrix.ViewArray(data, 4, 4)
	eval := vector.VectorComplexAlloc(4)
	evec := matrix.MatrixComplexAlloc(4, 4)

	w := eigen.NonsymmvAlloc(4)
	eigen.Nonsymmv(m.Matrix(), eval, evec, w)
	eigen.NonsymmvSort(eval, evec, eigen.EIGEN_SORT_ABS_DESC)

	for i := 0; i < 4; i++ {
		evalI := vector.ComplexGet(eval, i)
		evecI := matrix.ComplexColumn(evec, i)
		fmt.Printf("eigenvalue = %g + %gi\n",
			real(evalI), imag(evalI))
		fmt.Printf("eigenvector = \n")
		for j := 0; j < 4; j++ {
			z := vector.ComplexGet(evecI.Vector(), j)
			fmt.Printf("%g + %gi\n", real(z), imag(z))
		}
	}
}
