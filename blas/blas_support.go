package blas

/*
	#include <gsl/gsl_cblas.h>
*/
import "C"

type Index C.size_t

type Order int32

const (
	RowMajor Order = 101
	ColMajor Order = 102
)

type Transpose int32

const (
	NoTrans   Transpose = 111
	Trans     Transpose = 112
	ConjTrans Transpose = 113
)

type Uplo int32
type Diag int32
type Side int32

/*
enum CBLAS_TRANSPOSE {CblasNoTrans=111, CblasTrans=112, CblasConjTrans=113};
enum CBLAS_UPLO {CblasUpper=121, CblasLower=122};
enum CBLAS_DIAG {CblasNonUnit=131, CblasUnit=132};
enum CBLAS_SIDE {CblasLeft=141, CblasRight=142};
*/
