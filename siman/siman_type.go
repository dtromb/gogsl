package siman

/*
	#include <gsl/gsl_siman.h>
	#include <stdio.h>

*/
import "C"

import (
	//"fmt"
	"github.com/dtromb/gogsl"
	"github.com/dtromb/gogsl/rng"
	"reflect"
	"unsafe"
)

type GslSimanEnergyFunctionType func(xp interface{}) float64
type GslSimanStepFunctionType func(r *rng.GslRng, xp interface{}, stepSize float64)
type GslSimanMetricFunctionType func(xp interface{}, yp interface{}) float64
type GslSimanPrintFunctionType func(xp interface{})
type GslSimanCopyFunctionType func(src interface{}, dest interface{})
type GslSimanCopyConstructFunctionType func(xp interface{}) interface{}
type GslSimanDestroyFunctionType func(xp interface{})

type GslSimanImplementation struct {
	energyFn GslSimanEnergyFunctionType
	stepFn   GslSimanStepFunctionType
	metricFn GslSimanMetricFunctionType
	printFn  GslSimanPrintFunctionType
	copyFn   GslSimanCopyFunctionType
	ctorFn   GslSimanCopyConstructFunctionType
	freeFn   GslSimanDestroyFunctionType
	holdRefs map[uintptr]*GslSimanArgument
}

type GslSimanArgument struct {
	impl *GslSimanImplementation
	x    interface{}
}

type GslSimanParams struct {
	NumTries   int
	ItersFixed int
	StepSize   float64
	K          float64
	TInitial   float64
	Mu         float64
	TMin       float64
	CData      []byte
}

//export gslSimanEnergyFunctionCaller
func gslSimanEnergyFunctionCaller(xptr uintptr) float64 {
	//fmt.Println("E")
	xarg := (*GslSimanArgument)(unsafe.Pointer(xptr))
	//fmt.Printf("E inval=%f @%8.8X\n", *xarg.x.(*float64), reflect.ValueOf(xarg.x).Pointer())
	result := xarg.impl.energyFn(xarg.x)
	//fmt.Printf(" result=%f\n", result)
	return result
}

//export gslSimanStepFunctionCaller
func gslSimanStepFunctionCaller(rngptr uintptr, xptr uintptr, stepSize float64) {
	//fmt.Println("S")
	r := &rng.GslRng{}
	gogsl.InitializeGslStruct(r, rngptr)
	xarg := (*GslSimanArgument)(unsafe.Pointer(xptr))
	//fmt.Printf("S inval=%f @ %8.8x\n", *xarg.x.(*float64), reflect.ValueOf(xarg.x).Pointer())
	xarg.impl.stepFn(r, xarg.x, stepSize)
	//fmt.Printf("  outval=%f @ %8.8x\n", *xarg.x.(*float64), reflect.ValueOf(xarg.x).Pointer())
}

//export gslSimanMetricFunctionCaller
func gslSimanMetricFunctionCaller(xptr uintptr, yptr uintptr) float64 {
	//fmt.Println("M")
	xarg := (*GslSimanArgument)(unsafe.Pointer(xptr))
	yarg := (*GslSimanArgument)(unsafe.Pointer(yptr))
	return xarg.impl.metricFn(xarg.x, yarg.x)
}

//export gslSimanPrintFunctionCaller
func gslSimanPrintFunctionCaller(xptr uintptr) {
	//fmt.Println("P")
	xarg := (*GslSimanArgument)(unsafe.Pointer(xptr))
	xarg.impl.printFn(xarg.x)
}

//export gslSimanCopyFunctionCaller
func gslSimanCopyFunctionCaller(srcptr uintptr, dstptr uintptr) {
	//fmt.Println("C")
	srcarg := (*GslSimanArgument)(unsafe.Pointer(srcptr))
	dstarg := (*GslSimanArgument)(unsafe.Pointer(dstptr))
	if srcarg.impl.copyFn != nil {
		srcarg.impl.copyFn(srcarg.x, dstarg.x)
	} else {
		x := reflect.ValueOf(srcarg.x)
		xType := x.Type()
		y := reflect.ValueOf(dstarg.x)
		// XXX - Are these use cases the correct conventions to use?
		//       We want to mimic the C fixed-length behavior in a reasonable way...
		//
		// If it is a pointer, memcpy bytes sized by the base type.
		if xType.Kind() == reflect.Ptr {
			y.Elem().Set(x.Elem())
		} else {
			// Otherwise, simply assign
			panic("ugh")
			dstarg.x = srcarg.x
		}
		//fmt.Printf("C RV x is %s inval=%f @%8.8X, outval=%f @%8.8X\n",
		//		reflect.TypeOf(dstarg.x).String(),
		//		*srcarg.x.(*float64), reflect.ValueOf(srcarg.x).Pointer(),
		//		*dstarg.x.(*float64), reflect.ValueOf(dstarg.x).Pointer())
	}
}

//export gslSimanCopyConstructorCaller
func gslSimanCopyConstructorCaller(xptr uintptr) uintptr {
	//fmt.Println("CC")
	xarg := (*GslSimanArgument)(unsafe.Pointer(xptr))
	if xarg.impl.ctorFn != nil {
		nptr := xarg.impl.ctorFn(xarg.x)
		rv := &GslSimanArgument{impl: xarg.impl, x: nptr}
		rvptr := uintptr(unsafe.Pointer(rv))
		xarg.impl.holdRefs[rvptr] = rv
		//fmt.Printf("CC. RV x is %s inval=%f, outval=%f\n", reflect.TypeOf(rv.x).String(), *xarg.x.(*float64), *rv.x.(*float64))
		return rvptr
	} else {
		var rv *GslSimanArgument
		x := reflect.ValueOf(xarg.x)
		xType := x.Type()
		baseType := xType.Elem()
		if xType.Kind() == reflect.Ptr {
			nxptr := reflect.New(baseType)
			//C.memcpy(unsafe.Pointer(nxptr), unsafe.Pointer(x.Pointer()), C.size_t(baseType.Size()))
			nxptr.Elem().Set(x.Elem())
			rv = &GslSimanArgument{impl: xarg.impl, x: reflect.NewAt(baseType, unsafe.Pointer(nxptr.Pointer())).Interface()}
		} else {
			panic("fail")
			rv = &GslSimanArgument{impl: xarg.impl, x: xarg.x}
		}
		//fmt.Printf("CC RV x is %s inval=%f, outval=%f\n", reflect.TypeOf(rv.x).String(), *xarg.x.(*float64), *rv.x.(*float64))
		rvptr := uintptr(unsafe.Pointer(rv))
		xarg.impl.holdRefs[rvptr] = rv
		return rvptr
	}
}

//export gslSimanFreeFunctionCaller
func gslSimanFreeFunctionCaller(xptr uintptr) {
	//fmt.Println("D")
	xarg := (*GslSimanArgument)(unsafe.Pointer(xptr))
	delete(xarg.impl.holdRefs, xptr)
	if xarg.impl.freeFn != nil {
		xarg.impl.freeFn(xarg.x)
	}
}
