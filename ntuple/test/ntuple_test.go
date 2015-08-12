package test

import (
	"github.com/dtromb/gogsl/histogram"
	"github.com/dtromb/gogsl/ntuple"
	"github.com/dtromb/gogsl/randist"
	"github.com/dtromb/gogsl/rng"
	"io/ioutil"
	"os"
	"testing"
)

type Data struct {
	x float64
	y float64
	z float64
}

var TEMP_FILE_NAME string

func init() {
	tmpFile, err := ioutil.TempFile("", "_gogsl_test_ntuple")
	if err != nil {
		panic("Cannot get temporary file")
	}
	TEMP_FILE_NAME = tmpFile.Name()
}

func TestNtupleWrite(t *testing.T) {

	dataBuf := make([]Data, 1)
	nt := ntuple.Create(TEMP_FILE_NAME, dataBuf)

	rng.EnvSetup()
	T := rng.DefaultRngType()
	r := rng.RngAlloc(T)
	for i := 0; i < 10000; i++ {
		dataBuf[0].x = randist.Ugaussian(r)
		dataBuf[0].y = randist.Ugaussian(r)
		dataBuf[0].z = randist.Ugaussian(r)
		ntuple.Write(nt)
	}
	ntuple.Close(nt)
}

func selectionFn(ntupleData interface{}, params interface{}) int {
	slice := ntupleData.([]Data)
	scale := *params.(*float64)
	data := &slice[0]
	x := data.x
	y := data.y
	z := data.z
	E2 := x*x + y*y + z*z
	if E2 > scale {
		//fmt.Printf("%f %f %f = %f ACCEPT\n",x,y,z,E2)
		return 1
	}
	//fmt.Printf("%f %f %f = %f REJECT\n",x,y,z,E2)
	return 0
}

func valueFn(ntupleData interface{}, params interface{}) float64 {
	slice := ntupleData.([]Data)
	data := &slice[0]
	x := data.x
	y := data.y
	z := data.z
	//fmt.Printf("VALUE: %f\n",  x * x + y * y + z * z)
	return x*x + y*y + z*z
}

func TestNtuple(t *testing.T) {

	var lower float64 = 1.5

	dataBuf := make([]Data, 1)
	nt := ntuple.Open(TEMP_FILE_NAME, dataBuf)

	h := histogram.HistogramAlloc(100)
	histogram.SetRangesUniform(h, 0.0, 10.0)

	selection := ntuple.GslNtupleSelectionFunction{
		Function: selectionFn,
		Params:   &lower,
	}
	ntuple.InitializeGslNtupleSelectionFunction(&selection)

	value := ntuple.GslNtupleValueFunction{
		Function: valueFn,
	}
	ntuple.InitializeGslNtupleValueFunction(&value)

	ntuple.Project(h, nt, &value, &selection)
	histogram.Fprintf(os.Stdout, h, "%f", "%f")
}
