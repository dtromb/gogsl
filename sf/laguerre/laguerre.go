//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package laguerre

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_sf_result.h>
   #include <gsl/gsl_sf_laguerre.h>
*/
import "C"

import "github.com/dtromb/gogsl/sf"
import "unsafe"

func Laguerre1(a float64, x float64) float64 {
	return float64(C.gsl_sf_laguerre_1(C.double(a), C.double(x)))
}

func Laguerre2(a float64, x float64) float64 {
	return float64(C.gsl_sf_laguerre_2(C.double(a), C.double(x)))
}

func Laguerre3(a float64, x float64) float64 {
	return float64(C.gsl_sf_laguerre_3(C.double(a), C.double(x)))
}

func Laguerre1E(a float64, x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_laguerre_1_e(C.double(a), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Laguerre2E(a float64, x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_laguerre_2_e(C.double(a), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Laguerre3E(a float64, x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_laguerre_3_e(C.double(a), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func LaguerreN(n int32, a float64, x float64) float64 {
	return float64(C.gsl_sf_laguerre_n(C.int(n), C.double(a), C.double(x)))
}

func LaguerreNE(n int32, a float64, x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_laguerre_n_e(C.int(n), C.double(a), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}
