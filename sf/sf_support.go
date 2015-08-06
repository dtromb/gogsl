package sf

/*
	#include <gsl/gsl_sf_result.h>
	#include <stdlib.h>
	
	size_t get_gsl_sf_result_struct_size() {
		return sizeof(gsl_sf_result);
	}
		
	size_t get_gsl_sf_result_e10_struct_size() {
		return sizeof(gsl_sf_result_e10);
	}
	
	double get_gsl_sf_result_val(gsl_sf_result *r) {
		return r->val;
	}
	
	double get_gsl_sf_result_err(gsl_sf_result *r) {
		return r->err;
	}
		
	double get_gsl_sf_result_e10_val(gsl_sf_result_e10 *r) {
		return r->val;
	}
	
	double get_gsl_sf_result_e10_err(gsl_sf_result_e10 *r) {
		return r->err;
	}
	
	int get_gsl_sf_result_e10_e10(gsl_sf_result_e10 *r) {
		return r->e10;
	}
	
*/
import "C"

import (
	"reflect"
	"unsafe"
	"github.com/dtromb/gogsl"
)

var GSL_SF_RESULT_STRUCT_SIZE		int32 = int32(C.get_gsl_sf_result_struct_size())
var GSL_SF_RESULT_E10_STRUCT_SIZE	int32 = int32(C.get_gsl_sf_result_e10_struct_size())

// Result types follow a "stack allocated" pattern in C (see gsl manual)
// To keep the code looking and operation the same, we auto-initialize these 
// reference types into byte array fields they hold in (Go-managed) memory.
// This way, the value of a &GslSfResult{} for example can be passed to the
// sf functions that take that pointer.

func (r *GslSfResult) Ptr() uintptr {
	ptr := r.GslReference.Ptr()
	if ptr != 0 {
		return ptr
	}
	r.cData = make([]byte, GSL_SF_RESULT_STRUCT_SIZE)
	ptr = ((*reflect.SliceHeader)(unsafe.Pointer(&r.cData))).Data
	gogsl.InitializeGslStruct(r,ptr)
	return ptr
}

func (r *GslSfResult) Dispose() {
	r.cData = nil
}

func (r *GslSfResult) Val() float64 {
	return float64(C.get_gsl_sf_result_val((*C.gsl_sf_result)(unsafe.Pointer(r.Ptr()))))
}

func (r *GslSfResult) Err() float64 {
	return float64(C.get_gsl_sf_result_err((*C.gsl_sf_result)(unsafe.Pointer(r.Ptr()))))
}

func (r *GslSfResultE10) Ptr() uintptr {
	ptr := r.GslReference.Ptr()
	if ptr != 0 {
		return ptr
	}
	r.cData = make([]byte, GSL_SF_RESULT_E10_STRUCT_SIZE)
	ptr = ((*reflect.SliceHeader)(unsafe.Pointer(&r.cData))).Data
	gogsl.InitializeGslStruct(r,ptr)
	return ptr
}

func (r *GslSfResultE10) Dispose() {
	r.cData = nil
}

func (r *GslSfResultE10) Val() float64 {
	return float64(C.get_gsl_sf_result_e10_val((*C.gsl_sf_result_e10)(unsafe.Pointer(r.Ptr()))))
}

func (r *GslSfResultE10) Err() float64 {
	return float64(C.get_gsl_sf_result_e10_err((*C.gsl_sf_result_e10)(unsafe.Pointer(r.Ptr()))))
}

func (r *GslSfResultE10) E10() int {
	return int(C.get_gsl_sf_result_e10_e10((*C.gsl_sf_result_e10)(unsafe.Pointer(r.Ptr()))))
}
