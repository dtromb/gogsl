package test

import (
	"fmt"
	"github.com/dtromb/gogsl"
	"github.com/dtromb/gogsl/matrix"
	"github.com/dtromb/gogsl/multiroot"
	"github.com/dtromb/gogsl/vector"
	"testing"
)

type Rparams struct {
	A float64
	B float64
}

func PrintState(iter int, s *multiroot.GslMultirootFsolver) {
	fmt.Printf("iter = %3d x = % .3f % .3f f(x) = % .3e % .3e\n",
		iter,
		vector.Get(s.X_(), 0),
		vector.Get(s.X_(), 1),
		vector.Get(s.F_(), 0),
		vector.Get(s.F_(), 1))
}

func PrintStateFdf(iter int, s *multiroot.GslMultirootFdfsolver) {
	fmt.Printf("iter = %3d x = % .3f % .3f f(x) = % .3e % .3e\n",
		iter,
		vector.Get(s.X_(), 0),
		vector.Get(s.X_(), 1),
		vector.Get(s.F_(), 0),
		vector.Get(s.F_(), 1))
}

func Rosenbrock(x *vector.GslVector, params interface{}, f *vector.GslVector) int {
	rp := params.(*Rparams)
	x0 := vector.Get(x, 0)
	x1 := vector.Get(x, 1)
	y0 := rp.A * (1.0 - x0)
	y1 := rp.B * (x1 - x0*x0)
	vector.Set(f, 0, y0)
	vector.Set(f, 1, y1)
	return int(gogsl.GSL_SUCCESS)
}

func RosenbrockDf(x *vector.GslVector, params interface{}, j *matrix.GslMatrix) int {
	rp := params.(*Rparams)
	x0 := vector.Get(x, 0)
	matrix.Set(j, 0, 0, -rp.A)
	matrix.Set(j, 0, 1, 0)
	matrix.Set(j, 1, 0, -2*rp.B*x0)
	matrix.Set(j, 1, 1, rp.B)
	return int(gogsl.GSL_SUCCESS)
}

func RosenbrockFdf(x *vector.GslVector, params interface{}, f *vector.GslVector, j *matrix.GslMatrix) int {
	Rosenbrock(x, params, f)
	RosenbrockDf(x, params, j)
	return int(gogsl.GSL_SUCCESS)
}

func TestMultiroot(t *testing.T) {

	var n int = 2

	rp := &Rparams{
		A: 1.0,
		B: 10.0,
	}

	f := &multiroot.GslMultirootFunction{
		Function:  Rosenbrock,
		Dimension: n,
		Params:    rp,
	}
	multiroot.InitializeGslMultirootFunction(f)

	xInit := []float64{-10.0, -5.0}
	x := vector.VectorAlloc(n)

	vector.Set(x, 0, xInit[0])
	vector.Set(x, 1, xInit[1])

	T := multiroot.GSL_MULTIROOT_FSOLVER_HYBRIDS
	s := multiroot.FsolverAlloc(T, 2)
	multiroot.FsolverSet(s, f, x)

	var iter int

	PrintState(iter, s)

	for {
		iter += 1
		status := gogsl.GslError(multiroot.FsolverIterate(s))
		PrintState(iter, s)
		if status != gogsl.GSL_SUCCESS {
			break
		}
		status = gogsl.GslError(multiroot.TestResidual(s.F_(), 1e-7))
		if status != gogsl.GSL_CONTINUE || iter >= 1000 {
			break
		}
		fmt.Printf("status = %s\n", status.String())
	}
}

func TestMultirootFdf(t *testing.T) {

	var n int = 2

	rp := &Rparams{
		A: 1.0,
		B: 10.0,
	}

	f := &multiroot.GslMultirootFunctionFdf{
		Function:   Rosenbrock,
		Derivative: RosenbrockDf,
		Fdf:        RosenbrockFdf,
		Dimension:  n,
		Params:     rp,
	}
	multiroot.InitializeGslMultirootFunctionFdf(f)

	xInit := []float64{-10.0, -5.0}
	x := vector.VectorAlloc(n)

	vector.Set(x, 0, xInit[0])
	vector.Set(x, 1, xInit[1])

	T := multiroot.GSL_MULTIROOT_FDFSOLVER_GNEWTON
	s := multiroot.FdfsolverAlloc(T, 2)
	multiroot.FdfsolverSet(s, f, x)

	var iter int

	PrintStateFdf(iter, s)

	for {
		iter += 1
		status := gogsl.GslError(multiroot.FdfsolverIterate(s))
		PrintStateFdf(iter, s)
		if status != gogsl.GSL_SUCCESS {
			break
		}
		status = gogsl.GslError(multiroot.TestResidual(s.F_(), 1e-7))
		if status != gogsl.GSL_CONTINUE || iter >= 1000 {
			break
		}
		fmt.Printf("status = %s\n", status.String())
	}
}
