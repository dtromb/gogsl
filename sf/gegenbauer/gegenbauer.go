//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package gegenbauer

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_sf_result.h>
   #include <gsl/gsl_sf_gegenbauer.h>
*/
import "C"

import "github.com/dtromb/gogsl/sf"
import "unsafe"
import "reflect"

func Gegenpoly1(lambda float64, x float64) float64 {
   return float64(C.gsl_sf_gegenpoly_1(C.double(lambda), C.double(x)))
}

func Gegenpoly2(lambda float64, x float64) float64 {
   return float64(C.gsl_sf_gegenpoly_2(C.double(lambda), C.double(x)))
}

func Gegenpoly3(lambda float64, x float64) float64 {
   return float64(C.gsl_sf_gegenpoly_3(C.double(lambda), C.double(x)))
}

func Gegenpoly1E(lambda float64, x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_gegenpoly_1_e(C.double(lambda), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Gegenpoly2E(lambda float64, x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_gegenpoly_2_e(C.double(lambda), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Gegenpoly3E(lambda float64, x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_gegenpoly_3_e(C.double(lambda), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func GegenpolyN(n int32, lambda float64, x float64) float64 {
   return float64(C.gsl_sf_gegenpoly_n(C.int(n), C.double(lambda), C.double(x)))
}

func GegenpolyNE(n int32, lambda float64, x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_gegenpoly_n_e(C.int(n), C.double(lambda), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func GegenpolyArray(nmax int32, lambda float64, x float64, resultArray []float64) int32 {
   _slice_header_3 := (*reflect.SliceHeader)(unsafe.Pointer(&resultArray))
   return int32(C.gsl_sf_gegenpoly_array(C.int(nmax), C.double(lambda), C.double(x), (*C.double)(unsafe.Pointer(_slice_header_3.Data))))
}

