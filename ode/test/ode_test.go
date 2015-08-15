package test

import (
	"fmt"
	"github.com/dtromb/gogsl"
	"github.com/dtromb/gogsl/matrix"
	"github.com/dtromb/gogsl/ode"
	"testing"
)

func F(t float64, y []float64, f []float64, params interface{}) int {
	mu := *params.(*float64)
	f[0] = y[1]
	f[1] = -y[0] - mu*y[1]*(y[0]*y[0]-1)
	return int(gogsl.GSL_SUCCESS)
}

func Df(t float64, y []float64, dfdy []float64, dfdt []float64, params interface{}) int {
	mu := *params.(*float64)
	dfdyMat := matrix.ViewArray(dfdy, 2, 2)
	m := dfdyMat.Matrix()
	matrix.Set(m, 0, 0, 0.0)
	matrix.Set(m, 0, 1, 1.0)
	matrix.Set(m, 1, 0, -2.0*mu*y[0]*y[1]-1.0)
	matrix.Set(m, 1, 1, -mu*(y[0]*y[0]-1.0))
	dfdt[0] = 0.0
	dfdt[1] = 0.0
	return int(gogsl.GSL_SUCCESS)
}

func TestOde(test *testing.T) {
	var mu float64 = 10
	sys := &ode.GslOdeiv2System{
		Function:  F,
		Jacobian:  Df,
		Dimension: 2,
		Params:    &mu,
	}
	ode.InitializeGslFOdeiv2System(sys)

	d := ode.DriverAllocYNew(sys, ode.ODEIV2_STEP_TYPE_RK8PD, 1e-6, 1e-6, 0.0)

	var t float64 = 0.0
	var t1 float64 = 100.0
	var y []float64 = []float64{1.0, 0.0}

	for i := 1; i <= 100; i++ {
		ti := float64(i) * t1 / 100.0
		status := ode.DriverApply(d, &t, ti, y)
		if status != int32(gogsl.GSL_SUCCESS) {
			test.Error("error, return " + gogsl.GslError(status).String())
			break
		}
		fmt.Printf("%.5e %.5e %.5e\n", t, y[0], y[1])
	}
}
