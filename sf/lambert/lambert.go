//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package lambert

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_sf_result.h>
   #include <gsl/gsl_sf_lambert.h>
*/
import "C"

import "github.com/dtromb/gogsl/sf"
import "unsafe"

func W0(x float64) float64 {
	return float64(C.gsl_sf_lambert_W0(C.double(x)))
}

func W0E(x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_lambert_W0_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Wm1(x float64) float64 {
	return float64(C.gsl_sf_lambert_Wm1(C.double(x)))
}

func Wm1E(x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_lambert_Wm1_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}
