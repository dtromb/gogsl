//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package block

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_block.h>
*/
import "C"

import "unsafe"
import "github.com/dtromb/gogsl"
import "os"

type GslBlock struct {
	gogsl.GslReference
}

func BlockAlloc(n int) *GslBlock {
	_ref := C.gsl_block_alloc(C.size_t(n))
	_result := &GslBlock{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func BlockCalloc(n int) *GslBlock {
	_ref := C.gsl_block_calloc(C.size_t(n))
	_result := &GslBlock{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslBlock) Dispose() {
	C.gsl_block_free((*C.gsl_block)(unsafe.Pointer(x.Ptr())))
}

func BlockFwrite(stream *os.File, b *GslBlock) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_block_fwrite(_file_0, (*C.gsl_block)(unsafe.Pointer(b.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func BlockFread(stream *os.File, b *GslBlock) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_block_fread(_file_0, (*C.gsl_block)(unsafe.Pointer(b.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func BlockFprintf(stream *os.File, b *GslBlock, format string) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_string_2 := C.CString(format)
	_result := int32(C.gsl_block_fprintf(_file_0, (*C.gsl_block)(unsafe.Pointer(b.Ptr())), _string_2))
	C.fclose(_file_0)
	C.free(unsafe.Pointer(_string_2))
	return _result
}

func BlockFscanf(stream *os.File, b *GslBlock) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_block_fscanf(_file_0, (*C.gsl_block)(unsafe.Pointer(b.Ptr()))))
	C.fclose(_file_0)
	return _result
}

type GslBlockFloat struct {
	gogsl.GslReference
}

func BlockFloatAlloc(n int) *GslBlockFloat {
	_ref := C.gsl_block_float_alloc(C.size_t(n))
	_result := &GslBlockFloat{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func BlockFloatCalloc(n int) *GslBlockFloat {
	_ref := C.gsl_block_float_calloc(C.size_t(n))
	_result := &GslBlockFloat{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslBlockFloat) Dispose() {
	C.gsl_block_float_free((*C.gsl_block_float)(unsafe.Pointer(x.Ptr())))
}

func BlockFloatFwrite(stream *os.File, b *GslBlockFloat) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_block_float_fwrite(_file_0, (*C.gsl_block_float)(unsafe.Pointer(b.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func BlockFloatFread(stream *os.File, b *GslBlockFloat) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_block_float_fread(_file_0, (*C.gsl_block_float)(unsafe.Pointer(b.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func BlockFloatFprintf(stream *os.File, b *GslBlockFloat, format string) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_string_2 := C.CString(format)
	_result := int32(C.gsl_block_float_fprintf(_file_0, (*C.gsl_block_float)(unsafe.Pointer(b.Ptr())), _string_2))
	C.fclose(_file_0)
	C.free(unsafe.Pointer(_string_2))
	return _result
}

func BlockFloatFscanf(stream *os.File, b *GslBlockFloat) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_block_float_fscanf(_file_0, (*C.gsl_block_float)(unsafe.Pointer(b.Ptr()))))
	C.fclose(_file_0)
	return _result
}

type GslBlockInt struct {
	gogsl.GslReference
}

func BlockIntAlloc(n int) *GslBlockInt {
	_ref := C.gsl_block_int_alloc(C.size_t(n))
	_result := &GslBlockInt{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func BlockIntCalloc(n int) *GslBlockInt {
	_ref := C.gsl_block_int_calloc(C.size_t(n))
	_result := &GslBlockInt{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslBlockInt) Dispose() {
	C.gsl_block_int_free((*C.gsl_block_int)(unsafe.Pointer(x.Ptr())))
}

func BlockIntFwrite(stream *os.File, b *GslBlockInt) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_block_int_fwrite(_file_0, (*C.gsl_block_int)(unsafe.Pointer(b.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func BlockIntFread(stream *os.File, b *GslBlockInt) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_block_int_fread(_file_0, (*C.gsl_block_int)(unsafe.Pointer(b.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func BlockIntFprintf(stream *os.File, b *GslBlockInt, format string) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_string_2 := C.CString(format)
	_result := int32(C.gsl_block_int_fprintf(_file_0, (*C.gsl_block_int)(unsafe.Pointer(b.Ptr())), _string_2))
	C.fclose(_file_0)
	C.free(unsafe.Pointer(_string_2))
	return _result
}

func BlockIntFscanf(stream *os.File, b *GslBlockInt) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_block_int_fscanf(_file_0, (*C.gsl_block_int)(unsafe.Pointer(b.Ptr()))))
	C.fclose(_file_0)
	return _result
}

type GslBlockUint struct {
	gogsl.GslReference
}

func BlockUintAlloc(n int) *GslBlockUint {
	_ref := C.gsl_block_uint_alloc(C.size_t(n))
	_result := &GslBlockUint{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func BlockUintCalloc(n int) *GslBlockUint {
	_ref := C.gsl_block_uint_calloc(C.size_t(n))
	_result := &GslBlockUint{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslBlockUint) Dispose() {
	C.gsl_block_uint_free((*C.gsl_block_uint)(unsafe.Pointer(x.Ptr())))
}

func BlockUintFwrite(stream *os.File, b *GslBlockUint) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_block_uint_fwrite(_file_0, (*C.gsl_block_uint)(unsafe.Pointer(b.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func BlockUintFread(stream *os.File, b *GslBlockUint) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_block_uint_fread(_file_0, (*C.gsl_block_uint)(unsafe.Pointer(b.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func BlockUintFprintf(stream *os.File, b *GslBlockUint, format string) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_string_2 := C.CString(format)
	_result := int32(C.gsl_block_uint_fprintf(_file_0, (*C.gsl_block_uint)(unsafe.Pointer(b.Ptr())), _string_2))
	C.fclose(_file_0)
	C.free(unsafe.Pointer(_string_2))
	return _result
}

func BlockUintFscanf(stream *os.File, b *GslBlockUint) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_block_uint_fscanf(_file_0, (*C.gsl_block_uint)(unsafe.Pointer(b.Ptr()))))
	C.fclose(_file_0)
	return _result
}

type GslBlockLong struct {
	gogsl.GslReference
}

func BlockLongAlloc(n int) *GslBlockLong {
	_ref := C.gsl_block_long_alloc(C.size_t(n))
	_result := &GslBlockLong{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func BlockLongCalloc(n int) *GslBlockLong {
	_ref := C.gsl_block_long_calloc(C.size_t(n))
	_result := &GslBlockLong{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslBlockLong) Dispose() {
	C.gsl_block_long_free((*C.gsl_block_long)(unsafe.Pointer(x.Ptr())))
}

func BlockLongFwrite(stream *os.File, b *GslBlockLong) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_block_long_fwrite(_file_0, (*C.gsl_block_long)(unsafe.Pointer(b.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func BlockLongFread(stream *os.File, b *GslBlockLong) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_block_long_fread(_file_0, (*C.gsl_block_long)(unsafe.Pointer(b.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func BlockLongFprintf(stream *os.File, b *GslBlockLong, format string) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_string_2 := C.CString(format)
	_result := int32(C.gsl_block_long_fprintf(_file_0, (*C.gsl_block_long)(unsafe.Pointer(b.Ptr())), _string_2))
	C.fclose(_file_0)
	C.free(unsafe.Pointer(_string_2))
	return _result
}

func BlockLongFscanf(stream *os.File, b *GslBlockLong) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_block_long_fscanf(_file_0, (*C.gsl_block_long)(unsafe.Pointer(b.Ptr()))))
	C.fclose(_file_0)
	return _result
}

type GslBlockUlong struct {
	gogsl.GslReference
}

func BlockUlongAlloc(n int) *GslBlockUlong {
	_ref := C.gsl_block_ulong_alloc(C.size_t(n))
	_result := &GslBlockUlong{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func BlockUlongCalloc(n int) *GslBlockUlong {
	_ref := C.gsl_block_ulong_calloc(C.size_t(n))
	_result := &GslBlockUlong{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslBlockUlong) Dispose() {
	C.gsl_block_ulong_free((*C.gsl_block_ulong)(unsafe.Pointer(x.Ptr())))
}

func BlockUlongFwrite(stream *os.File, b *GslBlockUlong) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_block_ulong_fwrite(_file_0, (*C.gsl_block_ulong)(unsafe.Pointer(b.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func BlockUlongFread(stream *os.File, b *GslBlockUlong) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_block_ulong_fread(_file_0, (*C.gsl_block_ulong)(unsafe.Pointer(b.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func BlockUlongFprintf(stream *os.File, b *GslBlockUlong, format string) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_string_2 := C.CString(format)
	_result := int32(C.gsl_block_ulong_fprintf(_file_0, (*C.gsl_block_ulong)(unsafe.Pointer(b.Ptr())), _string_2))
	C.fclose(_file_0)
	C.free(unsafe.Pointer(_string_2))
	return _result
}

func BlockUlongFscanf(stream *os.File, b *GslBlockUlong) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_block_ulong_fscanf(_file_0, (*C.gsl_block_ulong)(unsafe.Pointer(b.Ptr()))))
	C.fclose(_file_0)
	return _result
}

type GslBlockShort struct {
	gogsl.GslReference
}

func BlockShortAlloc(n int) *GslBlockShort {
	_ref := C.gsl_block_short_alloc(C.size_t(n))
	_result := &GslBlockShort{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func BlockShortCalloc(n int) *GslBlockShort {
	_ref := C.gsl_block_short_calloc(C.size_t(n))
	_result := &GslBlockShort{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslBlockShort) Dispose() {
	C.gsl_block_short_free((*C.gsl_block_short)(unsafe.Pointer(x.Ptr())))
}

func BlockShortFwrite(stream *os.File, b *GslBlockShort) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_block_short_fwrite(_file_0, (*C.gsl_block_short)(unsafe.Pointer(b.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func BlockShortFread(stream *os.File, b *GslBlockShort) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_block_short_fread(_file_0, (*C.gsl_block_short)(unsafe.Pointer(b.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func BlockShortFprintf(stream *os.File, b *GslBlockShort, format string) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_string_2 := C.CString(format)
	_result := int32(C.gsl_block_short_fprintf(_file_0, (*C.gsl_block_short)(unsafe.Pointer(b.Ptr())), _string_2))
	C.fclose(_file_0)
	C.free(unsafe.Pointer(_string_2))
	return _result
}

func BlockShortFscanf(stream *os.File, b *GslBlockShort) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_block_short_fscanf(_file_0, (*C.gsl_block_short)(unsafe.Pointer(b.Ptr()))))
	C.fclose(_file_0)
	return _result
}

type GslBlockUshort struct {
	gogsl.GslReference
}

func BlockUshortAlloc(n int) *GslBlockUshort {
	_ref := C.gsl_block_ushort_alloc(C.size_t(n))
	_result := &GslBlockUshort{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func BlockUshortCalloc(n int) *GslBlockUshort {
	_ref := C.gsl_block_ushort_calloc(C.size_t(n))
	_result := &GslBlockUshort{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslBlockUshort) Dispose() {
	C.gsl_block_ushort_free((*C.gsl_block_ushort)(unsafe.Pointer(x.Ptr())))
}

func BlockUshortFwrite(stream *os.File, b *GslBlockUshort) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_block_ushort_fwrite(_file_0, (*C.gsl_block_ushort)(unsafe.Pointer(b.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func BlockUshortFread(stream *os.File, b *GslBlockUshort) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_block_ushort_fread(_file_0, (*C.gsl_block_ushort)(unsafe.Pointer(b.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func BlockUshortFprintf(stream *os.File, b *GslBlockUshort, format string) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_string_2 := C.CString(format)
	_result := int32(C.gsl_block_ushort_fprintf(_file_0, (*C.gsl_block_ushort)(unsafe.Pointer(b.Ptr())), _string_2))
	C.fclose(_file_0)
	C.free(unsafe.Pointer(_string_2))
	return _result
}

func BlockUshortFscanf(stream *os.File, b *GslBlockUshort) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_block_ushort_fscanf(_file_0, (*C.gsl_block_ushort)(unsafe.Pointer(b.Ptr()))))
	C.fclose(_file_0)
	return _result
}

type GslBlockChar struct {
	gogsl.GslReference
}

func BlockCharAlloc(n int) *GslBlockChar {
	_ref := C.gsl_block_char_alloc(C.size_t(n))
	_result := &GslBlockChar{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func BlockCharCalloc(n int) *GslBlockChar {
	_ref := C.gsl_block_char_calloc(C.size_t(n))
	_result := &GslBlockChar{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslBlockChar) Dispose() {
	C.gsl_block_char_free((*C.gsl_block_char)(unsafe.Pointer(x.Ptr())))
}

func BlockCharFwrite(stream *os.File, b *GslBlockChar) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_block_char_fwrite(_file_0, (*C.gsl_block_char)(unsafe.Pointer(b.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func BlockCharFread(stream *os.File, b *GslBlockChar) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_block_char_fread(_file_0, (*C.gsl_block_char)(unsafe.Pointer(b.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func BlockCharFprintf(stream *os.File, b *GslBlockChar, format string) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_string_2 := C.CString(format)
	_result := int32(C.gsl_block_char_fprintf(_file_0, (*C.gsl_block_char)(unsafe.Pointer(b.Ptr())), _string_2))
	C.fclose(_file_0)
	C.free(unsafe.Pointer(_string_2))
	return _result
}

func BlockCharFscanf(stream *os.File, b *GslBlockChar) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_block_char_fscanf(_file_0, (*C.gsl_block_char)(unsafe.Pointer(b.Ptr()))))
	C.fclose(_file_0)
	return _result
}

type GslBlockUchar struct {
	gogsl.GslReference
}

func BlockUcharAlloc(n int) *GslBlockUchar {
	_ref := C.gsl_block_uchar_alloc(C.size_t(n))
	_result := &GslBlockUchar{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func BlockUcharCalloc(n int) *GslBlockUchar {
	_ref := C.gsl_block_uchar_calloc(C.size_t(n))
	_result := &GslBlockUchar{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslBlockUchar) Dispose() {
	C.gsl_block_uchar_free((*C.gsl_block_uchar)(unsafe.Pointer(x.Ptr())))
}

func BlockUcharFwrite(stream *os.File, b *GslBlockUchar) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_block_uchar_fwrite(_file_0, (*C.gsl_block_uchar)(unsafe.Pointer(b.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func BlockUcharFread(stream *os.File, b *GslBlockUchar) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_block_uchar_fread(_file_0, (*C.gsl_block_uchar)(unsafe.Pointer(b.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func BlockUcharFprintf(stream *os.File, b *GslBlockUchar, format string) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_string_2 := C.CString(format)
	_result := int32(C.gsl_block_uchar_fprintf(_file_0, (*C.gsl_block_uchar)(unsafe.Pointer(b.Ptr())), _string_2))
	C.fclose(_file_0)
	C.free(unsafe.Pointer(_string_2))
	return _result
}

func BlockUcharFscanf(stream *os.File, b *GslBlockUchar) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_block_uchar_fscanf(_file_0, (*C.gsl_block_uchar)(unsafe.Pointer(b.Ptr()))))
	C.fclose(_file_0)
	return _result
}

type GslBlockComplex struct {
	gogsl.GslReference
}

func BlockComplexAlloc(n int) *GslBlockComplex {
	_ref := C.gsl_block_complex_alloc(C.size_t(n))
	_result := &GslBlockComplex{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func BlockComplexCalloc(n int) *GslBlockComplex {
	_ref := C.gsl_block_complex_calloc(C.size_t(n))
	_result := &GslBlockComplex{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslBlockComplex) Dispose() {
	C.gsl_block_complex_free((*C.gsl_block_complex)(unsafe.Pointer(x.Ptr())))
}

func BlockComplexFwrite(stream *os.File, b *GslBlockComplex) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_block_complex_fwrite(_file_0, (*C.gsl_block_complex)(unsafe.Pointer(b.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func BlockComplexFread(stream *os.File, b *GslBlockComplex) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_block_complex_fread(_file_0, (*C.gsl_block_complex)(unsafe.Pointer(b.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func BlockComplexFprintf(stream *os.File, b *GslBlockComplex, format string) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_string_2 := C.CString(format)
	_result := int32(C.gsl_block_complex_fprintf(_file_0, (*C.gsl_block_complex)(unsafe.Pointer(b.Ptr())), _string_2))
	C.fclose(_file_0)
	C.free(unsafe.Pointer(_string_2))
	return _result
}

func BlockComplexFscanf(stream *os.File, b *GslBlockComplex) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_block_complex_fscanf(_file_0, (*C.gsl_block_complex)(unsafe.Pointer(b.Ptr()))))
	C.fclose(_file_0)
	return _result
}

type GslBlockComplexFloat struct {
	gogsl.GslReference
}

func BlockComplexFloatAlloc(n int) *GslBlockComplexFloat {
	_ref := C.gsl_block_complex_float_alloc(C.size_t(n))
	_result := &GslBlockComplexFloat{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func BlockComplexFloatCalloc(n int) *GslBlockComplexFloat {
	_ref := C.gsl_block_complex_float_calloc(C.size_t(n))
	_result := &GslBlockComplexFloat{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslBlockComplexFloat) Dispose() {
	C.gsl_block_complex_float_free((*C.gsl_block_complex_float)(unsafe.Pointer(x.Ptr())))
}

func BlockComplexFloatFwrite(stream *os.File, b *GslBlockComplexFloat) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_block_complex_float_fwrite(_file_0, (*C.gsl_block_complex_float)(unsafe.Pointer(b.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func BlockComplexFloatFread(stream *os.File, b *GslBlockComplexFloat) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_block_complex_float_fread(_file_0, (*C.gsl_block_complex_float)(unsafe.Pointer(b.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func BlockComplexFloatFprintf(stream *os.File, b *GslBlockComplexFloat, format string) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_string_2 := C.CString(format)
	_result := int32(C.gsl_block_complex_float_fprintf(_file_0, (*C.gsl_block_complex_float)(unsafe.Pointer(b.Ptr())), _string_2))
	C.fclose(_file_0)
	C.free(unsafe.Pointer(_string_2))
	return _result
}

func BlockComplexFloatFscanf(stream *os.File, b *GslBlockComplexFloat) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_block_complex_float_fscanf(_file_0, (*C.gsl_block_complex_float)(unsafe.Pointer(b.Ptr()))))
	C.fclose(_file_0)
	return _result
}
