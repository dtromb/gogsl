//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package coupling

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_sf_result.h>
   #include <gsl/gsl_sf_coupling.h>
*/
import "C"

import "github.com/dtromb/gogsl/sf"
import "unsafe"

func GslSfCoupling3j(twoJa int32, twoJb int32, twoJc int32, twoMa int32, twoMb int32, twoMc int32) float64 {
	return float64(C.gsl_sf_coupling_3j(C.int(twoJa), C.int(twoJb), C.int(twoJc), C.int(twoMa), C.int(twoMb), C.int(twoMc)))
}

func GslSfCoupling3jE(twoJa int32, twoJb int32, twoJc int32, twoMa int32, twoMb int32, twoMc int32, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_coupling_3j_e(C.int(twoJa), C.int(twoJb), C.int(twoJc), C.int(twoMa), C.int(twoMb), C.int(twoMc), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func GslSfCoupling6j(twoJa int32, twoJb int32, twoJc int32, twoJd int32, twoJe int32, twoJf int32) float64 {
	return float64(C.gsl_sf_coupling_6j(C.int(twoJa), C.int(twoJb), C.int(twoJc), C.int(twoJd), C.int(twoJe), C.int(twoJf)))
}

func GslSfCoupling6jE(twoJa int32, twoJb int32, twoJc int32, twoJd int32, twoJe int32, twoJf int32, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_coupling_6j_e(C.int(twoJa), C.int(twoJb), C.int(twoJc), C.int(twoJd), C.int(twoJe), C.int(twoJf), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func GslSfCoupling9j(twoJa int32, twoJb int32, twoJc int32, twoJd int32, twoJe int32, twoJf int32, twoJg int32, twoJh int32, twoJi int32) float64 {
	return float64(C.gsl_sf_coupling_9j(C.int(twoJa), C.int(twoJb), C.int(twoJc), C.int(twoJd), C.int(twoJe), C.int(twoJf), C.int(twoJg), C.int(twoJh), C.int(twoJi)))
}

func GslSfCoupling9jE(twoJa int32, twoJb int32, twoJc int32, twoJd int32, twoJe int32, twoJf int32, twoJg int32, twoJh int32, twoJi int32, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_coupling_9j_e(C.int(twoJa), C.int(twoJb), C.int(twoJc), C.int(twoJd), C.int(twoJe), C.int(twoJf), C.int(twoJg), C.int(twoJh), C.int(twoJi), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}
