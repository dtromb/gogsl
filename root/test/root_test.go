package test

import (
	"fmt"
	"github.com/dtromb/gogsl"
	"github.com/dtromb/gogsl/root"
	"math"
	"testing"
)

type QuadraticParams struct {
	A float64
	B float64
	C float64
}

func Quadratic(x float64, params interface{}) float64 {
	qp := params.(*QuadraticParams)
	return (qp.A*x+qp.B)*x + qp.C
}

func QuadraticDeriv(x float64, params interface{}) float64 {
	qp := params.(*QuadraticParams)
	return 2.0*qp.A*x + qp.B
}

func QuadraticFdf(x float64, params interface{}) (float64, float64) {
	qp := params.(*QuadraticParams)
	return (qp.A*x+qp.B)*x + qp.C, 2.0*qp.A*x + qp.B
}

func TestRoot(t *testing.T) {
	var maxIter int = 100
	var rExpected float64 = math.Sqrt(5.0)
	var xLo float64 = 0.0
	var xHi float64 = 5.0

	qp := &QuadraticParams{
		A: 1.0,
		B: 0.0,
		C: -5.0,
	}

	f := &gogsl.GslFunction{
		Function: Quadratic,
		Params:   qp,
	}
	gogsl.InitializeGslFunction(f)

	T := root.GSL_ROOT_FSOLVER_BRENT
	s := root.FsolverAlloc(T)
	root.FsolverSet(s, f, xLo, xHi)
	fmt.Printf("using %s method\n", root.FsolverName(s))
	fmt.Printf("%5s [%9s, %9s] %9s %10s %9s\n",
		"iter", "lower", "upper", "root",
		"err", "err(est)")
	iter := 0
	for {
		iter += 1
		root.FsolverIterate(s)
		r := root.FsolverRoot(s)
		xLo = root.FsolverXLower(s)
		xHi = root.FsolverXUpper(s)
		status := gogsl.GslError(root.TestInterval(xLo, xHi, 0, 0.001))
		if status == gogsl.GSL_SUCCESS {
			fmt.Printf("Converged:\n")
		}
		fmt.Printf("%5d [%.7f, %.7f] %.7f %+.7f %.7f\n",
			iter, xLo, xHi, r, r-rExpected, xHi-xLo)
		if status != gogsl.GSL_CONTINUE || iter >= maxIter {
			break
		}
	}
}

func TestRootFdf(t *testing.T) {

	var maxIter int = 100
	var x float64 = 5.0
	var rExpected float64 = math.Sqrt(5.0)

	qp := &QuadraticParams{
		A: 1.0,
		B: 0.0,
		C: -5.0,
	}

	f := &gogsl.GslFunctionFdf{
		Function:   Quadratic,
		Derivative: QuadraticDeriv,
		Fdf:        QuadraticFdf,
		Params:     qp,
	}
	gogsl.InitializeGslFunctionFdf(f)

	T := root.GSL_ROOT_FDFSOLVER_NEWTON
	s := root.FdfsolverAlloc(T)
	root.FdfsolverSet(s, f, x)

	fmt.Printf("using %s method\n", root.FdfsolverName(s))
	fmt.Printf("%-5s %10s %10s %10s\n", "iter", "root", "err", "err(est)")

	iter := 0
	for {
		iter += 1
		root.FdfsolverIterate(s)
		x0 := x
		x = root.FdfsolverRoot(s)
		status := gogsl.GslError(root.TestDelta(x, x0, 0, 1e-3))
		if status == gogsl.GSL_SUCCESS {
			fmt.Printf("Converged:\n")
		}
		fmt.Printf("%5d %10.7f %+10.7f %10.7f\n", iter, x, x-rExpected, x-x0)
		if status != gogsl.GSL_CONTINUE || iter >= maxIter {
			break
		}
	}
}
