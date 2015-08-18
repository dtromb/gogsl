//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package debye

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_sf_result.h>
   #include <gsl/gsl_sf_debye.h>
*/
import "C"

import "github.com/dtromb/gogsl/sf"
import "unsafe"

func Debye1(x float64) float64 {
   return float64(C.gsl_sf_debye_1(C.double(x)))
}

func Debye1E(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_debye_1_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Debye2(x float64) float64 {
   return float64(C.gsl_sf_debye_2(C.double(x)))
}

func Debye2E(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_debye_2_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Debye3(x float64) float64 {
   return float64(C.gsl_sf_debye_3(C.double(x)))
}

func Debye3E(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_debye_3_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Debye4(x float64) float64 {
   return float64(C.gsl_sf_debye_4(C.double(x)))
}

func Debye4E(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_debye_4_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Debye5(x float64) float64 {
   return float64(C.gsl_sf_debye_5(C.double(x)))
}

func Debye5E(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_debye_5_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Debye6(x float64) float64 {
   return float64(C.gsl_sf_debye_6(C.double(x)))
}

func Debye6E(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_debye_6_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

