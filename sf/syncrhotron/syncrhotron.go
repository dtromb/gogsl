//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package syncrhotron

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_sf_result.h>
   #include <gsl/gsl_sf_synchrotron.h>
*/
import "C"

import "github.com/dtromb/gogsl/sf"
import "unsafe"

func Synchrotron1(x float64) float64 {
   return float64(C.gsl_sf_synchrotron_1(C.double(x)))
}

func Synchrotron1E(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_synchrotron_1_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Synchrotron2(x float64) float64 {
   return float64(C.gsl_sf_synchrotron_2(C.double(x)))
}

func Synchrotron2E(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_synchrotron_2_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

