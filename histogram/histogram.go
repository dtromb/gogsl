//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package histogram

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_histogram.h>
   #include <gsl/gsl_histogram2d.h>
   #include <unistd.h>
*/
import "C"

import "unsafe"
import "github.com/dtromb/gogsl"
import "reflect"
import "os"

type GslHistogram struct {
	gogsl.GslReference
}

type GslHistogramPdf struct {
	gogsl.GslReference
}

type GslHistogram2d struct {
	gogsl.GslReference
}

type GslHistogram2dPdf struct {
	gogsl.GslReference
}

func HistogramAlloc(n int) *GslHistogram {
	_ref := C.gsl_histogram_alloc(C.size_t(n))
	_result := &GslHistogram{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func SetRanges(h *GslHistogram, r []float64, size int) int32 {
	_slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&r))
	return int32(C.gsl_histogram_set_ranges((*C.gsl_histogram)(unsafe.Pointer(h.Ptr())), (*C.double)(unsafe.Pointer(_slice_header_1.Data)), C.size_t(size)))
}

func SetRangesUniform(h *GslHistogram, xmin float64, xmax float64) int32 {
	return int32(C.gsl_histogram_set_ranges_uniform((*C.gsl_histogram)(unsafe.Pointer(h.Ptr())), C.double(xmin), C.double(xmax)))
}

func (x *GslHistogram) Dispose() {
	C.gsl_histogram_free((*C.gsl_histogram)(unsafe.Pointer(x.Ptr())))
}

func Memcpy(dest *GslHistogram, src *GslHistogram) int32 {
	return int32(C.gsl_histogram_memcpy((*C.gsl_histogram)(unsafe.Pointer(dest.Ptr())), (*C.gsl_histogram)(unsafe.Pointer(src.Ptr()))))
}

func Clone(src *GslHistogram) *GslHistogram {
	_ref := C.gsl_histogram_clone((*C.gsl_histogram)(unsafe.Pointer(src.Ptr())))
	_result := &GslHistogram{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func Increment(h *GslHistogram, x float64) int32 {
	return int32(C.gsl_histogram_increment((*C.gsl_histogram)(unsafe.Pointer(h.Ptr())), C.double(x)))
}

func Accumulate(h *GslHistogram, x float64, weight float64) int32 {
	return int32(C.gsl_histogram_accumulate((*C.gsl_histogram)(unsafe.Pointer(h.Ptr())), C.double(x), C.double(weight)))
}

func Get(h *GslHistogram, i int) float64 {
	return float64(C.gsl_histogram_get((*C.gsl_histogram)(unsafe.Pointer(h.Ptr())), C.size_t(i)))
}

func GetRange(h *GslHistogram, i int) (int32, float64, float64) {
	var _outptr_2 C.double
	var _outptr_3 C.double
	_result := int32(C.gsl_histogram_get_range((*C.gsl_histogram)(unsafe.Pointer(h.Ptr())), C.size_t(i), &_outptr_2, &_outptr_3))
	return _result, *(*float64)(unsafe.Pointer(&_outptr_2)), *(*float64)(unsafe.Pointer(&_outptr_3))
}

func Max(h *GslHistogram) float64 {
	return float64(C.gsl_histogram_max((*C.gsl_histogram)(unsafe.Pointer(h.Ptr()))))
}

func Min(h *GslHistogram) float64 {
	return float64(C.gsl_histogram_min((*C.gsl_histogram)(unsafe.Pointer(h.Ptr()))))
}

func Bins(h *GslHistogram) int {
	return int(C.gsl_histogram_bins((*C.gsl_histogram)(unsafe.Pointer(h.Ptr()))))
}

func Reset(h *GslHistogram) {
	C.gsl_histogram_reset((*C.gsl_histogram)(unsafe.Pointer(h.Ptr())))
}

func Find(h *GslHistogram, x float64) (int32, int) {
	var _outptr_2 C.size_t
	_result := int32(C.gsl_histogram_find((*C.gsl_histogram)(unsafe.Pointer(h.Ptr())), C.double(x), &_outptr_2))
	return _result, *(*int)(unsafe.Pointer(&_outptr_2))
}

func MaxVal(h *GslHistogram) float64 {
	return float64(C.gsl_histogram_max_val((*C.gsl_histogram)(unsafe.Pointer(h.Ptr()))))
}

func MaxBin(h *GslHistogram) int {
	return int(C.gsl_histogram_max_bin((*C.gsl_histogram)(unsafe.Pointer(h.Ptr()))))
}

func MinVal(h *GslHistogram) float64 {
	return float64(C.gsl_histogram_min_val((*C.gsl_histogram)(unsafe.Pointer(h.Ptr()))))
}

func MinBin(h *GslHistogram) int {
	return int(C.gsl_histogram_min_bin((*C.gsl_histogram)(unsafe.Pointer(h.Ptr()))))
}

func Mean(h *GslHistogram) float64 {
	return float64(C.gsl_histogram_mean((*C.gsl_histogram)(unsafe.Pointer(h.Ptr()))))
}

func Sigma(h *GslHistogram) float64 {
	return float64(C.gsl_histogram_sigma((*C.gsl_histogram)(unsafe.Pointer(h.Ptr()))))
}

func Sum(h *GslHistogram) float64 {
	return float64(C.gsl_histogram_sum((*C.gsl_histogram)(unsafe.Pointer(h.Ptr()))))
}

func EqualBinsP(h1 *GslHistogram, h2 *GslHistogram) int32 {
	return int32(C.gsl_histogram_equal_bins_p((*C.gsl_histogram)(unsafe.Pointer(h1.Ptr())), (*C.gsl_histogram)(unsafe.Pointer(h2.Ptr()))))
}

func Add(h1 *GslHistogram, h2 *GslHistogram) int32 {
	return int32(C.gsl_histogram_add((*C.gsl_histogram)(unsafe.Pointer(h1.Ptr())), (*C.gsl_histogram)(unsafe.Pointer(h2.Ptr()))))
}

func Sub(h1 *GslHistogram, h2 *GslHistogram) int32 {
	return int32(C.gsl_histogram_sub((*C.gsl_histogram)(unsafe.Pointer(h1.Ptr())), (*C.gsl_histogram)(unsafe.Pointer(h2.Ptr()))))
}

func Mul(h1 *GslHistogram, h2 *GslHistogram) int32 {
	return int32(C.gsl_histogram_mul((*C.gsl_histogram)(unsafe.Pointer(h1.Ptr())), (*C.gsl_histogram)(unsafe.Pointer(h2.Ptr()))))
}

func Div(h1 *GslHistogram, h2 *GslHistogram) int32 {
	return int32(C.gsl_histogram_div((*C.gsl_histogram)(unsafe.Pointer(h1.Ptr())), (*C.gsl_histogram)(unsafe.Pointer(h2.Ptr()))))
}

func Scale(h *GslHistogram, scale float64) int32 {
	return int32(C.gsl_histogram_scale((*C.gsl_histogram)(unsafe.Pointer(h.Ptr())), C.double(scale)))
}

func Shift(h *GslHistogram, offset float64) int32 {
	return int32(C.gsl_histogram_shift((*C.gsl_histogram)(unsafe.Pointer(h.Ptr())), C.double(offset)))
}

func Fwrite(stream *os.File, h *GslHistogram) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_histogram_fwrite(_file_0, (*C.gsl_histogram)(unsafe.Pointer(h.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func Fread(stream *os.File, h *GslHistogram) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_histogram_fread(_file_0, (*C.gsl_histogram)(unsafe.Pointer(h.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func Fprintf(stream *os.File, h *GslHistogram, rangeFormat string, binFormat string) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_string_2 := C.CString(rangeFormat)
	_string_3 := C.CString(binFormat)
	_result := int32(C.gsl_histogram_fprintf(_file_0, (*C.gsl_histogram)(unsafe.Pointer(h.Ptr())), _string_2, _string_3))
	C.fclose(_file_0)
	C.free(unsafe.Pointer(_string_2))
	C.free(unsafe.Pointer(_string_3))
	return _result
}

func Fscanf(stream *os.File, h *GslHistogram) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_histogram_fscanf(_file_0, (*C.gsl_histogram)(unsafe.Pointer(h.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func PdfAlloc(n int) *GslHistogramPdf {
	_ref := C.gsl_histogram_pdf_alloc(C.size_t(n))
	_result := &GslHistogramPdf{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func PdfInit(p *GslHistogramPdf, h *GslHistogram) int32 {
	return int32(C.gsl_histogram_pdf_init((*C.gsl_histogram_pdf)(unsafe.Pointer(p.Ptr())), (*C.gsl_histogram)(unsafe.Pointer(h.Ptr()))))
}

func (x *GslHistogramPdf) Dispose() {
	C.gsl_histogram_pdf_free((*C.gsl_histogram_pdf)(unsafe.Pointer(x.Ptr())))
}

func PdfSample(p *GslHistogramPdf, r float64) float64 {
	return float64(C.gsl_histogram_pdf_sample((*C.gsl_histogram_pdf)(unsafe.Pointer(p.Ptr())), C.double(r)))
}

func Histogram2dAlloc(nx int, ny int) *GslHistogram2d {
	_ref := C.gsl_histogram2d_alloc(C.size_t(nx), C.size_t(ny))
	_result := &GslHistogram2d{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func Histogram2dSetRanges(h *GslHistogram2d, xrange []float64, xsize int, yrange []float64, ysize int) int32 {
	_slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&xrange))
	_slice_header_3 := (*reflect.SliceHeader)(unsafe.Pointer(&yrange))
	return int32(C.gsl_histogram2d_set_ranges((*C.gsl_histogram2d)(unsafe.Pointer(h.Ptr())), (*C.double)(unsafe.Pointer(_slice_header_1.Data)), C.size_t(xsize), (*C.double)(unsafe.Pointer(_slice_header_3.Data)), C.size_t(ysize)))
}

func Histogram2dSetRangesUniform(h *GslHistogram2d, xmin float64, xmax float64, ymin float64, ymax float64) int32 {
	return int32(C.gsl_histogram2d_set_ranges_uniform((*C.gsl_histogram2d)(unsafe.Pointer(h.Ptr())), C.double(xmin), C.double(xmax), C.double(ymin), C.double(ymax)))
}

func (x *GslHistogram2d) Dispose() {
	C.gsl_histogram2d_free((*C.gsl_histogram2d)(unsafe.Pointer(x.Ptr())))
}

func Histogram2dMemcpy(dest *GslHistogram2d, src *GslHistogram2d) int32 {
	return int32(C.gsl_histogram2d_memcpy((*C.gsl_histogram2d)(unsafe.Pointer(dest.Ptr())), (*C.gsl_histogram2d)(unsafe.Pointer(src.Ptr()))))
}

func Histogram2dClone(src *GslHistogram2d) *GslHistogram2d {
	_ref := C.gsl_histogram2d_clone((*C.gsl_histogram2d)(unsafe.Pointer(src.Ptr())))
	_result := &GslHistogram2d{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func Histogram2dIncrement(h *GslHistogram2d, x float64, y float64) int32 {
	return int32(C.gsl_histogram2d_increment((*C.gsl_histogram2d)(unsafe.Pointer(h.Ptr())), C.double(x), C.double(y)))
}

func Histogram2dAccumulate(h *GslHistogram2d, x float64, y float64, weight float64) int32 {
	return int32(C.gsl_histogram2d_accumulate((*C.gsl_histogram2d)(unsafe.Pointer(h.Ptr())), C.double(x), C.double(y), C.double(weight)))
}

func Histogram2dGet(h *GslHistogram2d, i int, j int) float64 {
	return float64(C.gsl_histogram2d_get((*C.gsl_histogram2d)(unsafe.Pointer(h.Ptr())), C.size_t(i), C.size_t(j)))
}

func Histogram2dGetXrange(h *GslHistogram2d, i int) (int32, float64, float64) {
	var _outptr_2 C.double
	var _outptr_3 C.double
	_result := int32(C.gsl_histogram2d_get_xrange((*C.gsl_histogram2d)(unsafe.Pointer(h.Ptr())), C.size_t(i), &_outptr_2, &_outptr_3))
	return _result, *(*float64)(unsafe.Pointer(&_outptr_2)), *(*float64)(unsafe.Pointer(&_outptr_3))
}

func Histogram2dGetYrange(h *GslHistogram2d, j int) (int32, float64, float64) {
	var _outptr_2 C.double
	var _outptr_3 C.double
	_result := int32(C.gsl_histogram2d_get_yrange((*C.gsl_histogram2d)(unsafe.Pointer(h.Ptr())), C.size_t(j), &_outptr_2, &_outptr_3))
	return _result, *(*float64)(unsafe.Pointer(&_outptr_2)), *(*float64)(unsafe.Pointer(&_outptr_3))
}

func Histogram2dXmax(h *GslHistogram2d) float64 {
	return float64(C.gsl_histogram2d_xmax((*C.gsl_histogram2d)(unsafe.Pointer(h.Ptr()))))
}

func Histogram2dXmin(h *GslHistogram2d) float64 {
	return float64(C.gsl_histogram2d_xmin((*C.gsl_histogram2d)(unsafe.Pointer(h.Ptr()))))
}

func Histogram2dNx(h *GslHistogram2d) int {
	return int(C.gsl_histogram2d_nx((*C.gsl_histogram2d)(unsafe.Pointer(h.Ptr()))))
}

func Histogram2dYmax(h *GslHistogram2d) float64 {
	return float64(C.gsl_histogram2d_ymax((*C.gsl_histogram2d)(unsafe.Pointer(h.Ptr()))))
}

func Histogram2dYmin(h *GslHistogram2d) float64 {
	return float64(C.gsl_histogram2d_ymin((*C.gsl_histogram2d)(unsafe.Pointer(h.Ptr()))))
}

func Histogram2dNy(h *GslHistogram2d) int {
	return int(C.gsl_histogram2d_ny((*C.gsl_histogram2d)(unsafe.Pointer(h.Ptr()))))
}

func Histogram2dReset(h *GslHistogram2d) {
	C.gsl_histogram2d_reset((*C.gsl_histogram2d)(unsafe.Pointer(h.Ptr())))
}

func Histogram2dFind(h *GslHistogram2d, x float64, y float64) (int32, int, int) {
	var _outptr_3 C.size_t
	var _outptr_4 C.size_t
	_result := int32(C.gsl_histogram2d_find((*C.gsl_histogram2d)(unsafe.Pointer(h.Ptr())), C.double(x), C.double(y), &_outptr_3, &_outptr_4))
	return _result, *(*int)(unsafe.Pointer(&_outptr_3)), *(*int)(unsafe.Pointer(&_outptr_4))
}

func Histogram2dMaxVal(h *GslHistogram2d) float64 {
	return float64(C.gsl_histogram2d_max_val((*C.gsl_histogram2d)(unsafe.Pointer(h.Ptr()))))
}

func Histogram2dMaxBin(h *GslHistogram2d) (int, int) {
	var _outptr_1 C.size_t
	var _outptr_2 C.size_t
	C.gsl_histogram2d_max_bin((*C.gsl_histogram2d)(unsafe.Pointer(h.Ptr())), &_outptr_1, &_outptr_2)
	return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2))
}

func Histogram2dMinVal(h *GslHistogram2d) float64 {
	return float64(C.gsl_histogram2d_min_val((*C.gsl_histogram2d)(unsafe.Pointer(h.Ptr()))))
}

func Histogram2dMinBin(h *GslHistogram2d) (int, int) {
	var _outptr_1 C.size_t
	var _outptr_2 C.size_t
	C.gsl_histogram2d_min_bin((*C.gsl_histogram2d)(unsafe.Pointer(h.Ptr())), &_outptr_1, &_outptr_2)
	return *(*int)(unsafe.Pointer(&_outptr_1)), *(*int)(unsafe.Pointer(&_outptr_2))
}

func Histogram2dXmean(h *GslHistogram2d) float64 {
	return float64(C.gsl_histogram2d_xmean((*C.gsl_histogram2d)(unsafe.Pointer(h.Ptr()))))
}

func Histogram2dYmean(h *GslHistogram2d) float64 {
	return float64(C.gsl_histogram2d_ymean((*C.gsl_histogram2d)(unsafe.Pointer(h.Ptr()))))
}

func Histogram2dXsigma(h *GslHistogram2d) float64 {
	return float64(C.gsl_histogram2d_xsigma((*C.gsl_histogram2d)(unsafe.Pointer(h.Ptr()))))
}

func Histogram2dYsigma(h *GslHistogram2d) float64 {
	return float64(C.gsl_histogram2d_ysigma((*C.gsl_histogram2d)(unsafe.Pointer(h.Ptr()))))
}

func Histogram2dCov(h *GslHistogram2d) float64 {
	return float64(C.gsl_histogram2d_cov((*C.gsl_histogram2d)(unsafe.Pointer(h.Ptr()))))
}

func Histogram2dSum(h *GslHistogram2d) float64 {
	return float64(C.gsl_histogram2d_sum((*C.gsl_histogram2d)(unsafe.Pointer(h.Ptr()))))
}

func Histogram2dEqualBinsP(h1 *GslHistogram2d, h2 *GslHistogram2d) int32 {
	return int32(C.gsl_histogram2d_equal_bins_p((*C.gsl_histogram2d)(unsafe.Pointer(h1.Ptr())), (*C.gsl_histogram2d)(unsafe.Pointer(h2.Ptr()))))
}

func Histogram2dAdd(h1 *GslHistogram2d, h2 *GslHistogram2d) int32 {
	return int32(C.gsl_histogram2d_add((*C.gsl_histogram2d)(unsafe.Pointer(h1.Ptr())), (*C.gsl_histogram2d)(unsafe.Pointer(h2.Ptr()))))
}

func Histogram2dSub(h1 *GslHistogram2d, h2 *GslHistogram2d) int32 {
	return int32(C.gsl_histogram2d_sub((*C.gsl_histogram2d)(unsafe.Pointer(h1.Ptr())), (*C.gsl_histogram2d)(unsafe.Pointer(h2.Ptr()))))
}

func Histogram2dMul(h1 *GslHistogram2d, h2 *GslHistogram2d) int32 {
	return int32(C.gsl_histogram2d_mul((*C.gsl_histogram2d)(unsafe.Pointer(h1.Ptr())), (*C.gsl_histogram2d)(unsafe.Pointer(h2.Ptr()))))
}

func Histogram2dDiv(h1 *GslHistogram2d, h2 *GslHistogram2d) int32 {
	return int32(C.gsl_histogram2d_div((*C.gsl_histogram2d)(unsafe.Pointer(h1.Ptr())), (*C.gsl_histogram2d)(unsafe.Pointer(h2.Ptr()))))
}

func Histogram2dScale(h *GslHistogram2d, scale float64) int32 {
	return int32(C.gsl_histogram2d_scale((*C.gsl_histogram2d)(unsafe.Pointer(h.Ptr())), C.double(scale)))
}

func Histogram2dShift(h *GslHistogram2d, offset float64) int32 {
	return int32(C.gsl_histogram2d_shift((*C.gsl_histogram2d)(unsafe.Pointer(h.Ptr())), C.double(offset)))
}

func Histogram2dFwrite(stream *os.File, h *GslHistogram2d) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_histogram2d_fwrite(_file_0, (*C.gsl_histogram2d)(unsafe.Pointer(h.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func Histogram2dFread(stream *os.File, h *GslHistogram2d) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_histogram2d_fread(_file_0, (*C.gsl_histogram2d)(unsafe.Pointer(h.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func Histogram2dFprintf(stream *os.File, h *GslHistogram2d, rangeFormat string, binFormat string) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_string_2 := C.CString(rangeFormat)
	_string_3 := C.CString(binFormat)
	_result := int32(C.gsl_histogram2d_fprintf(_file_0, (*C.gsl_histogram2d)(unsafe.Pointer(h.Ptr())), _string_2, _string_3))
	C.fclose(_file_0)
	C.free(unsafe.Pointer(_string_2))
	C.free(unsafe.Pointer(_string_3))
	return _result
}

func Histogram2dFscanf(stream *os.File, h *GslHistogram2d) int32 {
	_file_0 := C.fdopen(C.dup(C.int(stream.Fd())), (*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
	_result := int32(C.gsl_histogram2d_fscanf(_file_0, (*C.gsl_histogram2d)(unsafe.Pointer(h.Ptr()))))
	C.fclose(_file_0)
	return _result
}

func Histogram2dPdfAlloc(nx int, ny int) *GslHistogram2dPdf {
	_ref := C.gsl_histogram2d_pdf_alloc(C.size_t(nx), C.size_t(ny))
	_result := &GslHistogram2dPdf{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func Histogram2dPdfInit(p *GslHistogram2dPdf, h *GslHistogram2d) int32 {
	return int32(C.gsl_histogram2d_pdf_init((*C.gsl_histogram2d_pdf)(unsafe.Pointer(p.Ptr())), (*C.gsl_histogram2d)(unsafe.Pointer(h.Ptr()))))
}

func (x *GslHistogram2dPdf) Dispose() {
	C.gsl_histogram2d_pdf_free((*C.gsl_histogram2d_pdf)(unsafe.Pointer(x.Ptr())))
}

func Histogram2dPdfSample(p *GslHistogram2dPdf, r1 float64, r2 float64) (int32, float64, float64) {
	var _outptr_3 C.double
	var _outptr_4 C.double
	_result := int32(C.gsl_histogram2d_pdf_sample((*C.gsl_histogram2d_pdf)(unsafe.Pointer(p.Ptr())), C.double(r1), C.double(r2), &_outptr_3, &_outptr_4))
	return _result, *(*float64)(unsafe.Pointer(&_outptr_3)), *(*float64)(unsafe.Pointer(&_outptr_4))
}
