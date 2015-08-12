package multimin

/*

	#include <stdlib.h>
	#include <gsl/gsl_multimin.h>
	#include <gsl/gsl_vector.h>

	double _cgo_gsl_multimin_function_proxy(const gsl_vector * x, void * params) {
		return gslMultiminFunctionCaller(x, params);
	}

	double _cgo_gsl_multimin_function_fdf_f_proxy (const gsl_vector * x, void * params) {
		return gslMultiminFunctionFdfFCaller(x, params);
	}

	void _cgo_gsl_multimin_function_fdf_df_proxy (const gsl_vector * x, void * params, gsl_vector * g) {
		gslMultiminFunctionFdfDfCaller(x, params, g);
	}

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


*/
import "C"

import (
	"reflect"
	"unsafe"
)

var global_GSL_MULTIMIN_FUNCTION_STRUCT_SIZE uint32 = uint32(C.get_gsl_multimin_function_struct_size())
var global_GSL_MULTIMIN_FUNCTION_FDF_STRUCT_SIZE uint32 = uint32(C.get_gsl_multimin_function_fdf_struct_size())

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

func (gf GslMultiminFunction) CPtr() uintptr {
	return ((*reflect.SliceHeader)(unsafe.Pointer(&gf.cGslFunctionStruct))).Data
}

func (gf GslMultiminFunctionFdf) CPtr() uintptr {
	return ((*reflect.SliceHeader)(unsafe.Pointer(&gf.cGslFunctionStruct))).Data
}
