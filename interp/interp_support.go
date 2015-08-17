package interp

/*
	#include <gsl/gsl_interp.h>
*/
import "C"

import "unsafe"

var GSL_INTERP_LINEAR *GslInterpType = &GslInterpType{cPtr: uintptr(unsafe.Pointer(C.gsl_interp_linear))}
var GSL_INTERP_POLYNOMIAL *GslInterpType = &GslInterpType{cPtr: uintptr(unsafe.Pointer(C.gsl_interp_polynomial))}
var GSL_INTERP_CSPLINE *GslInterpType = &GslInterpType{cPtr: uintptr(unsafe.Pointer(C.gsl_interp_cspline))}
var GSL_INTERP_CSPLINE_PERIODIC *GslInterpType = &GslInterpType{cPtr: uintptr(unsafe.Pointer(C.gsl_interp_cspline_periodic))}
var GSL_INTERP_AKIMA *GslInterpType = &GslInterpType{cPtr: uintptr(unsafe.Pointer(C.gsl_interp_akima))}
var GSL_INTERP_AKIMA_PERIODIC *GslInterpType = &GslInterpType{cPtr: uintptr(unsafe.Pointer(C.gsl_interp_akima_periodic))}

func (it *GslInterpType) Ptr() uintptr {
	return it.cPtr
}

func (it *GslInterpType) Dispose() {
	it.cPtr = 0
}
