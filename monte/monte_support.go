package monte

/*
	#include <unistd.h>
	#include <string.h>
	#include <gsl/gsl_monte.h>
	#include <gsl/gsl_monte_miser.h>
	#include <gsl/gsl_monte_vegas.h>

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

	void set_miser_params(gsl_monte_miser_state *s, double alpha, double dither, double estimate_frac,
	                      int min_calls, int min_calls_per_bisection) {
		gsl_monte_miser_params P;
		bzero(&P, sizeof(gsl_monte_miser_params));
		P.alpha = alpha;
		P.dither = dither;
		P.estimate_frac = estimate_frac;
		P.min_calls = min_calls;
		P.min_calls_per_bisection = min_calls_per_bisection;
		gsl_monte_miser_params_set(s, &P);
	}

	void get_miser_params(gsl_monte_miser_state *s, double *alpha, double *dither, double *estimate_frac,
	                      int *min_calls, int *min_calls_per_bisection) {
		gsl_monte_miser_params P;
		bzero(&P, sizeof(gsl_monte_miser_params));
		gsl_monte_miser_params_get(s, &P);
		*alpha = P.alpha;
		*dither = P.dither;
		*estimate_frac = P.estimate_frac;
		*min_calls = P.min_calls;
		*min_calls_per_bisection = P.min_calls_per_bisection;
	}

	void set_vegas_params(gsl_monte_vegas_state *s, double alpha, int iterations, int stage, int mode,
	                      int verbose, FILE *ostream) {
		gsl_monte_vegas_params P;
		bzero(&P, sizeof(gsl_monte_vegas_params));
		P.alpha = alpha;
		P.iterations = iterations;
		P.stage = stage;
		P.mode = mode;
		P.verbose = verbose;
		P.ostream = ostream;
		gsl_monte_vegas_params_set(s, &P);
	}

	void get_vegas_params(gsl_monte_vegas_state *s, double *alpha, int *iterations, int *stage, int *mode,
	                      int *verbose, FILE **ostream) {
		gsl_monte_vegas_params P;
		bzero(&P, sizeof(gsl_monte_vegas_params));
		gsl_monte_vegas_params_get(s, &P);
		*alpha = P.alpha;
		*iterations = P.iterations;
		*stage = P.stage;
		*mode = P.mode;
		*verbose = P.verbose;
		*ostream = P.ostream;
	}

*/
import "C"

import (
	"github.com/dtromb/gogsl"
	"os"
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

func MiserParamsSet(s *GslMonteMiserState, params *GslMonteMiserParams) {
	C.set_miser_params((*C.gsl_monte_miser_state)(unsafe.Pointer(s.Ptr())),
		C.double(params.Alpha), C.double(params.Dither),
		C.double(params.EstimateFrac), C.int(params.MinCalls),
		C.int(params.MinCallsPerBisection))
}

func MiserParamsGet(s *GslMonteMiserState, params *GslMonteMiserParams) {
	var minCalls C.int
	var minCallsPerBisection C.int
	C.get_miser_params((*C.gsl_monte_miser_state)(unsafe.Pointer(s.Ptr())),
		(*C.double)(&params.Alpha), (*C.double)(&params.Dither),
		(*C.double)(&params.EstimateFrac), (*C.int)(&minCalls),
		(*C.int)(&minCallsPerBisection))
	params.MinCalls = int(minCalls)
	params.MinCallsPerBisection = int(minCallsPerBisection)
}

func VegasParamsSet(s *GslMonteVegasState, params *GslMonteVegasParams) {
	C.set_vegas_params((*C.gsl_monte_vegas_state)(unsafe.Pointer(s.Ptr())),
		C.double(params.Alpha), C.int(params.Iterations),
		C.int(params.Stage), C.int(params.Mode),
		C.int(params.Verbose),
		C.fdopen(C.dup(C.int(params.Ostream.Fd())),
			(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr()))))
}

func VegasParamsGet(s *GslMonteVegasState, params *GslMonteVegasParams) {
	var iterations C.int
	var stage C.int
	var mode C.int
	var verbose C.int
	var file *C.FILE
	C.get_vegas_params((*C.gsl_monte_vegas_state)(unsafe.Pointer(s.Ptr())),
		(*C.double)(&params.Alpha), &iterations,
		&stage, &mode, &verbose, &file)
	params.Iterations = int(iterations)
	params.Stage = int(stage)
	params.Mode = GslMonteVegasMode(int(mode))
	params.Verbose = int(verbose)
	params.Ostream = os.NewFile(uintptr(C.fileno(file)), "")
}
