//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package root

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_roots.h>
*/
import "C"

import "unsafe"
import "github.com/dtromb/gogsl"

type GslRootFsolver struct {
   gogsl.GslReference
}

type GslRootFsolverType struct {
   gogsl.GslReference
   cPtr uintptr
}

type GslRootFdfsolver struct {
   gogsl.GslReference
}

type GslRootFdfsolverType struct {
   gogsl.GslReference
   cPtr uintptr
}

func FsolverAlloc(t *GslRootFsolverType) *GslRootFsolver {
   _ref := C.gsl_root_fsolver_alloc((*C.gsl_root_fsolver_type)(unsafe.Pointer(t.Ptr())))
   _result := &GslRootFsolver{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func FdfsolverAlloc(t *GslRootFdfsolverType) *GslRootFdfsolver {
   _ref := C.gsl_root_fdfsolver_alloc((*C.gsl_root_fdfsolver_type)(unsafe.Pointer(t.Ptr())))
   _result := &GslRootFdfsolver{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func FsolverSet(s *GslRootFsolver, f *gogsl.GslFunction, xLower float64, xUpper float64) int32 {
   gogsl.InitializeGslFunction(f)
   return int32(C.gsl_root_fsolver_set((*C.gsl_root_fsolver)(unsafe.Pointer(s.Ptr())), (*C.gsl_function)(unsafe.Pointer(f.CPtr())), C.double(xLower), C.double(xUpper)))
}

func FdfsolverSet(s *GslRootFdfsolver, fdf *gogsl.GslFunctionFdf, root float64) int32 {
   gogsl.InitializeGslFunctionFdf(fdf)
   return int32(C.gsl_root_fdfsolver_set((*C.gsl_root_fdfsolver)(unsafe.Pointer(s.Ptr())), (*C.gsl_function_fdf)(unsafe.Pointer(fdf.CPtr())), C.double(root)))
}

func (x *GslRootFsolver) Dispose() {
   C.gsl_root_fsolver_free((*C.gsl_root_fsolver)(unsafe.Pointer(x.Ptr())))
}

func (x *GslRootFdfsolver) Dispose() {
   C.gsl_root_fdfsolver_free((*C.gsl_root_fdfsolver)(unsafe.Pointer(x.Ptr())))
}

func FsolverName(s *GslRootFsolver) string {
   return C.GoString(C.gsl_root_fsolver_name((*C.gsl_root_fsolver)(unsafe.Pointer(s.Ptr()))))
}

func FdfsolverName(s *GslRootFdfsolver) string {
   return C.GoString(C.gsl_root_fdfsolver_name((*C.gsl_root_fdfsolver)(unsafe.Pointer(s.Ptr()))))
}

func FsolverIterate(s *GslRootFsolver) int32 {
   return int32(C.gsl_root_fsolver_iterate((*C.gsl_root_fsolver)(unsafe.Pointer(s.Ptr()))))
}

func FdfsolverIterate(s *GslRootFdfsolver) int32 {
   return int32(C.gsl_root_fdfsolver_iterate((*C.gsl_root_fdfsolver)(unsafe.Pointer(s.Ptr()))))
}

func FsolverRoot(s *GslRootFsolver) float64 {
   return float64(C.gsl_root_fsolver_root((*C.gsl_root_fsolver)(unsafe.Pointer(s.Ptr()))))
}

func FdfsolverRoot(s *GslRootFdfsolver) float64 {
   return float64(C.gsl_root_fdfsolver_root((*C.gsl_root_fdfsolver)(unsafe.Pointer(s.Ptr()))))
}

func FsolverXLower(s *GslRootFsolver) float64 {
   return float64(C.gsl_root_fsolver_x_lower((*C.gsl_root_fsolver)(unsafe.Pointer(s.Ptr()))))
}

func FsolverXUpper(s *GslRootFsolver) float64 {
   return float64(C.gsl_root_fsolver_x_upper((*C.gsl_root_fsolver)(unsafe.Pointer(s.Ptr()))))
}

func TestInterval(xLower float64, xUpper float64, epsabs float64, epsrel float64) int32 {
   return int32(C.gsl_root_test_interval(C.double(xLower), C.double(xUpper), C.double(epsabs), C.double(epsrel)))
}

func TestDelta(x1 float64, x0 float64, epsabs float64, epsrel float64) int32 {
   return int32(C.gsl_root_test_delta(C.double(x1), C.double(x0), C.double(epsabs), C.double(epsrel)))
}

func TestResidual(f float64, epsabs float64) int32 {
   return int32(C.gsl_root_test_residual(C.double(f), C.double(epsabs)))
}

