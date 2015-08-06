package multiroot

/*
	#include <gsl/gsl_math.h>
	#include <stdlib.h>
	
*/
import "C" 

import (
	"unsafe"
	"github.com/dtromb/gogsl"
	"github.com/dtromb/gogsl/vector"
	"github.com/dtromb/gogsl/matrix"
)

type GslMultirootFunctionType func(x *vector.GslVector, params interface{}, f *vector.GslVector) int
type GslMultirootFunctionDfType func(x *vector.GslVector, params interface{}, df *matrix.GslMatrix) int
type GslMultirootFunctionFdfType func(x *vector.GslVector, params interface{}, f *vector.GslVector, df *matrix.GslMatrix) int

type GslCMultirootFunction uintptr
type GslCMultirootFunctionFdf uintptr

type GslMultirootFunction struct {
	Function GslMultirootFunctionType
	Dimension int
	Params interface{}
	cGslFunctionStruct []byte
}

type GslMultirootFunctionFdf struct {
	Function GslMultirootFunctionType
	Derivative GslMultirootFunctionDfType
	Fdf GslMultirootFunctionFdfType
	Dimension int
	Params interface{}
	cGslFunctionStruct []byte
}

//export gslMultirootFunctionCaller
func gslMultirootFunctionCaller(x uintptr, cFunParamPtr uintptr, f uintptr) int {
	xVect := &vector.GslVector{}
	gogsl.InitializeGslStruct(xVect,x)
	fVect := &vector.GslVector{}
	gogsl.InitializeGslStruct(fVect, f)
	gslf := (*GslMultirootFunction)(unsafe.Pointer(cFunParamPtr))
	return gslf.Function(xVect,gslf.Params,fVect)
}

//export gslMultirootFunctionFdfFCaller
func gslMultirootFunctionFdfFCaller(x uintptr, cFunParamPtr uintptr, f uintptr) int {
	xVect := &vector.GslVector{}
	gogsl.InitializeGslStruct(xVect,x)
	fVect := &vector.GslVector{}
	gogsl.InitializeGslStruct(fVect, f)
	gslf := (*GslMultirootFunctionFdf)(unsafe.Pointer(cFunParamPtr))
	return gslf.Function(xVect,gslf.Params,fVect)
}

//export gslMultirootFunctionFdfDfCaller
func gslMultirootFunctionFdfDfCaller(x uintptr, cFunParamPtr uintptr, df uintptr) int {
	xVect := &vector.GslVector{}
	gogsl.InitializeGslStruct(xVect,x)
	dfMat := &matrix.GslMatrix{}
	gogsl.InitializeGslStruct(dfMat, df)
	gslf := (*GslMultirootFunctionFdf)(unsafe.Pointer(cFunParamPtr))
	return gslf.Derivative(xVect,gslf.Params,dfMat)
}

//export gslMultirootFunctionFdfCaller
func gslMultirootFunctionFdfCaller(x uintptr, cFunParamPtr uintptr, f uintptr, df uintptr) int {
	xVect := &vector.GslVector{}
	gogsl.InitializeGslStruct(xVect,x)
	fVect := &vector.GslVector{}
	gogsl.InitializeGslStruct(fVect, f)
	dfMat := &matrix.GslMatrix{}
	gogsl.InitializeGslStruct(dfMat, df)
	gslf := (*GslMultirootFunctionFdf)(unsafe.Pointer(cFunParamPtr))
	return gslf.Fdf(xVect,gslf.Params,fVect,dfMat)
}
