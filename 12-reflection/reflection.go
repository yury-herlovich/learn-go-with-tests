package main

import "reflect"

func walk(x interface{}, fn func(string)) {
	elems := getValue(x)

	switch elems.Kind() {
	case reflect.String:
		fn(elems.String())
	case reflect.Slice, reflect.Array:
		for i := 0; i < elems.Len(); i++ {
			walk(elems.Index(i).Interface(), fn)
		}
	case reflect.Struct:
		for i := 0; i < elems.NumField(); i++ {
			walk(elems.Field(i).Interface(), fn)
		}
	case reflect.Map:
		for _, key := range elems.MapKeys() {
			walk(elems.MapIndex(key).Interface(), fn)
		}
	case reflect.Chan:
		for v, ok := elems.Recv(); ok; v, ok = elems.Recv() {
			walk(v.Interface(), fn)
		}
	case reflect.Func:
		fnResult := elems.Call(nil)
		for _, res := range fnResult {
			walk(res.Interface(), fn)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	elems := reflect.ValueOf(x)

	if elems.Kind() == reflect.Ptr {
		elems = elems.Elem()
	}

	return elems
}
