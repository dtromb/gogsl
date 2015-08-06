//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package stats

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_statistics.h>
   #include <gsl/gsl_complex.h>
*/
import "C"

import "reflect"
import "unsafe"

func Mean(data []float64, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_mean((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func Variance(data []float64, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_variance((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func VarianceM(data []float64, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_variance_m((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func Sd(data []float64, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_sd((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func SdM(data []float64, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_sd_m((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func Tss(data []float64, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_tss((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func TssM(data []float64, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_tss_m((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func VarianceWithFixedMean(data []float64, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_variance_with_fixed_mean((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func SdWithFixedMean(data []float64, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_sd_with_fixed_mean((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func Absdev(data []float64, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_absdev((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func AbsdevM(data []float64, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_absdev_m((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func Skew(data []float64, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_skew((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func SkewMSd(data []float64, stride int, n int, mean float64, sd float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_skew_m_sd((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean), C.double(sd)))
}

func Kurtosis(data []float64, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_kurtosis((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func KurtosisMSd(data []float64, stride int, n int, mean float64, sd float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_kurtosis_m_sd((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean), C.double(sd)))
}

func Lag1Autocorrelation(data []float64, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_lag1_autocorrelation((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func Lag1AutocorrelationM(data []float64, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_lag1_autocorrelation_m((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func Covariance(data1 []float64, stride1 int, data2 []float64, stride2 int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   return float64(C.gsl_stats_covariance((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n)))
}

func CovarianceM(data1 []float64, stride1 int, data2 []float64, stride2 int, n int, mean1 float64, mean2 float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   return float64(C.gsl_stats_covariance_m((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n), C.double(mean1), C.double(mean2)))
}

func Correlation(data1 []float64, stride1 int, data2 []float64, stride2 int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   return float64(C.gsl_stats_correlation((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n)))
}

func Spearman(data1 []float64, stride1 int, data2 []float64, stride2 int, n int, work []float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   _slice_header_5 := (*reflect.SliceHeader)(unsafe.Pointer(&work))
   return float64(C.gsl_stats_spearman((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n), (*C.double)(unsafe.Pointer(_slice_header_5.Data))))
}

func Wmean(w []float64, wstride int, data []float64, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&w))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_wmean((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(wstride), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func Wvariance(w []float64, wstride int, data []float64, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&w))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_wvariance((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(wstride), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func WvarianceM(w []float64, wstride int, data []float64, stride int, n int, wmean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&w))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_wvariance_m((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(wstride), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n), C.double(wmean)))
}

func Wsd(w []float64, wstride int, data []float64, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&w))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_wsd((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(wstride), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func WsdM(w []float64, wstride int, data []float64, stride int, n int, wmean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&w))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_wsd_m((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(wstride), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n), C.double(wmean)))
}

func WvarianceWithFixedMean(w []float64, wstride int, data []float64, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&w))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_wvariance_with_fixed_mean((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(wstride), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func WsdWithFixedMean(w []float64, wstride int, data []float64, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&w))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_wsd_with_fixed_mean((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(wstride), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func Wtss(w []float64, wstride int, data []float64, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&w))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_wtss((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(wstride), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func WtssM(w []float64, wstride int, data []float64, stride int, n int, wmean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&w))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_wtss_m((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(wstride), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n), C.double(wmean)))
}

func Wabsdev(w []float64, wstride int, data []float64, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&w))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_wabsdev((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(wstride), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func WabsdevM(w []float64, wstride int, data []float64, stride int, n int, wmean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&w))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_wabsdev_m((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(wstride), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n), C.double(wmean)))
}

func Wskew(w []float64, wstride int, data []float64, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&w))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_wskew((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(wstride), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func WskewMSd(w []float64, wstride int, data []float64, stride int, n int, wmean float64, wsd float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&w))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_wskew_m_sd((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(wstride), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n), C.double(wmean), C.double(wsd)))
}

func Wkurtosis(w []float64, wstride int, data []float64, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&w))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_wkurtosis((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(wstride), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func WkurtosisMSd(w []float64, wstride int, data []float64, stride int, n int, wmean float64, wsd float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&w))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_wkurtosis_m_sd((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(wstride), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n), C.double(wmean), C.double(wsd)))
}

func FloatWmean(w []float32, wstride int, data []float32, stride int, n int) float32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&w))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float32(C.gsl_stats_float_wmean((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(wstride), (*C.float)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func FloatWvariance(w []float32, wstride int, data []float32, stride int, n int) float32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&w))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float32(C.gsl_stats_float_wvariance((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(wstride), (*C.float)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func FloatWvarianceM(w []float32, wstride int, data []float32, stride int, n int, wmean float64) float32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&w))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float32(C.gsl_stats_float_wvariance_m((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(wstride), (*C.float)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n), C.double(wmean)))
}

func FloatWsd(w []float32, wstride int, data []float32, stride int, n int) float32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&w))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float32(C.gsl_stats_float_wsd((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(wstride), (*C.float)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func FloatWsdM(w []float32, wstride int, data []float32, stride int, n int, wmean float64) float32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&w))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float32(C.gsl_stats_float_wsd_m((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(wstride), (*C.float)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n), C.double(wmean)))
}

func FloatWvarianceWithFixedMean(w []float32, wstride int, data []float32, stride int, n int, mean float64) float32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&w))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float32(C.gsl_stats_float_wvariance_with_fixed_mean((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(wstride), (*C.float)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func FloatWsdWithFixedMean(w []float32, wstride int, data []float32, stride int, n int, mean float64) float32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&w))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float32(C.gsl_stats_float_wsd_with_fixed_mean((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(wstride), (*C.float)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func FloatWtss(w []float32, wstride int, data []float32, stride int, n int) float32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&w))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float32(C.gsl_stats_float_wtss((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(wstride), (*C.float)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func FloatWtssM(w []float32, wstride int, data []float32, stride int, n int, wmean float64) float32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&w))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float32(C.gsl_stats_float_wtss_m((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(wstride), (*C.float)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n), C.double(wmean)))
}

func FloatWabsdev(w []float32, wstride int, data []float32, stride int, n int) float32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&w))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float32(C.gsl_stats_float_wabsdev((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(wstride), (*C.float)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func FloatWabsdevM(w []float32, wstride int, data []float32, stride int, n int, wmean float64) float32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&w))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float32(C.gsl_stats_float_wabsdev_m((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(wstride), (*C.float)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n), C.double(wmean)))
}

func FloatWskew(w []float32, wstride int, data []float32, stride int, n int) float32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&w))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float32(C.gsl_stats_float_wskew((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(wstride), (*C.float)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func FloatWskewMSd(w []float32, wstride int, data []float32, stride int, n int, wmean float64, wsd float64) float32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&w))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float32(C.gsl_stats_float_wskew_m_sd((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(wstride), (*C.float)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n), C.double(wmean), C.double(wsd)))
}

func FloatWkurtosis(w []float32, wstride int, data []float32, stride int, n int) float32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&w))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float32(C.gsl_stats_float_wkurtosis((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(wstride), (*C.float)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func FloatWkurtosisMSd(w []float32, wstride int, data []float32, stride int, n int, wmean float64, wsd float64) float32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&w))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float32(C.gsl_stats_float_wkurtosis_m_sd((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(wstride), (*C.float)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n), C.double(wmean), C.double(wsd)))
}

func Max(data []float64, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_max((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func Min(data []float64, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_min((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func Minmax(data []float64, stride int, n int) (float64, float64) {
   var _outptr_0 C.double
   var _outptr_1 C.double
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   C.gsl_stats_minmax(&_outptr_0, &_outptr_1, (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n))
   return *(*float64)(unsafe.Pointer(&_outptr_0)), *(*float64)(unsafe.Pointer(&_outptr_1))
}

func MaxIndex(data []float64, stride int, n int) int {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return int(C.gsl_stats_max_index((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func MinIndex(data []float64, stride int, n int) int {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return int(C.gsl_stats_min_index((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func MinmaxIndex(data []float64, stride int, n int) (int, int) {
   var _outptr_0 C.size_t
   var _outptr_1 C.size_t
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   C.gsl_stats_minmax_index(&_outptr_0, &_outptr_1, (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n))
   return *(*int)(unsafe.Pointer(&_outptr_0)), *(*int)(unsafe.Pointer(&_outptr_1))
}

func MedianFromSortedData(sortedData []float64, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&sortedData))
   return float64(C.gsl_stats_median_from_sorted_data((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func QuantileFromSortedData(sortedData []float64, stride int, n int, f float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&sortedData))
   return float64(C.gsl_stats_quantile_from_sorted_data((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(f)))
}

func FloatMean(data []float32, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_float_mean((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func FloatVariance(data []float32, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_float_variance((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func FloatVarianceM(data []float32, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_float_variance_m((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func FloatSd(data []float32, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_float_sd((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func FloatSdM(data []float32, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_float_sd_m((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func FloatTss(data []float32, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_float_tss((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func FloatTssM(data []float32, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_float_tss_m((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func FloatVarianceWithFixedMean(data []float32, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_float_variance_with_fixed_mean((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func FloatSdWithFixedMean(data []float32, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_float_sd_with_fixed_mean((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func FloatAbsdev(data []float32, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_float_absdev((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func FloatAbsdevM(data []float32, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_float_absdev_m((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func FloatSkew(data []float32, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_float_skew((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func FloatSkewMSd(data []float32, stride int, n int, mean float64, sd float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_float_skew_m_sd((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean), C.double(sd)))
}

func FloatKurtosis(data []float32, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_float_kurtosis((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func FloatKurtosisMSd(data []float32, stride int, n int, mean float64, sd float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_float_kurtosis_m_sd((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean), C.double(sd)))
}

func FloatLag1Autocorrelation(data []float32, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_float_lag1_autocorrelation((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func FloatLag1AutocorrelationM(data []float32, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_float_lag1_autocorrelation_m((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func FloatCovariance(data1 []float32, stride1 int, data2 []float32, stride2 int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   return float64(C.gsl_stats_float_covariance((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.float)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n)))
}

func FloatCovarianceM(data1 []float32, stride1 int, data2 []float32, stride2 int, n int, mean1 float64, mean2 float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   return float64(C.gsl_stats_float_covariance_m((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.float)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n), C.double(mean1), C.double(mean2)))
}

func FloatCorrelation(data1 []float32, stride1 int, data2 []float32, stride2 int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   return float64(C.gsl_stats_float_correlation((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.float)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n)))
}

func FloatSpearman(data1 []float32, stride1 int, data2 []float32, stride2 int, n int, work []float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   _slice_header_5 := (*reflect.SliceHeader)(unsafe.Pointer(&work))
   return float64(C.gsl_stats_float_spearman((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.float)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n), (*C.double)(unsafe.Pointer(_slice_header_5.Data))))
}

func FloatMax(data []float32, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_float_max((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func FloatMin(data []float32, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_float_min((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func FloatMinmax(data []float32, stride int, n int) (float32, float32) {
   var _outptr_0 C.float
   var _outptr_1 C.float
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   C.gsl_stats_float_minmax(&_outptr_0, &_outptr_1, (*C.float)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n))
   return *(*float32)(unsafe.Pointer(&_outptr_0)), *(*float32)(unsafe.Pointer(&_outptr_1))
}

func FloatMaxIndex(data []float32, stride int, n int) int {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return int(C.gsl_stats_float_max_index((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func FloatMinIndex(data []float32, stride int, n int) int {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return int(C.gsl_stats_float_min_index((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func FloatMinmaxIndex(data []float32, stride int, n int) (int, int) {
   var _outptr_0 C.size_t
   var _outptr_1 C.size_t
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   C.gsl_stats_float_minmax_index(&_outptr_0, &_outptr_1, (*C.float)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n))
   return *(*int)(unsafe.Pointer(&_outptr_0)), *(*int)(unsafe.Pointer(&_outptr_1))
}

func FloatMedianFromSortedData(sortedData []float32, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&sortedData))
   return float64(C.gsl_stats_float_median_from_sorted_data((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func FloatQuantileFromSortedData(sortedData []float32, stride int, n int, f float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&sortedData))
   return float64(C.gsl_stats_float_quantile_from_sorted_data((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(f)))
}

func IntMean(data []int32, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_int_mean((*C.int)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func IntVariance(data []int32, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_int_variance((*C.int)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func IntVarianceM(data []int32, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_int_variance_m((*C.int)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func IntSd(data []int32, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_int_sd((*C.int)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func IntSdM(data []int32, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_int_sd_m((*C.int)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func IntTss(data []int32, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_int_tss((*C.int)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func IntTssM(data []int32, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_int_tss_m((*C.int)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func IntVarianceWithFixedMean(data []int32, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_int_variance_with_fixed_mean((*C.int)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func IntSdWithFixedMean(data []int32, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_int_sd_with_fixed_mean((*C.int)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func IntAbsdev(data []int32, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_int_absdev((*C.int)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func IntAbsdevM(data []int32, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_int_absdev_m((*C.int)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func IntSkew(data []int32, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_int_skew((*C.int)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func IntSkewMSd(data []int32, stride int, n int, mean float64, sd float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_int_skew_m_sd((*C.int)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean), C.double(sd)))
}

func IntKurtosis(data []int32, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_int_kurtosis((*C.int)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func IntKurtosisMSd(data []int32, stride int, n int, mean float64, sd float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_int_kurtosis_m_sd((*C.int)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean), C.double(sd)))
}

func IntLag1Autocorrelation(data []int32, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_int_lag1_autocorrelation((*C.int)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func IntLag1AutocorrelationM(data []int32, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_int_lag1_autocorrelation_m((*C.int)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func IntCovariance(data1 []int32, stride1 int, data2 []int32, stride2 int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   return float64(C.gsl_stats_int_covariance((*C.int)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.int)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n)))
}

func IntCovarianceM(data1 []int32, stride1 int, data2 []int32, stride2 int, n int, mean1 float64, mean2 float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   return float64(C.gsl_stats_int_covariance_m((*C.int)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.int)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n), C.double(mean1), C.double(mean2)))
}

func IntCorrelation(data1 []int32, stride1 int, data2 []int32, stride2 int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   return float64(C.gsl_stats_int_correlation((*C.int)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.int)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n)))
}

func IntSpearman(data1 []int32, stride1 int, data2 []int32, stride2 int, n int, work []float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   _slice_header_5 := (*reflect.SliceHeader)(unsafe.Pointer(&work))
   return float64(C.gsl_stats_int_spearman((*C.int)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.int)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n), (*C.double)(unsafe.Pointer(_slice_header_5.Data))))
}

func IntMax(data []int32, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_int_max((*C.int)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func IntMin(data []int32, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_int_min((*C.int)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func IntMinmax(data []int32, stride int, n int) (int32, int32) {
   var _outptr_0 C.int
   var _outptr_1 C.int
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   C.gsl_stats_int_minmax(&_outptr_0, &_outptr_1, (*C.int)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n))
   return *(*int32)(unsafe.Pointer(&_outptr_0)), *(*int32)(unsafe.Pointer(&_outptr_1))
}

func IntMaxIndex(data []int32, stride int, n int) int {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return int(C.gsl_stats_int_max_index((*C.int)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func IntMinIndex(data []int32, stride int, n int) int {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return int(C.gsl_stats_int_min_index((*C.int)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func IntMinmaxIndex(data []int32, stride int, n int) (int, int) {
   var _outptr_0 C.size_t
   var _outptr_1 C.size_t
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   C.gsl_stats_int_minmax_index(&_outptr_0, &_outptr_1, (*C.int)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n))
   return *(*int)(unsafe.Pointer(&_outptr_0)), *(*int)(unsafe.Pointer(&_outptr_1))
}

func IntMedianFromSortedData(sortedData []int32, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&sortedData))
   return float64(C.gsl_stats_int_median_from_sorted_data((*C.int)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func IntQuantileFromSortedData(sortedData []int32, stride int, n int, f float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&sortedData))
   return float64(C.gsl_stats_int_quantile_from_sorted_data((*C.int)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(f)))
}

func UintMean(data []uint32, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_uint_mean((*C.uint)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UintVariance(data []uint32, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_uint_variance((*C.uint)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UintVarianceM(data []uint32, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_uint_variance_m((*C.uint)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func UintSd(data []uint32, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_uint_sd((*C.uint)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UintSdM(data []uint32, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_uint_sd_m((*C.uint)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func UintTss(data []uint32, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_uint_tss((*C.uint)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UintTssM(data []uint32, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_uint_tss_m((*C.uint)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func UintVarianceWithFixedMean(data []uint32, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_uint_variance_with_fixed_mean((*C.uint)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func UintSdWithFixedMean(data []uint32, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_uint_sd_with_fixed_mean((*C.uint)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func UintAbsdev(data []uint32, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_uint_absdev((*C.uint)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UintAbsdevM(data []uint32, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_uint_absdev_m((*C.uint)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func UintSkew(data []uint32, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_uint_skew((*C.uint)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UintSkewMSd(data []uint32, stride int, n int, mean float64, sd float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_uint_skew_m_sd((*C.uint)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean), C.double(sd)))
}

func UintKurtosis(data []uint32, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_uint_kurtosis((*C.uint)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UintKurtosisMSd(data []uint32, stride int, n int, mean float64, sd float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_uint_kurtosis_m_sd((*C.uint)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean), C.double(sd)))
}

func UintLag1Autocorrelation(data []uint32, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_uint_lag1_autocorrelation((*C.uint)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UintLag1AutocorrelationM(data []uint32, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_uint_lag1_autocorrelation_m((*C.uint)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func UintCovariance(data1 []uint32, stride1 int, data2 []uint32, stride2 int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   return float64(C.gsl_stats_uint_covariance((*C.uint)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.uint)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n)))
}

func UintCovarianceM(data1 []uint32, stride1 int, data2 []uint32, stride2 int, n int, mean1 float64, mean2 float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   return float64(C.gsl_stats_uint_covariance_m((*C.uint)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.uint)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n), C.double(mean1), C.double(mean2)))
}

func UintCorrelation(data1 []uint32, stride1 int, data2 []uint32, stride2 int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   return float64(C.gsl_stats_uint_correlation((*C.uint)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.uint)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n)))
}

func UintSpearman(data1 []uint32, stride1 int, data2 []uint32, stride2 int, n int, work []float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   _slice_header_5 := (*reflect.SliceHeader)(unsafe.Pointer(&work))
   return float64(C.gsl_stats_uint_spearman((*C.uint)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.uint)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n), (*C.double)(unsafe.Pointer(_slice_header_5.Data))))
}

func UintMax(data []uint32, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_uint_max((*C.uint)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UintMin(data []uint32, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_uint_min((*C.uint)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UintMinmax(data []uint32, stride int, n int) (uint32, uint32) {
   var _outptr_0 C.uint
   var _outptr_1 C.uint
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   C.gsl_stats_uint_minmax(&_outptr_0, &_outptr_1, (*C.uint)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n))
   return *(*uint32)(unsafe.Pointer(&_outptr_0)), *(*uint32)(unsafe.Pointer(&_outptr_1))
}

func UintMaxIndex(data []uint32, stride int, n int) int {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return int(C.gsl_stats_uint_max_index((*C.uint)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UintMinIndex(data []uint32, stride int, n int) int {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return int(C.gsl_stats_uint_min_index((*C.uint)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UintMinmaxIndex(data []uint32, stride int, n int) (int, int) {
   var _outptr_0 C.size_t
   var _outptr_1 C.size_t
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   C.gsl_stats_uint_minmax_index(&_outptr_0, &_outptr_1, (*C.uint)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n))
   return *(*int)(unsafe.Pointer(&_outptr_0)), *(*int)(unsafe.Pointer(&_outptr_1))
}

func UintMedianFromSortedData(sortedData []uint32, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&sortedData))
   return float64(C.gsl_stats_uint_median_from_sorted_data((*C.uint)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UintQuantileFromSortedData(sortedData []uint32, stride int, n int, f float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&sortedData))
   return float64(C.gsl_stats_uint_quantile_from_sorted_data((*C.uint)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(f)))
}

func LongMean(data []uint, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_long_mean((*C.long)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func LongVariance(data []uint, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_long_variance((*C.long)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func LongVarianceM(data []uint, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_long_variance_m((*C.long)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func LongSd(data []uint, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_long_sd((*C.long)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func LongSdM(data []uint, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_long_sd_m((*C.long)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func LongTss(data []uint, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_long_tss((*C.long)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func LongTssM(data []uint, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_long_tss_m((*C.long)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func LongVarianceWithFixedMean(data []uint, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_long_variance_with_fixed_mean((*C.long)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func LongSdWithFixedMean(data []uint, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_long_sd_with_fixed_mean((*C.long)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func LongAbsdev(data []uint, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_long_absdev((*C.long)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func LongAbsdevM(data []uint, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_long_absdev_m((*C.long)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func LongSkew(data []uint, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_long_skew((*C.long)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func LongSkewMSd(data []uint, stride int, n int, mean float64, sd float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_long_skew_m_sd((*C.long)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean), C.double(sd)))
}

func LongKurtosis(data []uint, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_long_kurtosis((*C.long)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func LongKurtosisMSd(data []uint, stride int, n int, mean float64, sd float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_long_kurtosis_m_sd((*C.long)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean), C.double(sd)))
}

func LongLag1Autocorrelation(data []uint, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_long_lag1_autocorrelation((*C.long)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func LongLag1AutocorrelationM(data []uint, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_long_lag1_autocorrelation_m((*C.long)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func LongCovariance(data1 []uint, stride1 int, data2 []uint, stride2 int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   return float64(C.gsl_stats_long_covariance((*C.long)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.long)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n)))
}

func LongCovarianceM(data1 []uint, stride1 int, data2 []uint, stride2 int, n int, mean1 float64, mean2 float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   return float64(C.gsl_stats_long_covariance_m((*C.long)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.long)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n), C.double(mean1), C.double(mean2)))
}

func LongCorrelation(data1 []uint, stride1 int, data2 []uint, stride2 int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   return float64(C.gsl_stats_long_correlation((*C.long)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.long)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n)))
}

func LongSpearman(data1 []uint, stride1 int, data2 []uint, stride2 int, n int, work []float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   _slice_header_5 := (*reflect.SliceHeader)(unsafe.Pointer(&work))
   return float64(C.gsl_stats_long_spearman((*C.long)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.long)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n), (*C.double)(unsafe.Pointer(_slice_header_5.Data))))
}

func LongMax(data []uint, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_long_max((*C.long)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func LongMin(data []uint, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_long_min((*C.long)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func LongMinmax(data []uint, stride int, n int) (uint, uint) {
   var _outptr_0 C.long
   var _outptr_1 C.long
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   C.gsl_stats_long_minmax(&_outptr_0, &_outptr_1, (*C.long)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n))
   return *(*uint)(unsafe.Pointer(&_outptr_0)), *(*uint)(unsafe.Pointer(&_outptr_1))
}

func LongMaxIndex(data []uint, stride int, n int) int {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return int(C.gsl_stats_long_max_index((*C.long)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func LongMinIndex(data []uint, stride int, n int) int {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return int(C.gsl_stats_long_min_index((*C.long)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func LongMinmaxIndex(data []uint, stride int, n int) (int, int) {
   var _outptr_0 C.size_t
   var _outptr_1 C.size_t
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   C.gsl_stats_long_minmax_index(&_outptr_0, &_outptr_1, (*C.long)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n))
   return *(*int)(unsafe.Pointer(&_outptr_0)), *(*int)(unsafe.Pointer(&_outptr_1))
}

func LongMedianFromSortedData(sortedData []uint, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&sortedData))
   return float64(C.gsl_stats_long_median_from_sorted_data((*C.long)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func LongQuantileFromSortedData(sortedData []uint, stride int, n int, f float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&sortedData))
   return float64(C.gsl_stats_long_quantile_from_sorted_data((*C.long)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(f)))
}

func UlongMean(data []int, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_ulong_mean((*C.ulong)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UlongVariance(data []int, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_ulong_variance((*C.ulong)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UlongVarianceM(data []int, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_ulong_variance_m((*C.ulong)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func UlongSd(data []int, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_ulong_sd((*C.ulong)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UlongSdM(data []int, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_ulong_sd_m((*C.ulong)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func UlongTss(data []int, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_ulong_tss((*C.ulong)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UlongTssM(data []int, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_ulong_tss_m((*C.ulong)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func UlongVarianceWithFixedMean(data []int, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_ulong_variance_with_fixed_mean((*C.ulong)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func UlongSdWithFixedMean(data []int, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_ulong_sd_with_fixed_mean((*C.ulong)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func UlongAbsdev(data []int, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_ulong_absdev((*C.ulong)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UlongAbsdevM(data []int, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_ulong_absdev_m((*C.ulong)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func UlongSkew(data []int, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_ulong_skew((*C.ulong)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UlongSkewMSd(data []int, stride int, n int, mean float64, sd float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_ulong_skew_m_sd((*C.ulong)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean), C.double(sd)))
}

func UlongKurtosis(data []int, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_ulong_kurtosis((*C.ulong)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UlongKurtosisMSd(data []int, stride int, n int, mean float64, sd float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_ulong_kurtosis_m_sd((*C.ulong)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean), C.double(sd)))
}

func UlongLag1Autocorrelation(data []int, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_ulong_lag1_autocorrelation((*C.ulong)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UlongLag1AutocorrelationM(data []int, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_ulong_lag1_autocorrelation_m((*C.ulong)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func UlongCovariance(data1 []int, stride1 int, data2 []int, stride2 int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   return float64(C.gsl_stats_ulong_covariance((*C.ulong)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.ulong)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n)))
}

func UlongCovarianceM(data1 []int, stride1 int, data2 []int, stride2 int, n int, mean1 float64, mean2 float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   return float64(C.gsl_stats_ulong_covariance_m((*C.ulong)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.ulong)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n), C.double(mean1), C.double(mean2)))
}

func UlongCorrelation(data1 []int, stride1 int, data2 []int, stride2 int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   return float64(C.gsl_stats_ulong_correlation((*C.ulong)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.ulong)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n)))
}

func UlongSpearman(data1 []int, stride1 int, data2 []int, stride2 int, n int, work []float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   _slice_header_5 := (*reflect.SliceHeader)(unsafe.Pointer(&work))
   return float64(C.gsl_stats_ulong_spearman((*C.ulong)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.ulong)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n), (*C.double)(unsafe.Pointer(_slice_header_5.Data))))
}

func UlongMax(data []int, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_ulong_max((*C.ulong)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UlongMin(data []int, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_ulong_min((*C.ulong)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UlongMinmax(data []int, stride int, n int) (int, int) {
   var _outptr_0 C.ulong
   var _outptr_1 C.ulong
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   C.gsl_stats_ulong_minmax(&_outptr_0, &_outptr_1, (*C.ulong)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n))
   return *(*int)(unsafe.Pointer(&_outptr_0)), *(*int)(unsafe.Pointer(&_outptr_1))
}

func UlongMaxIndex(data []int, stride int, n int) int {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return int(C.gsl_stats_ulong_max_index((*C.ulong)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UlongMinIndex(data []int, stride int, n int) int {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return int(C.gsl_stats_ulong_min_index((*C.ulong)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UlongMinmaxIndex(data []int, stride int, n int) (int, int) {
   var _outptr_0 C.size_t
   var _outptr_1 C.size_t
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   C.gsl_stats_ulong_minmax_index(&_outptr_0, &_outptr_1, (*C.ulong)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n))
   return *(*int)(unsafe.Pointer(&_outptr_0)), *(*int)(unsafe.Pointer(&_outptr_1))
}

func UlongMedianFromSortedData(sortedData []int, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&sortedData))
   return float64(C.gsl_stats_ulong_median_from_sorted_data((*C.ulong)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UlongQuantileFromSortedData(sortedData []int, stride int, n int, f float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&sortedData))
   return float64(C.gsl_stats_ulong_quantile_from_sorted_data((*C.ulong)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(f)))
}

func ShortMean(data []int16, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_short_mean((*C.short)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func ShortVariance(data []int16, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_short_variance((*C.short)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func ShortVarianceM(data []int16, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_short_variance_m((*C.short)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func ShortSd(data []int16, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_short_sd((*C.short)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func ShortSdM(data []int16, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_short_sd_m((*C.short)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func ShortTss(data []int16, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_short_tss((*C.short)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func ShortTssM(data []int16, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_short_tss_m((*C.short)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func ShortVarianceWithFixedMean(data []int16, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_short_variance_with_fixed_mean((*C.short)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func ShortSdWithFixedMean(data []int16, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_short_sd_with_fixed_mean((*C.short)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func ShortAbsdev(data []int16, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_short_absdev((*C.short)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func ShortAbsdevM(data []int16, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_short_absdev_m((*C.short)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func ShortSkew(data []int16, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_short_skew((*C.short)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func ShortSkewMSd(data []int16, stride int, n int, mean float64, sd float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_short_skew_m_sd((*C.short)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean), C.double(sd)))
}

func ShortKurtosis(data []int16, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_short_kurtosis((*C.short)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func ShortKurtosisMSd(data []int16, stride int, n int, mean float64, sd float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_short_kurtosis_m_sd((*C.short)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean), C.double(sd)))
}

func ShortLag1Autocorrelation(data []int16, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_short_lag1_autocorrelation((*C.short)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func ShortLag1AutocorrelationM(data []int16, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_short_lag1_autocorrelation_m((*C.short)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func ShortCovariance(data1 []int16, stride1 int, data2 []int16, stride2 int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   return float64(C.gsl_stats_short_covariance((*C.short)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.short)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n)))
}

func ShortCovarianceM(data1 []int16, stride1 int, data2 []int16, stride2 int, n int, mean1 float64, mean2 float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   return float64(C.gsl_stats_short_covariance_m((*C.short)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.short)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n), C.double(mean1), C.double(mean2)))
}

func ShortCorrelation(data1 []int16, stride1 int, data2 []int16, stride2 int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   return float64(C.gsl_stats_short_correlation((*C.short)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.short)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n)))
}

func ShortSpearman(data1 []int16, stride1 int, data2 []int16, stride2 int, n int, work []float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   _slice_header_5 := (*reflect.SliceHeader)(unsafe.Pointer(&work))
   return float64(C.gsl_stats_short_spearman((*C.short)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.short)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n), (*C.double)(unsafe.Pointer(_slice_header_5.Data))))
}

func ShortMax(data []int16, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_short_max((*C.short)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func ShortMin(data []int16, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_short_min((*C.short)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func ShortMinmax(data []int16, stride int, n int) (int16, int16) {
   var _outptr_0 C.short
   var _outptr_1 C.short
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   C.gsl_stats_short_minmax(&_outptr_0, &_outptr_1, (*C.short)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n))
   return *(*int16)(unsafe.Pointer(&_outptr_0)), *(*int16)(unsafe.Pointer(&_outptr_1))
}

func ShortMaxIndex(data []int16, stride int, n int) int {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return int(C.gsl_stats_short_max_index((*C.short)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func ShortMinIndex(data []int16, stride int, n int) int {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return int(C.gsl_stats_short_min_index((*C.short)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func ShortMinmaxIndex(data []int16, stride int, n int) (int, int) {
   var _outptr_0 C.size_t
   var _outptr_1 C.size_t
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   C.gsl_stats_short_minmax_index(&_outptr_0, &_outptr_1, (*C.short)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n))
   return *(*int)(unsafe.Pointer(&_outptr_0)), *(*int)(unsafe.Pointer(&_outptr_1))
}

func ShortMedianFromSortedData(sortedData []int16, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&sortedData))
   return float64(C.gsl_stats_short_median_from_sorted_data((*C.short)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func ShortQuantileFromSortedData(sortedData []int16, stride int, n int, f float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&sortedData))
   return float64(C.gsl_stats_short_quantile_from_sorted_data((*C.short)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(f)))
}

func UshortMean(data []int16, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_ushort_mean((*C.ushort)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UshortVariance(data []int16, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_ushort_variance((*C.ushort)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UshortVarianceM(data []int16, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_ushort_variance_m((*C.ushort)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func UshortSd(data []int16, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_ushort_sd((*C.ushort)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UshortSdM(data []int16, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_ushort_sd_m((*C.ushort)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func UshortTss(data []int16, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_ushort_tss((*C.ushort)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UshortTssM(data []int16, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_ushort_tss_m((*C.ushort)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func UshortVarianceWithFixedMean(data []int16, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_ushort_variance_with_fixed_mean((*C.ushort)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func UshortSdWithFixedMean(data []int16, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_ushort_sd_with_fixed_mean((*C.ushort)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func UshortAbsdev(data []int16, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_ushort_absdev((*C.ushort)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UshortAbsdevM(data []int16, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_ushort_absdev_m((*C.ushort)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func UshortSkew(data []int16, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_ushort_skew((*C.ushort)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UshortSkewMSd(data []int16, stride int, n int, mean float64, sd float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_ushort_skew_m_sd((*C.ushort)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean), C.double(sd)))
}

func UshortKurtosis(data []int16, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_ushort_kurtosis((*C.ushort)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UshortKurtosisMSd(data []int16, stride int, n int, mean float64, sd float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_ushort_kurtosis_m_sd((*C.ushort)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean), C.double(sd)))
}

func UshortLag1Autocorrelation(data []int16, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_ushort_lag1_autocorrelation((*C.ushort)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UshortLag1AutocorrelationM(data []int16, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_ushort_lag1_autocorrelation_m((*C.ushort)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func UshortCovariance(data1 []int16, stride1 int, data2 []int16, stride2 int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   return float64(C.gsl_stats_ushort_covariance((*C.ushort)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.ushort)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n)))
}

func UshortCovarianceM(data1 []int16, stride1 int, data2 []int16, stride2 int, n int, mean1 float64, mean2 float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   return float64(C.gsl_stats_ushort_covariance_m((*C.ushort)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.ushort)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n), C.double(mean1), C.double(mean2)))
}

func UshortCorrelation(data1 []int16, stride1 int, data2 []int16, stride2 int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   return float64(C.gsl_stats_ushort_correlation((*C.ushort)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.ushort)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n)))
}

func UshortSpearman(data1 []int16, stride1 int, data2 []int16, stride2 int, n int, work []float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   _slice_header_5 := (*reflect.SliceHeader)(unsafe.Pointer(&work))
   return float64(C.gsl_stats_ushort_spearman((*C.ushort)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.ushort)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n), (*C.double)(unsafe.Pointer(_slice_header_5.Data))))
}

func UshortMax(data []int16, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_ushort_max((*C.ushort)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UshortMin(data []int16, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_ushort_min((*C.ushort)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UshortMinmax(data []int16, stride int, n int) (int16, int16) {
   var _outptr_0 C.ushort
   var _outptr_1 C.ushort
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   C.gsl_stats_ushort_minmax(&_outptr_0, &_outptr_1, (*C.ushort)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n))
   return *(*int16)(unsafe.Pointer(&_outptr_0)), *(*int16)(unsafe.Pointer(&_outptr_1))
}

func UshortMaxIndex(data []int16, stride int, n int) int {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return int(C.gsl_stats_ushort_max_index((*C.ushort)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UshortMinIndex(data []int16, stride int, n int) int {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return int(C.gsl_stats_ushort_min_index((*C.ushort)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UshortMinmaxIndex(data []int16, stride int, n int) (int, int) {
   var _outptr_0 C.size_t
   var _outptr_1 C.size_t
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   C.gsl_stats_ushort_minmax_index(&_outptr_0, &_outptr_1, (*C.ushort)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n))
   return *(*int)(unsafe.Pointer(&_outptr_0)), *(*int)(unsafe.Pointer(&_outptr_1))
}

func UshortMedianFromSortedData(sortedData []int16, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&sortedData))
   return float64(C.gsl_stats_ushort_median_from_sorted_data((*C.ushort)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UshortQuantileFromSortedData(sortedData []int16, stride int, n int, f float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&sortedData))
   return float64(C.gsl_stats_ushort_quantile_from_sorted_data((*C.ushort)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(f)))
}

func CharMean(data []int8, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_char_mean((*C.char)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func CharVariance(data []int8, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_char_variance((*C.char)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func CharVarianceM(data []int8, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_char_variance_m((*C.char)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func CharSd(data []int8, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_char_sd((*C.char)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func CharSdM(data []int8, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_char_sd_m((*C.char)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func CharTss(data []int8, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_char_tss((*C.char)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func CharTssM(data []int8, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_char_tss_m((*C.char)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func CharVarianceWithFixedMean(data []int8, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_char_variance_with_fixed_mean((*C.char)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func CharSdWithFixedMean(data []int8, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_char_sd_with_fixed_mean((*C.char)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func CharAbsdev(data []int8, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_char_absdev((*C.char)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func CharAbsdevM(data []int8, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_char_absdev_m((*C.char)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func CharSkew(data []int8, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_char_skew((*C.char)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func CharSkewMSd(data []int8, stride int, n int, mean float64, sd float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_char_skew_m_sd((*C.char)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean), C.double(sd)))
}

func CharKurtosis(data []int8, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_char_kurtosis((*C.char)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func CharKurtosisMSd(data []int8, stride int, n int, mean float64, sd float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_char_kurtosis_m_sd((*C.char)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean), C.double(sd)))
}

func CharLag1Autocorrelation(data []int8, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_char_lag1_autocorrelation((*C.char)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func CharLag1AutocorrelationM(data []int8, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_char_lag1_autocorrelation_m((*C.char)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func CharCovariance(data1 []int8, stride1 int, data2 []int8, stride2 int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   return float64(C.gsl_stats_char_covariance((*C.char)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.char)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n)))
}

func CharCovarianceM(data1 []int8, stride1 int, data2 []int8, stride2 int, n int, mean1 float64, mean2 float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   return float64(C.gsl_stats_char_covariance_m((*C.char)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.char)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n), C.double(mean1), C.double(mean2)))
}

func CharCorrelation(data1 []int8, stride1 int, data2 []int8, stride2 int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   return float64(C.gsl_stats_char_correlation((*C.char)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.char)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n)))
}

func CharSpearman(data1 []int8, stride1 int, data2 []int8, stride2 int, n int, work []float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   _slice_header_5 := (*reflect.SliceHeader)(unsafe.Pointer(&work))
   return float64(C.gsl_stats_char_spearman((*C.char)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.char)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n), (*C.double)(unsafe.Pointer(_slice_header_5.Data))))
}

func CharMax(data []int8, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_char_max((*C.char)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func CharMin(data []int8, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_char_min((*C.char)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func CharMinmax(data []int8, stride int, n int) (int8, int8) {
   var _outptr_0 C.char
   var _outptr_1 C.char
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   C.gsl_stats_char_minmax(&_outptr_0, &_outptr_1, (*C.char)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n))
   return *(*int8)(unsafe.Pointer(&_outptr_0)), *(*int8)(unsafe.Pointer(&_outptr_1))
}

func CharMaxIndex(data []int8, stride int, n int) int {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return int(C.gsl_stats_char_max_index((*C.char)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func CharMinIndex(data []int8, stride int, n int) int {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return int(C.gsl_stats_char_min_index((*C.char)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func CharMinmaxIndex(data []int8, stride int, n int) (int, int) {
   var _outptr_0 C.size_t
   var _outptr_1 C.size_t
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   C.gsl_stats_char_minmax_index(&_outptr_0, &_outptr_1, (*C.char)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n))
   return *(*int)(unsafe.Pointer(&_outptr_0)), *(*int)(unsafe.Pointer(&_outptr_1))
}

func CharMedianFromSortedData(sortedData []int8, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&sortedData))
   return float64(C.gsl_stats_char_median_from_sorted_data((*C.char)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func CharQuantileFromSortedData(sortedData []int8, stride int, n int, f float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&sortedData))
   return float64(C.gsl_stats_char_quantile_from_sorted_data((*C.char)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(f)))
}

func UcharMean(data []uint8, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_uchar_mean((*C.uchar)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UcharVariance(data []uint8, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_uchar_variance((*C.uchar)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UcharVarianceM(data []uint8, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_uchar_variance_m((*C.uchar)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func UcharSd(data []uint8, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_uchar_sd((*C.uchar)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UcharSdM(data []uint8, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_uchar_sd_m((*C.uchar)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func UcharTss(data []uint8, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_uchar_tss((*C.uchar)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UcharTssM(data []uint8, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_uchar_tss_m((*C.uchar)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func UcharVarianceWithFixedMean(data []uint8, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_uchar_variance_with_fixed_mean((*C.uchar)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func UcharSdWithFixedMean(data []uint8, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_uchar_sd_with_fixed_mean((*C.uchar)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func UcharAbsdev(data []uint8, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_uchar_absdev((*C.uchar)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UcharAbsdevM(data []uint8, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_uchar_absdev_m((*C.uchar)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func UcharSkew(data []uint8, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_uchar_skew((*C.uchar)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UcharSkewMSd(data []uint8, stride int, n int, mean float64, sd float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_uchar_skew_m_sd((*C.uchar)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean), C.double(sd)))
}

func UcharKurtosis(data []uint8, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_uchar_kurtosis((*C.uchar)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UcharKurtosisMSd(data []uint8, stride int, n int, mean float64, sd float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_uchar_kurtosis_m_sd((*C.uchar)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean), C.double(sd)))
}

func UcharLag1Autocorrelation(data []uint8, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_uchar_lag1_autocorrelation((*C.uchar)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UcharLag1AutocorrelationM(data []uint8, stride int, n int, mean float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_uchar_lag1_autocorrelation_m((*C.uchar)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(mean)))
}

func UcharCovariance(data1 []uint8, stride1 int, data2 []uint8, stride2 int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   return float64(C.gsl_stats_uchar_covariance((*C.uchar)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.uchar)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n)))
}

func UcharCovarianceM(data1 []uint8, stride1 int, data2 []uint8, stride2 int, n int, mean1 float64, mean2 float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   return float64(C.gsl_stats_uchar_covariance_m((*C.uchar)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.uchar)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n), C.double(mean1), C.double(mean2)))
}

func UcharCorrelation(data1 []uint8, stride1 int, data2 []uint8, stride2 int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   return float64(C.gsl_stats_uchar_correlation((*C.uchar)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.uchar)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n)))
}

func UcharSpearman(data1 []uint8, stride1 int, data2 []uint8, stride2 int, n int, work []float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   _slice_header_5 := (*reflect.SliceHeader)(unsafe.Pointer(&work))
   return float64(C.gsl_stats_uchar_spearman((*C.uchar)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.uchar)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n), (*C.double)(unsafe.Pointer(_slice_header_5.Data))))
}

func UcharMax(data []uint8, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_uchar_max((*C.uchar)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UcharMin(data []uint8, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return float64(C.gsl_stats_uchar_min((*C.uchar)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UcharMinmax(data []uint8, stride int, n int) (uint8, uint8) {
   var _outptr_0 C.uchar
   var _outptr_1 C.uchar
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   C.gsl_stats_uchar_minmax(&_outptr_0, &_outptr_1, (*C.uchar)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n))
   return *(*uint8)(unsafe.Pointer(&_outptr_0)), *(*uint8)(unsafe.Pointer(&_outptr_1))
}

func UcharMaxIndex(data []uint8, stride int, n int) int {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return int(C.gsl_stats_uchar_max_index((*C.uchar)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UcharMinIndex(data []uint8, stride int, n int) int {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   return int(C.gsl_stats_uchar_min_index((*C.uchar)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UcharMinmaxIndex(data []uint8, stride int, n int) (int, int) {
   var _outptr_0 C.size_t
   var _outptr_1 C.size_t
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   C.gsl_stats_uchar_minmax_index(&_outptr_0, &_outptr_1, (*C.uchar)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n))
   return *(*int)(unsafe.Pointer(&_outptr_0)), *(*int)(unsafe.Pointer(&_outptr_1))
}

func UcharMedianFromSortedData(sortedData []uint8, stride int, n int) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&sortedData))
   return float64(C.gsl_stats_uchar_median_from_sorted_data((*C.uchar)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n)))
}

func UcharQuantileFromSortedData(sortedData []uint8, stride int, n int, f float64) float64 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&sortedData))
   return float64(C.gsl_stats_uchar_quantile_from_sorted_data((*C.uchar)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n), C.double(f)))
}

