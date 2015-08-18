//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package interp

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_interp.h>
   #include <gsl/gsl_spline.h>
*/
import "C"

import "unsafe"
import "github.com/dtromb/gogsl"
import "reflect"

type GslInterp struct {
   gogsl.GslReference
}

type GslInterpType struct {
   gogsl.GslReference
   cPtr uintptr
}

type GslInterpAccel struct {
   gogsl.GslReference
}

type GslSpline struct {
   gogsl.GslReference
}

func Alloc(t *GslInterpType, size int) *GslInterp {
   _ref := C.gsl_interp_alloc((*C.gsl_interp_type)(unsafe.Pointer(t.Ptr())), C.size_t(size))
   _result := &GslInterp{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func Init(interp *GslInterp, xa []float64, ya []float64, size int) int32 {
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&xa))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&ya))
   return int32(C.gsl_interp_init((*C.gsl_interp)(unsafe.Pointer(interp.Ptr())), (*C.double)(unsafe.Pointer(_slice_header_1.Data)), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(size)))
}

func (x *GslInterp) Dispose() {
   C.gsl_interp_free((*C.gsl_interp)(unsafe.Pointer(x.Ptr())))
}

func Name(interp *GslInterp) string {
   return C.GoString(C.gsl_interp_name((*C.gsl_interp)(unsafe.Pointer(interp.Ptr()))))
}

func MinSize(interp *GslInterp) uint32 {
   return uint32(C.gsl_interp_min_size((*C.gsl_interp)(unsafe.Pointer(interp.Ptr()))))
}

func TypeMinSize(t *GslInterpType) uint32 {
   return uint32(C.gsl_interp_type_min_size((*C.gsl_interp_type)(unsafe.Pointer(t.Ptr()))))
}

func Bsearch(xArray []float64, x float64, indexLo int, indexHi int) int {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&xArray))
   return int(C.gsl_interp_bsearch((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.double(x), C.size_t(indexLo), C.size_t(indexHi)))
}

func AccelAlloc() *GslInterpAccel {
   _ref := C.gsl_interp_accel_alloc()
   _result := &GslInterpAccel{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func AccelFind(a *GslInterpAccel, xArray []float64, size int, x float64) int {
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&xArray))
   return int(C.gsl_interp_accel_find((*C.gsl_interp_accel)(unsafe.Pointer(a.Ptr())), (*C.double)(unsafe.Pointer(_slice_header_1.Data)), C.size_t(size), C.double(x)))
}

func AccelReset(acc *GslInterpAccel) int32 {
   return int32(C.gsl_interp_accel_reset((*C.gsl_interp_accel)(unsafe.Pointer(acc.Ptr()))))
}

func (x *GslInterpAccel) Dispose() {
   C.gsl_interp_accel_free((*C.gsl_interp_accel)(unsafe.Pointer(x.Ptr())))
}

func Eval(interp *GslInterp, xa []float64, ya []float64, x float64, acc *GslInterpAccel) float64 {
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&xa))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&ya))
   return float64(C.gsl_interp_eval((*C.gsl_interp)(unsafe.Pointer(interp.Ptr())), (*C.double)(unsafe.Pointer(_slice_header_1.Data)), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.double(x), (*C.gsl_interp_accel)(unsafe.Pointer(acc.Ptr()))))
}

func EvalE(interp *GslInterp, xa []float64, ya []float64, x float64, acc *GslInterpAccel) (int32, float64) {
   var _outptr_5 C.double
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&xa))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&ya))
   _result := int32(C.gsl_interp_eval_e((*C.gsl_interp)(unsafe.Pointer(interp.Ptr())), (*C.double)(unsafe.Pointer(_slice_header_1.Data)), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.double(x), (*C.gsl_interp_accel)(unsafe.Pointer(acc.Ptr())), &_outptr_5))
   return _result, *(*float64)(unsafe.Pointer(&_outptr_5))
}

func EvalDeriv(interp *GslInterp, xa []float64, ya []float64, x float64, acc *GslInterpAccel) float64 {
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&xa))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&ya))
   return float64(C.gsl_interp_eval_deriv((*C.gsl_interp)(unsafe.Pointer(interp.Ptr())), (*C.double)(unsafe.Pointer(_slice_header_1.Data)), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.double(x), (*C.gsl_interp_accel)(unsafe.Pointer(acc.Ptr()))))
}

func EvalDerivE(interp *GslInterp, xa []float64, ya []float64, x float64, acc *GslInterpAccel) (int32, float64) {
   var _outptr_5 C.double
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&xa))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&ya))
   _result := int32(C.gsl_interp_eval_deriv_e((*C.gsl_interp)(unsafe.Pointer(interp.Ptr())), (*C.double)(unsafe.Pointer(_slice_header_1.Data)), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.double(x), (*C.gsl_interp_accel)(unsafe.Pointer(acc.Ptr())), &_outptr_5))
   return _result, *(*float64)(unsafe.Pointer(&_outptr_5))
}

func EvalDeriv2(interp *GslInterp, xa []float64, ya []float64, x float64, acc *GslInterpAccel) float64 {
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&xa))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&ya))
   return float64(C.gsl_interp_eval_deriv2((*C.gsl_interp)(unsafe.Pointer(interp.Ptr())), (*C.double)(unsafe.Pointer(_slice_header_1.Data)), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.double(x), (*C.gsl_interp_accel)(unsafe.Pointer(acc.Ptr()))))
}

func EvalDeriv2E(interp *GslInterp, xa []float64, ya []float64, x float64, acc *GslInterpAccel) (int32, float64) {
   var _outptr_5 C.double
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&xa))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&ya))
   _result := int32(C.gsl_interp_eval_deriv2_e((*C.gsl_interp)(unsafe.Pointer(interp.Ptr())), (*C.double)(unsafe.Pointer(_slice_header_1.Data)), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.double(x), (*C.gsl_interp_accel)(unsafe.Pointer(acc.Ptr())), &_outptr_5))
   return _result, *(*float64)(unsafe.Pointer(&_outptr_5))
}

func EvalInteg(interp *GslInterp, xa []float64, ya []float64, a float64, b float64, acc *GslInterpAccel) float64 {
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&xa))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&ya))
   return float64(C.gsl_interp_eval_integ((*C.gsl_interp)(unsafe.Pointer(interp.Ptr())), (*C.double)(unsafe.Pointer(_slice_header_1.Data)), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.double(a), C.double(b), (*C.gsl_interp_accel)(unsafe.Pointer(acc.Ptr()))))
}

func EvalIntegE(interp *GslInterp, xa []float64, ya []float64, a float64, b float64, acc *GslInterpAccel) (int32, float64) {
   var _outptr_6 C.double
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&xa))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&ya))
   _result := int32(C.gsl_interp_eval_integ_e((*C.gsl_interp)(unsafe.Pointer(interp.Ptr())), (*C.double)(unsafe.Pointer(_slice_header_1.Data)), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.double(a), C.double(b), (*C.gsl_interp_accel)(unsafe.Pointer(acc.Ptr())), &_outptr_6))
   return _result, *(*float64)(unsafe.Pointer(&_outptr_6))
}

func SplineAlloc(t *GslInterpType, size int) *GslSpline {
   _ref := C.gsl_spline_alloc((*C.gsl_interp_type)(unsafe.Pointer(t.Ptr())), C.size_t(size))
   _result := &GslSpline{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func SplineInit(spline *GslSpline, xa []float64, ya []float64, size int) int32 {
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&xa))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&ya))
   return int32(C.gsl_spline_init((*C.gsl_spline)(unsafe.Pointer(spline.Ptr())), (*C.double)(unsafe.Pointer(_slice_header_1.Data)), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(size)))
}

func (x *GslSpline) Dispose() {
   C.gsl_spline_free((*C.gsl_spline)(unsafe.Pointer(x.Ptr())))
}

func SplineName(spline *GslSpline) string {
   return C.GoString(C.gsl_spline_name((*C.gsl_spline)(unsafe.Pointer(spline.Ptr()))))
}

func SplineMinSize(spline *GslSpline) uint32 {
   return uint32(C.gsl_spline_min_size((*C.gsl_spline)(unsafe.Pointer(spline.Ptr()))))
}

func SplineEval(spline *GslSpline, x float64, acc *GslInterpAccel) float64 {
   return float64(C.gsl_spline_eval((*C.gsl_spline)(unsafe.Pointer(spline.Ptr())), C.double(x), (*C.gsl_interp_accel)(unsafe.Pointer(acc.Ptr()))))
}

func SplineEvalE(spline *GslSpline, x float64, acc *GslInterpAccel) (int32, float64) {
   var _outptr_3 C.double
   _result := int32(C.gsl_spline_eval_e((*C.gsl_spline)(unsafe.Pointer(spline.Ptr())), C.double(x), (*C.gsl_interp_accel)(unsafe.Pointer(acc.Ptr())), &_outptr_3))
   return _result, *(*float64)(unsafe.Pointer(&_outptr_3))
}

func SplineEvalDeriv(spline *GslSpline, x float64, acc *GslInterpAccel) float64 {
   return float64(C.gsl_spline_eval_deriv((*C.gsl_spline)(unsafe.Pointer(spline.Ptr())), C.double(x), (*C.gsl_interp_accel)(unsafe.Pointer(acc.Ptr()))))
}

func SplineEvalDerivE(spline *GslSpline, x float64, acc *GslInterpAccel) (int32, float64) {
   var _outptr_3 C.double
   _result := int32(C.gsl_spline_eval_deriv_e((*C.gsl_spline)(unsafe.Pointer(spline.Ptr())), C.double(x), (*C.gsl_interp_accel)(unsafe.Pointer(acc.Ptr())), &_outptr_3))
   return _result, *(*float64)(unsafe.Pointer(&_outptr_3))
}

func SplineEvalDeriv2(spline *GslSpline, x float64, acc *GslInterpAccel) float64 {
   return float64(C.gsl_spline_eval_deriv2((*C.gsl_spline)(unsafe.Pointer(spline.Ptr())), C.double(x), (*C.gsl_interp_accel)(unsafe.Pointer(acc.Ptr()))))
}

func SplineEvalDeriv2E(spline *GslSpline, x float64, acc *GslInterpAccel) (int32, float64) {
   var _outptr_3 C.double
   _result := int32(C.gsl_spline_eval_deriv2_e((*C.gsl_spline)(unsafe.Pointer(spline.Ptr())), C.double(x), (*C.gsl_interp_accel)(unsafe.Pointer(acc.Ptr())), &_outptr_3))
   return _result, *(*float64)(unsafe.Pointer(&_outptr_3))
}

func SplineEvalInteg(spline *GslSpline, a float64, b float64, acc *GslInterpAccel) float64 {
   return float64(C.gsl_spline_eval_integ((*C.gsl_spline)(unsafe.Pointer(spline.Ptr())), C.double(a), C.double(b), (*C.gsl_interp_accel)(unsafe.Pointer(acc.Ptr()))))
}

func SplineEvalIntegE(spline *GslSpline, a float64, b float64, acc *GslInterpAccel) (int32, float64) {
   var _outptr_4 C.double
   _result := int32(C.gsl_spline_eval_integ_e((*C.gsl_spline)(unsafe.Pointer(spline.Ptr())), C.double(a), C.double(b), (*C.gsl_interp_accel)(unsafe.Pointer(acc.Ptr())), &_outptr_4))
   return _result, *(*float64)(unsafe.Pointer(&_outptr_4))
}

