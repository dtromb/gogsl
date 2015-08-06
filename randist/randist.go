//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package randist

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_randist.h>
   #include <gsl/gsl_cdf.h>
*/
import "C"

import "unsafe"
import "github.com/dtromb/gogsl"
import "github.com/dtromb/gogsl/rng"
import "reflect"

type GslRanDiscreteT struct {
   gogsl.GslReference
}

func Gaussian(r *rng.GslRng, sigma float64) float64 {
   return float64(C.gsl_ran_gaussian((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), C.double(sigma)))
}

func GaussianPdf(x float64, sigma float64) float64 {
   return float64(C.gsl_ran_gaussian_pdf(C.double(x), C.double(sigma)))
}

func GaussianZiggurat(r *rng.GslRng, sigma float64) float64 {
   return float64(C.gsl_ran_gaussian_ziggurat((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), C.double(sigma)))
}

func GaussianRatioMethod(r *rng.GslRng, sigma float64) float64 {
   return float64(C.gsl_ran_gaussian_ratio_method((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), C.double(sigma)))
}

func Ugaussian(r *rng.GslRng) float64 {
   return float64(C.gsl_ran_ugaussian((*C.gsl_rng)(unsafe.Pointer(r.Ptr()))))
}

func UgaussianPdf(x float64) float64 {
   return float64(C.gsl_ran_ugaussian_pdf(C.double(x)))
}

func UgaussianRatioMethod(r *rng.GslRng) float64 {
   return float64(C.gsl_ran_ugaussian_ratio_method((*C.gsl_rng)(unsafe.Pointer(r.Ptr()))))
}

func CdfGaussianP(x float64, sigma float64) float64 {
   return float64(C.gsl_cdf_gaussian_P(C.double(x), C.double(sigma)))
}

func CdfGaussianQ(x float64, sigma float64) float64 {
   return float64(C.gsl_cdf_gaussian_Q(C.double(x), C.double(sigma)))
}

func CdfGaussianPinv(p float64, sigma float64) float64 {
   return float64(C.gsl_cdf_gaussian_Pinv(C.double(p), C.double(sigma)))
}

func CdfGaussianQinv(q float64, sigma float64) float64 {
   return float64(C.gsl_cdf_gaussian_Qinv(C.double(q), C.double(sigma)))
}

func CdfUgaussianP(x float64) float64 {
   return float64(C.gsl_cdf_ugaussian_P(C.double(x)))
}

func CdfUgaussianQ(x float64) float64 {
   return float64(C.gsl_cdf_ugaussian_Q(C.double(x)))
}

func CdfUgaussianPinv(p float64) float64 {
   return float64(C.gsl_cdf_ugaussian_Pinv(C.double(p)))
}

func CdfUgaussianQinv(q float64) float64 {
   return float64(C.gsl_cdf_ugaussian_Qinv(C.double(q)))
}

func GaussianTail(r *rng.GslRng, a float64, sigma float64) float64 {
   return float64(C.gsl_ran_gaussian_tail((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), C.double(a), C.double(sigma)))
}

func GaussianTailPdf(x float64, a float64, sigma float64) float64 {
   return float64(C.gsl_ran_gaussian_tail_pdf(C.double(x), C.double(a), C.double(sigma)))
}

func UgaussianTail(r *rng.GslRng, a float64) float64 {
   return float64(C.gsl_ran_ugaussian_tail((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), C.double(a)))
}

func UgaussianTailPdf(x float64, a float64) float64 {
   return float64(C.gsl_ran_ugaussian_tail_pdf(C.double(x), C.double(a)))
}

func BivariateGaussian(r *rng.GslRng, sigmaX float64, sigmaY float64, rho float64) (float64, float64) {
   var _outptr_4 C.double
   var _outptr_5 C.double
   C.gsl_ran_bivariate_gaussian((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), C.double(sigmaX), C.double(sigmaY), C.double(rho), &_outptr_4, &_outptr_5)
   return *(*float64)(unsafe.Pointer(&_outptr_4)), *(*float64)(unsafe.Pointer(&_outptr_5))
}

func BivariateGaussianPdf(x float64, y float64, sigmaX float64, sigmaY float64, rho float64) float64 {
   return float64(C.gsl_ran_bivariate_gaussian_pdf(C.double(x), C.double(y), C.double(sigmaX), C.double(sigmaY), C.double(rho)))
}

func Exponential(r *rng.GslRng, mu float64) float64 {
   return float64(C.gsl_ran_exponential((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), C.double(mu)))
}

func ExponentialPdf(x float64, mu float64) float64 {
   return float64(C.gsl_ran_exponential_pdf(C.double(x), C.double(mu)))
}

func CdfExponentialP(x float64, mu float64) float64 {
   return float64(C.gsl_cdf_exponential_P(C.double(x), C.double(mu)))
}

func CdfExponentialQ(x float64, mu float64) float64 {
   return float64(C.gsl_cdf_exponential_Q(C.double(x), C.double(mu)))
}

func CdfExponentialPinv(p float64, mu float64) float64 {
   return float64(C.gsl_cdf_exponential_Pinv(C.double(p), C.double(mu)))
}

func CdfExponentialQinv(q float64, mu float64) float64 {
   return float64(C.gsl_cdf_exponential_Qinv(C.double(q), C.double(mu)))
}

func Laplace(r *rng.GslRng, a float64) float64 {
   return float64(C.gsl_ran_laplace((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), C.double(a)))
}

func LaplacePdf(x float64, a float64) float64 {
   return float64(C.gsl_ran_laplace_pdf(C.double(x), C.double(a)))
}

func CdfLaplaceP(x float64, a float64) float64 {
   return float64(C.gsl_cdf_laplace_P(C.double(x), C.double(a)))
}

func CdfLaplaceQ(x float64, a float64) float64 {
   return float64(C.gsl_cdf_laplace_Q(C.double(x), C.double(a)))
}

func CdfLaplacePinv(p float64, a float64) float64 {
   return float64(C.gsl_cdf_laplace_Pinv(C.double(p), C.double(a)))
}

func CdfLaplaceQinv(q float64, a float64) float64 {
   return float64(C.gsl_cdf_laplace_Qinv(C.double(q), C.double(a)))
}

func Exppow(r *rng.GslRng, a float64, b float64) float64 {
   return float64(C.gsl_ran_exppow((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), C.double(a), C.double(b)))
}

func ExppowPdf(x float64, a float64, b float64) float64 {
   return float64(C.gsl_ran_exppow_pdf(C.double(x), C.double(a), C.double(b)))
}

func CdfExppowP(x float64, a float64, b float64) float64 {
   return float64(C.gsl_cdf_exppow_P(C.double(x), C.double(a), C.double(b)))
}

func CdfExppowQ(x float64, a float64, b float64) float64 {
   return float64(C.gsl_cdf_exppow_Q(C.double(x), C.double(a), C.double(b)))
}

func Cauchy(r *rng.GslRng, a float64) float64 {
   return float64(C.gsl_ran_cauchy((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), C.double(a)))
}

func CauchyPdf(x float64, a float64) float64 {
   return float64(C.gsl_ran_cauchy_pdf(C.double(x), C.double(a)))
}

func CdfCauchyP(x float64, a float64) float64 {
   return float64(C.gsl_cdf_cauchy_P(C.double(x), C.double(a)))
}

func CdfCauchyQ(x float64, a float64) float64 {
   return float64(C.gsl_cdf_cauchy_Q(C.double(x), C.double(a)))
}

func CdfCauchyPinv(p float64, a float64) float64 {
   return float64(C.gsl_cdf_cauchy_Pinv(C.double(p), C.double(a)))
}

func CdfCauchyQinv(q float64, a float64) float64 {
   return float64(C.gsl_cdf_cauchy_Qinv(C.double(q), C.double(a)))
}

func Rayleigh(r *rng.GslRng, sigma float64) float64 {
   return float64(C.gsl_ran_rayleigh((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), C.double(sigma)))
}

func RayleighPdf(x float64, sigma float64) float64 {
   return float64(C.gsl_ran_rayleigh_pdf(C.double(x), C.double(sigma)))
}

func CdfRayleighP(x float64, sigma float64) float64 {
   return float64(C.gsl_cdf_rayleigh_P(C.double(x), C.double(sigma)))
}

func CdfRayleighQ(x float64, sigma float64) float64 {
   return float64(C.gsl_cdf_rayleigh_Q(C.double(x), C.double(sigma)))
}

func CdfRayleighPinv(p float64, sigma float64) float64 {
   return float64(C.gsl_cdf_rayleigh_Pinv(C.double(p), C.double(sigma)))
}

func CdfRayleighQinv(q float64, sigma float64) float64 {
   return float64(C.gsl_cdf_rayleigh_Qinv(C.double(q), C.double(sigma)))
}

func RayleighTail(r *rng.GslRng, a float64, sigma float64) float64 {
   return float64(C.gsl_ran_rayleigh_tail((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), C.double(a), C.double(sigma)))
}

func RayleighTailPdf(x float64, a float64, sigma float64) float64 {
   return float64(C.gsl_ran_rayleigh_tail_pdf(C.double(x), C.double(a), C.double(sigma)))
}

func Landau(r *rng.GslRng) float64 {
   return float64(C.gsl_ran_landau((*C.gsl_rng)(unsafe.Pointer(r.Ptr()))))
}

func LandauPdf(x float64) float64 {
   return float64(C.gsl_ran_landau_pdf(C.double(x)))
}

func Levy(r *rng.GslRng, c float64, alpha float64) float64 {
   return float64(C.gsl_ran_levy((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), C.double(c), C.double(alpha)))
}

func LevySkew(r *rng.GslRng, c float64, alpha float64, beta float64) float64 {
   return float64(C.gsl_ran_levy_skew((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), C.double(c), C.double(alpha), C.double(beta)))
}

func Gamma(r *rng.GslRng, a float64, b float64) float64 {
   return float64(C.gsl_ran_gamma((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), C.double(a), C.double(b)))
}

func GammaKnuth(r *rng.GslRng, a float64, b float64) float64 {
   return float64(C.gsl_ran_gamma_knuth((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), C.double(a), C.double(b)))
}

func GammaPdf(x float64, a float64, b float64) float64 {
   return float64(C.gsl_ran_gamma_pdf(C.double(x), C.double(a), C.double(b)))
}

func CdfGammaP(x float64, a float64, b float64) float64 {
   return float64(C.gsl_cdf_gamma_P(C.double(x), C.double(a), C.double(b)))
}

func CdfGammaQ(x float64, a float64, b float64) float64 {
   return float64(C.gsl_cdf_gamma_Q(C.double(x), C.double(a), C.double(b)))
}

func CdfGammaPinv(p float64, a float64, b float64) float64 {
   return float64(C.gsl_cdf_gamma_Pinv(C.double(p), C.double(a), C.double(b)))
}

func CdfGammaQinv(q float64, a float64, b float64) float64 {
   return float64(C.gsl_cdf_gamma_Qinv(C.double(q), C.double(a), C.double(b)))
}

func Flat(r *rng.GslRng, a float64, b float64) float64 {
   return float64(C.gsl_ran_flat((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), C.double(a), C.double(b)))
}

func FlatPdf(x float64, a float64, b float64) float64 {
   return float64(C.gsl_ran_flat_pdf(C.double(x), C.double(a), C.double(b)))
}

func CdfFlatP(x float64, a float64, b float64) float64 {
   return float64(C.gsl_cdf_flat_P(C.double(x), C.double(a), C.double(b)))
}

func CdfFlatQ(x float64, a float64, b float64) float64 {
   return float64(C.gsl_cdf_flat_Q(C.double(x), C.double(a), C.double(b)))
}

func CdfFlatPinv(p float64, a float64, b float64) float64 {
   return float64(C.gsl_cdf_flat_Pinv(C.double(p), C.double(a), C.double(b)))
}

func CdfFlatQinv(q float64, a float64, b float64) float64 {
   return float64(C.gsl_cdf_flat_Qinv(C.double(q), C.double(a), C.double(b)))
}

func Lognormal(r *rng.GslRng, zeta float64, sigma float64) float64 {
   return float64(C.gsl_ran_lognormal((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), C.double(zeta), C.double(sigma)))
}

func LognormalPdf(x float64, zeta float64, sigma float64) float64 {
   return float64(C.gsl_ran_lognormal_pdf(C.double(x), C.double(zeta), C.double(sigma)))
}

func CdfLognormalP(x float64, zeta float64, sigma float64) float64 {
   return float64(C.gsl_cdf_lognormal_P(C.double(x), C.double(zeta), C.double(sigma)))
}

func CdfLognormalQ(x float64, zeta float64, sigma float64) float64 {
   return float64(C.gsl_cdf_lognormal_Q(C.double(x), C.double(zeta), C.double(sigma)))
}

func CdfLognormalPinv(p float64, zeta float64, sigma float64) float64 {
   return float64(C.gsl_cdf_lognormal_Pinv(C.double(p), C.double(zeta), C.double(sigma)))
}

func CdfLognormalQinv(q float64, zeta float64, sigma float64) float64 {
   return float64(C.gsl_cdf_lognormal_Qinv(C.double(q), C.double(zeta), C.double(sigma)))
}

func Chisq(r *rng.GslRng, nu float64) float64 {
   return float64(C.gsl_ran_chisq((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), C.double(nu)))
}

func ChisqPdf(x float64, nu float64) float64 {
   return float64(C.gsl_ran_chisq_pdf(C.double(x), C.double(nu)))
}

func CdfChisqP(x float64, nu float64) float64 {
   return float64(C.gsl_cdf_chisq_P(C.double(x), C.double(nu)))
}

func CdfChisqQ(x float64, nu float64) float64 {
   return float64(C.gsl_cdf_chisq_Q(C.double(x), C.double(nu)))
}

func CdfChisqPinv(p float64, nu float64) float64 {
   return float64(C.gsl_cdf_chisq_Pinv(C.double(p), C.double(nu)))
}

func CdfChisqQinv(q float64, nu float64) float64 {
   return float64(C.gsl_cdf_chisq_Qinv(C.double(q), C.double(nu)))
}

func Fdist(r *rng.GslRng, nu1 float64, nu2 float64) float64 {
   return float64(C.gsl_ran_fdist((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), C.double(nu1), C.double(nu2)))
}

func FdistPdf(x float64, nu1 float64, nu2 float64) float64 {
   return float64(C.gsl_ran_fdist_pdf(C.double(x), C.double(nu1), C.double(nu2)))
}

func CdfFdistP(x float64, nu1 float64, nu2 float64) float64 {
   return float64(C.gsl_cdf_fdist_P(C.double(x), C.double(nu1), C.double(nu2)))
}

func CdfFdistQ(x float64, nu1 float64, nu2 float64) float64 {
   return float64(C.gsl_cdf_fdist_Q(C.double(x), C.double(nu1), C.double(nu2)))
}

func CdfFdistPinv(p float64, nu1 float64, nu2 float64) float64 {
   return float64(C.gsl_cdf_fdist_Pinv(C.double(p), C.double(nu1), C.double(nu2)))
}

func CdfFdistQinv(q float64, nu1 float64, nu2 float64) float64 {
   return float64(C.gsl_cdf_fdist_Qinv(C.double(q), C.double(nu1), C.double(nu2)))
}

func Tdist(r *rng.GslRng, nu float64) float64 {
   return float64(C.gsl_ran_tdist((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), C.double(nu)))
}

func TdistPdf(x float64, nu float64) float64 {
   return float64(C.gsl_ran_tdist_pdf(C.double(x), C.double(nu)))
}

func CdfTdistP(x float64, nu float64) float64 {
   return float64(C.gsl_cdf_tdist_P(C.double(x), C.double(nu)))
}

func CdfTdistQ(x float64, nu float64) float64 {
   return float64(C.gsl_cdf_tdist_Q(C.double(x), C.double(nu)))
}

func CdfTdistPinv(p float64, nu float64) float64 {
   return float64(C.gsl_cdf_tdist_Pinv(C.double(p), C.double(nu)))
}

func CdfTdistQinv(q float64, nu float64) float64 {
   return float64(C.gsl_cdf_tdist_Qinv(C.double(q), C.double(nu)))
}

func Beta(r *rng.GslRng, a float64, b float64) float64 {
   return float64(C.gsl_ran_beta((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), C.double(a), C.double(b)))
}

func BetaPdf(x float64, a float64, b float64) float64 {
   return float64(C.gsl_ran_beta_pdf(C.double(x), C.double(a), C.double(b)))
}

func CdfBetaP(x float64, a float64, b float64) float64 {
   return float64(C.gsl_cdf_beta_P(C.double(x), C.double(a), C.double(b)))
}

func CdfBetaQ(x float64, a float64, b float64) float64 {
   return float64(C.gsl_cdf_beta_Q(C.double(x), C.double(a), C.double(b)))
}

func CdfBetaPinv(p float64, a float64, b float64) float64 {
   return float64(C.gsl_cdf_beta_Pinv(C.double(p), C.double(a), C.double(b)))
}

func CdfBetaQinv(q float64, a float64, b float64) float64 {
   return float64(C.gsl_cdf_beta_Qinv(C.double(q), C.double(a), C.double(b)))
}

func Logistic(r *rng.GslRng, a float64) float64 {
   return float64(C.gsl_ran_logistic((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), C.double(a)))
}

func LogisticPdf(x float64, a float64) float64 {
   return float64(C.gsl_ran_logistic_pdf(C.double(x), C.double(a)))
}

func CdfLogisticP(x float64, a float64) float64 {
   return float64(C.gsl_cdf_logistic_P(C.double(x), C.double(a)))
}

func CdfLogisticQ(x float64, a float64) float64 {
   return float64(C.gsl_cdf_logistic_Q(C.double(x), C.double(a)))
}

func CdfLogisticPinv(p float64, a float64) float64 {
   return float64(C.gsl_cdf_logistic_Pinv(C.double(p), C.double(a)))
}

func CdfLogisticQinv(q float64, a float64) float64 {
   return float64(C.gsl_cdf_logistic_Qinv(C.double(q), C.double(a)))
}

func Pareto(r *rng.GslRng, a float64, b float64) float64 {
   return float64(C.gsl_ran_pareto((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), C.double(a), C.double(b)))
}

func ParetoPdf(x float64, a float64, b float64) float64 {
   return float64(C.gsl_ran_pareto_pdf(C.double(x), C.double(a), C.double(b)))
}

func CdfParetoP(x float64, a float64, b float64) float64 {
   return float64(C.gsl_cdf_pareto_P(C.double(x), C.double(a), C.double(b)))
}

func CdfParetoQ(x float64, a float64, b float64) float64 {
   return float64(C.gsl_cdf_pareto_Q(C.double(x), C.double(a), C.double(b)))
}

func CdfParetoPinv(p float64, a float64, b float64) float64 {
   return float64(C.gsl_cdf_pareto_Pinv(C.double(p), C.double(a), C.double(b)))
}

func CdfParetoQinv(q float64, a float64, b float64) float64 {
   return float64(C.gsl_cdf_pareto_Qinv(C.double(q), C.double(a), C.double(b)))
}

func Dir2d(r *rng.GslRng) (float64, float64) {
   var _outptr_1 C.double
   var _outptr_2 C.double
   C.gsl_ran_dir_2d((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), &_outptr_1, &_outptr_2)
   return *(*float64)(unsafe.Pointer(&_outptr_1)), *(*float64)(unsafe.Pointer(&_outptr_2))
}

func Dir2dTrigMethod(r *rng.GslRng) (float64, float64) {
   var _outptr_1 C.double
   var _outptr_2 C.double
   C.gsl_ran_dir_2d_trig_method((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), &_outptr_1, &_outptr_2)
   return *(*float64)(unsafe.Pointer(&_outptr_1)), *(*float64)(unsafe.Pointer(&_outptr_2))
}

func Dir3d(r *rng.GslRng) (float64, float64, float64) {
   var _outptr_1 C.double
   var _outptr_2 C.double
   var _outptr_3 C.double
   C.gsl_ran_dir_3d((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), &_outptr_1, &_outptr_2, &_outptr_3)
   return *(*float64)(unsafe.Pointer(&_outptr_1)), *(*float64)(unsafe.Pointer(&_outptr_2)), *(*float64)(unsafe.Pointer(&_outptr_3))
}

func DirNd(r *rng.GslRng, n int) (float64) {
   var _outptr_2 C.double
   C.gsl_ran_dir_nd((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), C.size_t(n), &_outptr_2)
   return *(*float64)(unsafe.Pointer(&_outptr_2))
}

func Weibull(r *rng.GslRng, a float64, b float64) float64 {
   return float64(C.gsl_ran_weibull((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), C.double(a), C.double(b)))
}

func WeibullPdf(x float64, a float64, b float64) float64 {
   return float64(C.gsl_ran_weibull_pdf(C.double(x), C.double(a), C.double(b)))
}

func CdfWeibullP(x float64, a float64, b float64) float64 {
   return float64(C.gsl_cdf_weibull_P(C.double(x), C.double(a), C.double(b)))
}

func CdfWeibullQ(x float64, a float64, b float64) float64 {
   return float64(C.gsl_cdf_weibull_Q(C.double(x), C.double(a), C.double(b)))
}

func CdfWeibullPinv(p float64, a float64, b float64) float64 {
   return float64(C.gsl_cdf_weibull_Pinv(C.double(p), C.double(a), C.double(b)))
}

func CdfWeibullQinv(q float64, a float64, b float64) float64 {
   return float64(C.gsl_cdf_weibull_Qinv(C.double(q), C.double(a), C.double(b)))
}

func Gumbel1(r *rng.GslRng, a float64, b float64) float64 {
   return float64(C.gsl_ran_gumbel1((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), C.double(a), C.double(b)))
}

func Gumbel1Pdf(x float64, a float64, b float64) float64 {
   return float64(C.gsl_ran_gumbel1_pdf(C.double(x), C.double(a), C.double(b)))
}

func CdfGumbel1P(x float64, a float64, b float64) float64 {
   return float64(C.gsl_cdf_gumbel1_P(C.double(x), C.double(a), C.double(b)))
}

func CdfGumbel1Q(x float64, a float64, b float64) float64 {
   return float64(C.gsl_cdf_gumbel1_Q(C.double(x), C.double(a), C.double(b)))
}

func CdfGumbel1Pinv(p float64, a float64, b float64) float64 {
   return float64(C.gsl_cdf_gumbel1_Pinv(C.double(p), C.double(a), C.double(b)))
}

func CdfGumbel1Qinv(q float64, a float64, b float64) float64 {
   return float64(C.gsl_cdf_gumbel1_Qinv(C.double(q), C.double(a), C.double(b)))
}

func Gumbel2(r *rng.GslRng, a float64, b float64) float64 {
   return float64(C.gsl_ran_gumbel2((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), C.double(a), C.double(b)))
}

func Gumbel2Pdf(x float64, a float64, b float64) float64 {
   return float64(C.gsl_ran_gumbel2_pdf(C.double(x), C.double(a), C.double(b)))
}

func CdfGumbel2P(x float64, a float64, b float64) float64 {
   return float64(C.gsl_cdf_gumbel2_P(C.double(x), C.double(a), C.double(b)))
}

func CdfGumbel2Q(x float64, a float64, b float64) float64 {
   return float64(C.gsl_cdf_gumbel2_Q(C.double(x), C.double(a), C.double(b)))
}

func CdfGumbel2Pinv(p float64, a float64, b float64) float64 {
   return float64(C.gsl_cdf_gumbel2_Pinv(C.double(p), C.double(a), C.double(b)))
}

func CdfGumbel2Qinv(q float64, a float64, b float64) float64 {
   return float64(C.gsl_cdf_gumbel2_Qinv(C.double(q), C.double(a), C.double(b)))
}

func Dirichlet(r *rng.GslRng, k int, alpha []float64, theta []float64) {
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&alpha))
   _slice_header_3 := (*reflect.SliceHeader)(unsafe.Pointer(&theta))
   C.gsl_ran_dirichlet((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), C.size_t(k), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), (*C.double)(unsafe.Pointer(_slice_header_3.Data)))
}

func DirichletPdf(k int, alpha []float64, theta []float64) float64 {
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&alpha))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&theta))
   return float64(C.gsl_ran_dirichlet_pdf(C.size_t(k), (*C.double)(unsafe.Pointer(_slice_header_1.Data)), (*C.double)(unsafe.Pointer(_slice_header_2.Data))))
}

func DirichletLnpdf(k int, alpha []float64, theta []float64) float64 {
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&alpha))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&theta))
   return float64(C.gsl_ran_dirichlet_lnpdf(C.size_t(k), (*C.double)(unsafe.Pointer(_slice_header_1.Data)), (*C.double)(unsafe.Pointer(_slice_header_2.Data))))
}

func DiscretePreproc(k int, p []float64) *GslRanDiscreteT {
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   _ref := C.gsl_ran_discrete_preproc(C.size_t(k), (*C.double)(unsafe.Pointer(_slice_header_1.Data)))
   _result := &GslRanDiscreteT{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func Discrete(r *rng.GslRng, g *GslRanDiscreteT) int {
   return int(C.gsl_ran_discrete((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), (*C.gsl_ran_discrete_t)(unsafe.Pointer(g.Ptr()))))
}

func DiscretePdf(k int, g *GslRanDiscreteT) float64 {
   return float64(C.gsl_ran_discrete_pdf(C.size_t(k), (*C.gsl_ran_discrete_t)(unsafe.Pointer(g.Ptr()))))
}

func (x *GslRanDiscreteT) Dispose() {
   C.gsl_ran_discrete_free((*C.gsl_ran_discrete_t)(unsafe.Pointer(x.Ptr())))
}

func Poisson(r *rng.GslRng, mu float64) uint32 {
   return uint32(C.gsl_ran_poisson((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), C.double(mu)))
}

func PoissonPdf(k uint32, mu float64) float64 {
   return float64(C.gsl_ran_poisson_pdf(C.uint(k), C.double(mu)))
}

func CdfPoissonP(k uint32, mu float64) float64 {
   return float64(C.gsl_cdf_poisson_P(C.uint(k), C.double(mu)))
}

func CdfPoissonQ(k uint32, mu float64) float64 {
   return float64(C.gsl_cdf_poisson_Q(C.uint(k), C.double(mu)))
}

func Bernoulli(r *rng.GslRng, p float64) uint32 {
   return uint32(C.gsl_ran_bernoulli((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), C.double(p)))
}

func BernoulliPdf(k uint32, p float64) float64 {
   return float64(C.gsl_ran_bernoulli_pdf(C.uint(k), C.double(p)))
}

func Binomial(r *rng.GslRng, p float64, n uint32) uint32 {
   return uint32(C.gsl_ran_binomial((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), C.double(p), C.uint(n)))
}

func BinomialPdf(k uint32, p float64, n uint32) float64 {
   return float64(C.gsl_ran_binomial_pdf(C.uint(k), C.double(p), C.uint(n)))
}

func CdfBinomialP(k uint32, p float64, n uint32) float64 {
   return float64(C.gsl_cdf_binomial_P(C.uint(k), C.double(p), C.uint(n)))
}

func CdfBinomialQ(k uint32, p float64, n uint32) float64 {
   return float64(C.gsl_cdf_binomial_Q(C.uint(k), C.double(p), C.uint(n)))
}

func Multinomial(r *rng.GslRng, k int, n uint32, p []float64, x []uint32) {
   _slice_header_3 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   _slice_header_4 := (*reflect.SliceHeader)(unsafe.Pointer(&x))
   C.gsl_ran_multinomial((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), C.size_t(k), C.uint(n), (*C.double)(unsafe.Pointer(_slice_header_3.Data)), (*C.uint)(unsafe.Pointer(_slice_header_4.Data)))
}

func MultinomialPdf(k int, p []float64, n []uint32) float64 {
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&n))
   return float64(C.gsl_ran_multinomial_pdf(C.size_t(k), (*C.double)(unsafe.Pointer(_slice_header_1.Data)), (*C.uint)(unsafe.Pointer(_slice_header_2.Data))))
}

func MultinomialLnpdf(k int, p []float64, n []uint32) float64 {
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&n))
   return float64(C.gsl_ran_multinomial_lnpdf(C.size_t(k), (*C.double)(unsafe.Pointer(_slice_header_1.Data)), (*C.uint)(unsafe.Pointer(_slice_header_2.Data))))
}

func NegativeBinomial(r *rng.GslRng, p float64, n float64) uint32 {
   return uint32(C.gsl_ran_negative_binomial((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), C.double(p), C.double(n)))
}

func NegativeBinomialPdf(k uint32, p float64, n float64) float64 {
   return float64(C.gsl_ran_negative_binomial_pdf(C.uint(k), C.double(p), C.double(n)))
}

func CdfNegativeBinomialP(k uint32, p float64, n float64) float64 {
   return float64(C.gsl_cdf_negative_binomial_P(C.uint(k), C.double(p), C.double(n)))
}

func CdfNegativeBinomialQ(k uint32, p float64, n float64) float64 {
   return float64(C.gsl_cdf_negative_binomial_Q(C.uint(k), C.double(p), C.double(n)))
}

func Pascal(r *rng.GslRng, p float64, n uint32) uint32 {
   return uint32(C.gsl_ran_pascal((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), C.double(p), C.uint(n)))
}

func PascalPdf(k uint32, p float64, n uint32) float64 {
   return float64(C.gsl_ran_pascal_pdf(C.uint(k), C.double(p), C.uint(n)))
}

func CdfPascalP(k uint32, p float64, n uint32) float64 {
   return float64(C.gsl_cdf_pascal_P(C.uint(k), C.double(p), C.uint(n)))
}

func CdfPascalQ(k uint32, p float64, n uint32) float64 {
   return float64(C.gsl_cdf_pascal_Q(C.uint(k), C.double(p), C.uint(n)))
}

func Geometric(r *rng.GslRng, p float64) uint32 {
   return uint32(C.gsl_ran_geometric((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), C.double(p)))
}

func GeometricPdf(k uint32, p float64) float64 {
   return float64(C.gsl_ran_geometric_pdf(C.uint(k), C.double(p)))
}

func CdfGeometricP(k uint32, p float64) float64 {
   return float64(C.gsl_cdf_geometric_P(C.uint(k), C.double(p)))
}

func CdfGeometricQ(k uint32, p float64) float64 {
   return float64(C.gsl_cdf_geometric_Q(C.uint(k), C.double(p)))
}

func Hypergeometric(r *rng.GslRng, n1 uint32, n2 uint32, t uint32) uint32 {
   return uint32(C.gsl_ran_hypergeometric((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), C.uint(n1), C.uint(n2), C.uint(t)))
}

func HypergeometricPdf(k uint32, n1 uint32, n2 uint32, t uint32) float64 {
   return float64(C.gsl_ran_hypergeometric_pdf(C.uint(k), C.uint(n1), C.uint(n2), C.uint(t)))
}

func CdfHypergeometricP(k uint32, n1 uint32, n2 uint32, t uint32) float64 {
   return float64(C.gsl_cdf_hypergeometric_P(C.uint(k), C.uint(n1), C.uint(n2), C.uint(t)))
}

func CdfHypergeometricQ(k uint32, n1 uint32, n2 uint32, t uint32) float64 {
   return float64(C.gsl_cdf_hypergeometric_Q(C.uint(k), C.uint(n1), C.uint(n2), C.uint(t)))
}

func Logarithmic(r *rng.GslRng, p float64) uint32 {
   return uint32(C.gsl_ran_logarithmic((*C.gsl_rng)(unsafe.Pointer(r.Ptr())), C.double(p)))
}

func LogarithmicPdf(k uint32, p float64) float64 {
   return float64(C.gsl_ran_logarithmic_pdf(C.uint(k), C.double(p)))
}

