package multiset

/*
	#include <gsl/gsl_multiset.h>

	void *get_multiset_data(gsl_multiset *p) {
		return p->data;
	}

*/
import "C"

import (
	"github.com/dtromb/gogsl"
	"reflect"
	"unsafe"
)

func (m *GslMultiset) Slice_() interface{} {
	baseType := gogsl.GOGSL_SIZE_T_TYPE
	sliceType := reflect.SliceOf(baseType)
	size := K(m)
	hdr := &reflect.SliceHeader{Len: size, Cap: size, Data: uintptr(C.get_multiset_data((*C.gsl_multiset)(unsafe.Pointer(m.Ptr()))))}
	return reflect.NewAt(sliceType, unsafe.Pointer(hdr)).Elem().Interface()
}
