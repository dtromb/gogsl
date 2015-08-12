package test

import (
	"fmt"
	"github.com/dtromb/gogsl"
	"github.com/dtromb/gogsl/multiset"
	"os"
	"testing"
)

func TestMultiset(t *testing.T) {
	fmt.Println("All multisets of {0,1,2,3} by size:\n")
	for i := 0; i <= 4; i++ {
		c := multiset.MultisetCalloc(4, i)
		for {
			fmt.Printf("{")
			multiset.Fprintf(os.Stdout, c, " %u")
			fmt.Printf(" }\n")
			r := gogsl.GslError(multiset.Next(c))
			if r != gogsl.GSL_SUCCESS {
				break
			}
		}
	}
}
