//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package transport

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_sf_result.h>
   #include <gsl/gsl_sf_transport.h>
*/
import "C"

import "github.com/dtromb/gogsl/sf"
import "unsafe"

func Transport2(x float64) float64 {
   return float64(C.gsl_sf_transport_2(C.double(x)))
}

func Transport2E(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_transport_2_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Transport3(x float64) float64 {
   return float64(C.gsl_sf_transport_3(C.double(x)))
}

func Transport3E(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_transport_3_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Transport4(x float64) float64 {
   return float64(C.gsl_sf_transport_4(C.double(x)))
}

func Transport4E(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_transport_4_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Transport5(x float64) float64 {
   return float64(C.gsl_sf_transport_5(C.double(x)))
}

func Transport5E(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_transport_5_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

