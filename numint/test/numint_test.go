package test

import (
	"fmt"
	"github.com/dtromb/gogsl"
	"github.com/dtromb/gogsl/numint"
	"math"
	"testing"
)

var w *numint.GslIntegrationWorkspace = numint.IntegrationWorkspaceAlloc(1000)

func F(x float64, params interface{}) float64 {
	alpha := params.(float64)
	f := math.Log(alpha*x) / math.Sqrt(x)
	return f
}

func TestNumint(t *testing.T) {
	gogsl.SetErrorHandler(&gogsl.GslErrorHandler{
		Handler: func(reason string, file string, line int, gslError gogsl.GslError) {
			fmt.Println(reason + " : " + gslError.String())
		},
	})
	expected := -4.0
	alpha := 1.0
	F := gogsl.GslFunction{Function: F, Params: alpha}
	status, result, err := numint.Qags(&F, 0, 1, 0, 0.0001, 1000, w)
	fmt.Printf("status          =  %s\n", gogsl.GslError(status).String())
	fmt.Printf("result          = % .18f\n", result)
	fmt.Printf("exact result    = % .18f\n", expected)
	fmt.Printf("estimated error = % .18f\n", err)
	fmt.Printf("actual error    = % .18f\n", result-expected)
	fmt.Printf("intervals =  %d\n", w.IntervalCount())
}
