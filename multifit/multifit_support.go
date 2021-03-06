package multifit

/*

	#include <stdlib.h>
	#include <gsl/gsl_multifit.h>
	#include <gsl/gsl_multifit_nlin.h>
	#include <gsl/gsl_vector.h>
	#include <gsl/gsl_matrix.h>

	extern int gslMultifitFunctionCaller(const gsl_vector *,void *, gsl_vector*);
	int _cgo_gsl_multifit_function_proxy(const gsl_vector * x, void * params, gsl_vector * f) {
		return gslMultifitFunctionCaller(x, params, f);
	}

	extern int gslMultifitFunctionFdfFCaller(const gsl_vector *, void *, gsl_vector *);
	int _cgo_gsl_multifit_function_fdf_f_proxy (const gsl_vector * x, void * params, gsl_vector * f) {
		return gslMultifitFunctionFdfFCaller(x, params, f);
	}

	extern int gslMultifitFunctionFdfDfCaller(const gsl_vector*, void *, gsl_matrix *);
	int _cgo_gsl_multifit_function_fdf_df_proxy(const gsl_vector * x, void * params, gsl_matrix * J) {
		return gslMultifitFunctionFdfDfCaller(x, params, J);
	}

	extern int gslMultifitFunctionFdfCaller(const gsl_vector *, void *, gsl_vector*, gsl_matrix*);
	int _cgo_gsl_multifit_function_fdf_proxy(const gsl_vector * x, void * params, gsl_vector * f, gsl_matrix * J) {
		return gslMultifitFunctionFdfCaller(x, params, f, J);
	}

	void _init_proxy_gsl_multifit_function(void *params, size_t dim, gsl_multifit_function *gslFunc) {
		gslFunc->f = _cgo_gsl_multifit_function_proxy;
		gslFunc->n = dim;
		gslFunc->params = params;
	}

	void _init_proxy_gsl_multifit_function_fdf(void *params, size_t dim, gsl_multifit_function_fdf *gslFunc) {
		gslFunc->f = _cgo_gsl_multifit_function_fdf_f_proxy;
		gslFunc->df = _cgo_gsl_multifit_function_fdf_df_proxy;
		gslFunc->fdf = _cgo_gsl_multifit_function_fdf_proxy;
		gslFunc->n = dim;
		gslFunc->params = params;
	}

	gsl_vector *get_gsl_multifit_fsolver_dx_ptr(gsl_multifit_fsolver *s) {
		return s->dx;
	}

	gsl_vector *get_gsl_multifit_fsolver_f_ptr(gsl_multifit_fsolver *s) {
		return s->f;
	}

	gsl_vector *get_gsl_multifit_fsolver_x_ptr(gsl_multifit_fsolver *s) {
		return s->x;
	}

	gsl_vector *get_gsl_multifit_fdfsolver_dx_ptr(gsl_multifit_fdfsolver *s) {
		return s->dx;
	}

	gsl_vector *get_gsl_multifit_fdfsolver_f_ptr(gsl_multifit_fdfsolver *s) {
		return s->f;
	}

	gsl_vector *get_gsl_multifit_fdfsolver_x_ptr(gsl_multifit_fdfsolver *s) {
		return s->x;
	}

	gsl_matrix *get_gsl_multifit_fdfsolver_j_ptr(gsl_multifit_fdfsolver *s) {
		return s->J;
	}

	size_t get_gsl_multifit_function_struct_size() {
		return sizeof(gsl_multifit_function);
	}

	size_t get_gsl_multifit_function_fdf_struct_size() {
		return sizeof(gsl_multifit_function_fdf);
	}

	void get_robust_statistics(const gsl_multifit_robust_workspace * w,
								double *sigma_ols, double *sigma_mad, double *sigma_rob, double *sigma,
								double *rsq, double *adj_rsq, double *rmse, double *sse,
								size_t *dof, size_t *numit, gsl_vector **weights, gsl_vector **R) {
		gsl_multifit_robust_stats rv;
		rv = gsl_multifit_robust_statistics(w);
		*sigma_ols = rv.sigma_ols;
		*sigma_mad = rv.sigma_mad;
		*sigma_rob = rv.sigma_rob;
		*sigma = rv.sigma;
		*rsq = rv.Rsq;
		*adj_rsq = rv.adj_Rsq;
		*rmse = rv.rmse;
		*sse = rv.sse;
		*dof = rv.dof;
		*numit = rv.numit;
		*weights = rv.weights;
		*R = rv.r;
	}

*/
import "C"

import (
	"github.com/dtromb/gogsl"
	"github.com/dtromb/gogsl/matrix"
	"github.com/dtromb/gogsl/vector"
	"reflect"
	"unsafe"
)

var global_GSL_MULTIFIT_FUNCTION_STRUCT_SIZE uint32 = uint32(C.get_gsl_multifit_function_struct_size())
var global_GSL_MULTIFIT_FUNCTION_FDF_STRUCT_SIZE uint32 = uint32(C.get_gsl_multifit_function_fdf_struct_size())

var GSL_MULTIFIT_ROBUST_DEFAULT *GslMultifitRobustType = &GslMultifitRobustType{cPtr: uintptr(unsafe.Pointer(C.gsl_multifit_robust_default))}
var GSL_MULTIFIT_ROBUST_BISQUARE *GslMultifitRobustType = &GslMultifitRobustType{cPtr: uintptr(unsafe.Pointer(C.gsl_multifit_robust_bisquare))}
var GSL_MULTIFIT_ROBUST_CAUCHY *GslMultifitRobustType = &GslMultifitRobustType{cPtr: uintptr(unsafe.Pointer(C.gsl_multifit_robust_cauchy))}
var GSL_MULTIFIT_ROBUST_FAIR *GslMultifitRobustType = &GslMultifitRobustType{cPtr: uintptr(unsafe.Pointer(C.gsl_multifit_robust_fair))}
var GSL_MULTIFIT_ROBUST_HUBER *GslMultifitRobustType = &GslMultifitRobustType{cPtr: uintptr(unsafe.Pointer(C.gsl_multifit_robust_huber))}
var GSL_MULTIFIT_ROBUST_OLS *GslMultifitRobustType = &GslMultifitRobustType{cPtr: uintptr(unsafe.Pointer(C.gsl_multifit_robust_ols))}
var GSL_MULTIFIT_ROBUST_WELSCH *GslMultifitRobustType = &GslMultifitRobustType{cPtr: uintptr(unsafe.Pointer(C.gsl_multifit_robust_welsch))}

var GSL_MULTIFIT_FDFSOLVER_LMSDER *GslMultifitFdfsolverType = &GslMultifitFdfsolverType{cPtr: uintptr(unsafe.Pointer(C.gsl_multifit_fdfsolver_lmsder))}
var GSL_MULTIFIT_FDFSOLVER_LMDER *GslMultifitFdfsolverType = &GslMultifitFdfsolverType{cPtr: uintptr(unsafe.Pointer(C.gsl_multifit_fdfsolver_lmder))}

//gsl_multifit_robust_stats gsl_multifit_robust_statistics (const gsl_multifit_robust_workspace * w)
func RobustStatistics(rw *GslMultifitRobustWorkspace) *GslMultifitRobustStatistics {
	var dof C.size_t
	var numit C.size_t
	var weightsPtr *C.gsl_vector
	var rPtr *C.gsl_vector
	res := &GslMultifitRobustStatistics{}
	C.get_robust_statistics((*C.gsl_multifit_robust_workspace)(unsafe.Pointer(rw.Ptr())),
		(*C.double)(&res.SigmaOls), (*C.double)(&res.SigmaMad), (*C.double)(&res.SigmaRob), (*C.double)(&res.Sigma),
		(*C.double)(&res.Rsq), (*C.double)(&res.AdjRsq), (*C.double)(&res.Rmse), (*C.double)(&res.Sse),
		&dof, &numit, &weightsPtr, &rPtr)
	res.Dof = int(dof)
	res.Numit = int(numit)
	weights := &vector.GslVector{}
	gogsl.InitializeGslStruct(weights, uintptr(unsafe.Pointer(weightsPtr)))
	r := &vector.GslVector{}
	gogsl.InitializeGslStruct(r, uintptr(unsafe.Pointer(rPtr)))
	res.Weights = vector.VectorAlloc(weights.Length())
	res.R = vector.VectorAlloc(r.Length())
	vector.Memcpy(res.Weights, weights)
	vector.Memcpy(res.R, r)
	return res
}

func InitializeGslMultifitFunction(ptr *GslMultifitFunction) {
	ptr.cGslFunctionStruct = make([]byte, 0, global_GSL_MULTIFIT_FUNCTION_STRUCT_SIZE)
	cPtr := ((*reflect.SliceHeader)(unsafe.Pointer(&ptr.cGslFunctionStruct))).Data
	C._init_proxy_gsl_multifit_function(unsafe.Pointer(ptr), C.size_t(ptr.Dimension), (*C.gsl_multifit_function)(unsafe.Pointer(cPtr)))
}

func InitializeGslMultifitFunctionFdf(ptr *GslMultifitFunctionFdf) {
	ptr.cGslFunctionStruct = make([]byte, 0, global_GSL_MULTIFIT_FUNCTION_FDF_STRUCT_SIZE)
	cPtr := ((*reflect.SliceHeader)(unsafe.Pointer(&ptr.cGslFunctionStruct))).Data
	C._init_proxy_gsl_multifit_function_fdf(unsafe.Pointer(ptr), C.size_t(ptr.Dimension), (*C.gsl_multifit_function_fdf)(unsafe.Pointer(cPtr)))
}

func (s *GslMultifitFsolver) X_() *vector.GslVector {
	cptr := C.get_gsl_multifit_fsolver_x_ptr((*C.gsl_multifit_fsolver)(unsafe.Pointer(s.Ptr())))
	v := &vector.GslVector{}
	gogsl.InitializeGslStruct(v, uintptr(unsafe.Pointer(cptr)))
	return v
}

func (s *GslMultifitFsolver) F_() *vector.GslVector {
	cptr := C.get_gsl_multifit_fsolver_f_ptr((*C.gsl_multifit_fsolver)(unsafe.Pointer(s.Ptr())))
	v := &vector.GslVector{}
	gogsl.InitializeGslStruct(v, uintptr(unsafe.Pointer(cptr)))
	return v
}

func (s *GslMultifitFsolver) Dx_() *vector.GslVector {
	cptr := C.get_gsl_multifit_fsolver_dx_ptr((*C.gsl_multifit_fsolver)(unsafe.Pointer(s.Ptr())))
	v := &vector.GslVector{}
	gogsl.InitializeGslStruct(v, uintptr(unsafe.Pointer(cptr)))
	return v
}

func (s *GslMultifitFdfsolver) X_() *vector.GslVector {
	cptr := C.get_gsl_multifit_fdfsolver_x_ptr((*C.gsl_multifit_fdfsolver)(unsafe.Pointer(s.Ptr())))
	v := &vector.GslVector{}
	gogsl.InitializeGslStruct(v, uintptr(unsafe.Pointer(cptr)))
	return v
}

func (s *GslMultifitFdfsolver) F_() *vector.GslVector {
	cptr := C.get_gsl_multifit_fdfsolver_f_ptr((*C.gsl_multifit_fdfsolver)(unsafe.Pointer(s.Ptr())))
	v := &vector.GslVector{}
	gogsl.InitializeGslStruct(v, uintptr(unsafe.Pointer(cptr)))
	return v
}

func (s *GslMultifitFdfsolver) Dx_() *vector.GslVector {
	cptr := C.get_gsl_multifit_fdfsolver_dx_ptr((*C.gsl_multifit_fdfsolver)(unsafe.Pointer(s.Ptr())))
	v := &vector.GslVector{}
	gogsl.InitializeGslStruct(v, uintptr(unsafe.Pointer(cptr)))
	return v
}

func (s *GslMultifitFdfsolver) J_() *matrix.GslMatrix {
	cptr := C.get_gsl_multifit_fdfsolver_j_ptr((*C.gsl_multifit_fdfsolver)(unsafe.Pointer(s.Ptr())))
	m := &matrix.GslMatrix{}
	gogsl.InitializeGslStruct(m, uintptr(unsafe.Pointer(cptr)))
	return m
}

func (gf GslMultifitFunction) CPtr() uintptr {
	return ((*reflect.SliceHeader)(unsafe.Pointer(&gf.cGslFunctionStruct))).Data
}

func (gf GslMultifitFunctionFdf) CPtr() uintptr {
	return ((*reflect.SliceHeader)(unsafe.Pointer(&gf.cGslFunctionStruct))).Data
}

func (rt *GslMultifitRobustType) Ptr() uintptr {
	return rt.cPtr
}

func (rt *GslMultifitRobustType) Dispose() {
	rt.cPtr = 0
}

func (st *GslMultifitFsolverType) Ptr() uintptr {
	return st.cPtr
}

func (st *GslMultifitFsolverType) Dispose() {
	st.cPtr = 0
}

func (st *GslMultifitFdfsolverType) Ptr() uintptr {
	return st.cPtr
}

func (st *GslMultifitFdfsolverType) Dispose() {
	st.cPtr = 0
}
