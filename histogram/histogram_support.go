package histogram

/*
	#include <gsl/gsl_histogram.h>
	#include <gsl/gsl_histogram2d.h>

	size_t histogram_get_size(gsl_histogram *h) {
		return h->n;
	}

	void *histogram_get_range_ptr(gsl_histogram *h, size_t ofs) {
		return &h->range[ofs];
	}

	void *histogram_get_value_ptr(gsl_histogram *h, size_t ofs) {
		return &h->bin[ofs];
	}

	size_t histogram2d_get_size_x(gsl_histogram2d *h) {
		return h->nx;
	}

	size_t histogram2d_get_size_y(gsl_histogram2d *h) {
		return h->ny;
	}

	void *histogram2d_get_range_ptr_x(gsl_histogram2d *h, size_t ofs) {
		return &h->xrange[ofs];
	}

	void *histogram2d_get_range_ptr_y(gsl_histogram2d *h, size_t ofs) {
		return &h->yrange[ofs];
	}

	// bin[i * ny + j]
	void *histogram2d_get_value_ptr(gsl_histogram2d *h, size_t ofs) {
		return &h->bin[ofs];
	}

*/
import "C"

import (
	"github.com/dtromb/gogsl"
	"github.com/dtromb/gogsl/matrix"
	"reflect"
	"unsafe"
)

func (h *GslHistogram) Size() int {
	return int(C.histogram_get_size((*C.gsl_histogram)(unsafe.Pointer(h.Ptr()))))
}

func (h *GslHistogram) RangeSlice(ofs, length int) []float64 {
	if length+ofs > h.Size() {
		gogsl.Error("histogram slice out of range", gogsl.GSL_ERANGE)
		return nil
	}
	ptr := C.histogram_get_range_ptr((*C.gsl_histogram)(unsafe.Pointer(h.Ptr())), C.size_t(ofs))
	srcHdr := &reflect.SliceHeader{Data: uintptr(ptr), Len: length, Cap: length}
	src := *(*[]float64)(unsafe.Pointer(srcHdr))
	dst := make([]float64, length)
	copy(dst, src)
	return dst
}

func (h *GslHistogram) RangeSlice_(ofs, length int) []float64 {
	if length+ofs > h.Size() {
		gogsl.Error("histogram slice out of range", gogsl.GSL_ERANGE)
		return nil
	}
	ptr := C.histogram_get_range_ptr((*C.gsl_histogram)(unsafe.Pointer(h.Ptr())), C.size_t(ofs))
	srcHdr := &reflect.SliceHeader{Data: uintptr(ptr), Len: length, Cap: length}
	src := *(*[]float64)(unsafe.Pointer(srcHdr))
	return src
}

func (h *GslHistogram) Slice(ofs, length int) []float64 {
	if length+ofs > h.Size() {
		gogsl.Error("histogram slice out of range", gogsl.GSL_ERANGE)
		return nil
	}
	ptr := C.histogram_get_value_ptr((*C.gsl_histogram)(unsafe.Pointer(h.Ptr())), C.size_t(ofs))
	srcHdr := &reflect.SliceHeader{Data: uintptr(ptr), Len: length, Cap: length}
	src := *(*[]float64)(unsafe.Pointer(srcHdr))
	dst := make([]float64, length)
	copy(dst, src)
	return dst
}

func (h *GslHistogram) Slice_(ofs, length int) []float64 {
	if length+ofs > h.Size() {
		gogsl.Error("histogram slice out of range", gogsl.GSL_ERANGE)
		return nil
	}
	ptr := C.histogram_get_value_ptr((*C.gsl_histogram)(unsafe.Pointer(h.Ptr())), C.size_t(ofs))
	srcHdr := &reflect.SliceHeader{Data: uintptr(ptr), Len: length, Cap: length}
	src := *(*[]float64)(unsafe.Pointer(srcHdr))
	return src
}

func (h *GslHistogram2d) Dim() []int {
	dx := int(C.histogram2d_get_size_x((*C.gsl_histogram2d)(unsafe.Pointer(h.Ptr()))))
	dy := int(C.histogram2d_get_size_y((*C.gsl_histogram2d)(unsafe.Pointer(h.Ptr()))))
	return []int{dx, dy}
}

func (h *GslHistogram2d) RangeSliceX(xofs, xlength int) []float64 {
	dim := h.Dim()
	if xlength+xofs > dim[0] {
		gogsl.Error("histogram slice out of range", gogsl.GSL_ERANGE)
		return nil
	}
	ptr := C.histogram2d_get_range_ptr_x((*C.gsl_histogram2d)(unsafe.Pointer(h.Ptr())), C.size_t(xofs))
	srcHdr := &reflect.SliceHeader{Data: uintptr(ptr), Len: xlength, Cap: xlength}
	src := *(*[]float64)(unsafe.Pointer(srcHdr))
	dst := make([]float64, xlength)
	copy(dst, src)
	return dst
}

func (h *GslHistogram2d) RangeSliceX_(xofs, xlength int) []float64 {
	dim := h.Dim()
	if xlength+xofs > dim[0] {
		gogsl.Error("histogram slice out of range", gogsl.GSL_ERANGE)
		return nil
	}
	ptr := C.histogram2d_get_range_ptr_x((*C.gsl_histogram2d)(unsafe.Pointer(h.Ptr())), C.size_t(xofs))
	srcHdr := &reflect.SliceHeader{Data: uintptr(ptr), Len: xlength, Cap: xlength}
	src := *(*[]float64)(unsafe.Pointer(srcHdr))
	return src
}

func (h *GslHistogram2d) RangeSliceY(yofs, ylength int) []float64 {
	dim := h.Dim()
	if ylength+yofs > dim[1] {
		gogsl.Error("histogram slice out of range", gogsl.GSL_ERANGE)
		return nil
	}
	ptr := C.histogram2d_get_range_ptr_y((*C.gsl_histogram2d)(unsafe.Pointer(h.Ptr())), C.size_t(yofs))
	srcHdr := &reflect.SliceHeader{Data: uintptr(ptr), Len: ylength, Cap: ylength}
	src := *(*[]float64)(unsafe.Pointer(srcHdr))
	dst := make([]float64, ylength)
	copy(dst, src)
	return dst
}

func (h *GslHistogram2d) RangeSliceY_(yofs, ylength int) []float64 {
	dim := h.Dim()
	if ylength+yofs > dim[0] {
		gogsl.Error("histogram slice out of range", gogsl.GSL_ERANGE)
		return nil
	}
	ptr := C.histogram2d_get_range_ptr_y((*C.gsl_histogram2d)(unsafe.Pointer(h.Ptr())), C.size_t(yofs))
	srcHdr := &reflect.SliceHeader{Data: uintptr(ptr), Len: ylength, Cap: ylength}
	src := *(*[]float64)(unsafe.Pointer(srcHdr))
	return src
}

func (h *GslHistogram2d) Slice(xofs, xlength, yofs, ylength int) *matrix.GslMatrix {
	dim := h.Dim()
	if xlength+xofs > dim[0] ||
		ylength+yofs > dim[1] {
		gogsl.Error("histogram slice out of range", gogsl.GSL_ERANGE)
		return nil
	}
	// bin[i * ny + j]
	ptrOfs := xofs*dim[1] + yofs
	ptr := C.histogram2d_get_value_ptr((*C.gsl_histogram2d)(unsafe.Pointer(h.Ptr())), C.size_t(ptrOfs))
	dimlen := dim[0]*dim[1] - ptrOfs
	srcHdr := &reflect.SliceHeader{Data: uintptr(ptr), Len: dimlen, Cap: dimlen}
	src := *(*[]float64)(unsafe.Pointer(srcHdr))
	view := matrix.ViewArrayWithTda(src, xlength, ylength, dim[1])
	dst := matrix.MatrixAlloc(xlength, ylength)
	matrix.Memcpy(dst, view.Matrix())
	return dst
}

func (h *GslHistogram2d) Slice_(xofs, xlength, yofs, ylength int) *matrix.GslMatrixView {
	dim := h.Dim()
	if xlength+xofs > dim[0] ||
		ylength+yofs > dim[1] {
		gogsl.Error("histogram slice out of range", gogsl.GSL_ERANGE)
		return nil
	}
	// bin[i * ny + j]
	ptrOfs := xofs*dim[1] + yofs
	ptr := C.histogram2d_get_value_ptr((*C.gsl_histogram2d)(unsafe.Pointer(h.Ptr())), C.size_t(ptrOfs))
	dimlen := dim[0]*dim[1] - ptrOfs
	srcHdr := &reflect.SliceHeader{Data: uintptr(ptr), Len: dimlen, Cap: dimlen}
	src := *(*[]float64)(unsafe.Pointer(srcHdr))
	view := matrix.ViewArrayWithTda(src, xlength, ylength, dim[1])
	return view
}
