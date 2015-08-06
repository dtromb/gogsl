//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package poly

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_poly.h>
*/
import "C"

import "unsafe"
import "github.com/dtromb/gogsl"
import "reflect"
import complex_ "github.com/dtromb/gogsl/complex"

type GslPolyComplexWorkspace struct {
   gogsl.GslReference
}

func Eval(c []float64, len int32, x float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&c))
   return float64(C.gsl_poly_eval((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.int(len), C.double(x)))
}

func ComplexEval(c []float64, len int32, z complex128) complex128 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&c))
   _arg_2 := complex_.GoComplexToGsl(z)
   _result := C.gsl_poly_complex_eval((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.int(len), *(*C.gsl_complex)(unsafe.Pointer(_arg_2)))
   return complex_.GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func ComplexPolyComplexEval(c []complex128, len int32, z complex128) complex128 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&c))
   _arg_2 := complex_.GoComplexToGsl(z)
   _result := C.gsl_complex_poly_complex_eval((*C.gsl_complex)(unsafe.Pointer(_slice_header_0.Data)), C.int(len), *(*C.gsl_complex)(unsafe.Pointer(_arg_2)))
   return complex_.GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func EvalDerivs(c []float64, lenc int, x float64, res []float64, lenres int) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&c))
   _slice_header_3 := (*reflect.SliceHeader)(unsafe.Pointer(&res))
   return int32(C.gsl_poly_eval_derivs((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(lenc), C.double(x), (*C.double)(unsafe.Pointer(_slice_header_3.Data)), C.size_t(lenres)))
}

func DdInit(dd []float64, xa []float64, ya []float64, size int) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&dd))
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&xa))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&ya))
   return int32(C.gsl_poly_dd_init((*C.double)(unsafe.Pointer(_slice_header_0.Data)), (*C.double)(unsafe.Pointer(_slice_header_1.Data)), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(size)))
}

func DdEval(dd []float64, xa []float64, size int, x float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&dd))
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&xa))
   return float64(C.gsl_poly_dd_eval((*C.double)(unsafe.Pointer(_slice_header_0.Data)), (*C.double)(unsafe.Pointer(_slice_header_1.Data)), C.size_t(size), C.double(x)))
}

func DdTaylor(c []float64, xp float64, dd []float64, xa []float64, size int, w []float64) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&c))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&dd))
   _slice_header_3 := (*reflect.SliceHeader)(unsafe.Pointer(&xa))
   _slice_header_5 := (*reflect.SliceHeader)(unsafe.Pointer(&w))
   return int32(C.gsl_poly_dd_taylor((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.double(xp), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), (*C.double)(unsafe.Pointer(_slice_header_3.Data)), C.size_t(size), (*C.double)(unsafe.Pointer(_slice_header_5.Data))))
}

func DdHermiteInit(dd []float64, za []float64, xa []float64, ya []float64, dya []float64, size int) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&dd))
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&za))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&xa))
   _slice_header_3 := (*reflect.SliceHeader)(unsafe.Pointer(&ya))
   _slice_header_4 := (*reflect.SliceHeader)(unsafe.Pointer(&dya))
   return int32(C.gsl_poly_dd_hermite_init((*C.double)(unsafe.Pointer(_slice_header_0.Data)), (*C.double)(unsafe.Pointer(_slice_header_1.Data)), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), (*C.double)(unsafe.Pointer(_slice_header_3.Data)), (*C.double)(unsafe.Pointer(_slice_header_4.Data)), C.size_t(size)))
}

func ComplexSolveQuadratic(a float64, b float64, c float64) (int32, complex128, complex128) {
   var _outptr_3 C.gsl_complex
   var _outptr_4 C.gsl_complex
   _result := int32(C.gsl_poly_complex_solve_quadratic(C.double(a), C.double(b), C.double(c), &_outptr_3, &_outptr_4))
   return _result, *(*complex128)(unsafe.Pointer(&_outptr_3)), *(*complex128)(unsafe.Pointer(&_outptr_4))
}

func ComplexSolveCubic(a float64, b float64, c float64) (int32, complex128, complex128, complex128) {
   var _outptr_3 C.gsl_complex
   var _outptr_4 C.gsl_complex
   var _outptr_5 C.gsl_complex
   _result := int32(C.gsl_poly_complex_solve_cubic(C.double(a), C.double(b), C.double(c), &_outptr_3, &_outptr_4, &_outptr_5))
   return _result, *(*complex128)(unsafe.Pointer(&_outptr_3)), *(*complex128)(unsafe.Pointer(&_outptr_4)), *(*complex128)(unsafe.Pointer(&_outptr_5))
}

func ComplexWorkspaceAlloc(n int) *GslPolyComplexWorkspace {
   _ref := C.gsl_poly_complex_workspace_alloc(C.size_t(n))
   _result := &GslPolyComplexWorkspace{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func (x *GslPolyComplexWorkspace) Dispose() {
   C.gsl_poly_complex_workspace_free((*C.gsl_poly_complex_workspace)(unsafe.Pointer(x.Ptr())))
}

func ComplexSolve(a []float64, n int, w *GslPolyComplexWorkspace, z []complex128) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&a))
   _slice_header_3 := (*reflect.SliceHeader)(unsafe.Pointer(&z))
   return int32(C.gsl_poly_complex_solve((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(n), (*C.gsl_poly_complex_workspace)(unsafe.Pointer(w.Ptr())), (*C.double)(unsafe.Pointer(_slice_header_3.Data))))
}

