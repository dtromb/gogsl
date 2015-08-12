package ntuple

/*
	#include <gsl/gsl_ntuple.h>
*/
import "C"

import (
	"reflect"
	"sync"
	"unsafe"
)

type GslNtupleSelectionFunctionType func(data interface{}, params interface{}) int
type GslNtupleValueFunctionType func(data interface{}, params interface{}) float64

type GslCNtupleSelectionFunction uintptr
type GslCNtupleValueFunction uintptr

type GslNtupleSelectionFunction struct {
	Function GslNtupleSelectionFunctionType
	Params   interface{}

	cStructData []byte
	lock        sync.Mutex
	ntuple      *GslNtuple // Transient value protected by lock.
	// XXX - Can we do this in some better way?
	//       The issue is that the Go crosscall function needs to wrap the C pointer
	//       into a Go slice, and the length and type information for that is not
	// 	   otherwise available...
}

type GslNtupleValueFunction struct {
	Function GslNtupleValueFunctionType
	Params   interface{}

	cStructData []byte
	lock        sync.Mutex
	ntuple      *GslNtuple // Transient value protected by lock.
}

//export gslNtupleSelectionFunctionCaller
func gslNtupleSelectionFunctionCaller(x uintptr, cFunParamPtr uintptr) int {
	ntsf := (*GslNtupleSelectionFunction)(unsafe.Pointer(cFunParamPtr))
	hdr := &reflect.SliceHeader{Data: x, Len: ntsf.ntuple.refLen, Cap: ntsf.ntuple.refLen}
	sliceType := reflect.SliceOf(ntsf.ntuple.baseType)
	slice := reflect.NewAt(sliceType, unsafe.Pointer(hdr)).Elem().Interface()
	return ntsf.Function(slice, ntsf.Params)
}

//export gslNtupleValueFunctionCaller
func gslNtupleValueFunctionCaller(x uintptr, cFunParamPtr uintptr) float64 {
	ntsf := (*GslNtupleValueFunction)(unsafe.Pointer(cFunParamPtr))
	hdr := &reflect.SliceHeader{Data: x, Len: ntsf.ntuple.refLen, Cap: ntsf.ntuple.refLen}
	sliceType := reflect.SliceOf(ntsf.ntuple.baseType)
	slice := reflect.NewAt(sliceType, unsafe.Pointer(hdr)).Elem().Interface()
	return ntsf.Function(slice, ntsf.Params)
}
