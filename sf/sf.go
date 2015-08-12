//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package sf

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_sf_result.h>
   #include <gsl/gsl_sf_elementary.h>
*/
import "C"

import "unsafe"
import "github.com/dtromb/gogsl"

type GslSfResult struct {
	gogsl.GslReference
	CData []byte
}

type GslSfResultE10 struct {
	gogsl.GslReference
	CData []byte
}

type GslMode uint32

func GslSfMultiplyE(x float64, y float64, result *GslSfResult) int32 {
	return int32(C.gsl_sf_multiply_e(C.double(x), C.double(y), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func GslSfMultiplyErrE(x float64, dx float64, y float64, dy float64, result *GslSfResult) int32 {
	return int32(C.gsl_sf_multiply_err_e(C.double(x), C.double(dx), C.double(y), C.double(dy), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}
