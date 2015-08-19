package test

import (
	"fmt"
	"github.com/dtromb/gogsl"
	"github.com/dtromb/gogsl/blas"
	"github.com/dtromb/gogsl/matrix"
	"github.com/dtromb/gogsl/multifit"
	"github.com/dtromb/gogsl/randist"
	"github.com/dtromb/gogsl/rng"
	"github.com/dtromb/gogsl/vector"
	"math"
	"testing"
)

type Data struct {
	N     int
	y     []float64
	sigma []float64
}

func ExpbF(x *vector.GslVector, params interface{}, f *vector.GslVector) int {

	data := params.(*Data)

	A := vector.Get(x, 0)
	lambda := vector.Get(x, 1)
	b := vector.Get(x, 2)

	for i := 0; i < data.N; i++ {
		//  Model Yi = A * exp(-lambda * i) + b
		yi := A*math.Exp(-lambda*float64(i)) + b
		vector.Set(f, i, (yi-data.y[i])/data.sigma[i])
	}

	return int(gogsl.GSL_SUCCESS)
}

func ExpbDf(x *vector.GslVector, params interface{}, j *matrix.GslMatrix) int {

	data := params.(*Data)

	A := vector.Get(x, 0)
	lambda := vector.Get(x, 1)

	for i := 0; i < data.N; i++ {
		e := math.Exp(-lambda * float64(i))
		matrix.Set(j, i, 0, e/data.sigma[i])
		matrix.Set(j, i, 1, -float64(i)*A*e/data.sigma[i])
		matrix.Set(j, i, 2, 1/data.sigma[i])
	}

	return int(gogsl.GSL_SUCCESS)
}

func ExpbFdf(x *vector.GslVector, params interface{}, f *vector.GslVector, j *matrix.GslMatrix) int {
	ExpbF(x, params, f)
	ExpbDf(x, params, j)
	return int(gogsl.GSL_SUCCESS)
}

func PrintState(iter int, s *multifit.GslMultifitFdfsolver) {
	fmt.Printf("iter: %3d x = % 15.8f % 15.8f % 15.8f |f(x)| = %g\n",
		iter, vector.Get(s.X_(), 0), vector.Get(s.X_(), 1),
		vector.Get(s.X_(), 2), blas.Dnrm2(s.F_()))
}

func TestMultifitNonlinear(t *testing.T) {

	var N int = 40
	var p int = 3

	covar := matrix.MatrixAlloc(p, p)
	xInit := []float64{1.0, 0.0, 0.0}
	x := vector.ViewArray(xInit, p)
	y := make([]float64, N)

	sigma := make([]float64, N)

	data := &Data{
		N:     N,
		y:     y,
		sigma: sigma,
	}

	rng.EnvSetup()
	T := rng.DefaultRngType()
	r := rng.RngAlloc(T)

	f := &multifit.GslMultifitFunctionFdf{
		Function:   ExpbF,
		Derivative: ExpbDf,
		Fdf:        ExpbFdf,
		Dimension:  N,
		Params:     data,
	}
	multifit.InitializeGslMultifitFunctionFdf(f)

	// This is the data to be fitted
	for i := 0; i < N; i++ {
		y[i] = 1.0 + 5*math.Exp(-0.1*float64(i)) + randist.Gaussian(r, 0.1)
		sigma[i] = 0.1
		fmt.Printf("data: %d %g %g\n", i, y[i], sigma[i])
	}

	st := multifit.GSL_MULTIFIT_FDFSOLVER_LMSDER
	s := multifit.FdfsolverAlloc(st, N, p)
	multifit.FdfsolverSet(s, f, x.Vector())

	var status gogsl.GslError
	var iter int
	PrintState(iter, s)

	for {
		iter += 1
		status = gogsl.GslError(multifit.FdfsolverIterate(s))
		fmt.Printf("status = %s\n", status.String())
		PrintState(iter, s)
		if status != gogsl.GSL_SUCCESS {
			break
		}
		status = gogsl.GslError(multifit.TestDelta(s.Dx_(), s.X_(), 1e-4, 1e-4))
		if status != gogsl.GSL_CONTINUE || iter >= 500 {
			break
		}
	}

	multifit.Covar(s.J_(), 0.0, covar)

	chi := blas.Dnrm2(s.F_())
	dof := N - p
	c := math.Max(1, chi/math.Sqrt(float64(dof)))

	fmt.Printf("chisq/dof = %g\n", math.Pow(chi, 2.0)/float64(dof))
	fmt.Printf("A      = %.5f +/- %.5f\n", vector.Get(s.X_(), 0), c*math.Sqrt(matrix.Get(covar, 0, 0)))
	fmt.Printf("lambda = %.5f +/- %.5f\n", vector.Get(s.X_(), 1), c*math.Sqrt(matrix.Get(covar, 1, 1)))
	fmt.Printf("b      = %.5f +/- %.5f\n", vector.Get(s.X_(), 2), c*math.Sqrt(matrix.Get(covar, 2, 2)))
	fmt.Printf("status = %s\n", status.String())
}
