package rng

/*
	#include <gsl/gsl_rng.h>
*/
import "C"

import (
	"unsafe"
)

type RngState struct {
}

func (rs *RngState) Dispose() {
}


var RNG_TYPE_MT19937 	*GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_mt19937))}
var RNG_TYPE_RANLXS0 	*GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_ranlxs0))}
var RNG_TYPE_RANLXS1 	*GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_ranlxs1))}
var RNG_TYPE_RANLXS2 	*GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_ranlxs2))}
var RNG_TYPE_RANLXD1 	*GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_ranlxd1))}
var RNG_TYPE_RANLXD2 	*GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_ranlxd2))}
var RNG_TYPE_RANLUX 	*GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_ranlux))}
var RNG_TYPE_RANLUX389	*GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_ranlux389))}
var RNG_TYPE_CMRG		*GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_cmrg))}
var RNG_TYPE_MRG		*GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_mrg))}
var RNG_TYPE_TAUS		*GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_taus))}
var RNG_TYPE_TAUS2		*GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_taus2))}
var RNG_TYPE_GFSR4		*GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_gfsr4))}

func DefaultRngType() *GslRngType {
	return &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_default))}
}

func (rs *GslRngType) Dispose() {
	rs.cPtr = 0
}

func (rs *GslRngType) Ptr() uintptr {
	return rs.cPtr
}

//void gsl_ran_sample (const gsl_rng * r, void * dest, size_t k, void * src, size_t n, size_t size)

