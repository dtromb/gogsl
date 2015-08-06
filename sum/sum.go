//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package sum

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_sum.h>
*/
import "C"

import "unsafe"
import "github.com/dtromb/gogsl"
import "reflect"

type GslSumLevinUWorkspace struct {
   gogsl.GslReference
}

type GslSumLevinUtruncWorkspace struct {
   gogsl.GslReference
}

func LevinUAlloc(n int) *GslSumLevinUWorkspace {
   _ref := C.gsl_sum_levin_u_alloc(C.size_t(n))
   _result := &GslSumLevinUWorkspace{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func (x *GslSumLevinUWorkspace) Dispose() {
   C.gsl_sum_levin_u_free((*C.gsl_sum_levin_u_workspace)(unsafe.Pointer(x.Ptr())))
}

func LevinUAccel(array []float64, arraySize int, w *GslSumLevinUWorkspace) (int32, float64, float64) {
   var _outptr_3 C.double
   var _outptr_4 C.double
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&array))
   _result := int32(C.gsl_sum_levin_u_accel((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(arraySize), (*C.gsl_sum_levin_u_workspace)(unsafe.Pointer(w.Ptr())), &_outptr_3, &_outptr_4))
   return _result, *(*float64)(unsafe.Pointer(&_outptr_3)), *(*float64)(unsafe.Pointer(&_outptr_4))
}

func LevinUtruncAlloc(n int) *GslSumLevinUtruncWorkspace {
   _ref := C.gsl_sum_levin_utrunc_alloc(C.size_t(n))
   _result := &GslSumLevinUtruncWorkspace{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func (x *GslSumLevinUtruncWorkspace) Dispose() {
   C.gsl_sum_levin_utrunc_free((*C.gsl_sum_levin_utrunc_workspace)(unsafe.Pointer(x.Ptr())))
}

func LevinUtruncAccel(array []float64, arraySize int, w *GslSumLevinUtruncWorkspace) (int32, float64, float64) {
   var _outptr_3 C.double
   var _outptr_4 C.double
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&array))
   _result := int32(C.gsl_sum_levin_utrunc_accel((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(arraySize), (*C.gsl_sum_levin_utrunc_workspace)(unsafe.Pointer(w.Ptr())), &_outptr_3, &_outptr_4))
   return _result, *(*float64)(unsafe.Pointer(&_outptr_3)), *(*float64)(unsafe.Pointer(&_outptr_4))
}

