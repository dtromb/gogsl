//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package permutation

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_permutation.h>
   #include <gsl/gsl_permute_vector.h>
   #include <unistd.h>
*/
import "C"

import "unsafe"
import "github.com/dtromb/gogsl"
import "github.com/dtromb/gogsl/vector"
import "os"

type GslPermutation struct {
	gogsl.GslReference
}

func PermutationAlloc(n int) *GslPermutation {
	_ref := C.gsl_permutation_alloc(C.size_t(n))
	_result := &GslPermutation{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func PermutationCalloc(n int) *GslPermutation {
	_ref := C.gsl_permutation_calloc(C.size_t(n))
	_result := &GslPermutation{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func PermutationInit(p *GslPermutation) {
	C.gsl_permutation_init((*C.gsl_permutation)(unsafe.Pointer(p.Ptr())))
}

func (x *GslPermutation) Dispose() {
	C.gsl_permutation_free((*C.gsl_permutation)(unsafe.Pointer(x.Ptr())))
}

func Memcpy(dest *GslPermutation, src *GslPermutation) int32 {
	return int32(C.gsl_permutation_memcpy((*C.gsl_permutation)(unsafe.Pointer(dest.Ptr())), (*C.gsl_permutation)(unsafe.Pointer(src.Ptr()))))
}

func Get(p *GslPermutation, i int) int {
	return int(C.gsl_permutation_get((*C.gsl_permutation)(unsafe.Pointer(p.Ptr())), C.size_t(i)))
}

func Swap(p *GslPermutation, i int, j int) int32 {
	return int32(C.gsl_permutation_swap((*C.gsl_permutation)(unsafe.Pointer(p.Ptr())), C.size_t(i), C.size_t(j)))
}

func Size(p *GslPermutation) int {
	return int(C.gsl_permutation_size((*C.gsl_permutation)(unsafe.Pointer(p.Ptr()))))
}

func Valid(p *GslPermutation) int32 {
	return int32(C.gsl_permutation_valid((*C.gsl_permutation)(unsafe.Pointer(p.Ptr()))))
}

func Reverse(p *GslPermutation) {
	C.gsl_permutation_reverse((*C.gsl_permutation)(unsafe.Pointer(p.Ptr())))
}

func Inverse(inv *GslPermutation, p *GslPermutation) int32 {
	return int32(C.gsl_permutation_inverse((*C.gsl_permutation)(unsafe.Pointer(inv.Ptr())), (*C.gsl_permutation)(unsafe.Pointer(p.Ptr()))))
}

func Next(p *GslPermutation) int32 {
	return int32(C.gsl_permutation_next((*C.gsl_permutation)(unsafe.Pointer(p.Ptr()))))
}

func Prev(p *GslPermutation) int32 {
	return int32(C.gsl_permutation_prev((*C.gsl_permutation)(unsafe.Pointer(p.Ptr()))))
}

func PermuteVector(p *GslPermutation, v *vector.GslVector) int32 {
	return int32(C.gsl_permute_vector((*C.gsl_permutation)(unsafe.Pointer(p.Ptr())), (*C.gsl_vector)(unsafe.Pointer(v.Ptr()))))
}

func PermuteVectorInverse(p *GslPermutation, v *vector.GslVector) int32 {
	return int32(C.gsl_permute_vector_inverse((*C.gsl_permutation)(unsafe.Pointer(p.Ptr())), (*C.gsl_vector)(unsafe.Pointer(v.Ptr()))))
}

func Mul(p *GslPermutation, pa *GslPermutation, pb *GslPermutation) int32 {
	return int32(C.gsl_permutation_mul((*C.gsl_permutation)(unsafe.Pointer(p.Ptr())), (*C.gsl_permutation)(unsafe.Pointer(pa.Ptr())), (*C.gsl_permutation)(unsafe.Pointer(pb.Ptr()))))
}

func Fwrite(stream *os.File, p *GslPermutation) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_permutation_fwrite(_file_0, (*C.gsl_permutation)(unsafe.Pointer(p.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func Fread(stream *os.File, p *GslPermutation) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_permutation_fread(_file_0, (*C.gsl_permutation)(unsafe.Pointer(p.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func Fprintf(stream *os.File, p *GslPermutation, format string) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_string_2 := C.CString(format)
	_result := int32(C.gsl_permutation_fprintf(_file_0, (*C.gsl_permutation)(unsafe.Pointer(p.Ptr())), _string_2))
	C.fclose(_file_0)
	C.free(unsafe.Pointer(_string_2))
	return _result
}

func Fscanf(stream *os.File, p *GslPermutation) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_permutation_fscanf(_file_0, (*C.gsl_permutation)(unsafe.Pointer(p.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func LinearToCanonical(q *GslPermutation, p *GslPermutation) int32 {
	return int32(C.gsl_permutation_linear_to_canonical((*C.gsl_permutation)(unsafe.Pointer(q.Ptr())), (*C.gsl_permutation)(unsafe.Pointer(p.Ptr()))))
}

func CanonicalToLinear(p *GslPermutation, q *GslPermutation) int32 {
	return int32(C.gsl_permutation_canonical_to_linear((*C.gsl_permutation)(unsafe.Pointer(p.Ptr())), (*C.gsl_permutation)(unsafe.Pointer(q.Ptr()))))
}

func Inversions(p *GslPermutation) int {
	return int(C.gsl_permutation_inversions((*C.gsl_permutation)(unsafe.Pointer(p.Ptr()))))
}

func LinearCycles(p *GslPermutation) int {
	return int(C.gsl_permutation_linear_cycles((*C.gsl_permutation)(unsafe.Pointer(p.Ptr()))))
}

func CanonicalCycles(q *GslPermutation) int {
	return int(C.gsl_permutation_canonical_cycles((*C.gsl_permutation)(unsafe.Pointer(q.Ptr()))))
}
