//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package log

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_sf_result.h>
   #include <gsl/gsl_sf_log.h>
*/
import "C"

import "github.com/dtromb/gogsl/sf"
import "unsafe"

func Log(x float64) float64 {
   return float64(C.gsl_sf_log(C.double(x)))
}

func LogE(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_log_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func LogAbs(x float64) float64 {
   return float64(C.gsl_sf_log_abs(C.double(x)))
}

func LogAbsE(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_log_abs_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func ComplexLogE(zr float64, zi float64, lnr *sf.GslSfResult, theta *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_complex_log_e(C.double(zr), C.double(zi), (*C.gsl_sf_result)(unsafe.Pointer(lnr.Ptr())), (*C.gsl_sf_result)(unsafe.Pointer(theta.Ptr()))))
}

func Log1plusx(x float64) float64 {
   return float64(C.gsl_sf_log_1plusx(C.double(x)))
}

func Log1plusxE(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_log_1plusx_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Log1plusxMx(x float64) float64 {
   return float64(C.gsl_sf_log_1plusx_mx(C.double(x)))
}

func Log1plusxMxE(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_log_1plusx_mx_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

