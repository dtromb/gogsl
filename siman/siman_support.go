package siman

/*

    #cgo pkg-config: --define-variable=prefix=. gsl
	#include <gsl/gsl_siman.h>
	#include <gsl/gsl_rng.h>
	#include <stdio.h>

	void siman_print_string(char *s) {
		printf("%s",s);
	}

	size_t get_gsl_siman_params_struct_size() {
		return sizeof(gsl_siman_params_t);
	}

	void _initialize_gsl_siman_params(gsl_siman_params_t *params,
									  int num_tries, int iters_fixed,
									  double step_size, double k,
									  double t_initial, double mu, double t_min) {
		params->n_tries = num_tries;
		params->iters_fixed_T = iters_fixed;
		params->step_size = step_size;
		params->k = k;
		params->t_initial = t_initial;
		params->mu_t = mu;
		params->t_min = t_min;
	}

	void breakpoint() {
		printf("");
	}

	extern double gslSimanEnergyFunctionCaller(void*);
	double _gsl_siman_energy_function_proxy(void *xp) {
		return gslSimanEnergyFunctionCaller(xp);
	}

	extern void gslSimanStepFunctionCaller(void*,void*,double);
	void _gsl_siman_step_function_proxy(gsl_rng *r, void *xp, double step_size) {
		//printf("CALLER: %8.8X\n", r);
		gslSimanStepFunctionCaller(r,xp,step_size);
	}

	extern double gslSimanMetricFunctionCaller(void*,void*);
	double _gsl_siman_metric_function_proxy(void *xp, void *yp) {
		return gslSimanMetricFunctionCaller(xp,yp);
	}

	extern void gslSimanPrintFunctionCaller(void*);
	void _gsl_siman_print_function_proxy(void *xp) {
		gslSimanPrintFunctionCaller(xp);
	}

	extern void gslSimanCopyFunctionCaller(void*,void*);
	void _gsl_siman_copy_function_proxy(void *xp, void *yp) {
		gslSimanCopyFunctionCaller(xp,yp);
	}

	extern void *gslSimanCopyConstructorCaller(void*);
	void *_gsl_siman_ctor_function_proxy(void *xp) {
		return gslSimanCopyConstructorCaller(xp);
	}

	extern void gslSimanFreeFunctionCaller(void*);
	void _gsl_siman_dtor_function_proxy(void *xp) {
		gslSimanFreeFunctionCaller(xp);
	}

*/
import "C"

import (
	"github.com/dtromb/gogsl/rng"
	"reflect"
	"unsafe"
)

var GSL_SIMAN_PARAMS_STRUCT_SIZE uint32 = uint32(C.get_gsl_siman_params_struct_size())

func InitializeGslSimanParams(params *GslSimanParams) {
	params.CData = make([]byte, GSL_SIMAN_PARAMS_STRUCT_SIZE)
	C._initialize_gsl_siman_params(
		(*C.gsl_siman_params_t)(unsafe.Pointer(params.CPtr())),
		C.int(params.NumTries),
		C.int(params.ItersFixed),
		C.double(params.StepSize),
		C.double(params.K),
		C.double(params.TInitial),
		C.double(params.Mu),
		C.double(params.TMin),
	)
}

func (sp *GslSimanParams) CPtr() uintptr {
	return (*reflect.SliceHeader)(unsafe.Pointer(&sp.CData)).Data
}

var GSL_SIMAN_ARGUMENT_SIZE int = int(reflect.TypeOf([]GslSimanArgument{}).Elem().Size())

func Solve(r *rng.GslRng, x0 interface{},
	ef GslSimanEnergyFunctionType,
	takeStep GslSimanStepFunctionType,
	distance GslSimanMetricFunctionType,
	printPosition GslSimanPrintFunctionType,
	copyfunc GslSimanCopyFunctionType,
	copyConstructor GslSimanCopyConstructFunctionType,
	destructor GslSimanDestroyFunctionType,
	params *GslSimanParams) {

	impl := &GslSimanImplementation{
		energyFn: ef,
		stepFn:   takeStep,
		metricFn: distance,
		printFn:  printPosition,
		copyFn:   copyfunc,
		ctorFn:   copyConstructor,
		freeFn:   destructor,
		holdRefs: make(map[uintptr]*GslSimanArgument),
	}

	initialArg := &GslSimanArgument{
		impl: impl,
		x:    x0,
	}
	InitializeGslSimanParams(params)
	var printFn *[0]byte
	if impl.printFn != nil {
		printFn = (*[0]byte)(C._gsl_siman_print_function_proxy)
	}
	C.gsl_siman_solve((*C.gsl_rng)(unsafe.Pointer(r.Ptr())),
		unsafe.Pointer(initialArg),
		(*[0]byte)(C._gsl_siman_energy_function_proxy),
		(*[0]byte)(C._gsl_siman_step_function_proxy),
		(*[0]byte)(C._gsl_siman_metric_function_proxy),
		printFn,
		(*[0]byte)(C._gsl_siman_copy_function_proxy),
		(*[0]byte)(C._gsl_siman_ctor_function_proxy),
		(*[0]byte)(C._gsl_siman_dtor_function_proxy),
		C.size_t(GSL_SIMAN_ARGUMENT_SIZE),
		*(*C.gsl_siman_params_t)(unsafe.Pointer(params.CPtr())))
}

func Print(s string) {
	cstr := C.CString(s)
	C.siman_print_string(cstr)
	C.free(unsafe.Pointer(cstr))
}
