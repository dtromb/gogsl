//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package dawson

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_sf_result.h>
   #include <gsl/gsl_sf_dawson.h>
*/
import "C"

import "github.com/dtromb/gogsl/sf"
import "unsafe"

func GslSfDawson(x float64) float64 {
   return float64(C.gsl_sf_dawson(C.double(x)))
}

func GslSfDawsonE(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_dawson_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

