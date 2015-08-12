//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package fft

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_fft.h>
   #include <gsl/gsl_fft_complex.h>
   #include <gsl/gsl_fft_real.h>
   #include <gsl/gsl_fft_halfcomplex.h>
*/
import "C"

import "unsafe"
import "github.com/dtromb/gogsl"
import "reflect"

type GslFftComplexWavetable struct {
	gogsl.GslReference
}

type GslFftComplexWorkspace struct {
	gogsl.GslReference
}

type GslFftRealWavetable struct {
	gogsl.GslReference
}

type GslFftRealWorkspace struct {
	gogsl.GslReference
}

type GslFftHalfcomplexWavetable struct {
	gogsl.GslReference
}

type GslFftHalfcomplexWorkspace struct {
	gogsl.GslReference
}

func ComplexRadix2Forward(data []complex128, stride int, n int) int32 {
	_slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	return int32(C.gsl_fft_complex_radix2_forward((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func ComplexRadix2Transform(data []complex128, stride int, n int, sign Direction) int32 {
	_slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	return int32(C.gsl_fft_complex_radix2_transform((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.gsl_fft_direction(sign)))
}

func ComplexRadix2Backward(data []complex128, stride int, n int) int32 {
	_slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	return int32(C.gsl_fft_complex_radix2_backward((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func ComplexRadix2Inverse(data []complex128, stride int, n int) int32 {
	_slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	return int32(C.gsl_fft_complex_radix2_inverse((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func ComplexRadix2DifForward(data []complex128, stride int, n int) int32 {
	_slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	return int32(C.gsl_fft_complex_radix2_dif_forward((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func ComplexRadix2DifTransform(data []complex128, stride int, n int, sign Direction) int32 {
	_slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	return int32(C.gsl_fft_complex_radix2_dif_transform((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.gsl_fft_direction(sign)))
}

func ComplexRadix2DifBackward(data []complex128, stride int, n int) int32 {
	_slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	return int32(C.gsl_fft_complex_radix2_dif_backward((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func ComplexRadix2DifInverse(data []complex128, stride int, n int) int32 {
	_slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	return int32(C.gsl_fft_complex_radix2_dif_inverse((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func ComplexWavetableAlloc(n int) *GslFftComplexWavetable {
	_ref := C.gsl_fft_complex_wavetable_alloc(C.size_t(n))
	_result := &GslFftComplexWavetable{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslFftComplexWavetable) Dispose() {
	C.gsl_fft_complex_wavetable_free((*C.gsl_fft_complex_wavetable)(unsafe.Pointer(x.Ptr())))
}

func ComplexWorkspaceAlloc(n int) *GslFftComplexWorkspace {
	_ref := C.gsl_fft_complex_workspace_alloc(C.size_t(n))
	_result := &GslFftComplexWorkspace{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslFftComplexWorkspace) Dispose() {
	C.gsl_fft_complex_workspace_free((*C.gsl_fft_complex_workspace)(unsafe.Pointer(x.Ptr())))
}

func ComplexForward(data []complex128, stride int, n int, wavetable *GslFftComplexWavetable, work *GslFftComplexWorkspace) int32 {
	_slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	return int32(C.gsl_fft_complex_forward((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), (*C.gsl_fft_complex_wavetable)(unsafe.Pointer(wavetable.Ptr())), (*C.gsl_fft_complex_workspace)(unsafe.Pointer(work.Ptr()))))
}

func ComplexTransform(data []complex128, stride int, n int, wavetable *GslFftComplexWavetable, work *GslFftComplexWorkspace, sign Direction) int32 {
	_slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	return int32(C.gsl_fft_complex_transform((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), (*C.gsl_fft_complex_wavetable)(unsafe.Pointer(wavetable.Ptr())), (*C.gsl_fft_complex_workspace)(unsafe.Pointer(work.Ptr())), C.gsl_fft_direction(sign)))
}

func ComplexBackward(data []complex128, stride int, n int, wavetable *GslFftComplexWavetable, work *GslFftComplexWorkspace) int32 {
	_slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	return int32(C.gsl_fft_complex_backward((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), (*C.gsl_fft_complex_wavetable)(unsafe.Pointer(wavetable.Ptr())), (*C.gsl_fft_complex_workspace)(unsafe.Pointer(work.Ptr()))))
}

func ComplexInverse(data []complex128, stride int, n int, wavetable *GslFftComplexWavetable, work *GslFftComplexWorkspace) int32 {
	_slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	return int32(C.gsl_fft_complex_inverse((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), (*C.gsl_fft_complex_wavetable)(unsafe.Pointer(wavetable.Ptr())), (*C.gsl_fft_complex_workspace)(unsafe.Pointer(work.Ptr()))))
}

func HalfcomplexRadix2Inverse(data []float64, stride int, n int) int32 {
	_slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	return int32(C.gsl_fft_halfcomplex_radix2_inverse((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func HalfcomplexRadix2Backward(data []float64, stride int, n int) int32 {
	_slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	return int32(C.gsl_fft_halfcomplex_radix2_backward((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func HalfcomplexRadix2Unpack(halfcomplexCoefficient []float64, complexCoefficient []complex128, stride int, n int) int32 {
	_slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&halfcomplexCoefficient))
	_slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&complexCoefficient))
	return int32(C.gsl_fft_halfcomplex_radix2_unpack((*C.double)(unsafe.Pointer(_slice_header_0.Data)), (*C.double)(unsafe.Pointer(_slice_header_1.Data)), C.size_t(stride), C.size_t(n)))
}

func RealWavetableAlloc(n int) *GslFftRealWavetable {
	_ref := C.gsl_fft_real_wavetable_alloc(C.size_t(n))
	_result := &GslFftRealWavetable{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func HalfcomplexWavetableAlloc(n int) *GslFftHalfcomplexWavetable {
	_ref := C.gsl_fft_halfcomplex_wavetable_alloc(C.size_t(n))
	_result := &GslFftHalfcomplexWavetable{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslFftRealWavetable) Dispose() {
	C.gsl_fft_real_wavetable_free((*C.gsl_fft_real_wavetable)(unsafe.Pointer(x.Ptr())))
}

func (x *GslFftHalfcomplexWavetable) Dispose() {
	C.gsl_fft_halfcomplex_wavetable_free((*C.gsl_fft_halfcomplex_wavetable)(unsafe.Pointer(x.Ptr())))
}

func RealWorkspaceAlloc(n int) *GslFftRealWorkspace {
	_ref := C.gsl_fft_real_workspace_alloc(C.size_t(n))
	_result := &GslFftRealWorkspace{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslFftRealWorkspace) Dispose() {
	C.gsl_fft_real_workspace_free((*C.gsl_fft_real_workspace)(unsafe.Pointer(x.Ptr())))
}

func RealTransform(data []float64, stride int, n int, wavetable *GslFftRealWavetable, work *GslFftRealWorkspace) int32 {
	_slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	return int32(C.gsl_fft_real_transform((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), (*C.gsl_fft_real_wavetable)(unsafe.Pointer(wavetable.Ptr())), (*C.gsl_fft_real_workspace)(unsafe.Pointer(work.Ptr()))))
}

func HalfcomplexTransform(data []float64, stride int, n int, wavetable *GslFftHalfcomplexWavetable, work *GslFftRealWorkspace) int32 {
	_slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	return int32(C.gsl_fft_halfcomplex_transform((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), (*C.gsl_fft_halfcomplex_wavetable)(unsafe.Pointer(wavetable.Ptr())), (*C.gsl_fft_real_workspace)(unsafe.Pointer(work.Ptr()))))
}

func RealUnpack(realCoefficient []float64, complexCoefficient []complex128, stride int, n int) int32 {
	_slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&realCoefficient))
	_slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&complexCoefficient))
	return int32(C.gsl_fft_real_unpack((*C.double)(unsafe.Pointer(_slice_header_0.Data)), (*C.double)(unsafe.Pointer(_slice_header_1.Data)), C.size_t(stride), C.size_t(n)))
}

func HalfcomplexUnpack(halfcomplexCoefficient []float64, complexCoefficient []complex128, stride int, n int) int32 {
	_slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&halfcomplexCoefficient))
	_slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&complexCoefficient))
	return int32(C.gsl_fft_halfcomplex_unpack((*C.double)(unsafe.Pointer(_slice_header_0.Data)), (*C.double)(unsafe.Pointer(_slice_header_1.Data)), C.size_t(stride), C.size_t(n)))
}
