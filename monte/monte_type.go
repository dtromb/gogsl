package monte

/*
	#include <gsl/gsl_math.h>
	#include <stdlib.h>

*/
import "C"

import (
	"os"
	"reflect"
	"unsafe"
)

type GslMonteFunctionType func(x []float64, dim int, params interface{}) float64

type GslCMonteFunction uintptr

type GslMonteFunction struct {
	Function           GslMonteFunctionType
	Params             interface{}
	Dim                int
	cGslFunctionStruct []byte
}

type GslMonteVegasMode int32

const (
	GSL_VEGAS_MODE_IMPORTANCE      GslMonteVegasMode = 1
	GSL_VEGAS_MODE_IMPORTANCE_ONLY GslMonteVegasMode = -1
	GSL_VEGAS_MODE_STRATIFIED      GslMonteVegasMode = 0
)

type GslMonteVegasParams struct {
	Alpha      float64
	Iterations int
	Stage      int
	Mode       GslMonteVegasMode
	Verbose    int
	Ostream    *os.File
}

type GslMonteMiserParams struct {
	EstimateFrac         float64
	MinCalls             int
	MinCallsPerBisection int
	Alpha                float64
	Dither               float64
}

//export gslMonteFunctionCaller
func gslMonteFunctionCaller(x uintptr, xlen int, cFunParamPtr uintptr) float64 {
	hdr := &reflect.SliceHeader{Len: xlen, Cap: xlen, Data: x}
	slice := *(*[]float64)(unsafe.Pointer(hdr))
	gslf := (*GslMonteFunction)(unsafe.Pointer(cFunParamPtr))
	return gslf.Function(slice, xlen, gslf.Params)
}
