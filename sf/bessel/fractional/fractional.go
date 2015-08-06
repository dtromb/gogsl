//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package fractional

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_sf_result.h>
   #include <gsl/gsl_sf_bessel.h>
*/
import "C"

import "github.com/dtromb/gogsl/sf"
import "unsafe"
import "reflect"

func Jnu(nu float64, x float64) float64 {
   return float64(C.gsl_sf_bessel_Jnu(C.double(nu), C.double(x)))
}

func JnuE(nu float64, x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_bessel_Jnu_e(C.double(nu), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func SequenceJnuE(nu float64, mode sf.GslMode, size int, v []float64) int32 {
   _slice_header_3 := (*reflect.SliceHeader)(unsafe.Pointer(&v))
   return int32(C.gsl_sf_bessel_sequence_Jnu_e(C.double(nu), C.gsl_mode_t(mode), C.size_t(size), (*C.double)(unsafe.Pointer(_slice_header_3.Data))))
}

func Ynu(nu float64, x float64) float64 {
   return float64(C.gsl_sf_bessel_Ynu(C.double(nu), C.double(x)))
}

func YnuE(nu float64, x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_bessel_Ynu_e(C.double(nu), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Inu(nu float64, x float64) float64 {
   return float64(C.gsl_sf_bessel_Inu(C.double(nu), C.double(x)))
}

func InuE(nu float64, x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_bessel_Inu_e(C.double(nu), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func InuScaled(nu float64, x float64) float64 {
   return float64(C.gsl_sf_bessel_Inu_scaled(C.double(nu), C.double(x)))
}

func InuScaledE(nu float64, x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_bessel_Inu_scaled_e(C.double(nu), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Knu(nu float64, x float64) float64 {
   return float64(C.gsl_sf_bessel_Knu(C.double(nu), C.double(x)))
}

func KnuE(nu float64, x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_bessel_Knu_e(C.double(nu), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func LnKnu(nu float64, x float64) float64 {
   return float64(C.gsl_sf_bessel_lnKnu(C.double(nu), C.double(x)))
}

func LnKnuE(nu float64, x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_bessel_lnKnu_e(C.double(nu), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func KnuScaled(nu float64, x float64) float64 {
   return float64(C.gsl_sf_bessel_Knu_scaled(C.double(nu), C.double(x)))
}

func KnuScaledE(nu float64, x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_bessel_Knu_scaled_e(C.double(nu), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func ZeroJ0(s uint32) float64 {
   return float64(C.gsl_sf_bessel_zero_J0(C.uint(s)))
}

func ZeroJ0E(s uint32, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_bessel_zero_J0_e(C.uint(s), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func ZeroJ1(s uint32) float64 {
   return float64(C.gsl_sf_bessel_zero_J1(C.uint(s)))
}

func ZeroJ1E(s uint32, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_bessel_zero_J1_e(C.uint(s), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func ZeroJnu(nu float64, s uint32) float64 {
   return float64(C.gsl_sf_bessel_zero_Jnu(C.double(nu), C.uint(s)))
}

func ZeroJnuE(nu float64, s uint32, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_bessel_zero_Jnu_e(C.double(nu), C.uint(s), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

