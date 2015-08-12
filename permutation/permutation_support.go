package permutation

/*
	#include <gsl/gsl_permutation.h>

	size_t get_permutation_size(gsl_permutation *p) {
		return p->size;
	}

	void *get_permutation_data(gsl_permutation *p) {
		return p->data;
	}

*/
import "C"

import (
	//"fmt"
	"github.com/dtromb/gogsl"
	"reflect"
	"unsafe"
)

func (p *GslPermutation) Len() int {
	return int(C.get_permutation_size((*C.gsl_permutation)(unsafe.Pointer(p.Ptr()))))
}

func (p *GslPermutation) Slice_() interface{} {
	baseType := gogsl.GOGSL_SIZE_T_TYPE
	sliceType := reflect.SliceOf(baseType)
	size := p.Len()
	hdr := &reflect.SliceHeader{Len: size, Cap: size, Data: uintptr(C.get_permutation_data((*C.gsl_permutation)(unsafe.Pointer(p.Ptr()))))}
	return reflect.NewAt(sliceType, unsafe.Pointer(hdr)).Elem().Interface()
}
