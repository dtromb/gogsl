package test

import (
	"fmt"
	"testing"
	"github.com/dtromb/gogsl/histogram"
	"github.com/dtromb/gogsl/rng"
)

func TestHistogram(t *testing.T) {

	h := histogram.Histogram2dAlloc(10,10)
	histogram.Histogram2dSetRangesUniform(h, 0.0, 1.0, 0.0, 1.0)
	histogram.Histogram2dAccumulate(h, 0.3, 0.3, 1)
	histogram.Histogram2dAccumulate(h, 0.8, 0.1, 5)
	histogram.Histogram2dAccumulate(h, 0.7, 0.9, 0.5)
	
	rng.EnvSetup()
	T := rng.DefaultRngType()
	r := rng.RngAlloc(T)
	
	hDim := h.Dim()
	p := histogram.Histogram2dPdfAlloc(hDim[0], hDim[1])
    histogram.Histogram2dPdfInit(p, h)
    for i := 0; i < 1000; i++ {
    	u := rng.Uniform(r);
    	v := rng.Uniform(r);
       	_, x, y := histogram.Histogram2dPdfSample(p, u, v)
        fmt.Printf ("%g %g\n", x, y);
    }
}