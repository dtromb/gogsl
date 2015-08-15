package test

import (
	"fmt"
	"github.com/dtromb/gogsl"
	"github.com/dtromb/gogsl/deriv"
	"math"
	"testing"
)

func F(x float64, params interface{}) float64 {
	return math.Pow(x, 1.5)
}

func TestDeriv(t *testing.T) {

	f := &gogsl.GslFunction{
		Function: F,
	}
	gogsl.InitializeGslFunction(f)

	fmt.Printf("f(x) = x^(3/2)\n")
	_, res, err := deriv.Central(f, 2.0, 1e-8)
	fmt.Printf("x = 2.0\n")
	fmt.Printf("f'(x) = %.10f +/- %.10f\n", res, err)
	fmt.Printf("exact = %.10f\n\n", 1.5*math.Sqrt(2.0))

	_, res, err = deriv.Forward(f, 0.0, 1e-8)
	fmt.Printf("x = 0.0\n")
	fmt.Printf("f'(x) = %.10f +/- %.10f\n", res, err)
	fmt.Printf("exact = %.10f\n\n", 0.0)
}
