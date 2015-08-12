package gogsl

import (
	"fmt"
	"testing"
)

func TestError(t *testing.T) {
	var gotErrorCalled bool
	handler := &GslErrorHandler{
		Handler: func(reason string, file string, line int, gslError GslError) {
			fmt.Printf("ERROR: %s at %s:%d (%s)\n", reason, file, line, gslError.String())
			gotErrorCalled = true
		},
	}
	SetErrorHandler(handler)
	Error("This should be a runaway iteration error.", GSL_ERUNAWAY)
	if !gotErrorCalled {
		t.Error("Custom error function was not called.")
	}
}
