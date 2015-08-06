package ode

/*
	#include <gsl/gsl_odeiv2.h>
	#include <stdlib.h>

*/
import "C" 

import (
	"unsafe"
	"reflect"
)

type GslOdeiv2FunctionType func(t float64, y []float64, dydt []float64, params interface{}) int
type GslOdeiv2JacobianFunctionType func(t float64, y []float64, dfdy []float64, dfdt []float64, params interface{}) int
 
type GslCOdeiv2SystemPtr uintptr

type GslOdeiv2System struct {
	Function GslOdeiv2FunctionType
	Jacobian GslOdeiv2JacobianFunctionType
	Dimension int
	Params interface{}
	cGslOdeiv2SystemStruct []byte
}

//export gslOdeiv2FunctionCaller
func gslOdeiv2FunctionCaller(t float64, yPtr uintptr, dydtPtr uintptr, cParams uintptr) int {
	sys := (*GslOdeiv2System)(unsafe.Pointer(cParams))
	yHdr := &reflect.SliceHeader{Len: sys.Dimension, Cap: sys.Dimension, Data: yPtr}
	y := *(*[]float64)(unsafe.Pointer(yHdr))
	dyHdr := &reflect.SliceHeader{Len: sys.Dimension, Cap: sys.Dimension, Data: dydtPtr}
	dydt := *(*[]float64)(unsafe.Pointer(dyHdr))
	return sys.Function(t,y,dydt,sys.Params)
}

//export gslOdeiv2JacobianFunctionCaller
func gslOdeiv2JacobianFunctionCaller(t float64, yPtr uintptr, dfdyPtr uintptr, dfdtPtr uintptr, cParams uintptr) int {
	sys := (*GslOdeiv2System)(unsafe.Pointer(cParams))
	yHdr := &reflect.SliceHeader{Len: sys.Dimension, Cap: sys.Dimension, Data: yPtr}
	y := *(*[]float64)(unsafe.Pointer(yHdr))
	n2 := sys.Dimension * sys.Dimension
	dfdyHdr := &reflect.SliceHeader{Len: n2, Cap: n2, Data: dfdyPtr}
	dfdy := *(*[]float64)(unsafe.Pointer(dfdyHdr))
	dfdtHdr := &reflect.SliceHeader{Len: sys.Dimension, Cap: sys.Dimension, Data: dfdtPtr}
	dfdt := *(*[]float64)(unsafe.Pointer(dfdtHdr))
	return sys.Jacobian(t,y,dfdy,dfdt,sys.Params)
}

