package test

import (
	"fmt"
	"github.com/dtromb/gogsl"
	"github.com/dtromb/gogsl/multimin"
	"github.com/dtromb/gogsl/vector"
	"testing"
)

/* Paraboloid centered on (p[0],p[1]), with
   scale factors (p[2],p[3]) and minimum p[4] */
func MyF(v *vector.GslVector, params interface{}) float64 {
	p := params.([]float64)
	x := vector.Get(v, 0)
	y := vector.Get(v, 1)
	return p[2]*(x-p[0])*(x-p[0]) +
		p[3]*(y-p[1])*(y-p[1]) +
		p[4]
}

func MyDf(v *vector.GslVector, params interface{}, df *vector.GslVector) {
	p := params.([]float64)
	x := vector.Get(v, 0)
	y := vector.Get(v, 1)
	vector.Set(df, 0, 2.0*p[2]*(x-p[0]))
	vector.Set(df, 1, 2.0*p[3]*(y-p[1]))
}

func MyFdf(v *vector.GslVector, params interface{}, df *vector.GslVector) float64 {
	MyDf(v, params, df)
	return MyF(v, params)
}

func TestMultimin(t *testing.T) {

	par := []float64{1.0, 2.0, 10.0, 20.0, 30.0}
	f := &multimin.GslMultiminFunction{
		Function:  MyF,
		Dimension: 2,
		Params:    par,
	}
	multimin.InitializeGslMultiminFunction(f)

	/* Starting point */
	x := vector.VectorAlloc(2)
	vector.Set(x, 0, 5.0)
	vector.Set(x, 1, 7.0)

	/* Set initial step sizes to 1 */
	ss := vector.VectorAlloc(2)
	vector.SetAll(ss, 1.0)

	T := multimin.GSL_MULTIMIN_FMINIMIZER_NSIMPLEX2
	s := multimin.FminimizerAlloc(T, 2)
	multimin.FminimizerSet(s, f, x, ss)

	iter := 0

	for {
		iter += 1
		status := gogsl.GslError(multimin.FminimizerIterate(s))
		if status != gogsl.GSL_SUCCESS {
			break
		}
		size := multimin.FminimizerSize(s)
		status = gogsl.GslError(multimin.TestSize(size, 1e-2))
		if status == gogsl.GSL_SUCCESS {
			fmt.Printf("converged to minimum at\n")
		}
		fmt.Printf("%5d %10.3e %10.3e f() = %7.3f size = %.3f\n",
			iter, vector.Get(s.X_(), 0), vector.Get(s.X_(), 1), s.Fval(), size)
		if status != gogsl.GSL_CONTINUE || iter >= 100 {
			break
		}
	}
}

func TestMultiminFdf(t *testing.T) {

	par := []float64{1.0, 2.0, 10.0, 20.0, 30.0}
	f := &multimin.GslMultiminFunctionFdf{
		Function:   MyF,
		Derivative: MyDf,
		Fdf:        MyFdf,
		Dimension:  2,
		Params:     par,
	}
	multimin.InitializeGslMultiminFunctionFdf(f)

	/* Starting point */
	x := vector.VectorAlloc(2)
	vector.Set(x, 0, 5.0)
	vector.Set(x, 1, 7.0)

	T := multimin.GSL_MULTIMIN_FDFMINIMIZER_CONJUGATE_FR
	s := multimin.FdfminimizerAlloc(T, 2)
	multimin.FdfminimizerSet(s, f, x, 0.01, 1e-4)

	iter := 0

	for {
		iter += 1
		status := gogsl.GslError(multimin.FdfminimizerIterate(s))
		if status != gogsl.GSL_SUCCESS {
			break
		}
		status = gogsl.GslError(multimin.TestGradient(s.Gradient_(), 1e-3))
		if status == gogsl.GSL_SUCCESS {
			fmt.Printf("Minimum found at:\n")
		}
		fmt.Printf("%5d %10.3e %10.3e f() = %7.3f\n",
			iter, vector.Get(s.X_(), 0), vector.Get(s.X_(), 1), s.Fval())
		if status != gogsl.GSL_CONTINUE || iter >= 100 {
			break
		}
	}
}
