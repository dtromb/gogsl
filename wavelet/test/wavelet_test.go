package test

import (
	"fmt"
	"github.com/dtromb/gogsl/sort"
	"github.com/dtromb/gogsl/wavelet"
	"math"
	"testing"
)

func TestWavelet(t *testing.T) {

	var n int = 256
	var nc int = 20

	data := make([]float64, n)
	for i := 0; i < n; i++ {
		x := float64(i) / float64(n) * 2.0 * math.Pi
		data[i] = (math.Sin(x) + math.Sin(2*x) + math.Sin(3*x) + math.Sin(4*x)) / 4.0
	}

	w := wavelet.WaveletAlloc(wavelet.GSL_WAVELET_DAUBECHIES, 4)
	work := wavelet.WaveletWorkspaceAlloc(n)

	data2 := make([]float64, n)
	copy(data2, data)
	wavelet.WaveletTransformForward(w, data2, 1, n, work)
	for i := 0; i < n; i++ {
		data2[i] = math.Abs(data2[i])
	}

	p := make([]int, n)
	sort.SortIndex(p, data2, 1, n)

	for i := 0; i < n; i++ {
		fmt.Printf("k[%d] = %g\n", p[i], data2[p[i]])
	}

	fmt.Printf("Truncating all but largest %d components...\n", nc)
	for i := 0; i+nc < n; i++ {
		data2[p[i]] = 0
	}

	wavelet.WaveletTransformInverse(w, data2, 1, n, work)
	for i := 0; i < n; i++ {
		x := float64(i) / float64(n) * 2.0 * math.Pi
		fmt.Printf("x=%g 	%g	%g	%g\n", x, data[i], data2[i], data[i]-data2[i])
	}
}
