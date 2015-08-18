//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package multimin

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_multimin.h>
*/
import "C"

import "unsafe"
import "github.com/dtromb/gogsl"
import "github.com/dtromb/gogsl/vector"

type GslMultiminFminimizer struct {
	gogsl.GslReference
}

type GslMultiminFminimizerType struct {
	gogsl.GslReference
	cPtr uintptr
}

type GslMultiminFdfminimizer struct {
	gogsl.GslReference
}

type GslMultiminFdfminimizerType struct {
	gogsl.GslReference
	cPtr uintptr
}

func FdfminimizerAlloc(t *GslMultiminFdfminimizerType, n int) *GslMultiminFdfminimizer {
	_ref := C.gsl_multimin_fdfminimizer_alloc((*C.gsl_multimin_fdfminimizer_type)(unsafe.Pointer(t.Ptr())), C.size_t(n))
	_result := &GslMultiminFdfminimizer{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func FminimizerAlloc(t *GslMultiminFminimizerType, n int) *GslMultiminFminimizer {
	_ref := C.gsl_multimin_fminimizer_alloc((*C.gsl_multimin_fminimizer_type)(unsafe.Pointer(t.Ptr())), C.size_t(n))
	_result := &GslMultiminFminimizer{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func FdfminimizerSet(s *GslMultiminFdfminimizer, fdf *GslMultiminFunctionFdf, x *vector.GslVector, stepSize float64, tol float64) int32 {
	return int32(C.gsl_multimin_fdfminimizer_set((*C.gsl_multimin_fdfminimizer)(unsafe.Pointer(s.Ptr())), (*C.gsl_multimin_function_fdf)(unsafe.Pointer(fdf.CPtr())), (*C.gsl_vector)(unsafe.Pointer(x.Ptr())), C.double(stepSize), C.double(tol)))
}

func FminimizerSet(s *GslMultiminFminimizer, f *GslMultiminFunction, x *vector.GslVector, stepSize *vector.GslVector) int32 {
	return int32(C.gsl_multimin_fminimizer_set((*C.gsl_multimin_fminimizer)(unsafe.Pointer(s.Ptr())), (*C.gsl_multimin_function)(unsafe.Pointer(f.CPtr())), (*C.gsl_vector)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector)(unsafe.Pointer(stepSize.Ptr()))))
}

func (x *GslMultiminFdfminimizer) Dispose() {
	C.gsl_multimin_fdfminimizer_free((*C.gsl_multimin_fdfminimizer)(unsafe.Pointer(x.Ptr())))
}

func (x *GslMultiminFminimizer) Dispose() {
	C.gsl_multimin_fminimizer_free((*C.gsl_multimin_fminimizer)(unsafe.Pointer(x.Ptr())))
}

func FdfminimizerName(s *GslMultiminFdfminimizer) string {
	return C.GoString(C.gsl_multimin_fdfminimizer_name((*C.gsl_multimin_fdfminimizer)(unsafe.Pointer(s.Ptr()))))
}

func FminimizerName(s *GslMultiminFminimizer) string {
	return C.GoString(C.gsl_multimin_fminimizer_name((*C.gsl_multimin_fminimizer)(unsafe.Pointer(s.Ptr()))))
}

func FdfminimizerIterate(s *GslMultiminFdfminimizer) int32 {
	return int32(C.gsl_multimin_fdfminimizer_iterate((*C.gsl_multimin_fdfminimizer)(unsafe.Pointer(s.Ptr()))))
}

func FminimizerIterate(s *GslMultiminFminimizer) int32 {
	return int32(C.gsl_multimin_fminimizer_iterate((*C.gsl_multimin_fminimizer)(unsafe.Pointer(s.Ptr()))))
}

func FdfminimizerX(s *GslMultiminFdfminimizer) *vector.GslVector {
	_ref := C.gsl_multimin_fdfminimizer_x((*C.gsl_multimin_fdfminimizer)(unsafe.Pointer(s.Ptr())))
	_result := &vector.GslVector{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func FminimizerX(s *GslMultiminFminimizer) *vector.GslVector {
	_ref := C.gsl_multimin_fminimizer_x((*C.gsl_multimin_fminimizer)(unsafe.Pointer(s.Ptr())))
	_result := &vector.GslVector{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func FdfminimizerMinimum(s *GslMultiminFdfminimizer) float64 {
	return float64(C.gsl_multimin_fdfminimizer_minimum((*C.gsl_multimin_fdfminimizer)(unsafe.Pointer(s.Ptr()))))
}

func FminimizerMinimum(s *GslMultiminFminimizer) float64 {
	return float64(C.gsl_multimin_fminimizer_minimum((*C.gsl_multimin_fminimizer)(unsafe.Pointer(s.Ptr()))))
}

func FdfminimizerGradient(s *GslMultiminFdfminimizer) *vector.GslVector {
	_ref := C.gsl_multimin_fdfminimizer_gradient((*C.gsl_multimin_fdfminimizer)(unsafe.Pointer(s.Ptr())))
	_result := &vector.GslVector{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func FminimizerSize(s *GslMultiminFminimizer) float64 {
	return float64(C.gsl_multimin_fminimizer_size((*C.gsl_multimin_fminimizer)(unsafe.Pointer(s.Ptr()))))
}

func FdfminimizerRestart(s *GslMultiminFdfminimizer) int32 {
	return int32(C.gsl_multimin_fdfminimizer_restart((*C.gsl_multimin_fdfminimizer)(unsafe.Pointer(s.Ptr()))))
}

func TestGradient(g *vector.GslVector, epsabs float64) int32 {
	return int32(C.gsl_multimin_test_gradient((*C.gsl_vector)(unsafe.Pointer(g.Ptr())), C.double(epsabs)))
}

func TestSize(size float64, epsabs float64) int32 {
	return int32(C.gsl_multimin_test_size(C.double(size), C.double(epsabs)))
}
