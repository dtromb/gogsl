//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package wavelet

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_wavelet.h>
   #include <gsl/gsl_wavelet2d.h>
*/
import "C"

import "unsafe"
import "github.com/dtromb/gogsl"
import "reflect"
import "github.com/dtromb/gogsl/matrix"

type GslWavelet struct {
   gogsl.GslReference
}

type GslWaveletType struct {
   gogsl.GslReference
   cPtr uintptr
}

type GslWaveletWorkspace struct {
   gogsl.GslReference
}

func WaveletAlloc(t *GslWaveletType, k int) *GslWavelet {
   _ref := C.gsl_wavelet_alloc((*C.gsl_wavelet_type)(unsafe.Pointer(t.Ptr())), C.size_t(k))
   _result := &GslWavelet{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func WaveletName(w *GslWavelet) string {
   return C.GoString(C.gsl_wavelet_name((*C.gsl_wavelet)(unsafe.Pointer(w.Ptr()))))
}

func (x *GslWavelet) Dispose() {
   C.gsl_wavelet_free((*C.gsl_wavelet)(unsafe.Pointer(x.Ptr())))
}

func WaveletWorkspaceAlloc(n int) *GslWaveletWorkspace {
   _ref := C.gsl_wavelet_workspace_alloc(C.size_t(n))
   _result := &GslWaveletWorkspace{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func (x *GslWaveletWorkspace) Dispose() {
   C.gsl_wavelet_workspace_free((*C.gsl_wavelet_workspace)(unsafe.Pointer(x.Ptr())))
}

func WaveletTransform(w *GslWavelet, data []float64, stride int, n int, dir GslWaveletDirection, work *GslWaveletWorkspace) int32 {
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return int32(C.gsl_wavelet_transform((*C.gsl_wavelet)(unsafe.Pointer(w.Ptr())), (*C.double)(unsafe.Pointer(_slice_header_1.Data)), C.size_t(stride), C.size_t(n), C.gsl_wavelet_direction(dir), (*C.gsl_wavelet_workspace)(unsafe.Pointer(work.Ptr()))))
}

func WaveletTransformForward(w *GslWavelet, data []float64, stride int, n int, work *GslWaveletWorkspace) int32 {
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return int32(C.gsl_wavelet_transform_forward((*C.gsl_wavelet)(unsafe.Pointer(w.Ptr())), (*C.double)(unsafe.Pointer(_slice_header_1.Data)), C.size_t(stride), C.size_t(n), (*C.gsl_wavelet_workspace)(unsafe.Pointer(work.Ptr()))))
}

func WaveletTransformInverse(w *GslWavelet, data []float64, stride int, n int, work *GslWaveletWorkspace) int32 {
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return int32(C.gsl_wavelet_transform_inverse((*C.gsl_wavelet)(unsafe.Pointer(w.Ptr())), (*C.double)(unsafe.Pointer(_slice_header_1.Data)), C.size_t(stride), C.size_t(n), (*C.gsl_wavelet_workspace)(unsafe.Pointer(work.Ptr()))))
}

func Wavelet2dTransform(w *GslWavelet, data []float64, tda int, size1 int, size2 int, dir GslWaveletDirection, work *GslWaveletWorkspace) int32 {
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return int32(C.gsl_wavelet2d_transform((*C.gsl_wavelet)(unsafe.Pointer(w.Ptr())), (*C.double)(unsafe.Pointer(_slice_header_1.Data)), C.size_t(tda), C.size_t(size1), C.size_t(size2), C.gsl_wavelet_direction(dir), (*C.gsl_wavelet_workspace)(unsafe.Pointer(work.Ptr()))))
}

func Wavelet2dTransformForward(w *GslWavelet, data []float64, tda int, size1 int, size2 int, work *GslWaveletWorkspace) int32 {
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return int32(C.gsl_wavelet2d_transform_forward((*C.gsl_wavelet)(unsafe.Pointer(w.Ptr())), (*C.double)(unsafe.Pointer(_slice_header_1.Data)), C.size_t(tda), C.size_t(size1), C.size_t(size2), (*C.gsl_wavelet_workspace)(unsafe.Pointer(work.Ptr()))))
}

func Wavelet2dTransformInverse(w *GslWavelet, data []float64, tda int, size1 int, size2 int, work *GslWaveletWorkspace) int32 {
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return int32(C.gsl_wavelet2d_transform_inverse((*C.gsl_wavelet)(unsafe.Pointer(w.Ptr())), (*C.double)(unsafe.Pointer(_slice_header_1.Data)), C.size_t(tda), C.size_t(size1), C.size_t(size2), (*C.gsl_wavelet_workspace)(unsafe.Pointer(work.Ptr()))))
}

func Wavelet2dTransformMatrix(w *GslWavelet, m *matrix.GslMatrix, dir GslWaveletDirection, work *GslWaveletWorkspace) int32 {
   return int32(C.gsl_wavelet2d_transform_matrix((*C.gsl_wavelet)(unsafe.Pointer(w.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(m.Ptr())), C.gsl_wavelet_direction(dir), (*C.gsl_wavelet_workspace)(unsafe.Pointer(work.Ptr()))))
}

func Wavelet2dTransformMatrixForward(w *GslWavelet, m *matrix.GslMatrix, work *GslWaveletWorkspace) int32 {
   return int32(C.gsl_wavelet2d_transform_matrix_forward((*C.gsl_wavelet)(unsafe.Pointer(w.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(m.Ptr())), (*C.gsl_wavelet_workspace)(unsafe.Pointer(work.Ptr()))))
}

func Wavelet2dTransformMatrixInverse(w *GslWavelet, m *matrix.GslMatrix, work *GslWaveletWorkspace) int32 {
   return int32(C.gsl_wavelet2d_transform_matrix_inverse((*C.gsl_wavelet)(unsafe.Pointer(w.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(m.Ptr())), (*C.gsl_wavelet_workspace)(unsafe.Pointer(work.Ptr()))))
}

func Wavelet2dNstransform(w *GslWavelet, data []float64, tda int, size1 int, size2 int, dir GslWaveletDirection, work *GslWaveletWorkspace) int32 {
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return int32(C.gsl_wavelet2d_nstransform((*C.gsl_wavelet)(unsafe.Pointer(w.Ptr())), (*C.double)(unsafe.Pointer(_slice_header_1.Data)), C.size_t(tda), C.size_t(size1), C.size_t(size2), C.gsl_wavelet_direction(dir), (*C.gsl_wavelet_workspace)(unsafe.Pointer(work.Ptr()))))
}

func Wavelet2dNstransformForward(w *GslWavelet, data []float64, tda int, size1 int, size2 int, work *GslWaveletWorkspace) int32 {
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return int32(C.gsl_wavelet2d_nstransform_forward((*C.gsl_wavelet)(unsafe.Pointer(w.Ptr())), (*C.double)(unsafe.Pointer(_slice_header_1.Data)), C.size_t(tda), C.size_t(size1), C.size_t(size2), (*C.gsl_wavelet_workspace)(unsafe.Pointer(work.Ptr()))))
}

func Wavelet2dNstransformInverse(w *GslWavelet, data []float64, tda int, size1 int, size2 int, work *GslWaveletWorkspace) int32 {
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return int32(C.gsl_wavelet2d_nstransform_inverse((*C.gsl_wavelet)(unsafe.Pointer(w.Ptr())), (*C.double)(unsafe.Pointer(_slice_header_1.Data)), C.size_t(tda), C.size_t(size1), C.size_t(size2), (*C.gsl_wavelet_workspace)(unsafe.Pointer(work.Ptr()))))
}

func Wavelet2dNstransformMatrix(w *GslWavelet, m *matrix.GslMatrix, dir GslWaveletDirection, work *GslWaveletWorkspace) int32 {
   return int32(C.gsl_wavelet2d_nstransform_matrix((*C.gsl_wavelet)(unsafe.Pointer(w.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(m.Ptr())), C.gsl_wavelet_direction(dir), (*C.gsl_wavelet_workspace)(unsafe.Pointer(work.Ptr()))))
}

func Wavelet2dNstransformMatrixForward(w *GslWavelet, m *matrix.GslMatrix, work *GslWaveletWorkspace) int32 {
   return int32(C.gsl_wavelet2d_nstransform_matrix_forward((*C.gsl_wavelet)(unsafe.Pointer(w.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(m.Ptr())), (*C.gsl_wavelet_workspace)(unsafe.Pointer(work.Ptr()))))
}

func Wavelet2dNstransformMatrixInverse(w *GslWavelet, m *matrix.GslMatrix, work *GslWaveletWorkspace) int32 {
   return int32(C.gsl_wavelet2d_nstransform_matrix_inverse((*C.gsl_wavelet)(unsafe.Pointer(w.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(m.Ptr())), (*C.gsl_wavelet_workspace)(unsafe.Pointer(work.Ptr()))))
}

