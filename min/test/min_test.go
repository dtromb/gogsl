package test

import (
	"fmt"
	"github.com/dtromb/gogsl"
	"github.com/dtromb/gogsl/min"
	"math"
	"testing"
)

func Fn1(x float64, params interface{}) float64 {
	return math.Cos(x) + 1.0
}

func TestMin(t *testing.T) {
	var maxIter int = 100
	var m float64 = 2.0
	var mExpected float64 = math.Pi
	var a float64 = 0.0
	var b float64 = 6.0

	var iter int

	f := &gogsl.GslFunction{
		Function: Fn1,
	}
	gogsl.InitializeGslFunction(f)

	T := min.GSL_MIN_FMINIMIZER_BRENT
	s := min.FminimizerAlloc(T)
	min.FminimizerSet(s, f, m, a, b)

	fmt.Printf("using %s method\n", min.FminimizerName(s))
	fmt.Printf("%5s [%9s, %9s] %9s %10s %9s\n", "iter", "lower", "upper", "min", "err", "err(est)")
	fmt.Printf("%5d [%.7f, %.7f] %.7f %+.7f %.7f\n", iter, a, b, m, m-mExpected, b-a)

	for {

		iter += 1
		min.FminimizerIterate(s)

		m = min.FminimizerXMinimum(s)
		a = min.FminimizerXLower(s)
		b = min.FminimizerXUpper(s)

		status := gogsl.GslError(min.TestInterval(a, b, 0.001, 0.0))
		if status == gogsl.GSL_SUCCESS {
			fmt.Printf("Converged:\n")
		}

		fmt.Printf("%5d [%.7f, %.7f] %.7f %+.7f %.7f\n", iter, a, b, m, m-mExpected, b-a)

		if status != gogsl.GSL_CONTINUE || iter >= maxIter {
			break
		}
	}
}
