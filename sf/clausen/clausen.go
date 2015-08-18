//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package clausen

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_sf_result.h>
   #include <gsl/gsl_sf_clausen.h>
*/
import "C"

import "github.com/dtromb/gogsl/sf"
import "unsafe"

func GslSfClausen(x float64) float64 {
   return float64(C.gsl_sf_clausen(C.double(x)))
}

func GslSfClausenE(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_clausen_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

