//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package mathieu

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_sf_result.h>
   #include <gsl/gsl_sf_mathieu.h>
*/
import "C"

import "unsafe"
import "github.com/dtromb/gogsl"
import "github.com/dtromb/gogsl/sf"
import "reflect"

type GslSfMathieuWorkspace struct {
   gogsl.GslReference
}

func MathieuWorkspaceAlloc(n int, qmax float64) *GslSfMathieuWorkspace {
   _ref := C.gsl_sf_mathieu_alloc(C.size_t(n), C.double(qmax))
   _result := &GslSfMathieuWorkspace{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func (x *GslSfMathieuWorkspace) Dispose() {
   C.gsl_sf_mathieu_free((*C.gsl_sf_mathieu_workspace)(unsafe.Pointer(x.Ptr())))
}

func A(n int32, q float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_mathieu_a(C.int(n), C.double(q), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func B(n int32, q float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_mathieu_b(C.int(n), C.double(q), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func AArray(orderMin int32, orderMax int32, q float64, work *GslSfMathieuWorkspace, resultArray []float64) int32 {
   _slice_header_4 := (*reflect.SliceHeader)(unsafe.Pointer(&resultArray))
   return int32(C.gsl_sf_mathieu_a_array(C.int(orderMin), C.int(orderMax), C.double(q), (*C.gsl_sf_mathieu_workspace)(unsafe.Pointer(work.Ptr())), (*C.double)(unsafe.Pointer(_slice_header_4.Data))))
}

func BArray(orderMin int32, orderMax int32, q float64, work *GslSfMathieuWorkspace, resultArray []float64) int32 {
   _slice_header_4 := (*reflect.SliceHeader)(unsafe.Pointer(&resultArray))
   return int32(C.gsl_sf_mathieu_b_array(C.int(orderMin), C.int(orderMax), C.double(q), (*C.gsl_sf_mathieu_workspace)(unsafe.Pointer(work.Ptr())), (*C.double)(unsafe.Pointer(_slice_header_4.Data))))
}

func Ce(n int32, q float64, x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_mathieu_ce(C.int(n), C.double(q), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Se(n int32, q float64, x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_mathieu_se(C.int(n), C.double(q), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func CeArray(nmin int32, nmax int32, q float64, x float64, work *GslSfMathieuWorkspace, resultArray []float64) int32 {
   _slice_header_5 := (*reflect.SliceHeader)(unsafe.Pointer(&resultArray))
   return int32(C.gsl_sf_mathieu_ce_array(C.int(nmin), C.int(nmax), C.double(q), C.double(x), (*C.gsl_sf_mathieu_workspace)(unsafe.Pointer(work.Ptr())), (*C.double)(unsafe.Pointer(_slice_header_5.Data))))
}

func SeArray(nmin int32, nmax int32, q float64, x float64, work *GslSfMathieuWorkspace, resultArray []float64) int32 {
   _slice_header_5 := (*reflect.SliceHeader)(unsafe.Pointer(&resultArray))
   return int32(C.gsl_sf_mathieu_se_array(C.int(nmin), C.int(nmax), C.double(q), C.double(x), (*C.gsl_sf_mathieu_workspace)(unsafe.Pointer(work.Ptr())), (*C.double)(unsafe.Pointer(_slice_header_5.Data))))
}

func Mc(j int32, n int32, q float64, x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_mathieu_Mc(C.int(j), C.int(n), C.double(q), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Ms(j int32, n int32, q float64, x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_mathieu_Ms(C.int(j), C.int(n), C.double(q), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func McArray(j int32, nmin int32, nmax int32, q float64, x float64, work *GslSfMathieuWorkspace, resultArray []float64) int32 {
   _slice_header_6 := (*reflect.SliceHeader)(unsafe.Pointer(&resultArray))
   return int32(C.gsl_sf_mathieu_Mc_array(C.int(j), C.int(nmin), C.int(nmax), C.double(q), C.double(x), (*C.gsl_sf_mathieu_workspace)(unsafe.Pointer(work.Ptr())), (*C.double)(unsafe.Pointer(_slice_header_6.Data))))
}

func MsArray(j int32, nmin int32, nmax int32, q float64, x float64, work *GslSfMathieuWorkspace, resultArray []float64) int32 {
   _slice_header_6 := (*reflect.SliceHeader)(unsafe.Pointer(&resultArray))
   return int32(C.gsl_sf_mathieu_Ms_array(C.int(j), C.int(nmin), C.int(nmax), C.double(q), C.double(x), (*C.gsl_sf_mathieu_workspace)(unsafe.Pointer(work.Ptr())), (*C.double)(unsafe.Pointer(_slice_header_6.Data))))
}

