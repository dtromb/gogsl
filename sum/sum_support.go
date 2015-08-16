package sum

/*
	#include <gsl/gsl_sum.h>

	size_t get_levin_u_workspace_size(gsl_sum_levin_u_workspace *ws) {
		return ws->size;
	}

	size_t get_levin_u_workspace_terms_used(gsl_sum_levin_u_workspace *ws) {
		return ws->terms_used;
	}

	double get_levin_u_workspace_sum_plain(gsl_sum_levin_u_workspace *ws) {
		return ws->sum_plain;
	}

	size_t get_levin_utrunc_workspace_size(gsl_sum_levin_utrunc_workspace *ws) {
		return ws->size;
	}

	size_t get_levin_utrunc_workspace_terms_used(gsl_sum_levin_utrunc_workspace *ws) {
		return ws->terms_used;
	}

	double get_levin_utrunc_workspace_sum_plain(gsl_sum_levin_utrunc_workspace *ws) {
		return ws->sum_plain;
	}
*/
import "C"

import "unsafe"

func (ws *GslSumLevinUWorkspace) Size() int {
	return int(C.get_levin_u_workspace_size((*C.gsl_sum_levin_u_workspace)(unsafe.Pointer(ws.Ptr()))))
}

func (ws *GslSumLevinUWorkspace) TermsUsed() int {
	return int(C.get_levin_u_workspace_terms_used((*C.gsl_sum_levin_u_workspace)(unsafe.Pointer(ws.Ptr()))))
}

func (ws *GslSumLevinUWorkspace) SumPlain() float64 {
	return float64(C.get_levin_u_workspace_sum_plain((*C.gsl_sum_levin_u_workspace)(unsafe.Pointer(ws.Ptr()))))
}

func (ws *GslSumLevinUtruncWorkspace) Size() int {
	return int(C.get_levin_u_workspace_size((*C.gsl_sum_levin_u_workspace)(unsafe.Pointer(ws.Ptr()))))
}

func (ws *GslSumLevinUtruncWorkspace) TermsUsed() int {
	return int(C.get_levin_u_workspace_terms_used((*C.gsl_sum_levin_u_workspace)(unsafe.Pointer(ws.Ptr()))))
}

func (ws *GslSumLevinUtruncWorkspace) SumPlain() float64 {
	return float64(C.get_levin_u_workspace_sum_plain((*C.gsl_sum_levin_u_workspace)(unsafe.Pointer(ws.Ptr()))))
}
