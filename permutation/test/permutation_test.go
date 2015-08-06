package test


import (
	"fmt"
	"testing"
	"os"
	"github.com/dtromb/gogsl/permutation"
	"github.com/dtromb/gogsl/rng"
	"github.com/dtromb/gogsl/randist"
)

func TestPermutation(t *testing.T) {
	var N int = 15
  	p := permutation.PermutationAlloc(N)
  	q := permutation.PermutationAlloc(N)
	rng.EnvSetup()
	T := rng.DefaultRngType()
	r := rng.RngAlloc(T)
	fmt.Printf ("initial permutation: ")
	permutation.PermutationInit(p)
	permutation.Fprintf(os.Stdout, p," %u")
	os.Stdout.Sync()
	fmt.Printf("\n")
	fmt.Printf (" random permutation: ")
	randist.Shuffle(r, p.Slice_(), p.Len())
	permutation.Fprintf(os.Stdout, p," %u")
    fmt.Printf("\n")
	fmt.Printf (" inverse permutation: ")
	permutation.Inverse(q,p)
	permutation.Fprintf(os.Stdout, q," %u")
    fmt.Printf("\n")
}
