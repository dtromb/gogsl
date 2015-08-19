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
import "reflect"

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
