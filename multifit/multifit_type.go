package multifit

/*
	#include <gsl/gsl_math.h>
	#include <stdlib.h>

*/
import "C"

import (
	"github.com/dtromb/gogsl"
	"github.com/dtromb/gogsl/matrix"
	"github.com/dtromb/gogsl/vector"
	"unsafe"
)

type GslMultifitFunctionType func(x *vector.GslVector, params interface{}, f *vector.GslVector) int
type GslMultifitFunctionDfType func(x *vector.GslVector, params interface{}, df *matrix.GslMatrix) int
type GslMultifitFunctionFdfType func(x *vector.GslVector, params interface{}, f *vector.GslVector, df *matrix.GslMatrix) int

type GslCMultifitFunction uintptr
type GslCMultifitFunctionFdf uintptr

type GslMultifitFunction struct {
	Function           GslMultifitFunctionType
	Dimension          int
	Params             interface{}
	cGslFunctionStruct []byte
}

type GslMultifitFunctionFdf struct {
	Function           GslMultifitFunctionType
	Derivative         GslMultifitFunctionDfType
	Fdf                GslMultifitFunctionFdfType
	Dimension          int
	Params             interface{}
	cGslFunctionStruct []byte
}

//export gslMultifitFunctionCaller
func gslMultifitFunctionCaller(x uintptr, cFunParamPtr uintptr, f uintptr) int {
	xVect := &vector.GslVector{}
	gogsl.InitializeGslStruct(xVect, x)
	fVect := &vector.GslVector{}
	gogsl.InitializeGslStruct(fVect, f)
	gslf := (*GslMultifitFunction)(unsafe.Pointer(cFunParamPtr))
	return gslf.Function(xVect, gslf.Params, fVect)
}

//export gslMultifitFunctionFdfFCaller
func gslMultifitFunctionFdfFCaller(x uintptr, cFunParamPtr uintptr, f uintptr) int {
	xVect := &vector.GslVector{}
	gogsl.InitializeGslStruct(xVect, x)
	fVect := &vector.GslVector{}
	gogsl.InitializeGslStruct(fVect, f)
	gslf := (*GslMultifitFunctionFdf)(unsafe.Pointer(cFunParamPtr))
	return gslf.Function(xVect, gslf.Params, fVect)
}

//export gslMultifitFunctionFdfDfCaller
func gslMultifitFunctionFdfDfCaller(x uintptr, cFunParamPtr uintptr, df uintptr) int {
	xVect := &vector.GslVector{}
	gogsl.InitializeGslStruct(xVect, x)
	dfMat := &matrix.GslMatrix{}
	gogsl.InitializeGslStruct(dfMat, df)
	gslf := (*GslMultifitFunctionFdf)(unsafe.Pointer(cFunParamPtr))
	return gslf.Derivative(xVect, gslf.Params, dfMat)
}

//export gslMultifitFunctionFdfCaller
func gslMultifitFunctionFdfCaller(x uintptr, cFunParamPtr uintptr, f uintptr, df uintptr) int {
	xVect := &vector.GslVector{}
	gogsl.InitializeGslStruct(xVect, x)
	fVect := &vector.GslVector{}
	gogsl.InitializeGslStruct(fVect, f)
	dfMat := &matrix.GslMatrix{}
	gogsl.InitializeGslStruct(dfMat, df)
	gslf := (*GslMultifitFunctionFdf)(unsafe.Pointer(cFunParamPtr))
	return gslf.Fdf(xVect, gslf.Params, fVect, dfMat)
}
