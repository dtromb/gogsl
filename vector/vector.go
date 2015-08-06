//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package vector

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_vector.h>
   #include <unistd.h>
*/
import "C"

import "unsafe"
import "github.com/dtromb/gogsl"
import "os"
import "reflect"
import complex_ "github.com/dtromb/gogsl/complex"

type GslVector struct {
   gogsl.GslReference
}

type GslVectorView struct {
   gogsl.GslReference
   CData []byte
}

type GslVectorConstView struct {
   gogsl.GslReference
   CData []byte
}

type GslVectorComplex struct {
   gogsl.GslReference
}

func VectorAlloc(n int) *GslVector {
   _ref := C.gsl_vector_alloc(C.size_t(n))
   _result := &GslVector{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func VectorCalloc(n int) *GslVector {
   _ref := C.gsl_vector_calloc(C.size_t(n))
   _result := &GslVector{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func (x *GslVector) Dispose() {
   C.gsl_vector_free((*C.gsl_vector)(unsafe.Pointer(x.Ptr())))
}

func Get(v *GslVector, i int) float64 {
   return float64(C.gsl_vector_get((*C.gsl_vector)(unsafe.Pointer(v.Ptr())), C.size_t(i)))
}

func Set(v *GslVector, i int, x float64) {
   C.gsl_vector_set((*C.gsl_vector)(unsafe.Pointer(v.Ptr())), C.size_t(i), C.double(x))
}

func SetAll(v *GslVector, x float64) {
   C.gsl_vector_set_all((*C.gsl_vector)(unsafe.Pointer(v.Ptr())), C.double(x))
}

func SetZero(v *GslVector) {
   C.gsl_vector_set_zero((*C.gsl_vector)(unsafe.Pointer(v.Ptr())))
}

func SetBasis(v *GslVector, i int) int32 {
   return int32(C.gsl_vector_set_basis((*C.gsl_vector)(unsafe.Pointer(v.Ptr())), C.size_t(i)))
}

func Fwrite(stream *os.File, v *GslVector) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _result := int32(C.gsl_vector_fwrite(_file_0, (*C.gsl_vector)(unsafe.Pointer(v.Ptr()))))
   C.fclose(_file_0)
   return _result
}

func Fread(stream *os.File, v *GslVector) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _result := int32(C.gsl_vector_fread(_file_0, (*C.gsl_vector)(unsafe.Pointer(v.Ptr()))))
   C.fclose(_file_0)
   return _result
}

func Fprintf(stream *os.File, v *GslVector, format string) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _string_2 := C.CString(format)
   _result := int32(C.gsl_vector_fprintf(_file_0, (*C.gsl_vector)(unsafe.Pointer(v.Ptr())), _string_2))
   C.fclose(_file_0)
   C.free(unsafe.Pointer(_string_2))
   return _result
}

func Fscanf(stream *os.File, v *GslVector) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _result := int32(C.gsl_vector_fscanf(_file_0, (*C.gsl_vector)(unsafe.Pointer(v.Ptr()))))
   C.fclose(_file_0)
   return _result
}

func Subvector(v *GslVector, offset int, n int) *GslVectorView {
   _ref := C.gsl_vector_subvector((*C.gsl_vector)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(n))
   _result := &GslVectorView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func ConstSubvector(v *GslVector, offset int, n int) *GslVectorConstView {
   _ref := C.gsl_vector_const_subvector((*C.gsl_vector)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(n))
   _result := &GslVectorConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func SubvectorWithStride(v *GslVector, offset int, stride int, n int) *GslVectorView {
   _ref := C.gsl_vector_subvector_with_stride((*C.gsl_vector)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(stride), C.size_t(n))
   _result := &GslVectorView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func ConstSubvectorWithStride(v *GslVector, offset int, stride int, n int) *GslVectorConstView {
   _ref := C.gsl_vector_const_subvector_with_stride((*C.gsl_vector)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(stride), C.size_t(n))
   _result := &GslVectorConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func ComplexReal(v *GslVectorComplex) *GslVectorView {
   _ref := C.gsl_vector_complex_real((*C.gsl_vector_complex)(unsafe.Pointer(v.Ptr())))
   _result := &GslVectorView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func ComplexConstReal(v *GslVectorComplex) *GslVectorConstView {
   _ref := C.gsl_vector_complex_const_real((*C.gsl_vector_complex)(unsafe.Pointer(v.Ptr())))
   _result := &GslVectorConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func ComplexImag(v *GslVectorComplex) *GslVectorView {
   _ref := C.gsl_vector_complex_imag((*C.gsl_vector_complex)(unsafe.Pointer(v.Ptr())))
   _result := &GslVectorView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func ComplexConstImag(v *GslVectorComplex) *GslVectorConstView {
   _ref := C.gsl_vector_complex_const_imag((*C.gsl_vector_complex)(unsafe.Pointer(v.Ptr())))
   _result := &GslVectorConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func Memcpy(dest *GslVector, src *GslVector) int32 {
   return int32(C.gsl_vector_memcpy((*C.gsl_vector)(unsafe.Pointer(dest.Ptr())), (*C.gsl_vector)(unsafe.Pointer(src.Ptr()))))
}

func Swap(v *GslVector, w *GslVector) int32 {
   return int32(C.gsl_vector_swap((*C.gsl_vector)(unsafe.Pointer(v.Ptr())), (*C.gsl_vector)(unsafe.Pointer(w.Ptr()))))
}

func SwapElements(v *GslVector, i int, j int) int32 {
   return int32(C.gsl_vector_swap_elements((*C.gsl_vector)(unsafe.Pointer(v.Ptr())), C.size_t(i), C.size_t(j)))
}

func Reverse(v *GslVector) int32 {
   return int32(C.gsl_vector_reverse((*C.gsl_vector)(unsafe.Pointer(v.Ptr()))))
}

func Add(a *GslVector, b *GslVector) int32 {
   return int32(C.gsl_vector_add((*C.gsl_vector)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector)(unsafe.Pointer(b.Ptr()))))
}

func Sub(a *GslVector, b *GslVector) int32 {
   return int32(C.gsl_vector_sub((*C.gsl_vector)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector)(unsafe.Pointer(b.Ptr()))))
}

func Mul(a *GslVector, b *GslVector) int32 {
   return int32(C.gsl_vector_mul((*C.gsl_vector)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector)(unsafe.Pointer(b.Ptr()))))
}

func Div(a *GslVector, b *GslVector) int32 {
   return int32(C.gsl_vector_div((*C.gsl_vector)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector)(unsafe.Pointer(b.Ptr()))))
}

func Scale(a *GslVector, x float64) int32 {
   return int32(C.gsl_vector_scale((*C.gsl_vector)(unsafe.Pointer(a.Ptr())), C.double(x)))
}

func AddConstant(a *GslVector, x float64) int32 {
   return int32(C.gsl_vector_add_constant((*C.gsl_vector)(unsafe.Pointer(a.Ptr())), C.double(x)))
}

func Max(v *GslVector) float64 {
   return float64(C.gsl_vector_max((*C.gsl_vector)(unsafe.Pointer(v.Ptr()))))
}

func Min(v *GslVector) float64 {
   return float64(C.gsl_vector_min((*C.gsl_vector)(unsafe.Pointer(v.Ptr()))))
}

func Minmax(v *GslVector) (float64, float64) {
   var _outptr_1 C.double
   var _outptr_2 C.double
   C.gsl_vector_minmax((*C.gsl_vector)(unsafe.Pointer(v.Ptr())), &_outptr_1, &_outptr_2)
   return *(*float64)(unsafe.Pointer(&_outptr_1)), *(*float64)(unsafe.Pointer(&_outptr_2))
}

func MaxIndex(v *GslVector) int {
   return int(C.gsl_vector_max_index((*C.gsl_vector)(unsafe.Pointer(v.Ptr()))))
}

func MinIndex(v *GslVector) int {
   return int(C.gsl_vector_min_index((*C.gsl_vector)(unsafe.Pointer(v.Ptr()))))
}

func MinmaxIndex(v *GslVector) (int, int) {
   var _outptr_1 C.size_t
   var _outptr_2 C.size_t
   C.gsl_vector_minmax_index((*C.gsl_vector)(unsafe.Pointer(v.Ptr())), &_outptr_1, &_outptr_2)
   return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2))
}

func Isnull(v *GslVector) int32 {
   return int32(C.gsl_vector_isnull((*C.gsl_vector)(unsafe.Pointer(v.Ptr()))))
}

func Ispos(v *GslVector) int32 {
   return int32(C.gsl_vector_ispos((*C.gsl_vector)(unsafe.Pointer(v.Ptr()))))
}

func Isneg(v *GslVector) int32 {
   return int32(C.gsl_vector_isneg((*C.gsl_vector)(unsafe.Pointer(v.Ptr()))))
}

func Isnonneg(v *GslVector) int32 {
   return int32(C.gsl_vector_isnonneg((*C.gsl_vector)(unsafe.Pointer(v.Ptr()))))
}

func Equal(u *GslVector, v *GslVector) int32 {
   return int32(C.gsl_vector_equal((*C.gsl_vector)(unsafe.Pointer(u.Ptr())), (*C.gsl_vector)(unsafe.Pointer(v.Ptr()))))
}

func ViewArray(base []float64, n int) *GslVectorView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_view_array((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(n))
   _result := &GslVectorView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func ConstViewArray(base []float64, n int) *GslVectorConstView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_const_view_array((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(n))
   _result := &GslVectorConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func ViewArrayWithStride(base []float64, stride int, n int) *GslVectorView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_view_array_with_stride((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n))
   _result := &GslVectorView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func ConstViewArrayWithStride(base []float64, stride int, n int) *GslVectorConstView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_const_view_array_with_stride((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n))
   _result := &GslVectorConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

type GslVectorFloat struct {
   gogsl.GslReference
}

type GslVectorFloatView struct {
   gogsl.GslReference
   CData []byte
}

type GslVectorFloatConstView struct {
   gogsl.GslReference
   CData []byte
}

func VectorFloatAlloc(n int) *GslVectorFloat {
   _ref := C.gsl_vector_float_alloc(C.size_t(n))
   _result := &GslVectorFloat{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func VectorFloatCalloc(n int) *GslVectorFloat {
   _ref := C.gsl_vector_float_calloc(C.size_t(n))
   _result := &GslVectorFloat{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func (x *GslVectorFloat) Dispose() {
   C.gsl_vector_float_free((*C.gsl_vector_float)(unsafe.Pointer(x.Ptr())))
}

func FloatGet(v *GslVectorFloat, i int) float32 {
   return float32(C.gsl_vector_float_get((*C.gsl_vector_float)(unsafe.Pointer(v.Ptr())), C.size_t(i)))
}

func FloatSet(v *GslVectorFloat, i int, x float32) {
   C.gsl_vector_float_set((*C.gsl_vector_float)(unsafe.Pointer(v.Ptr())), C.size_t(i), C.float(x))
}

func FloatSetAll(v *GslVectorFloat, x float32) {
   C.gsl_vector_float_set_all((*C.gsl_vector_float)(unsafe.Pointer(v.Ptr())), C.float(x))
}

func FloatSetZero(v *GslVectorFloat) {
   C.gsl_vector_float_set_zero((*C.gsl_vector_float)(unsafe.Pointer(v.Ptr())))
}

func FloatSetBasis(v *GslVectorFloat, i int) int32 {
   return int32(C.gsl_vector_float_set_basis((*C.gsl_vector_float)(unsafe.Pointer(v.Ptr())), C.size_t(i)))
}

func FloatFwrite(stream *os.File, v *GslVectorFloat) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _result := int32(C.gsl_vector_float_fwrite(_file_0, (*C.gsl_vector_float)(unsafe.Pointer(v.Ptr()))))
   C.fclose(_file_0)
   return _result
}

func FloatFread(stream *os.File, v *GslVectorFloat) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _result := int32(C.gsl_vector_float_fread(_file_0, (*C.gsl_vector_float)(unsafe.Pointer(v.Ptr()))))
   C.fclose(_file_0)
   return _result
}

func FloatFprintf(stream *os.File, v *GslVectorFloat, format string) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _string_2 := C.CString(format)
   _result := int32(C.gsl_vector_float_fprintf(_file_0, (*C.gsl_vector_float)(unsafe.Pointer(v.Ptr())), _string_2))
   C.fclose(_file_0)
   C.free(unsafe.Pointer(_string_2))
   return _result
}

func FloatFscanf(stream *os.File, v *GslVectorFloat) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _result := int32(C.gsl_vector_float_fscanf(_file_0, (*C.gsl_vector_float)(unsafe.Pointer(v.Ptr()))))
   C.fclose(_file_0)
   return _result
}

func FloatSubvector(v *GslVectorFloat, offset int, n int) *GslVectorFloatView {
   _ref := C.gsl_vector_float_subvector((*C.gsl_vector_float)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(n))
   _result := &GslVectorFloatView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func FloatConstSubvector(v *GslVectorFloat, offset int, n int) *GslVectorFloatConstView {
   _ref := C.gsl_vector_float_const_subvector((*C.gsl_vector_float)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(n))
   _result := &GslVectorFloatConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func FloatSubvectorWithStride(v *GslVectorFloat, offset int, stride int, n int) *GslVectorFloatView {
   _ref := C.gsl_vector_float_subvector_with_stride((*C.gsl_vector_float)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(stride), C.size_t(n))
   _result := &GslVectorFloatView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func FloatConstSubvectorWithStride(v *GslVectorFloat, offset int, stride int, n int) *GslVectorFloatConstView {
   _ref := C.gsl_vector_float_const_subvector_with_stride((*C.gsl_vector_float)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(stride), C.size_t(n))
   _result := &GslVectorFloatConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func FloatMemcpy(dest *GslVectorFloat, src *GslVectorFloat) int32 {
   return int32(C.gsl_vector_float_memcpy((*C.gsl_vector_float)(unsafe.Pointer(dest.Ptr())), (*C.gsl_vector_float)(unsafe.Pointer(src.Ptr()))))
}

func FloatSwap(v *GslVectorFloat, w *GslVectorFloat) int32 {
   return int32(C.gsl_vector_float_swap((*C.gsl_vector_float)(unsafe.Pointer(v.Ptr())), (*C.gsl_vector_float)(unsafe.Pointer(w.Ptr()))))
}

func FloatSwapElements(v *GslVectorFloat, i int, j int) int32 {
   return int32(C.gsl_vector_float_swap_elements((*C.gsl_vector_float)(unsafe.Pointer(v.Ptr())), C.size_t(i), C.size_t(j)))
}

func FloatReverse(v *GslVectorFloat) int32 {
   return int32(C.gsl_vector_float_reverse((*C.gsl_vector_float)(unsafe.Pointer(v.Ptr()))))
}

func FloatAdd(a *GslVectorFloat, b *GslVectorFloat) int32 {
   return int32(C.gsl_vector_float_add((*C.gsl_vector_float)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_float)(unsafe.Pointer(b.Ptr()))))
}

func FloatSub(a *GslVectorFloat, b *GslVectorFloat) int32 {
   return int32(C.gsl_vector_float_sub((*C.gsl_vector_float)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_float)(unsafe.Pointer(b.Ptr()))))
}

func FloatMul(a *GslVectorFloat, b *GslVectorFloat) int32 {
   return int32(C.gsl_vector_float_mul((*C.gsl_vector_float)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_float)(unsafe.Pointer(b.Ptr()))))
}

func FloatDiv(a *GslVectorFloat, b *GslVectorFloat) int32 {
   return int32(C.gsl_vector_float_div((*C.gsl_vector_float)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_float)(unsafe.Pointer(b.Ptr()))))
}

func FloatViewArray(base []float32, n int) *GslVectorFloatView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_float_view_array((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(n))
   _result := &GslVectorFloatView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func FloatConstViewArray(base []float32, n int) *GslVectorFloatConstView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_float_const_view_array((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(n))
   _result := &GslVectorFloatConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func FloatViewArrayWithStride(base []float32, stride int, n int) *GslVectorFloatView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_float_view_array_with_stride((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n))
   _result := &GslVectorFloatView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func FloatConstViewArrayWithStride(base []float32, stride int, n int) *GslVectorFloatConstView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_float_const_view_array_with_stride((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n))
   _result := &GslVectorFloatConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func FloatScale(a *GslVectorFloat, x float64) int32 {
   return int32(C.gsl_vector_float_scale((*C.gsl_vector_float)(unsafe.Pointer(a.Ptr())), C.double(x)))
}

func FloatAddConstant(a *GslVectorFloat, x float64) int32 {
   return int32(C.gsl_vector_float_add_constant((*C.gsl_vector_float)(unsafe.Pointer(a.Ptr())), C.double(x)))
}

func FloatMax(v *GslVectorFloat) float32 {
   return float32(C.gsl_vector_float_max((*C.gsl_vector_float)(unsafe.Pointer(v.Ptr()))))
}

func FloatMin(v *GslVectorFloat) float32 {
   return float32(C.gsl_vector_float_min((*C.gsl_vector_float)(unsafe.Pointer(v.Ptr()))))
}

func FloatMinmax(v *GslVectorFloat) (float32, float32) {
   var _outptr_1 C.float
   var _outptr_2 C.float
   C.gsl_vector_float_minmax((*C.gsl_vector_float)(unsafe.Pointer(v.Ptr())), &_outptr_1, &_outptr_2)
   return *(*float32)(unsafe.Pointer(&_outptr_1)), *(*float32)(unsafe.Pointer(&_outptr_2))
}

func FloatMaxIndex(v *GslVectorFloat) int {
   return int(C.gsl_vector_float_max_index((*C.gsl_vector_float)(unsafe.Pointer(v.Ptr()))))
}

func FloatMinIndex(v *GslVectorFloat) int {
   return int(C.gsl_vector_float_min_index((*C.gsl_vector_float)(unsafe.Pointer(v.Ptr()))))
}

func FloatMinmaxIndex(v *GslVectorFloat) (int, int) {
   var _outptr_1 C.size_t
   var _outptr_2 C.size_t
   C.gsl_vector_float_minmax_index((*C.gsl_vector_float)(unsafe.Pointer(v.Ptr())), &_outptr_1, &_outptr_2)
   return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2))
}

func FloatIsnull(v *GslVectorFloat) int32 {
   return int32(C.gsl_vector_float_isnull((*C.gsl_vector_float)(unsafe.Pointer(v.Ptr()))))
}

func FloatIspos(v *GslVectorFloat) int32 {
   return int32(C.gsl_vector_float_ispos((*C.gsl_vector_float)(unsafe.Pointer(v.Ptr()))))
}

func FloatIsneg(v *GslVectorFloat) int32 {
   return int32(C.gsl_vector_float_isneg((*C.gsl_vector_float)(unsafe.Pointer(v.Ptr()))))
}

func FloatIsnonneg(v *GslVectorFloat) int32 {
   return int32(C.gsl_vector_float_isnonneg((*C.gsl_vector_float)(unsafe.Pointer(v.Ptr()))))
}

func FloatEqual(u *GslVectorFloat, v *GslVectorFloat) int32 {
   return int32(C.gsl_vector_float_equal((*C.gsl_vector_float)(unsafe.Pointer(u.Ptr())), (*C.gsl_vector_float)(unsafe.Pointer(v.Ptr()))))
}

type GslVectorInt struct {
   gogsl.GslReference
}

type GslVectorIntView struct {
   gogsl.GslReference
   CData []byte
}

type GslVectorIntConstView struct {
   gogsl.GslReference
   CData []byte
}

func VectorIntAlloc(n int) *GslVectorInt {
   _ref := C.gsl_vector_int_alloc(C.size_t(n))
   _result := &GslVectorInt{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func VectorIntCalloc(n int) *GslVectorInt {
   _ref := C.gsl_vector_int_calloc(C.size_t(n))
   _result := &GslVectorInt{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func (x *GslVectorInt) Dispose() {
   C.gsl_vector_int_free((*C.gsl_vector_int)(unsafe.Pointer(x.Ptr())))
}

func IntGet(v *GslVectorInt, i int) int32 {
   return int32(C.gsl_vector_int_get((*C.gsl_vector_int)(unsafe.Pointer(v.Ptr())), C.size_t(i)))
}

func IntSet(v *GslVectorInt, i int, x int32) {
   C.gsl_vector_int_set((*C.gsl_vector_int)(unsafe.Pointer(v.Ptr())), C.size_t(i), C.int(x))
}

func IntSetAll(v *GslVectorInt, x int32) {
   C.gsl_vector_int_set_all((*C.gsl_vector_int)(unsafe.Pointer(v.Ptr())), C.int(x))
}

func IntSetZero(v *GslVectorInt) {
   C.gsl_vector_int_set_zero((*C.gsl_vector_int)(unsafe.Pointer(v.Ptr())))
}

func IntSetBasis(v *GslVectorInt, i int) int32 {
   return int32(C.gsl_vector_int_set_basis((*C.gsl_vector_int)(unsafe.Pointer(v.Ptr())), C.size_t(i)))
}

func IntFwrite(stream *os.File, v *GslVectorInt) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _result := int32(C.gsl_vector_int_fwrite(_file_0, (*C.gsl_vector_int)(unsafe.Pointer(v.Ptr()))))
   C.fclose(_file_0)
   return _result
}

func IntFread(stream *os.File, v *GslVectorInt) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _result := int32(C.gsl_vector_int_fread(_file_0, (*C.gsl_vector_int)(unsafe.Pointer(v.Ptr()))))
   C.fclose(_file_0)
   return _result
}

func IntFprintf(stream *os.File, v *GslVectorInt, format string) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _string_2 := C.CString(format)
   _result := int32(C.gsl_vector_int_fprintf(_file_0, (*C.gsl_vector_int)(unsafe.Pointer(v.Ptr())), _string_2))
   C.fclose(_file_0)
   C.free(unsafe.Pointer(_string_2))
   return _result
}

func IntFscanf(stream *os.File, v *GslVectorInt) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _result := int32(C.gsl_vector_int_fscanf(_file_0, (*C.gsl_vector_int)(unsafe.Pointer(v.Ptr()))))
   C.fclose(_file_0)
   return _result
}

func IntSubvector(v *GslVectorInt, offset int, n int) *GslVectorIntView {
   _ref := C.gsl_vector_int_subvector((*C.gsl_vector_int)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(n))
   _result := &GslVectorIntView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func IntConstSubvector(v *GslVectorInt, offset int, n int) *GslVectorIntConstView {
   _ref := C.gsl_vector_int_const_subvector((*C.gsl_vector_int)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(n))
   _result := &GslVectorIntConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func IntSubvectorWithStride(v *GslVectorInt, offset int, stride int, n int) *GslVectorIntView {
   _ref := C.gsl_vector_int_subvector_with_stride((*C.gsl_vector_int)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(stride), C.size_t(n))
   _result := &GslVectorIntView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func IntConstSubvectorWithStride(v *GslVectorInt, offset int, stride int, n int) *GslVectorIntConstView {
   _ref := C.gsl_vector_int_const_subvector_with_stride((*C.gsl_vector_int)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(stride), C.size_t(n))
   _result := &GslVectorIntConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func IntMemcpy(dest *GslVectorInt, src *GslVectorInt) int32 {
   return int32(C.gsl_vector_int_memcpy((*C.gsl_vector_int)(unsafe.Pointer(dest.Ptr())), (*C.gsl_vector_int)(unsafe.Pointer(src.Ptr()))))
}

func IntSwap(v *GslVectorInt, w *GslVectorInt) int32 {
   return int32(C.gsl_vector_int_swap((*C.gsl_vector_int)(unsafe.Pointer(v.Ptr())), (*C.gsl_vector_int)(unsafe.Pointer(w.Ptr()))))
}

func IntSwapElements(v *GslVectorInt, i int, j int) int32 {
   return int32(C.gsl_vector_int_swap_elements((*C.gsl_vector_int)(unsafe.Pointer(v.Ptr())), C.size_t(i), C.size_t(j)))
}

func IntReverse(v *GslVectorInt) int32 {
   return int32(C.gsl_vector_int_reverse((*C.gsl_vector_int)(unsafe.Pointer(v.Ptr()))))
}

func IntAdd(a *GslVectorInt, b *GslVectorInt) int32 {
   return int32(C.gsl_vector_int_add((*C.gsl_vector_int)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_int)(unsafe.Pointer(b.Ptr()))))
}

func IntSub(a *GslVectorInt, b *GslVectorInt) int32 {
   return int32(C.gsl_vector_int_sub((*C.gsl_vector_int)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_int)(unsafe.Pointer(b.Ptr()))))
}

func IntMul(a *GslVectorInt, b *GslVectorInt) int32 {
   return int32(C.gsl_vector_int_mul((*C.gsl_vector_int)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_int)(unsafe.Pointer(b.Ptr()))))
}

func IntDiv(a *GslVectorInt, b *GslVectorInt) int32 {
   return int32(C.gsl_vector_int_div((*C.gsl_vector_int)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_int)(unsafe.Pointer(b.Ptr()))))
}

func IntViewArray(base []int32, n int) *GslVectorIntView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_int_view_array((*C.int)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(n))
   _result := &GslVectorIntView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func IntConstViewArray(base []int32, n int) *GslVectorIntConstView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_int_const_view_array((*C.int)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(n))
   _result := &GslVectorIntConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func IntViewArrayWithStride(base []int32, stride int, n int) *GslVectorIntView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_int_view_array_with_stride((*C.int)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n))
   _result := &GslVectorIntView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func IntConstViewArrayWithStride(base []int32, stride int, n int) *GslVectorIntConstView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_int_const_view_array_with_stride((*C.int)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n))
   _result := &GslVectorIntConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func IntScale(a *GslVectorInt, x float64) int32 {
   return int32(C.gsl_vector_int_scale((*C.gsl_vector_int)(unsafe.Pointer(a.Ptr())), C.double(x)))
}

func IntAddConstant(a *GslVectorInt, x float64) int32 {
   return int32(C.gsl_vector_int_add_constant((*C.gsl_vector_int)(unsafe.Pointer(a.Ptr())), C.double(x)))
}

func IntMax(v *GslVectorInt) int32 {
   return int32(C.gsl_vector_int_max((*C.gsl_vector_int)(unsafe.Pointer(v.Ptr()))))
}

func IntMin(v *GslVectorInt) int32 {
   return int32(C.gsl_vector_int_min((*C.gsl_vector_int)(unsafe.Pointer(v.Ptr()))))
}

func IntMinmax(v *GslVectorInt) (int32, int32) {
   var _outptr_1 C.int
   var _outptr_2 C.int
   C.gsl_vector_int_minmax((*C.gsl_vector_int)(unsafe.Pointer(v.Ptr())), &_outptr_1, &_outptr_2)
   return *(*int32)(unsafe.Pointer(&_outptr_1)), *(*int32)(unsafe.Pointer(&_outptr_2))
}

func IntMaxIndex(v *GslVectorInt) int {
   return int(C.gsl_vector_int_max_index((*C.gsl_vector_int)(unsafe.Pointer(v.Ptr()))))
}

func IntMinIndex(v *GslVectorInt) int {
   return int(C.gsl_vector_int_min_index((*C.gsl_vector_int)(unsafe.Pointer(v.Ptr()))))
}

func IntMinmaxIndex(v *GslVectorInt) (int, int) {
   var _outptr_1 C.size_t
   var _outptr_2 C.size_t
   C.gsl_vector_int_minmax_index((*C.gsl_vector_int)(unsafe.Pointer(v.Ptr())), &_outptr_1, &_outptr_2)
   return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2))
}

func IntIsnull(v *GslVectorInt) int32 {
   return int32(C.gsl_vector_int_isnull((*C.gsl_vector_int)(unsafe.Pointer(v.Ptr()))))
}

func IntIspos(v *GslVectorInt) int32 {
   return int32(C.gsl_vector_int_ispos((*C.gsl_vector_int)(unsafe.Pointer(v.Ptr()))))
}

func IntIsneg(v *GslVectorInt) int32 {
   return int32(C.gsl_vector_int_isneg((*C.gsl_vector_int)(unsafe.Pointer(v.Ptr()))))
}

func IntIsnonneg(v *GslVectorInt) int32 {
   return int32(C.gsl_vector_int_isnonneg((*C.gsl_vector_int)(unsafe.Pointer(v.Ptr()))))
}

func IntEqual(u *GslVectorInt, v *GslVectorInt) int32 {
   return int32(C.gsl_vector_int_equal((*C.gsl_vector_int)(unsafe.Pointer(u.Ptr())), (*C.gsl_vector_int)(unsafe.Pointer(v.Ptr()))))
}

type GslVectorUint struct {
   gogsl.GslReference
}

type GslVectorUintView struct {
   gogsl.GslReference
   CData []byte
}

type GslVectorUintConstView struct {
   gogsl.GslReference
   CData []byte
}

func VectorUintAlloc(n int) *GslVectorUint {
   _ref := C.gsl_vector_uint_alloc(C.size_t(n))
   _result := &GslVectorUint{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func VectorUintCalloc(n int) *GslVectorUint {
   _ref := C.gsl_vector_uint_calloc(C.size_t(n))
   _result := &GslVectorUint{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func (x *GslVectorUint) Dispose() {
   C.gsl_vector_uint_free((*C.gsl_vector_uint)(unsafe.Pointer(x.Ptr())))
}

func UintGet(v *GslVectorUint, i int) uint32 {
   return uint32(C.gsl_vector_uint_get((*C.gsl_vector_uint)(unsafe.Pointer(v.Ptr())), C.size_t(i)))
}

func UintSet(v *GslVectorUint, i int, x uint32) {
   C.gsl_vector_uint_set((*C.gsl_vector_uint)(unsafe.Pointer(v.Ptr())), C.size_t(i), C.uint(x))
}

func UintSetAll(v *GslVectorUint, x uint32) {
   C.gsl_vector_uint_set_all((*C.gsl_vector_uint)(unsafe.Pointer(v.Ptr())), C.uint(x))
}

func UintSetZero(v *GslVectorUint) {
   C.gsl_vector_uint_set_zero((*C.gsl_vector_uint)(unsafe.Pointer(v.Ptr())))
}

func UintSetBasis(v *GslVectorUint, i int) int32 {
   return int32(C.gsl_vector_uint_set_basis((*C.gsl_vector_uint)(unsafe.Pointer(v.Ptr())), C.size_t(i)))
}

func UintFwrite(stream *os.File, v *GslVectorUint) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _result := int32(C.gsl_vector_uint_fwrite(_file_0, (*C.gsl_vector_uint)(unsafe.Pointer(v.Ptr()))))
   C.fclose(_file_0)
   return _result
}

func UintFread(stream *os.File, v *GslVectorUint) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _result := int32(C.gsl_vector_uint_fread(_file_0, (*C.gsl_vector_uint)(unsafe.Pointer(v.Ptr()))))
   C.fclose(_file_0)
   return _result
}

func UintFprintf(stream *os.File, v *GslVectorUint, format string) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _string_2 := C.CString(format)
   _result := int32(C.gsl_vector_uint_fprintf(_file_0, (*C.gsl_vector_uint)(unsafe.Pointer(v.Ptr())), _string_2))
   C.fclose(_file_0)
   C.free(unsafe.Pointer(_string_2))
   return _result
}

func UintFscanf(stream *os.File, v *GslVectorUint) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _result := int32(C.gsl_vector_uint_fscanf(_file_0, (*C.gsl_vector_uint)(unsafe.Pointer(v.Ptr()))))
   C.fclose(_file_0)
   return _result
}

func UintSubvector(v *GslVectorUint, offset int, n int) *GslVectorUintView {
   _ref := C.gsl_vector_uint_subvector((*C.gsl_vector_uint)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(n))
   _result := &GslVectorUintView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func UintConstSubvector(v *GslVectorUint, offset int, n int) *GslVectorUintConstView {
   _ref := C.gsl_vector_uint_const_subvector((*C.gsl_vector_uint)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(n))
   _result := &GslVectorUintConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func UintSubvectorWithStride(v *GslVectorUint, offset int, stride int, n int) *GslVectorUintView {
   _ref := C.gsl_vector_uint_subvector_with_stride((*C.gsl_vector_uint)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(stride), C.size_t(n))
   _result := &GslVectorUintView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func UintConstSubvectorWithStride(v *GslVectorUint, offset int, stride int, n int) *GslVectorUintConstView {
   _ref := C.gsl_vector_uint_const_subvector_with_stride((*C.gsl_vector_uint)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(stride), C.size_t(n))
   _result := &GslVectorUintConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func UintMemcpy(dest *GslVectorUint, src *GslVectorUint) int32 {
   return int32(C.gsl_vector_uint_memcpy((*C.gsl_vector_uint)(unsafe.Pointer(dest.Ptr())), (*C.gsl_vector_uint)(unsafe.Pointer(src.Ptr()))))
}

func UintSwap(v *GslVectorUint, w *GslVectorUint) int32 {
   return int32(C.gsl_vector_uint_swap((*C.gsl_vector_uint)(unsafe.Pointer(v.Ptr())), (*C.gsl_vector_uint)(unsafe.Pointer(w.Ptr()))))
}

func UintSwapElements(v *GslVectorUint, i int, j int) int32 {
   return int32(C.gsl_vector_uint_swap_elements((*C.gsl_vector_uint)(unsafe.Pointer(v.Ptr())), C.size_t(i), C.size_t(j)))
}

func UintReverse(v *GslVectorUint) int32 {
   return int32(C.gsl_vector_uint_reverse((*C.gsl_vector_uint)(unsafe.Pointer(v.Ptr()))))
}

func UintAdd(a *GslVectorUint, b *GslVectorUint) int32 {
   return int32(C.gsl_vector_uint_add((*C.gsl_vector_uint)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_uint)(unsafe.Pointer(b.Ptr()))))
}

func UintSub(a *GslVectorUint, b *GslVectorUint) int32 {
   return int32(C.gsl_vector_uint_sub((*C.gsl_vector_uint)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_uint)(unsafe.Pointer(b.Ptr()))))
}

func UintMul(a *GslVectorUint, b *GslVectorUint) int32 {
   return int32(C.gsl_vector_uint_mul((*C.gsl_vector_uint)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_uint)(unsafe.Pointer(b.Ptr()))))
}

func UintDiv(a *GslVectorUint, b *GslVectorUint) int32 {
   return int32(C.gsl_vector_uint_div((*C.gsl_vector_uint)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_uint)(unsafe.Pointer(b.Ptr()))))
}

func UintViewArray(base []uint32, n int) *GslVectorUintView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_uint_view_array((*C.uint)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(n))
   _result := &GslVectorUintView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func UintConstViewArray(base []uint32, n int) *GslVectorUintConstView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_uint_const_view_array((*C.uint)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(n))
   _result := &GslVectorUintConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func UintViewArrayWithStride(base []uint32, stride int, n int) *GslVectorUintView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_uint_view_array_with_stride((*C.uint)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n))
   _result := &GslVectorUintView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func UintConstViewArrayWithStride(base []uint32, stride int, n int) *GslVectorUintConstView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_uint_const_view_array_with_stride((*C.uint)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n))
   _result := &GslVectorUintConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func UintScale(a *GslVectorUint, x float64) int32 {
   return int32(C.gsl_vector_uint_scale((*C.gsl_vector_uint)(unsafe.Pointer(a.Ptr())), C.double(x)))
}

func UintAddConstant(a *GslVectorUint, x float64) int32 {
   return int32(C.gsl_vector_uint_add_constant((*C.gsl_vector_uint)(unsafe.Pointer(a.Ptr())), C.double(x)))
}

func UintMax(v *GslVectorUint) uint32 {
   return uint32(C.gsl_vector_uint_max((*C.gsl_vector_uint)(unsafe.Pointer(v.Ptr()))))
}

func UintMin(v *GslVectorUint) uint32 {
   return uint32(C.gsl_vector_uint_min((*C.gsl_vector_uint)(unsafe.Pointer(v.Ptr()))))
}

func UintMinmax(v *GslVectorUint) (uint32, uint32) {
   var _outptr_1 C.uint
   var _outptr_2 C.uint
   C.gsl_vector_uint_minmax((*C.gsl_vector_uint)(unsafe.Pointer(v.Ptr())), &_outptr_1, &_outptr_2)
   return *(*uint32)(unsafe.Pointer(&_outptr_1)), *(*uint32)(unsafe.Pointer(&_outptr_2))
}

func UintMaxIndex(v *GslVectorUint) int {
   return int(C.gsl_vector_uint_max_index((*C.gsl_vector_uint)(unsafe.Pointer(v.Ptr()))))
}

func UintMinIndex(v *GslVectorUint) int {
   return int(C.gsl_vector_uint_min_index((*C.gsl_vector_uint)(unsafe.Pointer(v.Ptr()))))
}

func UintMinmaxIndex(v *GslVectorUint) (int, int) {
   var _outptr_1 C.size_t
   var _outptr_2 C.size_t
   C.gsl_vector_uint_minmax_index((*C.gsl_vector_uint)(unsafe.Pointer(v.Ptr())), &_outptr_1, &_outptr_2)
   return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2))
}

func UintIsnull(v *GslVectorUint) int32 {
   return int32(C.gsl_vector_uint_isnull((*C.gsl_vector_uint)(unsafe.Pointer(v.Ptr()))))
}

func UintIspos(v *GslVectorUint) int32 {
   return int32(C.gsl_vector_uint_ispos((*C.gsl_vector_uint)(unsafe.Pointer(v.Ptr()))))
}

func UintIsneg(v *GslVectorUint) int32 {
   return int32(C.gsl_vector_uint_isneg((*C.gsl_vector_uint)(unsafe.Pointer(v.Ptr()))))
}

func UintIsnonneg(v *GslVectorUint) int32 {
   return int32(C.gsl_vector_uint_isnonneg((*C.gsl_vector_uint)(unsafe.Pointer(v.Ptr()))))
}

func UintEqual(u *GslVectorUint, v *GslVectorUint) int32 {
   return int32(C.gsl_vector_uint_equal((*C.gsl_vector_uint)(unsafe.Pointer(u.Ptr())), (*C.gsl_vector_uint)(unsafe.Pointer(v.Ptr()))))
}

type GslVectorLong struct {
   gogsl.GslReference
}

type GslVectorLongView struct {
   gogsl.GslReference
   CData []byte
}

type GslVectorLongConstView struct {
   gogsl.GslReference
   CData []byte
}

func VectorLongAlloc(n int) *GslVectorLong {
   _ref := C.gsl_vector_long_alloc(C.size_t(n))
   _result := &GslVectorLong{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func VectorLongCalloc(n int) *GslVectorLong {
   _ref := C.gsl_vector_long_calloc(C.size_t(n))
   _result := &GslVectorLong{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func (x *GslVectorLong) Dispose() {
   C.gsl_vector_long_free((*C.gsl_vector_long)(unsafe.Pointer(x.Ptr())))
}

func LongGet(v *GslVectorLong, i int) uint {
   return uint(C.gsl_vector_long_get((*C.gsl_vector_long)(unsafe.Pointer(v.Ptr())), C.size_t(i)))
}

func LongSet(v *GslVectorLong, i int, x uint) {
   C.gsl_vector_long_set((*C.gsl_vector_long)(unsafe.Pointer(v.Ptr())), C.size_t(i), C.long(x))
}

func LongSetAll(v *GslVectorLong, x uint) {
   C.gsl_vector_long_set_all((*C.gsl_vector_long)(unsafe.Pointer(v.Ptr())), C.long(x))
}

func LongSetZero(v *GslVectorLong) {
   C.gsl_vector_long_set_zero((*C.gsl_vector_long)(unsafe.Pointer(v.Ptr())))
}

func LongSetBasis(v *GslVectorLong, i int) int32 {
   return int32(C.gsl_vector_long_set_basis((*C.gsl_vector_long)(unsafe.Pointer(v.Ptr())), C.size_t(i)))
}

func LongFwrite(stream *os.File, v *GslVectorLong) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _result := int32(C.gsl_vector_long_fwrite(_file_0, (*C.gsl_vector_long)(unsafe.Pointer(v.Ptr()))))
   C.fclose(_file_0)
   return _result
}

func LongFread(stream *os.File, v *GslVectorLong) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _result := int32(C.gsl_vector_long_fread(_file_0, (*C.gsl_vector_long)(unsafe.Pointer(v.Ptr()))))
   C.fclose(_file_0)
   return _result
}

func LongFprintf(stream *os.File, v *GslVectorLong, format string) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _string_2 := C.CString(format)
   _result := int32(C.gsl_vector_long_fprintf(_file_0, (*C.gsl_vector_long)(unsafe.Pointer(v.Ptr())), _string_2))
   C.fclose(_file_0)
   C.free(unsafe.Pointer(_string_2))
   return _result
}

func LongFscanf(stream *os.File, v *GslVectorLong) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _result := int32(C.gsl_vector_long_fscanf(_file_0, (*C.gsl_vector_long)(unsafe.Pointer(v.Ptr()))))
   C.fclose(_file_0)
   return _result
}

func LongSubvector(v *GslVectorLong, offset int, n int) *GslVectorLongView {
   _ref := C.gsl_vector_long_subvector((*C.gsl_vector_long)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(n))
   _result := &GslVectorLongView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func LongConstSubvector(v *GslVectorLong, offset int, n int) *GslVectorLongConstView {
   _ref := C.gsl_vector_long_const_subvector((*C.gsl_vector_long)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(n))
   _result := &GslVectorLongConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func LongSubvectorWithStride(v *GslVectorLong, offset int, stride int, n int) *GslVectorLongView {
   _ref := C.gsl_vector_long_subvector_with_stride((*C.gsl_vector_long)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(stride), C.size_t(n))
   _result := &GslVectorLongView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func LongConstSubvectorWithStride(v *GslVectorLong, offset int, stride int, n int) *GslVectorLongConstView {
   _ref := C.gsl_vector_long_const_subvector_with_stride((*C.gsl_vector_long)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(stride), C.size_t(n))
   _result := &GslVectorLongConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func LongMemcpy(dest *GslVectorLong, src *GslVectorLong) int32 {
   return int32(C.gsl_vector_long_memcpy((*C.gsl_vector_long)(unsafe.Pointer(dest.Ptr())), (*C.gsl_vector_long)(unsafe.Pointer(src.Ptr()))))
}

func LongSwap(v *GslVectorLong, w *GslVectorLong) int32 {
   return int32(C.gsl_vector_long_swap((*C.gsl_vector_long)(unsafe.Pointer(v.Ptr())), (*C.gsl_vector_long)(unsafe.Pointer(w.Ptr()))))
}

func LongSwapElements(v *GslVectorLong, i int, j int) int32 {
   return int32(C.gsl_vector_long_swap_elements((*C.gsl_vector_long)(unsafe.Pointer(v.Ptr())), C.size_t(i), C.size_t(j)))
}

func LongReverse(v *GslVectorLong) int32 {
   return int32(C.gsl_vector_long_reverse((*C.gsl_vector_long)(unsafe.Pointer(v.Ptr()))))
}

func LongAdd(a *GslVectorLong, b *GslVectorLong) int32 {
   return int32(C.gsl_vector_long_add((*C.gsl_vector_long)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_long)(unsafe.Pointer(b.Ptr()))))
}

func LongSub(a *GslVectorLong, b *GslVectorLong) int32 {
   return int32(C.gsl_vector_long_sub((*C.gsl_vector_long)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_long)(unsafe.Pointer(b.Ptr()))))
}

func LongMul(a *GslVectorLong, b *GslVectorLong) int32 {
   return int32(C.gsl_vector_long_mul((*C.gsl_vector_long)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_long)(unsafe.Pointer(b.Ptr()))))
}

func LongDiv(a *GslVectorLong, b *GslVectorLong) int32 {
   return int32(C.gsl_vector_long_div((*C.gsl_vector_long)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_long)(unsafe.Pointer(b.Ptr()))))
}

func LongViewArray(base []uint, n int) *GslVectorLongView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_long_view_array((*C.long)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(n))
   _result := &GslVectorLongView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func LongConstViewArray(base []uint, n int) *GslVectorLongConstView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_long_const_view_array((*C.long)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(n))
   _result := &GslVectorLongConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func LongViewArrayWithStride(base []uint, stride int, n int) *GslVectorLongView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_long_view_array_with_stride((*C.long)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n))
   _result := &GslVectorLongView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func LongConstViewArrayWithStride(base []uint, stride int, n int) *GslVectorLongConstView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_long_const_view_array_with_stride((*C.long)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n))
   _result := &GslVectorLongConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func LongScale(a *GslVectorLong, x float64) int32 {
   return int32(C.gsl_vector_long_scale((*C.gsl_vector_long)(unsafe.Pointer(a.Ptr())), C.double(x)))
}

func LongAddConstant(a *GslVectorLong, x float64) int32 {
   return int32(C.gsl_vector_long_add_constant((*C.gsl_vector_long)(unsafe.Pointer(a.Ptr())), C.double(x)))
}

func LongMax(v *GslVectorLong) uint {
   return uint(C.gsl_vector_long_max((*C.gsl_vector_long)(unsafe.Pointer(v.Ptr()))))
}

func LongMin(v *GslVectorLong) uint {
   return uint(C.gsl_vector_long_min((*C.gsl_vector_long)(unsafe.Pointer(v.Ptr()))))
}

func LongMinmax(v *GslVectorLong) (uint, uint) {
   var _outptr_1 C.long
   var _outptr_2 C.long
   C.gsl_vector_long_minmax((*C.gsl_vector_long)(unsafe.Pointer(v.Ptr())), &_outptr_1, &_outptr_2)
   return *(*uint)(unsafe.Pointer(&_outptr_1)), *(*uint)(unsafe.Pointer(&_outptr_2))
}

func LongMaxIndex(v *GslVectorLong) int {
   return int(C.gsl_vector_long_max_index((*C.gsl_vector_long)(unsafe.Pointer(v.Ptr()))))
}

func LongMinIndex(v *GslVectorLong) int {
   return int(C.gsl_vector_long_min_index((*C.gsl_vector_long)(unsafe.Pointer(v.Ptr()))))
}

func LongMinmaxIndex(v *GslVectorLong) (int, int) {
   var _outptr_1 C.size_t
   var _outptr_2 C.size_t
   C.gsl_vector_long_minmax_index((*C.gsl_vector_long)(unsafe.Pointer(v.Ptr())), &_outptr_1, &_outptr_2)
   return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2))
}

func LongIsnull(v *GslVectorLong) int32 {
   return int32(C.gsl_vector_long_isnull((*C.gsl_vector_long)(unsafe.Pointer(v.Ptr()))))
}

func LongIspos(v *GslVectorLong) int32 {
   return int32(C.gsl_vector_long_ispos((*C.gsl_vector_long)(unsafe.Pointer(v.Ptr()))))
}

func LongIsneg(v *GslVectorLong) int32 {
   return int32(C.gsl_vector_long_isneg((*C.gsl_vector_long)(unsafe.Pointer(v.Ptr()))))
}

func LongIsnonneg(v *GslVectorLong) int32 {
   return int32(C.gsl_vector_long_isnonneg((*C.gsl_vector_long)(unsafe.Pointer(v.Ptr()))))
}

func LongEqual(u *GslVectorLong, v *GslVectorLong) int32 {
   return int32(C.gsl_vector_long_equal((*C.gsl_vector_long)(unsafe.Pointer(u.Ptr())), (*C.gsl_vector_long)(unsafe.Pointer(v.Ptr()))))
}

type GslVectorUlong struct {
   gogsl.GslReference
}

type GslVectorUlongView struct {
   gogsl.GslReference
   CData []byte
}

type GslVectorUlongConstView struct {
   gogsl.GslReference
   CData []byte
}

func VectorUlongAlloc(n int) *GslVectorUlong {
   _ref := C.gsl_vector_ulong_alloc(C.size_t(n))
   _result := &GslVectorUlong{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func VectorUlongCalloc(n int) *GslVectorUlong {
   _ref := C.gsl_vector_ulong_calloc(C.size_t(n))
   _result := &GslVectorUlong{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func (x *GslVectorUlong) Dispose() {
   C.gsl_vector_ulong_free((*C.gsl_vector_ulong)(unsafe.Pointer(x.Ptr())))
}

func UlongGet(v *GslVectorUlong, i int) int {
   return int(C.gsl_vector_ulong_get((*C.gsl_vector_ulong)(unsafe.Pointer(v.Ptr())), C.size_t(i)))
}

func UlongSet(v *GslVectorUlong, i int, x int) {
   C.gsl_vector_ulong_set((*C.gsl_vector_ulong)(unsafe.Pointer(v.Ptr())), C.size_t(i), C.ulong(x))
}

func UlongSetAll(v *GslVectorUlong, x int) {
   C.gsl_vector_ulong_set_all((*C.gsl_vector_ulong)(unsafe.Pointer(v.Ptr())), C.ulong(x))
}

func UlongSetZero(v *GslVectorUlong) {
   C.gsl_vector_ulong_set_zero((*C.gsl_vector_ulong)(unsafe.Pointer(v.Ptr())))
}

func UlongSetBasis(v *GslVectorUlong, i int) int32 {
   return int32(C.gsl_vector_ulong_set_basis((*C.gsl_vector_ulong)(unsafe.Pointer(v.Ptr())), C.size_t(i)))
}

func UlongFwrite(stream *os.File, v *GslVectorUlong) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _result := int32(C.gsl_vector_ulong_fwrite(_file_0, (*C.gsl_vector_ulong)(unsafe.Pointer(v.Ptr()))))
   C.fclose(_file_0)
   return _result
}

func UlongFread(stream *os.File, v *GslVectorUlong) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _result := int32(C.gsl_vector_ulong_fread(_file_0, (*C.gsl_vector_ulong)(unsafe.Pointer(v.Ptr()))))
   C.fclose(_file_0)
   return _result
}

func UlongFprintf(stream *os.File, v *GslVectorUlong, format string) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _string_2 := C.CString(format)
   _result := int32(C.gsl_vector_ulong_fprintf(_file_0, (*C.gsl_vector_ulong)(unsafe.Pointer(v.Ptr())), _string_2))
   C.fclose(_file_0)
   C.free(unsafe.Pointer(_string_2))
   return _result
}

func UlongFscanf(stream *os.File, v *GslVectorUlong) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _result := int32(C.gsl_vector_ulong_fscanf(_file_0, (*C.gsl_vector_ulong)(unsafe.Pointer(v.Ptr()))))
   C.fclose(_file_0)
   return _result
}

func UlongSubvector(v *GslVectorUlong, offset int, n int) *GslVectorUlongView {
   _ref := C.gsl_vector_ulong_subvector((*C.gsl_vector_ulong)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(n))
   _result := &GslVectorUlongView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func UlongConstSubvector(v *GslVectorUlong, offset int, n int) *GslVectorUlongConstView {
   _ref := C.gsl_vector_ulong_const_subvector((*C.gsl_vector_ulong)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(n))
   _result := &GslVectorUlongConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func UlongSubvectorWithStride(v *GslVectorUlong, offset int, stride int, n int) *GslVectorUlongView {
   _ref := C.gsl_vector_ulong_subvector_with_stride((*C.gsl_vector_ulong)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(stride), C.size_t(n))
   _result := &GslVectorUlongView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func UlongConstSubvectorWithStride(v *GslVectorUlong, offset int, stride int, n int) *GslVectorUlongConstView {
   _ref := C.gsl_vector_ulong_const_subvector_with_stride((*C.gsl_vector_ulong)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(stride), C.size_t(n))
   _result := &GslVectorUlongConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func UlongMemcpy(dest *GslVectorUlong, src *GslVectorUlong) int32 {
   return int32(C.gsl_vector_ulong_memcpy((*C.gsl_vector_ulong)(unsafe.Pointer(dest.Ptr())), (*C.gsl_vector_ulong)(unsafe.Pointer(src.Ptr()))))
}

func UlongSwap(v *GslVectorUlong, w *GslVectorUlong) int32 {
   return int32(C.gsl_vector_ulong_swap((*C.gsl_vector_ulong)(unsafe.Pointer(v.Ptr())), (*C.gsl_vector_ulong)(unsafe.Pointer(w.Ptr()))))
}

func UlongSwapElements(v *GslVectorUlong, i int, j int) int32 {
   return int32(C.gsl_vector_ulong_swap_elements((*C.gsl_vector_ulong)(unsafe.Pointer(v.Ptr())), C.size_t(i), C.size_t(j)))
}

func UlongReverse(v *GslVectorUlong) int32 {
   return int32(C.gsl_vector_ulong_reverse((*C.gsl_vector_ulong)(unsafe.Pointer(v.Ptr()))))
}

func UlongAdd(a *GslVectorUlong, b *GslVectorUlong) int32 {
   return int32(C.gsl_vector_ulong_add((*C.gsl_vector_ulong)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_ulong)(unsafe.Pointer(b.Ptr()))))
}

func UlongSub(a *GslVectorUlong, b *GslVectorUlong) int32 {
   return int32(C.gsl_vector_ulong_sub((*C.gsl_vector_ulong)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_ulong)(unsafe.Pointer(b.Ptr()))))
}

func UlongMul(a *GslVectorUlong, b *GslVectorUlong) int32 {
   return int32(C.gsl_vector_ulong_mul((*C.gsl_vector_ulong)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_ulong)(unsafe.Pointer(b.Ptr()))))
}

func UlongDiv(a *GslVectorUlong, b *GslVectorUlong) int32 {
   return int32(C.gsl_vector_ulong_div((*C.gsl_vector_ulong)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_ulong)(unsafe.Pointer(b.Ptr()))))
}

func UlongViewArray(base []int, n int) *GslVectorUlongView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_ulong_view_array((*C.ulong)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(n))
   _result := &GslVectorUlongView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func UlongConstViewArray(base []int, n int) *GslVectorUlongConstView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_ulong_const_view_array((*C.ulong)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(n))
   _result := &GslVectorUlongConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func UlongViewArrayWithStride(base []int, stride int, n int) *GslVectorUlongView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_ulong_view_array_with_stride((*C.ulong)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n))
   _result := &GslVectorUlongView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func UlongConstViewArrayWithStride(base []int, stride int, n int) *GslVectorUlongConstView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_ulong_const_view_array_with_stride((*C.ulong)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n))
   _result := &GslVectorUlongConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func UlongScale(a *GslVectorUlong, x float64) int32 {
   return int32(C.gsl_vector_ulong_scale((*C.gsl_vector_ulong)(unsafe.Pointer(a.Ptr())), C.double(x)))
}

func UlongAddConstant(a *GslVectorUlong, x float64) int32 {
   return int32(C.gsl_vector_ulong_add_constant((*C.gsl_vector_ulong)(unsafe.Pointer(a.Ptr())), C.double(x)))
}

func UlongMax(v *GslVectorUlong) int {
   return int(C.gsl_vector_ulong_max((*C.gsl_vector_ulong)(unsafe.Pointer(v.Ptr()))))
}

func UlongMin(v *GslVectorUlong) int {
   return int(C.gsl_vector_ulong_min((*C.gsl_vector_ulong)(unsafe.Pointer(v.Ptr()))))
}

func UlongMinmax(v *GslVectorUlong) (int, int) {
   var _outptr_1 C.ulong
   var _outptr_2 C.ulong
   C.gsl_vector_ulong_minmax((*C.gsl_vector_ulong)(unsafe.Pointer(v.Ptr())), &_outptr_1, &_outptr_2)
   return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2))
}

func UlongMaxIndex(v *GslVectorUlong) int {
   return int(C.gsl_vector_ulong_max_index((*C.gsl_vector_ulong)(unsafe.Pointer(v.Ptr()))))
}

func UlongMinIndex(v *GslVectorUlong) int {
   return int(C.gsl_vector_ulong_min_index((*C.gsl_vector_ulong)(unsafe.Pointer(v.Ptr()))))
}

func UlongMinmaxIndex(v *GslVectorUlong) (int, int) {
   var _outptr_1 C.size_t
   var _outptr_2 C.size_t
   C.gsl_vector_ulong_minmax_index((*C.gsl_vector_ulong)(unsafe.Pointer(v.Ptr())), &_outptr_1, &_outptr_2)
   return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2))
}

func UlongIsnull(v *GslVectorUlong) int32 {
   return int32(C.gsl_vector_ulong_isnull((*C.gsl_vector_ulong)(unsafe.Pointer(v.Ptr()))))
}

func UlongIspos(v *GslVectorUlong) int32 {
   return int32(C.gsl_vector_ulong_ispos((*C.gsl_vector_ulong)(unsafe.Pointer(v.Ptr()))))
}

func UlongIsneg(v *GslVectorUlong) int32 {
   return int32(C.gsl_vector_ulong_isneg((*C.gsl_vector_ulong)(unsafe.Pointer(v.Ptr()))))
}

func UlongIsnonneg(v *GslVectorUlong) int32 {
   return int32(C.gsl_vector_ulong_isnonneg((*C.gsl_vector_ulong)(unsafe.Pointer(v.Ptr()))))
}

func UlongEqual(u *GslVectorUlong, v *GslVectorUlong) int32 {
   return int32(C.gsl_vector_ulong_equal((*C.gsl_vector_ulong)(unsafe.Pointer(u.Ptr())), (*C.gsl_vector_ulong)(unsafe.Pointer(v.Ptr()))))
}

type GslVectorShort struct {
   gogsl.GslReference
}

type GslVectorShortView struct {
   gogsl.GslReference
   CData []byte
}

type GslVectorShortConstView struct {
   gogsl.GslReference
   CData []byte
}

func VectorShortAlloc(n int) *GslVectorShort {
   _ref := C.gsl_vector_short_alloc(C.size_t(n))
   _result := &GslVectorShort{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func VectorShortCalloc(n int) *GslVectorShort {
   _ref := C.gsl_vector_short_calloc(C.size_t(n))
   _result := &GslVectorShort{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func (x *GslVectorShort) Dispose() {
   C.gsl_vector_short_free((*C.gsl_vector_short)(unsafe.Pointer(x.Ptr())))
}

func ShortGet(v *GslVectorShort, i int) int16 {
   return int16(C.gsl_vector_short_get((*C.gsl_vector_short)(unsafe.Pointer(v.Ptr())), C.size_t(i)))
}

func ShortSet(v *GslVectorShort, i int, x int16) {
   C.gsl_vector_short_set((*C.gsl_vector_short)(unsafe.Pointer(v.Ptr())), C.size_t(i), C.short(x))
}

func ShortSetAll(v *GslVectorShort, x int16) {
   C.gsl_vector_short_set_all((*C.gsl_vector_short)(unsafe.Pointer(v.Ptr())), C.short(x))
}

func ShortSetZero(v *GslVectorShort) {
   C.gsl_vector_short_set_zero((*C.gsl_vector_short)(unsafe.Pointer(v.Ptr())))
}

func ShortSetBasis(v *GslVectorShort, i int) int32 {
   return int32(C.gsl_vector_short_set_basis((*C.gsl_vector_short)(unsafe.Pointer(v.Ptr())), C.size_t(i)))
}

func ShortFwrite(stream *os.File, v *GslVectorShort) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _result := int32(C.gsl_vector_short_fwrite(_file_0, (*C.gsl_vector_short)(unsafe.Pointer(v.Ptr()))))
   C.fclose(_file_0)
   return _result
}

func ShortFread(stream *os.File, v *GslVectorShort) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _result := int32(C.gsl_vector_short_fread(_file_0, (*C.gsl_vector_short)(unsafe.Pointer(v.Ptr()))))
   C.fclose(_file_0)
   return _result
}

func ShortFprintf(stream *os.File, v *GslVectorShort, format string) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _string_2 := C.CString(format)
   _result := int32(C.gsl_vector_short_fprintf(_file_0, (*C.gsl_vector_short)(unsafe.Pointer(v.Ptr())), _string_2))
   C.fclose(_file_0)
   C.free(unsafe.Pointer(_string_2))
   return _result
}

func ShortFscanf(stream *os.File, v *GslVectorShort) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _result := int32(C.gsl_vector_short_fscanf(_file_0, (*C.gsl_vector_short)(unsafe.Pointer(v.Ptr()))))
   C.fclose(_file_0)
   return _result
}

func ShortSubvector(v *GslVectorShort, offset int, n int) *GslVectorShortView {
   _ref := C.gsl_vector_short_subvector((*C.gsl_vector_short)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(n))
   _result := &GslVectorShortView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func ShortConstSubvector(v *GslVectorShort, offset int, n int) *GslVectorShortConstView {
   _ref := C.gsl_vector_short_const_subvector((*C.gsl_vector_short)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(n))
   _result := &GslVectorShortConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func ShortSubvectorWithStride(v *GslVectorShort, offset int, stride int, n int) *GslVectorShortView {
   _ref := C.gsl_vector_short_subvector_with_stride((*C.gsl_vector_short)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(stride), C.size_t(n))
   _result := &GslVectorShortView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func ShortConstSubvectorWithStride(v *GslVectorShort, offset int, stride int, n int) *GslVectorShortConstView {
   _ref := C.gsl_vector_short_const_subvector_with_stride((*C.gsl_vector_short)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(stride), C.size_t(n))
   _result := &GslVectorShortConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func ShortMemcpy(dest *GslVectorShort, src *GslVectorShort) int32 {
   return int32(C.gsl_vector_short_memcpy((*C.gsl_vector_short)(unsafe.Pointer(dest.Ptr())), (*C.gsl_vector_short)(unsafe.Pointer(src.Ptr()))))
}

func ShortSwap(v *GslVectorShort, w *GslVectorShort) int32 {
   return int32(C.gsl_vector_short_swap((*C.gsl_vector_short)(unsafe.Pointer(v.Ptr())), (*C.gsl_vector_short)(unsafe.Pointer(w.Ptr()))))
}

func ShortSwapElements(v *GslVectorShort, i int, j int) int32 {
   return int32(C.gsl_vector_short_swap_elements((*C.gsl_vector_short)(unsafe.Pointer(v.Ptr())), C.size_t(i), C.size_t(j)))
}

func ShortReverse(v *GslVectorShort) int32 {
   return int32(C.gsl_vector_short_reverse((*C.gsl_vector_short)(unsafe.Pointer(v.Ptr()))))
}

func ShortAdd(a *GslVectorShort, b *GslVectorShort) int32 {
   return int32(C.gsl_vector_short_add((*C.gsl_vector_short)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_short)(unsafe.Pointer(b.Ptr()))))
}

func ShortSub(a *GslVectorShort, b *GslVectorShort) int32 {
   return int32(C.gsl_vector_short_sub((*C.gsl_vector_short)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_short)(unsafe.Pointer(b.Ptr()))))
}

func ShortMul(a *GslVectorShort, b *GslVectorShort) int32 {
   return int32(C.gsl_vector_short_mul((*C.gsl_vector_short)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_short)(unsafe.Pointer(b.Ptr()))))
}

func ShortDiv(a *GslVectorShort, b *GslVectorShort) int32 {
   return int32(C.gsl_vector_short_div((*C.gsl_vector_short)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_short)(unsafe.Pointer(b.Ptr()))))
}

func ShortViewArray(base []int16, n int) *GslVectorShortView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_short_view_array((*C.short)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(n))
   _result := &GslVectorShortView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func ShortConstViewArray(base []int16, n int) *GslVectorShortConstView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_short_const_view_array((*C.short)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(n))
   _result := &GslVectorShortConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func ShortViewArrayWithStride(base []int16, stride int, n int) *GslVectorShortView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_short_view_array_with_stride((*C.short)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n))
   _result := &GslVectorShortView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func ShortConstViewArrayWithStride(base []int16, stride int, n int) *GslVectorShortConstView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_short_const_view_array_with_stride((*C.short)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n))
   _result := &GslVectorShortConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func ShortScale(a *GslVectorShort, x float64) int32 {
   return int32(C.gsl_vector_short_scale((*C.gsl_vector_short)(unsafe.Pointer(a.Ptr())), C.double(x)))
}

func ShortAddConstant(a *GslVectorShort, x float64) int32 {
   return int32(C.gsl_vector_short_add_constant((*C.gsl_vector_short)(unsafe.Pointer(a.Ptr())), C.double(x)))
}

func ShortMax(v *GslVectorShort) int16 {
   return int16(C.gsl_vector_short_max((*C.gsl_vector_short)(unsafe.Pointer(v.Ptr()))))
}

func ShortMin(v *GslVectorShort) int16 {
   return int16(C.gsl_vector_short_min((*C.gsl_vector_short)(unsafe.Pointer(v.Ptr()))))
}

func ShortMinmax(v *GslVectorShort) (int16, int16) {
   var _outptr_1 C.short
   var _outptr_2 C.short
   C.gsl_vector_short_minmax((*C.gsl_vector_short)(unsafe.Pointer(v.Ptr())), &_outptr_1, &_outptr_2)
   return *(*int16)(unsafe.Pointer(&_outptr_1)), *(*int16)(unsafe.Pointer(&_outptr_2))
}

func ShortMaxIndex(v *GslVectorShort) int {
   return int(C.gsl_vector_short_max_index((*C.gsl_vector_short)(unsafe.Pointer(v.Ptr()))))
}

func ShortMinIndex(v *GslVectorShort) int {
   return int(C.gsl_vector_short_min_index((*C.gsl_vector_short)(unsafe.Pointer(v.Ptr()))))
}

func ShortMinmaxIndex(v *GslVectorShort) (int, int) {
   var _outptr_1 C.size_t
   var _outptr_2 C.size_t
   C.gsl_vector_short_minmax_index((*C.gsl_vector_short)(unsafe.Pointer(v.Ptr())), &_outptr_1, &_outptr_2)
   return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2))
}

func ShortIsnull(v *GslVectorShort) int32 {
   return int32(C.gsl_vector_short_isnull((*C.gsl_vector_short)(unsafe.Pointer(v.Ptr()))))
}

func ShortIspos(v *GslVectorShort) int32 {
   return int32(C.gsl_vector_short_ispos((*C.gsl_vector_short)(unsafe.Pointer(v.Ptr()))))
}

func ShortIsneg(v *GslVectorShort) int32 {
   return int32(C.gsl_vector_short_isneg((*C.gsl_vector_short)(unsafe.Pointer(v.Ptr()))))
}

func ShortIsnonneg(v *GslVectorShort) int32 {
   return int32(C.gsl_vector_short_isnonneg((*C.gsl_vector_short)(unsafe.Pointer(v.Ptr()))))
}

func ShortEqual(u *GslVectorShort, v *GslVectorShort) int32 {
   return int32(C.gsl_vector_short_equal((*C.gsl_vector_short)(unsafe.Pointer(u.Ptr())), (*C.gsl_vector_short)(unsafe.Pointer(v.Ptr()))))
}

type GslVectorUshort struct {
   gogsl.GslReference
}

type GslVectorUshortView struct {
   gogsl.GslReference
   CData []byte
}

type GslVectorUshortConstView struct {
   gogsl.GslReference
   CData []byte
}

func VectorUshortAlloc(n int) *GslVectorUshort {
   _ref := C.gsl_vector_ushort_alloc(C.size_t(n))
   _result := &GslVectorUshort{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func VectorUshortCalloc(n int) *GslVectorUshort {
   _ref := C.gsl_vector_ushort_calloc(C.size_t(n))
   _result := &GslVectorUshort{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func (x *GslVectorUshort) Dispose() {
   C.gsl_vector_ushort_free((*C.gsl_vector_ushort)(unsafe.Pointer(x.Ptr())))
}

func UshortGet(v *GslVectorUshort, i int) int16 {
   return int16(C.gsl_vector_ushort_get((*C.gsl_vector_ushort)(unsafe.Pointer(v.Ptr())), C.size_t(i)))
}

func UshortSet(v *GslVectorUshort, i int, x int16) {
   C.gsl_vector_ushort_set((*C.gsl_vector_ushort)(unsafe.Pointer(v.Ptr())), C.size_t(i), C.ushort(x))
}

func UshortSetAll(v *GslVectorUshort, x int16) {
   C.gsl_vector_ushort_set_all((*C.gsl_vector_ushort)(unsafe.Pointer(v.Ptr())), C.ushort(x))
}

func UshortSetZero(v *GslVectorUshort) {
   C.gsl_vector_ushort_set_zero((*C.gsl_vector_ushort)(unsafe.Pointer(v.Ptr())))
}

func UshortSetBasis(v *GslVectorUshort, i int) int32 {
   return int32(C.gsl_vector_ushort_set_basis((*C.gsl_vector_ushort)(unsafe.Pointer(v.Ptr())), C.size_t(i)))
}

func UshortFwrite(stream *os.File, v *GslVectorUshort) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _result := int32(C.gsl_vector_ushort_fwrite(_file_0, (*C.gsl_vector_ushort)(unsafe.Pointer(v.Ptr()))))
   C.fclose(_file_0)
   return _result
}

func UshortFread(stream *os.File, v *GslVectorUshort) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _result := int32(C.gsl_vector_ushort_fread(_file_0, (*C.gsl_vector_ushort)(unsafe.Pointer(v.Ptr()))))
   C.fclose(_file_0)
   return _result
}

func UshortFprintf(stream *os.File, v *GslVectorUshort, format string) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _string_2 := C.CString(format)
   _result := int32(C.gsl_vector_ushort_fprintf(_file_0, (*C.gsl_vector_ushort)(unsafe.Pointer(v.Ptr())), _string_2))
   C.fclose(_file_0)
   C.free(unsafe.Pointer(_string_2))
   return _result
}

func UshortFscanf(stream *os.File, v *GslVectorUshort) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _result := int32(C.gsl_vector_ushort_fscanf(_file_0, (*C.gsl_vector_ushort)(unsafe.Pointer(v.Ptr()))))
   C.fclose(_file_0)
   return _result
}

func UshortSubvector(v *GslVectorUshort, offset int, n int) *GslVectorUshortView {
   _ref := C.gsl_vector_ushort_subvector((*C.gsl_vector_ushort)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(n))
   _result := &GslVectorUshortView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func UshortConstSubvector(v *GslVectorUshort, offset int, n int) *GslVectorUshortConstView {
   _ref := C.gsl_vector_ushort_const_subvector((*C.gsl_vector_ushort)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(n))
   _result := &GslVectorUshortConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func UshortSubvectorWithStride(v *GslVectorUshort, offset int, stride int, n int) *GslVectorUshortView {
   _ref := C.gsl_vector_ushort_subvector_with_stride((*C.gsl_vector_ushort)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(stride), C.size_t(n))
   _result := &GslVectorUshortView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func UshortConstSubvectorWithStride(v *GslVectorUshort, offset int, stride int, n int) *GslVectorUshortConstView {
   _ref := C.gsl_vector_ushort_const_subvector_with_stride((*C.gsl_vector_ushort)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(stride), C.size_t(n))
   _result := &GslVectorUshortConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func UshortMemcpy(dest *GslVectorUshort, src *GslVectorUshort) int32 {
   return int32(C.gsl_vector_ushort_memcpy((*C.gsl_vector_ushort)(unsafe.Pointer(dest.Ptr())), (*C.gsl_vector_ushort)(unsafe.Pointer(src.Ptr()))))
}

func UshortSwap(v *GslVectorUshort, w *GslVectorUshort) int32 {
   return int32(C.gsl_vector_ushort_swap((*C.gsl_vector_ushort)(unsafe.Pointer(v.Ptr())), (*C.gsl_vector_ushort)(unsafe.Pointer(w.Ptr()))))
}

func UshortSwapElements(v *GslVectorUshort, i int, j int) int32 {
   return int32(C.gsl_vector_ushort_swap_elements((*C.gsl_vector_ushort)(unsafe.Pointer(v.Ptr())), C.size_t(i), C.size_t(j)))
}

func UshortReverse(v *GslVectorUshort) int32 {
   return int32(C.gsl_vector_ushort_reverse((*C.gsl_vector_ushort)(unsafe.Pointer(v.Ptr()))))
}

func UshortAdd(a *GslVectorUshort, b *GslVectorUshort) int32 {
   return int32(C.gsl_vector_ushort_add((*C.gsl_vector_ushort)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_ushort)(unsafe.Pointer(b.Ptr()))))
}

func UshortSub(a *GslVectorUshort, b *GslVectorUshort) int32 {
   return int32(C.gsl_vector_ushort_sub((*C.gsl_vector_ushort)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_ushort)(unsafe.Pointer(b.Ptr()))))
}

func UshortMul(a *GslVectorUshort, b *GslVectorUshort) int32 {
   return int32(C.gsl_vector_ushort_mul((*C.gsl_vector_ushort)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_ushort)(unsafe.Pointer(b.Ptr()))))
}

func UshortDiv(a *GslVectorUshort, b *GslVectorUshort) int32 {
   return int32(C.gsl_vector_ushort_div((*C.gsl_vector_ushort)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_ushort)(unsafe.Pointer(b.Ptr()))))
}

func UshortViewArray(base []int16, n int) *GslVectorUshortView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_ushort_view_array((*C.ushort)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(n))
   _result := &GslVectorUshortView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func UshortConstViewArray(base []int16, n int) *GslVectorUshortConstView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_ushort_const_view_array((*C.ushort)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(n))
   _result := &GslVectorUshortConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func UshortViewArrayWithStride(base []int16, stride int, n int) *GslVectorUshortView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_ushort_view_array_with_stride((*C.ushort)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n))
   _result := &GslVectorUshortView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func UshortConstViewArrayWithStride(base []int16, stride int, n int) *GslVectorUshortConstView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_ushort_const_view_array_with_stride((*C.ushort)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n))
   _result := &GslVectorUshortConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func UshortScale(a *GslVectorUshort, x float64) int32 {
   return int32(C.gsl_vector_ushort_scale((*C.gsl_vector_ushort)(unsafe.Pointer(a.Ptr())), C.double(x)))
}

func UshortAddConstant(a *GslVectorUshort, x float64) int32 {
   return int32(C.gsl_vector_ushort_add_constant((*C.gsl_vector_ushort)(unsafe.Pointer(a.Ptr())), C.double(x)))
}

func UshortMax(v *GslVectorUshort) int16 {
   return int16(C.gsl_vector_ushort_max((*C.gsl_vector_ushort)(unsafe.Pointer(v.Ptr()))))
}

func UshortMin(v *GslVectorUshort) int16 {
   return int16(C.gsl_vector_ushort_min((*C.gsl_vector_ushort)(unsafe.Pointer(v.Ptr()))))
}

func UshortMinmax(v *GslVectorUshort) (int16, int16) {
   var _outptr_1 C.ushort
   var _outptr_2 C.ushort
   C.gsl_vector_ushort_minmax((*C.gsl_vector_ushort)(unsafe.Pointer(v.Ptr())), &_outptr_1, &_outptr_2)
   return *(*int16)(unsafe.Pointer(&_outptr_1)), *(*int16)(unsafe.Pointer(&_outptr_2))
}

func UshortMaxIndex(v *GslVectorUshort) int {
   return int(C.gsl_vector_ushort_max_index((*C.gsl_vector_ushort)(unsafe.Pointer(v.Ptr()))))
}

func UshortMinIndex(v *GslVectorUshort) int {
   return int(C.gsl_vector_ushort_min_index((*C.gsl_vector_ushort)(unsafe.Pointer(v.Ptr()))))
}

func UshortMinmaxIndex(v *GslVectorUshort) (int, int) {
   var _outptr_1 C.size_t
   var _outptr_2 C.size_t
   C.gsl_vector_ushort_minmax_index((*C.gsl_vector_ushort)(unsafe.Pointer(v.Ptr())), &_outptr_1, &_outptr_2)
   return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2))
}

func UshortIsnull(v *GslVectorUshort) int32 {
   return int32(C.gsl_vector_ushort_isnull((*C.gsl_vector_ushort)(unsafe.Pointer(v.Ptr()))))
}

func UshortIspos(v *GslVectorUshort) int32 {
   return int32(C.gsl_vector_ushort_ispos((*C.gsl_vector_ushort)(unsafe.Pointer(v.Ptr()))))
}

func UshortIsneg(v *GslVectorUshort) int32 {
   return int32(C.gsl_vector_ushort_isneg((*C.gsl_vector_ushort)(unsafe.Pointer(v.Ptr()))))
}

func UshortIsnonneg(v *GslVectorUshort) int32 {
   return int32(C.gsl_vector_ushort_isnonneg((*C.gsl_vector_ushort)(unsafe.Pointer(v.Ptr()))))
}

func UshortEqual(u *GslVectorUshort, v *GslVectorUshort) int32 {
   return int32(C.gsl_vector_ushort_equal((*C.gsl_vector_ushort)(unsafe.Pointer(u.Ptr())), (*C.gsl_vector_ushort)(unsafe.Pointer(v.Ptr()))))
}

type GslVectorChar struct {
   gogsl.GslReference
}

type GslVectorCharView struct {
   gogsl.GslReference
   CData []byte
}

type GslVectorCharConstView struct {
   gogsl.GslReference
   CData []byte
}

func VectorCharAlloc(n int) *GslVectorChar {
   _ref := C.gsl_vector_char_alloc(C.size_t(n))
   _result := &GslVectorChar{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func VectorCharCalloc(n int) *GslVectorChar {
   _ref := C.gsl_vector_char_calloc(C.size_t(n))
   _result := &GslVectorChar{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func (x *GslVectorChar) Dispose() {
   C.gsl_vector_char_free((*C.gsl_vector_char)(unsafe.Pointer(x.Ptr())))
}

func CharGet(v *GslVectorChar, i int) int8 {
   return int8(C.gsl_vector_char_get((*C.gsl_vector_char)(unsafe.Pointer(v.Ptr())), C.size_t(i)))
}

func CharSet(v *GslVectorChar, i int, x int8) {
   C.gsl_vector_char_set((*C.gsl_vector_char)(unsafe.Pointer(v.Ptr())), C.size_t(i), C.char(x))
}

func CharSetAll(v *GslVectorChar, x int8) {
   C.gsl_vector_char_set_all((*C.gsl_vector_char)(unsafe.Pointer(v.Ptr())), C.char(x))
}

func CharSetZero(v *GslVectorChar) {
   C.gsl_vector_char_set_zero((*C.gsl_vector_char)(unsafe.Pointer(v.Ptr())))
}

func CharSetBasis(v *GslVectorChar, i int) int32 {
   return int32(C.gsl_vector_char_set_basis((*C.gsl_vector_char)(unsafe.Pointer(v.Ptr())), C.size_t(i)))
}

func CharFwrite(stream *os.File, v *GslVectorChar) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _result := int32(C.gsl_vector_char_fwrite(_file_0, (*C.gsl_vector_char)(unsafe.Pointer(v.Ptr()))))
   C.fclose(_file_0)
   return _result
}

func CharFread(stream *os.File, v *GslVectorChar) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _result := int32(C.gsl_vector_char_fread(_file_0, (*C.gsl_vector_char)(unsafe.Pointer(v.Ptr()))))
   C.fclose(_file_0)
   return _result
}

func CharFprintf(stream *os.File, v *GslVectorChar, format string) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _string_2 := C.CString(format)
   _result := int32(C.gsl_vector_char_fprintf(_file_0, (*C.gsl_vector_char)(unsafe.Pointer(v.Ptr())), _string_2))
   C.fclose(_file_0)
   C.free(unsafe.Pointer(_string_2))
   return _result
}

func CharFscanf(stream *os.File, v *GslVectorChar) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _result := int32(C.gsl_vector_char_fscanf(_file_0, (*C.gsl_vector_char)(unsafe.Pointer(v.Ptr()))))
   C.fclose(_file_0)
   return _result
}

func CharSubvector(v *GslVectorChar, offset int, n int) *GslVectorCharView {
   _ref := C.gsl_vector_char_subvector((*C.gsl_vector_char)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(n))
   _result := &GslVectorCharView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func CharConstSubvector(v *GslVectorChar, offset int, n int) *GslVectorCharConstView {
   _ref := C.gsl_vector_char_const_subvector((*C.gsl_vector_char)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(n))
   _result := &GslVectorCharConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func CharSubvectorWithStride(v *GslVectorChar, offset int, stride int, n int) *GslVectorCharView {
   _ref := C.gsl_vector_char_subvector_with_stride((*C.gsl_vector_char)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(stride), C.size_t(n))
   _result := &GslVectorCharView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func CharConstSubvectorWithStride(v *GslVectorChar, offset int, stride int, n int) *GslVectorCharConstView {
   _ref := C.gsl_vector_char_const_subvector_with_stride((*C.gsl_vector_char)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(stride), C.size_t(n))
   _result := &GslVectorCharConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func CharMemcpy(dest *GslVectorChar, src *GslVectorChar) int32 {
   return int32(C.gsl_vector_char_memcpy((*C.gsl_vector_char)(unsafe.Pointer(dest.Ptr())), (*C.gsl_vector_char)(unsafe.Pointer(src.Ptr()))))
}

func CharSwap(v *GslVectorChar, w *GslVectorChar) int32 {
   return int32(C.gsl_vector_char_swap((*C.gsl_vector_char)(unsafe.Pointer(v.Ptr())), (*C.gsl_vector_char)(unsafe.Pointer(w.Ptr()))))
}

func CharSwapElements(v *GslVectorChar, i int, j int) int32 {
   return int32(C.gsl_vector_char_swap_elements((*C.gsl_vector_char)(unsafe.Pointer(v.Ptr())), C.size_t(i), C.size_t(j)))
}

func CharReverse(v *GslVectorChar) int32 {
   return int32(C.gsl_vector_char_reverse((*C.gsl_vector_char)(unsafe.Pointer(v.Ptr()))))
}

func CharAdd(a *GslVectorChar, b *GslVectorChar) int32 {
   return int32(C.gsl_vector_char_add((*C.gsl_vector_char)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_char)(unsafe.Pointer(b.Ptr()))))
}

func CharSub(a *GslVectorChar, b *GslVectorChar) int32 {
   return int32(C.gsl_vector_char_sub((*C.gsl_vector_char)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_char)(unsafe.Pointer(b.Ptr()))))
}

func CharMul(a *GslVectorChar, b *GslVectorChar) int32 {
   return int32(C.gsl_vector_char_mul((*C.gsl_vector_char)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_char)(unsafe.Pointer(b.Ptr()))))
}

func CharDiv(a *GslVectorChar, b *GslVectorChar) int32 {
   return int32(C.gsl_vector_char_div((*C.gsl_vector_char)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_char)(unsafe.Pointer(b.Ptr()))))
}

func CharViewArray(base []int8, n int) *GslVectorCharView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_char_view_array((*C.char)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(n))
   _result := &GslVectorCharView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func CharConstViewArray(base []int8, n int) *GslVectorCharConstView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_char_const_view_array((*C.char)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(n))
   _result := &GslVectorCharConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func CharViewArrayWithStride(base []int8, stride int, n int) *GslVectorCharView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_char_view_array_with_stride((*C.char)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n))
   _result := &GslVectorCharView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func CharConstViewArrayWithStride(base []int8, stride int, n int) *GslVectorCharConstView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_char_const_view_array_with_stride((*C.char)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n))
   _result := &GslVectorCharConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func CharScale(a *GslVectorChar, x float64) int32 {
   return int32(C.gsl_vector_char_scale((*C.gsl_vector_char)(unsafe.Pointer(a.Ptr())), C.double(x)))
}

func CharAddConstant(a *GslVectorChar, x float64) int32 {
   return int32(C.gsl_vector_char_add_constant((*C.gsl_vector_char)(unsafe.Pointer(a.Ptr())), C.double(x)))
}

func CharMax(v *GslVectorChar) int8 {
   return int8(C.gsl_vector_char_max((*C.gsl_vector_char)(unsafe.Pointer(v.Ptr()))))
}

func CharMin(v *GslVectorChar) int8 {
   return int8(C.gsl_vector_char_min((*C.gsl_vector_char)(unsafe.Pointer(v.Ptr()))))
}

func CharMinmax(v *GslVectorChar) (int8, int8) {
   var _outptr_1 C.char
   var _outptr_2 C.char
   C.gsl_vector_char_minmax((*C.gsl_vector_char)(unsafe.Pointer(v.Ptr())), &_outptr_1, &_outptr_2)
   return *(*int8)(unsafe.Pointer(&_outptr_1)), *(*int8)(unsafe.Pointer(&_outptr_2))
}

func CharMaxIndex(v *GslVectorChar) int {
   return int(C.gsl_vector_char_max_index((*C.gsl_vector_char)(unsafe.Pointer(v.Ptr()))))
}

func CharMinIndex(v *GslVectorChar) int {
   return int(C.gsl_vector_char_min_index((*C.gsl_vector_char)(unsafe.Pointer(v.Ptr()))))
}

func CharMinmaxIndex(v *GslVectorChar) (int, int) {
   var _outptr_1 C.size_t
   var _outptr_2 C.size_t
   C.gsl_vector_char_minmax_index((*C.gsl_vector_char)(unsafe.Pointer(v.Ptr())), &_outptr_1, &_outptr_2)
   return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2))
}

func CharIsnull(v *GslVectorChar) int32 {
   return int32(C.gsl_vector_char_isnull((*C.gsl_vector_char)(unsafe.Pointer(v.Ptr()))))
}

func CharIspos(v *GslVectorChar) int32 {
   return int32(C.gsl_vector_char_ispos((*C.gsl_vector_char)(unsafe.Pointer(v.Ptr()))))
}

func CharIsneg(v *GslVectorChar) int32 {
   return int32(C.gsl_vector_char_isneg((*C.gsl_vector_char)(unsafe.Pointer(v.Ptr()))))
}

func CharIsnonneg(v *GslVectorChar) int32 {
   return int32(C.gsl_vector_char_isnonneg((*C.gsl_vector_char)(unsafe.Pointer(v.Ptr()))))
}

func CharEqual(u *GslVectorChar, v *GslVectorChar) int32 {
   return int32(C.gsl_vector_char_equal((*C.gsl_vector_char)(unsafe.Pointer(u.Ptr())), (*C.gsl_vector_char)(unsafe.Pointer(v.Ptr()))))
}

type GslVectorUchar struct {
   gogsl.GslReference
}

type GslVectorUcharView struct {
   gogsl.GslReference
   CData []byte
}

type GslVectorUcharConstView struct {
   gogsl.GslReference
   CData []byte
}

func VectorUcharAlloc(n int) *GslVectorUchar {
   _ref := C.gsl_vector_uchar_alloc(C.size_t(n))
   _result := &GslVectorUchar{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func VectorUcharCalloc(n int) *GslVectorUchar {
   _ref := C.gsl_vector_uchar_calloc(C.size_t(n))
   _result := &GslVectorUchar{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func (x *GslVectorUchar) Dispose() {
   C.gsl_vector_uchar_free((*C.gsl_vector_uchar)(unsafe.Pointer(x.Ptr())))
}

func UcharGet(v *GslVectorUchar, i int) uint8 {
   return uint8(C.gsl_vector_uchar_get((*C.gsl_vector_uchar)(unsafe.Pointer(v.Ptr())), C.size_t(i)))
}

func UcharSet(v *GslVectorUchar, i int, x uint8) {
   C.gsl_vector_uchar_set((*C.gsl_vector_uchar)(unsafe.Pointer(v.Ptr())), C.size_t(i), C.uchar(x))
}

func UcharSetAll(v *GslVectorUchar, x uint8) {
   C.gsl_vector_uchar_set_all((*C.gsl_vector_uchar)(unsafe.Pointer(v.Ptr())), C.uchar(x))
}

func UcharSetZero(v *GslVectorUchar) {
   C.gsl_vector_uchar_set_zero((*C.gsl_vector_uchar)(unsafe.Pointer(v.Ptr())))
}

func UcharSetBasis(v *GslVectorUchar, i int) int32 {
   return int32(C.gsl_vector_uchar_set_basis((*C.gsl_vector_uchar)(unsafe.Pointer(v.Ptr())), C.size_t(i)))
}

func UcharFwrite(stream *os.File, v *GslVectorUchar) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _result := int32(C.gsl_vector_uchar_fwrite(_file_0, (*C.gsl_vector_uchar)(unsafe.Pointer(v.Ptr()))))
   C.fclose(_file_0)
   return _result
}

func UcharFread(stream *os.File, v *GslVectorUchar) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _result := int32(C.gsl_vector_uchar_fread(_file_0, (*C.gsl_vector_uchar)(unsafe.Pointer(v.Ptr()))))
   C.fclose(_file_0)
   return _result
}

func UcharFprintf(stream *os.File, v *GslVectorUchar, format string) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _string_2 := C.CString(format)
   _result := int32(C.gsl_vector_uchar_fprintf(_file_0, (*C.gsl_vector_uchar)(unsafe.Pointer(v.Ptr())), _string_2))
   C.fclose(_file_0)
   C.free(unsafe.Pointer(_string_2))
   return _result
}

func UcharFscanf(stream *os.File, v *GslVectorUchar) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _result := int32(C.gsl_vector_uchar_fscanf(_file_0, (*C.gsl_vector_uchar)(unsafe.Pointer(v.Ptr()))))
   C.fclose(_file_0)
   return _result
}

func UcharSubvector(v *GslVectorUchar, offset int, n int) *GslVectorUcharView {
   _ref := C.gsl_vector_uchar_subvector((*C.gsl_vector_uchar)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(n))
   _result := &GslVectorUcharView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func UcharConstSubvector(v *GslVectorUchar, offset int, n int) *GslVectorUcharConstView {
   _ref := C.gsl_vector_uchar_const_subvector((*C.gsl_vector_uchar)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(n))
   _result := &GslVectorUcharConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func UcharSubvectorWithStride(v *GslVectorUchar, offset int, stride int, n int) *GslVectorUcharView {
   _ref := C.gsl_vector_uchar_subvector_with_stride((*C.gsl_vector_uchar)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(stride), C.size_t(n))
   _result := &GslVectorUcharView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func UcharConstSubvectorWithStride(v *GslVectorUchar, offset int, stride int, n int) *GslVectorUcharConstView {
   _ref := C.gsl_vector_uchar_const_subvector_with_stride((*C.gsl_vector_uchar)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(stride), C.size_t(n))
   _result := &GslVectorUcharConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func UcharMemcpy(dest *GslVectorUchar, src *GslVectorUchar) int32 {
   return int32(C.gsl_vector_uchar_memcpy((*C.gsl_vector_uchar)(unsafe.Pointer(dest.Ptr())), (*C.gsl_vector_uchar)(unsafe.Pointer(src.Ptr()))))
}

func UcharSwap(v *GslVectorUchar, w *GslVectorUchar) int32 {
   return int32(C.gsl_vector_uchar_swap((*C.gsl_vector_uchar)(unsafe.Pointer(v.Ptr())), (*C.gsl_vector_uchar)(unsafe.Pointer(w.Ptr()))))
}

func UcharSwapElements(v *GslVectorUchar, i int, j int) int32 {
   return int32(C.gsl_vector_uchar_swap_elements((*C.gsl_vector_uchar)(unsafe.Pointer(v.Ptr())), C.size_t(i), C.size_t(j)))
}

func UcharReverse(v *GslVectorUchar) int32 {
   return int32(C.gsl_vector_uchar_reverse((*C.gsl_vector_uchar)(unsafe.Pointer(v.Ptr()))))
}

func UcharAdd(a *GslVectorUchar, b *GslVectorUchar) int32 {
   return int32(C.gsl_vector_uchar_add((*C.gsl_vector_uchar)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_uchar)(unsafe.Pointer(b.Ptr()))))
}

func UcharSub(a *GslVectorUchar, b *GslVectorUchar) int32 {
   return int32(C.gsl_vector_uchar_sub((*C.gsl_vector_uchar)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_uchar)(unsafe.Pointer(b.Ptr()))))
}

func UcharMul(a *GslVectorUchar, b *GslVectorUchar) int32 {
   return int32(C.gsl_vector_uchar_mul((*C.gsl_vector_uchar)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_uchar)(unsafe.Pointer(b.Ptr()))))
}

func UcharDiv(a *GslVectorUchar, b *GslVectorUchar) int32 {
   return int32(C.gsl_vector_uchar_div((*C.gsl_vector_uchar)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_uchar)(unsafe.Pointer(b.Ptr()))))
}

func UcharViewArray(base []uint8, n int) *GslVectorUcharView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_uchar_view_array((*C.uchar)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(n))
   _result := &GslVectorUcharView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func UcharConstViewArray(base []uint8, n int) *GslVectorUcharConstView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_uchar_const_view_array((*C.uchar)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(n))
   _result := &GslVectorUcharConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func UcharViewArrayWithStride(base []uint8, stride int, n int) *GslVectorUcharView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_uchar_view_array_with_stride((*C.uchar)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n))
   _result := &GslVectorUcharView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func UcharConstViewArrayWithStride(base []uint8, stride int, n int) *GslVectorUcharConstView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_uchar_const_view_array_with_stride((*C.uchar)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n))
   _result := &GslVectorUcharConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func UcharScale(a *GslVectorUchar, x float64) int32 {
   return int32(C.gsl_vector_uchar_scale((*C.gsl_vector_uchar)(unsafe.Pointer(a.Ptr())), C.double(x)))
}

func UcharAddConstant(a *GslVectorUchar, x float64) int32 {
   return int32(C.gsl_vector_uchar_add_constant((*C.gsl_vector_uchar)(unsafe.Pointer(a.Ptr())), C.double(x)))
}

func UcharMax(v *GslVectorUchar) uint8 {
   return uint8(C.gsl_vector_uchar_max((*C.gsl_vector_uchar)(unsafe.Pointer(v.Ptr()))))
}

func UcharMin(v *GslVectorUchar) uint8 {
   return uint8(C.gsl_vector_uchar_min((*C.gsl_vector_uchar)(unsafe.Pointer(v.Ptr()))))
}

func UcharMinmax(v *GslVectorUchar) (uint8, uint8) {
   var _outptr_1 C.uchar
   var _outptr_2 C.uchar
   C.gsl_vector_uchar_minmax((*C.gsl_vector_uchar)(unsafe.Pointer(v.Ptr())), &_outptr_1, &_outptr_2)
   return *(*uint8)(unsafe.Pointer(&_outptr_1)), *(*uint8)(unsafe.Pointer(&_outptr_2))
}

func UcharMaxIndex(v *GslVectorUchar) int {
   return int(C.gsl_vector_uchar_max_index((*C.gsl_vector_uchar)(unsafe.Pointer(v.Ptr()))))
}

func UcharMinIndex(v *GslVectorUchar) int {
   return int(C.gsl_vector_uchar_min_index((*C.gsl_vector_uchar)(unsafe.Pointer(v.Ptr()))))
}

func UcharMinmaxIndex(v *GslVectorUchar) (int, int) {
   var _outptr_1 C.size_t
   var _outptr_2 C.size_t
   C.gsl_vector_uchar_minmax_index((*C.gsl_vector_uchar)(unsafe.Pointer(v.Ptr())), &_outptr_1, &_outptr_2)
   return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2))
}

func UcharIsnull(v *GslVectorUchar) int32 {
   return int32(C.gsl_vector_uchar_isnull((*C.gsl_vector_uchar)(unsafe.Pointer(v.Ptr()))))
}

func UcharIspos(v *GslVectorUchar) int32 {
   return int32(C.gsl_vector_uchar_ispos((*C.gsl_vector_uchar)(unsafe.Pointer(v.Ptr()))))
}

func UcharIsneg(v *GslVectorUchar) int32 {
   return int32(C.gsl_vector_uchar_isneg((*C.gsl_vector_uchar)(unsafe.Pointer(v.Ptr()))))
}

func UcharIsnonneg(v *GslVectorUchar) int32 {
   return int32(C.gsl_vector_uchar_isnonneg((*C.gsl_vector_uchar)(unsafe.Pointer(v.Ptr()))))
}

func UcharEqual(u *GslVectorUchar, v *GslVectorUchar) int32 {
   return int32(C.gsl_vector_uchar_equal((*C.gsl_vector_uchar)(unsafe.Pointer(u.Ptr())), (*C.gsl_vector_uchar)(unsafe.Pointer(v.Ptr()))))
}

type GslVectorComplexView struct {
   gogsl.GslReference
   CData []byte
}

type GslVectorComplexConstView struct {
   gogsl.GslReference
   CData []byte
}

func VectorComplexAlloc(n int) *GslVectorComplex {
   _ref := C.gsl_vector_complex_alloc(C.size_t(n))
   _result := &GslVectorComplex{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func VectorComplexCalloc(n int) *GslVectorComplex {
   _ref := C.gsl_vector_complex_calloc(C.size_t(n))
   _result := &GslVectorComplex{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func (x *GslVectorComplex) Dispose() {
   C.gsl_vector_complex_free((*C.gsl_vector_complex)(unsafe.Pointer(x.Ptr())))
}

func ComplexGet(v *GslVectorComplex, i int) complex128 {
   _result := C.gsl_vector_complex_get((*C.gsl_vector_complex)(unsafe.Pointer(v.Ptr())), C.size_t(i))
   return complex_.GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func ComplexSet(v *GslVectorComplex, i int, x complex128) {
   _arg_2 := complex_.GoComplexToGsl(x)
   C.gsl_vector_complex_set((*C.gsl_vector_complex)(unsafe.Pointer(v.Ptr())), C.size_t(i), *(*C.gsl_complex)(unsafe.Pointer(_arg_2)))
}

func ComplexSetAll(v *GslVectorComplex, x complex128) {
   _arg_1 := complex_.GoComplexToGsl(x)
   C.gsl_vector_complex_set_all((*C.gsl_vector_complex)(unsafe.Pointer(v.Ptr())), *(*C.gsl_complex)(unsafe.Pointer(_arg_1)))
}

func ComplexSetZero(v *GslVectorComplex) {
   C.gsl_vector_complex_set_zero((*C.gsl_vector_complex)(unsafe.Pointer(v.Ptr())))
}

func ComplexSetBasis(v *GslVectorComplex, i int) int32 {
   return int32(C.gsl_vector_complex_set_basis((*C.gsl_vector_complex)(unsafe.Pointer(v.Ptr())), C.size_t(i)))
}

func ComplexFwrite(stream *os.File, v *GslVectorComplex) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _result := int32(C.gsl_vector_complex_fwrite(_file_0, (*C.gsl_vector_complex)(unsafe.Pointer(v.Ptr()))))
   C.fclose(_file_0)
   return _result
}

func ComplexFread(stream *os.File, v *GslVectorComplex) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _result := int32(C.gsl_vector_complex_fread(_file_0, (*C.gsl_vector_complex)(unsafe.Pointer(v.Ptr()))))
   C.fclose(_file_0)
   return _result
}

func ComplexFprintf(stream *os.File, v *GslVectorComplex, format string) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _string_2 := C.CString(format)
   _result := int32(C.gsl_vector_complex_fprintf(_file_0, (*C.gsl_vector_complex)(unsafe.Pointer(v.Ptr())), _string_2))
   C.fclose(_file_0)
   C.free(unsafe.Pointer(_string_2))
   return _result
}

func ComplexFscanf(stream *os.File, v *GslVectorComplex) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _result := int32(C.gsl_vector_complex_fscanf(_file_0, (*C.gsl_vector_complex)(unsafe.Pointer(v.Ptr()))))
   C.fclose(_file_0)
   return _result
}

func ComplexSubvector(v *GslVectorComplex, offset int, n int) *GslVectorComplexView {
   _ref := C.gsl_vector_complex_subvector((*C.gsl_vector_complex)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(n))
   _result := &GslVectorComplexView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func ComplexConstSubvector(v *GslVectorComplex, offset int, n int) *GslVectorComplexConstView {
   _ref := C.gsl_vector_complex_const_subvector((*C.gsl_vector_complex)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(n))
   _result := &GslVectorComplexConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func ComplexSubvectorWithStride(v *GslVectorComplex, offset int, stride int, n int) *GslVectorComplexView {
   _ref := C.gsl_vector_complex_subvector_with_stride((*C.gsl_vector_complex)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(stride), C.size_t(n))
   _result := &GslVectorComplexView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func ComplexConstSubvectorWithStride(v *GslVectorComplex, offset int, stride int, n int) *GslVectorComplexConstView {
   _ref := C.gsl_vector_complex_const_subvector_with_stride((*C.gsl_vector_complex)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(stride), C.size_t(n))
   _result := &GslVectorComplexConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func ComplexMemcpy(dest *GslVectorComplex, src *GslVectorComplex) int32 {
   return int32(C.gsl_vector_complex_memcpy((*C.gsl_vector_complex)(unsafe.Pointer(dest.Ptr())), (*C.gsl_vector_complex)(unsafe.Pointer(src.Ptr()))))
}

func ComplexSwap(v *GslVectorComplex, w *GslVectorComplex) int32 {
   return int32(C.gsl_vector_complex_swap((*C.gsl_vector_complex)(unsafe.Pointer(v.Ptr())), (*C.gsl_vector_complex)(unsafe.Pointer(w.Ptr()))))
}

func ComplexSwapElements(v *GslVectorComplex, i int, j int) int32 {
   return int32(C.gsl_vector_complex_swap_elements((*C.gsl_vector_complex)(unsafe.Pointer(v.Ptr())), C.size_t(i), C.size_t(j)))
}

func ComplexReverse(v *GslVectorComplex) int32 {
   return int32(C.gsl_vector_complex_reverse((*C.gsl_vector_complex)(unsafe.Pointer(v.Ptr()))))
}

func ComplexAdd(a *GslVectorComplex, b *GslVectorComplex) int32 {
   return int32(C.gsl_vector_complex_add((*C.gsl_vector_complex)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_complex)(unsafe.Pointer(b.Ptr()))))
}

func ComplexSub(a *GslVectorComplex, b *GslVectorComplex) int32 {
   return int32(C.gsl_vector_complex_sub((*C.gsl_vector_complex)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_complex)(unsafe.Pointer(b.Ptr()))))
}

func ComplexMul(a *GslVectorComplex, b *GslVectorComplex) int32 {
   return int32(C.gsl_vector_complex_mul((*C.gsl_vector_complex)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_complex)(unsafe.Pointer(b.Ptr()))))
}

func ComplexDiv(a *GslVectorComplex, b *GslVectorComplex) int32 {
   return int32(C.gsl_vector_complex_div((*C.gsl_vector_complex)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_complex)(unsafe.Pointer(b.Ptr()))))
}

func ComplexIsnull(v []GslVectorComplex) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&v))
   return int32(C.gsl_vector_complex_isnull((*C.gsl_vector_complex)(unsafe.Pointer(_slice_header_0.Data))))
}

func ComplexIspos(v *GslVectorComplex) int32 {
   return int32(C.gsl_vector_complex_ispos((*C.gsl_vector_complex)(unsafe.Pointer(v.Ptr()))))
}

func ComplexIsneg(v *GslVectorComplex) int32 {
   return int32(C.gsl_vector_complex_isneg((*C.gsl_vector_complex)(unsafe.Pointer(v.Ptr()))))
}

func ComplexIsnonneg(v *GslVectorComplex) int32 {
   return int32(C.gsl_vector_complex_isnonneg((*C.gsl_vector_complex)(unsafe.Pointer(v.Ptr()))))
}

func ComplexEqual(u *GslVectorComplex, v *GslVectorComplex) int32 {
   return int32(C.gsl_vector_complex_equal((*C.gsl_vector_complex)(unsafe.Pointer(u.Ptr())), (*C.gsl_vector_complex)(unsafe.Pointer(v.Ptr()))))
}

type GslVectorComplexFloat struct {
   gogsl.GslReference
}

type GslVectorComplexFloatView struct {
   gogsl.GslReference
   CData []byte
}

type GslVectorComplexFloatConstView struct {
   gogsl.GslReference
   CData []byte
}

func VectorComplexFloatAlloc(n int) *GslVectorComplexFloat {
   _ref := C.gsl_vector_complex_float_alloc(C.size_t(n))
   _result := &GslVectorComplexFloat{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func VectorComplexFloatCalloc(n int) *GslVectorComplexFloat {
   _ref := C.gsl_vector_complex_float_calloc(C.size_t(n))
   _result := &GslVectorComplexFloat{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func (x *GslVectorComplexFloat) Dispose() {
   C.gsl_vector_complex_float_free((*C.gsl_vector_complex_float)(unsafe.Pointer(x.Ptr())))
}

func ComplexFloatGet(v *GslVectorComplexFloat, i int) complex64 {
   _result := C.gsl_vector_complex_float_get((*C.gsl_vector_complex_float)(unsafe.Pointer(v.Ptr())), C.size_t(i))
   return complex_.GslComplexFloatToGo(uintptr(unsafe.Pointer(&_result)))
}

func ComplexFloatSet(v *GslVectorComplexFloat, i int, x complex64) {
   _arg_2 := complex_.GoComplexFloatToGsl(x)
   C.gsl_vector_complex_float_set((*C.gsl_vector_complex_float)(unsafe.Pointer(v.Ptr())), C.size_t(i), *(*C.gsl_complex_float)(unsafe.Pointer(_arg_2)))
}

func ComplexFloatSetAll(v *GslVectorComplexFloat, x complex64) {
   _arg_1 := complex_.GoComplexFloatToGsl(x)
   C.gsl_vector_complex_float_set_all((*C.gsl_vector_complex_float)(unsafe.Pointer(v.Ptr())), *(*C.gsl_complex_float)(unsafe.Pointer(_arg_1)))
}

func ComplexFloatSetZero(v *GslVectorComplexFloat) {
   C.gsl_vector_complex_float_set_zero((*C.gsl_vector_complex_float)(unsafe.Pointer(v.Ptr())))
}

func ComplexFloatSetBasis(v *GslVectorComplexFloat, i int) int32 {
   return int32(C.gsl_vector_complex_float_set_basis((*C.gsl_vector_complex_float)(unsafe.Pointer(v.Ptr())), C.size_t(i)))
}

func ComplexFloatFwrite(stream *os.File, v *GslVectorComplexFloat) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _result := int32(C.gsl_vector_complex_float_fwrite(_file_0, (*C.gsl_vector_complex_float)(unsafe.Pointer(v.Ptr()))))
   C.fclose(_file_0)
   return _result
}

func ComplexFloatFread(stream *os.File, v *GslVectorComplexFloat) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _result := int32(C.gsl_vector_complex_float_fread(_file_0, (*C.gsl_vector_complex_float)(unsafe.Pointer(v.Ptr()))))
   C.fclose(_file_0)
   return _result
}

func ComplexFloatFprintf(stream *os.File, v *GslVectorComplexFloat, format string) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _string_2 := C.CString(format)
   _result := int32(C.gsl_vector_complex_float_fprintf(_file_0, (*C.gsl_vector_complex_float)(unsafe.Pointer(v.Ptr())), _string_2))
   C.fclose(_file_0)
   C.free(unsafe.Pointer(_string_2))
   return _result
}

func ComplexFloatFscanf(stream *os.File, v *GslVectorComplexFloat) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _result := int32(C.gsl_vector_complex_float_fscanf(_file_0, (*C.gsl_vector_complex_float)(unsafe.Pointer(v.Ptr()))))
   C.fclose(_file_0)
   return _result
}

func ComplexFloatSubvector(v *GslVectorComplexFloat, offset int, n int) *GslVectorComplexFloatView {
   _ref := C.gsl_vector_complex_float_subvector((*C.gsl_vector_complex_float)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(n))
   _result := &GslVectorComplexFloatView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func ComplexFloatConstSubvector(v *GslVectorComplexFloat, offset int, n int) *GslVectorComplexFloatConstView {
   _ref := C.gsl_vector_complex_float_const_subvector((*C.gsl_vector_complex_float)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(n))
   _result := &GslVectorComplexFloatConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func ComplexFloatSubvectorWithStride(v *GslVectorComplexFloat, offset int, stride int, n int) *GslVectorComplexFloatView {
   _ref := C.gsl_vector_complex_float_subvector_with_stride((*C.gsl_vector_complex_float)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(stride), C.size_t(n))
   _result := &GslVectorComplexFloatView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func ComplexFloatConstSubvectorWithStride(v *GslVectorComplexFloat, offset int, stride int, n int) *GslVectorComplexFloatConstView {
   _ref := C.gsl_vector_complex_float_const_subvector_with_stride((*C.gsl_vector_complex_float)(unsafe.Pointer(v.Ptr())), C.size_t(offset), C.size_t(stride), C.size_t(n))
   _result := &GslVectorComplexFloatConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func ComplexFloatMemcpy(dest *GslVectorComplexFloat, src *GslVectorComplexFloat) int32 {
   return int32(C.gsl_vector_complex_float_memcpy((*C.gsl_vector_complex_float)(unsafe.Pointer(dest.Ptr())), (*C.gsl_vector_complex_float)(unsafe.Pointer(src.Ptr()))))
}

func ComplexFloatSwap(v *GslVectorComplexFloat, w *GslVectorComplexFloat) int32 {
   return int32(C.gsl_vector_complex_float_swap((*C.gsl_vector_complex_float)(unsafe.Pointer(v.Ptr())), (*C.gsl_vector_complex_float)(unsafe.Pointer(w.Ptr()))))
}

func ComplexFloatSwapElements(v *GslVectorComplexFloat, i int, j int) int32 {
   return int32(C.gsl_vector_complex_float_swap_elements((*C.gsl_vector_complex_float)(unsafe.Pointer(v.Ptr())), C.size_t(i), C.size_t(j)))
}

func ComplexFloatReverse(v *GslVectorComplexFloat) int32 {
   return int32(C.gsl_vector_complex_float_reverse((*C.gsl_vector_complex_float)(unsafe.Pointer(v.Ptr()))))
}

func ComplexFloatAdd(a *GslVectorComplexFloat, b *GslVectorComplexFloat) int32 {
   return int32(C.gsl_vector_complex_float_add((*C.gsl_vector_complex_float)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_complex_float)(unsafe.Pointer(b.Ptr()))))
}

func ComplexFloatSub(a *GslVectorComplexFloat, b *GslVectorComplexFloat) int32 {
   return int32(C.gsl_vector_complex_float_sub((*C.gsl_vector_complex_float)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_complex_float)(unsafe.Pointer(b.Ptr()))))
}

func ComplexFloatMul(a *GslVectorComplexFloat, b *GslVectorComplexFloat) int32 {
   return int32(C.gsl_vector_complex_float_mul((*C.gsl_vector_complex_float)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_complex_float)(unsafe.Pointer(b.Ptr()))))
}

func ComplexFloatDiv(a *GslVectorComplexFloat, b *GslVectorComplexFloat) int32 {
   return int32(C.gsl_vector_complex_float_div((*C.gsl_vector_complex_float)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_complex_float)(unsafe.Pointer(b.Ptr()))))
}

func ComplexFloatIsnull(v []GslVectorComplexFloat) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&v))
   return int32(C.gsl_vector_complex_float_isnull((*C.gsl_vector_complex_float)(unsafe.Pointer(_slice_header_0.Data))))
}

func ComplexFloatIspos(v *GslVectorComplexFloat) int32 {
   return int32(C.gsl_vector_complex_float_ispos((*C.gsl_vector_complex_float)(unsafe.Pointer(v.Ptr()))))
}

func ComplexFloatIsneg(v *GslVectorComplexFloat) int32 {
   return int32(C.gsl_vector_complex_float_isneg((*C.gsl_vector_complex_float)(unsafe.Pointer(v.Ptr()))))
}

func ComplexFloatIsnonneg(v *GslVectorComplexFloat) int32 {
   return int32(C.gsl_vector_complex_float_isnonneg((*C.gsl_vector_complex_float)(unsafe.Pointer(v.Ptr()))))
}

func ComplexFloatEqual(u *GslVectorComplexFloat, v *GslVectorComplexFloat) int32 {
   return int32(C.gsl_vector_complex_float_equal((*C.gsl_vector_complex_float)(unsafe.Pointer(u.Ptr())), (*C.gsl_vector_complex_float)(unsafe.Pointer(v.Ptr()))))
}

func ComplexScale(a *GslVectorComplex, x complex128) int32 {
   _arg_1 := complex_.GoComplexToGsl(x)
   return int32(C.gsl_vector_complex_scale((*C.gsl_vector_complex)(unsafe.Pointer(a.Ptr())), *(*C.gsl_complex)(unsafe.Pointer(_arg_1))))
}

func ComplexFloatScale(a *GslVectorComplexFloat, x complex64) int32 {
   _arg_1 := complex_.GoComplexFloatToGsl(x)
   return int32(C.gsl_vector_complex_float_scale((*C.gsl_vector_complex_float)(unsafe.Pointer(a.Ptr())), *(*C.gsl_complex_float)(unsafe.Pointer(_arg_1))))
}

func ComplexAddConstant(a *GslVectorComplex, x complex128) int32 {
   _arg_1 := complex_.GoComplexToGsl(x)
   return int32(C.gsl_vector_complex_add_constant((*C.gsl_vector_complex)(unsafe.Pointer(a.Ptr())), *(*C.gsl_complex)(unsafe.Pointer(_arg_1))))
}

func ComplexFloatAddConstant(a *GslVectorComplexFloat, x complex64) int32 {
   _arg_1 := complex_.GoComplexFloatToGsl(x)
   return int32(C.gsl_vector_complex_float_add_constant((*C.gsl_vector_complex_float)(unsafe.Pointer(a.Ptr())), *(*C.gsl_complex_float)(unsafe.Pointer(_arg_1))))
}

func ComplexViewArray(base []float64, n int) *GslVectorComplexView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_complex_view_array((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(n))
   _result := &GslVectorComplexView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func ComplexConstViewArray(base []float64, n int) *GslVectorComplexConstView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_complex_const_view_array((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(n))
   _result := &GslVectorComplexConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func ComplexViewArrayWithStride(base []float64, stride int, n int) *GslVectorComplexView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_complex_view_array_with_stride((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n))
   _result := &GslVectorComplexView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

func ComplexConstViewArrayWithStride(base []float64, stride int, n int) *GslVectorComplexConstView {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
   _ref := C.gsl_vector_complex_const_view_array_with_stride((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n))
   _result := &GslVectorComplexConstView{}
   _result.CData = make([]byte, unsafe.Sizeof(_ref))
   gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
   return _result
}

