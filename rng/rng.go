//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package rng

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_rng.h>
   #include <unistd.h>
*/
import "C"

import "unsafe"
import "github.com/dtromb/gogsl"
import "os"

type GslRng struct {
	gogsl.GslReference
}

type GslRngType struct {
	gogsl.GslReference
	cPtr uintptr
}

func RngAlloc(t *GslRngType) *GslRng {
	_ref := C.gsl_rng_alloc((*C.gsl_rng_type)(unsafe.Pointer(t.Ptr())))
	_result := &GslRng{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func Set(r *GslRng, s int) {
	C.gsl_rng_set((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), C.ulong(s))
}

func (x *GslRng) Dispose() {
	C.gsl_rng_free((*C.gsl_rng)(unsafe.Pointer(x.Ptr())))
}

func Get(r *GslRng) int {
	return int(C.gsl_rng_get((*C.gsl_rng)(unsafe.Pointer(r.Ptr()))))
}

func Uniform(r *GslRng) float64 {
	return float64(C.gsl_rng_uniform((*C.gsl_rng)(unsafe.Pointer(r.Ptr()))))
}

func UniformPos(r *GslRng) float64 {
	return float64(C.gsl_rng_uniform_pos((*C.gsl_rng)(unsafe.Pointer(r.Ptr()))))
}

func UniformInt(r *GslRng, n int) int {
	return int(C.gsl_rng_uniform_int((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), C.ulong(n)))
}

func Name(r *GslRng) string {
	return C.GoString(C.gsl_rng_name((*C.gsl_rng)(unsafe.Pointer(r.Ptr()))))
}

func Max(r *GslRng) int {
	return int(C.gsl_rng_max((*C.gsl_rng)(unsafe.Pointer(r.Ptr()))))
}

func Min(r *GslRng) int {
	return int(C.gsl_rng_min((*C.gsl_rng)(unsafe.Pointer(r.Ptr()))))
}

func State(r *GslRng) *RngState {
	return (*RngState)(C.gsl_rng_state((*C.gsl_rng)(unsafe.Pointer(r.Ptr()))))
}

func Size(r *GslRng) int {
	return int(C.gsl_rng_size((*C.gsl_rng)(unsafe.Pointer(r.Ptr()))))
}

func EnvSetup() *GslRngType {
	_ref := C.gsl_rng_env_setup()
	_result := &GslRngType{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func Memcpy(dest *GslRng, src *GslRng) int32 {
	return int32(C.gsl_rng_memcpy((*C.gsl_rng)(unsafe.Pointer(dest.Ptr())), (*C.gsl_rng)(unsafe.Pointer(src.Ptr()))))
}

func Clone(r *GslRng) *GslRng {
	_ref := C.gsl_rng_clone((*C.gsl_rng)(unsafe.Pointer(r.Ptr())))
	_result := &GslRng{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func Fwrite(stream *os.File, r *GslRng) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_rng_fwrite(_file_0, (*C.gsl_rng)(unsafe.Pointer(r.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func Fread(stream *os.File, r *GslRng) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_rng_fread(_file_0, (*C.gsl_rng)(unsafe.Pointer(r.Ptr()))))
	C.fclose(_file_0)
	return _result
}
