package test

import (
	"fmt"
	"github.com/dtromb/gogsl/matrix"
	"github.com/dtromb/gogsl/multifit"
	"github.com/dtromb/gogsl/randist"
	"github.com/dtromb/gogsl/rng"
	"github.com/dtromb/gogsl/vector"
	"math"
	"testing"
)

func MakeData(start float64, end float64, inc float64) []float64 {
	var data []float64
	rng.EnvSetup()
	T := rng.DefaultRngType()
	r := rng.RngAlloc(T)
	for x := start; x < end; x += inc {
		y0 := math.Exp(x)
		sigma := inc * y0
		dy := randist.Gaussian(r, sigma)
		data = append(data, []float64{x, y0 + dy, sigma}...)
	}
	return data
}

func TestWeightedQuadtricFit(t *testing.T) {

	var n int = 19

	data := MakeData(0.1, 2.0, 0.1)

	X := matrix.MatrixAlloc(n, 3)
	y := vector.VectorAlloc(n)
	w := vector.VectorAlloc(n)

	c := vector.VectorAlloc(3)
	cov := matrix.MatrixAlloc(3, 3)

	for i := 0; i < n; i++ {
		xi := data[i*3]
		yi := data[i*3+1]
		ei := data[i*3+2]
		fmt.Printf("%g %g +/- %g\n", xi, yi, ei)

		matrix.Set(X, i, 0, 1.0)
		matrix.Set(X, i, 1, xi)
		matrix.Set(X, i, 2, xi*xi)

		vector.Set(y, i, yi)
		vector.Set(w, i, 1.0/(ei*ei))
	}

	work := multifit.LinearAlloc(n, 3)
	_, chisq := multifit.Wlinear(X, w, y, c, cov, work)

	fmt.Printf("# best fit: Y = %g + %g X + %g X^2\n",
		vector.Get(c, 0), vector.Get(c, 1), vector.Get(c, 2))
	fmt.Printf("# covariance matrix:\n")
	fmt.Printf("[ %+.5e, %+.5e, %+.5e  \n",
		matrix.Get(cov, 0, 0), matrix.Get(cov, 0, 1), matrix.Get(cov, 0, 2))
	fmt.Printf("  %+.5e, %+.5e, %+.5e  \n",
		matrix.Get(cov, 1, 0), matrix.Get(cov, 1, 1), matrix.Get(cov, 1, 2))
	fmt.Printf("  %+.5e, %+.5e, %+.5e ]\n",
		matrix.Get(cov, 2, 0), matrix.Get(cov, 2, 1), matrix.Get(cov, 2, 2))
	fmt.Printf("# chisq = %g\n", chisq)
}

func DoFit(t *multifit.GslMultifitRobustType, x *matrix.GslMatrix, y *vector.GslVector,
	c *vector.GslVector, cov *matrix.GslMatrix) {
	dim := x.Dim()
	work := multifit.RobustAlloc(t, dim[0], dim[1])
	multifit.Robust(x, y, c, cov, work)
}

func TestRobust(t *testing.T) {
	var p int = 2 // linear fit

	var a float64 = 1.45 // data slope
	var b float64 = 3.88 // data intercept

	var n int = 20

	X := matrix.MatrixAlloc(n, p)
	x := vector.VectorAlloc(n)
	y := vector.VectorAlloc(n)

	c := vector.VectorAlloc(p)
	cOls := vector.VectorAlloc(p)
	cov := matrix.MatrixAlloc(p, p)

	r := rng.RngAlloc(rng.DefaultRngType())

	// generate linear dataset
	for i := 0; i < n-3; i++ {
		dx := 10.0 / (float64(n) - 1.0)
		ei := rng.Uniform(r)
		xi := -5.0 + float64(i)*dx
		yi := a*xi + b

		vector.Set(x, i, xi)
		vector.Set(y, i, yi+ei)
	}

	// add a few outliers
	vector.Set(x, n-3, 4.7)
	vector.Set(y, n-3, -8.3)

	vector.Set(x, n-2, 3.5)
	vector.Set(y, n-2, -6.7)

	vector.Set(x, n-1, 4.1)
	vector.Set(y, n-1, -6.0)

	// construct design matrix X for linear fit
	for i := 0; i < n; i++ {
		xi := vector.Get(x, i)
		matrix.Set(X, i, 0, 1.0)
		matrix.Set(X, i, 1, xi)
	}

	// perform robust and OLS fit
	DoFit(multifit.GSL_MULTIFIT_ROBUST_OLS, X, y, cOls, cov)
	DoFit(multifit.GSL_MULTIFIT_ROBUST_BISQUARE, X, y, c, cov)

	// output data and model
	for i := 0; i < n; i++ {
		xi := vector.Get(x, i)
		yi := vector.Get(y, i)
		v := matrix.Row(X, i).Vector()
		_, yRob, _ := multifit.RobustEst(v, c, cov)
		_, yOls, _ := multifit.RobustEst(v, cOls, cov)

		fmt.Printf("%g %g %g %g\n", xi, yi, yRob, yOls)
	}

	fmt.Printf("# best fit: Y = %g + %g X\n", vector.Get(c, 0), vector.Get(c, 1))
	fmt.Printf("# covariance matrix:\n")
	fmt.Printf("# [ %+.5e, %+.5e\n", matrix.Get(cov, 0, 0), matrix.Get(cov, 0, 1))
	fmt.Printf("#   %+.5e, %+.5e\n", matrix.Get(cov, 1, 0), matrix.Get(cov, 1, 1))
}
