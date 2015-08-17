package root

/*
	#include <gsl/gsl_roots.h>
*/
import "C"

import "unsafe"

var GSL_ROOT_FSOLVER_BISECTION *GslRootFsolverType = &GslRootFsolverType{cPtr: uintptr(unsafe.Pointer(C.gsl_root_fsolver_bisection))}
var GSL_ROOT_FSOLVER_FALSEPOS *GslRootFsolverType = &GslRootFsolverType{cPtr: uintptr(unsafe.Pointer(C.gsl_root_fsolver_falsepos))}
var GSL_ROOT_FSOLVER_BRENT *GslRootFsolverType = &GslRootFsolverType{cPtr: uintptr(unsafe.Pointer(C.gsl_root_fsolver_brent))}

var GSL_ROOT_FDFSOLVER_NEWTON *GslRootFdfsolverType = &GslRootFdfsolverType{cPtr: uintptr(unsafe.Pointer(C.gsl_root_fdfsolver_newton))}
var GSL_ROOT_FDFSOLVER_SECANT *GslRootFdfsolverType = &GslRootFdfsolverType{cPtr: uintptr(unsafe.Pointer(C.gsl_root_fdfsolver_secant))}
var GSL_ROOT_FDFSOLVER_STEFFENSON *GslRootFdfsolverType = &GslRootFdfsolverType{cPtr: uintptr(unsafe.Pointer(C.gsl_root_fdfsolver_steffenson))}

func (st *GslRootFsolverType) Ptr() uintptr {
	return st.cPtr
}

func (st *GslRootFsolverType) Dispose() {
	st.cPtr = 0
}

func (st *GslRootFdfsolverType) Ptr() uintptr {
	return st.cPtr
}

func (st *GslRootFdfsolverType) Dispose() {
	st.cPtr = 0
}
