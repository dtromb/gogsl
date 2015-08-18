package min

/*
	#include <gsl/gsl_min.h>
*/
import "C"

import "unsafe"

var GSL_MIN_FMINIMIZER_GOLDENSECTION *GslMinFminimizerType = &GslMinFminimizerType{cPtr: uintptr(unsafe.Pointer(C.gsl_min_fminimizer_goldensection))}
var GSL_MIN_FMINIMIZER_BRENT *GslMinFminimizerType = &GslMinFminimizerType{cPtr: uintptr(unsafe.Pointer(C.gsl_min_fminimizer_brent))}
var GSL_MIN_FMINIMIZER_QUAD_GOLDEN *GslMinFminimizerType = &GslMinFminimizerType{cPtr: uintptr(unsafe.Pointer(C.gsl_min_fminimizer_quad_golden))}

func (mt *GslMinFminimizerType) Ptr() uintptr {
	return mt.cPtr
}

func (mt *GslMinFminimizerType) Dispose() {
	mt.cPtr = 0
}
