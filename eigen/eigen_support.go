package eigen

/*
	#include <gsl/gsl_eigen.h>
*/
import "C"

type Sort uint32

var EIGEN_SORT_VAL_ASC		Sort	= C.GSL_EIGEN_SORT_VAL_ASC
var EIGEN_SORT_VAL_DESC		Sort	= C.GSL_EIGEN_SORT_VAL_DESC
var EIGEN_SORT_ABS_ASC		Sort	= C.GSL_EIGEN_SORT_ABS_ASC
var EIGEN_SORT_ABS_DESC		Sort	= C.GSL_EIGEN_SORT_ABS_DESC
