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

type GslMultifitLinearWorkspace struct {
	gogsl.GslReference
}

type GslMultifitRobustWorkspace struct {
	gogsl.GslReference
}

type GslMultifitRobustType struct {
	gogsl.GslReference
	cPtr uintptr
}

type GslMultifitFsolver struct {
	gogsl.GslReference
}

type GslMultifitFsolverType struct {
	gogsl.GslReference
	cPtr uintptr
}

type GslMultifitFdfsolver struct {
	gogsl.GslReference
}

type GslMultifitFdfsolverType struct {
	gogsl.GslReference
	cPtr uintptr
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

func LinearAlloc(n int, p int) *GslMultifitLinearWorkspace {
	_ref := C.gsl_multifit_linear_alloc(C.size_t(n), C.size_t(p))
	_result := &GslMultifitLinearWorkspace{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslMultifitLinearWorkspace) Dispose() {
	C.gsl_multifit_linear_free((*C.gsl_multifit_linear_workspace)(unsafe.Pointer(x.Ptr())))
}

func Linear(x *matrix.GslMatrix, y *vector.GslVector, c *vector.GslVector, cov *matrix.GslMatrix, work *GslMultifitLinearWorkspace) (int32, float64) {
	var _outptr_4 C.double
	_result := int32(C.gsl_multifit_linear((*C.gsl_matrix)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector)(unsafe.Pointer(y.Ptr())), (*C.gsl_vector)(unsafe.Pointer(c.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(cov.Ptr())), &_outptr_4, (*C.gsl_multifit_linear_workspace)(unsafe.Pointer(work.Ptr()))))
	return _result, *(*float64)(unsafe.Pointer(&_outptr_4))
}

func Wlinear(x *matrix.GslMatrix, w *vector.GslVector, y *vector.GslVector, c *vector.GslVector, cov *matrix.GslMatrix, work *GslMultifitLinearWorkspace) (int32, float64) {
	var _outptr_5 C.double
	_result := int32(C.gsl_multifit_wlinear((*C.gsl_matrix)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector)(unsafe.Pointer(w.Ptr())), (*C.gsl_vector)(unsafe.Pointer(y.Ptr())), (*C.gsl_vector)(unsafe.Pointer(c.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(cov.Ptr())), &_outptr_5, (*C.gsl_multifit_linear_workspace)(unsafe.Pointer(work.Ptr()))))
	return _result, *(*float64)(unsafe.Pointer(&_outptr_5))
}

func LinearSvd(x *matrix.GslMatrix, y *vector.GslVector, tol float64, c *vector.GslVector, cov *matrix.GslMatrix, work *GslMultifitLinearWorkspace) (int32, int, float64) {
	var _outptr_3 C.size_t
	var _outptr_6 C.double
	_result := int32(C.gsl_multifit_linear_svd((*C.gsl_matrix)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector)(unsafe.Pointer(y.Ptr())), C.double(tol), &_outptr_3, (*C.gsl_vector)(unsafe.Pointer(c.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(cov.Ptr())), &_outptr_6, (*C.gsl_multifit_linear_workspace)(unsafe.Pointer(work.Ptr()))))
	return _result, *(*int)(unsafe.Pointer(&_outptr_3)), *(*float64)(unsafe.Pointer(&_outptr_6))
}

func WlinearSvd(x *matrix.GslMatrix, w *vector.GslVector, y *vector.GslVector, tol float64, c *vector.GslVector, cov *matrix.GslMatrix, work *GslMultifitLinearWorkspace) (int32, int, float64) {
	var _outptr_4 C.size_t
	var _outptr_7 C.double
	_result := int32(C.gsl_multifit_wlinear_svd((*C.gsl_matrix)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector)(unsafe.Pointer(w.Ptr())), (*C.gsl_vector)(unsafe.Pointer(y.Ptr())), C.double(tol), &_outptr_4, (*C.gsl_vector)(unsafe.Pointer(c.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(cov.Ptr())), &_outptr_7, (*C.gsl_multifit_linear_workspace)(unsafe.Pointer(work.Ptr()))))
	return _result, *(*int)(unsafe.Pointer(&_outptr_4)), *(*float64)(unsafe.Pointer(&_outptr_7))
}

func LinearUsvd(x *matrix.GslMatrix, y *vector.GslVector, tol float64, c *vector.GslVector, cov *matrix.GslMatrix, work *GslMultifitLinearWorkspace) (int32, int, float64) {
	var _outptr_3 C.size_t
	var _outptr_6 C.double
	_result := int32(C.gsl_multifit_linear_usvd((*C.gsl_matrix)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector)(unsafe.Pointer(y.Ptr())), C.double(tol), &_outptr_3, (*C.gsl_vector)(unsafe.Pointer(c.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(cov.Ptr())), &_outptr_6, (*C.gsl_multifit_linear_workspace)(unsafe.Pointer(work.Ptr()))))
	return _result, *(*int)(unsafe.Pointer(&_outptr_3)), *(*float64)(unsafe.Pointer(&_outptr_6))
}

func WlinearUsvd(x *matrix.GslMatrix, w *vector.GslVector, y *vector.GslVector, tol float64, c *vector.GslVector, cov *matrix.GslMatrix, work *GslMultifitLinearWorkspace) (int32, int, float64) {
	var _outptr_4 C.size_t
	var _outptr_7 C.double
	_result := int32(C.gsl_multifit_wlinear_usvd((*C.gsl_matrix)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector)(unsafe.Pointer(w.Ptr())), (*C.gsl_vector)(unsafe.Pointer(y.Ptr())), C.double(tol), &_outptr_4, (*C.gsl_vector)(unsafe.Pointer(c.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(cov.Ptr())), &_outptr_7, (*C.gsl_multifit_linear_workspace)(unsafe.Pointer(work.Ptr()))))
	return _result, *(*int)(unsafe.Pointer(&_outptr_4)), *(*float64)(unsafe.Pointer(&_outptr_7))
}

func LinearEst(x *vector.GslVector, c *vector.GslVector, cov *matrix.GslMatrix) (int32, float64, float64) {
	var _outptr_3 C.double
	var _outptr_4 C.double
	_result := int32(C.gsl_multifit_linear_est((*C.gsl_vector)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector)(unsafe.Pointer(c.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(cov.Ptr())), &_outptr_3, &_outptr_4))
	return _result, *(*float64)(unsafe.Pointer(&_outptr_3)), *(*float64)(unsafe.Pointer(&_outptr_4))
}

func LinearResiduals(x *matrix.GslMatrix, y *vector.GslVector, c *vector.GslVector, r *vector.GslVector) int32 {
	return int32(C.gsl_multifit_linear_residuals((*C.gsl_matrix)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector)(unsafe.Pointer(y.Ptr())), (*C.gsl_vector)(unsafe.Pointer(c.Ptr())), (*C.gsl_vector)(unsafe.Pointer(r.Ptr()))))
}

func RobustAlloc(t *GslMultifitRobustType, n int, p int) *GslMultifitRobustWorkspace {
	_ref := C.gsl_multifit_robust_alloc((*C.gsl_multifit_robust_type)(unsafe.Pointer(t.Ptr())), C.size_t(n), C.size_t(p))
	_result := &GslMultifitRobustWorkspace{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslMultifitRobustWorkspace) Dispose() {
	C.gsl_multifit_robust_free((*C.gsl_multifit_robust_workspace)(unsafe.Pointer(x.Ptr())))
}

func RobustName(w *GslMultifitRobustWorkspace) string {
	return C.GoString(C.gsl_multifit_robust_name((*C.gsl_multifit_robust_workspace)(unsafe.Pointer(w.Ptr()))))
}

func RobustTune(tune float64, w *GslMultifitRobustWorkspace) int32 {
	return int32(C.gsl_multifit_robust_tune(C.double(tune), (*C.gsl_multifit_robust_workspace)(unsafe.Pointer(w.Ptr()))))
}

func Robust(x *matrix.GslMatrix, y *vector.GslVector, c *vector.GslVector, cov *matrix.GslMatrix, w *GslMultifitRobustWorkspace) int32 {
	return int32(C.gsl_multifit_robust((*C.gsl_matrix)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector)(unsafe.Pointer(y.Ptr())), (*C.gsl_vector)(unsafe.Pointer(c.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(cov.Ptr())), (*C.gsl_multifit_robust_workspace)(unsafe.Pointer(w.Ptr()))))
}

func RobustEst(x *vector.GslVector, c *vector.GslVector, cov *matrix.GslMatrix) (int32, float64, float64) {
	var _outptr_3 C.double
	var _outptr_4 C.double
	_result := int32(C.gsl_multifit_robust_est((*C.gsl_vector)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector)(unsafe.Pointer(c.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(cov.Ptr())), &_outptr_3, &_outptr_4))
	return _result, *(*float64)(unsafe.Pointer(&_outptr_3)), *(*float64)(unsafe.Pointer(&_outptr_4))
}
