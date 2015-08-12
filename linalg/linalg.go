//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package linalg

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_linalg.h>
*/
import "C"

import "github.com/dtromb/gogsl/matrix"
import "github.com/dtromb/gogsl/permutation"
import "unsafe"
import "github.com/dtromb/gogsl/vector"
import complex_ "github.com/dtromb/gogsl/complex"

func LUDecomp(a *matrix.GslMatrix, p *permutation.GslPermutation) (int32, int32) {
	var _outptr_2 C.int
	_result := int32(C.gsl_linalg_LU_decomp((*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_permutation)(unsafe.Pointer(p.Ptr())), &_outptr_2))
	return _result, *(*int32)(unsafe.Pointer(&_outptr_2))
}

func ComplexLUDecomp(a *matrix.GslMatrixComplex, p *permutation.GslPermutation) (int32, int32) {
	var _outptr_2 C.int
	_result := int32(C.gsl_linalg_complex_LU_decomp((*C.gsl_matrix_complex)(unsafe.Pointer(a.Ptr())), (*C.gsl_permutation)(unsafe.Pointer(p.Ptr())), &_outptr_2))
	return _result, *(*int32)(unsafe.Pointer(&_outptr_2))
}

func LUSolve(lU *matrix.GslMatrix, p *permutation.GslPermutation, b *vector.GslVector, x *vector.GslVector) int32 {
	return int32(C.gsl_linalg_LU_solve((*C.gsl_matrix)(unsafe.Pointer(lU.Ptr())), (*C.gsl_permutation)(unsafe.Pointer(p.Ptr())), (*C.gsl_vector)(unsafe.Pointer(b.Ptr())), (*C.gsl_vector)(unsafe.Pointer(x.Ptr()))))
}

func ComplexLUSolve(lU *matrix.GslMatrixComplex, p *permutation.GslPermutation, b *vector.GslVectorComplex, x *vector.GslVectorComplex) int32 {
	return int32(C.gsl_linalg_complex_LU_solve((*C.gsl_matrix_complex)(unsafe.Pointer(lU.Ptr())), (*C.gsl_permutation)(unsafe.Pointer(p.Ptr())), (*C.gsl_vector_complex)(unsafe.Pointer(b.Ptr())), (*C.gsl_vector_complex)(unsafe.Pointer(x.Ptr()))))
}

func LUSvx(lU *matrix.GslMatrix, p *permutation.GslPermutation, x *vector.GslVector) int32 {
	return int32(C.gsl_linalg_LU_svx((*C.gsl_matrix)(unsafe.Pointer(lU.Ptr())), (*C.gsl_permutation)(unsafe.Pointer(p.Ptr())), (*C.gsl_vector)(unsafe.Pointer(x.Ptr()))))
}

func ComplexLUSvx(lU *matrix.GslMatrixComplex, p *permutation.GslPermutation, x *vector.GslVectorComplex) int32 {
	return int32(C.gsl_linalg_complex_LU_svx((*C.gsl_matrix_complex)(unsafe.Pointer(lU.Ptr())), (*C.gsl_permutation)(unsafe.Pointer(p.Ptr())), (*C.gsl_vector_complex)(unsafe.Pointer(x.Ptr()))))
}

func LURefine(a *matrix.GslMatrix, lU *matrix.GslMatrix, p *permutation.GslPermutation, b *vector.GslVector, x *vector.GslVector, residual *vector.GslVector) int32 {
	return int32(C.gsl_linalg_LU_refine((*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(lU.Ptr())), (*C.gsl_permutation)(unsafe.Pointer(p.Ptr())), (*C.gsl_vector)(unsafe.Pointer(b.Ptr())), (*C.gsl_vector)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector)(unsafe.Pointer(residual.Ptr()))))
}

func ComplexLURefine(a *matrix.GslMatrixComplex, lU *matrix.GslMatrixComplex, p *permutation.GslPermutation, b *vector.GslVectorComplex, x *vector.GslVectorComplex, residual *vector.GslVectorComplex) int32 {
	return int32(C.gsl_linalg_complex_LU_refine((*C.gsl_matrix_complex)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_complex)(unsafe.Pointer(lU.Ptr())), (*C.gsl_permutation)(unsafe.Pointer(p.Ptr())), (*C.gsl_vector_complex)(unsafe.Pointer(b.Ptr())), (*C.gsl_vector_complex)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector_complex)(unsafe.Pointer(residual.Ptr()))))
}

func LUInvert(lU *matrix.GslMatrix, p *permutation.GslPermutation, inverse *matrix.GslMatrix) int32 {
	return int32(C.gsl_linalg_LU_invert((*C.gsl_matrix)(unsafe.Pointer(lU.Ptr())), (*C.gsl_permutation)(unsafe.Pointer(p.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(inverse.Ptr()))))
}

func ComplexLUInvert(lU *matrix.GslMatrixComplex, p *permutation.GslPermutation, inverse *matrix.GslMatrixComplex) int32 {
	return int32(C.gsl_linalg_complex_LU_invert((*C.gsl_matrix_complex)(unsafe.Pointer(lU.Ptr())), (*C.gsl_permutation)(unsafe.Pointer(p.Ptr())), (*C.gsl_matrix_complex)(unsafe.Pointer(inverse.Ptr()))))
}

func LUDet(lU *matrix.GslMatrix, signum int32) float64 {
	return float64(C.gsl_linalg_LU_det((*C.gsl_matrix)(unsafe.Pointer(lU.Ptr())), C.int(signum)))
}

func ComplexLUDet(lU *matrix.GslMatrixComplex, signum int32) complex128 {
	_result := C.gsl_linalg_complex_LU_det((*C.gsl_matrix_complex)(unsafe.Pointer(lU.Ptr())), C.int(signum))
	return complex_.GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func LULndet(lU *matrix.GslMatrix) float64 {
	return float64(C.gsl_linalg_LU_lndet((*C.gsl_matrix)(unsafe.Pointer(lU.Ptr()))))
}

func ComplexLULndet(lU *matrix.GslMatrixComplex) float64 {
	return float64(C.gsl_linalg_complex_LU_lndet((*C.gsl_matrix_complex)(unsafe.Pointer(lU.Ptr()))))
}

func LUSgndet(lU *matrix.GslMatrix, signum int32) int32 {
	return int32(C.gsl_linalg_LU_sgndet((*C.gsl_matrix)(unsafe.Pointer(lU.Ptr())), C.int(signum)))
}

func ComplexLUSgndet(lU *matrix.GslMatrixComplex, signum int32) complex128 {
	_result := C.gsl_linalg_complex_LU_sgndet((*C.gsl_matrix_complex)(unsafe.Pointer(lU.Ptr())), C.int(signum))
	return complex_.GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func QRDecomp(a *matrix.GslMatrix, tau *vector.GslVector) int32 {
	return int32(C.gsl_linalg_QR_decomp((*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector)(unsafe.Pointer(tau.Ptr()))))
}

func QRSolve(qR *matrix.GslMatrix, tau *vector.GslVector, b *vector.GslVector, x *vector.GslVector) int32 {
	return int32(C.gsl_linalg_QR_solve((*C.gsl_matrix)(unsafe.Pointer(qR.Ptr())), (*C.gsl_vector)(unsafe.Pointer(tau.Ptr())), (*C.gsl_vector)(unsafe.Pointer(b.Ptr())), (*C.gsl_vector)(unsafe.Pointer(x.Ptr()))))
}

func QRSvx(qR *matrix.GslMatrix, tau *vector.GslVector, x *vector.GslVector) int32 {
	return int32(C.gsl_linalg_QR_svx((*C.gsl_matrix)(unsafe.Pointer(qR.Ptr())), (*C.gsl_vector)(unsafe.Pointer(tau.Ptr())), (*C.gsl_vector)(unsafe.Pointer(x.Ptr()))))
}

func QRLssolve(qR *matrix.GslMatrix, tau *vector.GslVector, b *vector.GslVector, x *vector.GslVector, residual *vector.GslVector) int32 {
	return int32(C.gsl_linalg_QR_lssolve((*C.gsl_matrix)(unsafe.Pointer(qR.Ptr())), (*C.gsl_vector)(unsafe.Pointer(tau.Ptr())), (*C.gsl_vector)(unsafe.Pointer(b.Ptr())), (*C.gsl_vector)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector)(unsafe.Pointer(residual.Ptr()))))
}

func QRQTvec(qR *matrix.GslMatrix, tau *vector.GslVector, v *vector.GslVector) int32 {
	return int32(C.gsl_linalg_QR_QTvec((*C.gsl_matrix)(unsafe.Pointer(qR.Ptr())), (*C.gsl_vector)(unsafe.Pointer(tau.Ptr())), (*C.gsl_vector)(unsafe.Pointer(v.Ptr()))))
}

func QRQvec(qR *matrix.GslMatrix, tau *vector.GslVector, v *vector.GslVector) int32 {
	return int32(C.gsl_linalg_QR_Qvec((*C.gsl_matrix)(unsafe.Pointer(qR.Ptr())), (*C.gsl_vector)(unsafe.Pointer(tau.Ptr())), (*C.gsl_vector)(unsafe.Pointer(v.Ptr()))))
}

func QRQTmat(qR *matrix.GslMatrix, tau *vector.GslVector, a *matrix.GslMatrix) int32 {
	return int32(C.gsl_linalg_QR_QTmat((*C.gsl_matrix)(unsafe.Pointer(qR.Ptr())), (*C.gsl_vector)(unsafe.Pointer(tau.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(a.Ptr()))))
}

func QRRsolve(qR *matrix.GslMatrix, b *vector.GslVector, x *vector.GslVector) int32 {
	return int32(C.gsl_linalg_QR_Rsolve((*C.gsl_matrix)(unsafe.Pointer(qR.Ptr())), (*C.gsl_vector)(unsafe.Pointer(b.Ptr())), (*C.gsl_vector)(unsafe.Pointer(x.Ptr()))))
}

func QRRsvx(qR *matrix.GslMatrix, x *vector.GslVector) int32 {
	return int32(C.gsl_linalg_QR_Rsvx((*C.gsl_matrix)(unsafe.Pointer(qR.Ptr())), (*C.gsl_vector)(unsafe.Pointer(x.Ptr()))))
}

func QRUnpack(qR *matrix.GslMatrix, tau *vector.GslVector, q *matrix.GslMatrix, r *matrix.GslMatrix) int32 {
	return int32(C.gsl_linalg_QR_unpack((*C.gsl_matrix)(unsafe.Pointer(qR.Ptr())), (*C.gsl_vector)(unsafe.Pointer(tau.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(q.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(r.Ptr()))))
}

func QRQRsolve(q *matrix.GslMatrix, r *matrix.GslMatrix, b *vector.GslVector, x *vector.GslVector) int32 {
	return int32(C.gsl_linalg_QR_QRsolve((*C.gsl_matrix)(unsafe.Pointer(q.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(r.Ptr())), (*C.gsl_vector)(unsafe.Pointer(b.Ptr())), (*C.gsl_vector)(unsafe.Pointer(x.Ptr()))))
}

func QRUpdate(q *matrix.GslMatrix, r *matrix.GslMatrix, w *vector.GslVector, v *vector.GslVector) int32 {
	return int32(C.gsl_linalg_QR_update((*C.gsl_matrix)(unsafe.Pointer(q.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(r.Ptr())), (*C.gsl_vector)(unsafe.Pointer(w.Ptr())), (*C.gsl_vector)(unsafe.Pointer(v.Ptr()))))
}

func RSolve(r *matrix.GslMatrix, b *vector.GslVector, x *vector.GslVector) int32 {
	return int32(C.gsl_linalg_R_solve((*C.gsl_matrix)(unsafe.Pointer(r.Ptr())), (*C.gsl_vector)(unsafe.Pointer(b.Ptr())), (*C.gsl_vector)(unsafe.Pointer(x.Ptr()))))
}

func RSvx(r *matrix.GslMatrix, x *vector.GslVector) int32 {
	return int32(C.gsl_linalg_R_svx((*C.gsl_matrix)(unsafe.Pointer(r.Ptr())), (*C.gsl_vector)(unsafe.Pointer(x.Ptr()))))
}

func QRPTDecomp(a *matrix.GslMatrix, tau *vector.GslVector, p *permutation.GslPermutation, norm *vector.GslVector) (int32, int32) {
	var _outptr_3 C.int
	_result := int32(C.gsl_linalg_QRPT_decomp((*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector)(unsafe.Pointer(tau.Ptr())), (*C.gsl_permutation)(unsafe.Pointer(p.Ptr())), &_outptr_3, (*C.gsl_vector)(unsafe.Pointer(norm.Ptr()))))
	return _result, *(*int32)(unsafe.Pointer(&_outptr_3))
}

func QRPTDecomp2(a *matrix.GslMatrix, q *matrix.GslMatrix, r *matrix.GslMatrix, tau *vector.GslVector, p *permutation.GslPermutation, norm *vector.GslVector) (int32, int32) {
	var _outptr_5 C.int
	_result := int32(C.gsl_linalg_QRPT_decomp2((*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(q.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(r.Ptr())), (*C.gsl_vector)(unsafe.Pointer(tau.Ptr())), (*C.gsl_permutation)(unsafe.Pointer(p.Ptr())), &_outptr_5, (*C.gsl_vector)(unsafe.Pointer(norm.Ptr()))))
	return _result, *(*int32)(unsafe.Pointer(&_outptr_5))
}

func QRPTSolve(qR *matrix.GslMatrix, tau *vector.GslVector, p *permutation.GslPermutation, b *vector.GslVector, x *vector.GslVector) int32 {
	return int32(C.gsl_linalg_QRPT_solve((*C.gsl_matrix)(unsafe.Pointer(qR.Ptr())), (*C.gsl_vector)(unsafe.Pointer(tau.Ptr())), (*C.gsl_permutation)(unsafe.Pointer(p.Ptr())), (*C.gsl_vector)(unsafe.Pointer(b.Ptr())), (*C.gsl_vector)(unsafe.Pointer(x.Ptr()))))
}

func QRPTSvx(qR *matrix.GslMatrix, tau *vector.GslVector, p *permutation.GslPermutation, x *vector.GslVector) int32 {
	return int32(C.gsl_linalg_QRPT_svx((*C.gsl_matrix)(unsafe.Pointer(qR.Ptr())), (*C.gsl_vector)(unsafe.Pointer(tau.Ptr())), (*C.gsl_permutation)(unsafe.Pointer(p.Ptr())), (*C.gsl_vector)(unsafe.Pointer(x.Ptr()))))
}

func QRPTQRsolve(q *matrix.GslMatrix, r *matrix.GslMatrix, p *permutation.GslPermutation, b *vector.GslVector, x *vector.GslVector) int32 {
	return int32(C.gsl_linalg_QRPT_QRsolve((*C.gsl_matrix)(unsafe.Pointer(q.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(r.Ptr())), (*C.gsl_permutation)(unsafe.Pointer(p.Ptr())), (*C.gsl_vector)(unsafe.Pointer(b.Ptr())), (*C.gsl_vector)(unsafe.Pointer(x.Ptr()))))
}

func QRPTUpdate(q *matrix.GslMatrix, r *matrix.GslMatrix, p *permutation.GslPermutation, w *vector.GslVector, v *vector.GslVector) int32 {
	return int32(C.gsl_linalg_QRPT_update((*C.gsl_matrix)(unsafe.Pointer(q.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(r.Ptr())), (*C.gsl_permutation)(unsafe.Pointer(p.Ptr())), (*C.gsl_vector)(unsafe.Pointer(w.Ptr())), (*C.gsl_vector)(unsafe.Pointer(v.Ptr()))))
}

func QRPTRsolve(qR *matrix.GslMatrix, p *permutation.GslPermutation, b *vector.GslVector, x *vector.GslVector) int32 {
	return int32(C.gsl_linalg_QRPT_Rsolve((*C.gsl_matrix)(unsafe.Pointer(qR.Ptr())), (*C.gsl_permutation)(unsafe.Pointer(p.Ptr())), (*C.gsl_vector)(unsafe.Pointer(b.Ptr())), (*C.gsl_vector)(unsafe.Pointer(x.Ptr()))))
}

func QRPTRsvx(qR *matrix.GslMatrix, p *permutation.GslPermutation, x *vector.GslVector) int32 {
	return int32(C.gsl_linalg_QRPT_Rsvx((*C.gsl_matrix)(unsafe.Pointer(qR.Ptr())), (*C.gsl_permutation)(unsafe.Pointer(p.Ptr())), (*C.gsl_vector)(unsafe.Pointer(x.Ptr()))))
}

func SVDecomp(a *matrix.GslMatrix, v *matrix.GslMatrix, s *vector.GslVector, work *vector.GslVector) int32 {
	return int32(C.gsl_linalg_SV_decomp((*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(v.Ptr())), (*C.gsl_vector)(unsafe.Pointer(s.Ptr())), (*C.gsl_vector)(unsafe.Pointer(work.Ptr()))))
}

func SVDecompMod(a *matrix.GslMatrix, x *matrix.GslMatrix, v *matrix.GslMatrix, s *vector.GslVector, work *vector.GslVector) int32 {
	return int32(C.gsl_linalg_SV_decomp_mod((*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(x.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(v.Ptr())), (*C.gsl_vector)(unsafe.Pointer(s.Ptr())), (*C.gsl_vector)(unsafe.Pointer(work.Ptr()))))
}

func SVDecompJacobi(a *matrix.GslMatrix, v *matrix.GslMatrix, s *vector.GslVector) int32 {
	return int32(C.gsl_linalg_SV_decomp_jacobi((*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(v.Ptr())), (*C.gsl_vector)(unsafe.Pointer(s.Ptr()))))
}

func SVSolve(u *matrix.GslMatrix, v *matrix.GslMatrix, s *vector.GslVector, b *vector.GslVector, x *vector.GslVector) int32 {
	return int32(C.gsl_linalg_SV_solve((*C.gsl_matrix)(unsafe.Pointer(u.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(v.Ptr())), (*C.gsl_vector)(unsafe.Pointer(s.Ptr())), (*C.gsl_vector)(unsafe.Pointer(b.Ptr())), (*C.gsl_vector)(unsafe.Pointer(x.Ptr()))))
}

func SVLeverage(u *matrix.GslMatrix, h *vector.GslVector) int32 {
	return int32(C.gsl_linalg_SV_leverage((*C.gsl_matrix)(unsafe.Pointer(u.Ptr())), (*C.gsl_vector)(unsafe.Pointer(h.Ptr()))))
}

func CholeskyDecomp(a *matrix.GslMatrix) int32 {
	return int32(C.gsl_linalg_cholesky_decomp((*C.gsl_matrix)(unsafe.Pointer(a.Ptr()))))
}

func ComplexCholeskyDecomp(a *matrix.GslMatrixComplex) int32 {
	return int32(C.gsl_linalg_complex_cholesky_decomp((*C.gsl_matrix_complex)(unsafe.Pointer(a.Ptr()))))
}

func CholeskySolve(cholesky *matrix.GslMatrix, b *vector.GslVector, x *vector.GslVector) int32 {
	return int32(C.gsl_linalg_cholesky_solve((*C.gsl_matrix)(unsafe.Pointer(cholesky.Ptr())), (*C.gsl_vector)(unsafe.Pointer(b.Ptr())), (*C.gsl_vector)(unsafe.Pointer(x.Ptr()))))
}

func ComplexCholeskySolve(cholesky *matrix.GslMatrixComplex, b *vector.GslVectorComplex, x *vector.GslVectorComplex) int32 {
	return int32(C.gsl_linalg_complex_cholesky_solve((*C.gsl_matrix_complex)(unsafe.Pointer(cholesky.Ptr())), (*C.gsl_vector_complex)(unsafe.Pointer(b.Ptr())), (*C.gsl_vector_complex)(unsafe.Pointer(x.Ptr()))))
}

func CholeskySvx(cholesky *matrix.GslMatrix, x *vector.GslVector) int32 {
	return int32(C.gsl_linalg_cholesky_svx((*C.gsl_matrix)(unsafe.Pointer(cholesky.Ptr())), (*C.gsl_vector)(unsafe.Pointer(x.Ptr()))))
}

func ComplexCholeskySvx(cholesky *matrix.GslMatrixComplex, x *vector.GslVectorComplex) int32 {
	return int32(C.gsl_linalg_complex_cholesky_svx((*C.gsl_matrix_complex)(unsafe.Pointer(cholesky.Ptr())), (*C.gsl_vector_complex)(unsafe.Pointer(x.Ptr()))))
}

func CholeskyInvert(cholesky *matrix.GslMatrix) int32 {
	return int32(C.gsl_linalg_cholesky_invert((*C.gsl_matrix)(unsafe.Pointer(cholesky.Ptr()))))
}

func ComplexCholeskyInvert(cholesky *matrix.GslMatrixComplex) int32 {
	return int32(C.gsl_linalg_complex_cholesky_invert((*C.gsl_matrix_complex)(unsafe.Pointer(cholesky.Ptr()))))
}

func SymmtdDecomp(a *matrix.GslMatrix, tau *vector.GslVector) int32 {
	return int32(C.gsl_linalg_symmtd_decomp((*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector)(unsafe.Pointer(tau.Ptr()))))
}

func SymmtdUnpack(a *matrix.GslMatrix, tau *vector.GslVector, q *matrix.GslMatrix, diag *vector.GslVector, subdiag *vector.GslVector) int32 {
	return int32(C.gsl_linalg_symmtd_unpack((*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector)(unsafe.Pointer(tau.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(q.Ptr())), (*C.gsl_vector)(unsafe.Pointer(diag.Ptr())), (*C.gsl_vector)(unsafe.Pointer(subdiag.Ptr()))))
}

func SymmtdUnpackT(a *matrix.GslMatrix, diag *vector.GslVector, subdiag *vector.GslVector) int32 {
	return int32(C.gsl_linalg_symmtd_unpack_T((*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector)(unsafe.Pointer(diag.Ptr())), (*C.gsl_vector)(unsafe.Pointer(subdiag.Ptr()))))
}

func HermtdDecomp(a *matrix.GslMatrixComplex, tau *vector.GslVectorComplex) int32 {
	return int32(C.gsl_linalg_hermtd_decomp((*C.gsl_matrix_complex)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_complex)(unsafe.Pointer(tau.Ptr()))))
}

func HermtdUnpack(a *matrix.GslMatrixComplex, tau *vector.GslVectorComplex, u *matrix.GslMatrixComplex, diag *vector.GslVector, subdiag *vector.GslVector) int32 {
	return int32(C.gsl_linalg_hermtd_unpack((*C.gsl_matrix_complex)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_complex)(unsafe.Pointer(tau.Ptr())), (*C.gsl_matrix_complex)(unsafe.Pointer(u.Ptr())), (*C.gsl_vector)(unsafe.Pointer(diag.Ptr())), (*C.gsl_vector)(unsafe.Pointer(subdiag.Ptr()))))
}

func HermtdUnpackT(a *matrix.GslMatrixComplex, diag *vector.GslVector, subdiag *vector.GslVector) int32 {
	return int32(C.gsl_linalg_hermtd_unpack_T((*C.gsl_matrix_complex)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector)(unsafe.Pointer(diag.Ptr())), (*C.gsl_vector)(unsafe.Pointer(subdiag.Ptr()))))
}

func HessenbergDecomp(a *matrix.GslMatrix, tau *vector.GslVector) int32 {
	return int32(C.gsl_linalg_hessenberg_decomp((*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector)(unsafe.Pointer(tau.Ptr()))))
}

func HessenbergUnpack(h *matrix.GslMatrix, tau *vector.GslVector, u *matrix.GslMatrix) int32 {
	return int32(C.gsl_linalg_hessenberg_unpack((*C.gsl_matrix)(unsafe.Pointer(h.Ptr())), (*C.gsl_vector)(unsafe.Pointer(tau.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(u.Ptr()))))
}

func HessenbergUnpackAccum(h *matrix.GslMatrix, tau *vector.GslVector, v *matrix.GslMatrix) int32 {
	return int32(C.gsl_linalg_hessenberg_unpack_accum((*C.gsl_matrix)(unsafe.Pointer(h.Ptr())), (*C.gsl_vector)(unsafe.Pointer(tau.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(v.Ptr()))))
}

func HessenbergSetZero(h *matrix.GslMatrix) int32 {
	return int32(C.gsl_linalg_hessenberg_set_zero((*C.gsl_matrix)(unsafe.Pointer(h.Ptr()))))
}

func HesstriDecomp(a *matrix.GslMatrix, b *matrix.GslMatrix, u *matrix.GslMatrix, v *matrix.GslMatrix, work *vector.GslVector) int32 {
	return int32(C.gsl_linalg_hesstri_decomp((*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(b.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(u.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(v.Ptr())), (*C.gsl_vector)(unsafe.Pointer(work.Ptr()))))
}

func BidiagDecomp(a *matrix.GslMatrix, tauU *vector.GslVector, tauV *vector.GslVector) int32 {
	return int32(C.gsl_linalg_bidiag_decomp((*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector)(unsafe.Pointer(tauU.Ptr())), (*C.gsl_vector)(unsafe.Pointer(tauV.Ptr()))))
}

func BidiagUnpack(a *matrix.GslMatrix, tauU *vector.GslVector, u *matrix.GslMatrix, tauV *vector.GslVector, v *matrix.GslMatrix, diag *vector.GslVector, superdiag *vector.GslVector) int32 {
	return int32(C.gsl_linalg_bidiag_unpack((*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector)(unsafe.Pointer(tauU.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(u.Ptr())), (*C.gsl_vector)(unsafe.Pointer(tauV.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(v.Ptr())), (*C.gsl_vector)(unsafe.Pointer(diag.Ptr())), (*C.gsl_vector)(unsafe.Pointer(superdiag.Ptr()))))
}

func BidiagUnpack2(a *matrix.GslMatrix, tauU *vector.GslVector, tauV *vector.GslVector, v *matrix.GslMatrix) int32 {
	return int32(C.gsl_linalg_bidiag_unpack2((*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector)(unsafe.Pointer(tauU.Ptr())), (*C.gsl_vector)(unsafe.Pointer(tauV.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(v.Ptr()))))
}

func BidiagUnpackB(a *matrix.GslMatrix, diag *vector.GslVector, superdiag *vector.GslVector) int32 {
	return int32(C.gsl_linalg_bidiag_unpack_B((*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector)(unsafe.Pointer(diag.Ptr())), (*C.gsl_vector)(unsafe.Pointer(superdiag.Ptr()))))
}

func HouseholderTransform(v *vector.GslVector) float64 {
	return float64(C.gsl_linalg_householder_transform((*C.gsl_vector)(unsafe.Pointer(v.Ptr()))))
}

func ComplexHouseholderTransform(v *vector.GslVectorComplex) complex128 {
	_result := C.gsl_linalg_complex_householder_transform((*C.gsl_vector_complex)(unsafe.Pointer(v.Ptr())))
	return complex_.GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func HouseholderHm(tau float64, v *vector.GslVector, a *matrix.GslMatrix) int32 {
	return int32(C.gsl_linalg_householder_hm(C.double(tau), (*C.gsl_vector)(unsafe.Pointer(v.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(a.Ptr()))))
}

func ComplexHouseholderHm(tau complex128, v *vector.GslVectorComplex, a *matrix.GslMatrixComplex) int32 {
	_arg_0 := complex_.GoComplexToGsl(tau)
	return int32(C.gsl_linalg_complex_householder_hm(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)), (*C.gsl_vector_complex)(unsafe.Pointer(v.Ptr())), (*C.gsl_matrix_complex)(unsafe.Pointer(a.Ptr()))))
}

func HouseholderMh(tau float64, v *vector.GslVector, a *matrix.GslMatrix) int32 {
	return int32(C.gsl_linalg_householder_mh(C.double(tau), (*C.gsl_vector)(unsafe.Pointer(v.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(a.Ptr()))))
}

func ComplexHouseholderMh(tau complex128, v *vector.GslVectorComplex, a *matrix.GslMatrixComplex) int32 {
	_arg_0 := complex_.GoComplexToGsl(tau)
	return int32(C.gsl_linalg_complex_householder_mh(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)), (*C.gsl_vector_complex)(unsafe.Pointer(v.Ptr())), (*C.gsl_matrix_complex)(unsafe.Pointer(a.Ptr()))))
}

func HouseholderHv(tau float64, v *vector.GslVector, w *vector.GslVector) int32 {
	return int32(C.gsl_linalg_householder_hv(C.double(tau), (*C.gsl_vector)(unsafe.Pointer(v.Ptr())), (*C.gsl_vector)(unsafe.Pointer(w.Ptr()))))
}

func ComplexHouseholderHv(tau complex128, v *vector.GslVectorComplex, w *vector.GslVectorComplex) int32 {
	_arg_0 := complex_.GoComplexToGsl(tau)
	return int32(C.gsl_linalg_complex_householder_hv(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)), (*C.gsl_vector_complex)(unsafe.Pointer(v.Ptr())), (*C.gsl_vector_complex)(unsafe.Pointer(w.Ptr()))))
}

func HHSolve(a *matrix.GslMatrix, b *vector.GslVector, x *vector.GslVector) int32 {
	return int32(C.gsl_linalg_HH_solve((*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector)(unsafe.Pointer(b.Ptr())), (*C.gsl_vector)(unsafe.Pointer(x.Ptr()))))
}

func HHSvx(a *matrix.GslMatrix, x *vector.GslVector) int32 {
	return int32(C.gsl_linalg_HH_svx((*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector)(unsafe.Pointer(x.Ptr()))))
}

func SolveTridiag(diag *vector.GslVector, e *vector.GslVector, f *vector.GslVector, b *vector.GslVector, x *vector.GslVector) int32 {
	return int32(C.gsl_linalg_solve_tridiag((*C.gsl_vector)(unsafe.Pointer(diag.Ptr())), (*C.gsl_vector)(unsafe.Pointer(e.Ptr())), (*C.gsl_vector)(unsafe.Pointer(f.Ptr())), (*C.gsl_vector)(unsafe.Pointer(b.Ptr())), (*C.gsl_vector)(unsafe.Pointer(x.Ptr()))))
}

func SolveSymmTridiag(diag *vector.GslVector, e *vector.GslVector, b *vector.GslVector, x *vector.GslVector) int32 {
	return int32(C.gsl_linalg_solve_symm_tridiag((*C.gsl_vector)(unsafe.Pointer(diag.Ptr())), (*C.gsl_vector)(unsafe.Pointer(e.Ptr())), (*C.gsl_vector)(unsafe.Pointer(b.Ptr())), (*C.gsl_vector)(unsafe.Pointer(x.Ptr()))))
}

func SolveCycTridiag(diag *vector.GslVector, e *vector.GslVector, f *vector.GslVector, b *vector.GslVector, x *vector.GslVector) int32 {
	return int32(C.gsl_linalg_solve_cyc_tridiag((*C.gsl_vector)(unsafe.Pointer(diag.Ptr())), (*C.gsl_vector)(unsafe.Pointer(e.Ptr())), (*C.gsl_vector)(unsafe.Pointer(f.Ptr())), (*C.gsl_vector)(unsafe.Pointer(b.Ptr())), (*C.gsl_vector)(unsafe.Pointer(x.Ptr()))))
}

func SolveSymmCycTridiag(diag *vector.GslVector, e *vector.GslVector, b *vector.GslVector, x *vector.GslVector) int32 {
	return int32(C.gsl_linalg_solve_symm_cyc_tridiag((*C.gsl_vector)(unsafe.Pointer(diag.Ptr())), (*C.gsl_vector)(unsafe.Pointer(e.Ptr())), (*C.gsl_vector)(unsafe.Pointer(b.Ptr())), (*C.gsl_vector)(unsafe.Pointer(x.Ptr()))))
}

func BalanceMatrix(a *matrix.GslMatrix, d *vector.GslVector) int32 {
	return int32(C.gsl_linalg_balance_matrix((*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector)(unsafe.Pointer(d.Ptr()))))
}
