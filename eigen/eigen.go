//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package eigen

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_eigen.h>
*/
import "C"

import "unsafe"
import "github.com/dtromb/gogsl"
import "github.com/dtromb/gogsl/matrix"
import "github.com/dtromb/gogsl/vector"

type GslEigenSymmWorkspace struct {
	gogsl.GslReference
}

type GslEigenSymmvWorkspace struct {
	gogsl.GslReference
}

type GslEigenHermWorkspace struct {
	gogsl.GslReference
}

type GslEigenHermvWorkspace struct {
	gogsl.GslReference
}

type GslEigenNonsymmWorkspace struct {
	gogsl.GslReference
}

type GslEigenNonsymmvWorkspace struct {
	gogsl.GslReference
}

type GslEigenGensymmWorkspace struct {
	gogsl.GslReference
}

type GslEigenGensymmvWorkspace struct {
	gogsl.GslReference
}

type GslEigenGenhermWorkspace struct {
	gogsl.GslReference
}

type GslEigenGenhermvWorkspace struct {
	gogsl.GslReference
}

type GslEigenGenWorkspace struct {
	gogsl.GslReference
}

type GslEigenGenvWorkspace struct {
	gogsl.GslReference
}

func SymmAlloc(n int) *GslEigenSymmWorkspace {
	_ref := C.gsl_eigen_symm_alloc(C.size_t(n))
	_result := &GslEigenSymmWorkspace{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslEigenSymmWorkspace) Dispose() {
	C.gsl_eigen_symm_free((*C.gsl_eigen_symm_workspace)(unsafe.Pointer(x.Ptr())))
}

func Symm(a *matrix.GslMatrix, eval *vector.GslVector, w *GslEigenSymmWorkspace) int32 {
	return int32(C.gsl_eigen_symm((*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector)(unsafe.Pointer(eval.Ptr())), (*C.gsl_eigen_symm_workspace)(unsafe.Pointer(w.Ptr()))))
}

func SymmvAlloc(n int) *GslEigenSymmvWorkspace {
	_ref := C.gsl_eigen_symmv_alloc(C.size_t(n))
	_result := &GslEigenSymmvWorkspace{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslEigenSymmvWorkspace) Dispose() {
	C.gsl_eigen_symmv_free((*C.gsl_eigen_symmv_workspace)(unsafe.Pointer(x.Ptr())))
}

func Symmv(a *matrix.GslMatrix, eval *vector.GslVector, evec *matrix.GslMatrix, w *GslEigenSymmvWorkspace) int32 {
	return int32(C.gsl_eigen_symmv((*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector)(unsafe.Pointer(eval.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(evec.Ptr())), (*C.gsl_eigen_symmv_workspace)(unsafe.Pointer(w.Ptr()))))
}

func HermAlloc(n int) *GslEigenHermWorkspace {
	_ref := C.gsl_eigen_herm_alloc(C.size_t(n))
	_result := &GslEigenHermWorkspace{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslEigenHermWorkspace) Dispose() {
	C.gsl_eigen_herm_free((*C.gsl_eigen_herm_workspace)(unsafe.Pointer(x.Ptr())))
}

func Herm(a *matrix.GslMatrixComplex, eval *vector.GslVector, w *GslEigenHermWorkspace) int32 {
	return int32(C.gsl_eigen_herm((*C.gsl_matrix_complex)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector)(unsafe.Pointer(eval.Ptr())), (*C.gsl_eigen_herm_workspace)(unsafe.Pointer(w.Ptr()))))
}

func HermvAlloc(n int) *GslEigenHermvWorkspace {
	_ref := C.gsl_eigen_hermv_alloc(C.size_t(n))
	_result := &GslEigenHermvWorkspace{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslEigenHermvWorkspace) Dispose() {
	C.gsl_eigen_hermv_free((*C.gsl_eigen_hermv_workspace)(unsafe.Pointer(x.Ptr())))
}

func Hermv(a *matrix.GslMatrixComplex, eval *vector.GslVector, evec *matrix.GslMatrixComplex, w *GslEigenHermvWorkspace) int32 {
	return int32(C.gsl_eigen_hermv((*C.gsl_matrix_complex)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector)(unsafe.Pointer(eval.Ptr())), (*C.gsl_matrix_complex)(unsafe.Pointer(evec.Ptr())), (*C.gsl_eigen_hermv_workspace)(unsafe.Pointer(w.Ptr()))))
}

func NonsymmAlloc(n int) *GslEigenNonsymmWorkspace {
	_ref := C.gsl_eigen_nonsymm_alloc(C.size_t(n))
	_result := &GslEigenNonsymmWorkspace{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslEigenNonsymmWorkspace) Dispose() {
	C.gsl_eigen_nonsymm_free((*C.gsl_eigen_nonsymm_workspace)(unsafe.Pointer(x.Ptr())))
}

func NonsymmParams(computeT int32, balance int32, w *GslEigenNonsymmWorkspace) {
	C.gsl_eigen_nonsymm_params(C.int(computeT), C.int(balance), (*C.gsl_eigen_nonsymm_workspace)(unsafe.Pointer(w.Ptr())))
}

func Nonsymm(a *matrix.GslMatrix, eval *vector.GslVectorComplex, w *GslEigenNonsymmWorkspace) int32 {
	return int32(C.gsl_eigen_nonsymm((*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_complex)(unsafe.Pointer(eval.Ptr())), (*C.gsl_eigen_nonsymm_workspace)(unsafe.Pointer(w.Ptr()))))
}

func NonsymmZ(a *matrix.GslMatrix, eval *vector.GslVectorComplex, z *matrix.GslMatrix, w *GslEigenNonsymmWorkspace) int32 {
	return int32(C.gsl_eigen_nonsymm_Z((*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_complex)(unsafe.Pointer(eval.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(z.Ptr())), (*C.gsl_eigen_nonsymm_workspace)(unsafe.Pointer(w.Ptr()))))
}

func NonsymmvAlloc(n int) *GslEigenNonsymmvWorkspace {
	_ref := C.gsl_eigen_nonsymmv_alloc(C.size_t(n))
	_result := &GslEigenNonsymmvWorkspace{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslEigenNonsymmvWorkspace) Dispose() {
	C.gsl_eigen_nonsymmv_free((*C.gsl_eigen_nonsymmv_workspace)(unsafe.Pointer(x.Ptr())))
}

func NonsymmvParams(balance int32, w *GslEigenNonsymmvWorkspace) {
	C.gsl_eigen_nonsymmv_params(C.int(balance), (*C.gsl_eigen_nonsymmv_workspace)(unsafe.Pointer(w.Ptr())))
}

func Nonsymmv(a *matrix.GslMatrix, eval *vector.GslVectorComplex, evec *matrix.GslMatrixComplex, w *GslEigenNonsymmvWorkspace) int32 {
	return int32(C.gsl_eigen_nonsymmv((*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_complex)(unsafe.Pointer(eval.Ptr())), (*C.gsl_matrix_complex)(unsafe.Pointer(evec.Ptr())), (*C.gsl_eigen_nonsymmv_workspace)(unsafe.Pointer(w.Ptr()))))
}

func NonsymmvZ(a *matrix.GslMatrix, eval *vector.GslVectorComplex, evec *matrix.GslMatrixComplex, z *matrix.GslMatrix, w *GslEigenNonsymmvWorkspace) int32 {
	return int32(C.gsl_eigen_nonsymmv_Z((*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_complex)(unsafe.Pointer(eval.Ptr())), (*C.gsl_matrix_complex)(unsafe.Pointer(evec.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(z.Ptr())), (*C.gsl_eigen_nonsymmv_workspace)(unsafe.Pointer(w.Ptr()))))
}

func GensymmAlloc(n int) *GslEigenGensymmWorkspace {
	_ref := C.gsl_eigen_gensymm_alloc(C.size_t(n))
	_result := &GslEigenGensymmWorkspace{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslEigenGensymmWorkspace) Dispose() {
	C.gsl_eigen_gensymm_free((*C.gsl_eigen_gensymm_workspace)(unsafe.Pointer(x.Ptr())))
}

func Gensymm(a *matrix.GslMatrix, b *matrix.GslMatrix, eval *vector.GslVector, w *GslEigenGensymmWorkspace) int32 {
	return int32(C.gsl_eigen_gensymm((*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(b.Ptr())), (*C.gsl_vector)(unsafe.Pointer(eval.Ptr())), (*C.gsl_eigen_gensymm_workspace)(unsafe.Pointer(w.Ptr()))))
}

func GensymmvAlloc(n int) *GslEigenGensymmvWorkspace {
	_ref := C.gsl_eigen_gensymmv_alloc(C.size_t(n))
	_result := &GslEigenGensymmvWorkspace{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslEigenGensymmvWorkspace) Dispose() {
	C.gsl_eigen_gensymmv_free((*C.gsl_eigen_gensymmv_workspace)(unsafe.Pointer(x.Ptr())))
}

func Gensymmv(a *matrix.GslMatrix, b *matrix.GslMatrix, eval *vector.GslVector, evec *matrix.GslMatrix, w *GslEigenGensymmvWorkspace) int32 {
	return int32(C.gsl_eigen_gensymmv((*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(b.Ptr())), (*C.gsl_vector)(unsafe.Pointer(eval.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(evec.Ptr())), (*C.gsl_eigen_gensymmv_workspace)(unsafe.Pointer(w.Ptr()))))
}

func GenhermAlloc(n int) *GslEigenGenhermWorkspace {
	_ref := C.gsl_eigen_genherm_alloc(C.size_t(n))
	_result := &GslEigenGenhermWorkspace{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslEigenGenhermWorkspace) Dispose() {
	C.gsl_eigen_genherm_free((*C.gsl_eigen_genherm_workspace)(unsafe.Pointer(x.Ptr())))
}

func Genherm(a *matrix.GslMatrixComplex, b *matrix.GslMatrixComplex, eval *vector.GslVector, w *GslEigenGenhermWorkspace) int32 {
	return int32(C.gsl_eigen_genherm((*C.gsl_matrix_complex)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_complex)(unsafe.Pointer(b.Ptr())), (*C.gsl_vector)(unsafe.Pointer(eval.Ptr())), (*C.gsl_eigen_genherm_workspace)(unsafe.Pointer(w.Ptr()))))
}

func GenhermvAlloc(n int) *GslEigenGenhermvWorkspace {
	_ref := C.gsl_eigen_genhermv_alloc(C.size_t(n))
	_result := &GslEigenGenhermvWorkspace{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslEigenGenhermvWorkspace) Dispose() {
	C.gsl_eigen_genhermv_free((*C.gsl_eigen_genhermv_workspace)(unsafe.Pointer(x.Ptr())))
}

func Genhermv(a *matrix.GslMatrixComplex, b *matrix.GslMatrixComplex, eval *vector.GslVector, evec *matrix.GslMatrixComplex, w *GslEigenGenhermvWorkspace) int32 {
	return int32(C.gsl_eigen_genhermv((*C.gsl_matrix_complex)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_complex)(unsafe.Pointer(b.Ptr())), (*C.gsl_vector)(unsafe.Pointer(eval.Ptr())), (*C.gsl_matrix_complex)(unsafe.Pointer(evec.Ptr())), (*C.gsl_eigen_genhermv_workspace)(unsafe.Pointer(w.Ptr()))))
}

func GenAlloc(n int) *GslEigenGenWorkspace {
	_ref := C.gsl_eigen_gen_alloc(C.size_t(n))
	_result := &GslEigenGenWorkspace{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslEigenGenWorkspace) Dispose() {
	C.gsl_eigen_gen_free((*C.gsl_eigen_gen_workspace)(unsafe.Pointer(x.Ptr())))
}

func GenParams(computeS int32, computeT int32, balance int32, w *GslEigenGenWorkspace) {
	C.gsl_eigen_gen_params(C.int(computeS), C.int(computeT), C.int(balance), (*C.gsl_eigen_gen_workspace)(unsafe.Pointer(w.Ptr())))
}

func Gen(a *matrix.GslMatrix, b *matrix.GslMatrix, alpha *vector.GslVectorComplex, beta *vector.GslVector, w *GslEigenGenWorkspace) int32 {
	return int32(C.gsl_eigen_gen((*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(b.Ptr())), (*C.gsl_vector_complex)(unsafe.Pointer(alpha.Ptr())), (*C.gsl_vector)(unsafe.Pointer(beta.Ptr())), (*C.gsl_eigen_gen_workspace)(unsafe.Pointer(w.Ptr()))))
}

func GenQZ(a *matrix.GslMatrix, b *matrix.GslMatrix, alpha *vector.GslVectorComplex, beta *vector.GslVector, q *matrix.GslMatrix, z *matrix.GslMatrix, w *GslEigenGenWorkspace) int32 {
	return int32(C.gsl_eigen_gen_QZ((*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(b.Ptr())), (*C.gsl_vector_complex)(unsafe.Pointer(alpha.Ptr())), (*C.gsl_vector)(unsafe.Pointer(beta.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(q.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(z.Ptr())), (*C.gsl_eigen_gen_workspace)(unsafe.Pointer(w.Ptr()))))
}

func GenvAlloc(n int) *GslEigenGenvWorkspace {
	_ref := C.gsl_eigen_genv_alloc(C.size_t(n))
	_result := &GslEigenGenvWorkspace{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslEigenGenvWorkspace) Dispose() {
	C.gsl_eigen_genv_free((*C.gsl_eigen_genv_workspace)(unsafe.Pointer(x.Ptr())))
}

func Genv(a *matrix.GslMatrix, b *matrix.GslMatrix, alpha *vector.GslVectorComplex, beta *vector.GslVector, evec *matrix.GslMatrixComplex, w *GslEigenGenvWorkspace) int32 {
	return int32(C.gsl_eigen_genv((*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(b.Ptr())), (*C.gsl_vector_complex)(unsafe.Pointer(alpha.Ptr())), (*C.gsl_vector)(unsafe.Pointer(beta.Ptr())), (*C.gsl_matrix_complex)(unsafe.Pointer(evec.Ptr())), (*C.gsl_eigen_genv_workspace)(unsafe.Pointer(w.Ptr()))))
}

func GenvQZ(a *matrix.GslMatrix, b *matrix.GslMatrix, alpha *vector.GslVectorComplex, beta *vector.GslVector, evec *matrix.GslMatrixComplex, q *matrix.GslMatrix, z *matrix.GslMatrix, w *GslEigenGenvWorkspace) int32 {
	return int32(C.gsl_eigen_genv_QZ((*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(b.Ptr())), (*C.gsl_vector_complex)(unsafe.Pointer(alpha.Ptr())), (*C.gsl_vector)(unsafe.Pointer(beta.Ptr())), (*C.gsl_matrix_complex)(unsafe.Pointer(evec.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(q.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(z.Ptr())), (*C.gsl_eigen_genv_workspace)(unsafe.Pointer(w.Ptr()))))
}

func SymmvSort(eval *vector.GslVector, evec *matrix.GslMatrix, sortType Sort) int32 {
	return int32(C.gsl_eigen_symmv_sort((*C.gsl_vector)(unsafe.Pointer(eval.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(evec.Ptr())), C.gsl_eigen_sort_t(sortType)))
}

func HermvSort(eval *vector.GslVector, evec *matrix.GslMatrixComplex, sortType Sort) int32 {
	return int32(C.gsl_eigen_hermv_sort((*C.gsl_vector)(unsafe.Pointer(eval.Ptr())), (*C.gsl_matrix_complex)(unsafe.Pointer(evec.Ptr())), C.gsl_eigen_sort_t(sortType)))
}

func NonsymmvSort(eval *vector.GslVectorComplex, evec *matrix.GslMatrixComplex, sortType Sort) int32 {
	return int32(C.gsl_eigen_nonsymmv_sort((*C.gsl_vector_complex)(unsafe.Pointer(eval.Ptr())), (*C.gsl_matrix_complex)(unsafe.Pointer(evec.Ptr())), C.gsl_eigen_sort_t(sortType)))
}

func GensymmvSort(eval *vector.GslVector, evec *matrix.GslMatrix, sortType Sort) int32 {
	return int32(C.gsl_eigen_gensymmv_sort((*C.gsl_vector)(unsafe.Pointer(eval.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(evec.Ptr())), C.gsl_eigen_sort_t(sortType)))
}

func GenhermvSort(eval *vector.GslVector, evec *matrix.GslMatrixComplex, sortType Sort) int32 {
	return int32(C.gsl_eigen_genhermv_sort((*C.gsl_vector)(unsafe.Pointer(eval.Ptr())), (*C.gsl_matrix_complex)(unsafe.Pointer(evec.Ptr())), C.gsl_eigen_sort_t(sortType)))
}

func GenvSort(alpha *vector.GslVectorComplex, beta *vector.GslVector, evec *matrix.GslMatrixComplex, sortType Sort) int32 {
	return int32(C.gsl_eigen_genv_sort((*C.gsl_vector_complex)(unsafe.Pointer(alpha.Ptr())), (*C.gsl_vector)(unsafe.Pointer(beta.Ptr())), (*C.gsl_matrix_complex)(unsafe.Pointer(evec.Ptr())), C.gsl_eigen_sort_t(sortType)))
}
