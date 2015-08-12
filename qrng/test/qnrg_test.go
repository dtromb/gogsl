package test

import (
	"fmt"
	"github.com/dtromb/gogsl/qrng"
	"testing"
)

func TestQnrg(t *testing.T) {
	q := qrng.QrngAlloc(qrng.QRNG_SOBOL, 2)
	v := make([]float64, 2)
	for i := 0; i < 1024; i++ {
		qrng.Get(q, v)
		fmt.Printf("%.5f %.5f\n", v[0], v[1])
	}
}
