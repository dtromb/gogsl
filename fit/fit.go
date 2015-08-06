//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package fit

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_fit.h>
   #include <gsl/gsl_multifit.h>
   #include <gsl/gsl_vector.h>
   #include <gsl/gsl_matrix.h>
*/
import "C"

import "unsafe"
import "github.com/dtromb/gogsl"
import "reflect"
import "github.com/dtromb/gogsl/matrix"
import "github.com/dtromb/gogsl/vector"

type GslMultifitLinearWorkspace struct {
   gogsl.GslReference
}

type GslMultifitRobustWorkspace struct {
   gogsl.GslReference
}

type GslMultifitRobustType struct {
   gogsl.GslReference
}

func Linear(x []float64, xstride int, y []float64, ystride int, n int) (int32, float64, float64, float64, float64, float64, float64) {
   var _outptr_5 C.double
   var _outptr_6 C.double
   var _outptr_7 C.double
   var _outptr_8 C.double
   var _outptr_9 C.double
   var _outptr_10 C.double
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&x))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&y))
   _result := int32(C.gsl_fit_linear((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(xstride), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(ystride), C.size_t(n), &_outptr_5, &_outptr_6, &_outptr_7, &_outptr_8, &_outptr_9, &_outptr_10))
   return _result, *(*float64)(unsafe.Pointer(&_outptr_5)), *(*float64)(unsafe.Pointer(&_outptr_6)), *(*float64)(unsafe.Pointer(&_outptr_7)), *(*float64)(unsafe.Pointer(&_outptr_8)), *(*float64)(unsafe.Pointer(&_outptr_9)), *(*float64)(unsafe.Pointer(&_outptr_10))
}

func Wlinear(x []float64, xstride int, w []float64, wstride int, y []float64, ystride int, n int) (int32, float64, float64, float64, float64, float64, float64) {
   var _outptr_7 C.double
   var _outptr_8 C.double
   var _outptr_9 C.double
   var _outptr_10 C.double
   var _outptr_11 C.double
   var _outptr_12 C.double
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&x))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&w))
   _slice_header_4 := (*reflect.SliceHeader)(unsafe.Pointer(&y))
   _result := int32(C.gsl_fit_wlinear((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(xstride), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(wstride), (*C.double)(unsafe.Pointer(_slice_header_4.Data)), C.size_t(ystride), C.size_t(n), &_outptr_7, &_outptr_8, &_outptr_9, &_outptr_10, &_outptr_11, &_outptr_12))
   return _result, *(*float64)(unsafe.Pointer(&_outptr_7)), *(*float64)(unsafe.Pointer(&_outptr_8)), *(*float64)(unsafe.Pointer(&_outptr_9)), *(*float64)(unsafe.Pointer(&_outptr_10)), *(*float64)(unsafe.Pointer(&_outptr_11)), *(*float64)(unsafe.Pointer(&_outptr_12))
}

func LinearEst(x float64, c0 float64, c1 float64, cov00 float64, cov01 float64, cov11 float64) (int32, float64, float64) {
   var _outptr_6 C.double
   var _outptr_7 C.double
   _result := int32(C.gsl_fit_linear_est(C.double(x), C.double(c0), C.double(c1), C.double(cov00), C.double(cov01), C.double(cov11), &_outptr_6, &_outptr_7))
   return _result, *(*float64)(unsafe.Pointer(&_outptr_6)), *(*float64)(unsafe.Pointer(&_outptr_7))
}

func Mul(x []float64, xstride int, y []float64, ystride int, n int) (int32, float64, float64, float64) {
   var _outptr_5 C.double
   var _outptr_6 C.double
   var _outptr_7 C.double
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&x))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&y))
   _result := int32(C.gsl_fit_mul((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(xstride), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(ystride), C.size_t(n), &_outptr_5, &_outptr_6, &_outptr_7))
   return _result, *(*float64)(unsafe.Pointer(&_outptr_5)), *(*float64)(unsafe.Pointer(&_outptr_6)), *(*float64)(unsafe.Pointer(&_outptr_7))
}

func Wmul(x []float64, xstride int, w []float64, wstride int, y []float64, ystride int, n int) (int32, float64, float64, float64) {
   var _outptr_7 C.double
   var _outptr_8 C.double
   var _outptr_9 C.double
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&x))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&w))
   _slice_header_4 := (*reflect.SliceHeader)(unsafe.Pointer(&y))
   _result := int32(C.gsl_fit_wmul((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(xstride), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(wstride), (*C.double)(unsafe.Pointer(_slice_header_4.Data)), C.size_t(ystride), C.size_t(n), &_outptr_7, &_outptr_8, &_outptr_9))
   return _result, *(*float64)(unsafe.Pointer(&_outptr_7)), *(*float64)(unsafe.Pointer(&_outptr_8)), *(*float64)(unsafe.Pointer(&_outptr_9))
}

func MulEst(x float64, c1 float64, cov11 float64) (int32, float64, float64) {
   var _outptr_3 C.double
   var _outptr_4 C.double
   _result := int32(C.gsl_fit_mul_est(C.double(x), C.double(c1), C.double(cov11), &_outptr_3, &_outptr_4))
   return _result, *(*float64)(unsafe.Pointer(&_outptr_3)), *(*float64)(unsafe.Pointer(&_outptr_4))
}

func GslMultifitLinearAlloc(n int, p int) *GslMultifitLinearWorkspace {
   _ref := C.gsl_multifit_linear_alloc(C.size_t(n), C.size_t(p))
   _result := &GslMultifitLinearWorkspace{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func (x *GslMultifitLinearWorkspace) Dispose() {
   C.gsl_multifit_linear_free((*C.gsl_multifit_linear_workspace)(unsafe.Pointer(x.Ptr())))
}

func GslMultifitLinear(x *matrix.GslMatrix, y *vector.GslVector, c *vector.GslVector, cov *matrix.GslMatrix, work *GslMultifitLinearWorkspace) (int32, float64) {
   var _outptr_4 C.double
   _result := int32(C.gsl_multifit_linear((*C.gsl_matrix)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector)(unsafe.Pointer(y.Ptr())), (*C.gsl_vector)(unsafe.Pointer(c.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(cov.Ptr())), &_outptr_4, (*C.gsl_multifit_linear_workspace)(unsafe.Pointer(work.Ptr()))))
   return _result, *(*float64)(unsafe.Pointer(&_outptr_4))
}

func GslMultifitWlinear(x *matrix.GslMatrix, w *vector.GslVector, y *vector.GslVector, c *vector.GslVector, cov *matrix.GslMatrix, work *GslMultifitLinearWorkspace) (int32, float64) {
   var _outptr_5 C.double
   _result := int32(C.gsl_multifit_wlinear((*C.gsl_matrix)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector)(unsafe.Pointer(w.Ptr())), (*C.gsl_vector)(unsafe.Pointer(y.Ptr())), (*C.gsl_vector)(unsafe.Pointer(c.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(cov.Ptr())), &_outptr_5, (*C.gsl_multifit_linear_workspace)(unsafe.Pointer(work.Ptr()))))
   return _result, *(*float64)(unsafe.Pointer(&_outptr_5))
}

func GslMultifitLinearSvd(x *matrix.GslMatrix, y *vector.GslVector, tol float64, c *vector.GslVector, cov *matrix.GslMatrix, work *GslMultifitLinearWorkspace) (int32, int, float64) {
   var _outptr_3 C.size_t
   var _outptr_6 C.double
   _result := int32(C.gsl_multifit_linear_svd((*C.gsl_matrix)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector)(unsafe.Pointer(y.Ptr())), C.double(tol), &_outptr_3, (*C.gsl_vector)(unsafe.Pointer(c.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(cov.Ptr())), &_outptr_6, (*C.gsl_multifit_linear_workspace)(unsafe.Pointer(work.Ptr()))))
   return _result, *(*int)(unsafe.Pointer(&_outptr_3)), *(*float64)(unsafe.Pointer(&_outptr_6))
}

func GslMultifitWlinearSvd(x *matrix.GslMatrix, w *vector.GslVector, y *vector.GslVector, tol float64, c *vector.GslVector, cov *matrix.GslMatrix, work *GslMultifitLinearWorkspace) (int32, int, float64) {
   var _outptr_4 C.size_t
   var _outptr_7 C.double
   _result := int32(C.gsl_multifit_wlinear_svd((*C.gsl_matrix)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector)(unsafe.Pointer(w.Ptr())), (*C.gsl_vector)(unsafe.Pointer(y.Ptr())), C.double(tol), &_outptr_4, (*C.gsl_vector)(unsafe.Pointer(c.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(cov.Ptr())), &_outptr_7, (*C.gsl_multifit_linear_workspace)(unsafe.Pointer(work.Ptr()))))
   return _result, *(*int)(unsafe.Pointer(&_outptr_4)), *(*float64)(unsafe.Pointer(&_outptr_7))
}

func GslMultifitLinearUsvd(x *matrix.GslMatrix, y *vector.GslVector, tol float64, c *vector.GslVector, cov *matrix.GslMatrix, work *GslMultifitLinearWorkspace) (int32, int, float64) {
   var _outptr_3 C.size_t
   var _outptr_6 C.double
   _result := int32(C.gsl_multifit_linear_usvd((*C.gsl_matrix)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector)(unsafe.Pointer(y.Ptr())), C.double(tol), &_outptr_3, (*C.gsl_vector)(unsafe.Pointer(c.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(cov.Ptr())), &_outptr_6, (*C.gsl_multifit_linear_workspace)(unsafe.Pointer(work.Ptr()))))
   return _result, *(*int)(unsafe.Pointer(&_outptr_3)), *(*float64)(unsafe.Pointer(&_outptr_6))
}

func GslMultifitWlinearUsvd(x *matrix.GslMatrix, w *vector.GslVector, y *vector.GslVector, tol float64, c *vector.GslVector, cov *matrix.GslMatrix, work *GslMultifitLinearWorkspace) (int32, int, float64) {
   var _outptr_4 C.size_t
   var _outptr_7 C.double
   _result := int32(C.gsl_multifit_wlinear_usvd((*C.gsl_matrix)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector)(unsafe.Pointer(w.Ptr())), (*C.gsl_vector)(unsafe.Pointer(y.Ptr())), C.double(tol), &_outptr_4, (*C.gsl_vector)(unsafe.Pointer(c.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(cov.Ptr())), &_outptr_7, (*C.gsl_multifit_linear_workspace)(unsafe.Pointer(work.Ptr()))))
   return _result, *(*int)(unsafe.Pointer(&_outptr_4)), *(*float64)(unsafe.Pointer(&_outptr_7))
}

func GslMultifitLinearEst(x *vector.GslVector, c *vector.GslVector, cov *matrix.GslMatrix) (int32, float64, float64) {
   var _outptr_3 C.double
   var _outptr_4 C.double
   _result := int32(C.gsl_multifit_linear_est((*C.gsl_vector)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector)(unsafe.Pointer(c.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(cov.Ptr())), &_outptr_3, &_outptr_4))
   return _result, *(*float64)(unsafe.Pointer(&_outptr_3)), *(*float64)(unsafe.Pointer(&_outptr_4))
}

func GslMultifitLinearResiduals(x *matrix.GslMatrix, y *vector.GslVector, c *vector.GslVector, r *vector.GslVector) int32 {
   return int32(C.gsl_multifit_linear_residuals((*C.gsl_matrix)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector)(unsafe.Pointer(y.Ptr())), (*C.gsl_vector)(unsafe.Pointer(c.Ptr())), (*C.gsl_vector)(unsafe.Pointer(r.Ptr()))))
}

func GslMultifitRobustAlloc(t *GslMultifitRobustType, n int, p int) *GslMultifitRobustWorkspace {
   _ref := C.gsl_multifit_robust_alloc((*C.gsl_multifit_robust_type)(unsafe.Pointer(t.Ptr())), C.size_t(n), C.size_t(p))
   _result := &GslMultifitRobustWorkspace{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func (x *GslMultifitRobustWorkspace) Dispose() {
   C.gsl_multifit_robust_free((*C.gsl_multifit_robust_workspace)(unsafe.Pointer(x.Ptr())))
}

func GslMultifitRobustName(w *GslMultifitRobustWorkspace) string {
   return C.GoString(C.gsl_multifit_robust_name((*C.gsl_multifit_robust_workspace)(unsafe.Pointer(w.Ptr()))))
}

func GslMultifitRobustTune(tune float64, w *GslMultifitRobustWorkspace) int32 {
   return int32(C.gsl_multifit_robust_tune(C.double(tune), (*C.gsl_multifit_robust_workspace)(unsafe.Pointer(w.Ptr()))))
}

func GslMultifitRobust(x *matrix.GslMatrix, y *vector.GslVector, c *vector.GslVector, cov *matrix.GslMatrix, w *GslMultifitRobustWorkspace) int32 {
   return int32(C.gsl_multifit_robust((*C.gsl_matrix)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector)(unsafe.Pointer(y.Ptr())), (*C.gsl_vector)(unsafe.Pointer(c.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(cov.Ptr())), (*C.gsl_multifit_robust_workspace)(unsafe.Pointer(w.Ptr()))))
}

func GslMultifitRobustEst(x *vector.GslVector, c *vector.GslVector, cov *matrix.GslMatrix) (int32, float64, float64) {
   var _outptr_3 C.double
   var _outptr_4 C.double
   _result := int32(C.gsl_multifit_robust_est((*C.gsl_vector)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector)(unsafe.Pointer(c.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(cov.Ptr())), &_outptr_3, &_outptr_4))
   return _result, *(*float64)(unsafe.Pointer(&_outptr_3)), *(*float64)(unsafe.Pointer(&_outptr_4))
}

