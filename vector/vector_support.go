package vector

/*
	#include <gsl/gsl_vector.h>
	
	void *get_gsl_vector_data(gsl_vector *v) {
		return v->data;
	}
	
	size_t get_gsl_vector_length(gsl_vector *v) {
		return v->size;
	}
	
	size_t get_gsl_vector_stride(gsl_vector *v) {
		return v->stride;
	}
	
	gsl_vector *get_vector_view_vector(gsl_vector_view *v) {
		return &v->vector;
	}
	
	gsl_vector_float *get_vector_float_view_vector(gsl_vector_float_view *v) {
		return &v->vector;
	}
	
	gsl_vector_int *get_vector_int_view_vector(gsl_vector_int_view *v) {
		return &v->vector;
	}
	
	gsl_vector_uint *get_vector_uint_view_vector(gsl_vector_uint_view *v) {
		return &v->vector;
	}
	
	gsl_vector_long *get_vector_long_view_vector(gsl_vector_long_view *v) {
		return &v->vector;
	}
	
	gsl_vector_ulong *get_vector_ulong_view_vector(gsl_vector_ulong_view *v) {
		return &v->vector;
	}
	
	gsl_vector_short *get_vector_short_view_vector(gsl_vector_short_view *v) {
		return &v->vector;
	}
	
	gsl_vector_ushort *get_vector_ushort_view_vector(gsl_vector_ushort_view *v) {
		return &v->vector;
	}
	
	gsl_vector_char *get_vector_char_view_vector(gsl_vector_char_view *v) {
		return &v->vector;
	}
	
	gsl_vector_uchar *get_vector_uchar_view_vector(gsl_vector_uchar_view *v) {
		return &v->vector;
	}
	
	gsl_vector_complex *get_vector_complex_view_vector(gsl_vector_complex_view *v) {
		return &v->vector;
	}
	
	gsl_vector_complex_float *get_vector_complex_float_view_vector(gsl_vector_complex_float_view *v) {
		return &v->vector;
	}
	
*/
import "C"

import (
	"unsafe"
	"github.com/dtromb/gogsl"
)

// XXX - Do vector accessors for all the types.
func (gvv *GslVectorView) Dispose() {
}
func (vv *GslVectorView) Vector() *GslVector {
	v := &GslVector{}
	ptr := C.get_vector_view_vector((*C.gsl_vector_view)(unsafe.Pointer(vv.Ptr())))
	gogsl.InitializeGslStruct(v, uintptr(unsafe.Pointer(ptr)))
	return v
}
func (gvv *GslVectorView) Data() []byte {
	return gvv.CData
}

func (gvcv *GslVectorConstView) Dispose() {
	
}
func (gvv *GslVectorConstView) Data() []byte {
	return gvv.CData
}

func (gvv *GslVectorFloatView) Dispose() {
	
}

func (gvv *GslVectorFloatView) Data() []byte {
	return gvv.CData
}
func (vv *GslVectorFloatView) Vector() *GslVectorFloat {
	v := &GslVectorFloat{}
	ptr := C.get_vector_float_view_vector((*C.gsl_vector_float_view)(unsafe.Pointer(vv.Ptr())))
	gogsl.InitializeGslStruct(v, uintptr(unsafe.Pointer(ptr)))
	return v
}

func (gvcv *GslVectorFloatConstView) Dispose() {
	
}
func (gvv *GslVectorFloatConstView) Data() []byte {
	return gvv.CData
}

func (gvv *GslVectorIntView) Dispose() {
	
}
func (gvv *GslVectorIntView) Data() []byte {
	return gvv.CData
}
func (vv *GslVectorIntView) Vector() *GslVectorInt {
	v := &GslVectorInt{}
	ptr := C.get_vector_int_view_vector((*C.gsl_vector_int_view)(unsafe.Pointer(vv.Ptr())))
	gogsl.InitializeGslStruct(v, uintptr(unsafe.Pointer(ptr)))
	return v
}

func (gvcv *GslVectorIntConstView) Dispose() {
	
}
func (gvv *GslVectorIntConstView) Data() []byte {
	return gvv.CData
}

func (gvv *GslVectorUintView) Dispose() {
	
}
func (gvv *GslVectorUintView) Data() []byte {
	return gvv.CData
}
func (vv *GslVectorUintView) Vector() *GslVectorUint {
	v := &GslVectorUint{}
	ptr := C.get_vector_uint_view_vector((*C.gsl_vector_uint_view)(unsafe.Pointer(vv.Ptr())))
	gogsl.InitializeGslStruct(v, uintptr(unsafe.Pointer(ptr)))
	return v
}

func (gvcv *GslVectorUintConstView) Dispose() {
	
}
func (gvv *GslVectorUintConstView) Data() []byte {
	return gvv.CData
}

func (gvv *GslVectorLongView) Dispose() {
	
}
func (gvv *GslVectorLongView) Data() []byte {
	return gvv.CData
}
func (vv *GslVectorLongView) Vector() *GslVectorLong {
	v := &GslVectorLong{}
	ptr := C.get_vector_long_view_vector((*C.gsl_vector_long_view)(unsafe.Pointer(vv.Ptr())))
	gogsl.InitializeGslStruct(v, uintptr(unsafe.Pointer(ptr)))
	return v
}

func (gvcv *GslVectorLongConstView) Dispose() {
	
}
func (gvv *GslVectorLongConstView) Data() []byte {
	return gvv.CData
}

func (gvv *GslVectorUlongView) Dispose() {
	
}
func (gvv *GslVectorUlongView) Data() []byte {
	return gvv.CData
}
func (vv *GslVectorUlongView) Vector() *GslVectorUlong {
	v := &GslVectorUlong{}
	ptr := C.get_vector_ulong_view_vector((*C.gsl_vector_ulong_view)(unsafe.Pointer(vv.Ptr())))
	gogsl.InitializeGslStruct(v, uintptr(unsafe.Pointer(ptr)))
	return v
}

func (gvcv *GslVectorUlongConstView) Dispose() {
	
}
func (gvv *GslVectorUlongConstView) Data() []byte {
	return gvv.CData
}

func (gvv *GslVectorShortView) Dispose() {
	
}
func (gvv *GslVectorShortView) Data() []byte {
	return gvv.CData
}
func (vv *GslVectorShortView) Vector() *GslVectorShort {
	v := &GslVectorShort{}
	ptr := C.get_vector_short_view_vector((*C.gsl_vector_short_view)(unsafe.Pointer(vv.Ptr())))
	gogsl.InitializeGslStruct(v, uintptr(unsafe.Pointer(ptr)))
	return v
}

func (gvcv *GslVectorShortConstView) Dispose() {
	
}
func (gvv *GslVectorShortConstView) Data() []byte {
	return gvv.CData
}

func (gvv *GslVectorUshortView) Dispose() {
	
}
func (gvv *GslVectorUshortView) Data() []byte {
	return gvv.CData
}
func (vv *GslVectorUshortView) Vector() *GslVectorUshort {
	v := &GslVectorUshort{}
	ptr := C.get_vector_ushort_view_vector((*C.gsl_vector_ushort_view)(unsafe.Pointer(vv.Ptr())))
	gogsl.InitializeGslStruct(v, uintptr(unsafe.Pointer(ptr)))
	return v
}

func (gvcv *GslVectorUshortConstView) Dispose() {
	
}
func (gvv *GslVectorUshortConstView) Data() []byte {
	return gvv.CData
}

func (gvv *GslVectorCharView) Dispose() {
	
}
func (gvv *GslVectorCharView) Data() []byte {
	return gvv.CData
}
func (vv *GslVectorCharView) Vector() *GslVectorChar {
	v := &GslVectorChar{}
	ptr := C.get_vector_char_view_vector((*C.gsl_vector_char_view)(unsafe.Pointer(vv.Ptr())))
	gogsl.InitializeGslStruct(v, uintptr(unsafe.Pointer(ptr)))
	return v
}

func (gvcv *GslVectorCharConstView) Dispose() {
	
}
func (gvv *GslVectorCharConstView) Data() []byte {
	return gvv.CData
}

func (gvv *GslVectorUcharView) Dispose() {
	
}
func (gvv *GslVectorUcharView) Data() []byte {
	return gvv.CData
}
func (vv *GslVectorUcharView) Vector() *GslVectorUchar {
	v := &GslVectorUchar{}
	ptr := C.get_vector_uchar_view_vector((*C.gsl_vector_uchar_view)(unsafe.Pointer(vv.Ptr())))
	gogsl.InitializeGslStruct(v, uintptr(unsafe.Pointer(ptr)))
	return v
}

func (gvcv *GslVectorUcharConstView) Dispose() {
	
}
func (gvv *GslVectorUcharConstView) Data() []byte {
	return gvv.CData
}

func (gvv *GslVectorComplexView) Dispose() {
	
}
func (gvv *GslVectorComplexView) Data() []byte {
	return gvv.CData
}
func (vv *GslVectorComplexView) Vector() *GslVectorComplex {
	v := &GslVectorComplex{}
	ptr := C.get_vector_complex_view_vector((*C.gsl_vector_complex_view)(unsafe.Pointer(vv.Ptr())))
	gogsl.InitializeGslStruct(v, uintptr(unsafe.Pointer(ptr)))
	return v
}

func (gvcv *GslVectorComplexConstView) Dispose() {
	
}
func (gvv *GslVectorComplexConstView) Data() []byte {
	return gvv.CData
}

func (gvv *GslVectorComplexFloatView) Dispose() {
	
}
func (gvv *GslVectorComplexFloatView) Data() []byte {
	return gvv.CData
}
func (vv *GslVectorComplexFloatView) Vector() *GslVectorComplexFloat {
	v := &GslVectorComplexFloat{}
	ptr := C.get_vector_complex_float_view_vector((*C.gsl_vector_complex_float_view)(unsafe.Pointer(vv.Ptr())))
	gogsl.InitializeGslStruct(v, uintptr(unsafe.Pointer(ptr)))
	return v
}

func (gvcv *GslVectorComplexFloatConstView) Dispose() {
	
}
func (gvv *GslVectorComplexFloatConstView) Data() []byte {
	return gvv.CData
}


