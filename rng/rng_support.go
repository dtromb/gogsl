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

var RNG_TYPE_MT19937 *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_mt19937))}
var RNG_TYPE_RANLXS0 *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_ranlxs0))}
var RNG_TYPE_RANLXS1 *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_ranlxs1))}
var RNG_TYPE_RANLXS2 *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_ranlxs2))}
var RNG_TYPE_RANLXD1 *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_ranlxd1))}
var RNG_TYPE_RANLXD2 *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_ranlxd2))}
var RNG_TYPE_RANLUX *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_ranlux))}
var RNG_TYPE_RANLUX389 *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_ranlux389))}
var RNG_TYPE_CMRG *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_cmrg))}
var RNG_TYPE_MRG *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_mrg))}
var RNG_TYPE_TAUS *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_taus))}
var RNG_TYPE_TAUS2 *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_taus2))}
var RNG_TYPE_GFSR4 *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_gfsr4))}
var RNG_TYPE_RAND *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_rand))}
var RNG_TYPE_RANDOM_BSD *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_random_bsd))}
var RNG_TYPE_RANDOM_LIBC5 *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_random_libc5))}
var RNG_TYPE_RANDOM_GLIBC2 *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_random_glibc2))}
var RNG_TYPE_RANDOM8_BSD *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_random8_bsd))}
var RNG_TYPE_RANDOM32_BSD *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_random32_bsd))}
var RNG_TYPE_RANDOM64_BSD *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_random64_bsd))}
var RNG_TYPE_RANDOM128_BSD *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_random128_bsd))}
var RNG_TYPE_RANDOM256_BSD *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_random256_bsd))}
var RNG_TYPE_RAND48 *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_rand48))}
var RNG_TYPE_RANF *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_ranf))}
var RNG_TYPE_RANMAR *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_ranmar))}
var RNG_TYPE_R250 *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_r250))}
var RNG_TYPE_TT800 *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_tt800))}
var RNG_TYPE_VAX *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_vax))}
var RNG_TYPE_TRANSPUTER *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_transputer))}
var RNG_TYPE_RANDU *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_randu))}
var RNG_TYPE_MINSTD *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_minstd))}
var RNG_TYPE_UNI *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_uni))}
var RNG_TYPE_UNI32 *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_uni32))}
var RNG_TYPE_SLATEC *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_slatec))}
var RNG_TYPE_ZUF *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_zuf))}
var RNG_TYPE_KNUTHRAN2 *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_knuthran2))}
var RNG_TYPE_KNUTHRAN2002 *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_knuthran2002))}
var RNG_TYPE_KNUTHRAN *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_knuthran))}
var RNG_TYPE_BOROSH13 *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_borosh13))}
var RNG_TYPE_FISHMAN18 *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_fishman18))}
var RNG_TYPE_FISHMAN20 *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_fishman20))}
var RNG_TYPE_LECUYER21 *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_lecuyer21))}
var RNG_TYPE_WATERMAN14 *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_waterman14))}
var RNG_TYPE_FISHMAN2X *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_fishman2x))}
var RNG_TYPE_COVEYOU *GslRngType = &GslRngType{cPtr: uintptr(unsafe.Pointer(C.gsl_rng_coveyou))}

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
