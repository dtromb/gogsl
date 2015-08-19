package randist

/*
	#include <gsl/gsl_rng.h>
	#include <gsl/gsl_randist.h>
*/
import "C"

import (
	"github.com/dtromb/gogsl"
	"github.com/dtromb/gogsl/rng"
	"reflect"
	"unsafe"
)

// void gsl_ran_sample (const gsl_rng * r, void * dest, size_t k, void * src, size_t n, size_t size)
func Choose(rng *rng.GslRng, dest interface{}, src interface{}, n int) int {
	// data must be a slice
	destDataType := reflect.TypeOf(dest)
	if destDataType.Kind() != reflect.Slice {
		gogsl.Error("rng.Choose() must have a slice as its second argument", gogsl.GSL_EINVAL)
		return int(gogsl.GSL_EINVAL)
	}
	srcDataType := reflect.TypeOf(src)
	if srcDataType.Kind() != reflect.Slice {
		gogsl.Error("rng.Choose() must have a slice as its third argument", gogsl.GSL_EINVAL)
		return int(gogsl.GSL_EINVAL)
	}
	if destDataType != srcDataType {
		gogsl.Error("rng.Choose() source and destination data types must match", gogsl.GSL_EINVAL)
		return int(gogsl.GSL_EINVAL)
	}
	if n < 0 {
		gogsl.Error("length may not be negative in rng.Choose()", gogsl.GSL_EINVAL)
		return int(gogsl.GSL_EINVAL)
	}
	baseType := destDataType.Elem()
	destVal := reflect.ValueOf(dest)
	srcVal := reflect.ValueOf(src)
	sliceN := destVal.Len()
	if sliceN < n {
		gogsl.Error("sample length too large for slice in rng.Choose()", gogsl.GSL_EINVAL)
		return int(gogsl.GSL_EINVAL)
	}
	if srcVal.Len() < n {
		gogsl.Error("sample length too large for source in rng.Choose()", gogsl.GSL_EINVAL)
		return int(gogsl.GSL_EINVAL)
	}
	elementLen := baseType.Size()
	srcAddr := srcVal.Index(0).UnsafeAddr()
	destAddr := destVal.Index(0).UnsafeAddr()
	return int(C.gsl_ran_choose((*C.gsl_rng)(unsafe.Pointer(rng.Ptr())),
		unsafe.Pointer(destAddr), C.size_t(sliceN),
		unsafe.Pointer(srcAddr), C.size_t(srcVal.Len()),
		C.size_t(elementLen)))
}

// void gsl_ran_sample (const gsl_rng * r, void * dest, size_t k, void * src, size_t n, size_t size)
func Sample(rng *rng.GslRng, dest interface{}, src interface{}, n int) {
	// data must be a slice
	destDataType := reflect.TypeOf(dest)
	if destDataType.Kind() != reflect.Slice {
		gogsl.Error("rng.Sample() must have a slice as its second argument", gogsl.GSL_EINVAL)
		return
	}
	srcDataType := reflect.TypeOf(src)
	if srcDataType.Kind() != reflect.Slice {
		gogsl.Error("rng.Sample() must have a slice as its third argument", gogsl.GSL_EINVAL)
		return
	}
	if destDataType != srcDataType {
		gogsl.Error("rng.Sample() source and destination data types must match", gogsl.GSL_EINVAL)
		return
	}
	if n < 0 {
		gogsl.Error("length may not be negative in rng.Sample()", gogsl.GSL_EINVAL)
		return
	}
	baseType := destDataType.Elem()
	destVal := reflect.ValueOf(dest)
	srcVal := reflect.ValueOf(src)
	sliceN := destVal.Len()
	if sliceN < n {
		gogsl.Error("sample length too large for slice in rng.Sample()", gogsl.GSL_EINVAL)
		return
	}
	elementLen := baseType.Size()
	srcAddr := srcVal.Index(0).UnsafeAddr()
	destAddr := destVal.Index(0).UnsafeAddr()
	C.gsl_ran_sample((*C.gsl_rng)(unsafe.Pointer(rng.Ptr())),
		unsafe.Pointer(destAddr), C.size_t(sliceN),
		unsafe.Pointer(srcAddr), C.size_t(srcVal.Len()),
		C.size_t(elementLen))
}

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
	if n < 0 {
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
