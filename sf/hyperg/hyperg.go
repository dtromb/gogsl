//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package hyperg

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_sf_result.h>
   #include <gsl/gsl_sf_hyperg.h>
*/
import "C"

import "github.com/dtromb/gogsl/sf"
import "unsafe"

func Hyperg0F1(c float64, x float64) float64 {
	return float64(C.gsl_sf_hyperg_0F1(C.double(c), C.double(x)))
}

func Hyperg0F1E(c float64, x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_hyperg_0F1_e(C.double(c), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Hyperg1F1Int(m int32, n int32, x float64) float64 {
	return float64(C.gsl_sf_hyperg_1F1_int(C.int(m), C.int(n), C.double(x)))
}

func Hyperg1F1IntE(m int32, n int32, x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_hyperg_1F1_int_e(C.int(m), C.int(n), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Hyperg1F1(a float64, b float64, x float64) float64 {
	return float64(C.gsl_sf_hyperg_1F1(C.double(a), C.double(b), C.double(x)))
}

func Hyperg1F1E(a float64, b float64, x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_hyperg_1F1_e(C.double(a), C.double(b), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func HypergUInt(m int32, n int32, x float64) float64 {
	return float64(C.gsl_sf_hyperg_U_int(C.int(m), C.int(n), C.double(x)))
}

func HypergUIntE(m int32, n int32, x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_hyperg_U_int_e(C.int(m), C.int(n), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func HypergUIntE10E(m int32, n int32, x float64, result *sf.GslSfResultE10) int32 {
	return int32(C.gsl_sf_hyperg_U_int_e10_e(C.int(m), C.int(n), C.double(x), (*C.gsl_sf_result_e10)(unsafe.Pointer(result.Ptr()))))
}

func HypergU(a float64, b float64, x float64) float64 {
	return float64(C.gsl_sf_hyperg_U(C.double(a), C.double(b), C.double(x)))
}

func HypergUE(a float64, b float64, x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_hyperg_U_e(C.double(a), C.double(b), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func HypergUE10E(a float64, b float64, x float64, result *sf.GslSfResultE10) int32 {
	return int32(C.gsl_sf_hyperg_U_e10_e(C.double(a), C.double(b), C.double(x), (*C.gsl_sf_result_e10)(unsafe.Pointer(result.Ptr()))))
}

func Hyperg2F1(a float64, b float64, c float64, x float64) float64 {
	return float64(C.gsl_sf_hyperg_2F1(C.double(a), C.double(b), C.double(c), C.double(x)))
}

func Hyperg2F1E(a float64, b float64, c float64, x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_hyperg_2F1_e(C.double(a), C.double(b), C.double(c), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Hyperg2F1Conj(aR float64, aI float64, c float64, x float64) float64 {
	return float64(C.gsl_sf_hyperg_2F1_conj(C.double(aR), C.double(aI), C.double(c), C.double(x)))
}

func Hyperg2F1ConjE(aR float64, aI float64, c float64, x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_hyperg_2F1_conj_e(C.double(aR), C.double(aI), C.double(c), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Hyperg2F1Renorm(a float64, b float64, c float64, x float64) float64 {
	return float64(C.gsl_sf_hyperg_2F1_renorm(C.double(a), C.double(b), C.double(c), C.double(x)))
}

func Hyperg2F1RenormE(a float64, b float64, c float64, x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_hyperg_2F1_renorm_e(C.double(a), C.double(b), C.double(c), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Hyperg2F1ConjRenorm(aR float64, aI float64, c float64, x float64) float64 {
	return float64(C.gsl_sf_hyperg_2F1_conj_renorm(C.double(aR), C.double(aI), C.double(c), C.double(x)))
}

func Hyperg2F1ConjRenormE(aR float64, aI float64, c float64, x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_hyperg_2F1_conj_renorm_e(C.double(aR), C.double(aI), C.double(c), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Hyperg2F0(a float64, b float64, x float64) float64 {
	return float64(C.gsl_sf_hyperg_2F0(C.double(a), C.double(b), C.double(x)))
}

func Hyperg2F0E(a float64, b float64, x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_hyperg_2F0_e(C.double(a), C.double(b), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}
