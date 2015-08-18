//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package blas

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_blas.h>
   #include <unistd.h>
*/
import "C"

import "github.com/dtromb/gogsl/vector"
import "unsafe"
import complex_ "github.com/dtromb/gogsl/complex"
import "reflect"
import "github.com/dtromb/gogsl/matrix"

func Sdsdot(alpha float32, x *vector.GslVectorFloat, y *vector.GslVectorFloat) (int32, float32) {
   var _outptr_3 C.float
   _result := int32(C.gsl_blas_sdsdot(C.float(alpha), (*C.gsl_vector_float)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector_float)(unsafe.Pointer(y.Ptr())), &_outptr_3))
   return _result, *(*float32)(unsafe.Pointer(&_outptr_3))
}

func Sdot(x *vector.GslVectorFloat, y *vector.GslVectorFloat) (int32, float32) {
   var _outptr_2 C.float
   _result := int32(C.gsl_blas_sdot((*C.gsl_vector_float)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector_float)(unsafe.Pointer(y.Ptr())), &_outptr_2))
   return _result, *(*float32)(unsafe.Pointer(&_outptr_2))
}

func Dsdot(x *vector.GslVectorFloat, y *vector.GslVectorFloat) (int32, float64) {
   var _outptr_2 C.double
   _result := int32(C.gsl_blas_dsdot((*C.gsl_vector_float)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector_float)(unsafe.Pointer(y.Ptr())), &_outptr_2))
   return _result, *(*float64)(unsafe.Pointer(&_outptr_2))
}

func Ddot(x *vector.GslVector, y *vector.GslVector) (int32, float64) {
   var _outptr_2 C.double
   _result := int32(C.gsl_blas_ddot((*C.gsl_vector)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector)(unsafe.Pointer(y.Ptr())), &_outptr_2))
   return _result, *(*float64)(unsafe.Pointer(&_outptr_2))
}

func Cdotu(x *vector.GslVectorComplexFloat, y *vector.GslVectorComplexFloat) (int32, complex64) {
   var _outptr_2 C.gsl_complex_float
   _result := int32(C.gsl_blas_cdotu((*C.gsl_vector_complex_float)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector_complex_float)(unsafe.Pointer(y.Ptr())), &_outptr_2))
   return _result, *(*complex64)(unsafe.Pointer(&_outptr_2))
}

func Zdotu(x *vector.GslVectorComplex, y *vector.GslVectorComplex) (int32, complex128) {
   var _outptr_2 C.gsl_complex
   _result := int32(C.gsl_blas_zdotu((*C.gsl_vector_complex)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector_complex)(unsafe.Pointer(y.Ptr())), &_outptr_2))
   return _result, *(*complex128)(unsafe.Pointer(&_outptr_2))
}

func Cdotc(x *vector.GslVectorComplexFloat, y *vector.GslVectorComplexFloat) (int32, complex64) {
   var _outptr_2 C.gsl_complex_float
   _result := int32(C.gsl_blas_cdotc((*C.gsl_vector_complex_float)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector_complex_float)(unsafe.Pointer(y.Ptr())), &_outptr_2))
   return _result, *(*complex64)(unsafe.Pointer(&_outptr_2))
}

func Zdotc(x *vector.GslVectorComplex, y *vector.GslVectorComplex) (int32, complex128) {
   var _outptr_2 C.gsl_complex
   _result := int32(C.gsl_blas_zdotc((*C.gsl_vector_complex)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector_complex)(unsafe.Pointer(y.Ptr())), &_outptr_2))
   return _result, *(*complex128)(unsafe.Pointer(&_outptr_2))
}

func Snrm2(x *vector.GslVectorFloat) float32 {
   return float32(C.gsl_blas_snrm2((*C.gsl_vector_float)(unsafe.Pointer(x.Ptr()))))
}

func Dnrm2(x *vector.GslVector) float64 {
   return float64(C.gsl_blas_dnrm2((*C.gsl_vector)(unsafe.Pointer(x.Ptr()))))
}

func Scnrm2(x *vector.GslVectorComplexFloat) float32 {
   return float32(C.gsl_blas_scnrm2((*C.gsl_vector_complex_float)(unsafe.Pointer(x.Ptr()))))
}

func Dznrm2(x *vector.GslVectorComplex) float64 {
   return float64(C.gsl_blas_dznrm2((*C.gsl_vector_complex)(unsafe.Pointer(x.Ptr()))))
}

func Sasum(x *vector.GslVectorFloat) float32 {
   return float32(C.gsl_blas_sasum((*C.gsl_vector_float)(unsafe.Pointer(x.Ptr()))))
}

func Dasum(x *vector.GslVector) float64 {
   return float64(C.gsl_blas_dasum((*C.gsl_vector)(unsafe.Pointer(x.Ptr()))))
}

func Scasum(x *vector.GslVectorComplexFloat) float32 {
   return float32(C.gsl_blas_scasum((*C.gsl_vector_complex_float)(unsafe.Pointer(x.Ptr()))))
}

func Dzasum(x *vector.GslVectorComplex) float64 {
   return float64(C.gsl_blas_dzasum((*C.gsl_vector_complex)(unsafe.Pointer(x.Ptr()))))
}

func Isamax(x *vector.GslVectorFloat) Index {
   return Index(C.gsl_blas_isamax((*C.gsl_vector_float)(unsafe.Pointer(x.Ptr()))))
}

func Idamax(x *vector.GslVector) Index {
   return Index(C.gsl_blas_idamax((*C.gsl_vector)(unsafe.Pointer(x.Ptr()))))
}

func Icamax(x *vector.GslVectorComplexFloat) Index {
   return Index(C.gsl_blas_icamax((*C.gsl_vector_complex_float)(unsafe.Pointer(x.Ptr()))))
}

func Izamax(x *vector.GslVectorComplex) Index {
   return Index(C.gsl_blas_izamax((*C.gsl_vector_complex)(unsafe.Pointer(x.Ptr()))))
}

func Sswap(x *vector.GslVectorFloat, y *vector.GslVectorFloat) int32 {
   return int32(C.gsl_blas_sswap((*C.gsl_vector_float)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector_float)(unsafe.Pointer(y.Ptr()))))
}

func Dswap(x *vector.GslVector, y *vector.GslVector) int32 {
   return int32(C.gsl_blas_dswap((*C.gsl_vector)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector)(unsafe.Pointer(y.Ptr()))))
}

func Cswap(x *vector.GslVectorComplexFloat, y *vector.GslVectorComplexFloat) int32 {
   return int32(C.gsl_blas_cswap((*C.gsl_vector_complex_float)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector_complex_float)(unsafe.Pointer(y.Ptr()))))
}

func Zswap(x *vector.GslVectorComplex, y *vector.GslVectorComplex) int32 {
   return int32(C.gsl_blas_zswap((*C.gsl_vector_complex)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector_complex)(unsafe.Pointer(y.Ptr()))))
}

func Scopy(x *vector.GslVectorFloat, y *vector.GslVectorFloat) int32 {
   return int32(C.gsl_blas_scopy((*C.gsl_vector_float)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector_float)(unsafe.Pointer(y.Ptr()))))
}

func Dcopy(x *vector.GslVector, y *vector.GslVector) int32 {
   return int32(C.gsl_blas_dcopy((*C.gsl_vector)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector)(unsafe.Pointer(y.Ptr()))))
}

func Ccopy(x *vector.GslVectorComplexFloat, y *vector.GslVectorComplexFloat) int32 {
   return int32(C.gsl_blas_ccopy((*C.gsl_vector_complex_float)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector_complex_float)(unsafe.Pointer(y.Ptr()))))
}

func Zcopy(x *vector.GslVectorComplex, y *vector.GslVectorComplex) int32 {
   return int32(C.gsl_blas_zcopy((*C.gsl_vector_complex)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector_complex)(unsafe.Pointer(y.Ptr()))))
}

func Saxpy(alpha float32, x *vector.GslVectorFloat, y *vector.GslVectorFloat) int32 {
   return int32(C.gsl_blas_saxpy(C.float(alpha), (*C.gsl_vector_float)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector_float)(unsafe.Pointer(y.Ptr()))))
}

func Daxpy(alpha float64, x *vector.GslVector, y *vector.GslVector) int32 {
   return int32(C.gsl_blas_daxpy(C.double(alpha), (*C.gsl_vector)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector)(unsafe.Pointer(y.Ptr()))))
}

func Caxpy(alpha complex64, x *vector.GslVectorComplexFloat, y *vector.GslVectorComplexFloat) int32 {
   _arg_0 := complex_.GoComplexFloatToGsl(alpha)
   return int32(C.gsl_blas_caxpy(*(*C.gsl_complex_float)(unsafe.Pointer(_arg_0)), (*C.gsl_vector_complex_float)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector_complex_float)(unsafe.Pointer(y.Ptr()))))
}

func Zaxpy(alpha complex128, x *vector.GslVectorComplex, y *vector.GslVectorComplex) int32 {
   _arg_0 := complex_.GoComplexToGsl(alpha)
   return int32(C.gsl_blas_zaxpy(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)), (*C.gsl_vector_complex)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector_complex)(unsafe.Pointer(y.Ptr()))))
}

func Sscal(alpha float32, x *vector.GslVectorFloat) {
   C.gsl_blas_sscal(C.float(alpha), (*C.gsl_vector_float)(unsafe.Pointer(x.Ptr())))
}

func Dscal(alpha float64, x *vector.GslVector) {
   C.gsl_blas_dscal(C.double(alpha), (*C.gsl_vector)(unsafe.Pointer(x.Ptr())))
}

func Cscal(alpha complex64, x *vector.GslVectorComplexFloat) {
   _arg_0 := complex_.GoComplexFloatToGsl(alpha)
   C.gsl_blas_cscal(*(*C.gsl_complex_float)(unsafe.Pointer(_arg_0)), (*C.gsl_vector_complex_float)(unsafe.Pointer(x.Ptr())))
}

func Zscal(alpha complex128, x *vector.GslVectorComplex) {
   _arg_0 := complex_.GoComplexToGsl(alpha)
   C.gsl_blas_zscal(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)), (*C.gsl_vector_complex)(unsafe.Pointer(x.Ptr())))
}

func Csscal(alpha float32, x *vector.GslVectorComplexFloat) {
   C.gsl_blas_csscal(C.float(alpha), (*C.gsl_vector_complex_float)(unsafe.Pointer(x.Ptr())))
}

func Zdscal(alpha float64, x *vector.GslVectorComplex) {
   C.gsl_blas_zdscal(C.double(alpha), (*C.gsl_vector_complex)(unsafe.Pointer(x.Ptr())))
}

func Srotg(a []float32, b []float32, c []float32, s []float32) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&a))
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&b))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&c))
   _slice_header_3 := (*reflect.SliceHeader)(unsafe.Pointer(&s))
   return int32(C.gsl_blas_srotg((*C.float)(unsafe.Pointer(_slice_header_0.Data)), (*C.float)(unsafe.Pointer(_slice_header_1.Data)), (*C.float)(unsafe.Pointer(_slice_header_2.Data)), (*C.float)(unsafe.Pointer(_slice_header_3.Data))))
}

func Drotg(a []float64, b []float64, c []float64, s []float64) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&a))
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&b))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&c))
   _slice_header_3 := (*reflect.SliceHeader)(unsafe.Pointer(&s))
   return int32(C.gsl_blas_drotg((*C.double)(unsafe.Pointer(_slice_header_0.Data)), (*C.double)(unsafe.Pointer(_slice_header_1.Data)), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), (*C.double)(unsafe.Pointer(_slice_header_3.Data))))
}

func Srot(x *vector.GslVectorFloat, y *vector.GslVectorFloat, c float32, s float32) int32 {
   return int32(C.gsl_blas_srot((*C.gsl_vector_float)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector_float)(unsafe.Pointer(y.Ptr())), C.float(c), C.float(s)))
}

func Drot(x *vector.GslVector, y *vector.GslVector, c float64, s float64) int32 {
   return int32(C.gsl_blas_drot((*C.gsl_vector)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector)(unsafe.Pointer(y.Ptr())), C.double(c), C.double(s)))
}

func Srotmg(d1 []float32, d2 []float32, b1 []float32, b2 float32, p []float32) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&d1))
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&d2))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&b1))
   _slice_header_4 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   return int32(C.gsl_blas_srotmg((*C.float)(unsafe.Pointer(_slice_header_0.Data)), (*C.float)(unsafe.Pointer(_slice_header_1.Data)), (*C.float)(unsafe.Pointer(_slice_header_2.Data)), C.float(b2), (*C.float)(unsafe.Pointer(_slice_header_4.Data))))
}

func Drotmg(d1 []float64, d2 []float64, b1 []float64, b2 float64, p []float64) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&d1))
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&d2))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&b1))
   _slice_header_4 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   return int32(C.gsl_blas_drotmg((*C.double)(unsafe.Pointer(_slice_header_0.Data)), (*C.double)(unsafe.Pointer(_slice_header_1.Data)), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.double(b2), (*C.double)(unsafe.Pointer(_slice_header_4.Data))))
}

func Srotm(x *vector.GslVectorFloat, y *vector.GslVectorFloat, p []float32) int32 {
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   return int32(C.gsl_blas_srotm((*C.gsl_vector_float)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector_float)(unsafe.Pointer(y.Ptr())), (*C.float)(unsafe.Pointer(_slice_header_2.Data))))
}

func Drotm(x *vector.GslVector, y *vector.GslVector, p []float64) int32 {
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   return int32(C.gsl_blas_drotm((*C.gsl_vector)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector)(unsafe.Pointer(y.Ptr())), (*C.double)(unsafe.Pointer(_slice_header_2.Data))))
}

func Sgemv(transA Transpose, alpha float32, a *matrix.GslMatrixFloat, x *vector.GslVectorFloat, beta float32, y *vector.GslVectorFloat) int32 {
   return int32(C.gsl_blas_sgemv(C.CBLAS_TRANSPOSE_t(transA), C.float(alpha), (*C.gsl_matrix_float)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_float)(unsafe.Pointer(x.Ptr())), C.float(beta), (*C.gsl_vector_float)(unsafe.Pointer(y.Ptr()))))
}

func Dgemv(transA Transpose, alpha float64, a *matrix.GslMatrix, x *vector.GslVector, beta float64, y *vector.GslVector) int32 {
   return int32(C.gsl_blas_dgemv(C.CBLAS_TRANSPOSE_t(transA), C.double(alpha), (*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector)(unsafe.Pointer(x.Ptr())), C.double(beta), (*C.gsl_vector)(unsafe.Pointer(y.Ptr()))))
}

func Cgemv(transA Transpose, alpha complex64, a *matrix.GslMatrixComplexFloat, x *vector.GslVectorComplexFloat, beta complex64, y *vector.GslVectorComplexFloat) int32 {
   _arg_1 := complex_.GoComplexFloatToGsl(alpha)
   _arg_4 := complex_.GoComplexFloatToGsl(beta)
   return int32(C.gsl_blas_cgemv(C.CBLAS_TRANSPOSE_t(transA), *(*C.gsl_complex_float)(unsafe.Pointer(_arg_1)), (*C.gsl_matrix_complex_float)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_complex_float)(unsafe.Pointer(x.Ptr())), *(*C.gsl_complex_float)(unsafe.Pointer(_arg_4)), (*C.gsl_vector_complex_float)(unsafe.Pointer(y.Ptr()))))
}

func Zgemv(transA Transpose, alpha complex128, a *matrix.GslMatrixComplex, x *vector.GslVectorComplex, beta complex128, y *vector.GslVectorComplex) int32 {
   _arg_1 := complex_.GoComplexToGsl(alpha)
   _arg_4 := complex_.GoComplexToGsl(beta)
   return int32(C.gsl_blas_zgemv(C.CBLAS_TRANSPOSE_t(transA), *(*C.gsl_complex)(unsafe.Pointer(_arg_1)), (*C.gsl_matrix_complex)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_complex)(unsafe.Pointer(x.Ptr())), *(*C.gsl_complex)(unsafe.Pointer(_arg_4)), (*C.gsl_vector_complex)(unsafe.Pointer(y.Ptr()))))
}

func Strmv(uplo Uplo, transA Transpose, diag Diag, a *matrix.GslMatrixFloat, x *vector.GslVectorFloat) int32 {
   return int32(C.gsl_blas_strmv(C.CBLAS_UPLO_t(uplo), C.CBLAS_TRANSPOSE_t(transA), C.CBLAS_DIAG_t(diag), (*C.gsl_matrix_float)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_float)(unsafe.Pointer(x.Ptr()))))
}

func Dtrmv(uplo Uplo, transA Transpose, diag Diag, a *matrix.GslMatrix, x *vector.GslVector) int32 {
   return int32(C.gsl_blas_dtrmv(C.CBLAS_UPLO_t(uplo), C.CBLAS_TRANSPOSE_t(transA), C.CBLAS_DIAG_t(diag), (*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector)(unsafe.Pointer(x.Ptr()))))
}

func Ctrmv(uplo Uplo, transA Transpose, diag Diag, a *matrix.GslMatrixComplexFloat, x *vector.GslVectorComplexFloat) int32 {
   return int32(C.gsl_blas_ctrmv(C.CBLAS_UPLO_t(uplo), C.CBLAS_TRANSPOSE_t(transA), C.CBLAS_DIAG_t(diag), (*C.gsl_matrix_complex_float)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_complex_float)(unsafe.Pointer(x.Ptr()))))
}

func Ztrmv(uplo Uplo, transA Transpose, diag Diag, a *matrix.GslMatrixComplex, x *vector.GslVectorComplex) int32 {
   return int32(C.gsl_blas_ztrmv(C.CBLAS_UPLO_t(uplo), C.CBLAS_TRANSPOSE_t(transA), C.CBLAS_DIAG_t(diag), (*C.gsl_matrix_complex)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_complex)(unsafe.Pointer(x.Ptr()))))
}

func Strsv(uplo Uplo, transA Transpose, diag Diag, a *matrix.GslMatrixFloat, x *vector.GslVectorFloat) int32 {
   return int32(C.gsl_blas_strsv(C.CBLAS_UPLO_t(uplo), C.CBLAS_TRANSPOSE_t(transA), C.CBLAS_DIAG_t(diag), (*C.gsl_matrix_float)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_float)(unsafe.Pointer(x.Ptr()))))
}

func Dtrsv(uplo Uplo, transA Transpose, diag Diag, a *matrix.GslMatrix, x *vector.GslVector) int32 {
   return int32(C.gsl_blas_dtrsv(C.CBLAS_UPLO_t(uplo), C.CBLAS_TRANSPOSE_t(transA), C.CBLAS_DIAG_t(diag), (*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector)(unsafe.Pointer(x.Ptr()))))
}

func Ctrsv(uplo Uplo, transA Transpose, diag Diag, a *matrix.GslMatrixComplexFloat, x *vector.GslVectorComplexFloat) int32 {
   return int32(C.gsl_blas_ctrsv(C.CBLAS_UPLO_t(uplo), C.CBLAS_TRANSPOSE_t(transA), C.CBLAS_DIAG_t(diag), (*C.gsl_matrix_complex_float)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_complex_float)(unsafe.Pointer(x.Ptr()))))
}

func Ztrsv(uplo Uplo, transA Transpose, diag Diag, a *matrix.GslMatrixComplex, x *vector.GslVectorComplex) int32 {
   return int32(C.gsl_blas_ztrsv(C.CBLAS_UPLO_t(uplo), C.CBLAS_TRANSPOSE_t(transA), C.CBLAS_DIAG_t(diag), (*C.gsl_matrix_complex)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_complex)(unsafe.Pointer(x.Ptr()))))
}

func Ssymv(uplo Uplo, alpha float32, a *matrix.GslMatrixFloat, x *vector.GslVectorFloat, beta float32, y *vector.GslVectorFloat) int32 {
   return int32(C.gsl_blas_ssymv(C.CBLAS_UPLO_t(uplo), C.float(alpha), (*C.gsl_matrix_float)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_float)(unsafe.Pointer(x.Ptr())), C.float(beta), (*C.gsl_vector_float)(unsafe.Pointer(y.Ptr()))))
}

func Dsymv(uplo Uplo, alpha float64, a *matrix.GslMatrix, x *vector.GslVector, beta float64, y *vector.GslVector) int32 {
   return int32(C.gsl_blas_dsymv(C.CBLAS_UPLO_t(uplo), C.double(alpha), (*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector)(unsafe.Pointer(x.Ptr())), C.double(beta), (*C.gsl_vector)(unsafe.Pointer(y.Ptr()))))
}

func Chemv(uplo Uplo, alpha complex64, a *matrix.GslMatrixComplexFloat, x *vector.GslVectorComplexFloat, beta complex64, y *vector.GslVectorComplexFloat) int32 {
   _arg_1 := complex_.GoComplexFloatToGsl(alpha)
   _arg_4 := complex_.GoComplexFloatToGsl(beta)
   return int32(C.gsl_blas_chemv(C.CBLAS_UPLO_t(uplo), *(*C.gsl_complex_float)(unsafe.Pointer(_arg_1)), (*C.gsl_matrix_complex_float)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_complex_float)(unsafe.Pointer(x.Ptr())), *(*C.gsl_complex_float)(unsafe.Pointer(_arg_4)), (*C.gsl_vector_complex_float)(unsafe.Pointer(y.Ptr()))))
}

func Zhemv(uplo Uplo, alpha complex128, a *matrix.GslMatrixComplex, x *vector.GslVectorComplex, beta complex128, y *vector.GslVectorComplex) int32 {
   _arg_1 := complex_.GoComplexToGsl(alpha)
   _arg_4 := complex_.GoComplexToGsl(beta)
   return int32(C.gsl_blas_zhemv(C.CBLAS_UPLO_t(uplo), *(*C.gsl_complex)(unsafe.Pointer(_arg_1)), (*C.gsl_matrix_complex)(unsafe.Pointer(a.Ptr())), (*C.gsl_vector_complex)(unsafe.Pointer(x.Ptr())), *(*C.gsl_complex)(unsafe.Pointer(_arg_4)), (*C.gsl_vector_complex)(unsafe.Pointer(y.Ptr()))))
}

func Sger(alpha float32, x *vector.GslVectorFloat, y *vector.GslVectorFloat, a *matrix.GslMatrixFloat) int32 {
   return int32(C.gsl_blas_sger(C.float(alpha), (*C.gsl_vector_float)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector_float)(unsafe.Pointer(y.Ptr())), (*C.gsl_matrix_float)(unsafe.Pointer(a.Ptr()))))
}

func Dger(alpha float64, x *vector.GslVector, y *vector.GslVector, a *matrix.GslMatrix) int32 {
   return int32(C.gsl_blas_dger(C.double(alpha), (*C.gsl_vector)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector)(unsafe.Pointer(y.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(a.Ptr()))))
}

func Cgeru(alpha complex64, x *vector.GslVectorComplexFloat, y *vector.GslVectorComplexFloat, a *matrix.GslMatrixComplexFloat) int32 {
   _arg_0 := complex_.GoComplexFloatToGsl(alpha)
   return int32(C.gsl_blas_cgeru(*(*C.gsl_complex_float)(unsafe.Pointer(_arg_0)), (*C.gsl_vector_complex_float)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector_complex_float)(unsafe.Pointer(y.Ptr())), (*C.gsl_matrix_complex_float)(unsafe.Pointer(a.Ptr()))))
}

func Zgeru(alpha complex128, x *vector.GslVectorComplex, y *vector.GslVectorComplex, a *matrix.GslMatrixComplex) int32 {
   _arg_0 := complex_.GoComplexToGsl(alpha)
   return int32(C.gsl_blas_zgeru(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)), (*C.gsl_vector_complex)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector_complex)(unsafe.Pointer(y.Ptr())), (*C.gsl_matrix_complex)(unsafe.Pointer(a.Ptr()))))
}

func Cgerc(alpha complex64, x *vector.GslVectorComplexFloat, y *vector.GslVectorComplexFloat, a *matrix.GslMatrixComplexFloat) int32 {
   _arg_0 := complex_.GoComplexFloatToGsl(alpha)
   return int32(C.gsl_blas_cgerc(*(*C.gsl_complex_float)(unsafe.Pointer(_arg_0)), (*C.gsl_vector_complex_float)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector_complex_float)(unsafe.Pointer(y.Ptr())), (*C.gsl_matrix_complex_float)(unsafe.Pointer(a.Ptr()))))
}

func Zgerc(alpha complex128, x *vector.GslVectorComplex, y *vector.GslVectorComplex, a *matrix.GslMatrixComplex) int32 {
   _arg_0 := complex_.GoComplexToGsl(alpha)
   return int32(C.gsl_blas_zgerc(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)), (*C.gsl_vector_complex)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector_complex)(unsafe.Pointer(y.Ptr())), (*C.gsl_matrix_complex)(unsafe.Pointer(a.Ptr()))))
}

func Ssyr(uplo Uplo, alpha float32, x *vector.GslVectorFloat, a *matrix.GslMatrixFloat) int32 {
   return int32(C.gsl_blas_ssyr(C.CBLAS_UPLO_t(uplo), C.float(alpha), (*C.gsl_vector_float)(unsafe.Pointer(x.Ptr())), (*C.gsl_matrix_float)(unsafe.Pointer(a.Ptr()))))
}

func Dsyr(uplo Uplo, alpha float64, x *vector.GslVector, a *matrix.GslMatrix) int32 {
   return int32(C.gsl_blas_dsyr(C.CBLAS_UPLO_t(uplo), C.double(alpha), (*C.gsl_vector)(unsafe.Pointer(x.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(a.Ptr()))))
}

func Cher(uplo Uplo, alpha float32, x *vector.GslVectorComplexFloat, a *matrix.GslMatrixComplexFloat) int32 {
   return int32(C.gsl_blas_cher(C.CBLAS_UPLO_t(uplo), C.float(alpha), (*C.gsl_vector_complex_float)(unsafe.Pointer(x.Ptr())), (*C.gsl_matrix_complex_float)(unsafe.Pointer(a.Ptr()))))
}

func Zher(uplo Uplo, alpha float64, x *vector.GslVectorComplex, a *matrix.GslMatrixComplex) int32 {
   return int32(C.gsl_blas_zher(C.CBLAS_UPLO_t(uplo), C.double(alpha), (*C.gsl_vector_complex)(unsafe.Pointer(x.Ptr())), (*C.gsl_matrix_complex)(unsafe.Pointer(a.Ptr()))))
}

func Ssyr2(uplo Uplo, alpha float32, x *vector.GslVectorFloat, y *vector.GslVectorFloat, a *matrix.GslMatrixFloat) int32 {
   return int32(C.gsl_blas_ssyr2(C.CBLAS_UPLO_t(uplo), C.float(alpha), (*C.gsl_vector_float)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector_float)(unsafe.Pointer(y.Ptr())), (*C.gsl_matrix_float)(unsafe.Pointer(a.Ptr()))))
}

func Dsyr2(uplo Uplo, alpha float64, x *vector.GslVector, y *vector.GslVector, a *matrix.GslMatrix) int32 {
   return int32(C.gsl_blas_dsyr2(C.CBLAS_UPLO_t(uplo), C.double(alpha), (*C.gsl_vector)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector)(unsafe.Pointer(y.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(a.Ptr()))))
}

func Cher2(uplo Uplo, alpha complex64, x *vector.GslVectorComplexFloat, y *vector.GslVectorComplexFloat, a *matrix.GslMatrixComplexFloat) int32 {
   _arg_1 := complex_.GoComplexFloatToGsl(alpha)
   return int32(C.gsl_blas_cher2(C.CBLAS_UPLO_t(uplo), *(*C.gsl_complex_float)(unsafe.Pointer(_arg_1)), (*C.gsl_vector_complex_float)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector_complex_float)(unsafe.Pointer(y.Ptr())), (*C.gsl_matrix_complex_float)(unsafe.Pointer(a.Ptr()))))
}

func Zher2(uplo Uplo, alpha complex128, x *vector.GslVectorComplex, y *vector.GslVectorComplex, a *matrix.GslMatrixComplex) int32 {
   _arg_1 := complex_.GoComplexToGsl(alpha)
   return int32(C.gsl_blas_zher2(C.CBLAS_UPLO_t(uplo), *(*C.gsl_complex)(unsafe.Pointer(_arg_1)), (*C.gsl_vector_complex)(unsafe.Pointer(x.Ptr())), (*C.gsl_vector_complex)(unsafe.Pointer(y.Ptr())), (*C.gsl_matrix_complex)(unsafe.Pointer(a.Ptr()))))
}

func Sgemm(transA Transpose, transB Transpose, alpha float32, a *matrix.GslMatrixFloat, b *matrix.GslMatrixFloat, beta float32, c *matrix.GslMatrixFloat) int32 {
   return int32(C.gsl_blas_sgemm(C.CBLAS_TRANSPOSE_t(transA), C.CBLAS_TRANSPOSE_t(transB), C.float(alpha), (*C.gsl_matrix_float)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_float)(unsafe.Pointer(b.Ptr())), C.float(beta), (*C.gsl_matrix_float)(unsafe.Pointer(c.Ptr()))))
}

func Dgemm(transA Transpose, transB Transpose, alpha float64, a *matrix.GslMatrix, b *matrix.GslMatrix, beta float64, c *matrix.GslMatrix) int32 {
   return int32(C.gsl_blas_dgemm(C.CBLAS_TRANSPOSE_t(transA), C.CBLAS_TRANSPOSE_t(transB), C.double(alpha), (*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(b.Ptr())), C.double(beta), (*C.gsl_matrix)(unsafe.Pointer(c.Ptr()))))
}

func Cgemm(transA Transpose, transB Transpose, alpha complex64, a *matrix.GslMatrixComplexFloat, b *matrix.GslMatrixComplexFloat, beta complex64, c *matrix.GslMatrixComplexFloat) int32 {
   _arg_2 := complex_.GoComplexFloatToGsl(alpha)
   _arg_5 := complex_.GoComplexFloatToGsl(beta)
   return int32(C.gsl_blas_cgemm(C.CBLAS_TRANSPOSE_t(transA), C.CBLAS_TRANSPOSE_t(transB), *(*C.gsl_complex_float)(unsafe.Pointer(_arg_2)), (*C.gsl_matrix_complex_float)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_complex_float)(unsafe.Pointer(b.Ptr())), *(*C.gsl_complex_float)(unsafe.Pointer(_arg_5)), (*C.gsl_matrix_complex_float)(unsafe.Pointer(c.Ptr()))))
}

func Zgemm(transA Transpose, transB Transpose, alpha complex128, a *matrix.GslMatrixComplex, b *matrix.GslMatrixComplex, beta complex128, c *matrix.GslMatrixComplex) int32 {
   _arg_2 := complex_.GoComplexToGsl(alpha)
   _arg_5 := complex_.GoComplexToGsl(beta)
   return int32(C.gsl_blas_zgemm(C.CBLAS_TRANSPOSE_t(transA), C.CBLAS_TRANSPOSE_t(transB), *(*C.gsl_complex)(unsafe.Pointer(_arg_2)), (*C.gsl_matrix_complex)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_complex)(unsafe.Pointer(b.Ptr())), *(*C.gsl_complex)(unsafe.Pointer(_arg_5)), (*C.gsl_matrix_complex)(unsafe.Pointer(c.Ptr()))))
}

func Ssymm(side Side, uplo Uplo, alpha float32, a *matrix.GslMatrixFloat, b *matrix.GslMatrixFloat, beta float32, c *matrix.GslMatrixFloat) int32 {
   return int32(C.gsl_blas_ssymm(C.CBLAS_SIDE_t(side), C.CBLAS_UPLO_t(uplo), C.float(alpha), (*C.gsl_matrix_float)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_float)(unsafe.Pointer(b.Ptr())), C.float(beta), (*C.gsl_matrix_float)(unsafe.Pointer(c.Ptr()))))
}

func Dsymm(side Side, uplo Uplo, alpha float64, a *matrix.GslMatrix, b *matrix.GslMatrix, beta float64, c *matrix.GslMatrix) int32 {
   return int32(C.gsl_blas_dsymm(C.CBLAS_SIDE_t(side), C.CBLAS_UPLO_t(uplo), C.double(alpha), (*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(b.Ptr())), C.double(beta), (*C.gsl_matrix)(unsafe.Pointer(c.Ptr()))))
}

func Csymm(side Side, uplo Uplo, alpha complex64, a *matrix.GslMatrixComplexFloat, b *matrix.GslMatrixComplexFloat, beta complex64, c *matrix.GslMatrixComplexFloat) int32 {
   _arg_2 := complex_.GoComplexFloatToGsl(alpha)
   _arg_5 := complex_.GoComplexFloatToGsl(beta)
   return int32(C.gsl_blas_csymm(C.CBLAS_SIDE_t(side), C.CBLAS_UPLO_t(uplo), *(*C.gsl_complex_float)(unsafe.Pointer(_arg_2)), (*C.gsl_matrix_complex_float)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_complex_float)(unsafe.Pointer(b.Ptr())), *(*C.gsl_complex_float)(unsafe.Pointer(_arg_5)), (*C.gsl_matrix_complex_float)(unsafe.Pointer(c.Ptr()))))
}

func Zsymm(side Side, uplo Uplo, alpha complex128, a *matrix.GslMatrixComplex, b *matrix.GslMatrixComplex, beta complex128, c *matrix.GslMatrixComplex) int32 {
   _arg_2 := complex_.GoComplexToGsl(alpha)
   _arg_5 := complex_.GoComplexToGsl(beta)
   return int32(C.gsl_blas_zsymm(C.CBLAS_SIDE_t(side), C.CBLAS_UPLO_t(uplo), *(*C.gsl_complex)(unsafe.Pointer(_arg_2)), (*C.gsl_matrix_complex)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_complex)(unsafe.Pointer(b.Ptr())), *(*C.gsl_complex)(unsafe.Pointer(_arg_5)), (*C.gsl_matrix_complex)(unsafe.Pointer(c.Ptr()))))
}

func Chemm(side Side, uplo Uplo, alpha complex64, a *matrix.GslMatrixComplexFloat, b *matrix.GslMatrixComplexFloat, beta complex64, c *matrix.GslMatrixComplexFloat) int32 {
   _arg_2 := complex_.GoComplexFloatToGsl(alpha)
   _arg_5 := complex_.GoComplexFloatToGsl(beta)
   return int32(C.gsl_blas_chemm(C.CBLAS_SIDE_t(side), C.CBLAS_UPLO_t(uplo), *(*C.gsl_complex_float)(unsafe.Pointer(_arg_2)), (*C.gsl_matrix_complex_float)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_complex_float)(unsafe.Pointer(b.Ptr())), *(*C.gsl_complex_float)(unsafe.Pointer(_arg_5)), (*C.gsl_matrix_complex_float)(unsafe.Pointer(c.Ptr()))))
}

func Zhemm(side Side, uplo Uplo, alpha complex128, a *matrix.GslMatrixComplex, b *matrix.GslMatrixComplex, beta complex128, c *matrix.GslMatrixComplex) int32 {
   _arg_2 := complex_.GoComplexToGsl(alpha)
   _arg_5 := complex_.GoComplexToGsl(beta)
   return int32(C.gsl_blas_zhemm(C.CBLAS_SIDE_t(side), C.CBLAS_UPLO_t(uplo), *(*C.gsl_complex)(unsafe.Pointer(_arg_2)), (*C.gsl_matrix_complex)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_complex)(unsafe.Pointer(b.Ptr())), *(*C.gsl_complex)(unsafe.Pointer(_arg_5)), (*C.gsl_matrix_complex)(unsafe.Pointer(c.Ptr()))))
}

func Strmm(side Side, uplo Uplo, transA Transpose, diag Diag, alpha float32, a *matrix.GslMatrixFloat, b *matrix.GslMatrixFloat) int32 {
   return int32(C.gsl_blas_strmm(C.CBLAS_SIDE_t(side), C.CBLAS_UPLO_t(uplo), C.CBLAS_TRANSPOSE_t(transA), C.CBLAS_DIAG_t(diag), C.float(alpha), (*C.gsl_matrix_float)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_float)(unsafe.Pointer(b.Ptr()))))
}

func Dtrmm(side Side, uplo Uplo, transA Transpose, diag Diag, alpha float64, a *matrix.GslMatrix, b *matrix.GslMatrix) int32 {
   return int32(C.gsl_blas_dtrmm(C.CBLAS_SIDE_t(side), C.CBLAS_UPLO_t(uplo), C.CBLAS_TRANSPOSE_t(transA), C.CBLAS_DIAG_t(diag), C.double(alpha), (*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(b.Ptr()))))
}

func Ctrmm(side Side, uplo Uplo, transA Transpose, diag Diag, alpha complex64, a *matrix.GslMatrixComplexFloat, b *matrix.GslMatrixComplexFloat) int32 {
   _arg_4 := complex_.GoComplexFloatToGsl(alpha)
   return int32(C.gsl_blas_ctrmm(C.CBLAS_SIDE_t(side), C.CBLAS_UPLO_t(uplo), C.CBLAS_TRANSPOSE_t(transA), C.CBLAS_DIAG_t(diag), *(*C.gsl_complex_float)(unsafe.Pointer(_arg_4)), (*C.gsl_matrix_complex_float)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_complex_float)(unsafe.Pointer(b.Ptr()))))
}

func Ztrmm(side Side, uplo Uplo, transA Transpose, diag Diag, alpha complex128, a *matrix.GslMatrixComplex, b *matrix.GslMatrixComplex) int32 {
   _arg_4 := complex_.GoComplexToGsl(alpha)
   return int32(C.gsl_blas_ztrmm(C.CBLAS_SIDE_t(side), C.CBLAS_UPLO_t(uplo), C.CBLAS_TRANSPOSE_t(transA), C.CBLAS_DIAG_t(diag), *(*C.gsl_complex)(unsafe.Pointer(_arg_4)), (*C.gsl_matrix_complex)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_complex)(unsafe.Pointer(b.Ptr()))))
}

func Strsm(side Side, uplo Uplo, transA Transpose, diag Diag, alpha float32, a *matrix.GslMatrixFloat, b *matrix.GslMatrixFloat) int32 {
   return int32(C.gsl_blas_strsm(C.CBLAS_SIDE_t(side), C.CBLAS_UPLO_t(uplo), C.CBLAS_TRANSPOSE_t(transA), C.CBLAS_DIAG_t(diag), C.float(alpha), (*C.gsl_matrix_float)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_float)(unsafe.Pointer(b.Ptr()))))
}

func Dtrsm(side Side, uplo Uplo, transA Transpose, diag Diag, alpha float64, a *matrix.GslMatrix, b *matrix.GslMatrix) int32 {
   return int32(C.gsl_blas_dtrsm(C.CBLAS_SIDE_t(side), C.CBLAS_UPLO_t(uplo), C.CBLAS_TRANSPOSE_t(transA), C.CBLAS_DIAG_t(diag), C.double(alpha), (*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(b.Ptr()))))
}

func Ctrsm(side Side, uplo Uplo, transA Transpose, diag Diag, alpha complex64, a *matrix.GslMatrixComplexFloat, b *matrix.GslMatrixComplexFloat) int32 {
   _arg_4 := complex_.GoComplexFloatToGsl(alpha)
   return int32(C.gsl_blas_ctrsm(C.CBLAS_SIDE_t(side), C.CBLAS_UPLO_t(uplo), C.CBLAS_TRANSPOSE_t(transA), C.CBLAS_DIAG_t(diag), *(*C.gsl_complex_float)(unsafe.Pointer(_arg_4)), (*C.gsl_matrix_complex_float)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_complex_float)(unsafe.Pointer(b.Ptr()))))
}

func Ztrsm(side Side, uplo Uplo, transA Transpose, diag Diag, alpha complex128, a *matrix.GslMatrixComplex, b *matrix.GslMatrixComplex) int32 {
   _arg_4 := complex_.GoComplexToGsl(alpha)
   return int32(C.gsl_blas_ztrsm(C.CBLAS_SIDE_t(side), C.CBLAS_UPLO_t(uplo), C.CBLAS_TRANSPOSE_t(transA), C.CBLAS_DIAG_t(diag), *(*C.gsl_complex)(unsafe.Pointer(_arg_4)), (*C.gsl_matrix_complex)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_complex)(unsafe.Pointer(b.Ptr()))))
}

func Ssyrk(uplo Uplo, trans Transpose, alpha float32, a *matrix.GslMatrixFloat, beta float32, c *matrix.GslMatrixFloat) int32 {
   return int32(C.gsl_blas_ssyrk(C.CBLAS_UPLO_t(uplo), C.CBLAS_TRANSPOSE_t(trans), C.float(alpha), (*C.gsl_matrix_float)(unsafe.Pointer(a.Ptr())), C.float(beta), (*C.gsl_matrix_float)(unsafe.Pointer(c.Ptr()))))
}

func Dsyrk(uplo Uplo, trans Transpose, alpha float64, a *matrix.GslMatrix, beta float64, c *matrix.GslMatrix) int32 {
   return int32(C.gsl_blas_dsyrk(C.CBLAS_UPLO_t(uplo), C.CBLAS_TRANSPOSE_t(trans), C.double(alpha), (*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), C.double(beta), (*C.gsl_matrix)(unsafe.Pointer(c.Ptr()))))
}

func Csyrk(uplo Uplo, trans Transpose, alpha complex64, a *matrix.GslMatrixComplexFloat, beta complex64, c *matrix.GslMatrixComplexFloat) int32 {
   _arg_2 := complex_.GoComplexFloatToGsl(alpha)
   _arg_4 := complex_.GoComplexFloatToGsl(beta)
   return int32(C.gsl_blas_csyrk(C.CBLAS_UPLO_t(uplo), C.CBLAS_TRANSPOSE_t(trans), *(*C.gsl_complex_float)(unsafe.Pointer(_arg_2)), (*C.gsl_matrix_complex_float)(unsafe.Pointer(a.Ptr())), *(*C.gsl_complex_float)(unsafe.Pointer(_arg_4)), (*C.gsl_matrix_complex_float)(unsafe.Pointer(c.Ptr()))))
}

func Zsyrk(uplo Uplo, trans Transpose, alpha complex128, a *matrix.GslMatrixComplex, beta complex128, c *matrix.GslMatrixComplex) int32 {
   _arg_2 := complex_.GoComplexToGsl(alpha)
   _arg_4 := complex_.GoComplexToGsl(beta)
   return int32(C.gsl_blas_zsyrk(C.CBLAS_UPLO_t(uplo), C.CBLAS_TRANSPOSE_t(trans), *(*C.gsl_complex)(unsafe.Pointer(_arg_2)), (*C.gsl_matrix_complex)(unsafe.Pointer(a.Ptr())), *(*C.gsl_complex)(unsafe.Pointer(_arg_4)), (*C.gsl_matrix_complex)(unsafe.Pointer(c.Ptr()))))
}

func Cherk(uplo Uplo, trans Transpose, alpha float32, a *matrix.GslMatrixComplexFloat, beta float32, c *matrix.GslMatrixComplexFloat) int32 {
   return int32(C.gsl_blas_cherk(C.CBLAS_UPLO_t(uplo), C.CBLAS_TRANSPOSE_t(trans), C.float(alpha), (*C.gsl_matrix_complex_float)(unsafe.Pointer(a.Ptr())), C.float(beta), (*C.gsl_matrix_complex_float)(unsafe.Pointer(c.Ptr()))))
}

func Zherk(uplo Uplo, trans Transpose, alpha float64, a *matrix.GslMatrixComplex, beta float64, c *matrix.GslMatrixComplex) int32 {
   return int32(C.gsl_blas_zherk(C.CBLAS_UPLO_t(uplo), C.CBLAS_TRANSPOSE_t(trans), C.double(alpha), (*C.gsl_matrix_complex)(unsafe.Pointer(a.Ptr())), C.double(beta), (*C.gsl_matrix_complex)(unsafe.Pointer(c.Ptr()))))
}

func Ssyr2k(uplo Uplo, trans Transpose, alpha float32, a *matrix.GslMatrixFloat, b *matrix.GslMatrixFloat, beta float32, c *matrix.GslMatrixFloat) int32 {
   return int32(C.gsl_blas_ssyr2k(C.CBLAS_UPLO_t(uplo), C.CBLAS_TRANSPOSE_t(trans), C.float(alpha), (*C.gsl_matrix_float)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_float)(unsafe.Pointer(b.Ptr())), C.float(beta), (*C.gsl_matrix_float)(unsafe.Pointer(c.Ptr()))))
}

func Dsyr2k(uplo Uplo, trans Transpose, alpha float64, a *matrix.GslMatrix, b *matrix.GslMatrix, beta float64, c *matrix.GslMatrix) int32 {
   return int32(C.gsl_blas_dsyr2k(C.CBLAS_UPLO_t(uplo), C.CBLAS_TRANSPOSE_t(trans), C.double(alpha), (*C.gsl_matrix)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix)(unsafe.Pointer(b.Ptr())), C.double(beta), (*C.gsl_matrix)(unsafe.Pointer(c.Ptr()))))
}

func Csyr2k(uplo Uplo, trans Transpose, alpha complex64, a *matrix.GslMatrixComplexFloat, b *matrix.GslMatrixComplexFloat, beta complex64, c *matrix.GslMatrixComplexFloat) int32 {
   _arg_2 := complex_.GoComplexFloatToGsl(alpha)
   _arg_5 := complex_.GoComplexFloatToGsl(beta)
   return int32(C.gsl_blas_csyr2k(C.CBLAS_UPLO_t(uplo), C.CBLAS_TRANSPOSE_t(trans), *(*C.gsl_complex_float)(unsafe.Pointer(_arg_2)), (*C.gsl_matrix_complex_float)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_complex_float)(unsafe.Pointer(b.Ptr())), *(*C.gsl_complex_float)(unsafe.Pointer(_arg_5)), (*C.gsl_matrix_complex_float)(unsafe.Pointer(c.Ptr()))))
}

func Zsyr2k(uplo Uplo, trans Transpose, alpha complex128, a *matrix.GslMatrixComplex, b *matrix.GslMatrixComplex, beta complex128, c *matrix.GslMatrixComplex) int32 {
   _arg_2 := complex_.GoComplexToGsl(alpha)
   _arg_5 := complex_.GoComplexToGsl(beta)
   return int32(C.gsl_blas_zsyr2k(C.CBLAS_UPLO_t(uplo), C.CBLAS_TRANSPOSE_t(trans), *(*C.gsl_complex)(unsafe.Pointer(_arg_2)), (*C.gsl_matrix_complex)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_complex)(unsafe.Pointer(b.Ptr())), *(*C.gsl_complex)(unsafe.Pointer(_arg_5)), (*C.gsl_matrix_complex)(unsafe.Pointer(c.Ptr()))))
}

func Cher2k(uplo Uplo, trans Transpose, alpha complex64, a *matrix.GslMatrixComplexFloat, b *matrix.GslMatrixComplexFloat, beta float32, c *matrix.GslMatrixComplexFloat) int32 {
   _arg_2 := complex_.GoComplexFloatToGsl(alpha)
   return int32(C.gsl_blas_cher2k(C.CBLAS_UPLO_t(uplo), C.CBLAS_TRANSPOSE_t(trans), *(*C.gsl_complex_float)(unsafe.Pointer(_arg_2)), (*C.gsl_matrix_complex_float)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_complex_float)(unsafe.Pointer(b.Ptr())), C.float(beta), (*C.gsl_matrix_complex_float)(unsafe.Pointer(c.Ptr()))))
}

func Zher2k(uplo Uplo, trans Transpose, alpha complex128, a *matrix.GslMatrixComplex, b *matrix.GslMatrixComplex, beta float64, c *matrix.GslMatrixComplex) int32 {
   _arg_2 := complex_.GoComplexToGsl(alpha)
   return int32(C.gsl_blas_zher2k(C.CBLAS_UPLO_t(uplo), C.CBLAS_TRANSPOSE_t(trans), *(*C.gsl_complex)(unsafe.Pointer(_arg_2)), (*C.gsl_matrix_complex)(unsafe.Pointer(a.Ptr())), (*C.gsl_matrix_complex)(unsafe.Pointer(b.Ptr())), C.double(beta), (*C.gsl_matrix_complex)(unsafe.Pointer(c.Ptr()))))
}

