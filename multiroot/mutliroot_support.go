package multiroot

/*

	#include <stdlib.h>
	#include <gsl/gsl_multiroots.h>
	#include <gsl/gsl_vector.h>
	#include <gsl/gsl_matrix.h>
	#include <stdio.h>

    extern int gslMultirootFunctionCaller(const gsl_vector *x, void *params, gsl_vector *f);
	int _cgo_gsl_multiroot_function_proxy(const gsl_vector * x, void * params, gsl_vector * f) {
		return gslMultirootFunctionCaller(x, params, f);
	}

  	extern int gslMultirootFunctionFdfFCaller(const gsl_vector *x, void *params, gsl_vector *f);
	int _cgo_gsl_multiroot_function_fdf_f_proxy (const gsl_vector * x, void * params, gsl_vector * f) {
		return gslMultirootFunctionFdfFCaller(x, params, f);
	}

	extern int gslMultirootFunctionFdfDfCaller(const gsl_vector *x, void *params, gsl_matrix *J);
	int _cgo_gsl_multiroot_function_fdf_df_proxy(const gsl_vector * x, void * params, gsl_matrix * J) {
		return gslMultirootFunctionFdfDfCaller(x, params, J);
	}

	extern int gslMultirootFunctionFdfCaller(const gsl_vector *x, void *params, gsl_vector *f, gsl_matrix *J);
	int _cgo_gsl_multiroot_function_fdf_proxy(const gsl_vector * x, void * params, gsl_vector * f, gsl_matrix * J) {
		return gslMultirootFunctionFdfCaller(x, params, f, J);
	}

	void _init_proxy_gsl_multiroot_function(void *params, size_t dim, gsl_multiroot_function *gslFunc) {
		gslFunc->f = _cgo_gsl_multiroot_function_proxy;
		gslFunc->n = dim;
		gslFunc->params = params;
	}

	void _init_proxy_gsl_multiroot_function_fdf(void *params, size_t dim, gsl_multiroot_function_fdf *gslFunc) {
		gslFunc->f = _cgo_gsl_multiroot_function_fdf_f_proxy;
		gslFunc->df = _cgo_gsl_multiroot_function_fdf_df_proxy;
		gslFunc->fdf = _cgo_gsl_multiroot_function_fdf_proxy;
		gslFunc->n = dim;
		gslFunc->params = params;
	}


	size_t get_gsl_multiroot_function_struct_size() {
		return sizeof(gsl_multiroot_function);
	}

	size_t get_gsl_multiroot_function_fdf_struct_size() {
		return sizeof(gsl_multiroot_function_fdf);
	}

	gsl_vector *get_gsl_mutliroot_solver_f_ptr(gsl_multiroot_fsolver *s) {
		return s->f;
	}

	gsl_vector *get_gsl_mutliroot_solver_x_ptr(gsl_multiroot_fsolver *s) {
		return s->x;
	}

	gsl_vector *get_gsl_mutliroot_solver_dx_ptr(gsl_multiroot_fsolver *s) {
		return s->dx;
	}

	gsl_vector *get_gsl_mutliroot_fdfsolver_f_ptr(gsl_multiroot_fdfsolver *s) {
		return s->f;
	}

	gsl_vector *get_gsl_mutliroot_fdfsolver_x_ptr(gsl_multiroot_fdfsolver *s) {
		return s->x;
	}

	gsl_vector *get_gsl_mutliroot_fdfsolver_dx_ptr(gsl_multiroot_fdfsolver *s) {
		return s->dx;
	}
*/
import "C"

import (
	"github.com/dtromb/gogsl"
	"github.com/dtromb/gogsl/vector"
	"reflect"
	"unsafe"
)

var global_GSL_MULTIROOT_FUNCTION_STRUCT_SIZE uint32 = uint32(C.get_gsl_multiroot_function_struct_size())
var global_GSL_MULTIROOT_FUNCTION_FDF_STRUCT_SIZE uint32 = uint32(C.get_gsl_multiroot_function_fdf_struct_size())

var GSL_MULTIROOT_FSOLVER_HYBRIDS *GslMultirootFsolverType = &GslMultirootFsolverType{cPtr: uintptr(unsafe.Pointer(C.gsl_multiroot_fsolver_hybrids))}
var GSL_MULTIROOT_FSOLVER_HYBRID *GslMultirootFsolverType = &GslMultirootFsolverType{cPtr: uintptr(unsafe.Pointer(C.gsl_multiroot_fsolver_hybrid))}
var GSL_MULTIROOT_FSOLVER_DNEWTON *GslMultirootFsolverType = &GslMultirootFsolverType{cPtr: uintptr(unsafe.Pointer(C.gsl_multiroot_fsolver_dnewton))}
var GSL_MULTIROOT_FSOLVER_BROYDEN *GslMultirootFsolverType = &GslMultirootFsolverType{cPtr: uintptr(unsafe.Pointer(C.gsl_multiroot_fsolver_broyden))}

var GSL_MULTIROOT_FDFSOLVER_HYBRIDSJ *GslMultirootFdfsolverType = &GslMultirootFdfsolverType{cPtr: uintptr(unsafe.Pointer(C.gsl_multiroot_fdfsolver_hybridsj))}
var GSL_MULTIROOT_FDFSOLVER_HYBRIDJ *GslMultirootFdfsolverType = &GslMultirootFdfsolverType{cPtr: uintptr(unsafe.Pointer(C.gsl_multiroot_fdfsolver_hybridj))}
var GSL_MULTIROOT_FDFSOLVER_NEWTON *GslMultirootFdfsolverType = &GslMultirootFdfsolverType{cPtr: uintptr(unsafe.Pointer(C.gsl_multiroot_fdfsolver_newton))}
var GSL_MULTIROOT_FDFSOLVER_GNEWTON *GslMultirootFdfsolverType = &GslMultirootFdfsolverType{cPtr: uintptr(unsafe.Pointer(C.gsl_multiroot_fdfsolver_gnewton))}

func (s *GslMultirootFsolver) X_() *vector.GslVector {
	cptr := C.get_gsl_mutliroot_solver_x_ptr((*C.gsl_multiroot_fsolver)(unsafe.Pointer(s.Ptr())))
	v := &vector.GslVector{}
	gogsl.InitializeGslStruct(v, uintptr(unsafe.Pointer(cptr)))
	return v
}

func (s *GslMultirootFsolver) F_() *vector.GslVector {
	cptr := C.get_gsl_mutliroot_solver_f_ptr((*C.gsl_multiroot_fsolver)(unsafe.Pointer(s.Ptr())))
	v := &vector.GslVector{}
	gogsl.InitializeGslStruct(v, uintptr(unsafe.Pointer(cptr)))
	return v
}

func (s *GslMultirootFsolver) Dx_() *vector.GslVector {
	cptr := C.get_gsl_mutliroot_solver_dx_ptr((*C.gsl_multiroot_fsolver)(unsafe.Pointer(s.Ptr())))
	v := &vector.GslVector{}
	gogsl.InitializeGslStruct(v, uintptr(unsafe.Pointer(cptr)))
	return v
}

func (s *GslMultirootFdfsolver) X_() *vector.GslVector {
	cptr := C.get_gsl_mutliroot_fdfsolver_x_ptr((*C.gsl_multiroot_fdfsolver)(unsafe.Pointer(s.Ptr())))
	v := &vector.GslVector{}
	gogsl.InitializeGslStruct(v, uintptr(unsafe.Pointer(cptr)))
	return v
}

func (s *GslMultirootFdfsolver) F_() *vector.GslVector {
	cptr := C.get_gsl_mutliroot_fdfsolver_f_ptr((*C.gsl_multiroot_fdfsolver)(unsafe.Pointer(s.Ptr())))
	v := &vector.GslVector{}
	gogsl.InitializeGslStruct(v, uintptr(unsafe.Pointer(cptr)))
	return v
}

func (s *GslMultirootFdfsolver) Dx_() *vector.GslVector {
	cptr := C.get_gsl_mutliroot_fdfsolver_dx_ptr((*C.gsl_multiroot_fdfsolver)(unsafe.Pointer(s.Ptr())))
	v := &vector.GslVector{}
	gogsl.InitializeGslStruct(v, uintptr(unsafe.Pointer(cptr)))
	return v
}

func InitializeGslMultirootFunction(ptr *GslMultirootFunction) {
	ptr.cGslFunctionStruct = make([]byte, 0, global_GSL_MULTIROOT_FUNCTION_STRUCT_SIZE)
	cPtr := ((*reflect.SliceHeader)(unsafe.Pointer(&ptr.cGslFunctionStruct))).Data
	C._init_proxy_gsl_multiroot_function(unsafe.Pointer(ptr), C.size_t(ptr.Dimension), (*C.gsl_multiroot_function)(unsafe.Pointer(cPtr)))
}

func InitializeGslMultirootFunctionFdf(ptr *GslMultirootFunctionFdf) {
	ptr.cGslFunctionStruct = make([]byte, 0, global_GSL_MULTIROOT_FUNCTION_FDF_STRUCT_SIZE)
	cPtr := ((*reflect.SliceHeader)(unsafe.Pointer(&ptr.cGslFunctionStruct))).Data
	C._init_proxy_gsl_multiroot_function_fdf(unsafe.Pointer(ptr), C.size_t(ptr.Dimension), (*C.gsl_multiroot_function_fdf)(unsafe.Pointer(cPtr)))
}

func (gf GslMultirootFunction) CPtr() uintptr {
	return ((*reflect.SliceHeader)(unsafe.Pointer(&gf.cGslFunctionStruct))).Data
}

func (gf GslMultirootFunctionFdf) CPtr() uintptr {
	return ((*reflect.SliceHeader)(unsafe.Pointer(&gf.cGslFunctionStruct))).Data
}

func (mft *GslMultirootFsolverType) Ptr() uintptr {
	return mft.cPtr
}

func (mft *GslMultirootFsolverType) Dispose() {
	mft.cPtr = 0
}

func (mft *GslMultirootFdfsolverType) Ptr() uintptr {
	return mft.cPtr
}

func (mft *GslMultirootFdfsolverType) Dispose() {
	mft.cPtr = 0
}
