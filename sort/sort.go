//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package sort

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_complex.h>
   #include <gsl/gsl_heapsort.h>
   #include <gsl/gsl_sort.h>
   #include <gsl/gsl_sort_vector.h>
*/
import "C"

import "reflect"
import "unsafe"
import "github.com/dtromb/gogsl/vector"
import "github.com/dtromb/gogsl/permutation"

func Sort(data []float64, stride int, n int) {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   C.gsl_sort((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n))
}

func Sort2(data1 []float64, stride1 int, data2 []float64, stride2 int, n int) {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   C.gsl_sort2((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n))
}

func SortVector(v *vector.GslVector) {
   C.gsl_sort_vector((*C.gsl_vector)(unsafe.Pointer(v.Ptr())))
}

func SortVector2(v1 *vector.GslVector, v2 *vector.GslVector) {
   C.gsl_sort_vector2((*C.gsl_vector)(unsafe.Pointer(v1.Ptr())), (*C.gsl_vector)(unsafe.Pointer(v2.Ptr())))
}

func SortIndex(p []int, data []float64, stride int, n int) {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   C.gsl_sort_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), (*C.double)(unsafe.Pointer(_slice_header_1.Data)), C.size_t(stride), C.size_t(n))
}

func SortVectorIndex(p *permutation.GslPermutation, v *vector.GslVector) int32 {
   return int32(C.gsl_sort_vector_index((*C.gsl_permutation)(unsafe.Pointer(p.Ptr())), (*C.gsl_vector)(unsafe.Pointer(v.Ptr()))))
}

func SortSmallest(dest []float64, k int, src []float64, stride int, n int) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&dest))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
   return int32(C.gsl_sort_smallest((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func SortLargest(dest []float64, k int, src []float64, stride int, n int) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&dest))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
   return int32(C.gsl_sort_largest((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func SortVectorSmallest(dest []float64, k int, v *vector.GslVector) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&dest))
   return int32(C.gsl_sort_vector_smallest((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.gsl_vector)(unsafe.Pointer(v.Ptr()))))
}

func SortVectorLargest(dest []float64, k int, v *vector.GslVector) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&dest))
   return int32(C.gsl_sort_vector_largest((*C.double)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.gsl_vector)(unsafe.Pointer(v.Ptr()))))
}

func SortSmallestIndex(p []int, k int, src []float64, stride int, n int) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
   return int32(C.gsl_sort_smallest_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func SortLargestIndex(p []int, k int, src []float64, stride int, n int) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
   return int32(C.gsl_sort_largest_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func SortVectorSmallestIndex(p []int, k int, v *vector.GslVector) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   return int32(C.gsl_sort_vector_smallest_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.gsl_vector)(unsafe.Pointer(v.Ptr()))))
}

func SortVectorLargestIndex(p []int, k int, v *vector.GslVector) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   return int32(C.gsl_sort_vector_largest_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.gsl_vector)(unsafe.Pointer(v.Ptr()))))
}

func SortFloat(data []float32, stride int, n int) {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   C.gsl_sort_float((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n))
}

func Sort2Float(data1 []float32, stride1 int, data2 []float32, stride2 int, n int) {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   C.gsl_sort2_float((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.float)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n))
}

func SortVectorFloat(v *vector.GslVectorFloat) {
   C.gsl_sort_vector_float((*C.gsl_vector_float)(unsafe.Pointer(v.Ptr())))
}

func SortVector2Float(v1 *vector.GslVectorFloat, v2 *vector.GslVectorFloat) {
   C.gsl_sort_vector2_float((*C.gsl_vector_float)(unsafe.Pointer(v1.Ptr())), (*C.gsl_vector_float)(unsafe.Pointer(v2.Ptr())))
}

func SortFloatIndex(p []int, data []float32, stride int, n int) {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   C.gsl_sort_float_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), (*C.float)(unsafe.Pointer(_slice_header_1.Data)), C.size_t(stride), C.size_t(n))
}

func SortVectorFloatIndex(p *permutation.GslPermutation, v *vector.GslVectorFloat) int32 {
   return int32(C.gsl_sort_vector_float_index((*C.gsl_permutation)(unsafe.Pointer(p.Ptr())), (*C.gsl_vector_float)(unsafe.Pointer(v.Ptr()))))
}

func SortFloatSmallest(dest []float32, k int, src []float32, stride int, n int) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&dest))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
   return int32(C.gsl_sort_float_smallest((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.float)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func SortFloatLargest(dest []float32, k int, src []float32, stride int, n int) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&dest))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
   return int32(C.gsl_sort_float_largest((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.float)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func SortVectorFloatSmallest(dest []float32, k int, v *vector.GslVectorFloat) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&dest))
   return int32(C.gsl_sort_vector_float_smallest((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.gsl_vector_float)(unsafe.Pointer(v.Ptr()))))
}

func SortVectorFloatLargest(dest []float32, k int, v *vector.GslVectorFloat) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&dest))
   return int32(C.gsl_sort_vector_float_largest((*C.float)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.gsl_vector_float)(unsafe.Pointer(v.Ptr()))))
}

func SortFloatSmallestIndex(p []int, k int, src []float32, stride int, n int) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
   return int32(C.gsl_sort_float_smallest_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.float)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func SortFloatLargestIndex(p []int, k int, src []float32, stride int, n int) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
   return int32(C.gsl_sort_float_largest_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.float)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func SortVectorFloatSmallestIndex(p []int, k int, v *vector.GslVectorFloat) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   return int32(C.gsl_sort_vector_float_smallest_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.gsl_vector_float)(unsafe.Pointer(v.Ptr()))))
}

func SortVectorFloatLargestIndex(p []int, k int, v *vector.GslVectorFloat) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   return int32(C.gsl_sort_vector_float_largest_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.gsl_vector_float)(unsafe.Pointer(v.Ptr()))))
}

func SortInt(data []int32, stride int, n int) {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   C.gsl_sort_int((*C.int)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n))
}

func Sort2Int(data1 []int32, stride1 int, data2 []int32, stride2 int, n int) {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   C.gsl_sort2_int((*C.int)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.int)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n))
}

func SortVectorInt(v *vector.GslVectorInt) {
   C.gsl_sort_vector_int((*C.gsl_vector_int)(unsafe.Pointer(v.Ptr())))
}

func SortVector2Int(v1 *vector.GslVectorInt, v2 *vector.GslVectorInt) {
   C.gsl_sort_vector2_int((*C.gsl_vector_int)(unsafe.Pointer(v1.Ptr())), (*C.gsl_vector_int)(unsafe.Pointer(v2.Ptr())))
}

func SortIntIndex(p []int, data []int32, stride int, n int) {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   C.gsl_sort_int_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), (*C.int)(unsafe.Pointer(_slice_header_1.Data)), C.size_t(stride), C.size_t(n))
}

func SortVectorIntIndex(p *permutation.GslPermutation, v *vector.GslVectorInt) int32 {
   return int32(C.gsl_sort_vector_int_index((*C.gsl_permutation)(unsafe.Pointer(p.Ptr())), (*C.gsl_vector_int)(unsafe.Pointer(v.Ptr()))))
}

func SortIntSmallest(dest []int32, k int, src []int32, stride int, n int) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&dest))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
   return int32(C.gsl_sort_int_smallest((*C.int)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.int)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func SortIntLargest(dest []int32, k int, src []int32, stride int, n int) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&dest))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
   return int32(C.gsl_sort_int_largest((*C.int)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.int)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func SortVectorIntSmallest(dest []int32, k int, v *vector.GslVectorInt) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&dest))
   return int32(C.gsl_sort_vector_int_smallest((*C.int)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.gsl_vector_int)(unsafe.Pointer(v.Ptr()))))
}

func SortVectorIntLargest(dest []int32, k int, v *vector.GslVectorInt) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&dest))
   return int32(C.gsl_sort_vector_int_largest((*C.int)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.gsl_vector_int)(unsafe.Pointer(v.Ptr()))))
}

func SortIntSmallestIndex(p []int, k int, src []int32, stride int, n int) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
   return int32(C.gsl_sort_int_smallest_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.int)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func SortIntLargestIndex(p []int, k int, src []int32, stride int, n int) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
   return int32(C.gsl_sort_int_largest_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.int)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func SortVectorIntSmallestIndex(p []int, k int, v *vector.GslVectorInt) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   return int32(C.gsl_sort_vector_int_smallest_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.gsl_vector_int)(unsafe.Pointer(v.Ptr()))))
}

func SortVectorIntLargestIndex(p []int, k int, v *vector.GslVectorInt) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   return int32(C.gsl_sort_vector_int_largest_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.gsl_vector_int)(unsafe.Pointer(v.Ptr()))))
}

func SortUint(data []uint32, stride int, n int) {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   C.gsl_sort_uint((*C.uint)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n))
}

func Sort2Uint(data1 []uint32, stride1 int, data2 []uint32, stride2 int, n int) {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   C.gsl_sort2_uint((*C.uint)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.uint)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n))
}

func SortVectorUint(v *vector.GslVectorUint) {
   C.gsl_sort_vector_uint((*C.gsl_vector_uint)(unsafe.Pointer(v.Ptr())))
}

func SortVector2Uint(v1 *vector.GslVectorUint, v2 *vector.GslVectorUint) {
   C.gsl_sort_vector2_uint((*C.gsl_vector_uint)(unsafe.Pointer(v1.Ptr())), (*C.gsl_vector_uint)(unsafe.Pointer(v2.Ptr())))
}

func SortUintIndex(p []int, data []uint32, stride int, n int) {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   C.gsl_sort_uint_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), (*C.uint)(unsafe.Pointer(_slice_header_1.Data)), C.size_t(stride), C.size_t(n))
}

func SortVectorUintIndex(p *permutation.GslPermutation, v *vector.GslVectorUint) int32 {
   return int32(C.gsl_sort_vector_uint_index((*C.gsl_permutation)(unsafe.Pointer(p.Ptr())), (*C.gsl_vector_uint)(unsafe.Pointer(v.Ptr()))))
}

func SortUintSmallest(dest []uint32, k int, src []uint32, stride int, n int) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&dest))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
   return int32(C.gsl_sort_uint_smallest((*C.uint)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.uint)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func SortUintLargest(dest []uint32, k int, src []uint32, stride int, n int) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&dest))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
   return int32(C.gsl_sort_uint_largest((*C.uint)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.uint)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func SortVectorUintSmallest(dest []uint32, k int, v *vector.GslVectorUint) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&dest))
   return int32(C.gsl_sort_vector_uint_smallest((*C.uint)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.gsl_vector_uint)(unsafe.Pointer(v.Ptr()))))
}

func SortVectorUintLargest(dest []uint32, k int, v *vector.GslVectorUint) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&dest))
   return int32(C.gsl_sort_vector_uint_largest((*C.uint)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.gsl_vector_uint)(unsafe.Pointer(v.Ptr()))))
}

func SortUintSmallestIndex(p []int, k int, src []uint32, stride int, n int) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
   return int32(C.gsl_sort_uint_smallest_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.uint)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func SortUintLargestIndex(p []int, k int, src []uint32, stride int, n int) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
   return int32(C.gsl_sort_uint_largest_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.uint)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func SortVectorUintSmallestIndex(p []int, k int, v *vector.GslVectorUint) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   return int32(C.gsl_sort_vector_uint_smallest_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.gsl_vector_uint)(unsafe.Pointer(v.Ptr()))))
}

func SortVectorUintLargestIndex(p []int, k int, v *vector.GslVectorUint) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   return int32(C.gsl_sort_vector_uint_largest_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.gsl_vector_uint)(unsafe.Pointer(v.Ptr()))))
}

func SortLong(data []uint, stride int, n int) {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   C.gsl_sort_long((*C.long)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n))
}

func Sort2Long(data1 []uint, stride1 int, data2 []uint, stride2 int, n int) {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   C.gsl_sort2_long((*C.long)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.long)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n))
}

func SortVectorLong(v *vector.GslVectorLong) {
   C.gsl_sort_vector_long((*C.gsl_vector_long)(unsafe.Pointer(v.Ptr())))
}

func SortVector2Long(v1 *vector.GslVectorLong, v2 *vector.GslVectorLong) {
   C.gsl_sort_vector2_long((*C.gsl_vector_long)(unsafe.Pointer(v1.Ptr())), (*C.gsl_vector_long)(unsafe.Pointer(v2.Ptr())))
}

func SortLongIndex(p []int, data []uint, stride int, n int) {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   C.gsl_sort_long_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), (*C.long)(unsafe.Pointer(_slice_header_1.Data)), C.size_t(stride), C.size_t(n))
}

func SortVectorLongIndex(p *permutation.GslPermutation, v *vector.GslVectorLong) int32 {
   return int32(C.gsl_sort_vector_long_index((*C.gsl_permutation)(unsafe.Pointer(p.Ptr())), (*C.gsl_vector_long)(unsafe.Pointer(v.Ptr()))))
}

func SortLongSmallest(dest []uint, k int, src []uint, stride int, n int) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&dest))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
   return int32(C.gsl_sort_long_smallest((*C.long)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.long)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func SortLongLargest(dest []uint, k int, src []uint, stride int, n int) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&dest))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
   return int32(C.gsl_sort_long_largest((*C.long)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.long)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func SortVectorLongSmallest(dest []uint, k int, v *vector.GslVectorLong) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&dest))
   return int32(C.gsl_sort_vector_long_smallest((*C.long)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.gsl_vector_long)(unsafe.Pointer(v.Ptr()))))
}

func SortVectorLongLargest(dest []uint, k int, v *vector.GslVectorLong) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&dest))
   return int32(C.gsl_sort_vector_long_largest((*C.long)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.gsl_vector_long)(unsafe.Pointer(v.Ptr()))))
}

func SortLongSmallestIndex(p []int, k int, src []uint, stride int, n int) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
   return int32(C.gsl_sort_long_smallest_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.long)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func SortLongLargestIndex(p []int, k int, src []uint, stride int, n int) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
   return int32(C.gsl_sort_long_largest_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.long)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func SortVectorLongSmallestIndex(p []int, k int, v *vector.GslVectorLong) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   return int32(C.gsl_sort_vector_long_smallest_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.gsl_vector_long)(unsafe.Pointer(v.Ptr()))))
}

func SortVectorLongLargestIndex(p []int, k int, v *vector.GslVectorLong) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   return int32(C.gsl_sort_vector_long_largest_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.gsl_vector_long)(unsafe.Pointer(v.Ptr()))))
}

func SortUlong(data []int, stride int, n int) {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   C.gsl_sort_ulong((*C.ulong)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n))
}

func Sort2Ulong(data1 []int, stride1 int, data2 []int, stride2 int, n int) {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   C.gsl_sort2_ulong((*C.ulong)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.ulong)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n))
}

func SortVectorUlong(v *vector.GslVectorUlong) {
   C.gsl_sort_vector_ulong((*C.gsl_vector_ulong)(unsafe.Pointer(v.Ptr())))
}

func SortVector2Ulong(v1 *vector.GslVectorUlong, v2 *vector.GslVectorUlong) {
   C.gsl_sort_vector2_ulong((*C.gsl_vector_ulong)(unsafe.Pointer(v1.Ptr())), (*C.gsl_vector_ulong)(unsafe.Pointer(v2.Ptr())))
}

func SortUlongIndex(p []int, data []int, stride int, n int) {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   C.gsl_sort_ulong_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), (*C.ulong)(unsafe.Pointer(_slice_header_1.Data)), C.size_t(stride), C.size_t(n))
}

func SortVectorUlongIndex(p *permutation.GslPermutation, v *vector.GslVectorUlong) int32 {
   return int32(C.gsl_sort_vector_ulong_index((*C.gsl_permutation)(unsafe.Pointer(p.Ptr())), (*C.gsl_vector_ulong)(unsafe.Pointer(v.Ptr()))))
}

func SortUlongSmallest(dest []int, k int, src []int, stride int, n int) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&dest))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
   return int32(C.gsl_sort_ulong_smallest((*C.ulong)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.ulong)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func SortUlongLargest(dest []int, k int, src []int, stride int, n int) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&dest))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
   return int32(C.gsl_sort_ulong_largest((*C.ulong)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.ulong)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func SortVectorUlongSmallest(dest []int, k int, v *vector.GslVectorUlong) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&dest))
   return int32(C.gsl_sort_vector_ulong_smallest((*C.ulong)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.gsl_vector_ulong)(unsafe.Pointer(v.Ptr()))))
}

func SortVectorUlongLargest(dest []int, k int, v *vector.GslVectorUlong) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&dest))
   return int32(C.gsl_sort_vector_ulong_largest((*C.ulong)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.gsl_vector_ulong)(unsafe.Pointer(v.Ptr()))))
}

func SortUlongSmallestIndex(p []int, k int, src []int, stride int, n int) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
   return int32(C.gsl_sort_ulong_smallest_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.ulong)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func SortUlongLargestIndex(p []int, k int, src []int, stride int, n int) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
   return int32(C.gsl_sort_ulong_largest_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.ulong)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func SortVectorUlongSmallestIndex(p []int, k int, v *vector.GslVectorUlong) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   return int32(C.gsl_sort_vector_ulong_smallest_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.gsl_vector_ulong)(unsafe.Pointer(v.Ptr()))))
}

func SortVectorUlongLargestIndex(p []int, k int, v *vector.GslVectorUlong) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   return int32(C.gsl_sort_vector_ulong_largest_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.gsl_vector_ulong)(unsafe.Pointer(v.Ptr()))))
}

func SortShort(data []int16, stride int, n int) {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   C.gsl_sort_short((*C.short)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n))
}

func Sort2Short(data1 []int16, stride1 int, data2 []int16, stride2 int, n int) {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   C.gsl_sort2_short((*C.short)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.short)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n))
}

func SortVectorShort(v *vector.GslVectorShort) {
   C.gsl_sort_vector_short((*C.gsl_vector_short)(unsafe.Pointer(v.Ptr())))
}

func SortVector2Short(v1 *vector.GslVectorShort, v2 *vector.GslVectorShort) {
   C.gsl_sort_vector2_short((*C.gsl_vector_short)(unsafe.Pointer(v1.Ptr())), (*C.gsl_vector_short)(unsafe.Pointer(v2.Ptr())))
}

func SortShortIndex(p []int, data []int16, stride int, n int) {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   C.gsl_sort_short_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), (*C.short)(unsafe.Pointer(_slice_header_1.Data)), C.size_t(stride), C.size_t(n))
}

func SortVectorShortIndex(p *permutation.GslPermutation, v *vector.GslVectorShort) int32 {
   return int32(C.gsl_sort_vector_short_index((*C.gsl_permutation)(unsafe.Pointer(p.Ptr())), (*C.gsl_vector_short)(unsafe.Pointer(v.Ptr()))))
}

func SortShortSmallest(dest []int16, k int, src []int16, stride int, n int) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&dest))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
   return int32(C.gsl_sort_short_smallest((*C.short)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.short)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func SortShortLargest(dest []int16, k int, src []int16, stride int, n int) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&dest))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
   return int32(C.gsl_sort_short_largest((*C.short)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.short)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func SortVectorShortSmallest(dest []int16, k int, v *vector.GslVectorShort) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&dest))
   return int32(C.gsl_sort_vector_short_smallest((*C.short)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.gsl_vector_short)(unsafe.Pointer(v.Ptr()))))
}

func SortVectorShortLargest(dest []int16, k int, v *vector.GslVectorShort) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&dest))
   return int32(C.gsl_sort_vector_short_largest((*C.short)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.gsl_vector_short)(unsafe.Pointer(v.Ptr()))))
}

func SortShortSmallestIndex(p []int, k int, src []int16, stride int, n int) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
   return int32(C.gsl_sort_short_smallest_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.short)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func SortShortLargestIndex(p []int, k int, src []int16, stride int, n int) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
   return int32(C.gsl_sort_short_largest_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.short)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func SortVectorShortSmallestIndex(p []int, k int, v *vector.GslVectorShort) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   return int32(C.gsl_sort_vector_short_smallest_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.gsl_vector_short)(unsafe.Pointer(v.Ptr()))))
}

func SortVectorShortLargestIndex(p []int, k int, v *vector.GslVectorShort) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   return int32(C.gsl_sort_vector_short_largest_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.gsl_vector_short)(unsafe.Pointer(v.Ptr()))))
}

func SortUshort(data []int16, stride int, n int) {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   C.gsl_sort_ushort((*C.ushort)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n))
}

func Sort2Ushort(data1 []int16, stride1 int, data2 []int16, stride2 int, n int) {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   C.gsl_sort2_ushort((*C.ushort)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.ushort)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n))
}

func SortVectorUshort(v *vector.GslVectorUshort) {
   C.gsl_sort_vector_ushort((*C.gsl_vector_ushort)(unsafe.Pointer(v.Ptr())))
}

func SortVector2Ushort(v1 *vector.GslVectorUshort, v2 *vector.GslVectorUshort) {
   C.gsl_sort_vector2_ushort((*C.gsl_vector_ushort)(unsafe.Pointer(v1.Ptr())), (*C.gsl_vector_ushort)(unsafe.Pointer(v2.Ptr())))
}

func SortUshortIndex(p []int, data []int16, stride int, n int) {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   C.gsl_sort_ushort_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), (*C.ushort)(unsafe.Pointer(_slice_header_1.Data)), C.size_t(stride), C.size_t(n))
}

func SortVectorUshortIndex(p *permutation.GslPermutation, v *vector.GslVectorUshort) int32 {
   return int32(C.gsl_sort_vector_ushort_index((*C.gsl_permutation)(unsafe.Pointer(p.Ptr())), (*C.gsl_vector_ushort)(unsafe.Pointer(v.Ptr()))))
}

func SortUshortSmallest(dest []int16, k int, src []int16, stride int, n int) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&dest))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
   return int32(C.gsl_sort_ushort_smallest((*C.ushort)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.ushort)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func SortUshortLargest(dest []int16, k int, src []int16, stride int, n int) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&dest))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
   return int32(C.gsl_sort_ushort_largest((*C.ushort)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.ushort)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func SortVectorUshortSmallest(dest []int16, k int, v *vector.GslVectorUshort) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&dest))
   return int32(C.gsl_sort_vector_ushort_smallest((*C.ushort)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.gsl_vector_ushort)(unsafe.Pointer(v.Ptr()))))
}

func SortVectorUshortLargest(dest []int16, k int, v *vector.GslVectorUshort) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&dest))
   return int32(C.gsl_sort_vector_ushort_largest((*C.ushort)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.gsl_vector_ushort)(unsafe.Pointer(v.Ptr()))))
}

func SortUshortSmallestIndex(p []int, k int, src []int16, stride int, n int) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
   return int32(C.gsl_sort_ushort_smallest_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.ushort)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func SortUshortLargestIndex(p []int, k int, src []int16, stride int, n int) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
   return int32(C.gsl_sort_ushort_largest_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.ushort)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func SortVectorUshortSmallestIndex(p []int, k int, v *vector.GslVectorUshort) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   return int32(C.gsl_sort_vector_ushort_smallest_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.gsl_vector_ushort)(unsafe.Pointer(v.Ptr()))))
}

func SortVectorUshortLargestIndex(p []int, k int, v *vector.GslVectorUshort) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   return int32(C.gsl_sort_vector_ushort_largest_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.gsl_vector_ushort)(unsafe.Pointer(v.Ptr()))))
}

func SortChar(data []int8, stride int, n int) {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   C.gsl_sort_char((*C.char)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n))
}

func Sort2Char(data1 []int8, stride1 int, data2 []int8, stride2 int, n int) {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   C.gsl_sort2_char((*C.char)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.char)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n))
}

func SortVectorChar(v *vector.GslVectorChar) {
   C.gsl_sort_vector_char((*C.gsl_vector_char)(unsafe.Pointer(v.Ptr())))
}

func SortVector2Char(v1 *vector.GslVectorChar, v2 *vector.GslVectorChar) {
   C.gsl_sort_vector2_char((*C.gsl_vector_char)(unsafe.Pointer(v1.Ptr())), (*C.gsl_vector_char)(unsafe.Pointer(v2.Ptr())))
}

func SortCharIndex(p []int, data []int8, stride int, n int) {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   C.gsl_sort_char_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), (*C.char)(unsafe.Pointer(_slice_header_1.Data)), C.size_t(stride), C.size_t(n))
}

func SortVectorCharIndex(p *permutation.GslPermutation, v *vector.GslVectorChar) int32 {
   return int32(C.gsl_sort_vector_char_index((*C.gsl_permutation)(unsafe.Pointer(p.Ptr())), (*C.gsl_vector_char)(unsafe.Pointer(v.Ptr()))))
}

func SortCharSmallest(dest []int8, k int, src []int8, stride int, n int) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&dest))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
   return int32(C.gsl_sort_char_smallest((*C.char)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.char)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func SortCharLargest(dest []int8, k int, src []int8, stride int, n int) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&dest))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
   return int32(C.gsl_sort_char_largest((*C.char)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.char)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func SortVectorCharSmallest(dest []int8, k int, v *vector.GslVectorChar) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&dest))
   return int32(C.gsl_sort_vector_char_smallest((*C.char)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.gsl_vector_char)(unsafe.Pointer(v.Ptr()))))
}

func SortVectorCharLargest(dest []int8, k int, v *vector.GslVectorChar) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&dest))
   return int32(C.gsl_sort_vector_char_largest((*C.char)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.gsl_vector_char)(unsafe.Pointer(v.Ptr()))))
}

func SortCharSmallestIndex(p []int, k int, src []int8, stride int, n int) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
   return int32(C.gsl_sort_char_smallest_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.char)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func SortCharLargestIndex(p []int, k int, src []int8, stride int, n int) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
   return int32(C.gsl_sort_char_largest_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.char)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func SortVectorCharSmallestIndex(p []int, k int, v *vector.GslVectorChar) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   return int32(C.gsl_sort_vector_char_smallest_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.gsl_vector_char)(unsafe.Pointer(v.Ptr()))))
}

func SortVectorCharLargestIndex(p []int, k int, v *vector.GslVectorChar) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   return int32(C.gsl_sort_vector_char_largest_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.gsl_vector_char)(unsafe.Pointer(v.Ptr()))))
}

func SortUchar(data []uint8, stride int, n int) {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   C.gsl_sort_uchar((*C.uchar)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride), C.size_t(n))
}

func Sort2Uchar(data1 []uint8, stride1 int, data2 []uint8, stride2 int, n int) {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&data1))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&data2))
   C.gsl_sort2_uchar((*C.uchar)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(stride1), (*C.uchar)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride2), C.size_t(n))
}

func SortVectorUchar(v *vector.GslVectorUchar) {
   C.gsl_sort_vector_uchar((*C.gsl_vector_uchar)(unsafe.Pointer(v.Ptr())))
}

func SortVector2Uchar(v1 *vector.GslVectorUchar, v2 *vector.GslVectorUchar) {
   C.gsl_sort_vector2_uchar((*C.gsl_vector_uchar)(unsafe.Pointer(v1.Ptr())), (*C.gsl_vector_uchar)(unsafe.Pointer(v2.Ptr())))
}

func SortUcharIndex(p []int, data []uint8, stride int, n int) {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
   C.gsl_sort_uchar_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), (*C.uchar)(unsafe.Pointer(_slice_header_1.Data)), C.size_t(stride), C.size_t(n))
}

func SortVectorUcharIndex(p *permutation.GslPermutation, v *vector.GslVectorUchar) int32 {
   return int32(C.gsl_sort_vector_uchar_index((*C.gsl_permutation)(unsafe.Pointer(p.Ptr())), (*C.gsl_vector_uchar)(unsafe.Pointer(v.Ptr()))))
}

func SortUcharSmallest(dest []uint8, k int, src []uint8, stride int, n int) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&dest))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
   return int32(C.gsl_sort_uchar_smallest((*C.uchar)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.uchar)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func SortUcharLargest(dest []uint8, k int, src []uint8, stride int, n int) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&dest))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
   return int32(C.gsl_sort_uchar_largest((*C.uchar)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.uchar)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func SortVectorUcharSmallest(dest []uint8, k int, v *vector.GslVectorUchar) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&dest))
   return int32(C.gsl_sort_vector_uchar_smallest((*C.uchar)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.gsl_vector_uchar)(unsafe.Pointer(v.Ptr()))))
}

func SortVectorUcharLargest(dest []uint8, k int, v *vector.GslVectorUchar) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&dest))
   return int32(C.gsl_sort_vector_uchar_largest((*C.uchar)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.gsl_vector_uchar)(unsafe.Pointer(v.Ptr()))))
}

func SortUcharSmallestIndex(p []int, k int, src []uint8, stride int, n int) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
   return int32(C.gsl_sort_uchar_smallest_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.uchar)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func SortUcharLargestIndex(p []int, k int, src []uint8, stride int, n int) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
   return int32(C.gsl_sort_uchar_largest_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.uchar)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(stride), C.size_t(n)))
}

func SortVectorUcharSmallestIndex(p []int, k int, v *vector.GslVectorUchar) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   return int32(C.gsl_sort_vector_uchar_smallest_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.gsl_vector_uchar)(unsafe.Pointer(v.Ptr()))))
}

func SortVectorUcharLargestIndex(p []int, k int, v *vector.GslVectorUchar) int32 {
   _slice_header_0 := (*reflect.SliceHeader)(unsafe.Pointer(&p))
   return int32(C.gsl_sort_vector_uchar_largest_index((*C.size_t)(unsafe.Pointer(_slice_header_0.Data)), C.size_t(k), (*C.gsl_vector_uchar)(unsafe.Pointer(v.Ptr()))))
}

