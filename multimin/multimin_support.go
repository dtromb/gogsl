package multimin

/*

	#include <stdlib.h>
	#include <gsl/gsl_multimin.h>
	#include <gsl/gsl_vector.h>

	extern double gslMultiminFunctionCaller(const gsl_vector*,void*);
	double _cgo_gsl_multimin_function_proxy(const gsl_vector * x, void * params) {
		return gslMultiminFunctionCaller(x, params);
	}

	extern double gslMultiminFunctionFdfFCaller(const gsl_vector*,void*);
	double _cgo_gsl_multimin_function_fdf_f_proxy (const gsl_vector * x, void * params) {
		return gslMultiminFunctionFdfFCaller(x, params);
	}

	extern void gslMultiminFunctionFdfDfCaller(const gsl_vector*,void*,gsl_vector*);
	void _cgo_gsl_multimin_function_fdf_df_proxy (const gsl_vector * x, void * params, gsl_vector * g) {
		gslMultiminFunctionFdfDfCaller(x, params, g);
	}

	extern double gslMultiminFunctionFdfCaller(const gsl_vector*,void*,gsl_vector*);
	void _cgo_gsl_multimin_function_fdf_proxy (const gsl_vector * x, void * params, double * f, gsl_vector * g) {
		double res;
		res = gslMultiminFunctionFdfCaller(x, params, g);
		if (f) *f = res;
	}

	void _init_proxy_gsl_multimin_function(void *params, size_t dim, gsl_multimin_function *gslFunc) {
		gslFunc->f = _cgo_gsl_multimin_function_proxy;
		gslFunc->n = dim;
		gslFunc->params = params;
	}

	void _init_proxy_gsl_multimin_function_fdf(void *params, size_t dim, gsl_multimin_function_fdf *gslFunc) {
		gslFunc->f = _cgo_gsl_multimin_function_fdf_f_proxy;
		gslFunc->df = _cgo_gsl_multimin_function_fdf_df_proxy;
		gslFunc->fdf = _cgo_gsl_multimin_function_fdf_proxy;
		gslFunc->n = dim;
		gslFunc->params = params;
	}

	size_t get_gsl_multimin_function_struct_size() {
		return sizeof(gsl_multimin_function);
	}

	size_t get_gsl_multimin_function_fdf_struct_size() {
		return sizeof(gsl_multimin_function_fdf);
	}

	double get_gsl_mutlimin_fminimizer_fval(gsl_multimin_fminimizer *s) {
		return s->fval;
	}

	gsl_vector *get_gsl_mutlimin_fminimizer_x_ptr(gsl_multimin_fminimizer *s) {
		return s->x;
	}

	double get_gsl_mutlimin_fminimizer_size(gsl_multimin_fminimizer *s) {
		return s->size;
	}

	double get_gsl_mutlimin_fdfminimizer_fval(gsl_multimin_fdfminimizer *s) {
		return s->f;
	}

	gsl_vector *get_gsl_mutlimin_fdfminimizer_x_ptr(gsl_multimin_fdfminimizer *s) {
		return s->x;
	}

	gsl_vector *get_gsl_mutlimin_fdfminimizer_dx_ptr(gsl_multimin_fdfminimizer *s) {
		return s->dx;
	}

	gsl_vector *get_gsl_mutlimin_fdfminimizer_gradient_ptr(gsl_multimin_fdfminimizer *s) {
		return s->gradient;
	}


*/
import "C"

import (
	"github.com/dtromb/gogsl"
	"github.com/dtromb/gogsl/vector"
	"reflect"
	"unsafe"
)

var global_GSL_MULTIMIN_FUNCTION_STRUCT_SIZE uint32 = uint32(C.get_gsl_multimin_function_struct_size())
var global_GSL_MULTIMIN_FUNCTION_FDF_STRUCT_SIZE uint32 = uint32(C.get_gsl_multimin_function_fdf_struct_size())

var GSL_MULTIMIN_FMINIMIZER_NSIMPLEX *GslMultiminFminimizerType = &GslMultiminFminimizerType{cPtr: uintptr(unsafe.Pointer(C.gsl_multimin_fminimizer_nmsimplex))}
var GSL_MULTIMIN_FMINIMIZER_NSIMPLEX2 *GslMultiminFminimizerType = &GslMultiminFminimizerType{cPtr: uintptr(unsafe.Pointer(C.gsl_multimin_fminimizer_nmsimplex2))}
var GSL_MULTIMIN_FMINIMIZER_NSIMPLEX2RAND *GslMultiminFminimizerType = &GslMultiminFminimizerType{cPtr: uintptr(unsafe.Pointer(C.gsl_multimin_fminimizer_nmsimplex2rand))}

var GSL_MULTIMIN_FDFMINIMIZER_CONJUGATE_FR *GslMultiminFdfminimizerType = &GslMultiminFdfminimizerType{cPtr: uintptr(unsafe.Pointer(C.gsl_multimin_fdfminimizer_conjugate_fr))}
var GSL_MULTIMIN_FDFMINIMIZER_CONJUGATE_PR *GslMultiminFdfminimizerType = &GslMultiminFdfminimizerType{cPtr: uintptr(unsafe.Pointer(C.gsl_multimin_fdfminimizer_conjugate_pr))}
var GSL_MULTIMIN_FDFMINIMIZER_VECTOR_BFGS2 *GslMultiminFdfminimizerType = &GslMultiminFdfminimizerType{cPtr: uintptr(unsafe.Pointer(C.gsl_multimin_fdfminimizer_vector_bfgs2))}
var GSL_MULTIMIN_FDFMINIMIZER_VECTOR_BFGS *GslMultiminFdfminimizerType = &GslMultiminFdfminimizerType{cPtr: uintptr(unsafe.Pointer(C.gsl_multimin_fdfminimizer_vector_bfgs))}
var GSL_MULTIMIN_FDFMINIMIZER_STEEPEST_DESCENT *GslMultiminFdfminimizerType = &GslMultiminFdfminimizerType{cPtr: uintptr(unsafe.Pointer(C.gsl_multimin_fdfminimizer_steepest_descent))}

func InitializeGslMultiminFunction(ptr *GslMultiminFunction) {
	ptr.cGslFunctionStruct = make([]byte, 0, global_GSL_MULTIMIN_FUNCTION_STRUCT_SIZE)
	cPtr := ((*reflect.SliceHeader)(unsafe.Pointer(&ptr.cGslFunctionStruct))).Data
	C._init_proxy_gsl_multimin_function(unsafe.Pointer(ptr), C.size_t(ptr.Dimension), (*C.gsl_multimin_function)(unsafe.Pointer(cPtr)))
}

func InitializeGslMultiminFunctionFdf(ptr *GslMultiminFunctionFdf) {
	ptr.cGslFunctionStruct = make([]byte, 0, global_GSL_MULTIMIN_FUNCTION_FDF_STRUCT_SIZE)
	cPtr := ((*reflect.SliceHeader)(unsafe.Pointer(&ptr.cGslFunctionStruct))).Data
	C._init_proxy_gsl_multimin_function_fdf(unsafe.Pointer(ptr), C.size_t(ptr.Dimension), (*C.gsl_multimin_function_fdf)(unsafe.Pointer(cPtr)))
}

func (s *GslMultiminFminimizer) Fval() float64 {
	return float64(C.get_gsl_mutlimin_fminimizer_fval((*C.gsl_multimin_fminimizer)(unsafe.Pointer(s.Ptr()))))
}

func (s *GslMultiminFminimizer) Size() float64 {
	return float64(C.get_gsl_mutlimin_fminimizer_size((*C.gsl_multimin_fminimizer)(unsafe.Pointer(s.Ptr()))))
}

func (s *GslMultiminFminimizer) X_() *vector.GslVector {
	cptr := C.get_gsl_mutlimin_fminimizer_x_ptr((*C.gsl_multimin_fminimizer)(unsafe.Pointer(s.Ptr())))
	v := &vector.GslVector{}
	gogsl.InitializeGslStruct(v, uintptr(unsafe.Pointer(cptr)))
	return v
}

func (s *GslMultiminFdfminimizer) Fval() float64 {
	return float64(C.get_gsl_mutlimin_fdfminimizer_fval((*C.gsl_multimin_fdfminimizer)(unsafe.Pointer(s.Ptr()))))
}

func (s *GslMultiminFdfminimizer) X_() *vector.GslVector {
	cptr := C.get_gsl_mutlimin_fdfminimizer_x_ptr((*C.gsl_multimin_fdfminimizer)(unsafe.Pointer(s.Ptr())))
	v := &vector.GslVector{}
	gogsl.InitializeGslStruct(v, uintptr(unsafe.Pointer(cptr)))
	return v
}

func (s *GslMultiminFdfminimizer) Dx_() *vector.GslVector {
	cptr := C.get_gsl_mutlimin_fdfminimizer_dx_ptr((*C.gsl_multimin_fdfminimizer)(unsafe.Pointer(s.Ptr())))
	v := &vector.GslVector{}
	gogsl.InitializeGslStruct(v, uintptr(unsafe.Pointer(cptr)))
	return v
}

func (s *GslMultiminFdfminimizer) Gradient_() *vector.GslVector {
	cptr := C.get_gsl_mutlimin_fdfminimizer_gradient_ptr((*C.gsl_multimin_fdfminimizer)(unsafe.Pointer(s.Ptr())))
	v := &vector.GslVector{}
	gogsl.InitializeGslStruct(v, uintptr(unsafe.Pointer(cptr)))
	return v
}

func (gf GslMultiminFunction) CPtr() uintptr {
	return ((*reflect.SliceHeader)(unsafe.Pointer(&gf.cGslFunctionStruct))).Data
}

func (gf GslMultiminFunctionFdf) CPtr() uintptr {
	return ((*reflect.SliceHeader)(unsafe.Pointer(&gf.cGslFunctionStruct))).Data
}

func (mt *GslMultiminFminimizerType) Ptr() uintptr {
	return mt.cPtr
}

func (mt *GslMultiminFminimizerType) Dispose() {
	mt.cPtr = 0
}

func (mt *GslMultiminFdfminimizerType) Ptr() uintptr {
	return mt.cPtr
}

func (mt *GslMultiminFdfminimizerType) Dispose() {
	mt.cPtr = 0
}
