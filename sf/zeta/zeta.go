//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package zeta

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_sf_result.h>
   #include <gsl/gsl_sf_zeta.h>
*/
import "C"

import "github.com/dtromb/gogsl/sf"
import "unsafe"

func ZetaInt(n int32) float64 {
   return float64(C.gsl_sf_zeta_int(C.int(n)))
}

func ZetaIntE(n int32, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_zeta_int_e(C.int(n), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Zeta(s float64) float64 {
   return float64(C.gsl_sf_zeta(C.double(s)))
}

func ZetaE(s float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_zeta_e(C.double(s), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Zetam1Int(n int32) float64 {
   return float64(C.gsl_sf_zetam1_int(C.int(n)))
}

func Zetam1IntE(n int32, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_zetam1_int_e(C.int(n), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Zetam1(s float64) float64 {
   return float64(C.gsl_sf_zetam1(C.double(s)))
}

func Zetam1E(s float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_zetam1_e(C.double(s), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Hzeta(s float64, q float64) float64 {
   return float64(C.gsl_sf_hzeta(C.double(s), C.double(q)))
}

func HzetaE(s float64, q float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_hzeta_e(C.double(s), C.double(q), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func EtaInt(n int32) float64 {
   return float64(C.gsl_sf_eta_int(C.int(n)))
}

func EtaIntE(n int32, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_eta_int_e(C.int(n), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Eta(s float64) float64 {
   return float64(C.gsl_sf_eta(C.double(s)))
}

func EtaE(s float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_eta_e(C.double(s), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

