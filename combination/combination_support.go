package combination

/*
	#include <gsl/gsl_combination.h>

	void *get_combination_data(gsl_combination *p) {
		return p->data;
	}

*/
import "C"

import (
	"github.com/dtromb/gogsl"
	"reflect"
	"unsafe"
)

func (c *GslCombination) Slice_() interface{} {
	baseType := gogsl.GOGSL_SIZE_T_TYPE
	sliceType := reflect.SliceOf(baseType)
	size := K(c)
	hdr := &reflect.SliceHeader{Len: size, Cap: size, Data: uintptr(C.get_combination_data((*C.gsl_combination)(unsafe.Pointer(c.Ptr()))))}
	return reflect.NewAt(sliceType, unsafe.Pointer(hdr)).Elem().Interface()
}
