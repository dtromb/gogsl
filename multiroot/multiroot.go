//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package multiroot

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_multiroots.h>
   #include <gsl/gsl_vector.h>
   #include <gsl/gsl_matrix.h>
*/
import "C"

import "unsafe"
import "github.com/dtromb/gogsl"
import "github.com/dtromb/gogsl/vector"

type GslMultirootFsolver struct {
	gogsl.GslReference
}

type GslMultirootFsolverType struct {
	gogsl.GslReference
	cPtr uintptr
}

type GslMultirootFdfsolver struct {
	gogsl.GslReference
}

type GslMultirootFdfsolverType struct {
	gogsl.GslReference
	cPtr uintptr
}

func FsolverAlloc(t *GslMultirootFsolverType, n int) *GslMultirootFsolver {
	_ref := C.gsl_multiroot_fsolver_alloc((*C.gsl_multiroot_fsolver_type)(unsafe.Pointer(t.Ptr())), C.size_t(n))
	_result := &GslMultirootFsolver{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func FdfsolverAlloc(t *GslMultirootFdfsolverType, n int) *GslMultirootFdfsolver {
	_ref := C.gsl_multiroot_fdfsolver_alloc((*C.gsl_multiroot_fdfsolver_type)(unsafe.Pointer(t.Ptr())), C.size_t(n))
	_result := &GslMultirootFdfsolver{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func FsolverSet(s *GslMultirootFsolver, f *GslMultirootFunction, x *vector.GslVector) int32 {
	return int32(C.gsl_multiroot_fsolver_set((*C.gsl_multiroot_fsolver)(unsafe.Pointer(s.Ptr())), (*C.gsl_multiroot_function)(unsafe.Pointer(f.CPtr())), (*C.gsl_vector)(unsafe.Pointer(x.Ptr()))))
}

func FdfsolverSet(s *GslMultirootFdfsolver, fdf *GslMultirootFunctionFdf, x *vector.GslVector) int32 {
	return int32(C.gsl_multiroot_fdfsolver_set((*C.gsl_multiroot_fdfsolver)(unsafe.Pointer(s.Ptr())), (*C.gsl_multiroot_function_fdf)(unsafe.Pointer(fdf.CPtr())), (*C.gsl_vector)(unsafe.Pointer(x.Ptr()))))
}

func (x *GslMultirootFsolver) Dispose() {
	C.gsl_multiroot_fsolver_free((*C.gsl_multiroot_fsolver)(unsafe.Pointer(x.Ptr())))
}

func (x *GslMultirootFdfsolver) Dispose() {
	C.gsl_multiroot_fdfsolver_free((*C.gsl_multiroot_fdfsolver)(unsafe.Pointer(x.Ptr())))
}

func FsolverName(s *GslMultirootFsolver) string {
	return C.GoString(C.gsl_multiroot_fsolver_name((*C.gsl_multiroot_fsolver)(unsafe.Pointer(s.Ptr()))))
}

func FdfsolverName(s *GslMultirootFdfsolver) string {
	return C.GoString(C.gsl_multiroot_fdfsolver_name((*C.gsl_multiroot_fdfsolver)(unsafe.Pointer(s.Ptr()))))
}

func FsolverIterate(s *GslMultirootFsolver) int32 {
	return int32(C.gsl_multiroot_fsolver_iterate((*C.gsl_multiroot_fsolver)(unsafe.Pointer(s.Ptr()))))
}

func FdfsolverIterate(s *GslMultirootFdfsolver) int32 {
	return int32(C.gsl_multiroot_fdfsolver_iterate((*C.gsl_multiroot_fdfsolver)(unsafe.Pointer(s.Ptr()))))
}

func FsolverRoot(s *GslMultirootFsolver) *vector.GslVector {
	_ref := C.gsl_multiroot_fsolver_root((*C.gsl_multiroot_fsolver)(unsafe.Pointer(s.Ptr())))
	_result := &vector.GslVector{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func FdfsolverRoot(s *GslMultirootFdfsolver) *vector.GslVector {
	_ref := C.gsl_multiroot_fdfsolver_root((*C.gsl_multiroot_fdfsolver)(unsafe.Pointer(s.Ptr())))
	_result := &vector.GslVector{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func FsolverF(s *GslMultirootFsolver) *vector.GslVector {
	_ref := C.gsl_multiroot_fsolver_f((*C.gsl_multiroot_fsolver)(unsafe.Pointer(s.Ptr())))
	_result := &vector.GslVector{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func FdfsolverF(s *GslMultirootFdfsolver) *vector.GslVector {
	_ref := C.gsl_multiroot_fdfsolver_f((*C.gsl_multiroot_fdfsolver)(unsafe.Pointer(s.Ptr())))
	_result := &vector.GslVector{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func FsolverDx(s *GslMultirootFsolver) *vector.GslVector {
	_ref := C.gsl_multiroot_fsolver_dx((*C.gsl_multiroot_fsolver)(unsafe.Pointer(s.Ptr())))
	_result := &vector.GslVector{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func FdfsolverDx(s *GslMultirootFdfsolver) *vector.GslVector {
	_ref := C.gsl_multiroot_fdfsolver_dx((*C.gsl_multiroot_fdfsolver)(unsafe.Pointer(s.Ptr())))
	_result := &vector.GslVector{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func TestDelta(dx *vector.GslVector, x *vector.GslVector, epsabs float64, epsrel float64) int32 {
	return int32(C.gsl_multiroot_test_delta((*C.gsl_vector)(unsafe.Pointer(dx.Ptr())), (*C.gsl_vector)(unsafe.Pointer(x.Ptr())), C.double(epsabs), C.double(epsrel)))
}

func TestResidual(f *vector.GslVector, epsabs float64) int32 {
	return int32(C.gsl_multiroot_test_residual((*C.gsl_vector)(unsafe.Pointer(f.Ptr())), C.double(epsabs)))
}
