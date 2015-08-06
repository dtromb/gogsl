package multiroot

/*

	#include <stdlib.h>
	#include <gsl/gsl_multiroots.h>
	#include <gsl/gsl_vector.h>
	#include <gsl/gsl_matrix.h>
	
	int _cgo_gsl_multiroot_function_proxy(const gsl_vector * x, void * params, gsl_vector * f) {
		return gslMultirootFunctionCaller(x, params, f);
	}	
	
	int _cgo_gsl_multiroot_function_fdf_f_proxy (const gsl_vector * x, void * params, gsl_vector * f) {
		return gslMultirootFunctionFdfFCaller(x, params, f);
	}	
	
	int _cgo_gsl_multiroot_function_fdf_df_proxy(const gsl_vector * x, void * params, gsl_matrix * J) {
		return gslMultirootFunctionFdfDfCaller(x, params, J);
	}	
	
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
	
	
*/
import "C"

import (
	"reflect"
	"unsafe"
)

var global_GSL_MULTIROOT_FUNCTION_STRUCT_SIZE 		uint32 = uint32(C.get_gsl_multiroot_function_struct_size())
var global_GSL_MULTIROOT_FUNCTION_FDF_STRUCT_SIZE 	uint32 = uint32(C.get_gsl_multiroot_function_fdf_struct_size())

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

