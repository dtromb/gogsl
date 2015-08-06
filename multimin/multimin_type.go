package multimin

/*
	#include <gsl/gsl_math.h>
	#include <stdlib.h>
	
*/
import "C" 

import (
	"unsafe"
	"github.com/dtromb/gogsl"
	"github.com/dtromb/gogsl/vector"
)

type GslMultiminFunctionType func(x *vector.GslVector, params interface{}) float64
type GslMultiminFunctionDfType func(x *vector.GslVector, params interface{}, df *vector.GslVector) 
type GslMultiminFunctionFdfType func(x *vector.GslVector, params interface{}, df *vector.GslVector) float64

type GslCMultiminFunction uintptr
type GslCMultiminFunctionFdf uintptr

type GslMultiminFunction struct {
	Function GslMultiminFunctionType
	Dimension int
	Params interface{}
	cGslFunctionStruct []byte
}

type GslMultiminFunctionFdf struct {
	Function GslMultiminFunctionType
	Derivative GslMultiminFunctionDfType
	Fdf GslMultiminFunctionFdfType
	Dimension int
	Params interface{}
	cGslFunctionStruct []byte
}

//export gslMultiminFunctionCaller
func gslMultiminFunctionCaller(x uintptr, cFunParamPtr uintptr) float64 {
	xVect := &vector.GslVector{}
	gogsl.InitializeGslStruct(xVect,x)
	gslf := (*GslMultiminFunction)(unsafe.Pointer(cFunParamPtr))
	return gslf.Function(xVect,gslf.Params)
}

//export gslMultiminFunctionFdfFCaller
func gslMultiminFunctionFdfFCaller(x uintptr, cFunParamPtr uintptr) float64 {
	xVect := &vector.GslVector{}
	gogsl.InitializeGslStruct(xVect,x)
	gslf := (*GslMultiminFunctionFdf)(unsafe.Pointer(cFunParamPtr))
	return gslf.Function(xVect,gslf.Params)
}

//export gslMultiminFunctionFdfDfCaller
func gslMultiminFunctionFdfDfCaller(x uintptr, cFunParamPtr uintptr, df uintptr) {
	xVect := &vector.GslVector{}
	gogsl.InitializeGslStruct(xVect,x)
	dfVect := &vector.GslVector{}
	gogsl.InitializeGslStruct(dfVect, df)
	gslf := (*GslMultiminFunctionFdf)(unsafe.Pointer(cFunParamPtr))
	gslf.Derivative(xVect,gslf.Params,dfVect)
}

//export gslMultiminFunctionFdfCaller
func gslMultiminFunctionFdfCaller(x uintptr, cFunParamPtr uintptr, df uintptr) float64 {
	xVect := &vector.GslVector{}
	gogsl.InitializeGslStruct(xVect,x)
	dfVect := &vector.GslVector{}
	gogsl.InitializeGslStruct(dfVect, df)
	gslf := (*GslMultiminFunctionFdf)(unsafe.Pointer(cFunParamPtr))
	return gslf.Fdf(xVect,gslf.Params,dfVect)
}
