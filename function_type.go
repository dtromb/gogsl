package gogsl

/*
	#include <gsl/gsl_math.h>
	#include <stdlib.h>

*/
import "C"

import (
	"unsafe"
)

type GslFunctionType func(x float64, params interface{}) float64
type GslFunctionFdfType func(x float64, params interface{}) (float64, float64)

type GslCFunction uintptr
type GslCFunctionFdf uintptr

type GslFunction struct {
	Function           GslFunctionType
	Params             interface{}
	cGslFunctionStruct []byte
}

type GslFunctionFdf struct {
	Function           GslFunctionType
	Derivative         GslFunctionType
	Fdf                GslFunctionFdfType
	Params             interface{}
	cGslFunctionStruct []byte
}

//export gslFunctionCaller
func gslFunctionCaller(x float64, cFunParamPtr uintptr) float64 {
	gslf := (*GslFunction)(unsafe.Pointer(cFunParamPtr))
	return gslf.Function(x, gslf.Params)
}

//export gslFunctionFdfFCaller
func gslFunctionFdfFCaller(x float64, cFunParamPtr uintptr) float64 {
	gslf := (*GslFunctionFdf)(unsafe.Pointer(cFunParamPtr))
	return gslf.Function(x, gslf.Params)
}

//export gslFunctionFdfDfCaller
func gslFunctionFdfDfCaller(x float64, cFunParamPtr uintptr) float64 {
	gslf := (*GslFunctionFdf)(unsafe.Pointer(cFunParamPtr))
	return gslf.Derivative(x, gslf.Params)
}

//export gslFunctionFdfCaller
func gslFunctionFdfCaller(x float64, cFunParamPtr uintptr) (float64, float64) {
	gslf := (*GslFunctionFdf)(unsafe.Pointer(cFunParamPtr))
	return gslf.Fdf(x, gslf.Params)
}
