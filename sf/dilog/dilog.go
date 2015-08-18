//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package dilog

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_sf_result.h>
   #include <gsl/gsl_sf_dilog.h>
*/
import "C"

import "github.com/dtromb/gogsl/sf"
import "unsafe"


func Dilog(x float64) float64 {
   return float64(C.gsl_sf_dilog(C.double(x)))
}

func DilogE(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_dilog_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func ComplexDilogE(r float64, theta float64, resultRe *sf.GslSfResult, resultIm *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_complex_dilog_e(C.double(r), C.double(theta), (*C.gsl_sf_result)(unsafe.Pointer(resultRe.Ptr())), (*C.gsl_sf_result)(unsafe.Pointer(resultIm.Ptr()))))
}

