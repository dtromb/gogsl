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
		C: -0.5,
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
