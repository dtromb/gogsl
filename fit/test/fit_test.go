package test

import (
	"fmt"
	"github.com/dtromb/gogsl/fit"
	"math"
	"testing"
)

func TestFit(t *testing.T) {

	var n int = 4
	var x []float64 = []float64{1970, 1980, 1990, 2000}
	var y []float64 = []float64{12, 11, 14, 13}
	var w []float64 = []float64{0.1, 0.2, 0.3, 0.4}

	_, c0, c1, cov00, cov01, cov11, chisq := fit.Wlinear(x, 1, w, 1, y, 1, n)
	fmt.Printf("# best fit: Y = %g + %g X\n", c0, c1)
	fmt.Printf("# covariance matrix:\n")
	fmt.Printf("# [ %g, %g\n#   %g, %g]\n", cov00, cov01, cov01, cov11)
	fmt.Printf("# chisq = %g\n", chisq)

	for i := 0; i < n; i++ {
		fmt.Printf("data: %g %g %g\n", x[i], y[i], 1/math.Sqrt(w[i]))
	}
	fmt.Print("\n")

	for i := -30; i < 130; i++ {
		xf := x[0] + (float64(i)/100.0)*(x[n-1]-x[0])
		_, yf, err := fit.LinearEst(xf, c0, c1, cov00, cov01, cov11)

		fmt.Printf("fit: %g %g\n", xf, yf)
		fmt.Printf("hi : %g %g\n", xf, yf+err)
		fmt.Printf("lo : %g %g\n", xf, yf-err)
	}
}
