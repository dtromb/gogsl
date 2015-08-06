//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package chebyshev

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <stdlib.h>
   #include <gsl/gsl_chebyshev.h>
*/
import "C"

import "unsafe"
import "github.com/dtromb/gogsl"

type GslChebSeries struct {
   gogsl.GslReference
}

func Alloc(n int) *GslChebSeries {
   _ref := C.gsl_cheb_alloc(C.size_t(n))
   _result := &GslChebSeries{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func (x *GslChebSeries) Dispose() {
   C.gsl_cheb_free((*C.gsl_cheb_series)(unsafe.Pointer(x.Ptr())))
}

func Init(cs *GslChebSeries, f *gogsl.GslFunction, a float64, b float64) int32 {
   gogsl.InitializeGslFunction(f)
   return int32(C.gsl_cheb_init((*C.gsl_cheb_series)(unsafe.Pointer(cs.Ptr())), (*C.gsl_function)(unsafe.Pointer(f.CPtr())), C.double(a), C.double(b)))
}

func Order(cs *GslChebSeries) int {
   return int(C.gsl_cheb_order((*C.gsl_cheb_series)(unsafe.Pointer(cs.Ptr()))))
}

func Size(cs *GslChebSeries) int {
   return int(C.gsl_cheb_size((*C.gsl_cheb_series)(unsafe.Pointer(cs.Ptr()))))
}

func Eval(cs *GslChebSeries, x float64) float64 {
   return float64(C.gsl_cheb_eval((*C.gsl_cheb_series)(unsafe.Pointer(cs.Ptr())), C.double(x)))
}

func EvalErr(cs *GslChebSeries, x float64) (int32, float64, float64) {
   var _outptr_2 C.double
   var _outptr_3 C.double
   _result := int32(C.gsl_cheb_eval_err((*C.gsl_cheb_series)(unsafe.Pointer(cs.Ptr())), C.double(x), &_outptr_2, &_outptr_3))
   return _result, *(*float64)(unsafe.Pointer(&_outptr_2)), *(*float64)(unsafe.Pointer(&_outptr_3))
}

func EvalN(cs *GslChebSeries, order int, x float64) float64 {
   return float64(C.gsl_cheb_eval_n((*C.gsl_cheb_series)(unsafe.Pointer(cs.Ptr())), C.size_t(order), C.double(x)))
}

func EvalNErr(cs *GslChebSeries, order int, x float64) (int32, float64, float64) {
   var _outptr_3 C.double
   var _outptr_4 C.double
   _result := int32(C.gsl_cheb_eval_n_err((*C.gsl_cheb_series)(unsafe.Pointer(cs.Ptr())), C.size_t(order), C.double(x), &_outptr_3, &_outptr_4))
   return _result, *(*float64)(unsafe.Pointer(&_outptr_3)), *(*float64)(unsafe.Pointer(&_outptr_4))
}

func CalcDeriv(deriv *GslChebSeries, cs *GslChebSeries) int32 {
   return int32(C.gsl_cheb_calc_deriv((*C.gsl_cheb_series)(unsafe.Pointer(deriv.Ptr())), (*C.gsl_cheb_series)(unsafe.Pointer(cs.Ptr()))))
}

func CalcInteg(integ *GslChebSeries, cs *GslChebSeries) int32 {
   return int32(C.gsl_cheb_calc_integ((*C.gsl_cheb_series)(unsafe.Pointer(integ.Ptr())), (*C.gsl_cheb_series)(unsafe.Pointer(cs.Ptr()))))
}

