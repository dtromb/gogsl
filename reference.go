package gogsl

/*
	#include <string.h>
*/
import "C"

import (
	"errors"
	//"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

type CReference interface {
	Ptr() uintptr
	AliasGroup() *ReferenceList
	Dispose()
}

type ReferenceList struct {
	ref CReference
	nxt *ReferenceList
	prv *ReferenceList
}

type GslReference struct {
	x    uintptr
	rgrp *ReferenceList
}

type CStaticReference interface {
	CReference
	Data() []byte
}

func GslReferenceForceClear(ref *GslReference) {
	ref.x = 0
}

func GslReferenceFinalizer(ref CReference) {
	//fmt.Println("FINALIZER")
	gslRef := reflect.ValueOf(ref).Elem().FieldByName("GslReference").Addr().Interface().(*GslReference)
	if gslRef.x != 0 {
		rgrp := gslRef.rgrp
		if rgrp == nil {
			ref.Dispose()
		} else {
			rl := gslRef.rgrp
			if rl.nxt == nil && rl.prv == nil {
				ref.Dispose()
			} else {
				if rl.nxt != nil {
					rl.nxt.prv = rl.prv
				}
				if rl.prv != nil {
					rl.prv.nxt = rl.nxt
				}
			}
		}
	}
}

func InitializeGslReference(ref CReference, ptr uintptr) error {
	return gslReference(ref, ptr, true)
}

func InitializeGslStruct(ref CReference, ptr uintptr) error {
	return gslReference(ref, ptr, false)
}

func InitializeGslStatic(ref CStaticReference, ptr uintptr) error {
	dataSlice := ref.Data()
	dataPtr := (*reflect.SliceHeader)(unsafe.Pointer(&dataSlice)).Data
	C.memcpy(unsafe.Pointer(dataPtr), unsafe.Pointer(ptr), C.size_t(len(dataSlice)))
	return gslReference(ref, dataPtr, false)
}

func gslReference(ref CReference, ptr uintptr, finalize bool) error {
	// Find the GslReference embed reflectively.
	ifVal := reflect.ValueOf(ref)
	if ifVal.Kind() == reflect.Ptr {
		ifVal = ifVal.Elem()
	} else {
		return errors.New("First argument to InitializeGslReference() must contain a pointer")
	}
	if ifVal.Kind() != reflect.Struct {
		return errors.New("First argument to InitializeGslReference() must contain a pointer to struct")
	}
	embed := ifVal.FieldByName("GslReference")
	if !embed.IsValid() || (embed.Type() != reflect.TypeOf([]GslReference{}).Elem()) {
		return errors.New("First argument to InitializeGslReference() must embed GslReference")
	}
	gslRef := embed.Addr().Interface().(*GslReference)
	if gslRef.x != 0 {
		return errors.New("attempt to reinitialize active reference")
	}
	gslRef.x = ptr
	//fmt.Printf("GSL REF PTR: %x\n", uintptr(unsafe.Pointer(gslRef)))
	runtime.SetFinalizer(ifVal.Addr().Interface().(CReference), GslReferenceFinalizer)
	return nil
}

func (ref *GslReference) Ptr() uintptr {
	return ref.x
}

func (ref *GslReference) AliasGroup() *ReferenceList {
	panic("unimplemented")
}
