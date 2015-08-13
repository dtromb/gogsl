package test

import (
	"fmt"
	"github.com/dtromb/gogsl/monte"
	"github.com/dtromb/gogsl/rng"
	"math"
	"testing"
)

var exact float64 = 1.3932039296856768591842462603255

func DisplayResults(title string, result float64, err float64) {
	fmt.Printf("%s ==================\n", title)
	fmt.Printf("result = % .6f\n", result)
	fmt.Printf("sigma  = % .6f\n", err)
	fmt.Printf("exact  = % .6f\n", exact)
	fmt.Printf("error  = % .6f = %.2g sigma\n", result-exact,
		math.Abs((result-exact)/err))
}

func g(k []float64, dim int, params interface{}) float64 {
	A := 1.0 / (math.Pi * math.Pi * math.Pi)
	return A / (1.0 - math.Cos(k[0])*math.Cos(k[1])*math.Cos(k[2]))
}

func TestMonte(t *testing.T) {

	xl := []float64{0, 0, 0}
	xu := []float64{math.Pi, math.Pi, math.Pi}

	G := &monte.GslMonteFunction{
		Function: g,
		Dim:      3,
	}
	monte.InitializeGslMonteFunction(G)

	var calls int = 500000

	rng.EnvSetup()
	T := rng.DefaultRngType()
	r := rng.RngAlloc(T)

	sp := monte.PlainAlloc(3)
	_, res, err := monte.PlainIntegrate(G, xl, xu, 3, calls, r, sp)
	DisplayResults("plain", res, err)

	sm := monte.MiserAlloc(3)
	_, res, err = monte.MiserIntegrate(G, xl, xu, 3, calls, r, sm)
	DisplayResults("miser", res, err)

	sv := monte.VegasAlloc(3)
	_, res, err = monte.VegasIntegrate(G, xl, xu, 3, 10000, r, sv)
	DisplayResults("vegas warm-up", res, err)

	fmt.Printf("converging...\n")
	for {
		_, res, err = monte.VegasIntegrate(G, xl, xu, 3, calls/5, r, sv)
		fmt.Printf("result = % .6f sigma = % .6f "+
			"chisq/dof = %.1f\n", res, err, monte.VegasChisq(sv))
		if math.Abs(monte.VegasChisq(sv))-1.0 <= 0.5 {
			break
		}
	}
}
