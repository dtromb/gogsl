package wavelet

/*
	#include <gsl/gsl_wavelet.h>
*/
import "C"

import "unsafe"

type GslWaveletDirection int

const (
	GSL_WAVELET_DIRECTION_FORWARD  GslWaveletDirection = +1
	GSL_WAVELET_DIRECTION_BACKWARD GslWaveletDirection = -1
)

var GSL_WAVELET_DAUBECHIES *GslWaveletType = &GslWaveletType{cPtr: uintptr(unsafe.Pointer(C.gsl_wavelet_daubechies))}
var GSL_WAVELET_DAUBECHIES_CENTERED *GslWaveletType = &GslWaveletType{cPtr: uintptr(unsafe.Pointer(C.gsl_wavelet_daubechies_centered))}
var GSL_WAVELET_HAAR *GslWaveletType = &GslWaveletType{cPtr: uintptr(unsafe.Pointer(C.gsl_wavelet_haar))}
var GSL_WAVELET_HAAR_CENTERED *GslWaveletType = &GslWaveletType{cPtr: uintptr(unsafe.Pointer(C.gsl_wavelet_haar_centered))}
var GSL_WAVELET_BSPLINE *GslWaveletType = &GslWaveletType{cPtr: uintptr(unsafe.Pointer(C.gsl_wavelet_bspline))}
var GSL_WAVELET_BSPLINE_CENTERED *GslWaveletType = &GslWaveletType{cPtr: uintptr(unsafe.Pointer(C.gsl_wavelet_bspline_centered))}

func (wt *GslWaveletType) Ptr() uintptr {
	return wt.cPtr
}

func (wt *GslWaveletType) Dispose() {
	wt.cPtr = 0
}
