package numint

/*
	#include <gsl/gsl_integration.h>
	#include <stdio.h>

	size_t get_integration_workspace_interval_count(gsl_integration_workspace *w) {
		return w->size;
	}

*/
import "C"

import "unsafe"

type IntegrationRule int32

const (
	GSL_INTEG_GAUSS15 IntegrationRule = 1
	GSL_INTEG_GAUSS21 IntegrationRule = 2
	GSL_INTEG_GAUSS31 IntegrationRule = 3
	GSL_INTEG_GAUSS41 IntegrationRule = 4
	GSL_INTEG_GAUSS51 IntegrationRule = 5
	GSL_INTEG_GAUSS61 IntegrationRule = 6
)

func (w *GslIntegrationWorkspace) IntervalCount() int32 {
	return int32(C.get_integration_workspace_interval_count((*C.gsl_integration_workspace)(unsafe.Pointer(w.Ptr()))))
}
