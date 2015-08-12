//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package expint

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_sf_result.h>
   #include <gsl/gsl_sf_expint.h>
*/
import "C"

import "github.com/dtromb/gogsl/sf"
import "unsafe"

func E1(x float64) float64 {
	return float64(C.gsl_sf_expint_E1(C.double(x)))
}

func E1E(x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_expint_E1_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func E2(x float64) float64 {
	return float64(C.gsl_sf_expint_E2(C.double(x)))
}

func E2E(x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_expint_E2_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func En(n int32, x float64) float64 {
	return float64(C.gsl_sf_expint_En(C.int(n), C.double(x)))
}

func EnE(n int32, x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_expint_En_e(C.int(n), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Ei(x float64) float64 {
	return float64(C.gsl_sf_expint_Ei(C.double(x)))
}

func EiE(x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_expint_Ei_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Shi(x float64) float64 {
	return float64(C.gsl_sf_Shi(C.double(x)))
}

func ShiE(x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_Shi_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Chi(x float64) float64 {
	return float64(C.gsl_sf_Chi(C.double(x)))
}

func ChiE(x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_Chi_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Ei3(x float64) float64 {
	return float64(C.gsl_sf_expint_3(C.double(x)))
}

func Ei3e(x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_expint_3_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Si(x float64) float64 {
	return float64(C.gsl_sf_Si(C.double(x)))
}

func SiE(x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_Si_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Ci(x float64) float64 {
	return float64(C.gsl_sf_Ci(C.double(x)))
}

func CiE(x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_Ci_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Atanint(x float64) float64 {
	return float64(C.gsl_sf_atanint(C.double(x)))
}

func AtanintE(x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_atanint_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}
