package test

import (
	"fmt"
	"github.com/dtromb/gogsl/bspline"
	"github.com/dtromb/gogsl/matrix"
	"github.com/dtromb/gogsl/multifit"
	"github.com/dtromb/gogsl/randist"
	"github.com/dtromb/gogsl/rng"
	"github.com/dtromb/gogsl/stats"
	"github.com/dtromb/gogsl/vector"
	"math"
	"testing"
)

func TestBspline(t *testing.T) {

	var n int = 200
	var ncoeffs int = 12
	var nbreak int = ncoeffs - 2

	rng.EnvSetup()
	r := rng.RngAlloc(rng.DefaultRngType())

	// allocate a cubic bspline workspace (k = 4)
	bw := bspline.Alloc(4, nbreak)
	B := vector.VectorAlloc(ncoeffs)

	x := vector.VectorAlloc(n)
	y := vector.VectorAlloc(n)
	X := matrix.MatrixAlloc(n, ncoeffs)
	c := vector.VectorAlloc(ncoeffs)
	w := vector.VectorAlloc(n)
	cov := matrix.MatrixAlloc(ncoeffs, ncoeffs)
	mw := multifit.LinearAlloc(n, ncoeffs)

	fmt.Printf("#m=0,S=0\n")
	// this is the data to be fitted
	for i := 0; i < n; i++ {
		xi := (15.0 / (float64(n) - 1)) * float64(i)
		yi := math.Cos(xi) * math.Exp(-0.1*xi)
		sigma := 0.1 * yi
		dy := randist.Gaussian(r, sigma)
		vector.Set(x, i, xi)
		vector.Set(y, i, yi+dy)
		vector.Set(w, i, 1.0/(sigma*sigma))
		fmt.Printf("%f %f\n", xi, yi+dy)
	}

	// use uniform breakpoints on [0, 15]
	bspline.KnotsUniform(0.0, 15.0, bw)

	// construct the fit matrix X
	for i := 0; i < n; i++ {
		xi := vector.Get(x, i)

		// compute B_j(xi) for all j
		bspline.Eval(xi, B, bw)

		// fill in row i of X
		for j := 0; j < ncoeffs; j++ {
			matrix.Set(X, i, j, vector.Get(B, j))
		}
	}

	// do the fit
	_, chisq := multifit.Wlinear(X, w, y, c, cov, mw)
	dof := float64(n - ncoeffs)
	tss := stats.Wtss(w.Data_(), w.Stride(), y.Data_(), y.Stride(), n)
	rsq := 1.0 - chisq/tss
	fmt.Printf("chisq/dof = %e, Rsq = %f\n", chisq/dof, rsq)

	fmt.Printf("#m=1,S=0\n")
	for xi := 0.0; xi < 15.0; xi += 0.1 {
		bspline.Eval(xi, B, bw)
		_, yi, _ := multifit.LinearEst(B, c, cov)
		fmt.Printf("%f %f\n", xi, yi)
	}
}
