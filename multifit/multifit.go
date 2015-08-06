//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package multifit

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_multifit.h>
   #include <gsl/gsl_multifit_nlin.h>
*/
import "C"

import "unsafe"
import "github.com/dtromb/gogsl"
import "github.com/dtromb/gogsl/vector"
import "github.com/dtromb/gogsl/matrix"

type GslMultifitFsolver struct {
   gogsl.GslReference
}

type GslMultifitFsolverType struct {
   gogsl.GslReference
}

type GslMultifitFdfsolver struct {
   gogsl.GslReference
}

type GslMultifitFdfsolverType struct {
   gogsl.GslReference
}

func FsolverAlloc(t *GslMultifitFsolverType, n int, p int) *GslMultifitFsolver {
   _ref := C.gsl_multifit_fsolver_alloc((*C.gsl_multifit_fsolver_type)(unsafe.Pointer(t.Ptr())), C.size_t(n), C.size_t(p))
   _result := &GslMultifitFsolver{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func FdfsolverAlloc(t *GslMultifitFdfsolverType, n int, p int) *GslMultifitFdfsolver {
   _ref := C.gsl_multifit_fdfsolver_alloc((*C.gsl_multifit_fdfsolver_type)(unsafe.Pointer(t.Ptr())), C.size_t(n), C.size_t(p))
   _result := &GslMultifitFdfsolver{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func FsolverSet(s *GslMultifitFsolver, f *GslMultifitFunction, x *vector.GslVector) int32 {
   return int32(C.gsl_multifit_fsolver_set((*C.gsl_multifit_fsolver)(unsafe.Pointer(s.Ptr())), (*C.gsl_multifit_function)(unsafe.Pointer(f.CPtr())), (*C.gsl_vector)(unsafe.Pointer(x.Ptr()))))
}

func FdfsolverSet(s *GslMultifitFdfsolver, fdf *GslMultifitFunctionFdf, x *vector.GslVector) int32 {
   return int32(C.gsl_multifit_fdfsolver_set((*C.gsl_multifit_fdfsolver)(unsafe.Pointer(s.Ptr())), (*C.gsl_multifit_function_fdf)(unsafe.Pointer(fdf.CPtr())), (*C.gsl_vector)(unsafe.Pointer(x.Ptr()))))
}

func (x *GslMultifitFsolver) Dispose() {
   C.gsl_multifit_fsolver_free((*C.gsl_multifit_fsolver)(unsafe.Pointer(x.Ptr())))
}

func (x *GslMultifitFdfsolver) Dispose() {
   C.gsl_multifit_fdfsolver_free((*C.gsl_multifit_fdfsolver)(unsafe.Pointer(x.Ptr())))
}

func FsolverName(s *GslMultifitFsolver) string {
   return C.GoString(C.gsl_multifit_fsolver_name((*C.gsl_multifit_fsolver)(unsafe.Pointer(s.Ptr()))))
}

func FdfsolverName(s *GslMultifitFdfsolver) string {
   return C.GoString(C.gsl_multifit_fdfsolver_name((*C.gsl_multifit_fdfsolver)(unsafe.Pointer(s.Ptr()))))
}

func FdfsolverDifDf(x *vector.GslVector, fdf *GslMultifitFunctionFdf, f *vector.GslVector, j *matrix.GslMatrix) int32 {
   return int32(C.gsl_multifit_fdfsolver_dif_df((*C.gsl_vector)(unsafe.Pointer(x.Ptr())), (*C.gsl_multifit_function_fdf)(unsafe.Pointer(fdf.CPtr())), (*C.gsl_vector)(unsafe.Pointer(f.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(j.Ptr()))))
}

func FdfsolverDifFdf(x *vector.GslVector, fdf *GslMultifitFunctionFdf, f *vector.GslVector, j *matrix.GslMatrix) int32 {
   return int32(C.gsl_multifit_fdfsolver_dif_fdf((*C.gsl_vector)(unsafe.Pointer(x.Ptr())), (*C.gsl_multifit_function_fdf)(unsafe.Pointer(fdf.CPtr())), (*C.gsl_vector)(unsafe.Pointer(f.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(j.Ptr()))))
}

func FsolverIterate(s *GslMultifitFsolver) int32 {
   return int32(C.gsl_multifit_fsolver_iterate((*C.gsl_multifit_fsolver)(unsafe.Pointer(s.Ptr()))))
}

func FdfsolverIterate(s *GslMultifitFdfsolver) int32 {
   return int32(C.gsl_multifit_fdfsolver_iterate((*C.gsl_multifit_fdfsolver)(unsafe.Pointer(s.Ptr()))))
}

func FsolverPosition(s *GslMultifitFsolver) *vector.GslVector {
   _ref := C.gsl_multifit_fsolver_position((*C.gsl_multifit_fsolver)(unsafe.Pointer(s.Ptr())))
   _result := &vector.GslVector{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func FdfsolverPosition(s *GslMultifitFdfsolver) *vector.GslVector {
   _ref := C.gsl_multifit_fdfsolver_position((*C.gsl_multifit_fdfsolver)(unsafe.Pointer(s.Ptr())))
   _result := &vector.GslVector{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func TestDelta(dx *vector.GslVector, x *vector.GslVector, epsabs float64, epsrel float64) int32 {
   return int32(C.gsl_multifit_test_delta((*C.gsl_vector)(unsafe.Pointer(dx.Ptr())), (*C.gsl_vector)(unsafe.Pointer(x.Ptr())), C.double(epsabs), C.double(epsrel)))
}

func TestGradient(g *vector.GslVector, epsabs float64) int32 {
   return int32(C.gsl_multifit_test_gradient((*C.gsl_vector)(unsafe.Pointer(g.Ptr())), C.double(epsabs)))
}

func Gradient(j *matrix.GslMatrix, f *vector.GslVector, g *vector.GslVector) int32 {
   return int32(C.gsl_multifit_gradient((*C.gsl_matrix)(unsafe.Pointer(j.Ptr())), (*C.gsl_vector)(unsafe.Pointer(f.Ptr())), (*C.gsl_vector)(unsafe.Pointer(g.Ptr()))))
}

func FsolverDriver(s *GslMultifitFsolver, maxiter int, epsabs float64, epsrel float64) int32 {
   return int32(C.gsl_multifit_fsolver_driver((*C.gsl_multifit_fsolver)(unsafe.Pointer(s.Ptr())), C.size_t(maxiter), C.double(epsabs), C.double(epsrel)))
}

func FdfsolverDriver(s *GslMultifitFdfsolver, maxiter int, epsabs float64, epsrel float64) int32 {
   return int32(C.gsl_multifit_fdfsolver_driver((*C.gsl_multifit_fdfsolver)(unsafe.Pointer(s.Ptr())), C.size_t(maxiter), C.double(epsabs), C.double(epsrel)))
}

func Covar(j *matrix.GslMatrix, epsrel float64, covar *matrix.GslMatrix) int32 {
   return int32(C.gsl_multifit_covar((*C.gsl_matrix)(unsafe.Pointer(j.Ptr())), C.double(epsrel), (*C.gsl_matrix)(unsafe.Pointer(covar.Ptr()))))
}

