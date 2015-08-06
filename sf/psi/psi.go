//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package psi

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_sf_result.h>
   #include <gsl/gsl_sf_psi.h>
*/
import "C"

import "github.com/dtromb/gogsl/sf"
import "unsafe"

func PsiInt(n int32) float64 {
   return float64(C.gsl_sf_psi_int(C.int(n)))
}

func PsiIntE(n int32, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_psi_int_e(C.int(n), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Psi(x float64) float64 {
   return float64(C.gsl_sf_psi(C.double(x)))
}

func PsiE(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_psi_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Psi1piy(y float64) float64 {
   return float64(C.gsl_sf_psi_1piy(C.double(y)))
}

func Psi1piyE(y float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_psi_1piy_e(C.double(y), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Psi1Int(n int32) float64 {
   return float64(C.gsl_sf_psi_1_int(C.int(n)))
}

func Psi1IntE(n int32, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_psi_1_int_e(C.int(n), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Psi1(x float64) float64 {
   return float64(C.gsl_sf_psi_1(C.double(x)))
}

func Psi1E(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_psi_1_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func PsiN(n int32, x float64) float64 {
   return float64(C.gsl_sf_psi_n(C.int(n), C.double(x)))
}

func PsiNE(n int32, x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_psi_n_e(C.int(n), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

