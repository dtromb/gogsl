package matrix

/*

	#include <gsl/gsl_matrix.h>

	gsl_matrix *get_matrix_view_matrix(gsl_matrix_view *view) {
		return &view->matrix;
	}

	size_t get_matrix_size1(gsl_matrix *m) {
		return m->size1;
	}

	size_t get_matrix_size2(gsl_matrix *m) {
		return m->size2;
	}

	size_t get_matrix_view_struct_size() {
		return sizeof(gsl_matrix_view);
	}

*/
import "C"

import (
	"github.com/dtromb/gogsl"
	"unsafe"
)

var GSL_MATRIX_VIEW_STRUCT_SIZE int32 = int32(C.get_matrix_view_struct_size())

func (gmv *GslMatrixView) Dispose() {

}
func (v *GslMatrixView) Data() []byte {
	return v.CData
}

func (gmcv *GslMatrixConstView) Dispose() {

}
func (v *GslMatrixConstView) Data() []byte {
	return v.CData
}

func (gmv *GslMatrixFloatView) Dispose() {

}
func (v *GslMatrixFloatView) Data() []byte {
	return v.CData
}

func (gmcv *GslMatrixFloatConstView) Dispose() {

}
func (v *GslMatrixFloatConstView) Data() []byte {
	return v.CData
}

func (gmv *GslMatrixIntView) Dispose() {

}
func (v *GslMatrixIntView) Data() []byte {
	return v.CData
}

func (gmcv *GslMatrixIntConstView) Dispose() {

}
func (v *GslMatrixIntConstView) Data() []byte {
	return v.CData
}

func (gmv *GslMatrixUintView) Dispose() {

}
func (v *GslMatrixUintView) Data() []byte {
	return v.CData
}

func (gmcv *GslMatrixUintConstView) Dispose() {

}
func (v *GslMatrixUintConstView) Data() []byte {
	return v.CData
}

func (gmv *GslMatrixLongView) Dispose() {

}
func (v *GslMatrixLongView) Data() []byte {
	return v.CData
}

func (gmcv *GslMatrixLongConstView) Dispose() {

}
func (v *GslMatrixLongConstView) Data() []byte {
	return v.CData
}

func (gmv *GslMatrixUlongView) Dispose() {

}
func (v *GslMatrixUlongView) Data() []byte {
	return v.CData
}

func (gmcv *GslMatrixUlongConstView) Dispose() {

}
func (v *GslMatrixUlongConstView) Data() []byte {
	return v.CData
}

func (gmv *GslMatrixShortView) Dispose() {

}
func (v *GslMatrixShortView) Data() []byte {
	return v.CData
}

func (gmcv *GslMatrixShortConstView) Dispose() {

}
func (v *GslMatrixShortConstView) Data() []byte {
	return v.CData
}

func (gmv *GslMatrixUshortView) Dispose() {

}
func (v *GslMatrixUshortView) Data() []byte {
	return v.CData
}

func (gmcv *GslMatrixUshortConstView) Dispose() {

}
func (v *GslMatrixUshortConstView) Data() []byte {
	return v.CData
}

func (gmv *GslMatrixCharView) Dispose() {

}
func (v *GslMatrixCharView) Data() []byte {
	return v.CData
}

func (gmcv *GslMatrixCharConstView) Dispose() {

}
func (v *GslMatrixCharConstView) Data() []byte {
	return v.CData
}

func (gmv *GslMatrixUcharView) Dispose() {

}
func (v *GslMatrixUcharView) Data() []byte {
	return v.CData
}

func (gmcv *GslMatrixUcharConstView) Dispose() {

}
func (v *GslMatrixUcharConstView) Data() []byte {
	return v.CData
}

func (gmv *GslMatrixComplexView) Dispose() {

}
func (v *GslMatrixComplexView) Data() []byte {
	return v.CData
}

func (gmcv *GslMatrixComplexConstView) Dispose() {

}
func (v *GslMatrixComplexConstView) Data() []byte {
	return v.CData
}

func (gmv *GslMatrixComplexFloatView) Dispose() {

}
func (v *GslMatrixComplexFloatView) Data() []byte {
	return v.CData
}

func (gmcv *GslMatrixComplexFloatConstView) Dispose() {

}
func (v *GslMatrixComplexFloatConstView) Data() []byte {
	return v.CData
}

// Unconvinced of the usefulness of these.
//double * gsl_matrix_ptr (gsl_matrix * m, size_t i, size_t j)
//const double * gsl_matrix_const_ptr (const gsl_matrix * m, size_t i, size_t j)

func (m *GslMatrix) Dim() []int {
	ptr := (*C.gsl_matrix)(unsafe.Pointer(m.Ptr()))
	a := C.get_matrix_size1(ptr)
	b := C.get_matrix_size2(ptr)
	return []int{int(a), int(b)}
}

// XXX - Do this with the others...
func (v *GslMatrixView) Matrix() *GslMatrix {
	m := &GslMatrix{}
	ptr := C.get_matrix_view_matrix((*C.gsl_matrix_view)(unsafe.Pointer(v.Ptr())))
	gogsl.InitializeGslStruct(m, uintptr(unsafe.Pointer(ptr)))
	return m
}
