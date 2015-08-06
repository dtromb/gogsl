package randist

/*
	#include <gsl/gsl_rng.h>
	#include <gsl/gsl_randist.h>
*/
import "C"

import (
	"reflect"
	"unsafe"
	"github.com/dtromb/gogsl"
	"github.com/dtromb/gogsl/rng"
)

func Shuffle(rng *rng.GslRng, data interface{}, n int) {
	// data must be a slice
	dataType := reflect.TypeOf(data)
	if dataType.Kind() != reflect.Slice {
		gogsl.Error("rng.Shuffle() must have a slice as its second argument", gogsl.GSL_EINVAL)
		return
	}
	baseType := dataType.Elem()
	dataVal := reflect.ValueOf(data)
	sliceN := dataVal.Len()
	if sliceN < 0 {
		gogsl.Error("shuffle length may not be negative in rng.Shuffle()", gogsl.GSL_EINVAL)
		return
	}
	if sliceN < n {
		gogsl.Error("shuffle length too large for slice in rng.Shuffle()", gogsl.GSL_EINVAL)
		return
	}
	elementLen := baseType.Size()
	addr := dataVal.Index(0).UnsafeAddr()
	
	//hdr := (*reflect.SliceHeader)(unsafe.Pointer(addr))
	//void gsl_ran_shuffle (const gsl_rng * r, void * base, size_t n, size_t size)
	C.gsl_ran_shuffle((*C.gsl_rng)(unsafe.Pointer(rng.Ptr())), 
	                  unsafe.Pointer(addr),
					  C.size_t(n), C.size_t(elementLen))
}