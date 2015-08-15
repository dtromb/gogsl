package test

import (
	"fmt"
	"github.com/dtromb/gogsl/rng"
	"github.com/dtromb/gogsl/siman"
	"math"
	"testing"
)

/* how many points do we try before stepping */
var N_TRIES int = 200

/* how many iterations for each T? */
var ITERS_FIXED_T int = 1000

/* max step size in random walk */
var STEP_SIZE float64 = 1.0

/* Boltzmann constant */
var K float64 = 1.0

/* initial temperature */
var T_INITIAL float64 = 0.008

/* damping factor for temperature */
var MU_T float64 = 1.003
var T_MIN float64 = 2.0e-6

/* now some functions to test in one dimension */
func E1(params interface{}) float64 {
	x := *params.(*float64)
	res := math.Exp(-math.Pow((x-1.0), 2.0)) * math.Sin(8*x)
	//fmt.Printf("E(%f) = %g\n", x, res)
	return res
}

func M1(xp interface{}, yp interface{}) float64 {
	x := *xp.(*float64)
	y := *yp.(*float64)
	return math.Abs(x - y)
}

func S1(r *rng.GslRng, xp interface{}, stepSize float64) {
	ptr := xp.(*float64)
	oldX := *ptr
	u := rng.Uniform(r)
	newX := u*2*stepSize - stepSize + oldX
	*ptr = newX
}

func P1(xp interface{}) {
	// XXX - Need to resolve this stdio / go io thing.
	siman.Print(fmt.Sprintf(" %12g", *xp.(*float64)))
}

func TestSiman(t *testing.T) {

	params := &siman.GslSimanParams{
		NumTries:   N_TRIES,
		ItersFixed: ITERS_FIXED_T,
		StepSize:   STEP_SIZE,
		K:          K,
		TInitial:   T_INITIAL,
		Mu:         MU_T,
		TMin:       T_MIN,
	}
	siman.InitializeGslSimanParams(params)

	var xInitial float64 = 15.5

	rng.EnvSetup()
	T := rng.DefaultRngType()
	r := rng.RngAlloc(T)
	fmt.Println(rng.Uniform(r))
	siman.Solve(r, &xInitial, E1, S1, M1, P1, nil, nil, nil, params)
}
