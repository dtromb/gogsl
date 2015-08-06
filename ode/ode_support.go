package ode


/*

	#include <stdlib.h>
	#include <gsl/gsl_odeiv2.h>
	
	int _cgo_gsl_odeiv2_function_proxy(double t, const double y[], double dydt[], void * params) {
		return gslOdeiv2FunctionCaller(t, y, dydt, params);
	}
	
	int _cgo_gsl_odeiv2_jacobian_function_proxy(double t, const double y[], double * dfdy, double dfdt[], void * params) {
		return gslOdeiv2JacobianFunctionCaller(t, y, dfdy, dfdt, params);
	}
	
	void _init_proxy_gsl_odeiv2_system(void *params, size_t dim, gsl_odeiv2_system *system) {
		system->function = _cgo_gsl_odeiv2_function_proxy;
		system->jacobian = _cgo_gsl_odeiv2_jacobian_function_proxy;
		system->params = params;
	}
	
	size_t get_gsl_odeiv2_system_struct_size() {
		return sizeof(gsl_odeiv2_system);
	}
	
*/
import "C"

import (
	"reflect"
	"unsafe"
)

var global_GSL_Odeiv2_SYSTEM_STRUCT_SIZE 	uint32 = uint32(C.get_gsl_odeiv2_system_struct_size())

func InitializeGslFOdeiv2System(ptr *GslOdeiv2System) {
	ptr.cGslOdeiv2SystemStruct = make([]byte, 0, global_GSL_Odeiv2_SYSTEM_STRUCT_SIZE)
	cPtr := ((*reflect.SliceHeader)(unsafe.Pointer(&ptr.cGslOdeiv2SystemStruct))).Data
	C._init_proxy_gsl_odeiv2_system(unsafe.Pointer(ptr), C.size_t(ptr.Dimension), (*C.gsl_odeiv2_system)(unsafe.Pointer(cPtr)))
}

func (sys GslOdeiv2System) CPtr() uintptr {
	return ((*reflect.SliceHeader)(unsafe.Pointer(&sys.cGslOdeiv2SystemStruct))).Data
}