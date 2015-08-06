package gogsl

/*	
	#include <stdlib.h>
	
	int get_c_size_t_size() {
	int c;
	size_t k;
		k = k - 1;
		while (k) {
			k >>= 1;
			c += 1;
		}
		return c;
	}
	
	int get_c_short_size() {
	int c;
	unsigned short k;
		k = k - 1;
		while (k) {
			k >>= 1;
			c += 1;
		}
		return c;
	}
		
	int get_c_int_size() {
	int c;
	unsigned int k;
		k = k - 1;
		while (k) {
			k >>= 1;
			c += 1;
		}
		return c;
	}
	
	int get_c_long_size() {
	int c;
	unsigned long k;
		k = k - 1;
		while (k) {
			k >>= 1;
			c += 1;
		}
		return c;
	}
	
	int get_c_long_long_size() {
	int c;
	unsigned long long k;
		k = k - 1;
		while (k) {
			k >>= 1;
			c += 1;
		}
		return c;
	}
*/
import "C"

import (
	"errors"
	"reflect"
)

var GOGSL_SIZE_T_TYPE reflect.Type
var GOGSL_INT_TYPE reflect.Type
var GOGSL_UINT_TYPE reflect.Type
var GOGSL_SHORT_TYPE reflect.Type
var GOGSL_USHORT_TYPE reflect.Type
var GOGSL_LONG_TYPE reflect.Type
var GOGSL_ULONG_TYPE reflect.Type
var GOGSL_LONG_LONG_TYPE reflect.Type
var GOGSL_ULONG_LONG_TYPE reflect.Type
var GOGSL_C16_TYPE reflect.Type
var GOGSL_C32_TYPE reflect.Type
var GOGSL_C64_TYPE reflect.Type

func init() {
	// XXX - Check a type size signature hash against a generation-time value here,
	// 		 to detect if regeneration is necessary, and issue warning if so.
	var err error
	if GOGSL_SIZE_T_TYPE, err = GetNumberType(false, GetCSizeTSize()); err != nil {
		panic(err.Error())
	}	
	if GOGSL_INT_TYPE, err = GetNumberType(true, GetCIntSize()); err != nil {
		panic(err.Error())
	}	
	if GOGSL_UINT_TYPE, err = GetNumberType(false, GetCIntSize()); err != nil {
		panic(err.Error())
	}	
	if GOGSL_SHORT_TYPE, err = GetNumberType(true, GetCShortSize()); err != nil {
		panic(err.Error())
	}	
	if GOGSL_USHORT_TYPE, err = GetNumberType(false, GetCShortSize()); err != nil {
		panic(err.Error())
	}	
	if GOGSL_LONG_TYPE, err = GetNumberType(true, GetCLongSize()); err != nil {
		panic(err.Error())
	}	
	if GOGSL_ULONG_TYPE, err = GetNumberType(false, GetCLongSize()); err != nil {
		panic(err.Error())
	}	
	if GOGSL_LONG_LONG_TYPE, err = GetNumberType(false, GetCLongLongSize()); err != nil {
		panic(err.Error())
	}	
	if GOGSL_ULONG_LONG_TYPE, err = GetNumberType(false, GetCLongLongSize()); err != nil {
		panic(err.Error())
	}
}

func GetIntSize() int {
	var c, k uint
	k = k - 1
	for k > 0 {
		c += 1
		k >>= 1
	}
	return int(c)
}

func GetCSizeTSize() int {
	return int(C.get_c_size_t_size())
}

func GetCShortSize() int {
	return int(C.get_c_short_size())
}

func GetCIntSize() int {
	return int(C.get_c_int_size())
}

func GetCLongSize() int {
	return int(C.get_c_long_size())
}

func GetCLongLongSize() int {
	return int(C.get_c_long_long_size())
}

func GetNumberType(signed bool, n int) (reflect.Type, error) {
	intSize := GetIntSize()
	if intSize == n {
		if signed {
			return reflect.TypeOf([]uint{}).Elem(), nil
		}
		return reflect.TypeOf([]int{}).Elem(), nil
	}
	switch n {
		case 8: {
			if signed {
				return reflect.TypeOf([]int8{}).Elem(), nil
			} else {
				return reflect.TypeOf([]byte{}).Elem(), nil
			}
		}
		case 16: {
			if signed {
				return reflect.TypeOf([]int16{}).Elem(), nil
			} else {
				return reflect.TypeOf([]uint16{}).Elem(), nil
			}
		}
		case 32: {
			if signed {
				return reflect.TypeOf([]int32{}).Elem(), nil
			} else {
				return reflect.TypeOf([]uint32{}).Elem(), nil
			}
		}
		case 64: {
			if signed {
				return reflect.TypeOf([]int64{}).Elem(), nil
			} else {
				return reflect.TypeOf([]uint64{}).Elem(), nil
			}
		}
	}
	return nil, errors.New("no default mapping for numeric type available")
}