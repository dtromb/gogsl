package test

import (
	"fmt"
	"github.com/dtromb/gogsl/interp"
	"math"
	"testing"
)

func TestInterp(t *testing.T) {

	var xi float64
	var yi float64

	x := make([]float64, 10)
	y := make([]float64, 10)

	fmt.Printf("#m=0,S=2\n")
	for i := 0; i < 10; i++ {
		x[i] = float64(i) + 0.5*math.Sin(float64(i))
		y[i] = float64(i) + math.Cos(float64(i*i))
		fmt.Printf("%g %g\n", x[i], y[i])
	}

	fmt.Printf("#m=1,S=0\n")
	acc := interp.AccelAlloc()
	spline := interp.SplineAlloc(interp.GSL_INTERP_CSPLINE, 10)
	interp.SplineInit(spline, x, y, 10)
	for xi = x[0]; xi < x[9]; xi += 0.01 {
		yi = interp.SplineEval(spline, xi, acc)
		fmt.Printf("%g %g\n", xi, yi)
	}
}

func TestInterp2(t *testing.T) {

	var N int = 4
	var x []float64 = []float64{0.00, 0.10, 0.27, 0.30}
	var y []float64 = []float64{0.15, 0.70, -0.10, 0.15}
	// Note: y[0] == y[3] for periodic data

	acc := interp.AccelAlloc()
	stype := interp.GSL_INTERP_CSPLINE_PERIODIC
	spline := interp.SplineAlloc(stype, N)

	var xi float64
	var yi float64

	fmt.Printf("#m=0,S=5\n")
	for i := 0; i < N; i++ {
		fmt.Printf("%g %g\n", x[i], y[i])
	}

	fmt.Printf("#m=1,S=0\n")
	interp.SplineInit(spline, x, y, N)
	for i := 0; i <= 100; i++ {
		xi = (1-float64(i)/100.0)*x[0] + (float64(i)/100.0)*x[N-1]
		yi = interp.SplineEval(spline, xi, acc)
		fmt.Printf("%g %g\n", xi, yi)
	}
}
