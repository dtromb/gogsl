package sort

type GslComparisonFn func(interface{}, interface{}) int

//void gsl_heapsort (void * array, size_t count, size_t size, gsl_comparison_fn_t compare)
func Heapsort(array []interface{}, count int, size int, fn GslComparisonFn) int {
	return 0
}

// int gsl_heapsort_index (size_t * p, const void * array, size_t count, size_t size, gsl_comparison_fn_t compare)
func HeapsortIndex(p []int, array []interface{}, count int, size int, fn GslComparisonFn) int {
	return 0
}
