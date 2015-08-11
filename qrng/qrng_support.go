package qrng

/*
	#include <gsl/gsl_qrng.h>
*/
import "C"

import "unsafe"

type QrngState struct {
	
}

var QNRG_NIEDERREITER_2 		*GslQrngType = &GslQrngType{cPtr: uintptr(unsafe.Pointer(C.gsl_qrng_niederreiter_2))}
var QRNG_SOBOL		 			*GslQrngType = &GslQrngType{cPtr: uintptr(unsafe.Pointer(C.gsl_qrng_sobol))}
var QRNG_HALTON		 			*GslQrngType = &GslQrngType{cPtr: uintptr(unsafe.Pointer(C.gsl_qrng_halton))}
var QRNG_REVERSEHALTON 			*GslQrngType = &GslQrngType{cPtr: uintptr(unsafe.Pointer(C.gsl_qrng_reversehalton))}


func (qt *GslQrngType) Ptr() uintptr {
	return qt.cPtr
}