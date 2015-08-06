//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package bspline

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_bspline.h>
*/
import "C"

import "unsafe"
import "github.com/dtromb/gogsl"
import "github.com/dtromb/gogsl/vector"
import "github.com/dtromb/gogsl/matrix"

type GslBsplineWorkspace struct {
   gogsl.GslReference
}

type GslBsplineDerivWorkspace struct {
   gogsl.GslReference
}

func Alloc(k int, nbreak int) *GslBsplineWorkspace {
   _ref := C.gsl_bspline_alloc(C.size_t(k), C.size_t(nbreak))
   _result := &GslBsplineWorkspace{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func (x *GslBsplineWorkspace) Dispose() {
   C.gsl_bspline_free((*C.gsl_bspline_workspace)(unsafe.Pointer(x.Ptr())))
}

func DerivAlloc(k int) *GslBsplineDerivWorkspace {
   _ref := C.gsl_bspline_deriv_alloc(C.size_t(k))
   _result := &GslBsplineDerivWorkspace{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func (x *GslBsplineDerivWorkspace) Dispose() {
   C.gsl_bspline_deriv_free((*C.gsl_bspline_deriv_workspace)(unsafe.Pointer(x.Ptr())))
}

func Knots(breakpts *vector.GslVector, w *GslBsplineWorkspace) int32 {
   return int32(C.gsl_bspline_knots((*C.gsl_vector)(unsafe.Pointer(breakpts.Ptr())), (*C.gsl_bspline_workspace)(unsafe.Pointer(w.Ptr()))))
}

func KnotsUniform(a float64, b float64, w *GslBsplineWorkspace) int32 {
   return int32(C.gsl_bspline_knots_uniform(C.double(a), C.double(b), (*C.gsl_bspline_workspace)(unsafe.Pointer(w.Ptr()))))
}

func Eval(x float64, b *vector.GslVector, w *GslBsplineWorkspace) int32 {
   return int32(C.gsl_bspline_eval(C.double(x), (*C.gsl_vector)(unsafe.Pointer(b.Ptr())), (*C.gsl_bspline_workspace)(unsafe.Pointer(w.Ptr()))))
}

func EvalNonzero(x float64, bk *vector.GslVector, w *GslBsplineWorkspace) (int32, int, int) {
   var _outptr_2 C.size_t
   var _outptr_3 C.size_t
   _result := int32(C.gsl_bspline_eval_nonzero(C.double(x), (*C.gsl_vector)(unsafe.Pointer(bk.Ptr())), &_outptr_2, &_outptr_3, (*C.gsl_bspline_workspace)(unsafe.Pointer(w.Ptr()))))
   return _result, *(*int)(unsafe.Pointer(&_outptr_2)), *(*int)(unsafe.Pointer(&_outptr_3))
}

func Ncoeffs(w *GslBsplineWorkspace) int {
   return int(C.gsl_bspline_ncoeffs((*C.gsl_bspline_workspace)(unsafe.Pointer(w.Ptr()))))
}

func DerivEval(x float64, nderiv int, dB *matrix.GslMatrix, w *GslBsplineWorkspace, dw *GslBsplineDerivWorkspace) int32 {
   return int32(C.gsl_bspline_deriv_eval(C.double(x), C.size_t(nderiv), (*C.gsl_matrix)(unsafe.Pointer(dB.Ptr())), (*C.gsl_bspline_workspace)(unsafe.Pointer(w.Ptr())), (*C.gsl_bspline_deriv_workspace)(unsafe.Pointer(dw.Ptr()))))
}

func DerivEvalNonzero(x float64, nderiv int, dB *matrix.GslMatrix, w *GslBsplineWorkspace, dw *GslBsplineDerivWorkspace) (int32, int, int) {
   var _outptr_3 C.size_t
   var _outptr_4 C.size_t
   _result := int32(C.gsl_bspline_deriv_eval_nonzero(C.double(x), C.size_t(nderiv), (*C.gsl_matrix)(unsafe.Pointer(dB.Ptr())), &_outptr_3, &_outptr_4, (*C.gsl_bspline_workspace)(unsafe.Pointer(w.Ptr())), (*C.gsl_bspline_deriv_workspace)(unsafe.Pointer(dw.Ptr()))))
   return _result, *(*int)(unsafe.Pointer(&_outptr_3)), *(*int)(unsafe.Pointer(&_outptr_4))
}

func GrevilleAbscissa(i int, w *GslBsplineWorkspace) float64 {
   return float64(C.gsl_bspline_greville_abscissa(C.size_t(i), (*C.gsl_bspline_workspace)(unsafe.Pointer(w.Ptr()))))
}

