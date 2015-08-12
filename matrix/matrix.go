//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package matrix

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_matrix.h>
   #include <unistd.h>
*/
import "C"

import "unsafe"
import "github.com/dtromb/gogsl"
import "os"
import "reflect"
import "github.com/dtromb/gogsl/vector"
import complex_ "github.com/dtromb/gogsl/complex"

type GslMatrix struct {
	gogsl.GslReference
}

type GslMatrixView struct {
	gogsl.GslReference
	CData []byte
}

type GslMatrixConstView struct {
	gogsl.GslReference
	CData []byte
}

func MatrixAlloc(n1 int, n2 int) *GslMatrix {
	_ref := C.gsl_matrix_alloc(C.size_t(n1), C.size_t(n2))
	_result := &GslMatrix{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func MatrixCalloc(n1 int, n2 int) *GslMatrix {
	_ref := C.gsl_matrix_calloc(C.size_t(n1), C.size_t(n2))
	_result := &GslMatrix{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslMatrix) Dispose() {
	C.gsl_matrix_free((*C.gsl_matrix)(unsafe.Pointer(x.Ptr())))
}

func Get(m *GslMatrix, i int, j int) float64 {
	return float64(C.gsl_matrix_get((*C.gsl_matrix)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func Set(m *GslMatrix, i int, j int, x float64) {
	C.gsl_matrix_set((*C.gsl_matrix)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j), C.double(x))
}

func SetAll(m *GslMatrix, x float64) {
	C.gsl_matrix_set_all((*C.gsl_matrix)(unsafe.Pointer(m.Ptr())), C.double(x))
}

func SetZero(m *GslMatrix) {
	C.gsl_matrix_set_zero((*C.gsl_matrix)(unsafe.Pointer(m.Ptr())))
}

func SetIdentity(m *GslMatrix) {
	C.gsl_matrix_set_identity((*C.gsl_matrix)(unsafe.Pointer(m.Ptr())))
}

func Fwrite(stream *os.File, m *GslMatrix) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_matrix_fwrite(_file_0, (*C.gsl_matrix)(unsafe.Pointer(m.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func Fread(stream *os.File, m *GslMatrix) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_matrix_fread(_file_0, (*C.gsl_matrix)(unsafe.Pointer(m.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func Fprintf(stream *os.File, m *GslMatrix, format string) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_string_2 := C.CString(format)
	_result := int32(C.gsl_matrix_fprintf(_file_0, (*C.gsl_matrix)(unsafe.Pointer(m.Ptr())), _string_2))
	C.fclose(_file_0)
	C.free(unsafe.Pointer(_string_2))
	return _result
}

func Fscanf(stream *os.File, m *GslMatrix) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_matrix_fscanf(_file_0, (*C.gsl_matrix)(unsafe.Pointer(m.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func ViewArray(base []float64, n1 int, n2 int) *GslMatrixView {
	_slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
	_ref := C.gsl_matrix_view_array((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ConstViewArray(base []float64, n1 int, n2 int) *GslMatrixConstView {
	_slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
	_ref := C.gsl_matrix_const_view_array((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ViewArrayWithTda(base []float64, n1 int, n2 int, tda int) *GslMatrixView {
	_slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
	_ref := C.gsl_matrix_view_array_with_tda((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(n1), C.size_t(n2), C.size_t(tda))
	_result := &GslMatrixView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ConstViewArrayWithTda(base []float64, n1 int, n2 int, tda int) *GslMatrixConstView {
	_slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&base))
	_ref := C.gsl_matrix_const_view_array_with_tda((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(n1), C.size_t(n2), C.size_t(tda))
	_result := &GslMatrixConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func Submatrix(m *GslMatrix, k1 int, k2 int, n1 int, n2 int) *GslMatrixView {
	_ref := C.gsl_matrix_submatrix((*C.gsl_matrix)(unsafe.Pointer(m.Ptr())), C.size_t(k1), C.size_t(k2), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ConstSubmatrix(m *GslMatrix, k1 int, k2 int, n1 int, n2 int) *GslMatrixConstView {
	_ref := C.gsl_matrix_const_submatrix((*C.gsl_matrix)(unsafe.Pointer(m.Ptr())), C.size_t(k1), C.size_t(k2), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ViewVector(v *vector.GslVector, n1 int, n2 int) *GslMatrixView {
	_ref := C.gsl_matrix_view_vector((*C.gsl_vector)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ConstViewVector(v *vector.GslVector, n1 int, n2 int) *GslMatrixConstView {
	_ref := C.gsl_matrix_const_view_vector((*C.gsl_vector)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ViewVectorWithTda(v *vector.GslVector, n1 int, n2 int, tda int) *GslMatrixView {
	_ref := C.gsl_matrix_view_vector_with_tda((*C.gsl_vector)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2), C.size_t(tda))
	_result := &GslMatrixView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ConstViewVectorWithTda(v *vector.GslVector, n1 int, n2 int, tda int) *GslMatrixConstView {
	_ref := C.gsl_matrix_const_view_vector_with_tda((*C.gsl_vector)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2), C.size_t(tda))
	_result := &GslMatrixConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func Row(m *GslMatrix, i int) *vector.GslVectorView {
	_ref := C.gsl_matrix_row((*C.gsl_matrix)(unsafe.Pointer(m.Ptr())), C.size_t(i))
	_result := &vector.GslVectorView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ConstRow(m *GslMatrix, i int) *vector.GslVectorConstView {
	_ref := C.gsl_matrix_const_row((*C.gsl_matrix)(unsafe.Pointer(m.Ptr())), C.size_t(i))
	_result := &vector.GslVectorConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func Column(m *GslMatrix, j int) *vector.GslVectorView {
	_ref := C.gsl_matrix_column((*C.gsl_matrix)(unsafe.Pointer(m.Ptr())), C.size_t(j))
	_result := &vector.GslVectorView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ConstColumn(m *GslMatrix, j int) *vector.GslVectorConstView {
	_ref := C.gsl_matrix_const_column((*C.gsl_matrix)(unsafe.Pointer(m.Ptr())), C.size_t(j))
	_result := &vector.GslVectorConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func Subrow(m *GslMatrix, i int, offset int, n int) *vector.GslVectorView {
	_ref := C.gsl_matrix_subrow((*C.gsl_matrix)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ConstSubrow(m *GslMatrix, i int, offset int, n int) *vector.GslVectorConstView {
	_ref := C.gsl_matrix_const_subrow((*C.gsl_matrix)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func Subcolumn(m *GslMatrix, j int, offset int, n int) *vector.GslVectorView {
	_ref := C.gsl_matrix_subcolumn((*C.gsl_matrix)(unsafe.Pointer(m.Ptr())), C.size_t(j), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ConstSubcolumn(m *GslMatrix, j int, offset int, n int) *vector.GslVectorConstView {
	_ref := C.gsl_matrix_const_subcolumn((*C.gsl_matrix)(unsafe.Pointer(m.Ptr())), C.size_t(j), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func Diagonal(m *GslMatrix) *vector.GslVectorView {
	_ref := C.gsl_matrix_diagonal((*C.gsl_matrix)(unsafe.Pointer(m.Ptr())))
	_result := &vector.GslVectorView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ConstDiagonal(m *GslMatrix) *vector.GslVectorConstView {
	_ref := C.gsl_matrix_const_diagonal((*C.gsl_matrix)(unsafe.Pointer(m.Ptr())))
	_result := &vector.GslVectorConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func Subdiagonal(m *GslMatrix, k int) *vector.GslVectorView {
	_ref := C.gsl_matrix_subdiagonal((*C.gsl_matrix)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ConstSubdiagonal(m *GslMatrix, k int) *vector.GslVectorConstView {
	_ref := C.gsl_matrix_const_subdiagonal((*C.gsl_matrix)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func Superdiagonal(m *GslMatrix, k int) *vector.GslVectorView {
	_ref := C.gsl_matrix_superdiagonal((*C.gsl_matrix)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ConstSuperdiagonal(m *GslMatrix, k int) *vector.GslVectorConstView {
	_ref := C.gsl_matrix_const_superdiagonal((*C.gsl_matrix)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func Memcpy(dest *GslMatrix, src *GslMatrix) int32 {
	return int32(C.gsl_matrix_memcpy((*C.gsl_matrix)(unsafe.Pointer(dest.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(src.Ptr()))))
}

func Swap(m1 *GslMatrix, m2 *GslMatrix) int32 {
	return int32(C.gsl_matrix_swap((*C.gsl_matrix)(unsafe.Pointer(m1.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(m2.Ptr()))))
}

func GetRow(v *vector.GslVector, m *GslMatrix, i int) int32 {
	return int32(C.gsl_matrix_get_row((*C.gsl_vector)(unsafe.Pointer(v.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(m.Ptr())), C.size_t(i)))
}

func GetCol(v *vector.GslVector, m *GslMatrix, j int) int32 {
	return int32(C.gsl_matrix_get_col((*C.gsl_vector)(unsafe.Pointer(v.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(m.Ptr())), C.size_t(j)))
}

func SetRow(m *GslMatrix, i int, v *vector.GslVector) int32 {
	return int32(C.gsl_matrix_set_row((*C.gsl_matrix)(unsafe.Pointer(m.Ptr())), C.size_t(i), (*C.gsl_vector)(unsafe.Pointer(v.Ptr()))))
}

func SetCol(m *GslMatrix, j int, v *vector.GslVector) int32 {
	return int32(C.gsl_matrix_set_col((*C.gsl_matrix)(unsafe.Pointer(m.Ptr())), C.size_t(j), (*C.gsl_vector)(unsafe.Pointer(v.Ptr()))))
}

func SwapRows(m *GslMatrix, i int, j int) int32 {
	return int32(C.gsl_matrix_swap_rows((*C.gsl_matrix)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func SwapColumns(m *GslMatrix, i int, j int) int32 {
	return int32(C.gsl_matrix_swap_columns((*C.gsl_matrix)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func SwapRowcol(m *GslMatrix, i int, j int) int32 {
	return int32(C.gsl_matrix_swap_rowcol((*C.gsl_matrix)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func TransposeMemcpy(dest *GslMatrix, src *GslMatrix) int32 {
	return int32(C.gsl_matrix_transpose_memcpy((*C.gsl_matrix)(unsafe.Pointer(dest.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(src.Ptr()))))
}

func Transpose(m *GslMatrix) int32 {
	return int32(C.gsl_matrix_transpose((*C.gsl_matrix)(unsafe.Pointer(m.Ptr()))))
}

func Add(a *GslMatrix, b *GslMatrix) int32 {
	return int32(C.gsl_matrix_add((*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(b.Ptr()))))
}

func Sub(a *GslMatrix, b *GslMatrix) int32 {
	return int32(C.gsl_matrix_sub((*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(b.Ptr()))))
}

func MulElements(a *GslMatrix, b *GslMatrix) int32 {
	return int32(C.gsl_matrix_mul_elements((*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(b.Ptr()))))
}

func DivElements(a *GslMatrix, b *GslMatrix) int32 {
	return int32(C.gsl_matrix_div_elements((*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(b.Ptr()))))
}

func Scale(a *GslMatrix, x float64) int32 {
	return int32(C.gsl_matrix_scale((*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), C.double(x)))
}

func AddConstant(a *GslMatrix, x float64) int32 {
	return int32(C.gsl_matrix_add_constant((*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), C.double(x)))
}

func Max(m *GslMatrix) float64 {
	return float64(C.gsl_matrix_max((*C.gsl_matrix)(unsafe.Pointer(m.Ptr()))))
}

func Min(m *GslMatrix) float64 {
	return float64(C.gsl_matrix_min((*C.gsl_matrix)(unsafe.Pointer(m.Ptr()))))
}

func Minmax(m *GslMatrix) (float64, float64) {
	var _outptr_1 C.double
	var _outptr_2 C.double
	C.gsl_matrix_minmax((*C.gsl_matrix)(unsafe.Pointer(m.Ptr())), &_outptr_1, &_outptr_2)
	return *(*float64)(unsafe.Pointer(&_outptr_1)), *(*float64)(unsafe.Pointer(&_outptr_2))
}

func MaxIndex(m *GslMatrix) (int, int) {
	var _outptr_1 C.size_t
	var _outptr_2 C.size_t
	C.gsl_matrix_max_index((*C.gsl_matrix)(unsafe.Pointer(m.Ptr())), &_outptr_1, &_outptr_2)
	return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2))
}

func MinIndex(m *GslMatrix) (int, int) {
	var _outptr_1 C.size_t
	var _outptr_2 C.size_t
	C.gsl_matrix_min_index((*C.gsl_matrix)(unsafe.Pointer(m.Ptr())), &_outptr_1, &_outptr_2)
	return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2))
}

func MinmaxIndex(m *GslMatrix) (int, int, int, int) {
	var _outptr_1 C.size_t
	var _outptr_2 C.size_t
	var _outptr_3 C.size_t
	var _outptr_4 C.size_t
	C.gsl_matrix_minmax_index((*C.gsl_matrix)(unsafe.Pointer(m.Ptr())), &_outptr_1, &_outptr_2, &_outptr_3, &_outptr_4)
	return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2)), *(*int)(unsafe.Pointer(&_outptr_3)), *(*int)(unsafe.Pointer(&_outptr_4))
}

func Isnull(m *GslMatrix) int32 {
	return int32(C.gsl_matrix_isnull((*C.gsl_matrix)(unsafe.Pointer(m.Ptr()))))
}

func Ispos(m *GslMatrix) int32 {
	return int32(C.gsl_matrix_ispos((*C.gsl_matrix)(unsafe.Pointer(m.Ptr()))))
}

func Isneg(m *GslMatrix) int32 {
	return int32(C.gsl_matrix_isneg((*C.gsl_matrix)(unsafe.Pointer(m.Ptr()))))
}

func Isnonneg(m *GslMatrix) int32 {
	return int32(C.gsl_matrix_isnonneg((*C.gsl_matrix)(unsafe.Pointer(m.Ptr()))))
}

func Equal(a *GslMatrix, b *GslMatrix) int32 {
	return int32(C.gsl_matrix_equal((*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(b.Ptr()))))
}

type GslMatrixFloat struct {
	gogsl.GslReference
}

type GslMatrixFloatView struct {
	gogsl.GslReference
	CData []byte
}

type GslMatrixFloatConstView struct {
	gogsl.GslReference
	CData []byte
}

func MatrixFloatAlloc(n1 int, n2 int) *GslMatrixFloat {
	_ref := C.gsl_matrix_float_alloc(C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixFloat{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func MatrixFloatCalloc(n1 int, n2 int) *GslMatrixFloat {
	_ref := C.gsl_matrix_float_calloc(C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixFloat{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslMatrixFloat) Dispose() {
	C.gsl_matrix_float_free((*C.gsl_matrix_float)(unsafe.Pointer(x.Ptr())))
}

func FloatGet(m *GslMatrixFloat, i int, j int) float32 {
	return float32(C.gsl_matrix_float_get((*C.gsl_matrix_float)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func FloatSet(m *GslMatrixFloat, i int, j int, x float32) {
	C.gsl_matrix_float_set((*C.gsl_matrix_float)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j), C.float(x))
}

func FloatSetAll(m *GslMatrixFloat, x float32) {
	C.gsl_matrix_float_set_all((*C.gsl_matrix_float)(unsafe.Pointer(m.Ptr())), C.float(x))
}

func FloatSetZero(m *GslMatrixFloat) {
	C.gsl_matrix_float_set_zero((*C.gsl_matrix_float)(unsafe.Pointer(m.Ptr())))
}

func FloatSetIdentity(m *GslMatrixFloat) {
	C.gsl_matrix_float_set_identity((*C.gsl_matrix_float)(unsafe.Pointer(m.Ptr())))
}

func FloatFwrite(stream *os.File, m *GslMatrixFloat) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_matrix_float_fwrite(_file_0, (*C.gsl_matrix_float)(unsafe.Pointer(m.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func FloatFread(stream *os.File, m *GslMatrixFloat) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_matrix_float_fread(_file_0, (*C.gsl_matrix_float)(unsafe.Pointer(m.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func FloatFprintf(stream *os.File, m *GslMatrixFloat, format string) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_string_2 := C.CString(format)
	_result := int32(C.gsl_matrix_float_fprintf(_file_0, (*C.gsl_matrix_float)(unsafe.Pointer(m.Ptr())), _string_2))
	C.fclose(_file_0)
	C.free(unsafe.Pointer(_string_2))
	return _result
}

func FloatFscanf(stream *os.File, m *GslMatrixFloat) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_matrix_float_fscanf(_file_0, (*C.gsl_matrix_float)(unsafe.Pointer(m.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func FloatSubmatrix(m *GslMatrixFloat, k1 int, k2 int, n1 int, n2 int) *GslMatrixFloatView {
	_ref := C.gsl_matrix_float_submatrix((*C.gsl_matrix_float)(unsafe.Pointer(m.Ptr())), C.size_t(k1), C.size_t(k2), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixFloatView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func FloatConstSubmatrix(m *GslMatrixFloat, k1 int, k2 int, n1 int, n2 int) *GslMatrixFloatConstView {
	_ref := C.gsl_matrix_float_const_submatrix((*C.gsl_matrix_float)(unsafe.Pointer(m.Ptr())), C.size_t(k1), C.size_t(k2), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixFloatConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func FloatViewVector(v *vector.GslVectorFloat, n1 int, n2 int) *GslMatrixFloatView {
	_ref := C.gsl_matrix_float_view_vector((*C.gsl_vector_float)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixFloatView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func FloatConstViewVector(v *vector.GslVectorFloat, n1 int, n2 int) *GslMatrixFloatConstView {
	_ref := C.gsl_matrix_float_const_view_vector((*C.gsl_vector_float)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixFloatConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func FloatViewVectorWithTda(v *vector.GslVectorFloat, n1 int, n2 int, tda int) *GslMatrixFloatView {
	_ref := C.gsl_matrix_float_view_vector_with_tda((*C.gsl_vector_float)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2), C.size_t(tda))
	_result := &GslMatrixFloatView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func FloatConstViewVectorWithTda(v *vector.GslVectorFloat, n1 int, n2 int, tda int) *GslMatrixFloatConstView {
	_ref := C.gsl_matrix_float_const_view_vector_with_tda((*C.gsl_vector_float)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2), C.size_t(tda))
	_result := &GslMatrixFloatConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func FloatRow(m *GslMatrixFloat, i int) *vector.GslVectorFloatView {
	_ref := C.gsl_matrix_float_row((*C.gsl_matrix_float)(unsafe.Pointer(m.Ptr())), C.size_t(i))
	_result := &vector.GslVectorFloatView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func FloatConstRow(m *GslMatrixFloat, i int) *vector.GslVectorFloatConstView {
	_ref := C.gsl_matrix_float_const_row((*C.gsl_matrix_float)(unsafe.Pointer(m.Ptr())), C.size_t(i))
	_result := &vector.GslVectorFloatConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func FloatColumn(m *GslMatrixFloat, j int) *vector.GslVectorFloatView {
	_ref := C.gsl_matrix_float_column((*C.gsl_matrix_float)(unsafe.Pointer(m.Ptr())), C.size_t(j))
	_result := &vector.GslVectorFloatView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func FloatConstColumn(m *GslMatrixFloat, j int) *vector.GslVectorFloatConstView {
	_ref := C.gsl_matrix_float_const_column((*C.gsl_matrix_float)(unsafe.Pointer(m.Ptr())), C.size_t(j))
	_result := &vector.GslVectorFloatConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func FloatSubrow(m *GslMatrixFloat, i int, offset int, n int) *vector.GslVectorFloatView {
	_ref := C.gsl_matrix_float_subrow((*C.gsl_matrix_float)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorFloatView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func FloatConstSubrow(m *GslMatrixFloat, i int, offset int, n int) *vector.GslVectorFloatConstView {
	_ref := C.gsl_matrix_float_const_subrow((*C.gsl_matrix_float)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorFloatConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func FloatSubcolumn(m *GslMatrixFloat, j int, offset int, n int) *vector.GslVectorFloatView {
	_ref := C.gsl_matrix_float_subcolumn((*C.gsl_matrix_float)(unsafe.Pointer(m.Ptr())), C.size_t(j), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorFloatView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func FloatConstSubcolumn(m *GslMatrixFloat, j int, offset int, n int) *vector.GslVectorFloatConstView {
	_ref := C.gsl_matrix_float_const_subcolumn((*C.gsl_matrix_float)(unsafe.Pointer(m.Ptr())), C.size_t(j), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorFloatConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func FloatDiagonal(m *GslMatrixFloat) *vector.GslVectorFloatView {
	_ref := C.gsl_matrix_float_diagonal((*C.gsl_matrix_float)(unsafe.Pointer(m.Ptr())))
	_result := &vector.GslVectorFloatView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func FloatConstDiagonal(m *GslMatrixFloat) *vector.GslVectorFloatConstView {
	_ref := C.gsl_matrix_float_const_diagonal((*C.gsl_matrix_float)(unsafe.Pointer(m.Ptr())))
	_result := &vector.GslVectorFloatConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func FloatSubdiagonal(m *GslMatrixFloat, k int) *vector.GslVectorFloatView {
	_ref := C.gsl_matrix_float_subdiagonal((*C.gsl_matrix_float)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorFloatView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func FloatConstSubdiagonal(m *GslMatrixFloat, k int) *vector.GslVectorFloatConstView {
	_ref := C.gsl_matrix_float_const_subdiagonal((*C.gsl_matrix_float)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorFloatConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func FloatSuperdiagonal(m *GslMatrixFloat, k int) *vector.GslVectorFloatView {
	_ref := C.gsl_matrix_float_superdiagonal((*C.gsl_matrix_float)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorFloatView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func FloatConstSuperdiagonal(m *GslMatrixFloat, k int) *vector.GslVectorFloatConstView {
	_ref := C.gsl_matrix_float_const_superdiagonal((*C.gsl_matrix_float)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorFloatConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func FloatMemcpy(dest *GslMatrixFloat, src *GslMatrixFloat) int32 {
	return int32(C.gsl_matrix_float_memcpy((*C.gsl_matrix_float)(unsafe.Pointer(dest.Ptr())), (*C.gsl_matrix_float)(unsafe.Pointer(src.Ptr()))))
}

func FloatSwap(m1 *GslMatrixFloat, m2 *GslMatrixFloat) int32 {
	return int32(C.gsl_matrix_float_swap((*C.gsl_matrix_float)(unsafe.Pointer(m1.Ptr())), (*C.gsl_matrix_float)(unsafe.Pointer(m2.Ptr()))))
}

func FloatGetRow(v *vector.GslVectorFloat, m *GslMatrixFloat, i int) int32 {
	return int32(C.gsl_matrix_float_get_row((*C.gsl_vector_float)(unsafe.Pointer(v.Ptr())), (*C.gsl_matrix_float)(unsafe.Pointer(m.Ptr())), C.size_t(i)))
}

func FloatGetCol(v *vector.GslVectorFloat, m *GslMatrixFloat, j int) int32 {
	return int32(C.gsl_matrix_float_get_col((*C.gsl_vector_float)(unsafe.Pointer(v.Ptr())), (*C.gsl_matrix_float)(unsafe.Pointer(m.Ptr())), C.size_t(j)))
}

func FloatSetRow(m *GslMatrixFloat, i int, v *vector.GslVectorFloat) int32 {
	return int32(C.gsl_matrix_float_set_row((*C.gsl_matrix_float)(unsafe.Pointer(m.Ptr())), C.size_t(i), (*C.gsl_vector_float)(unsafe.Pointer(v.Ptr()))))
}

func FloatSetCol(m *GslMatrixFloat, j int, v *vector.GslVectorFloat) int32 {
	return int32(C.gsl_matrix_float_set_col((*C.gsl_matrix_float)(unsafe.Pointer(m.Ptr())), C.size_t(j), (*C.gsl_vector_float)(unsafe.Pointer(v.Ptr()))))
}

func FloatSwapRows(m *GslMatrixFloat, i int, j int) int32 {
	return int32(C.gsl_matrix_float_swap_rows((*C.gsl_matrix_float)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func FloatSwapColumns(m *GslMatrixFloat, i int, j int) int32 {
	return int32(C.gsl_matrix_float_swap_columns((*C.gsl_matrix_float)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func FloatSwapRowcol(m *GslMatrixFloat, i int, j int) int32 {
	return int32(C.gsl_matrix_float_swap_rowcol((*C.gsl_matrix_float)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func FloatTransposeMemcpy(dest *GslMatrixFloat, src *GslMatrixFloat) int32 {
	return int32(C.gsl_matrix_float_transpose_memcpy((*C.gsl_matrix_float)(unsafe.Pointer(dest.Ptr())), (*C.gsl_matrix_float)(unsafe.Pointer(src.Ptr()))))
}

func FloatTranspose(m *GslMatrixFloat) int32 {
	return int32(C.gsl_matrix_float_transpose((*C.gsl_matrix_float)(unsafe.Pointer(m.Ptr()))))
}

func FloatAdd(a *GslMatrixFloat, b *GslMatrixFloat) int32 {
	return int32(C.gsl_matrix_float_add((*C.gsl_matrix_float)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_float)(unsafe.Pointer(b.Ptr()))))
}

func FloatSub(a *GslMatrixFloat, b *GslMatrixFloat) int32 {
	return int32(C.gsl_matrix_float_sub((*C.gsl_matrix_float)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_float)(unsafe.Pointer(b.Ptr()))))
}

func FloatMulElements(a *GslMatrixFloat, b *GslMatrixFloat) int32 {
	return int32(C.gsl_matrix_float_mul_elements((*C.gsl_matrix_float)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_float)(unsafe.Pointer(b.Ptr()))))
}

func FloatDivElements(a *GslMatrixFloat, b *GslMatrixFloat) int32 {
	return int32(C.gsl_matrix_float_div_elements((*C.gsl_matrix_float)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_float)(unsafe.Pointer(b.Ptr()))))
}

func FloatScale(a *GslMatrixFloat, x float64) int32 {
	return int32(C.gsl_matrix_float_scale((*C.gsl_matrix_float)(unsafe.Pointer(a.Ptr())), C.double(x)))
}

func FloatAddConstant(a *GslMatrixFloat, x float64) int32 {
	return int32(C.gsl_matrix_float_add_constant((*C.gsl_matrix_float)(unsafe.Pointer(a.Ptr())), C.double(x)))
}

func FloatMax(m *GslMatrixFloat) float32 {
	return float32(C.gsl_matrix_float_max((*C.gsl_matrix_float)(unsafe.Pointer(m.Ptr()))))
}

func FloatMin(m *GslMatrixFloat) float32 {
	return float32(C.gsl_matrix_float_min((*C.gsl_matrix_float)(unsafe.Pointer(m.Ptr()))))
}

func FloatMinmax(m *GslMatrixFloat) (float32, float32) {
	var _outptr_1 C.float
	var _outptr_2 C.float
	C.gsl_matrix_float_minmax((*C.gsl_matrix_float)(unsafe.Pointer(m.Ptr())), &_outptr_1, &_outptr_2)
	return *(*float32)(unsafe.Pointer(&_outptr_1)), *(*float32)(unsafe.Pointer(&_outptr_2))
}

func FloatMaxIndex(m *GslMatrixFloat) (int, int) {
	var _outptr_1 C.size_t
	var _outptr_2 C.size_t
	C.gsl_matrix_float_max_index((*C.gsl_matrix_float)(unsafe.Pointer(m.Ptr())), &_outptr_1, &_outptr_2)
	return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2))
}

func FloatMinIndex(m *GslMatrixFloat) (int, int) {
	var _outptr_1 C.size_t
	var _outptr_2 C.size_t
	C.gsl_matrix_float_min_index((*C.gsl_matrix_float)(unsafe.Pointer(m.Ptr())), &_outptr_1, &_outptr_2)
	return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2))
}

func FloatMinmaxIndex(m *GslMatrixFloat) (int, int, int, int) {
	var _outptr_1 C.size_t
	var _outptr_2 C.size_t
	var _outptr_3 C.size_t
	var _outptr_4 C.size_t
	C.gsl_matrix_float_minmax_index((*C.gsl_matrix_float)(unsafe.Pointer(m.Ptr())), &_outptr_1, &_outptr_2, &_outptr_3, &_outptr_4)
	return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2)), *(*int)(unsafe.Pointer(&_outptr_3)), *(*int)(unsafe.Pointer(&_outptr_4))
}

func FloatIsnull(m *GslMatrixFloat) int32 {
	return int32(C.gsl_matrix_float_isnull((*C.gsl_matrix_float)(unsafe.Pointer(m.Ptr()))))
}

func FloatIspos(m *GslMatrixFloat) int32 {
	return int32(C.gsl_matrix_float_ispos((*C.gsl_matrix_float)(unsafe.Pointer(m.Ptr()))))
}

func FloatIsneg(m *GslMatrixFloat) int32 {
	return int32(C.gsl_matrix_float_isneg((*C.gsl_matrix_float)(unsafe.Pointer(m.Ptr()))))
}

func FloatIsnonneg(m *GslMatrixFloat) int32 {
	return int32(C.gsl_matrix_float_isnonneg((*C.gsl_matrix_float)(unsafe.Pointer(m.Ptr()))))
}

func FloatEqual(a *GslMatrixFloat, b *GslMatrixFloat) int32 {
	return int32(C.gsl_matrix_float_equal((*C.gsl_matrix_float)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_float)(unsafe.Pointer(b.Ptr()))))
}

type GslMatrixInt struct {
	gogsl.GslReference
}

type GslMatrixIntView struct {
	gogsl.GslReference
	CData []byte
}

type GslMatrixIntConstView struct {
	gogsl.GslReference
	CData []byte
}

func MatrixIntAlloc(n1 int, n2 int) *GslMatrixInt {
	_ref := C.gsl_matrix_int_alloc(C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixInt{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func MatrixIntCalloc(n1 int, n2 int) *GslMatrixInt {
	_ref := C.gsl_matrix_int_calloc(C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixInt{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslMatrixInt) Dispose() {
	C.gsl_matrix_int_free((*C.gsl_matrix_int)(unsafe.Pointer(x.Ptr())))
}

func IntGet(m *GslMatrixInt, i int, j int) int32 {
	return int32(C.gsl_matrix_int_get((*C.gsl_matrix_int)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func IntSet(m *GslMatrixInt, i int, j int, x int32) {
	C.gsl_matrix_int_set((*C.gsl_matrix_int)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j), C.int(x))
}

func IntSetAll(m *GslMatrixInt, x int32) {
	C.gsl_matrix_int_set_all((*C.gsl_matrix_int)(unsafe.Pointer(m.Ptr())), C.int(x))
}

func IntSetZero(m *GslMatrixInt) {
	C.gsl_matrix_int_set_zero((*C.gsl_matrix_int)(unsafe.Pointer(m.Ptr())))
}

func IntSetIdentity(m *GslMatrixInt) {
	C.gsl_matrix_int_set_identity((*C.gsl_matrix_int)(unsafe.Pointer(m.Ptr())))
}

func IntFwrite(stream *os.File, m *GslMatrixInt) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_matrix_int_fwrite(_file_0, (*C.gsl_matrix_int)(unsafe.Pointer(m.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func IntFread(stream *os.File, m *GslMatrixInt) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_matrix_int_fread(_file_0, (*C.gsl_matrix_int)(unsafe.Pointer(m.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func IntFprintf(stream *os.File, m *GslMatrixInt, format string) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_string_2 := C.CString(format)
	_result := int32(C.gsl_matrix_int_fprintf(_file_0, (*C.gsl_matrix_int)(unsafe.Pointer(m.Ptr())), _string_2))
	C.fclose(_file_0)
	C.free(unsafe.Pointer(_string_2))
	return _result
}

func IntFscanf(stream *os.File, m *GslMatrixInt) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_matrix_int_fscanf(_file_0, (*C.gsl_matrix_int)(unsafe.Pointer(m.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func IntSubmatrix(m *GslMatrixInt, k1 int, k2 int, n1 int, n2 int) *GslMatrixIntView {
	_ref := C.gsl_matrix_int_submatrix((*C.gsl_matrix_int)(unsafe.Pointer(m.Ptr())), C.size_t(k1), C.size_t(k2), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixIntView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func IntConstSubmatrix(m *GslMatrixInt, k1 int, k2 int, n1 int, n2 int) *GslMatrixIntConstView {
	_ref := C.gsl_matrix_int_const_submatrix((*C.gsl_matrix_int)(unsafe.Pointer(m.Ptr())), C.size_t(k1), C.size_t(k2), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixIntConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func IntViewVector(v *vector.GslVectorInt, n1 int, n2 int) *GslMatrixIntView {
	_ref := C.gsl_matrix_int_view_vector((*C.gsl_vector_int)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixIntView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func IntConstViewVector(v *vector.GslVectorInt, n1 int, n2 int) *GslMatrixIntConstView {
	_ref := C.gsl_matrix_int_const_view_vector((*C.gsl_vector_int)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixIntConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func IntViewVectorWithTda(v *vector.GslVectorInt, n1 int, n2 int, tda int) *GslMatrixIntView {
	_ref := C.gsl_matrix_int_view_vector_with_tda((*C.gsl_vector_int)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2), C.size_t(tda))
	_result := &GslMatrixIntView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func IntConstViewVectorWithTda(v *vector.GslVectorInt, n1 int, n2 int, tda int) *GslMatrixIntConstView {
	_ref := C.gsl_matrix_int_const_view_vector_with_tda((*C.gsl_vector_int)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2), C.size_t(tda))
	_result := &GslMatrixIntConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func IntRow(m *GslMatrixInt, i int) *vector.GslVectorIntView {
	_ref := C.gsl_matrix_int_row((*C.gsl_matrix_int)(unsafe.Pointer(m.Ptr())), C.size_t(i))
	_result := &vector.GslVectorIntView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func IntConstRow(m *GslMatrixInt, i int) *vector.GslVectorIntConstView {
	_ref := C.gsl_matrix_int_const_row((*C.gsl_matrix_int)(unsafe.Pointer(m.Ptr())), C.size_t(i))
	_result := &vector.GslVectorIntConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func IntColumn(m *GslMatrixInt, j int) *vector.GslVectorIntView {
	_ref := C.gsl_matrix_int_column((*C.gsl_matrix_int)(unsafe.Pointer(m.Ptr())), C.size_t(j))
	_result := &vector.GslVectorIntView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func IntConstColumn(m *GslMatrixInt, j int) *vector.GslVectorIntConstView {
	_ref := C.gsl_matrix_int_const_column((*C.gsl_matrix_int)(unsafe.Pointer(m.Ptr())), C.size_t(j))
	_result := &vector.GslVectorIntConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func IntSubrow(m *GslMatrixInt, i int, offset int, n int) *vector.GslVectorIntView {
	_ref := C.gsl_matrix_int_subrow((*C.gsl_matrix_int)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorIntView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func IntConstSubrow(m *GslMatrixInt, i int, offset int, n int) *vector.GslVectorIntConstView {
	_ref := C.gsl_matrix_int_const_subrow((*C.gsl_matrix_int)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorIntConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func IntSubcolumn(m *GslMatrixInt, j int, offset int, n int) *vector.GslVectorIntView {
	_ref := C.gsl_matrix_int_subcolumn((*C.gsl_matrix_int)(unsafe.Pointer(m.Ptr())), C.size_t(j), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorIntView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func IntConstSubcolumn(m *GslMatrixInt, j int, offset int, n int) *vector.GslVectorIntConstView {
	_ref := C.gsl_matrix_int_const_subcolumn((*C.gsl_matrix_int)(unsafe.Pointer(m.Ptr())), C.size_t(j), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorIntConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func IntDiagonal(m *GslMatrixInt) *vector.GslVectorIntView {
	_ref := C.gsl_matrix_int_diagonal((*C.gsl_matrix_int)(unsafe.Pointer(m.Ptr())))
	_result := &vector.GslVectorIntView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func IntConstDiagonal(m *GslMatrixInt) *vector.GslVectorIntConstView {
	_ref := C.gsl_matrix_int_const_diagonal((*C.gsl_matrix_int)(unsafe.Pointer(m.Ptr())))
	_result := &vector.GslVectorIntConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func IntSubdiagonal(m *GslMatrixInt, k int) *vector.GslVectorIntView {
	_ref := C.gsl_matrix_int_subdiagonal((*C.gsl_matrix_int)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorIntView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func IntConstSubdiagonal(m *GslMatrixInt, k int) *vector.GslVectorIntConstView {
	_ref := C.gsl_matrix_int_const_subdiagonal((*C.gsl_matrix_int)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorIntConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func IntSuperdiagonal(m *GslMatrixInt, k int) *vector.GslVectorIntView {
	_ref := C.gsl_matrix_int_superdiagonal((*C.gsl_matrix_int)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorIntView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func IntConstSuperdiagonal(m *GslMatrixInt, k int) *vector.GslVectorIntConstView {
	_ref := C.gsl_matrix_int_const_superdiagonal((*C.gsl_matrix_int)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorIntConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func IntMemcpy(dest *GslMatrixInt, src *GslMatrixInt) int32 {
	return int32(C.gsl_matrix_int_memcpy((*C.gsl_matrix_int)(unsafe.Pointer(dest.Ptr())), (*C.gsl_matrix_int)(unsafe.Pointer(src.Ptr()))))
}

func IntSwap(m1 *GslMatrixInt, m2 *GslMatrixInt) int32 {
	return int32(C.gsl_matrix_int_swap((*C.gsl_matrix_int)(unsafe.Pointer(m1.Ptr())), (*C.gsl_matrix_int)(unsafe.Pointer(m2.Ptr()))))
}

func IntGetRow(v *vector.GslVectorInt, m *GslMatrixInt, i int) int32 {
	return int32(C.gsl_matrix_int_get_row((*C.gsl_vector_int)(unsafe.Pointer(v.Ptr())), (*C.gsl_matrix_int)(unsafe.Pointer(m.Ptr())), C.size_t(i)))
}

func IntGetCol(v *vector.GslVectorInt, m *GslMatrixInt, j int) int32 {
	return int32(C.gsl_matrix_int_get_col((*C.gsl_vector_int)(unsafe.Pointer(v.Ptr())), (*C.gsl_matrix_int)(unsafe.Pointer(m.Ptr())), C.size_t(j)))
}

func IntSetRow(m *GslMatrixInt, i int, v *vector.GslVectorInt) int32 {
	return int32(C.gsl_matrix_int_set_row((*C.gsl_matrix_int)(unsafe.Pointer(m.Ptr())), C.size_t(i), (*C.gsl_vector_int)(unsafe.Pointer(v.Ptr()))))
}

func IntSetCol(m *GslMatrixInt, j int, v *vector.GslVectorInt) int32 {
	return int32(C.gsl_matrix_int_set_col((*C.gsl_matrix_int)(unsafe.Pointer(m.Ptr())), C.size_t(j), (*C.gsl_vector_int)(unsafe.Pointer(v.Ptr()))))
}

func IntSwapRows(m *GslMatrixInt, i int, j int) int32 {
	return int32(C.gsl_matrix_int_swap_rows((*C.gsl_matrix_int)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func IntSwapColumns(m *GslMatrixInt, i int, j int) int32 {
	return int32(C.gsl_matrix_int_swap_columns((*C.gsl_matrix_int)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func IntSwapRowcol(m *GslMatrixInt, i int, j int) int32 {
	return int32(C.gsl_matrix_int_swap_rowcol((*C.gsl_matrix_int)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func IntTransposeMemcpy(dest *GslMatrixInt, src *GslMatrixInt) int32 {
	return int32(C.gsl_matrix_int_transpose_memcpy((*C.gsl_matrix_int)(unsafe.Pointer(dest.Ptr())), (*C.gsl_matrix_int)(unsafe.Pointer(src.Ptr()))))
}

func IntTranspose(m *GslMatrixInt) int32 {
	return int32(C.gsl_matrix_int_transpose((*C.gsl_matrix_int)(unsafe.Pointer(m.Ptr()))))
}

func IntAdd(a *GslMatrixInt, b *GslMatrixInt) int32 {
	return int32(C.gsl_matrix_int_add((*C.gsl_matrix_int)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_int)(unsafe.Pointer(b.Ptr()))))
}

func IntSub(a *GslMatrixInt, b *GslMatrixInt) int32 {
	return int32(C.gsl_matrix_int_sub((*C.gsl_matrix_int)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_int)(unsafe.Pointer(b.Ptr()))))
}

func IntMulElements(a *GslMatrixInt, b *GslMatrixInt) int32 {
	return int32(C.gsl_matrix_int_mul_elements((*C.gsl_matrix_int)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_int)(unsafe.Pointer(b.Ptr()))))
}

func IntDivElements(a *GslMatrixInt, b *GslMatrixInt) int32 {
	return int32(C.gsl_matrix_int_div_elements((*C.gsl_matrix_int)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_int)(unsafe.Pointer(b.Ptr()))))
}

func IntScale(a *GslMatrixInt, x float64) int32 {
	return int32(C.gsl_matrix_int_scale((*C.gsl_matrix_int)(unsafe.Pointer(a.Ptr())), C.double(x)))
}

func IntAddConstant(a *GslMatrixInt, x float64) int32 {
	return int32(C.gsl_matrix_int_add_constant((*C.gsl_matrix_int)(unsafe.Pointer(a.Ptr())), C.double(x)))
}

func IntMax(m *GslMatrixInt) int32 {
	return int32(C.gsl_matrix_int_max((*C.gsl_matrix_int)(unsafe.Pointer(m.Ptr()))))
}

func IntMin(m *GslMatrixInt) int32 {
	return int32(C.gsl_matrix_int_min((*C.gsl_matrix_int)(unsafe.Pointer(m.Ptr()))))
}

func IntMinmax(m *GslMatrixInt) (int32, int32) {
	var _outptr_1 C.int
	var _outptr_2 C.int
	C.gsl_matrix_int_minmax((*C.gsl_matrix_int)(unsafe.Pointer(m.Ptr())), &_outptr_1, &_outptr_2)
	return *(*int32)(unsafe.Pointer(&_outptr_1)), *(*int32)(unsafe.Pointer(&_outptr_2))
}

func IntMaxIndex(m *GslMatrixInt) (int, int) {
	var _outptr_1 C.size_t
	var _outptr_2 C.size_t
	C.gsl_matrix_int_max_index((*C.gsl_matrix_int)(unsafe.Pointer(m.Ptr())), &_outptr_1, &_outptr_2)
	return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2))
}

func IntMinIndex(m *GslMatrixInt) (int, int) {
	var _outptr_1 C.size_t
	var _outptr_2 C.size_t
	C.gsl_matrix_int_min_index((*C.gsl_matrix_int)(unsafe.Pointer(m.Ptr())), &_outptr_1, &_outptr_2)
	return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2))
}

func IntMinmaxIndex(m *GslMatrixInt) (int, int, int, int) {
	var _outptr_1 C.size_t
	var _outptr_2 C.size_t
	var _outptr_3 C.size_t
	var _outptr_4 C.size_t
	C.gsl_matrix_int_minmax_index((*C.gsl_matrix_int)(unsafe.Pointer(m.Ptr())), &_outptr_1, &_outptr_2, &_outptr_3, &_outptr_4)
	return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2)), *(*int)(unsafe.Pointer(&_outptr_3)), *(*int)(unsafe.Pointer(&_outptr_4))
}

func IntIsnull(m *GslMatrixInt) int32 {
	return int32(C.gsl_matrix_int_isnull((*C.gsl_matrix_int)(unsafe.Pointer(m.Ptr()))))
}

func IntIspos(m *GslMatrixInt) int32 {
	return int32(C.gsl_matrix_int_ispos((*C.gsl_matrix_int)(unsafe.Pointer(m.Ptr()))))
}

func IntIsneg(m *GslMatrixInt) int32 {
	return int32(C.gsl_matrix_int_isneg((*C.gsl_matrix_int)(unsafe.Pointer(m.Ptr()))))
}

func IntIsnonneg(m *GslMatrixInt) int32 {
	return int32(C.gsl_matrix_int_isnonneg((*C.gsl_matrix_int)(unsafe.Pointer(m.Ptr()))))
}

func IntEqual(a *GslMatrixInt, b *GslMatrixInt) int32 {
	return int32(C.gsl_matrix_int_equal((*C.gsl_matrix_int)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_int)(unsafe.Pointer(b.Ptr()))))
}

type GslMatrixUint struct {
	gogsl.GslReference
}

type GslMatrixUintView struct {
	gogsl.GslReference
	CData []byte
}

type GslMatrixUintConstView struct {
	gogsl.GslReference
	CData []byte
}

func MatrixUintAlloc(n1 int, n2 int) *GslMatrixUint {
	_ref := C.gsl_matrix_uint_alloc(C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixUint{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func MatrixUintCalloc(n1 int, n2 int) *GslMatrixUint {
	_ref := C.gsl_matrix_uint_calloc(C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixUint{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslMatrixUint) Dispose() {
	C.gsl_matrix_uint_free((*C.gsl_matrix_uint)(unsafe.Pointer(x.Ptr())))
}

func UintGet(m *GslMatrixUint, i int, j int) uint32 {
	return uint32(C.gsl_matrix_uint_get((*C.gsl_matrix_uint)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func UintSet(m *GslMatrixUint, i int, j int, x uint32) {
	C.gsl_matrix_uint_set((*C.gsl_matrix_uint)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j), C.uint(x))
}

func UintSetAll(m *GslMatrixUint, x uint32) {
	C.gsl_matrix_uint_set_all((*C.gsl_matrix_uint)(unsafe.Pointer(m.Ptr())), C.uint(x))
}

func UintSetZero(m *GslMatrixUint) {
	C.gsl_matrix_uint_set_zero((*C.gsl_matrix_uint)(unsafe.Pointer(m.Ptr())))
}

func UintSetIdentity(m *GslMatrixUint) {
	C.gsl_matrix_uint_set_identity((*C.gsl_matrix_uint)(unsafe.Pointer(m.Ptr())))
}

func UintFwrite(stream *os.File, m *GslMatrixUint) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_matrix_uint_fwrite(_file_0, (*C.gsl_matrix_uint)(unsafe.Pointer(m.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func UintFread(stream *os.File, m *GslMatrixUint) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_matrix_uint_fread(_file_0, (*C.gsl_matrix_uint)(unsafe.Pointer(m.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func UintFprintf(stream *os.File, m *GslMatrixUint, format string) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_string_2 := C.CString(format)
	_result := int32(C.gsl_matrix_uint_fprintf(_file_0, (*C.gsl_matrix_uint)(unsafe.Pointer(m.Ptr())), _string_2))
	C.fclose(_file_0)
	C.free(unsafe.Pointer(_string_2))
	return _result
}

func UintFscanf(stream *os.File, m *GslMatrixUint) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_matrix_uint_fscanf(_file_0, (*C.gsl_matrix_uint)(unsafe.Pointer(m.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func UintSubmatrix(m *GslMatrixUint, k1 int, k2 int, n1 int, n2 int) *GslMatrixUintView {
	_ref := C.gsl_matrix_uint_submatrix((*C.gsl_matrix_uint)(unsafe.Pointer(m.Ptr())), C.size_t(k1), C.size_t(k2), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixUintView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UintConstSubmatrix(m *GslMatrixUint, k1 int, k2 int, n1 int, n2 int) *GslMatrixUintConstView {
	_ref := C.gsl_matrix_uint_const_submatrix((*C.gsl_matrix_uint)(unsafe.Pointer(m.Ptr())), C.size_t(k1), C.size_t(k2), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixUintConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UintViewVector(v *vector.GslVectorUint, n1 int, n2 int) *GslMatrixUintView {
	_ref := C.gsl_matrix_uint_view_vector((*C.gsl_vector_uint)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixUintView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UintConstViewVector(v *vector.GslVectorUint, n1 int, n2 int) *GslMatrixUintConstView {
	_ref := C.gsl_matrix_uint_const_view_vector((*C.gsl_vector_uint)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixUintConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UintViewVectorWithTda(v *vector.GslVectorUint, n1 int, n2 int, tda int) *GslMatrixUintView {
	_ref := C.gsl_matrix_uint_view_vector_with_tda((*C.gsl_vector_uint)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2), C.size_t(tda))
	_result := &GslMatrixUintView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UintConstViewVectorWithTda(v *vector.GslVectorUint, n1 int, n2 int, tda int) *GslMatrixUintConstView {
	_ref := C.gsl_matrix_uint_const_view_vector_with_tda((*C.gsl_vector_uint)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2), C.size_t(tda))
	_result := &GslMatrixUintConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UintRow(m *GslMatrixUint, i int) *vector.GslVectorUintView {
	_ref := C.gsl_matrix_uint_row((*C.gsl_matrix_uint)(unsafe.Pointer(m.Ptr())), C.size_t(i))
	_result := &vector.GslVectorUintView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UintConstRow(m *GslMatrixUint, i int) *vector.GslVectorUintConstView {
	_ref := C.gsl_matrix_uint_const_row((*C.gsl_matrix_uint)(unsafe.Pointer(m.Ptr())), C.size_t(i))
	_result := &vector.GslVectorUintConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UintColumn(m *GslMatrixUint, j int) *vector.GslVectorUintView {
	_ref := C.gsl_matrix_uint_column((*C.gsl_matrix_uint)(unsafe.Pointer(m.Ptr())), C.size_t(j))
	_result := &vector.GslVectorUintView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UintConstColumn(m *GslMatrixUint, j int) *vector.GslVectorUintConstView {
	_ref := C.gsl_matrix_uint_const_column((*C.gsl_matrix_uint)(unsafe.Pointer(m.Ptr())), C.size_t(j))
	_result := &vector.GslVectorUintConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UintSubrow(m *GslMatrixUint, i int, offset int, n int) *vector.GslVectorUintView {
	_ref := C.gsl_matrix_uint_subrow((*C.gsl_matrix_uint)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorUintView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UintConstSubrow(m *GslMatrixUint, i int, offset int, n int) *vector.GslVectorUintConstView {
	_ref := C.gsl_matrix_uint_const_subrow((*C.gsl_matrix_uint)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorUintConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UintSubcolumn(m *GslMatrixUint, j int, offset int, n int) *vector.GslVectorUintView {
	_ref := C.gsl_matrix_uint_subcolumn((*C.gsl_matrix_uint)(unsafe.Pointer(m.Ptr())), C.size_t(j), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorUintView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UintConstSubcolumn(m *GslMatrixUint, j int, offset int, n int) *vector.GslVectorUintConstView {
	_ref := C.gsl_matrix_uint_const_subcolumn((*C.gsl_matrix_uint)(unsafe.Pointer(m.Ptr())), C.size_t(j), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorUintConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UintDiagonal(m *GslMatrixUint) *vector.GslVectorUintView {
	_ref := C.gsl_matrix_uint_diagonal((*C.gsl_matrix_uint)(unsafe.Pointer(m.Ptr())))
	_result := &vector.GslVectorUintView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UintConstDiagonal(m *GslMatrixUint) *vector.GslVectorUintConstView {
	_ref := C.gsl_matrix_uint_const_diagonal((*C.gsl_matrix_uint)(unsafe.Pointer(m.Ptr())))
	_result := &vector.GslVectorUintConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UintSubdiagonal(m *GslMatrixUint, k int) *vector.GslVectorUintView {
	_ref := C.gsl_matrix_uint_subdiagonal((*C.gsl_matrix_uint)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorUintView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UintConstSubdiagonal(m *GslMatrixUint, k int) *vector.GslVectorUintConstView {
	_ref := C.gsl_matrix_uint_const_subdiagonal((*C.gsl_matrix_uint)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorUintConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UintSuperdiagonal(m *GslMatrixUint, k int) *vector.GslVectorUintView {
	_ref := C.gsl_matrix_uint_superdiagonal((*C.gsl_matrix_uint)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorUintView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UintConstSuperdiagonal(m *GslMatrixUint, k int) *vector.GslVectorUintConstView {
	_ref := C.gsl_matrix_uint_const_superdiagonal((*C.gsl_matrix_uint)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorUintConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UintMemcpy(dest *GslMatrixUint, src *GslMatrixUint) int32 {
	return int32(C.gsl_matrix_uint_memcpy((*C.gsl_matrix_uint)(unsafe.Pointer(dest.Ptr())), (*C.gsl_matrix_uint)(unsafe.Pointer(src.Ptr()))))
}

func UintSwap(m1 *GslMatrixUint, m2 *GslMatrixUint) int32 {
	return int32(C.gsl_matrix_uint_swap((*C.gsl_matrix_uint)(unsafe.Pointer(m1.Ptr())), (*C.gsl_matrix_uint)(unsafe.Pointer(m2.Ptr()))))
}

func UintGetRow(v *vector.GslVectorUint, m *GslMatrixUint, i int) int32 {
	return int32(C.gsl_matrix_uint_get_row((*C.gsl_vector_uint)(unsafe.Pointer(v.Ptr())), (*C.gsl_matrix_uint)(unsafe.Pointer(m.Ptr())), C.size_t(i)))
}

func UintGetCol(v *vector.GslVectorUint, m *GslMatrixUint, j int) int32 {
	return int32(C.gsl_matrix_uint_get_col((*C.gsl_vector_uint)(unsafe.Pointer(v.Ptr())), (*C.gsl_matrix_uint)(unsafe.Pointer(m.Ptr())), C.size_t(j)))
}

func UintSetRow(m *GslMatrixUint, i int, v *vector.GslVectorUint) int32 {
	return int32(C.gsl_matrix_uint_set_row((*C.gsl_matrix_uint)(unsafe.Pointer(m.Ptr())), C.size_t(i), (*C.gsl_vector_uint)(unsafe.Pointer(v.Ptr()))))
}

func UintSetCol(m *GslMatrixUint, j int, v *vector.GslVectorUint) int32 {
	return int32(C.gsl_matrix_uint_set_col((*C.gsl_matrix_uint)(unsafe.Pointer(m.Ptr())), C.size_t(j), (*C.gsl_vector_uint)(unsafe.Pointer(v.Ptr()))))
}

func UintSwapRows(m *GslMatrixUint, i int, j int) int32 {
	return int32(C.gsl_matrix_uint_swap_rows((*C.gsl_matrix_uint)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func UintSwapColumns(m *GslMatrixUint, i int, j int) int32 {
	return int32(C.gsl_matrix_uint_swap_columns((*C.gsl_matrix_uint)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func UintSwapRowcol(m *GslMatrixUint, i int, j int) int32 {
	return int32(C.gsl_matrix_uint_swap_rowcol((*C.gsl_matrix_uint)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func UintTransposeMemcpy(dest *GslMatrixUint, src *GslMatrixUint) int32 {
	return int32(C.gsl_matrix_uint_transpose_memcpy((*C.gsl_matrix_uint)(unsafe.Pointer(dest.Ptr())), (*C.gsl_matrix_uint)(unsafe.Pointer(src.Ptr()))))
}

func UintTranspose(m *GslMatrixUint) int32 {
	return int32(C.gsl_matrix_uint_transpose((*C.gsl_matrix_uint)(unsafe.Pointer(m.Ptr()))))
}

func UintAdd(a *GslMatrixUint, b *GslMatrixUint) int32 {
	return int32(C.gsl_matrix_uint_add((*C.gsl_matrix_uint)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_uint)(unsafe.Pointer(b.Ptr()))))
}

func UintSub(a *GslMatrixUint, b *GslMatrixUint) int32 {
	return int32(C.gsl_matrix_uint_sub((*C.gsl_matrix_uint)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_uint)(unsafe.Pointer(b.Ptr()))))
}

func UintMulElements(a *GslMatrixUint, b *GslMatrixUint) int32 {
	return int32(C.gsl_matrix_uint_mul_elements((*C.gsl_matrix_uint)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_uint)(unsafe.Pointer(b.Ptr()))))
}

func UintDivElements(a *GslMatrixUint, b *GslMatrixUint) int32 {
	return int32(C.gsl_matrix_uint_div_elements((*C.gsl_matrix_uint)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_uint)(unsafe.Pointer(b.Ptr()))))
}

func UintScale(a *GslMatrixUint, x float64) int32 {
	return int32(C.gsl_matrix_uint_scale((*C.gsl_matrix_uint)(unsafe.Pointer(a.Ptr())), C.double(x)))
}

func UintAddConstant(a *GslMatrixUint, x float64) int32 {
	return int32(C.gsl_matrix_uint_add_constant((*C.gsl_matrix_uint)(unsafe.Pointer(a.Ptr())), C.double(x)))
}

func UintMax(m *GslMatrixUint) uint32 {
	return uint32(C.gsl_matrix_uint_max((*C.gsl_matrix_uint)(unsafe.Pointer(m.Ptr()))))
}

func UintMin(m *GslMatrixUint) uint32 {
	return uint32(C.gsl_matrix_uint_min((*C.gsl_matrix_uint)(unsafe.Pointer(m.Ptr()))))
}

func UintMinmax(m *GslMatrixUint) (uint32, uint32) {
	var _outptr_1 C.uint
	var _outptr_2 C.uint
	C.gsl_matrix_uint_minmax((*C.gsl_matrix_uint)(unsafe.Pointer(m.Ptr())), &_outptr_1, &_outptr_2)
	return *(*uint32)(unsafe.Pointer(&_outptr_1)), *(*uint32)(unsafe.Pointer(&_outptr_2))
}

func UintMaxIndex(m *GslMatrixUint) (int, int) {
	var _outptr_1 C.size_t
	var _outptr_2 C.size_t
	C.gsl_matrix_uint_max_index((*C.gsl_matrix_uint)(unsafe.Pointer(m.Ptr())), &_outptr_1, &_outptr_2)
	return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2))
}

func UintMinIndex(m *GslMatrixUint) (int, int) {
	var _outptr_1 C.size_t
	var _outptr_2 C.size_t
	C.gsl_matrix_uint_min_index((*C.gsl_matrix_uint)(unsafe.Pointer(m.Ptr())), &_outptr_1, &_outptr_2)
	return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2))
}

func UintMinmaxIndex(m *GslMatrixUint) (int, int, int, int) {
	var _outptr_1 C.size_t
	var _outptr_2 C.size_t
	var _outptr_3 C.size_t
	var _outptr_4 C.size_t
	C.gsl_matrix_uint_minmax_index((*C.gsl_matrix_uint)(unsafe.Pointer(m.Ptr())), &_outptr_1, &_outptr_2, &_outptr_3, &_outptr_4)
	return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2)), *(*int)(unsafe.Pointer(&_outptr_3)), *(*int)(unsafe.Pointer(&_outptr_4))
}

func UintIsnull(m *GslMatrixUint) int32 {
	return int32(C.gsl_matrix_uint_isnull((*C.gsl_matrix_uint)(unsafe.Pointer(m.Ptr()))))
}

func UintIspos(m *GslMatrixUint) int32 {
	return int32(C.gsl_matrix_uint_ispos((*C.gsl_matrix_uint)(unsafe.Pointer(m.Ptr()))))
}

func UintIsneg(m *GslMatrixUint) int32 {
	return int32(C.gsl_matrix_uint_isneg((*C.gsl_matrix_uint)(unsafe.Pointer(m.Ptr()))))
}

func UintIsnonneg(m *GslMatrixUint) int32 {
	return int32(C.gsl_matrix_uint_isnonneg((*C.gsl_matrix_uint)(unsafe.Pointer(m.Ptr()))))
}

func UintEqual(a *GslMatrixUint, b *GslMatrixUint) int32 {
	return int32(C.gsl_matrix_uint_equal((*C.gsl_matrix_uint)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_uint)(unsafe.Pointer(b.Ptr()))))
}

type GslMatrixLong struct {
	gogsl.GslReference
}

type GslMatrixLongView struct {
	gogsl.GslReference
	CData []byte
}

type GslMatrixLongConstView struct {
	gogsl.GslReference
	CData []byte
}

func MatrixLongAlloc(n1 int, n2 int) *GslMatrixLong {
	_ref := C.gsl_matrix_long_alloc(C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixLong{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func MatrixLongCalloc(n1 int, n2 int) *GslMatrixLong {
	_ref := C.gsl_matrix_long_calloc(C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixLong{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslMatrixLong) Dispose() {
	C.gsl_matrix_long_free((*C.gsl_matrix_long)(unsafe.Pointer(x.Ptr())))
}

func LongGet(m *GslMatrixLong, i int, j int) uint {
	return uint(C.gsl_matrix_long_get((*C.gsl_matrix_long)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func LongSet(m *GslMatrixLong, i int, j int, x uint) {
	C.gsl_matrix_long_set((*C.gsl_matrix_long)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j), C.long(x))
}

func LongSetAll(m *GslMatrixLong, x uint) {
	C.gsl_matrix_long_set_all((*C.gsl_matrix_long)(unsafe.Pointer(m.Ptr())), C.long(x))
}

func LongSetZero(m *GslMatrixLong) {
	C.gsl_matrix_long_set_zero((*C.gsl_matrix_long)(unsafe.Pointer(m.Ptr())))
}

func LongSetIdentity(m *GslMatrixLong) {
	C.gsl_matrix_long_set_identity((*C.gsl_matrix_long)(unsafe.Pointer(m.Ptr())))
}

func LongFwrite(stream *os.File, m *GslMatrixLong) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_matrix_long_fwrite(_file_0, (*C.gsl_matrix_long)(unsafe.Pointer(m.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func LongFread(stream *os.File, m *GslMatrixLong) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_matrix_long_fread(_file_0, (*C.gsl_matrix_long)(unsafe.Pointer(m.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func LongFprintf(stream *os.File, m *GslMatrixLong, format string) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_string_2 := C.CString(format)
	_result := int32(C.gsl_matrix_long_fprintf(_file_0, (*C.gsl_matrix_long)(unsafe.Pointer(m.Ptr())), _string_2))
	C.fclose(_file_0)
	C.free(unsafe.Pointer(_string_2))
	return _result
}

func LongFscanf(stream *os.File, m *GslMatrixLong) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_matrix_long_fscanf(_file_0, (*C.gsl_matrix_long)(unsafe.Pointer(m.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func LongSubmatrix(m *GslMatrixLong, k1 int, k2 int, n1 int, n2 int) *GslMatrixLongView {
	_ref := C.gsl_matrix_long_submatrix((*C.gsl_matrix_long)(unsafe.Pointer(m.Ptr())), C.size_t(k1), C.size_t(k2), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixLongView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func LongConstSubmatrix(m *GslMatrixLong, k1 int, k2 int, n1 int, n2 int) *GslMatrixLongConstView {
	_ref := C.gsl_matrix_long_const_submatrix((*C.gsl_matrix_long)(unsafe.Pointer(m.Ptr())), C.size_t(k1), C.size_t(k2), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixLongConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func LongViewVector(v *vector.GslVectorLong, n1 int, n2 int) *GslMatrixLongView {
	_ref := C.gsl_matrix_long_view_vector((*C.gsl_vector_long)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixLongView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func LongConstViewVector(v *vector.GslVectorLong, n1 int, n2 int) *GslMatrixLongConstView {
	_ref := C.gsl_matrix_long_const_view_vector((*C.gsl_vector_long)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixLongConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func LongViewVectorWithTda(v *vector.GslVectorLong, n1 int, n2 int, tda int) *GslMatrixLongView {
	_ref := C.gsl_matrix_long_view_vector_with_tda((*C.gsl_vector_long)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2), C.size_t(tda))
	_result := &GslMatrixLongView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func LongConstViewVectorWithTda(v *vector.GslVectorLong, n1 int, n2 int, tda int) *GslMatrixLongConstView {
	_ref := C.gsl_matrix_long_const_view_vector_with_tda((*C.gsl_vector_long)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2), C.size_t(tda))
	_result := &GslMatrixLongConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func LongRow(m *GslMatrixLong, i int) *vector.GslVectorLongView {
	_ref := C.gsl_matrix_long_row((*C.gsl_matrix_long)(unsafe.Pointer(m.Ptr())), C.size_t(i))
	_result := &vector.GslVectorLongView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func LongConstRow(m *GslMatrixLong, i int) *vector.GslVectorLongConstView {
	_ref := C.gsl_matrix_long_const_row((*C.gsl_matrix_long)(unsafe.Pointer(m.Ptr())), C.size_t(i))
	_result := &vector.GslVectorLongConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func LongColumn(m *GslMatrixLong, j int) *vector.GslVectorLongView {
	_ref := C.gsl_matrix_long_column((*C.gsl_matrix_long)(unsafe.Pointer(m.Ptr())), C.size_t(j))
	_result := &vector.GslVectorLongView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func LongConstColumn(m *GslMatrixLong, j int) *vector.GslVectorLongConstView {
	_ref := C.gsl_matrix_long_const_column((*C.gsl_matrix_long)(unsafe.Pointer(m.Ptr())), C.size_t(j))
	_result := &vector.GslVectorLongConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func LongSubrow(m *GslMatrixLong, i int, offset int, n int) *vector.GslVectorLongView {
	_ref := C.gsl_matrix_long_subrow((*C.gsl_matrix_long)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorLongView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func LongConstSubrow(m *GslMatrixLong, i int, offset int, n int) *vector.GslVectorLongConstView {
	_ref := C.gsl_matrix_long_const_subrow((*C.gsl_matrix_long)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorLongConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func LongSubcolumn(m *GslMatrixLong, j int, offset int, n int) *vector.GslVectorLongView {
	_ref := C.gsl_matrix_long_subcolumn((*C.gsl_matrix_long)(unsafe.Pointer(m.Ptr())), C.size_t(j), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorLongView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func LongConstSubcolumn(m *GslMatrixLong, j int, offset int, n int) *vector.GslVectorLongConstView {
	_ref := C.gsl_matrix_long_const_subcolumn((*C.gsl_matrix_long)(unsafe.Pointer(m.Ptr())), C.size_t(j), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorLongConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func LongDiagonal(m *GslMatrixLong) *vector.GslVectorLongView {
	_ref := C.gsl_matrix_long_diagonal((*C.gsl_matrix_long)(unsafe.Pointer(m.Ptr())))
	_result := &vector.GslVectorLongView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func LongConstDiagonal(m *GslMatrixLong) *vector.GslVectorLongConstView {
	_ref := C.gsl_matrix_long_const_diagonal((*C.gsl_matrix_long)(unsafe.Pointer(m.Ptr())))
	_result := &vector.GslVectorLongConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func LongSubdiagonal(m *GslMatrixLong, k int) *vector.GslVectorLongView {
	_ref := C.gsl_matrix_long_subdiagonal((*C.gsl_matrix_long)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorLongView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func LongConstSubdiagonal(m *GslMatrixLong, k int) *vector.GslVectorLongConstView {
	_ref := C.gsl_matrix_long_const_subdiagonal((*C.gsl_matrix_long)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorLongConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func LongSuperdiagonal(m *GslMatrixLong, k int) *vector.GslVectorLongView {
	_ref := C.gsl_matrix_long_superdiagonal((*C.gsl_matrix_long)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorLongView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func LongConstSuperdiagonal(m *GslMatrixLong, k int) *vector.GslVectorLongConstView {
	_ref := C.gsl_matrix_long_const_superdiagonal((*C.gsl_matrix_long)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorLongConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func LongMemcpy(dest *GslMatrixLong, src *GslMatrixLong) int32 {
	return int32(C.gsl_matrix_long_memcpy((*C.gsl_matrix_long)(unsafe.Pointer(dest.Ptr())), (*C.gsl_matrix_long)(unsafe.Pointer(src.Ptr()))))
}

func LongSwap(m1 *GslMatrixLong, m2 *GslMatrixLong) int32 {
	return int32(C.gsl_matrix_long_swap((*C.gsl_matrix_long)(unsafe.Pointer(m1.Ptr())), (*C.gsl_matrix_long)(unsafe.Pointer(m2.Ptr()))))
}

func LongGetRow(v *vector.GslVectorLong, m *GslMatrixLong, i int) int32 {
	return int32(C.gsl_matrix_long_get_row((*C.gsl_vector_long)(unsafe.Pointer(v.Ptr())), (*C.gsl_matrix_long)(unsafe.Pointer(m.Ptr())), C.size_t(i)))
}

func LongGetCol(v *vector.GslVectorLong, m *GslMatrixLong, j int) int32 {
	return int32(C.gsl_matrix_long_get_col((*C.gsl_vector_long)(unsafe.Pointer(v.Ptr())), (*C.gsl_matrix_long)(unsafe.Pointer(m.Ptr())), C.size_t(j)))
}

func LongSetRow(m *GslMatrixLong, i int, v *vector.GslVectorLong) int32 {
	return int32(C.gsl_matrix_long_set_row((*C.gsl_matrix_long)(unsafe.Pointer(m.Ptr())), C.size_t(i), (*C.gsl_vector_long)(unsafe.Pointer(v.Ptr()))))
}

func LongSetCol(m *GslMatrixLong, j int, v *vector.GslVectorLong) int32 {
	return int32(C.gsl_matrix_long_set_col((*C.gsl_matrix_long)(unsafe.Pointer(m.Ptr())), C.size_t(j), (*C.gsl_vector_long)(unsafe.Pointer(v.Ptr()))))
}

func LongSwapRows(m *GslMatrixLong, i int, j int) int32 {
	return int32(C.gsl_matrix_long_swap_rows((*C.gsl_matrix_long)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func LongSwapColumns(m *GslMatrixLong, i int, j int) int32 {
	return int32(C.gsl_matrix_long_swap_columns((*C.gsl_matrix_long)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func LongSwapRowcol(m *GslMatrixLong, i int, j int) int32 {
	return int32(C.gsl_matrix_long_swap_rowcol((*C.gsl_matrix_long)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func LongTransposeMemcpy(dest *GslMatrixLong, src *GslMatrixLong) int32 {
	return int32(C.gsl_matrix_long_transpose_memcpy((*C.gsl_matrix_long)(unsafe.Pointer(dest.Ptr())), (*C.gsl_matrix_long)(unsafe.Pointer(src.Ptr()))))
}

func LongTranspose(m *GslMatrixLong) int32 {
	return int32(C.gsl_matrix_long_transpose((*C.gsl_matrix_long)(unsafe.Pointer(m.Ptr()))))
}

func LongAdd(a *GslMatrixLong, b *GslMatrixLong) int32 {
	return int32(C.gsl_matrix_long_add((*C.gsl_matrix_long)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_long)(unsafe.Pointer(b.Ptr()))))
}

func LongSub(a *GslMatrixLong, b *GslMatrixLong) int32 {
	return int32(C.gsl_matrix_long_sub((*C.gsl_matrix_long)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_long)(unsafe.Pointer(b.Ptr()))))
}

func LongMulElements(a *GslMatrixLong, b *GslMatrixLong) int32 {
	return int32(C.gsl_matrix_long_mul_elements((*C.gsl_matrix_long)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_long)(unsafe.Pointer(b.Ptr()))))
}

func LongDivElements(a *GslMatrixLong, b *GslMatrixLong) int32 {
	return int32(C.gsl_matrix_long_div_elements((*C.gsl_matrix_long)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_long)(unsafe.Pointer(b.Ptr()))))
}

func LongScale(a *GslMatrixLong, x float64) int32 {
	return int32(C.gsl_matrix_long_scale((*C.gsl_matrix_long)(unsafe.Pointer(a.Ptr())), C.double(x)))
}

func LongAddConstant(a *GslMatrixLong, x float64) int32 {
	return int32(C.gsl_matrix_long_add_constant((*C.gsl_matrix_long)(unsafe.Pointer(a.Ptr())), C.double(x)))
}

func LongMax(m *GslMatrixLong) uint {
	return uint(C.gsl_matrix_long_max((*C.gsl_matrix_long)(unsafe.Pointer(m.Ptr()))))
}

func LongMin(m *GslMatrixLong) uint {
	return uint(C.gsl_matrix_long_min((*C.gsl_matrix_long)(unsafe.Pointer(m.Ptr()))))
}

func LongMinmax(m *GslMatrixLong) (uint, uint) {
	var _outptr_1 C.long
	var _outptr_2 C.long
	C.gsl_matrix_long_minmax((*C.gsl_matrix_long)(unsafe.Pointer(m.Ptr())), &_outptr_1, &_outptr_2)
	return *(*uint)(unsafe.Pointer(&_outptr_1)), *(*uint)(unsafe.Pointer(&_outptr_2))
}

func LongMaxIndex(m *GslMatrixLong) (int, int) {
	var _outptr_1 C.size_t
	var _outptr_2 C.size_t
	C.gsl_matrix_long_max_index((*C.gsl_matrix_long)(unsafe.Pointer(m.Ptr())), &_outptr_1, &_outptr_2)
	return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2))
}

func LongMinIndex(m *GslMatrixLong) (int, int) {
	var _outptr_1 C.size_t
	var _outptr_2 C.size_t
	C.gsl_matrix_long_min_index((*C.gsl_matrix_long)(unsafe.Pointer(m.Ptr())), &_outptr_1, &_outptr_2)
	return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2))
}

func LongMinmaxIndex(m *GslMatrixLong) (int, int, int, int) {
	var _outptr_1 C.size_t
	var _outptr_2 C.size_t
	var _outptr_3 C.size_t
	var _outptr_4 C.size_t
	C.gsl_matrix_long_minmax_index((*C.gsl_matrix_long)(unsafe.Pointer(m.Ptr())), &_outptr_1, &_outptr_2, &_outptr_3, &_outptr_4)
	return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2)), *(*int)(unsafe.Pointer(&_outptr_3)), *(*int)(unsafe.Pointer(&_outptr_4))
}

func LongIsnull(m *GslMatrixLong) int32 {
	return int32(C.gsl_matrix_long_isnull((*C.gsl_matrix_long)(unsafe.Pointer(m.Ptr()))))
}

func LongIspos(m *GslMatrixLong) int32 {
	return int32(C.gsl_matrix_long_ispos((*C.gsl_matrix_long)(unsafe.Pointer(m.Ptr()))))
}

func LongIsneg(m *GslMatrixLong) int32 {
	return int32(C.gsl_matrix_long_isneg((*C.gsl_matrix_long)(unsafe.Pointer(m.Ptr()))))
}

func LongIsnonneg(m *GslMatrixLong) int32 {
	return int32(C.gsl_matrix_long_isnonneg((*C.gsl_matrix_long)(unsafe.Pointer(m.Ptr()))))
}

func LongEqual(a *GslMatrixLong, b *GslMatrixLong) int32 {
	return int32(C.gsl_matrix_long_equal((*C.gsl_matrix_long)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_long)(unsafe.Pointer(b.Ptr()))))
}

type GslMatrixUlong struct {
	gogsl.GslReference
}

type GslMatrixUlongView struct {
	gogsl.GslReference
	CData []byte
}

type GslMatrixUlongConstView struct {
	gogsl.GslReference
	CData []byte
}

func MatrixUlongAlloc(n1 int, n2 int) *GslMatrixUlong {
	_ref := C.gsl_matrix_ulong_alloc(C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixUlong{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func MatrixUlongCalloc(n1 int, n2 int) *GslMatrixUlong {
	_ref := C.gsl_matrix_ulong_calloc(C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixUlong{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslMatrixUlong) Dispose() {
	C.gsl_matrix_ulong_free((*C.gsl_matrix_ulong)(unsafe.Pointer(x.Ptr())))
}

func UlongGet(m *GslMatrixUlong, i int, j int) int {
	return int(C.gsl_matrix_ulong_get((*C.gsl_matrix_ulong)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func UlongSet(m *GslMatrixUlong, i int, j int, x int) {
	C.gsl_matrix_ulong_set((*C.gsl_matrix_ulong)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j), C.ulong(x))
}

func UlongSetAll(m *GslMatrixUlong, x int) {
	C.gsl_matrix_ulong_set_all((*C.gsl_matrix_ulong)(unsafe.Pointer(m.Ptr())), C.ulong(x))
}

func UlongSetZero(m *GslMatrixUlong) {
	C.gsl_matrix_ulong_set_zero((*C.gsl_matrix_ulong)(unsafe.Pointer(m.Ptr())))
}

func UlongSetIdentity(m *GslMatrixUlong) {
	C.gsl_matrix_ulong_set_identity((*C.gsl_matrix_ulong)(unsafe.Pointer(m.Ptr())))
}

func UlongFwrite(stream *os.File, m *GslMatrixUlong) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_matrix_ulong_fwrite(_file_0, (*C.gsl_matrix_ulong)(unsafe.Pointer(m.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func UlongFread(stream *os.File, m *GslMatrixUlong) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_matrix_ulong_fread(_file_0, (*C.gsl_matrix_ulong)(unsafe.Pointer(m.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func UlongFprintf(stream *os.File, m *GslMatrixUlong, format string) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_string_2 := C.CString(format)
	_result := int32(C.gsl_matrix_ulong_fprintf(_file_0, (*C.gsl_matrix_ulong)(unsafe.Pointer(m.Ptr())), _string_2))
	C.fclose(_file_0)
	C.free(unsafe.Pointer(_string_2))
	return _result
}

func UlongFscanf(stream *os.File, m *GslMatrixUlong) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_matrix_ulong_fscanf(_file_0, (*C.gsl_matrix_ulong)(unsafe.Pointer(m.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func UlongSubmatrix(m *GslMatrixUlong, k1 int, k2 int, n1 int, n2 int) *GslMatrixUlongView {
	_ref := C.gsl_matrix_ulong_submatrix((*C.gsl_matrix_ulong)(unsafe.Pointer(m.Ptr())), C.size_t(k1), C.size_t(k2), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixUlongView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UlongConstSubmatrix(m *GslMatrixUlong, k1 int, k2 int, n1 int, n2 int) *GslMatrixUlongConstView {
	_ref := C.gsl_matrix_ulong_const_submatrix((*C.gsl_matrix_ulong)(unsafe.Pointer(m.Ptr())), C.size_t(k1), C.size_t(k2), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixUlongConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UlongViewVector(v *vector.GslVectorUlong, n1 int, n2 int) *GslMatrixUlongView {
	_ref := C.gsl_matrix_ulong_view_vector((*C.gsl_vector_ulong)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixUlongView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UlongConstViewVector(v *vector.GslVectorUlong, n1 int, n2 int) *GslMatrixUlongConstView {
	_ref := C.gsl_matrix_ulong_const_view_vector((*C.gsl_vector_ulong)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixUlongConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UlongViewVectorWithTda(v *vector.GslVectorUlong, n1 int, n2 int, tda int) *GslMatrixUlongView {
	_ref := C.gsl_matrix_ulong_view_vector_with_tda((*C.gsl_vector_ulong)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2), C.size_t(tda))
	_result := &GslMatrixUlongView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UlongConstViewVectorWithTda(v *vector.GslVectorUlong, n1 int, n2 int, tda int) *GslMatrixUlongConstView {
	_ref := C.gsl_matrix_ulong_const_view_vector_with_tda((*C.gsl_vector_ulong)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2), C.size_t(tda))
	_result := &GslMatrixUlongConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UlongRow(m *GslMatrixUlong, i int) *vector.GslVectorUlongView {
	_ref := C.gsl_matrix_ulong_row((*C.gsl_matrix_ulong)(unsafe.Pointer(m.Ptr())), C.size_t(i))
	_result := &vector.GslVectorUlongView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UlongConstRow(m *GslMatrixUlong, i int) *vector.GslVectorUlongConstView {
	_ref := C.gsl_matrix_ulong_const_row((*C.gsl_matrix_ulong)(unsafe.Pointer(m.Ptr())), C.size_t(i))
	_result := &vector.GslVectorUlongConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UlongColumn(m *GslMatrixUlong, j int) *vector.GslVectorUlongView {
	_ref := C.gsl_matrix_ulong_column((*C.gsl_matrix_ulong)(unsafe.Pointer(m.Ptr())), C.size_t(j))
	_result := &vector.GslVectorUlongView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UlongConstColumn(m *GslMatrixUlong, j int) *vector.GslVectorUlongConstView {
	_ref := C.gsl_matrix_ulong_const_column((*C.gsl_matrix_ulong)(unsafe.Pointer(m.Ptr())), C.size_t(j))
	_result := &vector.GslVectorUlongConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UlongSubrow(m *GslMatrixUlong, i int, offset int, n int) *vector.GslVectorUlongView {
	_ref := C.gsl_matrix_ulong_subrow((*C.gsl_matrix_ulong)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorUlongView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UlongConstSubrow(m *GslMatrixUlong, i int, offset int, n int) *vector.GslVectorUlongConstView {
	_ref := C.gsl_matrix_ulong_const_subrow((*C.gsl_matrix_ulong)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorUlongConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UlongSubcolumn(m *GslMatrixUlong, j int, offset int, n int) *vector.GslVectorUlongView {
	_ref := C.gsl_matrix_ulong_subcolumn((*C.gsl_matrix_ulong)(unsafe.Pointer(m.Ptr())), C.size_t(j), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorUlongView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UlongConstSubcolumn(m *GslMatrixUlong, j int, offset int, n int) *vector.GslVectorUlongConstView {
	_ref := C.gsl_matrix_ulong_const_subcolumn((*C.gsl_matrix_ulong)(unsafe.Pointer(m.Ptr())), C.size_t(j), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorUlongConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UlongDiagonal(m *GslMatrixUlong) *vector.GslVectorUlongView {
	_ref := C.gsl_matrix_ulong_diagonal((*C.gsl_matrix_ulong)(unsafe.Pointer(m.Ptr())))
	_result := &vector.GslVectorUlongView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UlongConstDiagonal(m *GslMatrixUlong) *vector.GslVectorUlongConstView {
	_ref := C.gsl_matrix_ulong_const_diagonal((*C.gsl_matrix_ulong)(unsafe.Pointer(m.Ptr())))
	_result := &vector.GslVectorUlongConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UlongSubdiagonal(m *GslMatrixUlong, k int) *vector.GslVectorUlongView {
	_ref := C.gsl_matrix_ulong_subdiagonal((*C.gsl_matrix_ulong)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorUlongView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UlongConstSubdiagonal(m *GslMatrixUlong, k int) *vector.GslVectorUlongConstView {
	_ref := C.gsl_matrix_ulong_const_subdiagonal((*C.gsl_matrix_ulong)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorUlongConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UlongSuperdiagonal(m *GslMatrixUlong, k int) *vector.GslVectorUlongView {
	_ref := C.gsl_matrix_ulong_superdiagonal((*C.gsl_matrix_ulong)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorUlongView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UlongConstSuperdiagonal(m *GslMatrixUlong, k int) *vector.GslVectorUlongConstView {
	_ref := C.gsl_matrix_ulong_const_superdiagonal((*C.gsl_matrix_ulong)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorUlongConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UlongMemcpy(dest *GslMatrixUlong, src *GslMatrixUlong) int32 {
	return int32(C.gsl_matrix_ulong_memcpy((*C.gsl_matrix_ulong)(unsafe.Pointer(dest.Ptr())), (*C.gsl_matrix_ulong)(unsafe.Pointer(src.Ptr()))))
}

func UlongSwap(m1 *GslMatrixUlong, m2 *GslMatrixUlong) int32 {
	return int32(C.gsl_matrix_ulong_swap((*C.gsl_matrix_ulong)(unsafe.Pointer(m1.Ptr())), (*C.gsl_matrix_ulong)(unsafe.Pointer(m2.Ptr()))))
}

func UlongGetRow(v *vector.GslVectorUlong, m *GslMatrixUlong, i int) int32 {
	return int32(C.gsl_matrix_ulong_get_row((*C.gsl_vector_ulong)(unsafe.Pointer(v.Ptr())), (*C.gsl_matrix_ulong)(unsafe.Pointer(m.Ptr())), C.size_t(i)))
}

func UlongGetCol(v *vector.GslVectorUlong, m *GslMatrixUlong, j int) int32 {
	return int32(C.gsl_matrix_ulong_get_col((*C.gsl_vector_ulong)(unsafe.Pointer(v.Ptr())), (*C.gsl_matrix_ulong)(unsafe.Pointer(m.Ptr())), C.size_t(j)))
}

func UlongSetRow(m *GslMatrixUlong, i int, v *vector.GslVectorUlong) int32 {
	return int32(C.gsl_matrix_ulong_set_row((*C.gsl_matrix_ulong)(unsafe.Pointer(m.Ptr())), C.size_t(i), (*C.gsl_vector_ulong)(unsafe.Pointer(v.Ptr()))))
}

func UlongSetCol(m *GslMatrixUlong, j int, v *vector.GslVectorUlong) int32 {
	return int32(C.gsl_matrix_ulong_set_col((*C.gsl_matrix_ulong)(unsafe.Pointer(m.Ptr())), C.size_t(j), (*C.gsl_vector_ulong)(unsafe.Pointer(v.Ptr()))))
}

func UlongSwapRows(m *GslMatrixUlong, i int, j int) int32 {
	return int32(C.gsl_matrix_ulong_swap_rows((*C.gsl_matrix_ulong)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func UlongSwapColumns(m *GslMatrixUlong, i int, j int) int32 {
	return int32(C.gsl_matrix_ulong_swap_columns((*C.gsl_matrix_ulong)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func UlongSwapRowcol(m *GslMatrixUlong, i int, j int) int32 {
	return int32(C.gsl_matrix_ulong_swap_rowcol((*C.gsl_matrix_ulong)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func UlongTransposeMemcpy(dest *GslMatrixUlong, src *GslMatrixUlong) int32 {
	return int32(C.gsl_matrix_ulong_transpose_memcpy((*C.gsl_matrix_ulong)(unsafe.Pointer(dest.Ptr())), (*C.gsl_matrix_ulong)(unsafe.Pointer(src.Ptr()))))
}

func UlongTranspose(m *GslMatrixUlong) int32 {
	return int32(C.gsl_matrix_ulong_transpose((*C.gsl_matrix_ulong)(unsafe.Pointer(m.Ptr()))))
}

func UlongAdd(a *GslMatrixUlong, b *GslMatrixUlong) int32 {
	return int32(C.gsl_matrix_ulong_add((*C.gsl_matrix_ulong)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_ulong)(unsafe.Pointer(b.Ptr()))))
}

func UlongSub(a *GslMatrixUlong, b *GslMatrixUlong) int32 {
	return int32(C.gsl_matrix_ulong_sub((*C.gsl_matrix_ulong)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_ulong)(unsafe.Pointer(b.Ptr()))))
}

func UlongMulElements(a *GslMatrixUlong, b *GslMatrixUlong) int32 {
	return int32(C.gsl_matrix_ulong_mul_elements((*C.gsl_matrix_ulong)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_ulong)(unsafe.Pointer(b.Ptr()))))
}

func UlongDivElements(a *GslMatrixUlong, b *GslMatrixUlong) int32 {
	return int32(C.gsl_matrix_ulong_div_elements((*C.gsl_matrix_ulong)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_ulong)(unsafe.Pointer(b.Ptr()))))
}

func UlongScale(a *GslMatrixUlong, x float64) int32 {
	return int32(C.gsl_matrix_ulong_scale((*C.gsl_matrix_ulong)(unsafe.Pointer(a.Ptr())), C.double(x)))
}

func UlongAddConstant(a *GslMatrixUlong, x float64) int32 {
	return int32(C.gsl_matrix_ulong_add_constant((*C.gsl_matrix_ulong)(unsafe.Pointer(a.Ptr())), C.double(x)))
}

func UlongMax(m *GslMatrixUlong) int {
	return int(C.gsl_matrix_ulong_max((*C.gsl_matrix_ulong)(unsafe.Pointer(m.Ptr()))))
}

func UlongMin(m *GslMatrixUlong) int {
	return int(C.gsl_matrix_ulong_min((*C.gsl_matrix_ulong)(unsafe.Pointer(m.Ptr()))))
}

func UlongMinmax(m *GslMatrixUlong) (int, int) {
	var _outptr_1 C.ulong
	var _outptr_2 C.ulong
	C.gsl_matrix_ulong_minmax((*C.gsl_matrix_ulong)(unsafe.Pointer(m.Ptr())), &_outptr_1, &_outptr_2)
	return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2))
}

func UlongMaxIndex(m *GslMatrixUlong) (int, int) {
	var _outptr_1 C.size_t
	var _outptr_2 C.size_t
	C.gsl_matrix_ulong_max_index((*C.gsl_matrix_ulong)(unsafe.Pointer(m.Ptr())), &_outptr_1, &_outptr_2)
	return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2))
}

func UlongMinIndex(m *GslMatrixUlong) (int, int) {
	var _outptr_1 C.size_t
	var _outptr_2 C.size_t
	C.gsl_matrix_ulong_min_index((*C.gsl_matrix_ulong)(unsafe.Pointer(m.Ptr())), &_outptr_1, &_outptr_2)
	return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2))
}

func UlongMinmaxIndex(m *GslMatrixUlong) (int, int, int, int) {
	var _outptr_1 C.size_t
	var _outptr_2 C.size_t
	var _outptr_3 C.size_t
	var _outptr_4 C.size_t
	C.gsl_matrix_ulong_minmax_index((*C.gsl_matrix_ulong)(unsafe.Pointer(m.Ptr())), &_outptr_1, &_outptr_2, &_outptr_3, &_outptr_4)
	return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2)), *(*int)(unsafe.Pointer(&_outptr_3)), *(*int)(unsafe.Pointer(&_outptr_4))
}

func UlongIsnull(m *GslMatrixUlong) int32 {
	return int32(C.gsl_matrix_ulong_isnull((*C.gsl_matrix_ulong)(unsafe.Pointer(m.Ptr()))))
}

func UlongIspos(m *GslMatrixUlong) int32 {
	return int32(C.gsl_matrix_ulong_ispos((*C.gsl_matrix_ulong)(unsafe.Pointer(m.Ptr()))))
}

func UlongIsneg(m *GslMatrixUlong) int32 {
	return int32(C.gsl_matrix_ulong_isneg((*C.gsl_matrix_ulong)(unsafe.Pointer(m.Ptr()))))
}

func UlongIsnonneg(m *GslMatrixUlong) int32 {
	return int32(C.gsl_matrix_ulong_isnonneg((*C.gsl_matrix_ulong)(unsafe.Pointer(m.Ptr()))))
}

func UlongEqual(a *GslMatrixUlong, b *GslMatrixUlong) int32 {
	return int32(C.gsl_matrix_ulong_equal((*C.gsl_matrix_ulong)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_ulong)(unsafe.Pointer(b.Ptr()))))
}

type GslMatrixShort struct {
	gogsl.GslReference
}

type GslMatrixShortView struct {
	gogsl.GslReference
	CData []byte
}

type GslMatrixShortConstView struct {
	gogsl.GslReference
	CData []byte
}

func MatrixShortAlloc(n1 int, n2 int) *GslMatrixShort {
	_ref := C.gsl_matrix_short_alloc(C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixShort{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func MatrixShortCalloc(n1 int, n2 int) *GslMatrixShort {
	_ref := C.gsl_matrix_short_calloc(C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixShort{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslMatrixShort) Dispose() {
	C.gsl_matrix_short_free((*C.gsl_matrix_short)(unsafe.Pointer(x.Ptr())))
}

func ShortGet(m *GslMatrixShort, i int, j int) int16 {
	return int16(C.gsl_matrix_short_get((*C.gsl_matrix_short)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func ShortSet(m *GslMatrixShort, i int, j int, x int16) {
	C.gsl_matrix_short_set((*C.gsl_matrix_short)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j), C.short(x))
}

func ShortSetAll(m *GslMatrixShort, x int16) {
	C.gsl_matrix_short_set_all((*C.gsl_matrix_short)(unsafe.Pointer(m.Ptr())), C.short(x))
}

func ShortSetZero(m *GslMatrixShort) {
	C.gsl_matrix_short_set_zero((*C.gsl_matrix_short)(unsafe.Pointer(m.Ptr())))
}

func ShortSetIdentity(m *GslMatrixShort) {
	C.gsl_matrix_short_set_identity((*C.gsl_matrix_short)(unsafe.Pointer(m.Ptr())))
}

func ShortFwrite(stream *os.File, m *GslMatrixShort) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_matrix_short_fwrite(_file_0, (*C.gsl_matrix_short)(unsafe.Pointer(m.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func ShortFread(stream *os.File, m *GslMatrixShort) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_matrix_short_fread(_file_0, (*C.gsl_matrix_short)(unsafe.Pointer(m.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func ShortFprintf(stream *os.File, m *GslMatrixShort, format string) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_string_2 := C.CString(format)
	_result := int32(C.gsl_matrix_short_fprintf(_file_0, (*C.gsl_matrix_short)(unsafe.Pointer(m.Ptr())), _string_2))
	C.fclose(_file_0)
	C.free(unsafe.Pointer(_string_2))
	return _result
}

func ShortFscanf(stream *os.File, m *GslMatrixShort) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_matrix_short_fscanf(_file_0, (*C.gsl_matrix_short)(unsafe.Pointer(m.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func ShortSubmatrix(m *GslMatrixShort, k1 int, k2 int, n1 int, n2 int) *GslMatrixShortView {
	_ref := C.gsl_matrix_short_submatrix((*C.gsl_matrix_short)(unsafe.Pointer(m.Ptr())), C.size_t(k1), C.size_t(k2), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixShortView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ShortConstSubmatrix(m *GslMatrixShort, k1 int, k2 int, n1 int, n2 int) *GslMatrixShortConstView {
	_ref := C.gsl_matrix_short_const_submatrix((*C.gsl_matrix_short)(unsafe.Pointer(m.Ptr())), C.size_t(k1), C.size_t(k2), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixShortConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ShortViewVector(v *vector.GslVectorShort, n1 int, n2 int) *GslMatrixShortView {
	_ref := C.gsl_matrix_short_view_vector((*C.gsl_vector_short)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixShortView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ShortConstViewVector(v *vector.GslVectorShort, n1 int, n2 int) *GslMatrixShortConstView {
	_ref := C.gsl_matrix_short_const_view_vector((*C.gsl_vector_short)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixShortConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ShortViewVectorWithTda(v *vector.GslVectorShort, n1 int, n2 int, tda int) *GslMatrixShortView {
	_ref := C.gsl_matrix_short_view_vector_with_tda((*C.gsl_vector_short)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2), C.size_t(tda))
	_result := &GslMatrixShortView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ShortConstViewVectorWithTda(v *vector.GslVectorShort, n1 int, n2 int, tda int) *GslMatrixShortConstView {
	_ref := C.gsl_matrix_short_const_view_vector_with_tda((*C.gsl_vector_short)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2), C.size_t(tda))
	_result := &GslMatrixShortConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ShortRow(m *GslMatrixShort, i int) *vector.GslVectorShortView {
	_ref := C.gsl_matrix_short_row((*C.gsl_matrix_short)(unsafe.Pointer(m.Ptr())), C.size_t(i))
	_result := &vector.GslVectorShortView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ShortConstRow(m *GslMatrixShort, i int) *vector.GslVectorShortConstView {
	_ref := C.gsl_matrix_short_const_row((*C.gsl_matrix_short)(unsafe.Pointer(m.Ptr())), C.size_t(i))
	_result := &vector.GslVectorShortConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ShortColumn(m *GslMatrixShort, j int) *vector.GslVectorShortView {
	_ref := C.gsl_matrix_short_column((*C.gsl_matrix_short)(unsafe.Pointer(m.Ptr())), C.size_t(j))
	_result := &vector.GslVectorShortView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ShortConstColumn(m *GslMatrixShort, j int) *vector.GslVectorShortConstView {
	_ref := C.gsl_matrix_short_const_column((*C.gsl_matrix_short)(unsafe.Pointer(m.Ptr())), C.size_t(j))
	_result := &vector.GslVectorShortConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ShortSubrow(m *GslMatrixShort, i int, offset int, n int) *vector.GslVectorShortView {
	_ref := C.gsl_matrix_short_subrow((*C.gsl_matrix_short)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorShortView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ShortConstSubrow(m *GslMatrixShort, i int, offset int, n int) *vector.GslVectorShortConstView {
	_ref := C.gsl_matrix_short_const_subrow((*C.gsl_matrix_short)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorShortConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ShortSubcolumn(m *GslMatrixShort, j int, offset int, n int) *vector.GslVectorShortView {
	_ref := C.gsl_matrix_short_subcolumn((*C.gsl_matrix_short)(unsafe.Pointer(m.Ptr())), C.size_t(j), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorShortView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ShortConstSubcolumn(m *GslMatrixShort, j int, offset int, n int) *vector.GslVectorShortConstView {
	_ref := C.gsl_matrix_short_const_subcolumn((*C.gsl_matrix_short)(unsafe.Pointer(m.Ptr())), C.size_t(j), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorShortConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ShortDiagonal(m *GslMatrixShort) *vector.GslVectorShortView {
	_ref := C.gsl_matrix_short_diagonal((*C.gsl_matrix_short)(unsafe.Pointer(m.Ptr())))
	_result := &vector.GslVectorShortView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ShortConstDiagonal(m *GslMatrixShort) *vector.GslVectorShortConstView {
	_ref := C.gsl_matrix_short_const_diagonal((*C.gsl_matrix_short)(unsafe.Pointer(m.Ptr())))
	_result := &vector.GslVectorShortConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ShortSubdiagonal(m *GslMatrixShort, k int) *vector.GslVectorShortView {
	_ref := C.gsl_matrix_short_subdiagonal((*C.gsl_matrix_short)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorShortView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ShortConstSubdiagonal(m *GslMatrixShort, k int) *vector.GslVectorShortConstView {
	_ref := C.gsl_matrix_short_const_subdiagonal((*C.gsl_matrix_short)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorShortConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ShortSuperdiagonal(m *GslMatrixShort, k int) *vector.GslVectorShortView {
	_ref := C.gsl_matrix_short_superdiagonal((*C.gsl_matrix_short)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorShortView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ShortConstSuperdiagonal(m *GslMatrixShort, k int) *vector.GslVectorShortConstView {
	_ref := C.gsl_matrix_short_const_superdiagonal((*C.gsl_matrix_short)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorShortConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ShortMemcpy(dest *GslMatrixShort, src *GslMatrixShort) int32 {
	return int32(C.gsl_matrix_short_memcpy((*C.gsl_matrix_short)(unsafe.Pointer(dest.Ptr())), (*C.gsl_matrix_short)(unsafe.Pointer(src.Ptr()))))
}

func ShortSwap(m1 *GslMatrixShort, m2 *GslMatrixShort) int32 {
	return int32(C.gsl_matrix_short_swap((*C.gsl_matrix_short)(unsafe.Pointer(m1.Ptr())), (*C.gsl_matrix_short)(unsafe.Pointer(m2.Ptr()))))
}

func ShortGetRow(v *vector.GslVectorShort, m *GslMatrixShort, i int) int32 {
	return int32(C.gsl_matrix_short_get_row((*C.gsl_vector_short)(unsafe.Pointer(v.Ptr())), (*C.gsl_matrix_short)(unsafe.Pointer(m.Ptr())), C.size_t(i)))
}

func ShortGetCol(v *vector.GslVectorShort, m *GslMatrixShort, j int) int32 {
	return int32(C.gsl_matrix_short_get_col((*C.gsl_vector_short)(unsafe.Pointer(v.Ptr())), (*C.gsl_matrix_short)(unsafe.Pointer(m.Ptr())), C.size_t(j)))
}

func ShortSetRow(m *GslMatrixShort, i int, v *vector.GslVectorShort) int32 {
	return int32(C.gsl_matrix_short_set_row((*C.gsl_matrix_short)(unsafe.Pointer(m.Ptr())), C.size_t(i), (*C.gsl_vector_short)(unsafe.Pointer(v.Ptr()))))
}

func ShortSetCol(m *GslMatrixShort, j int, v *vector.GslVectorShort) int32 {
	return int32(C.gsl_matrix_short_set_col((*C.gsl_matrix_short)(unsafe.Pointer(m.Ptr())), C.size_t(j), (*C.gsl_vector_short)(unsafe.Pointer(v.Ptr()))))
}

func ShortSwapRows(m *GslMatrixShort, i int, j int) int32 {
	return int32(C.gsl_matrix_short_swap_rows((*C.gsl_matrix_short)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func ShortSwapColumns(m *GslMatrixShort, i int, j int) int32 {
	return int32(C.gsl_matrix_short_swap_columns((*C.gsl_matrix_short)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func ShortSwapRowcol(m *GslMatrixShort, i int, j int) int32 {
	return int32(C.gsl_matrix_short_swap_rowcol((*C.gsl_matrix_short)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func ShortTransposeMemcpy(dest *GslMatrixShort, src *GslMatrixShort) int32 {
	return int32(C.gsl_matrix_short_transpose_memcpy((*C.gsl_matrix_short)(unsafe.Pointer(dest.Ptr())), (*C.gsl_matrix_short)(unsafe.Pointer(src.Ptr()))))
}

func ShortTranspose(m *GslMatrixShort) int32 {
	return int32(C.gsl_matrix_short_transpose((*C.gsl_matrix_short)(unsafe.Pointer(m.Ptr()))))
}

func ShortAdd(a *GslMatrixShort, b *GslMatrixShort) int32 {
	return int32(C.gsl_matrix_short_add((*C.gsl_matrix_short)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_short)(unsafe.Pointer(b.Ptr()))))
}

func ShortSub(a *GslMatrixShort, b *GslMatrixShort) int32 {
	return int32(C.gsl_matrix_short_sub((*C.gsl_matrix_short)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_short)(unsafe.Pointer(b.Ptr()))))
}

func ShortMulElements(a *GslMatrixShort, b *GslMatrixShort) int32 {
	return int32(C.gsl_matrix_short_mul_elements((*C.gsl_matrix_short)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_short)(unsafe.Pointer(b.Ptr()))))
}

func ShortDivElements(a *GslMatrixShort, b *GslMatrixShort) int32 {
	return int32(C.gsl_matrix_short_div_elements((*C.gsl_matrix_short)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_short)(unsafe.Pointer(b.Ptr()))))
}

func ShortScale(a *GslMatrixShort, x float64) int32 {
	return int32(C.gsl_matrix_short_scale((*C.gsl_matrix_short)(unsafe.Pointer(a.Ptr())), C.double(x)))
}

func ShortAddConstant(a *GslMatrixShort, x float64) int32 {
	return int32(C.gsl_matrix_short_add_constant((*C.gsl_matrix_short)(unsafe.Pointer(a.Ptr())), C.double(x)))
}

func ShortMax(m *GslMatrixShort) int16 {
	return int16(C.gsl_matrix_short_max((*C.gsl_matrix_short)(unsafe.Pointer(m.Ptr()))))
}

func ShortMin(m *GslMatrixShort) int16 {
	return int16(C.gsl_matrix_short_min((*C.gsl_matrix_short)(unsafe.Pointer(m.Ptr()))))
}

func ShortMinmax(m *GslMatrixShort) (int16, int16) {
	var _outptr_1 C.short
	var _outptr_2 C.short
	C.gsl_matrix_short_minmax((*C.gsl_matrix_short)(unsafe.Pointer(m.Ptr())), &_outptr_1, &_outptr_2)
	return *(*int16)(unsafe.Pointer(&_outptr_1)), *(*int16)(unsafe.Pointer(&_outptr_2))
}

func ShortMaxIndex(m *GslMatrixShort) (int, int) {
	var _outptr_1 C.size_t
	var _outptr_2 C.size_t
	C.gsl_matrix_short_max_index((*C.gsl_matrix_short)(unsafe.Pointer(m.Ptr())), &_outptr_1, &_outptr_2)
	return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2))
}

func ShortMinIndex(m *GslMatrixShort) (int, int) {
	var _outptr_1 C.size_t
	var _outptr_2 C.size_t
	C.gsl_matrix_short_min_index((*C.gsl_matrix_short)(unsafe.Pointer(m.Ptr())), &_outptr_1, &_outptr_2)
	return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2))
}

func ShortMinmaxIndex(m *GslMatrixShort) (int, int, int, int) {
	var _outptr_1 C.size_t
	var _outptr_2 C.size_t
	var _outptr_3 C.size_t
	var _outptr_4 C.size_t
	C.gsl_matrix_short_minmax_index((*C.gsl_matrix_short)(unsafe.Pointer(m.Ptr())), &_outptr_1, &_outptr_2, &_outptr_3, &_outptr_4)
	return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2)), *(*int)(unsafe.Pointer(&_outptr_3)), *(*int)(unsafe.Pointer(&_outptr_4))
}

func ShortIsnull(m *GslMatrixShort) int32 {
	return int32(C.gsl_matrix_short_isnull((*C.gsl_matrix_short)(unsafe.Pointer(m.Ptr()))))
}

func ShortIspos(m *GslMatrixShort) int32 {
	return int32(C.gsl_matrix_short_ispos((*C.gsl_matrix_short)(unsafe.Pointer(m.Ptr()))))
}

func ShortIsneg(m *GslMatrixShort) int32 {
	return int32(C.gsl_matrix_short_isneg((*C.gsl_matrix_short)(unsafe.Pointer(m.Ptr()))))
}

func ShortIsnonneg(m *GslMatrixShort) int32 {
	return int32(C.gsl_matrix_short_isnonneg((*C.gsl_matrix_short)(unsafe.Pointer(m.Ptr()))))
}

func ShortEqual(a *GslMatrixShort, b *GslMatrixShort) int32 {
	return int32(C.gsl_matrix_short_equal((*C.gsl_matrix_short)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_short)(unsafe.Pointer(b.Ptr()))))
}

type GslMatrixUshort struct {
	gogsl.GslReference
}

type GslMatrixUshortView struct {
	gogsl.GslReference
	CData []byte
}

type GslMatrixUshortConstView struct {
	gogsl.GslReference
	CData []byte
}

func MatrixUshortAlloc(n1 int, n2 int) *GslMatrixUshort {
	_ref := C.gsl_matrix_ushort_alloc(C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixUshort{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func MatrixUshortCalloc(n1 int, n2 int) *GslMatrixUshort {
	_ref := C.gsl_matrix_ushort_calloc(C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixUshort{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslMatrixUshort) Dispose() {
	C.gsl_matrix_ushort_free((*C.gsl_matrix_ushort)(unsafe.Pointer(x.Ptr())))
}

func UshortGet(m *GslMatrixUshort, i int, j int) int16 {
	return int16(C.gsl_matrix_ushort_get((*C.gsl_matrix_ushort)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func UshortSet(m *GslMatrixUshort, i int, j int, x int16) {
	C.gsl_matrix_ushort_set((*C.gsl_matrix_ushort)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j), C.ushort(x))
}

func UshortSetAll(m *GslMatrixUshort, x int16) {
	C.gsl_matrix_ushort_set_all((*C.gsl_matrix_ushort)(unsafe.Pointer(m.Ptr())), C.ushort(x))
}

func UshortSetZero(m *GslMatrixUshort) {
	C.gsl_matrix_ushort_set_zero((*C.gsl_matrix_ushort)(unsafe.Pointer(m.Ptr())))
}

func UshortSetIdentity(m *GslMatrixUshort) {
	C.gsl_matrix_ushort_set_identity((*C.gsl_matrix_ushort)(unsafe.Pointer(m.Ptr())))
}

func UshortFwrite(stream *os.File, m *GslMatrixUshort) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_matrix_ushort_fwrite(_file_0, (*C.gsl_matrix_ushort)(unsafe.Pointer(m.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func UshortFread(stream *os.File, m *GslMatrixUshort) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_matrix_ushort_fread(_file_0, (*C.gsl_matrix_ushort)(unsafe.Pointer(m.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func UshortFprintf(stream *os.File, m *GslMatrixUshort, format string) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_string_2 := C.CString(format)
	_result := int32(C.gsl_matrix_ushort_fprintf(_file_0, (*C.gsl_matrix_ushort)(unsafe.Pointer(m.Ptr())), _string_2))
	C.fclose(_file_0)
	C.free(unsafe.Pointer(_string_2))
	return _result
}

func UshortFscanf(stream *os.File, m *GslMatrixUshort) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_matrix_ushort_fscanf(_file_0, (*C.gsl_matrix_ushort)(unsafe.Pointer(m.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func UshortSubmatrix(m *GslMatrixUshort, k1 int, k2 int, n1 int, n2 int) *GslMatrixUshortView {
	_ref := C.gsl_matrix_ushort_submatrix((*C.gsl_matrix_ushort)(unsafe.Pointer(m.Ptr())), C.size_t(k1), C.size_t(k2), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixUshortView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UshortConstSubmatrix(m *GslMatrixUshort, k1 int, k2 int, n1 int, n2 int) *GslMatrixUshortConstView {
	_ref := C.gsl_matrix_ushort_const_submatrix((*C.gsl_matrix_ushort)(unsafe.Pointer(m.Ptr())), C.size_t(k1), C.size_t(k2), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixUshortConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UshortViewVector(v *vector.GslVectorUshort, n1 int, n2 int) *GslMatrixUshortView {
	_ref := C.gsl_matrix_ushort_view_vector((*C.gsl_vector_ushort)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixUshortView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UshortConstViewVector(v *vector.GslVectorUshort, n1 int, n2 int) *GslMatrixUshortConstView {
	_ref := C.gsl_matrix_ushort_const_view_vector((*C.gsl_vector_ushort)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixUshortConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UshortViewVectorWithTda(v *vector.GslVectorUshort, n1 int, n2 int, tda int) *GslMatrixUshortView {
	_ref := C.gsl_matrix_ushort_view_vector_with_tda((*C.gsl_vector_ushort)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2), C.size_t(tda))
	_result := &GslMatrixUshortView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UshortConstViewVectorWithTda(v *vector.GslVectorUshort, n1 int, n2 int, tda int) *GslMatrixUshortConstView {
	_ref := C.gsl_matrix_ushort_const_view_vector_with_tda((*C.gsl_vector_ushort)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2), C.size_t(tda))
	_result := &GslMatrixUshortConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UshortRow(m *GslMatrixUshort, i int) *vector.GslVectorUshortView {
	_ref := C.gsl_matrix_ushort_row((*C.gsl_matrix_ushort)(unsafe.Pointer(m.Ptr())), C.size_t(i))
	_result := &vector.GslVectorUshortView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UshortConstRow(m *GslMatrixUshort, i int) *vector.GslVectorUshortConstView {
	_ref := C.gsl_matrix_ushort_const_row((*C.gsl_matrix_ushort)(unsafe.Pointer(m.Ptr())), C.size_t(i))
	_result := &vector.GslVectorUshortConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UshortColumn(m *GslMatrixUshort, j int) *vector.GslVectorUshortView {
	_ref := C.gsl_matrix_ushort_column((*C.gsl_matrix_ushort)(unsafe.Pointer(m.Ptr())), C.size_t(j))
	_result := &vector.GslVectorUshortView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UshortConstColumn(m *GslMatrixUshort, j int) *vector.GslVectorUshortConstView {
	_ref := C.gsl_matrix_ushort_const_column((*C.gsl_matrix_ushort)(unsafe.Pointer(m.Ptr())), C.size_t(j))
	_result := &vector.GslVectorUshortConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UshortSubrow(m *GslMatrixUshort, i int, offset int, n int) *vector.GslVectorUshortView {
	_ref := C.gsl_matrix_ushort_subrow((*C.gsl_matrix_ushort)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorUshortView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UshortConstSubrow(m *GslMatrixUshort, i int, offset int, n int) *vector.GslVectorUshortConstView {
	_ref := C.gsl_matrix_ushort_const_subrow((*C.gsl_matrix_ushort)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorUshortConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UshortSubcolumn(m *GslMatrixUshort, j int, offset int, n int) *vector.GslVectorUshortView {
	_ref := C.gsl_matrix_ushort_subcolumn((*C.gsl_matrix_ushort)(unsafe.Pointer(m.Ptr())), C.size_t(j), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorUshortView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UshortConstSubcolumn(m *GslMatrixUshort, j int, offset int, n int) *vector.GslVectorUshortConstView {
	_ref := C.gsl_matrix_ushort_const_subcolumn((*C.gsl_matrix_ushort)(unsafe.Pointer(m.Ptr())), C.size_t(j), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorUshortConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UshortDiagonal(m *GslMatrixUshort) *vector.GslVectorUshortView {
	_ref := C.gsl_matrix_ushort_diagonal((*C.gsl_matrix_ushort)(unsafe.Pointer(m.Ptr())))
	_result := &vector.GslVectorUshortView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UshortConstDiagonal(m *GslMatrixUshort) *vector.GslVectorUshortConstView {
	_ref := C.gsl_matrix_ushort_const_diagonal((*C.gsl_matrix_ushort)(unsafe.Pointer(m.Ptr())))
	_result := &vector.GslVectorUshortConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UshortSubdiagonal(m *GslMatrixUshort, k int) *vector.GslVectorUshortView {
	_ref := C.gsl_matrix_ushort_subdiagonal((*C.gsl_matrix_ushort)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorUshortView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UshortConstSubdiagonal(m *GslMatrixUshort, k int) *vector.GslVectorUshortConstView {
	_ref := C.gsl_matrix_ushort_const_subdiagonal((*C.gsl_matrix_ushort)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorUshortConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UshortSuperdiagonal(m *GslMatrixUshort, k int) *vector.GslVectorUshortView {
	_ref := C.gsl_matrix_ushort_superdiagonal((*C.gsl_matrix_ushort)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorUshortView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UshortConstSuperdiagonal(m *GslMatrixUshort, k int) *vector.GslVectorUshortConstView {
	_ref := C.gsl_matrix_ushort_const_superdiagonal((*C.gsl_matrix_ushort)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorUshortConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UshortMemcpy(dest *GslMatrixUshort, src *GslMatrixUshort) int32 {
	return int32(C.gsl_matrix_ushort_memcpy((*C.gsl_matrix_ushort)(unsafe.Pointer(dest.Ptr())), (*C.gsl_matrix_ushort)(unsafe.Pointer(src.Ptr()))))
}

func UshortSwap(m1 *GslMatrixUshort, m2 *GslMatrixUshort) int32 {
	return int32(C.gsl_matrix_ushort_swap((*C.gsl_matrix_ushort)(unsafe.Pointer(m1.Ptr())), (*C.gsl_matrix_ushort)(unsafe.Pointer(m2.Ptr()))))
}

func UshortGetRow(v *vector.GslVectorUshort, m *GslMatrixUshort, i int) int32 {
	return int32(C.gsl_matrix_ushort_get_row((*C.gsl_vector_ushort)(unsafe.Pointer(v.Ptr())), (*C.gsl_matrix_ushort)(unsafe.Pointer(m.Ptr())), C.size_t(i)))
}

func UshortGetCol(v *vector.GslVectorUshort, m *GslMatrixUshort, j int) int32 {
	return int32(C.gsl_matrix_ushort_get_col((*C.gsl_vector_ushort)(unsafe.Pointer(v.Ptr())), (*C.gsl_matrix_ushort)(unsafe.Pointer(m.Ptr())), C.size_t(j)))
}

func UshortSetRow(m *GslMatrixUshort, i int, v *vector.GslVectorUshort) int32 {
	return int32(C.gsl_matrix_ushort_set_row((*C.gsl_matrix_ushort)(unsafe.Pointer(m.Ptr())), C.size_t(i), (*C.gsl_vector_ushort)(unsafe.Pointer(v.Ptr()))))
}

func UshortSetCol(m *GslMatrixUshort, j int, v *vector.GslVectorUshort) int32 {
	return int32(C.gsl_matrix_ushort_set_col((*C.gsl_matrix_ushort)(unsafe.Pointer(m.Ptr())), C.size_t(j), (*C.gsl_vector_ushort)(unsafe.Pointer(v.Ptr()))))
}

func UshortSwapRows(m *GslMatrixUshort, i int, j int) int32 {
	return int32(C.gsl_matrix_ushort_swap_rows((*C.gsl_matrix_ushort)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func UshortSwapColumns(m *GslMatrixUshort, i int, j int) int32 {
	return int32(C.gsl_matrix_ushort_swap_columns((*C.gsl_matrix_ushort)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func UshortSwapRowcol(m *GslMatrixUshort, i int, j int) int32 {
	return int32(C.gsl_matrix_ushort_swap_rowcol((*C.gsl_matrix_ushort)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func UshortTransposeMemcpy(dest *GslMatrixUshort, src *GslMatrixUshort) int32 {
	return int32(C.gsl_matrix_ushort_transpose_memcpy((*C.gsl_matrix_ushort)(unsafe.Pointer(dest.Ptr())), (*C.gsl_matrix_ushort)(unsafe.Pointer(src.Ptr()))))
}

func UshortTranspose(m *GslMatrixUshort) int32 {
	return int32(C.gsl_matrix_ushort_transpose((*C.gsl_matrix_ushort)(unsafe.Pointer(m.Ptr()))))
}

func UshortAdd(a *GslMatrixUshort, b *GslMatrixUshort) int32 {
	return int32(C.gsl_matrix_ushort_add((*C.gsl_matrix_ushort)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_ushort)(unsafe.Pointer(b.Ptr()))))
}

func UshortSub(a *GslMatrixUshort, b *GslMatrixUshort) int32 {
	return int32(C.gsl_matrix_ushort_sub((*C.gsl_matrix_ushort)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_ushort)(unsafe.Pointer(b.Ptr()))))
}

func UshortMulElements(a *GslMatrixUshort, b *GslMatrixUshort) int32 {
	return int32(C.gsl_matrix_ushort_mul_elements((*C.gsl_matrix_ushort)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_ushort)(unsafe.Pointer(b.Ptr()))))
}

func UshortDivElements(a *GslMatrixUshort, b *GslMatrixUshort) int32 {
	return int32(C.gsl_matrix_ushort_div_elements((*C.gsl_matrix_ushort)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_ushort)(unsafe.Pointer(b.Ptr()))))
}

func UshortScale(a *GslMatrixUshort, x float64) int32 {
	return int32(C.gsl_matrix_ushort_scale((*C.gsl_matrix_ushort)(unsafe.Pointer(a.Ptr())), C.double(x)))
}

func UshortAddConstant(a *GslMatrixUshort, x float64) int32 {
	return int32(C.gsl_matrix_ushort_add_constant((*C.gsl_matrix_ushort)(unsafe.Pointer(a.Ptr())), C.double(x)))
}

func UshortMax(m *GslMatrixUshort) int16 {
	return int16(C.gsl_matrix_ushort_max((*C.gsl_matrix_ushort)(unsafe.Pointer(m.Ptr()))))
}

func UshortMin(m *GslMatrixUshort) int16 {
	return int16(C.gsl_matrix_ushort_min((*C.gsl_matrix_ushort)(unsafe.Pointer(m.Ptr()))))
}

func UshortMinmax(m *GslMatrixUshort) (int16, int16) {
	var _outptr_1 C.ushort
	var _outptr_2 C.ushort
	C.gsl_matrix_ushort_minmax((*C.gsl_matrix_ushort)(unsafe.Pointer(m.Ptr())), &_outptr_1, &_outptr_2)
	return *(*int16)(unsafe.Pointer(&_outptr_1)), *(*int16)(unsafe.Pointer(&_outptr_2))
}

func UshortMaxIndex(m *GslMatrixUshort) (int, int) {
	var _outptr_1 C.size_t
	var _outptr_2 C.size_t
	C.gsl_matrix_ushort_max_index((*C.gsl_matrix_ushort)(unsafe.Pointer(m.Ptr())), &_outptr_1, &_outptr_2)
	return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2))
}

func UshortMinIndex(m *GslMatrixUshort) (int, int) {
	var _outptr_1 C.size_t
	var _outptr_2 C.size_t
	C.gsl_matrix_ushort_min_index((*C.gsl_matrix_ushort)(unsafe.Pointer(m.Ptr())), &_outptr_1, &_outptr_2)
	return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2))
}

func UshortMinmaxIndex(m *GslMatrixUshort) (int, int, int, int) {
	var _outptr_1 C.size_t
	var _outptr_2 C.size_t
	var _outptr_3 C.size_t
	var _outptr_4 C.size_t
	C.gsl_matrix_ushort_minmax_index((*C.gsl_matrix_ushort)(unsafe.Pointer(m.Ptr())), &_outptr_1, &_outptr_2, &_outptr_3, &_outptr_4)
	return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2)), *(*int)(unsafe.Pointer(&_outptr_3)), *(*int)(unsafe.Pointer(&_outptr_4))
}

func UshortIsnull(m *GslMatrixUshort) int32 {
	return int32(C.gsl_matrix_ushort_isnull((*C.gsl_matrix_ushort)(unsafe.Pointer(m.Ptr()))))
}

func UshortIspos(m *GslMatrixUshort) int32 {
	return int32(C.gsl_matrix_ushort_ispos((*C.gsl_matrix_ushort)(unsafe.Pointer(m.Ptr()))))
}

func UshortIsneg(m *GslMatrixUshort) int32 {
	return int32(C.gsl_matrix_ushort_isneg((*C.gsl_matrix_ushort)(unsafe.Pointer(m.Ptr()))))
}

func UshortIsnonneg(m *GslMatrixUshort) int32 {
	return int32(C.gsl_matrix_ushort_isnonneg((*C.gsl_matrix_ushort)(unsafe.Pointer(m.Ptr()))))
}

func UshortEqual(a *GslMatrixUshort, b *GslMatrixUshort) int32 {
	return int32(C.gsl_matrix_ushort_equal((*C.gsl_matrix_ushort)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_ushort)(unsafe.Pointer(b.Ptr()))))
}

type GslMatrixChar struct {
	gogsl.GslReference
}

type GslMatrixCharView struct {
	gogsl.GslReference
	CData []byte
}

type GslMatrixCharConstView struct {
	gogsl.GslReference
	CData []byte
}

func MatrixCharAlloc(n1 int, n2 int) *GslMatrixChar {
	_ref := C.gsl_matrix_char_alloc(C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixChar{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func MatrixCharCalloc(n1 int, n2 int) *GslMatrixChar {
	_ref := C.gsl_matrix_char_calloc(C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixChar{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslMatrixChar) Dispose() {
	C.gsl_matrix_char_free((*C.gsl_matrix_char)(unsafe.Pointer(x.Ptr())))
}

func CharGet(m *GslMatrixChar, i int, j int) int8 {
	return int8(C.gsl_matrix_char_get((*C.gsl_matrix_char)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func CharSet(m *GslMatrixChar, i int, j int, x int8) {
	C.gsl_matrix_char_set((*C.gsl_matrix_char)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j), C.char(x))
}

func CharSetAll(m *GslMatrixChar, x int8) {
	C.gsl_matrix_char_set_all((*C.gsl_matrix_char)(unsafe.Pointer(m.Ptr())), C.char(x))
}

func CharSetZero(m *GslMatrixChar) {
	C.gsl_matrix_char_set_zero((*C.gsl_matrix_char)(unsafe.Pointer(m.Ptr())))
}

func CharSetIdentity(m *GslMatrixChar) {
	C.gsl_matrix_char_set_identity((*C.gsl_matrix_char)(unsafe.Pointer(m.Ptr())))
}

func CharFwrite(stream *os.File, m *GslMatrixChar) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_matrix_char_fwrite(_file_0, (*C.gsl_matrix_char)(unsafe.Pointer(m.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func CharFread(stream *os.File, m *GslMatrixChar) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_matrix_char_fread(_file_0, (*C.gsl_matrix_char)(unsafe.Pointer(m.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func CharFprintf(stream *os.File, m *GslMatrixChar, format string) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_string_2 := C.CString(format)
	_result := int32(C.gsl_matrix_char_fprintf(_file_0, (*C.gsl_matrix_char)(unsafe.Pointer(m.Ptr())), _string_2))
	C.fclose(_file_0)
	C.free(unsafe.Pointer(_string_2))
	return _result
}

func CharFscanf(stream *os.File, m *GslMatrixChar) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_matrix_char_fscanf(_file_0, (*C.gsl_matrix_char)(unsafe.Pointer(m.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func CharSubmatrix(m *GslMatrixChar, k1 int, k2 int, n1 int, n2 int) *GslMatrixCharView {
	_ref := C.gsl_matrix_char_submatrix((*C.gsl_matrix_char)(unsafe.Pointer(m.Ptr())), C.size_t(k1), C.size_t(k2), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixCharView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func CharConstSubmatrix(m *GslMatrixChar, k1 int, k2 int, n1 int, n2 int) *GslMatrixCharConstView {
	_ref := C.gsl_matrix_char_const_submatrix((*C.gsl_matrix_char)(unsafe.Pointer(m.Ptr())), C.size_t(k1), C.size_t(k2), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixCharConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func CharViewVector(v *vector.GslVectorChar, n1 int, n2 int) *GslMatrixCharView {
	_ref := C.gsl_matrix_char_view_vector((*C.gsl_vector_char)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixCharView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func CharConstViewVector(v *vector.GslVectorChar, n1 int, n2 int) *GslMatrixCharConstView {
	_ref := C.gsl_matrix_char_const_view_vector((*C.gsl_vector_char)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixCharConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func CharViewVectorWithTda(v *vector.GslVectorChar, n1 int, n2 int, tda int) *GslMatrixCharView {
	_ref := C.gsl_matrix_char_view_vector_with_tda((*C.gsl_vector_char)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2), C.size_t(tda))
	_result := &GslMatrixCharView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func CharConstViewVectorWithTda(v *vector.GslVectorChar, n1 int, n2 int, tda int) *GslMatrixCharConstView {
	_ref := C.gsl_matrix_char_const_view_vector_with_tda((*C.gsl_vector_char)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2), C.size_t(tda))
	_result := &GslMatrixCharConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func CharRow(m *GslMatrixChar, i int) *vector.GslVectorCharView {
	_ref := C.gsl_matrix_char_row((*C.gsl_matrix_char)(unsafe.Pointer(m.Ptr())), C.size_t(i))
	_result := &vector.GslVectorCharView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func CharConstRow(m *GslMatrixChar, i int) *vector.GslVectorCharConstView {
	_ref := C.gsl_matrix_char_const_row((*C.gsl_matrix_char)(unsafe.Pointer(m.Ptr())), C.size_t(i))
	_result := &vector.GslVectorCharConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func CharColumn(m *GslMatrixChar, j int) *vector.GslVectorCharView {
	_ref := C.gsl_matrix_char_column((*C.gsl_matrix_char)(unsafe.Pointer(m.Ptr())), C.size_t(j))
	_result := &vector.GslVectorCharView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func CharConstColumn(m *GslMatrixChar, j int) *vector.GslVectorCharConstView {
	_ref := C.gsl_matrix_char_const_column((*C.gsl_matrix_char)(unsafe.Pointer(m.Ptr())), C.size_t(j))
	_result := &vector.GslVectorCharConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func CharSubrow(m *GslMatrixChar, i int, offset int, n int) *vector.GslVectorCharView {
	_ref := C.gsl_matrix_char_subrow((*C.gsl_matrix_char)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorCharView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func CharConstSubrow(m *GslMatrixChar, i int, offset int, n int) *vector.GslVectorCharConstView {
	_ref := C.gsl_matrix_char_const_subrow((*C.gsl_matrix_char)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorCharConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func CharSubcolumn(m *GslMatrixChar, j int, offset int, n int) *vector.GslVectorCharView {
	_ref := C.gsl_matrix_char_subcolumn((*C.gsl_matrix_char)(unsafe.Pointer(m.Ptr())), C.size_t(j), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorCharView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func CharConstSubcolumn(m *GslMatrixChar, j int, offset int, n int) *vector.GslVectorCharConstView {
	_ref := C.gsl_matrix_char_const_subcolumn((*C.gsl_matrix_char)(unsafe.Pointer(m.Ptr())), C.size_t(j), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorCharConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func CharDiagonal(m *GslMatrixChar) *vector.GslVectorCharView {
	_ref := C.gsl_matrix_char_diagonal((*C.gsl_matrix_char)(unsafe.Pointer(m.Ptr())))
	_result := &vector.GslVectorCharView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func CharConstDiagonal(m *GslMatrixChar) *vector.GslVectorCharConstView {
	_ref := C.gsl_matrix_char_const_diagonal((*C.gsl_matrix_char)(unsafe.Pointer(m.Ptr())))
	_result := &vector.GslVectorCharConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func CharSubdiagonal(m *GslMatrixChar, k int) *vector.GslVectorCharView {
	_ref := C.gsl_matrix_char_subdiagonal((*C.gsl_matrix_char)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorCharView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func CharConstSubdiagonal(m *GslMatrixChar, k int) *vector.GslVectorCharConstView {
	_ref := C.gsl_matrix_char_const_subdiagonal((*C.gsl_matrix_char)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorCharConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func CharSuperdiagonal(m *GslMatrixChar, k int) *vector.GslVectorCharView {
	_ref := C.gsl_matrix_char_superdiagonal((*C.gsl_matrix_char)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorCharView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func CharConstSuperdiagonal(m *GslMatrixChar, k int) *vector.GslVectorCharConstView {
	_ref := C.gsl_matrix_char_const_superdiagonal((*C.gsl_matrix_char)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorCharConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func CharMemcpy(dest *GslMatrixChar, src *GslMatrixChar) int32 {
	return int32(C.gsl_matrix_char_memcpy((*C.gsl_matrix_char)(unsafe.Pointer(dest.Ptr())), (*C.gsl_matrix_char)(unsafe.Pointer(src.Ptr()))))
}

func CharSwap(m1 *GslMatrixChar, m2 *GslMatrixChar) int32 {
	return int32(C.gsl_matrix_char_swap((*C.gsl_matrix_char)(unsafe.Pointer(m1.Ptr())), (*C.gsl_matrix_char)(unsafe.Pointer(m2.Ptr()))))
}

func CharGetRow(v *vector.GslVectorChar, m *GslMatrixChar, i int) int32 {
	return int32(C.gsl_matrix_char_get_row((*C.gsl_vector_char)(unsafe.Pointer(v.Ptr())), (*C.gsl_matrix_char)(unsafe.Pointer(m.Ptr())), C.size_t(i)))
}

func CharGetCol(v *vector.GslVectorChar, m *GslMatrixChar, j int) int32 {
	return int32(C.gsl_matrix_char_get_col((*C.gsl_vector_char)(unsafe.Pointer(v.Ptr())), (*C.gsl_matrix_char)(unsafe.Pointer(m.Ptr())), C.size_t(j)))
}

func CharSetRow(m *GslMatrixChar, i int, v *vector.GslVectorChar) int32 {
	return int32(C.gsl_matrix_char_set_row((*C.gsl_matrix_char)(unsafe.Pointer(m.Ptr())), C.size_t(i), (*C.gsl_vector_char)(unsafe.Pointer(v.Ptr()))))
}

func CharSetCol(m *GslMatrixChar, j int, v *vector.GslVectorChar) int32 {
	return int32(C.gsl_matrix_char_set_col((*C.gsl_matrix_char)(unsafe.Pointer(m.Ptr())), C.size_t(j), (*C.gsl_vector_char)(unsafe.Pointer(v.Ptr()))))
}

func CharSwapRows(m *GslMatrixChar, i int, j int) int32 {
	return int32(C.gsl_matrix_char_swap_rows((*C.gsl_matrix_char)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func CharSwapColumns(m *GslMatrixChar, i int, j int) int32 {
	return int32(C.gsl_matrix_char_swap_columns((*C.gsl_matrix_char)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func CharSwapRowcol(m *GslMatrixChar, i int, j int) int32 {
	return int32(C.gsl_matrix_char_swap_rowcol((*C.gsl_matrix_char)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func CharTransposeMemcpy(dest *GslMatrixChar, src *GslMatrixChar) int32 {
	return int32(C.gsl_matrix_char_transpose_memcpy((*C.gsl_matrix_char)(unsafe.Pointer(dest.Ptr())), (*C.gsl_matrix_char)(unsafe.Pointer(src.Ptr()))))
}

func CharTranspose(m *GslMatrixChar) int32 {
	return int32(C.gsl_matrix_char_transpose((*C.gsl_matrix_char)(unsafe.Pointer(m.Ptr()))))
}

func CharAdd(a *GslMatrixChar, b *GslMatrixChar) int32 {
	return int32(C.gsl_matrix_char_add((*C.gsl_matrix_char)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_char)(unsafe.Pointer(b.Ptr()))))
}

func CharSub(a *GslMatrixChar, b *GslMatrixChar) int32 {
	return int32(C.gsl_matrix_char_sub((*C.gsl_matrix_char)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_char)(unsafe.Pointer(b.Ptr()))))
}

func CharMulElements(a *GslMatrixChar, b *GslMatrixChar) int32 {
	return int32(C.gsl_matrix_char_mul_elements((*C.gsl_matrix_char)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_char)(unsafe.Pointer(b.Ptr()))))
}

func CharDivElements(a *GslMatrixChar, b *GslMatrixChar) int32 {
	return int32(C.gsl_matrix_char_div_elements((*C.gsl_matrix_char)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_char)(unsafe.Pointer(b.Ptr()))))
}

func CharScale(a *GslMatrixChar, x float64) int32 {
	return int32(C.gsl_matrix_char_scale((*C.gsl_matrix_char)(unsafe.Pointer(a.Ptr())), C.double(x)))
}

func CharAddConstant(a *GslMatrixChar, x float64) int32 {
	return int32(C.gsl_matrix_char_add_constant((*C.gsl_matrix_char)(unsafe.Pointer(a.Ptr())), C.double(x)))
}

func CharMax(m *GslMatrixChar) int8 {
	return int8(C.gsl_matrix_char_max((*C.gsl_matrix_char)(unsafe.Pointer(m.Ptr()))))
}

func CharMin(m *GslMatrixChar) int8 {
	return int8(C.gsl_matrix_char_min((*C.gsl_matrix_char)(unsafe.Pointer(m.Ptr()))))
}

func CharMinmax(m *GslMatrixChar) (int8, int8) {
	var _outptr_1 C.char
	var _outptr_2 C.char
	C.gsl_matrix_char_minmax((*C.gsl_matrix_char)(unsafe.Pointer(m.Ptr())), &_outptr_1, &_outptr_2)
	return *(*int8)(unsafe.Pointer(&_outptr_1)), *(*int8)(unsafe.Pointer(&_outptr_2))
}

func CharMaxIndex(m *GslMatrixChar) (int, int) {
	var _outptr_1 C.size_t
	var _outptr_2 C.size_t
	C.gsl_matrix_char_max_index((*C.gsl_matrix_char)(unsafe.Pointer(m.Ptr())), &_outptr_1, &_outptr_2)
	return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2))
}

func CharMinIndex(m *GslMatrixChar) (int, int) {
	var _outptr_1 C.size_t
	var _outptr_2 C.size_t
	C.gsl_matrix_char_min_index((*C.gsl_matrix_char)(unsafe.Pointer(m.Ptr())), &_outptr_1, &_outptr_2)
	return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2))
}

func CharMinmaxIndex(m *GslMatrixChar) (int, int, int, int) {
	var _outptr_1 C.size_t
	var _outptr_2 C.size_t
	var _outptr_3 C.size_t
	var _outptr_4 C.size_t
	C.gsl_matrix_char_minmax_index((*C.gsl_matrix_char)(unsafe.Pointer(m.Ptr())), &_outptr_1, &_outptr_2, &_outptr_3, &_outptr_4)
	return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2)), *(*int)(unsafe.Pointer(&_outptr_3)), *(*int)(unsafe.Pointer(&_outptr_4))
}

func CharIsnull(m *GslMatrixChar) int32 {
	return int32(C.gsl_matrix_char_isnull((*C.gsl_matrix_char)(unsafe.Pointer(m.Ptr()))))
}

func CharIspos(m *GslMatrixChar) int32 {
	return int32(C.gsl_matrix_char_ispos((*C.gsl_matrix_char)(unsafe.Pointer(m.Ptr()))))
}

func CharIsneg(m *GslMatrixChar) int32 {
	return int32(C.gsl_matrix_char_isneg((*C.gsl_matrix_char)(unsafe.Pointer(m.Ptr()))))
}

func CharIsnonneg(m *GslMatrixChar) int32 {
	return int32(C.gsl_matrix_char_isnonneg((*C.gsl_matrix_char)(unsafe.Pointer(m.Ptr()))))
}

func CharEqual(a *GslMatrixChar, b *GslMatrixChar) int32 {
	return int32(C.gsl_matrix_char_equal((*C.gsl_matrix_char)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_char)(unsafe.Pointer(b.Ptr()))))
}

type GslMatrixUchar struct {
	gogsl.GslReference
}

type GslMatrixUcharView struct {
	gogsl.GslReference
	CData []byte
}

type GslMatrixUcharConstView struct {
	gogsl.GslReference
	CData []byte
}

func MatrixUcharAlloc(n1 int, n2 int) *GslMatrixUchar {
	_ref := C.gsl_matrix_uchar_alloc(C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixUchar{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func MatrixUcharCalloc(n1 int, n2 int) *GslMatrixUchar {
	_ref := C.gsl_matrix_uchar_calloc(C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixUchar{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslMatrixUchar) Dispose() {
	C.gsl_matrix_uchar_free((*C.gsl_matrix_uchar)(unsafe.Pointer(x.Ptr())))
}

func UcharGet(m *GslMatrixUchar, i int, j int) uint8 {
	return uint8(C.gsl_matrix_uchar_get((*C.gsl_matrix_uchar)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func UcharSet(m *GslMatrixUchar, i int, j int, x uint8) {
	C.gsl_matrix_uchar_set((*C.gsl_matrix_uchar)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j), C.uchar(x))
}

func UcharSetAll(m *GslMatrixUchar, x uint8) {
	C.gsl_matrix_uchar_set_all((*C.gsl_matrix_uchar)(unsafe.Pointer(m.Ptr())), C.uchar(x))
}

func UcharSetZero(m *GslMatrixUchar) {
	C.gsl_matrix_uchar_set_zero((*C.gsl_matrix_uchar)(unsafe.Pointer(m.Ptr())))
}

func UcharSetIdentity(m *GslMatrixUchar) {
	C.gsl_matrix_uchar_set_identity((*C.gsl_matrix_uchar)(unsafe.Pointer(m.Ptr())))
}

func UcharFwrite(stream *os.File, m *GslMatrixUchar) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_matrix_uchar_fwrite(_file_0, (*C.gsl_matrix_uchar)(unsafe.Pointer(m.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func UcharFread(stream *os.File, m *GslMatrixUchar) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_matrix_uchar_fread(_file_0, (*C.gsl_matrix_uchar)(unsafe.Pointer(m.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func UcharFprintf(stream *os.File, m *GslMatrixUchar, format string) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_string_2 := C.CString(format)
	_result := int32(C.gsl_matrix_uchar_fprintf(_file_0, (*C.gsl_matrix_uchar)(unsafe.Pointer(m.Ptr())), _string_2))
	C.fclose(_file_0)
	C.free(unsafe.Pointer(_string_2))
	return _result
}

func UcharFscanf(stream *os.File, m *GslMatrixUchar) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_matrix_uchar_fscanf(_file_0, (*C.gsl_matrix_uchar)(unsafe.Pointer(m.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func UcharSubmatrix(m *GslMatrixUchar, k1 int, k2 int, n1 int, n2 int) *GslMatrixUcharView {
	_ref := C.gsl_matrix_uchar_submatrix((*C.gsl_matrix_uchar)(unsafe.Pointer(m.Ptr())), C.size_t(k1), C.size_t(k2), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixUcharView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UcharConstSubmatrix(m *GslMatrixUchar, k1 int, k2 int, n1 int, n2 int) *GslMatrixUcharConstView {
	_ref := C.gsl_matrix_uchar_const_submatrix((*C.gsl_matrix_uchar)(unsafe.Pointer(m.Ptr())), C.size_t(k1), C.size_t(k2), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixUcharConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UcharViewVector(v *vector.GslVectorUchar, n1 int, n2 int) *GslMatrixUcharView {
	_ref := C.gsl_matrix_uchar_view_vector((*C.gsl_vector_uchar)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixUcharView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UcharConstViewVector(v *vector.GslVectorUchar, n1 int, n2 int) *GslMatrixUcharConstView {
	_ref := C.gsl_matrix_uchar_const_view_vector((*C.gsl_vector_uchar)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixUcharConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UcharViewVectorWithTda(v *vector.GslVectorUchar, n1 int, n2 int, tda int) *GslMatrixUcharView {
	_ref := C.gsl_matrix_uchar_view_vector_with_tda((*C.gsl_vector_uchar)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2), C.size_t(tda))
	_result := &GslMatrixUcharView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UcharConstViewVectorWithTda(v *vector.GslVectorUchar, n1 int, n2 int, tda int) *GslMatrixUcharConstView {
	_ref := C.gsl_matrix_uchar_const_view_vector_with_tda((*C.gsl_vector_uchar)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2), C.size_t(tda))
	_result := &GslMatrixUcharConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UcharRow(m *GslMatrixUchar, i int) *vector.GslVectorUcharView {
	_ref := C.gsl_matrix_uchar_row((*C.gsl_matrix_uchar)(unsafe.Pointer(m.Ptr())), C.size_t(i))
	_result := &vector.GslVectorUcharView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UcharConstRow(m *GslMatrixUchar, i int) *vector.GslVectorUcharConstView {
	_ref := C.gsl_matrix_uchar_const_row((*C.gsl_matrix_uchar)(unsafe.Pointer(m.Ptr())), C.size_t(i))
	_result := &vector.GslVectorUcharConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UcharColumn(m *GslMatrixUchar, j int) *vector.GslVectorUcharView {
	_ref := C.gsl_matrix_uchar_column((*C.gsl_matrix_uchar)(unsafe.Pointer(m.Ptr())), C.size_t(j))
	_result := &vector.GslVectorUcharView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UcharConstColumn(m *GslMatrixUchar, j int) *vector.GslVectorUcharConstView {
	_ref := C.gsl_matrix_uchar_const_column((*C.gsl_matrix_uchar)(unsafe.Pointer(m.Ptr())), C.size_t(j))
	_result := &vector.GslVectorUcharConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UcharSubrow(m *GslMatrixUchar, i int, offset int, n int) *vector.GslVectorUcharView {
	_ref := C.gsl_matrix_uchar_subrow((*C.gsl_matrix_uchar)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorUcharView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UcharConstSubrow(m *GslMatrixUchar, i int, offset int, n int) *vector.GslVectorUcharConstView {
	_ref := C.gsl_matrix_uchar_const_subrow((*C.gsl_matrix_uchar)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorUcharConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UcharSubcolumn(m *GslMatrixUchar, j int, offset int, n int) *vector.GslVectorUcharView {
	_ref := C.gsl_matrix_uchar_subcolumn((*C.gsl_matrix_uchar)(unsafe.Pointer(m.Ptr())), C.size_t(j), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorUcharView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UcharConstSubcolumn(m *GslMatrixUchar, j int, offset int, n int) *vector.GslVectorUcharConstView {
	_ref := C.gsl_matrix_uchar_const_subcolumn((*C.gsl_matrix_uchar)(unsafe.Pointer(m.Ptr())), C.size_t(j), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorUcharConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UcharDiagonal(m *GslMatrixUchar) *vector.GslVectorUcharView {
	_ref := C.gsl_matrix_uchar_diagonal((*C.gsl_matrix_uchar)(unsafe.Pointer(m.Ptr())))
	_result := &vector.GslVectorUcharView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UcharConstDiagonal(m *GslMatrixUchar) *vector.GslVectorUcharConstView {
	_ref := C.gsl_matrix_uchar_const_diagonal((*C.gsl_matrix_uchar)(unsafe.Pointer(m.Ptr())))
	_result := &vector.GslVectorUcharConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UcharSubdiagonal(m *GslMatrixUchar, k int) *vector.GslVectorUcharView {
	_ref := C.gsl_matrix_uchar_subdiagonal((*C.gsl_matrix_uchar)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorUcharView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UcharConstSubdiagonal(m *GslMatrixUchar, k int) *vector.GslVectorUcharConstView {
	_ref := C.gsl_matrix_uchar_const_subdiagonal((*C.gsl_matrix_uchar)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorUcharConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UcharSuperdiagonal(m *GslMatrixUchar, k int) *vector.GslVectorUcharView {
	_ref := C.gsl_matrix_uchar_superdiagonal((*C.gsl_matrix_uchar)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorUcharView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UcharConstSuperdiagonal(m *GslMatrixUchar, k int) *vector.GslVectorUcharConstView {
	_ref := C.gsl_matrix_uchar_const_superdiagonal((*C.gsl_matrix_uchar)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorUcharConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func UcharMemcpy(dest *GslMatrixUchar, src *GslMatrixUchar) int32 {
	return int32(C.gsl_matrix_uchar_memcpy((*C.gsl_matrix_uchar)(unsafe.Pointer(dest.Ptr())), (*C.gsl_matrix_uchar)(unsafe.Pointer(src.Ptr()))))
}

func UcharSwap(m1 *GslMatrixUchar, m2 *GslMatrixUchar) int32 {
	return int32(C.gsl_matrix_uchar_swap((*C.gsl_matrix_uchar)(unsafe.Pointer(m1.Ptr())), (*C.gsl_matrix_uchar)(unsafe.Pointer(m2.Ptr()))))
}

func UcharGetRow(v *vector.GslVectorUchar, m *GslMatrixUchar, i int) int32 {
	return int32(C.gsl_matrix_uchar_get_row((*C.gsl_vector_uchar)(unsafe.Pointer(v.Ptr())), (*C.gsl_matrix_uchar)(unsafe.Pointer(m.Ptr())), C.size_t(i)))
}

func UcharGetCol(v *vector.GslVectorUchar, m *GslMatrixUchar, j int) int32 {
	return int32(C.gsl_matrix_uchar_get_col((*C.gsl_vector_uchar)(unsafe.Pointer(v.Ptr())), (*C.gsl_matrix_uchar)(unsafe.Pointer(m.Ptr())), C.size_t(j)))
}

func UcharSetRow(m *GslMatrixUchar, i int, v *vector.GslVectorUchar) int32 {
	return int32(C.gsl_matrix_uchar_set_row((*C.gsl_matrix_uchar)(unsafe.Pointer(m.Ptr())), C.size_t(i), (*C.gsl_vector_uchar)(unsafe.Pointer(v.Ptr()))))
}

func UcharSetCol(m *GslMatrixUchar, j int, v *vector.GslVectorUchar) int32 {
	return int32(C.gsl_matrix_uchar_set_col((*C.gsl_matrix_uchar)(unsafe.Pointer(m.Ptr())), C.size_t(j), (*C.gsl_vector_uchar)(unsafe.Pointer(v.Ptr()))))
}

func UcharSwapRows(m *GslMatrixUchar, i int, j int) int32 {
	return int32(C.gsl_matrix_uchar_swap_rows((*C.gsl_matrix_uchar)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func UcharSwapColumns(m *GslMatrixUchar, i int, j int) int32 {
	return int32(C.gsl_matrix_uchar_swap_columns((*C.gsl_matrix_uchar)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func UcharSwapRowcol(m *GslMatrixUchar, i int, j int) int32 {
	return int32(C.gsl_matrix_uchar_swap_rowcol((*C.gsl_matrix_uchar)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func UcharTransposeMemcpy(dest *GslMatrixUchar, src *GslMatrixUchar) int32 {
	return int32(C.gsl_matrix_uchar_transpose_memcpy((*C.gsl_matrix_uchar)(unsafe.Pointer(dest.Ptr())), (*C.gsl_matrix_uchar)(unsafe.Pointer(src.Ptr()))))
}

func UcharTranspose(m *GslMatrixUchar) int32 {
	return int32(C.gsl_matrix_uchar_transpose((*C.gsl_matrix_uchar)(unsafe.Pointer(m.Ptr()))))
}

func UcharAdd(a *GslMatrixUchar, b *GslMatrixUchar) int32 {
	return int32(C.gsl_matrix_uchar_add((*C.gsl_matrix_uchar)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_uchar)(unsafe.Pointer(b.Ptr()))))
}

func UcharSub(a *GslMatrixUchar, b *GslMatrixUchar) int32 {
	return int32(C.gsl_matrix_uchar_sub((*C.gsl_matrix_uchar)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_uchar)(unsafe.Pointer(b.Ptr()))))
}

func UcharMulElements(a *GslMatrixUchar, b *GslMatrixUchar) int32 {
	return int32(C.gsl_matrix_uchar_mul_elements((*C.gsl_matrix_uchar)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_uchar)(unsafe.Pointer(b.Ptr()))))
}

func UcharDivElements(a *GslMatrixUchar, b *GslMatrixUchar) int32 {
	return int32(C.gsl_matrix_uchar_div_elements((*C.gsl_matrix_uchar)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_uchar)(unsafe.Pointer(b.Ptr()))))
}

func UcharScale(a *GslMatrixUchar, x float64) int32 {
	return int32(C.gsl_matrix_uchar_scale((*C.gsl_matrix_uchar)(unsafe.Pointer(a.Ptr())), C.double(x)))
}

func UcharAddConstant(a *GslMatrixUchar, x float64) int32 {
	return int32(C.gsl_matrix_uchar_add_constant((*C.gsl_matrix_uchar)(unsafe.Pointer(a.Ptr())), C.double(x)))
}

func UcharMax(m *GslMatrixUchar) uint8 {
	return uint8(C.gsl_matrix_uchar_max((*C.gsl_matrix_uchar)(unsafe.Pointer(m.Ptr()))))
}

func UcharMin(m *GslMatrixUchar) uint8 {
	return uint8(C.gsl_matrix_uchar_min((*C.gsl_matrix_uchar)(unsafe.Pointer(m.Ptr()))))
}

func UcharMinmax(m *GslMatrixUchar) (uint8, uint8) {
	var _outptr_1 C.uchar
	var _outptr_2 C.uchar
	C.gsl_matrix_uchar_minmax((*C.gsl_matrix_uchar)(unsafe.Pointer(m.Ptr())), &_outptr_1, &_outptr_2)
	return *(*uint8)(unsafe.Pointer(&_outptr_1)), *(*uint8)(unsafe.Pointer(&_outptr_2))
}

func UcharMaxIndex(m *GslMatrixUchar) (int, int) {
	var _outptr_1 C.size_t
	var _outptr_2 C.size_t
	C.gsl_matrix_uchar_max_index((*C.gsl_matrix_uchar)(unsafe.Pointer(m.Ptr())), &_outptr_1, &_outptr_2)
	return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2))
}

func UcharMinIndex(m *GslMatrixUchar) (int, int) {
	var _outptr_1 C.size_t
	var _outptr_2 C.size_t
	C.gsl_matrix_uchar_min_index((*C.gsl_matrix_uchar)(unsafe.Pointer(m.Ptr())), &_outptr_1, &_outptr_2)
	return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2))
}

func UcharMinmaxIndex(m *GslMatrixUchar) (int, int, int, int) {
	var _outptr_1 C.size_t
	var _outptr_2 C.size_t
	var _outptr_3 C.size_t
	var _outptr_4 C.size_t
	C.gsl_matrix_uchar_minmax_index((*C.gsl_matrix_uchar)(unsafe.Pointer(m.Ptr())), &_outptr_1, &_outptr_2, &_outptr_3, &_outptr_4)
	return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2)), *(*int)(unsafe.Pointer(&_outptr_3)), *(*int)(unsafe.Pointer(&_outptr_4))
}

func UcharIsnull(m *GslMatrixUchar) int32 {
	return int32(C.gsl_matrix_uchar_isnull((*C.gsl_matrix_uchar)(unsafe.Pointer(m.Ptr()))))
}

func UcharIspos(m *GslMatrixUchar) int32 {
	return int32(C.gsl_matrix_uchar_ispos((*C.gsl_matrix_uchar)(unsafe.Pointer(m.Ptr()))))
}

func UcharIsneg(m *GslMatrixUchar) int32 {
	return int32(C.gsl_matrix_uchar_isneg((*C.gsl_matrix_uchar)(unsafe.Pointer(m.Ptr()))))
}

func UcharIsnonneg(m *GslMatrixUchar) int32 {
	return int32(C.gsl_matrix_uchar_isnonneg((*C.gsl_matrix_uchar)(unsafe.Pointer(m.Ptr()))))
}

func UcharEqual(a *GslMatrixUchar, b *GslMatrixUchar) int32 {
	return int32(C.gsl_matrix_uchar_equal((*C.gsl_matrix_uchar)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_uchar)(unsafe.Pointer(b.Ptr()))))
}

type GslMatrixComplex struct {
	gogsl.GslReference
}

type GslMatrixComplexView struct {
	gogsl.GslReference
	CData []byte
}

type GslMatrixComplexConstView struct {
	gogsl.GslReference
	CData []byte
}

func MatrixComplexAlloc(n1 int, n2 int) *GslMatrixComplex {
	_ref := C.gsl_matrix_complex_alloc(C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixComplex{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func MatrixComplexCalloc(n1 int, n2 int) *GslMatrixComplex {
	_ref := C.gsl_matrix_complex_calloc(C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixComplex{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslMatrixComplex) Dispose() {
	C.gsl_matrix_complex_free((*C.gsl_matrix_complex)(unsafe.Pointer(x.Ptr())))
}

func ComplexGet(m *GslMatrixComplex, i int, j int) complex128 {
	_result := C.gsl_matrix_complex_get((*C.gsl_matrix_complex)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j))
	return complex_.GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func ComplexSet(m *GslMatrixComplex, i int, j int, x complex128) {
	_arg_3 := complex_.GoComplexToGsl(x)
	C.gsl_matrix_complex_set((*C.gsl_matrix_complex)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j), *(*C.gsl_complex)(unsafe.Pointer(_arg_3)))
}

func ComplexSetAll(m *GslMatrixComplex, x complex128) {
	_arg_1 := complex_.GoComplexToGsl(x)
	C.gsl_matrix_complex_set_all((*C.gsl_matrix_complex)(unsafe.Pointer(m.Ptr())), *(*C.gsl_complex)(unsafe.Pointer(_arg_1)))
}

func ComplexSetZero(m *GslMatrixComplex) {
	C.gsl_matrix_complex_set_zero((*C.gsl_matrix_complex)(unsafe.Pointer(m.Ptr())))
}

func ComplexSetIdentity(m *GslMatrixComplex) {
	C.gsl_matrix_complex_set_identity((*C.gsl_matrix_complex)(unsafe.Pointer(m.Ptr())))
}

func ComplexFwrite(stream *os.File, m *GslMatrixComplex) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_matrix_complex_fwrite(_file_0, (*C.gsl_matrix_complex)(unsafe.Pointer(m.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func ComplexFread(stream *os.File, m *GslMatrixComplex) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_matrix_complex_fread(_file_0, (*C.gsl_matrix_complex)(unsafe.Pointer(m.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func ComplexFprintf(stream *os.File, m *GslMatrixComplex, format string) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_string_2 := C.CString(format)
	_result := int32(C.gsl_matrix_complex_fprintf(_file_0, (*C.gsl_matrix_complex)(unsafe.Pointer(m.Ptr())), _string_2))
	C.fclose(_file_0)
	C.free(unsafe.Pointer(_string_2))
	return _result
}

func ComplexFscanf(stream *os.File, m *GslMatrixComplex) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_matrix_complex_fscanf(_file_0, (*C.gsl_matrix_complex)(unsafe.Pointer(m.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func ComplexSubmatrix(m *GslMatrixComplex, k1 int, k2 int, n1 int, n2 int) *GslMatrixComplexView {
	_ref := C.gsl_matrix_complex_submatrix((*C.gsl_matrix_complex)(unsafe.Pointer(m.Ptr())), C.size_t(k1), C.size_t(k2), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixComplexView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ComplexConstSubmatrix(m *GslMatrixComplex, k1 int, k2 int, n1 int, n2 int) *GslMatrixComplexConstView {
	_ref := C.gsl_matrix_complex_const_submatrix((*C.gsl_matrix_complex)(unsafe.Pointer(m.Ptr())), C.size_t(k1), C.size_t(k2), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixComplexConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ComplexViewVector(v *vector.GslVectorComplex, n1 int, n2 int) *GslMatrixComplexView {
	_ref := C.gsl_matrix_complex_view_vector((*C.gsl_vector_complex)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixComplexView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ComplexConstViewVector(v *vector.GslVectorComplex, n1 int, n2 int) *GslMatrixComplexConstView {
	_ref := C.gsl_matrix_complex_const_view_vector((*C.gsl_vector_complex)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixComplexConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ComplexViewVectorWithTda(v *vector.GslVectorComplex, n1 int, n2 int, tda int) *GslMatrixComplexView {
	_ref := C.gsl_matrix_complex_view_vector_with_tda((*C.gsl_vector_complex)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2), C.size_t(tda))
	_result := &GslMatrixComplexView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ComplexConstViewVectorWithTda(v *vector.GslVectorComplex, n1 int, n2 int, tda int) *GslMatrixComplexConstView {
	_ref := C.gsl_matrix_complex_const_view_vector_with_tda((*C.gsl_vector_complex)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2), C.size_t(tda))
	_result := &GslMatrixComplexConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ComplexRow(m *GslMatrixComplex, i int) *vector.GslVectorComplexView {
	_ref := C.gsl_matrix_complex_row((*C.gsl_matrix_complex)(unsafe.Pointer(m.Ptr())), C.size_t(i))
	_result := &vector.GslVectorComplexView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ComplexConstRow(m *GslMatrixComplex, i int) *vector.GslVectorComplexConstView {
	_ref := C.gsl_matrix_complex_const_row((*C.gsl_matrix_complex)(unsafe.Pointer(m.Ptr())), C.size_t(i))
	_result := &vector.GslVectorComplexConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ComplexColumn(m *GslMatrixComplex, j int) *vector.GslVectorComplexView {
	_ref := C.gsl_matrix_complex_column((*C.gsl_matrix_complex)(unsafe.Pointer(m.Ptr())), C.size_t(j))
	_result := &vector.GslVectorComplexView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ComplexConstColumn(m *GslMatrixComplex, j int) *vector.GslVectorComplexConstView {
	_ref := C.gsl_matrix_complex_const_column((*C.gsl_matrix_complex)(unsafe.Pointer(m.Ptr())), C.size_t(j))
	_result := &vector.GslVectorComplexConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ComplexSubrow(m *GslMatrixComplex, i int, offset int, n int) *vector.GslVectorComplexView {
	_ref := C.gsl_matrix_complex_subrow((*C.gsl_matrix_complex)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorComplexView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ComplexConstSubrow(m *GslMatrixComplex, i int, offset int, n int) *vector.GslVectorComplexConstView {
	_ref := C.gsl_matrix_complex_const_subrow((*C.gsl_matrix_complex)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorComplexConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ComplexSubcolumn(m *GslMatrixComplex, j int, offset int, n int) *vector.GslVectorComplexView {
	_ref := C.gsl_matrix_complex_subcolumn((*C.gsl_matrix_complex)(unsafe.Pointer(m.Ptr())), C.size_t(j), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorComplexView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ComplexConstSubcolumn(m *GslMatrixComplex, j int, offset int, n int) *vector.GslVectorComplexConstView {
	_ref := C.gsl_matrix_complex_const_subcolumn((*C.gsl_matrix_complex)(unsafe.Pointer(m.Ptr())), C.size_t(j), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorComplexConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ComplexDiagonal(m *GslMatrixComplex) *vector.GslVectorComplexView {
	_ref := C.gsl_matrix_complex_diagonal((*C.gsl_matrix_complex)(unsafe.Pointer(m.Ptr())))
	_result := &vector.GslVectorComplexView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ComplexConstDiagonal(m *GslMatrixComplex) *vector.GslVectorComplexConstView {
	_ref := C.gsl_matrix_complex_const_diagonal((*C.gsl_matrix_complex)(unsafe.Pointer(m.Ptr())))
	_result := &vector.GslVectorComplexConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ComplexSubdiagonal(m *GslMatrixComplex, k int) *vector.GslVectorComplexView {
	_ref := C.gsl_matrix_complex_subdiagonal((*C.gsl_matrix_complex)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorComplexView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ComplexConstSubdiagonal(m *GslMatrixComplex, k int) *vector.GslVectorComplexConstView {
	_ref := C.gsl_matrix_complex_const_subdiagonal((*C.gsl_matrix_complex)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorComplexConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ComplexSuperdiagonal(m *GslMatrixComplex, k int) *vector.GslVectorComplexView {
	_ref := C.gsl_matrix_complex_superdiagonal((*C.gsl_matrix_complex)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorComplexView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ComplexConstSuperdiagonal(m *GslMatrixComplex, k int) *vector.GslVectorComplexConstView {
	_ref := C.gsl_matrix_complex_const_superdiagonal((*C.gsl_matrix_complex)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorComplexConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ComplexMemcpy(dest *GslMatrixComplex, src *GslMatrixComplex) int32 {
	return int32(C.gsl_matrix_complex_memcpy((*C.gsl_matrix_complex)(unsafe.Pointer(dest.Ptr())), (*C.gsl_matrix_complex)(unsafe.Pointer(src.Ptr()))))
}

func ComplexSwap(m1 *GslMatrixComplex, m2 *GslMatrixComplex) int32 {
	return int32(C.gsl_matrix_complex_swap((*C.gsl_matrix_complex)(unsafe.Pointer(m1.Ptr())), (*C.gsl_matrix_complex)(unsafe.Pointer(m2.Ptr()))))
}

func ComplexGetRow(v *vector.GslVectorComplex, m *GslMatrixComplex, i int) int32 {
	return int32(C.gsl_matrix_complex_get_row((*C.gsl_vector_complex)(unsafe.Pointer(v.Ptr())), (*C.gsl_matrix_complex)(unsafe.Pointer(m.Ptr())), C.size_t(i)))
}

func ComplexGetCol(v *vector.GslVectorComplex, m *GslMatrixComplex, j int) int32 {
	return int32(C.gsl_matrix_complex_get_col((*C.gsl_vector_complex)(unsafe.Pointer(v.Ptr())), (*C.gsl_matrix_complex)(unsafe.Pointer(m.Ptr())), C.size_t(j)))
}

func ComplexSetRow(m *GslMatrixComplex, i int, v *vector.GslVectorComplex) int32 {
	return int32(C.gsl_matrix_complex_set_row((*C.gsl_matrix_complex)(unsafe.Pointer(m.Ptr())), C.size_t(i), (*C.gsl_vector_complex)(unsafe.Pointer(v.Ptr()))))
}

func ComplexSetCol(m *GslMatrixComplex, j int, v *vector.GslVectorComplex) int32 {
	return int32(C.gsl_matrix_complex_set_col((*C.gsl_matrix_complex)(unsafe.Pointer(m.Ptr())), C.size_t(j), (*C.gsl_vector_complex)(unsafe.Pointer(v.Ptr()))))
}

func ComplexSwapRows(m *GslMatrixComplex, i int, j int) int32 {
	return int32(C.gsl_matrix_complex_swap_rows((*C.gsl_matrix_complex)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func ComplexSwapColumns(m *GslMatrixComplex, i int, j int) int32 {
	return int32(C.gsl_matrix_complex_swap_columns((*C.gsl_matrix_complex)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func ComplexSwapRowcol(m *GslMatrixComplex, i int, j int) int32 {
	return int32(C.gsl_matrix_complex_swap_rowcol((*C.gsl_matrix_complex)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func ComplexTransposeMemcpy(dest *GslMatrixComplex, src *GslMatrixComplex) int32 {
	return int32(C.gsl_matrix_complex_transpose_memcpy((*C.gsl_matrix_complex)(unsafe.Pointer(dest.Ptr())), (*C.gsl_matrix_complex)(unsafe.Pointer(src.Ptr()))))
}

func ComplexTranspose(m *GslMatrixComplex) int32 {
	return int32(C.gsl_matrix_complex_transpose((*C.gsl_matrix_complex)(unsafe.Pointer(m.Ptr()))))
}

func ComplexAdd(a *GslMatrixComplex, b *GslMatrixComplex) int32 {
	return int32(C.gsl_matrix_complex_add((*C.gsl_matrix_complex)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_complex)(unsafe.Pointer(b.Ptr()))))
}

func ComplexSub(a *GslMatrixComplex, b *GslMatrixComplex) int32 {
	return int32(C.gsl_matrix_complex_sub((*C.gsl_matrix_complex)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_complex)(unsafe.Pointer(b.Ptr()))))
}

func ComplexMulElements(a *GslMatrixComplex, b *GslMatrixComplex) int32 {
	return int32(C.gsl_matrix_complex_mul_elements((*C.gsl_matrix_complex)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_complex)(unsafe.Pointer(b.Ptr()))))
}

func ComplexDivElements(a *GslMatrixComplex, b *GslMatrixComplex) int32 {
	return int32(C.gsl_matrix_complex_div_elements((*C.gsl_matrix_complex)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_complex)(unsafe.Pointer(b.Ptr()))))
}

func ComplexIsnull(m *GslMatrixComplex) int32 {
	return int32(C.gsl_matrix_complex_isnull((*C.gsl_matrix_complex)(unsafe.Pointer(m.Ptr()))))
}

func ComplexIspos(m *GslMatrixComplex) int32 {
	return int32(C.gsl_matrix_complex_ispos((*C.gsl_matrix_complex)(unsafe.Pointer(m.Ptr()))))
}

func ComplexIsneg(m *GslMatrixComplex) int32 {
	return int32(C.gsl_matrix_complex_isneg((*C.gsl_matrix_complex)(unsafe.Pointer(m.Ptr()))))
}

func ComplexIsnonneg(m *GslMatrixComplex) int32 {
	return int32(C.gsl_matrix_complex_isnonneg((*C.gsl_matrix_complex)(unsafe.Pointer(m.Ptr()))))
}

func ComplexEqual(a *GslMatrixComplex, b *GslMatrixComplex) int32 {
	return int32(C.gsl_matrix_complex_equal((*C.gsl_matrix_complex)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_complex)(unsafe.Pointer(b.Ptr()))))
}

type GslMatrixComplexFloat struct {
	gogsl.GslReference
}

type GslMatrixComplexFloatView struct {
	gogsl.GslReference
	CData []byte
}

type GslMatrixComplexFloatConstView struct {
	gogsl.GslReference
	CData []byte
}

func MatrixComplexFloatAlloc(n1 int, n2 int) *GslMatrixComplexFloat {
	_ref := C.gsl_matrix_complex_float_alloc(C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixComplexFloat{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func MatrixComplexFloatCalloc(n1 int, n2 int) *GslMatrixComplexFloat {
	_ref := C.gsl_matrix_complex_float_calloc(C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixComplexFloat{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslMatrixComplexFloat) Dispose() {
	C.gsl_matrix_complex_float_free((*C.gsl_matrix_complex_float)(unsafe.Pointer(x.Ptr())))
}

func ComplexFloatGet(m *GslMatrixComplexFloat, i int, j int) complex64 {
	_result := C.gsl_matrix_complex_float_get((*C.gsl_matrix_complex_float)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j))
	return complex_.GslComplexFloatToGo(uintptr(unsafe.Pointer(&_result)))
}

func ComplexFloatSet(m *GslMatrixComplexFloat, i int, j int, x complex64) {
	_arg_3 := complex_.GoComplexFloatToGsl(x)
	C.gsl_matrix_complex_float_set((*C.gsl_matrix_complex_float)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j), *(*C.gsl_complex_float)(unsafe.Pointer(_arg_3)))
}

func ComplexFloatSetAll(m *GslMatrixComplexFloat, x complex64) {
	_arg_1 := complex_.GoComplexFloatToGsl(x)
	C.gsl_matrix_complex_float_set_all((*C.gsl_matrix_complex_float)(unsafe.Pointer(m.Ptr())), *(*C.gsl_complex_float)(unsafe.Pointer(_arg_1)))
}

func ComplexFloatSetZero(m *GslMatrixComplexFloat) {
	C.gsl_matrix_complex_float_set_zero((*C.gsl_matrix_complex_float)(unsafe.Pointer(m.Ptr())))
}

func ComplexFloatSetIdentity(m *GslMatrixComplexFloat) {
	C.gsl_matrix_complex_float_set_identity((*C.gsl_matrix_complex_float)(unsafe.Pointer(m.Ptr())))
}

func ComplexFloatFwrite(stream *os.File, m *GslMatrixComplexFloat) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_matrix_complex_float_fwrite(_file_0, (*C.gsl_matrix_complex_float)(unsafe.Pointer(m.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func ComplexFloatFread(stream *os.File, m *GslMatrixComplexFloat) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_matrix_complex_float_fread(_file_0, (*C.gsl_matrix_complex_float)(unsafe.Pointer(m.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func ComplexFloatFprintf(stream *os.File, m *GslMatrixComplexFloat, format string) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_string_2 := C.CString(format)
	_result := int32(C.gsl_matrix_complex_float_fprintf(_file_0, (*C.gsl_matrix_complex_float)(unsafe.Pointer(m.Ptr())), _string_2))
	C.fclose(_file_0)
	C.free(unsafe.Pointer(_string_2))
	return _result
}

func ComplexFloatFscanf(stream *os.File, m *GslMatrixComplexFloat) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_matrix_complex_float_fscanf(_file_0, (*C.gsl_matrix_complex_float)(unsafe.Pointer(m.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func ComplexFloatSubmatrix(m *GslMatrixComplexFloat, k1 int, k2 int, n1 int, n2 int) *GslMatrixComplexFloatView {
	_ref := C.gsl_matrix_complex_float_submatrix((*C.gsl_matrix_complex_float)(unsafe.Pointer(m.Ptr())), C.size_t(k1), C.size_t(k2), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixComplexFloatView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ComplexFloatConstSubmatrix(m *GslMatrixComplexFloat, k1 int, k2 int, n1 int, n2 int) *GslMatrixComplexFloatConstView {
	_ref := C.gsl_matrix_complex_float_const_submatrix((*C.gsl_matrix_complex_float)(unsafe.Pointer(m.Ptr())), C.size_t(k1), C.size_t(k2), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixComplexFloatConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ComplexFloatViewVector(v *vector.GslVectorComplexFloat, n1 int, n2 int) *GslMatrixComplexFloatView {
	_ref := C.gsl_matrix_complex_float_view_vector((*C.gsl_vector_complex_float)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixComplexFloatView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ComplexFloatConstViewVector(v *vector.GslVectorComplexFloat, n1 int, n2 int) *GslMatrixComplexFloatConstView {
	_ref := C.gsl_matrix_complex_float_const_view_vector((*C.gsl_vector_complex_float)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2))
	_result := &GslMatrixComplexFloatConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ComplexFloatViewVectorWithTda(v *vector.GslVectorComplexFloat, n1 int, n2 int, tda int) *GslMatrixComplexFloatView {
	_ref := C.gsl_matrix_complex_float_view_vector_with_tda((*C.gsl_vector_complex_float)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2), C.size_t(tda))
	_result := &GslMatrixComplexFloatView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ComplexFloatConstViewVectorWithTda(v *vector.GslVectorComplexFloat, n1 int, n2 int, tda int) *GslMatrixComplexFloatConstView {
	_ref := C.gsl_matrix_complex_float_const_view_vector_with_tda((*C.gsl_vector_complex_float)(unsafe.Pointer(v.Ptr())), C.size_t(n1), C.size_t(n2), C.size_t(tda))
	_result := &GslMatrixComplexFloatConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ComplexFloatRow(m *GslMatrixComplexFloat, i int) *vector.GslVectorComplexFloatView {
	_ref := C.gsl_matrix_complex_float_row((*C.gsl_matrix_complex_float)(unsafe.Pointer(m.Ptr())), C.size_t(i))
	_result := &vector.GslVectorComplexFloatView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ComplexFloatConstRow(m *GslMatrixComplexFloat, i int) *vector.GslVectorComplexFloatConstView {
	_ref := C.gsl_matrix_complex_float_const_row((*C.gsl_matrix_complex_float)(unsafe.Pointer(m.Ptr())), C.size_t(i))
	_result := &vector.GslVectorComplexFloatConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ComplexFloatColumn(m *GslMatrixComplexFloat, j int) *vector.GslVectorComplexFloatView {
	_ref := C.gsl_matrix_complex_float_column((*C.gsl_matrix_complex_float)(unsafe.Pointer(m.Ptr())), C.size_t(j))
	_result := &vector.GslVectorComplexFloatView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ComplexFloatConstColumn(m *GslMatrixComplexFloat, j int) *vector.GslVectorComplexFloatConstView {
	_ref := C.gsl_matrix_complex_float_const_column((*C.gsl_matrix_complex_float)(unsafe.Pointer(m.Ptr())), C.size_t(j))
	_result := &vector.GslVectorComplexFloatConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ComplexFloatSubrow(m *GslMatrixComplexFloat, i int, offset int, n int) *vector.GslVectorComplexFloatView {
	_ref := C.gsl_matrix_complex_float_subrow((*C.gsl_matrix_complex_float)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorComplexFloatView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ComplexFloatConstSubrow(m *GslMatrixComplexFloat, i int, offset int, n int) *vector.GslVectorComplexFloatConstView {
	_ref := C.gsl_matrix_complex_float_const_subrow((*C.gsl_matrix_complex_float)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorComplexFloatConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ComplexFloatSubcolumn(m *GslMatrixComplexFloat, j int, offset int, n int) *vector.GslVectorComplexFloatView {
	_ref := C.gsl_matrix_complex_float_subcolumn((*C.gsl_matrix_complex_float)(unsafe.Pointer(m.Ptr())), C.size_t(j), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorComplexFloatView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ComplexFloatConstSubcolumn(m *GslMatrixComplexFloat, j int, offset int, n int) *vector.GslVectorComplexFloatConstView {
	_ref := C.gsl_matrix_complex_float_const_subcolumn((*C.gsl_matrix_complex_float)(unsafe.Pointer(m.Ptr())), C.size_t(j), C.size_t(offset), C.size_t(n))
	_result := &vector.GslVectorComplexFloatConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ComplexFloatDiagonal(m *GslMatrixComplexFloat) *vector.GslVectorComplexFloatView {
	_ref := C.gsl_matrix_complex_float_diagonal((*C.gsl_matrix_complex_float)(unsafe.Pointer(m.Ptr())))
	_result := &vector.GslVectorComplexFloatView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ComplexFloatConstDiagonal(m *GslMatrixComplexFloat) *vector.GslVectorComplexFloatConstView {
	_ref := C.gsl_matrix_complex_float_const_diagonal((*C.gsl_matrix_complex_float)(unsafe.Pointer(m.Ptr())))
	_result := &vector.GslVectorComplexFloatConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ComplexFloatSubdiagonal(m *GslMatrixComplexFloat, k int) *vector.GslVectorComplexFloatView {
	_ref := C.gsl_matrix_complex_float_subdiagonal((*C.gsl_matrix_complex_float)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorComplexFloatView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ComplexFloatConstSubdiagonal(m *GslMatrixComplexFloat, k int) *vector.GslVectorComplexFloatConstView {
	_ref := C.gsl_matrix_complex_float_const_subdiagonal((*C.gsl_matrix_complex_float)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorComplexFloatConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ComplexFloatSuperdiagonal(m *GslMatrixComplexFloat, k int) *vector.GslVectorComplexFloatView {
	_ref := C.gsl_matrix_complex_float_superdiagonal((*C.gsl_matrix_complex_float)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorComplexFloatView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ComplexFloatConstSuperdiagonal(m *GslMatrixComplexFloat, k int) *vector.GslVectorComplexFloatConstView {
	_ref := C.gsl_matrix_complex_float_const_superdiagonal((*C.gsl_matrix_complex_float)(unsafe.Pointer(m.Ptr())), C.size_t(k))
	_result := &vector.GslVectorComplexFloatConstView{}
	_result.CData = make([]byte, unsafe.Sizeof(_ref))
	gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))
	return _result
}

func ComplexFloatMemcpy(dest *GslMatrixComplexFloat, src *GslMatrixComplexFloat) int32 {
	return int32(C.gsl_matrix_complex_float_memcpy((*C.gsl_matrix_complex_float)(unsafe.Pointer(dest.Ptr())), (*C.gsl_matrix_complex_float)(unsafe.Pointer(src.Ptr()))))
}

func ComplexFloatSwap(m1 *GslMatrixComplexFloat, m2 *GslMatrixComplexFloat) int32 {
	return int32(C.gsl_matrix_complex_float_swap((*C.gsl_matrix_complex_float)(unsafe.Pointer(m1.Ptr())), (*C.gsl_matrix_complex_float)(unsafe.Pointer(m2.Ptr()))))
}

func ComplexFloatGetRow(v *vector.GslVectorComplexFloat, m *GslMatrixComplexFloat, i int) int32 {
	return int32(C.gsl_matrix_complex_float_get_row((*C.gsl_vector_complex_float)(unsafe.Pointer(v.Ptr())), (*C.gsl_matrix_complex_float)(unsafe.Pointer(m.Ptr())), C.size_t(i)))
}

func ComplexFloatGetCol(v *vector.GslVectorComplexFloat, m *GslMatrixComplexFloat, j int) int32 {
	return int32(C.gsl_matrix_complex_float_get_col((*C.gsl_vector_complex_float)(unsafe.Pointer(v.Ptr())), (*C.gsl_matrix_complex_float)(unsafe.Pointer(m.Ptr())), C.size_t(j)))
}

func ComplexFloatSetRow(m *GslMatrixComplexFloat, i int, v *vector.GslVectorComplexFloat) int32 {
	return int32(C.gsl_matrix_complex_float_set_row((*C.gsl_matrix_complex_float)(unsafe.Pointer(m.Ptr())), C.size_t(i), (*C.gsl_vector_complex_float)(unsafe.Pointer(v.Ptr()))))
}

func ComplexFloatSetCol(m *GslMatrixComplexFloat, j int, v *vector.GslVectorComplexFloat) int32 {
	return int32(C.gsl_matrix_complex_float_set_col((*C.gsl_matrix_complex_float)(unsafe.Pointer(m.Ptr())), C.size_t(j), (*C.gsl_vector_complex_float)(unsafe.Pointer(v.Ptr()))))
}

func ComplexFloatSwapRows(m *GslMatrixComplexFloat, i int, j int) int32 {
	return int32(C.gsl_matrix_complex_float_swap_rows((*C.gsl_matrix_complex_float)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func ComplexFloatSwapColumns(m *GslMatrixComplexFloat, i int, j int) int32 {
	return int32(C.gsl_matrix_complex_float_swap_columns((*C.gsl_matrix_complex_float)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func ComplexFloatSwapRowcol(m *GslMatrixComplexFloat, i int, j int) int32 {
	return int32(C.gsl_matrix_complex_float_swap_rowcol((*C.gsl_matrix_complex_float)(unsafe.Pointer(m.Ptr())), C.size_t(i), C.size_t(j)))
}

func ComplexFloatTransposeMemcpy(dest *GslMatrixComplexFloat, src *GslMatrixComplexFloat) int32 {
	return int32(C.gsl_matrix_complex_float_transpose_memcpy((*C.gsl_matrix_complex_float)(unsafe.Pointer(dest.Ptr())), (*C.gsl_matrix_complex_float)(unsafe.Pointer(src.Ptr()))))
}

func ComplexFloatTranspose(m *GslMatrixComplexFloat) int32 {
	return int32(C.gsl_matrix_complex_float_transpose((*C.gsl_matrix_complex_float)(unsafe.Pointer(m.Ptr()))))
}

func ComplexFloatAdd(a *GslMatrixComplexFloat, b *GslMatrixComplexFloat) int32 {
	return int32(C.gsl_matrix_complex_float_add((*C.gsl_matrix_complex_float)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_complex_float)(unsafe.Pointer(b.Ptr()))))
}

func ComplexFloatSub(a *GslMatrixComplexFloat, b *GslMatrixComplexFloat) int32 {
	return int32(C.gsl_matrix_complex_float_sub((*C.gsl_matrix_complex_float)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_complex_float)(unsafe.Pointer(b.Ptr()))))
}

func ComplexFloatMulElements(a *GslMatrixComplexFloat, b *GslMatrixComplexFloat) int32 {
	return int32(C.gsl_matrix_complex_float_mul_elements((*C.gsl_matrix_complex_float)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_complex_float)(unsafe.Pointer(b.Ptr()))))
}

func ComplexFloatDivElements(a *GslMatrixComplexFloat, b *GslMatrixComplexFloat) int32 {
	return int32(C.gsl_matrix_complex_float_div_elements((*C.gsl_matrix_complex_float)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_complex_float)(unsafe.Pointer(b.Ptr()))))
}

func ComplexFloatIsnull(m *GslMatrixComplexFloat) int32 {
	return int32(C.gsl_matrix_complex_float_isnull((*C.gsl_matrix_complex_float)(unsafe.Pointer(m.Ptr()))))
}

func ComplexFloatIspos(m *GslMatrixComplexFloat) int32 {
	return int32(C.gsl_matrix_complex_float_ispos((*C.gsl_matrix_complex_float)(unsafe.Pointer(m.Ptr()))))
}

func ComplexFloatIsneg(m *GslMatrixComplexFloat) int32 {
	return int32(C.gsl_matrix_complex_float_isneg((*C.gsl_matrix_complex_float)(unsafe.Pointer(m.Ptr()))))
}

func ComplexFloatIsnonneg(m *GslMatrixComplexFloat) int32 {
	return int32(C.gsl_matrix_complex_float_isnonneg((*C.gsl_matrix_complex_float)(unsafe.Pointer(m.Ptr()))))
}

func ComplexFloatEqual(a *GslMatrixComplexFloat, b *GslMatrixComplexFloat) int32 {
	return int32(C.gsl_matrix_complex_float_equal((*C.gsl_matrix_complex_float)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_complex_float)(unsafe.Pointer(b.Ptr()))))
}

func ComplexScale(a *GslMatrixComplex, x complex128) int32 {
	_arg_1 := complex_.GoComplexToGsl(x)
	return int32(C.gsl_matrix_complex_scale((*C.gsl_matrix_complex)(unsafe.Pointer(a.Ptr())), *(*C.gsl_complex)(unsafe.Pointer(_arg_1))))
}

func ComplexFloatScale(a *GslMatrixComplexFloat, x complex64) int32 {
	_arg_1 := complex_.GoComplexFloatToGsl(x)
	return int32(C.gsl_matrix_complex_float_scale((*C.gsl_matrix_complex_float)(unsafe.Pointer(a.Ptr())), *(*C.gsl_complex_float)(unsafe.Pointer(_arg_1))))
}

func ComplexAddConstant(a *GslMatrixComplex, x complex128) int32 {
	_arg_1 := complex_.GoComplexToGsl(x)
	return int32(C.gsl_matrix_complex_add_constant((*C.gsl_matrix_complex)(unsafe.Pointer(a.Ptr())), *(*C.gsl_complex)(unsafe.Pointer(_arg_1))))
}

func ComplexFloatAddConstant(a *GslMatrixComplexFloat, x complex64) int32 {
	_arg_1 := complex_.GoComplexFloatToGsl(x)
	return int32(C.gsl_matrix_complex_float_add_constant((*C.gsl_matrix_complex_float)(unsafe.Pointer(a.Ptr())), *(*C.gsl_complex_float)(unsafe.Pointer(_arg_1))))
}
