//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package erf

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_sf_result.h>
   #include <gsl/gsl_sf_erf.h>
*/
import "C"

import "github.com/dtromb/gogsl/sf"
import "unsafe"

func Erf(x float64) float64 {
   return float64(C.gsl_sf_erf(C.double(x)))
}

func ErfE(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_erf_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Erfc(x float64) float64 {
   return float64(C.gsl_sf_erfc(C.double(x)))
}

func ErfcE(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_erfc_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func LogErfc(x float64) float64 {
   return float64(C.gsl_sf_log_erfc(C.double(x)))
}

func LogErfcE(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_log_erfc_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func ErfZ(x float64) float64 {
   return float64(C.gsl_sf_erf_Z(C.double(x)))
}

func ErfZE(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_erf_Z_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func ErfQ(x float64) float64 {
   return float64(C.gsl_sf_erf_Q(C.double(x)))
}

func ErfQE(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_erf_Q_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Hazard(x float64) float64 {
   return float64(C.gsl_sf_hazard(C.double(x)))
}

func HazardE(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_hazard_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

