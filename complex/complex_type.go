package complex

/*
	#include <gsl/gsl_complex.h>
	
	gsl_complex make_gsl_complex(double r, double i) {
		gsl_complex rv;
		rv.dat[0] = r;
		rv.dat[1] = i;
		return rv;
	}
	
	gsl_complex_float make_gsl_complex_float(float r, float i) {
		gsl_complex_float rv;
		rv.dat[0] = r;
		rv.dat[1] = i;
		return rv;
	}
	
*/
import "C"

import "unsafe"

func GslComplexToGo(pGslComplex uintptr) complex128 {
	rl := pGslComplex
	im := rl + 8
	return complex(*(*float64)(unsafe.Pointer(rl)),
	 			   *(*float64)(unsafe.Pointer(im)))
}

func GoComplexToGsl(goComplex complex128) uintptr {
	gslCmp := C.make_gsl_complex(C.double(real(goComplex)), C.double(imag(goComplex)))
	return uintptr(unsafe.Pointer(&gslCmp))
}

func GslComplexFloatToGo(pGslComplex uintptr) complex64 {
	rl := pGslComplex
	im := rl + 4
	return complex(*(*float32)(unsafe.Pointer(rl)),
	 			   *(*float32)(unsafe.Pointer(im)))
}

func GoComplexFloatToGsl(goComplex complex64) uintptr {
	gslCmp := C.make_gsl_complex_float(C.float(real(goComplex)), C.float(imag(goComplex)))
	return uintptr(unsafe.Pointer(&gslCmp))
}