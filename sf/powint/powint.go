//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package powint

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_sf_result.h>
   #include <gsl/gsl_sf_pow_int.h>
*/
import "C"

import "github.com/dtromb/gogsl/sf"
import "unsafe"

func PowInt(x float64, n int32) float64 {
	return float64(C.gsl_sf_pow_int(C.double(x), C.int(n)))
}

func PowIntE(x float64, n int32, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_pow_int_e(C.double(x), C.int(n), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}
