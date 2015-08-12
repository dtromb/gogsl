package multifit

/*

	#include <stdlib.h>
	#include <gsl/gsl_multifit_nlin.h>
	#include <gsl/gsl_vector.h>
	#include <gsl/gsl_matrix.h>

	int _cgo_gsl_multifit_function_proxy(const gsl_vector * x, void * params, gsl_vector * f) {
		return gslMultifitFunctionCaller(x, params, f);
	}

	int _cgo_gsl_multifit_function_fdf_f_proxy (const gsl_vector * x, void * params, gsl_vector * f) {
		return gslMultifitFunctionFdfFCaller(x, params, f);
	}

	int _cgo_gsl_multifit_function_fdf_df_proxy(const gsl_vector * x, void * params, gsl_matrix * J) {
		return gslMultifitFunctionFdfDfCaller(x, params, J);
	}

	int _cgo_gsl_multifit_function_fdf_proxy(const gsl_vector * x, void * params, gsl_vector * f, gsl_matrix * J) {
		return gslMultifitFunctionFdfCaller(x, params, f, J);
	}

	void _init_proxy_gsl_multifit_function(void *params, size_t dim, gsl_multifit_function *gslFunc) {
		gslFunc->f = _cgo_gsl_multifit_function_proxy;
		gslFunc->n = dim;
		gslFunc->params = params;
	}

	void _init_proxy_gsl_multifit_function_fdf(void *params, size_t dim, gsl_multifit_function_fdf *gslFunc) {
		gslFunc->f = _cgo_gsl_multifit_function_fdf_f_proxy;
		gslFunc->df = _cgo_gsl_multifit_function_fdf_df_proxy;
		gslFunc->fdf = _cgo_gsl_multifit_function_fdf_proxy;
		gslFunc->n = dim;
		gslFunc->params = params;
	}


	size_t get_gsl_multifit_function_struct_size() {
		return sizeof(gsl_multifit_function);
	}

	size_t get_gsl_multifit_function_fdf_struct_size() {
		return sizeof(gsl_multifit_function_fdf);
	}


*/
import "C"

import (
	"reflect"
	"unsafe"
)

var global_GSL_MULTIFIT_FUNCTION_STRUCT_SIZE uint32 = uint32(C.get_gsl_multifit_function_struct_size())
var global_GSL_MULTIFIT_FUNCTION_FDF_STRUCT_SIZE uint32 = uint32(C.get_gsl_multifit_function_fdf_struct_size())

func InitializeGslMultifitFunction(ptr *GslMultifitFunction) {
	ptr.cGslFunctionStruct = make([]byte, 0, global_GSL_MULTIFIT_FUNCTION_STRUCT_SIZE)
	cPtr := ((*reflect.SliceHeader)(unsafe.Pointer(&ptr.cGslFunctionStruct))).Data
	C._init_proxy_gsl_multifit_function(unsafe.Pointer(ptr), C.size_t(ptr.Dimension), (*C.gsl_multifit_function)(unsafe.Pointer(cPtr)))
}

func InitializeGslMultifitFunctionFdf(ptr *GslMultifitFunctionFdf) {
	ptr.cGslFunctionStruct = make([]byte, 0, global_GSL_MULTIFIT_FUNCTION_FDF_STRUCT_SIZE)
	cPtr := ((*reflect.SliceHeader)(unsafe.Pointer(&ptr.cGslFunctionStruct))).Data
	C._init_proxy_gsl_multifit_function_fdf(unsafe.Pointer(ptr), C.size_t(ptr.Dimension), (*C.gsl_multifit_function_fdf)(unsafe.Pointer(cPtr)))
}

func (gf GslMultifitFunction) CPtr() uintptr {
	return ((*reflect.SliceHeader)(unsafe.Pointer(&gf.cGslFunctionStruct))).Data
}

func (gf GslMultifitFunctionFdf) CPtr() uintptr {
	return ((*reflect.SliceHeader)(unsafe.Pointer(&gf.cGslFunctionStruct))).Data
}
