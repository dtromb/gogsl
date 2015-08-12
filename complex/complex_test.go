package complex

import (
	"fmt"
	"testing"
)

func TestComplex(t *testing.T) {
	fmt.Printf("%f\n", Abs(complex(3, 4)))
	m := Mul(complex(1, 2), complex(3, 4))
	fmt.Printf("%f+%fi\n", real(m), imag(m))
}
