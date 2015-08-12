//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package exp

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_sf_result.h>
   #include <gsl/gsl_sf_exp.h>
*/
import "C"

import "github.com/dtromb/gogsl/sf"
import "unsafe"

func Exp(x float64) float64 {
	return float64(C.gsl_sf_exp(C.double(x)))
}

func ExpE(x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_exp_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func ExpE10E(x float64, result *sf.GslSfResultE10) int32 {
	return int32(C.gsl_sf_exp_e10_e(C.double(x), (*C.gsl_sf_result_e10)(unsafe.Pointer(result.Ptr()))))
}

func ExpMult(x float64, y float64) float64 {
	return float64(C.gsl_sf_exp_mult(C.double(x), C.double(y)))
}

func ExpMultE(x float64, y float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_exp_mult_e(C.double(x), C.double(y), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func ExpMultE10E(x float64, y float64, result *sf.GslSfResultE10) int32 {
	return int32(C.gsl_sf_exp_mult_e10_e(C.double(x), C.double(y), (*C.gsl_sf_result_e10)(unsafe.Pointer(result.Ptr()))))
}

func Expm1(x float64) float64 {
	return float64(C.gsl_sf_expm1(C.double(x)))
}

func Expm1E(x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_expm1_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Exprel(x float64) float64 {
	return float64(C.gsl_sf_exprel(C.double(x)))
}

func ExprelE(x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_exprel_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Exprel2(x float64) float64 {
	return float64(C.gsl_sf_exprel_2(C.double(x)))
}

func Exprel2E(x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_exprel_2_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func ExprelN(n int32, x float64) float64 {
	return float64(C.gsl_sf_exprel_n(C.int(n), C.double(x)))
}

func ExprelNE(n int32, x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_exprel_n_e(C.int(n), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func ExpErrE(x float64, dx float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_exp_err_e(C.double(x), C.double(dx), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func ExpErrE10E(x float64, dx float64, result *sf.GslSfResultE10) int32 {
	return int32(C.gsl_sf_exp_err_e10_e(C.double(x), C.double(dx), (*C.gsl_sf_result_e10)(unsafe.Pointer(result.Ptr()))))
}

func ExpMultErrE(x float64, dx float64, y float64, dy float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_exp_mult_err_e(C.double(x), C.double(dx), C.double(y), C.double(dy), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func ExpMultErrE10E(x float64, dx float64, y float64, dy float64, result *sf.GslSfResultE10) int32 {
	return int32(C.gsl_sf_exp_mult_err_e10_e(C.double(x), C.double(dx), C.double(y), C.double(dy), (*C.gsl_sf_result_e10)(unsafe.Pointer(result.Ptr()))))
}
