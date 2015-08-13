package monte

/*
	#include <unistd.h>
	#include <gsl/gsl_monte.h>

	extern double gslMonteFunctionCaller(void *x, size_t xlen, void *params);

	double _cgo_gsl_monte_function_proxy(double *x, size_t dim, void * params) {
		return gslMonteFunctionCaller(x, dim, params);
	}

	void _init_proxy_gsl_monte_function(void *params, size_t dim, gsl_monte_function *gslMonteFunc) {
		gslMonteFunc->f = _cgo_gsl_monte_function_proxy;
		gslMonteFunc->dim = dim;
		gslMonteFunc->params = params;
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

var GSL_MONTE_FUNCTION_STRUCT_SIZE uint32 = uint32(C.get_gsl_monte_function_struct_size())

func InitializeGslMonteFunction(ptr *GslMonteFunction) {
	ptr.cGslFunctionStruct = make([]byte, 0, GSL_MONTE_FUNCTION_STRUCT_SIZE)
	cPtr := ((*reflect.SliceHeader)(unsafe.Pointer(&ptr.cGslFunctionStruct))).Data
	C._init_proxy_gsl_monte_function(unsafe.Pointer(ptr), C.size_t(ptr.Dim), (*C.gsl_monte_function)(unsafe.Pointer(cPtr)))
}

func (gf GslMonteFunction) CPtr() uintptr {
	return ((*reflect.SliceHeader)(unsafe.Pointer(&gf.cGslFunctionStruct))).Data
}
