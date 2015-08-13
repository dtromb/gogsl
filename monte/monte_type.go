package monte

/*
	#include <gsl/gsl_math.h>
	#include <stdlib.h>

*/
import "C"

import (
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

//export gslMonteFunctionCaller
func gslMonteFunctionCaller(x uintptr, xlen int, cFunParamPtr uintptr) float64 {
	hdr := &reflect.SliceHeader{Len: xlen, Cap: xlen, Data: x}
	slice := *(*[]float64)(unsafe.Pointer(hdr))
	gslf := (*GslMonteFunction)(unsafe.Pointer(cFunParamPtr))
	return gslf.Function(slice, xlen, gslf.Params)
}
