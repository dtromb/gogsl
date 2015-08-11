package gogsl

/*

	#include <stdlib.h>
	#include <gsl/gsl_math.h>
	#include <gsl/gsl_monte.h>
	
	extern double gslFunctionCaller(double, void*);

	double _cgo_gsl_function_proxy(double x, void *params) {
		return gslFunctionCaller(x, params);
	}	
	
	double _cgo_gsl_function_fdf_f_proxy(double x, void *params) {
		return gslFunctionFdfFCaller(x, params);
	}	
	
	double _cgo_gsl_function_fdf_df_proxy(double x, void *params) {
		return gslFunctionFdfDfCaller(x, params);
	}	
	
	typedef struct gslFunctionFdfCaller_return {
		double f;
		double df;
	} gslFunctionFdfCaller_return;
	
	extern gslFunctionFdfCaller_return gslFunctionFdfCaller(double x, void *params);
	
	void _cgo_gsl_function_fdf_proxy(double x, void *params, double *f, double *df) {
		gslFunctionFdfCaller_return foo;
		foo = gslFunctionFdfCaller(x, params);
		if (f) *f = foo.f;
		if (df) *df = foo.df;
	}
	
	double _cgo_gsl_monte_function_proxy(double *x, size_t dim, void * params) {
		return gslMonteFunctionCaller(x, params);
	}
	
	void _init_proxy_gsl_function(void *params, gsl_function *gslFunc) {
		gslFunc->function = _cgo_gsl_function_proxy;
		gslFunc->params = params;
	}
		
	void _init_proxy_gsl_function_fdf(void *params, gsl_function_fdf *gslFunc) {
		gslFunc->f = _cgo_gsl_function_fdf_f_proxy;
		gslFunc->df = _cgo_gsl_function_fdf_df_proxy;
		gslFunc->fdf = _cgo_gsl_function_fdf_proxy;
		gslFunc->params = params;
	}
	
	void _init_proxy_gsl_monte_function(void *params, size_t dim, gsl_monte_function *gslMonteFunc) {
		gslMonteFunc->f = _cgo_gsl_monte_function_proxy;
		gslMonteFunc->dim = dim;
		gslMonteFunc->params = params;
	}
	
	size_t get_gsl_function_struct_size() {
		return sizeof(gsl_function);
	}
	
	size_t get_gsl_function_fdf_struct_size() {
		return sizeof(gsl_function_fdf);
	}
	
	size_t get_gsl_monte_function_struct_size() {
		return sizeof(gsl_monte_function);
	}
	
*/
import "C"

import (
	"reflect"
	"unsafe"
)

var global_GSL_FUNCTION_STRUCT_SIZE 		uint32 = uint32(C.get_gsl_function_struct_size())
var global_GSL_FUNCTION_FDF_STRUCT_SIZE 	uint32 = uint32(C.get_gsl_function_fdf_struct_size())
var global_GSL_MONTE_FUNCTION_STRUCT_SIZE 	uint32 = uint32(C.get_gsl_monte_function_struct_size())

// These structures are allocated in the usual fashion and initialized from desired call parameters.
func InitializeGslFunction(ptr *GslFunction) {
	ptr.cGslFunctionStruct = make([]byte, 0, global_GSL_FUNCTION_STRUCT_SIZE)
	cPtr := ((*reflect.SliceHeader)(unsafe.Pointer(&ptr.cGslFunctionStruct))).Data
	C._init_proxy_gsl_function(unsafe.Pointer(ptr), (*C.gsl_function)(unsafe.Pointer(cPtr)))
}

func InitializeGslFunctionFdf(ptr *GslFunctionFdf) {
	ptr.cGslFunctionStruct = make([]byte, 0, global_GSL_FUNCTION_FDF_STRUCT_SIZE)
	cPtr := ((*reflect.SliceHeader)(unsafe.Pointer(&ptr.cGslFunctionStruct))).Data
	C._init_proxy_gsl_function_fdf(unsafe.Pointer(ptr), (*C.gsl_function_fdf)(unsafe.Pointer(cPtr)))
}

func InitializeGslMonteFunction(ptr *GslMonteFunction) {
	ptr.cGslFunctionStruct = make([]byte, 0, global_GSL_MONTE_FUNCTION_STRUCT_SIZE)
	cPtr := ((*reflect.SliceHeader)(unsafe.Pointer(&ptr.cGslFunctionStruct))).Data
	C._init_proxy_gsl_monte_function(unsafe.Pointer(ptr), C.size_t(ptr.Dim), (*C.gsl_monte_function)(unsafe.Pointer(cPtr)))
}

// These pointers can be cast to the package-local types (eg. *C.gsl_function) to perform the call.
func (gf GslFunction) CPtr() uintptr {
	return ((*reflect.SliceHeader)(unsafe.Pointer(&gf.cGslFunctionStruct))).Data
}

func (gf GslFunctionFdf) CPtr() uintptr {
	return ((*reflect.SliceHeader)(unsafe.Pointer(&gf.cGslFunctionStruct))).Data
}

func (gf GslMonteFunction) CPtr() uintptr {
	return ((*reflect.SliceHeader)(unsafe.Pointer(&gf.cGslFunctionStruct))).Data
}

