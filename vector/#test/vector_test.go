package test

import (
	"fmt"
	"testing"
	"github.com/dtromb/gogsl"
	"github.com/dtromb/gogsl/vector"
)

func TestVector(t *testing.T) {
	var stopDoingIt bool
	gogsl.SetErrorHandler(&gogsl.GslErrorHandler{
		Handler: func(reason string, file string, line int, gslError gogsl.GslError) {
			fmt.Println("Caught error: "+reason)
			stopDoingIt = true
		},
	})
	v := vector.VectorAlloc(3)
	for i := 0; i < 3; i++ {
    	vector.Set(v, i, float64(1.23) + float64(i))
    }
    for i := 0; i < 100; i++ { // OUT OF RANGE ERROR 
      	fmt.Printf("v_%d = %g\n", i, vector.Get(v, i));
		if stopDoingIt {
			break
		}
    }
}

