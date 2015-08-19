package permutation

/*
	#include <gsl/gsl_permutation.h>
	#include <gsl/gsl_permute_double.h>

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

// We make a decisive judgement call here and break the mapping - replacing size_t *
// with a permutation object.  It's difficult to get a []C.size_t in Go and easy to
// get a permutation object.

//int gsl_permute (const size_t * p, double * data, size_t stride, size_t n)
func Permute(p *GslPermutation, data []float64, stride int, n int) {
	C.gsl_permute((*C.size_t)(C.get_permutation_data((*C.gsl_permutation)(unsafe.Pointer(p.Ptr())))),
		(*C.double)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&data)).Data)),
		C.size_t(stride), C.size_t(n))
}

//int gsl_permute_inverse (const size_t * p, double * data, size_t stride, size_t n)
func PermuteInverse(p *GslPermutation, data []float64, stride int, n int) {
	C.gsl_permute_inverse((*C.size_t)(C.get_permutation_data((*C.gsl_permutation)(unsafe.Pointer(p.Ptr())))),
		(*C.double)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&data)).Data)),
		C.size_t(stride), C.size_t(n))
}

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
