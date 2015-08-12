package test

import (
	"fmt"
	"github.com/dtromb/gogsl"
	"github.com/dtromb/gogsl/combination"
	"os"
	"testing"
)

func TestCombination(t *testing.T) {
	maxiter := 25
	var iter int
	fmt.Printf("All subsets of {0,1,2,3} by size:\n")
	for i := 0; i <= 4; i++ {
		c := combination.CombinationCalloc(4, i)
		for {
			fmt.Printf("{")
			combination.Fprintf(os.Stdout, c, " %u")
			fmt.Printf(" }\n")
			r := gogsl.GslError(combination.Next(c))
			if r != gogsl.GSL_SUCCESS {
				break
			}
			if iter > maxiter {
				t.Error("iteration limit reached")
				break
			}
			iter += 1
		}
	}
}
