package reflection

import (
	"fmt"
	"reflect"
)

/*
Traverse recursively through a struct and all its fields.
If a field or an array contain a nil ptr, the function
will return true and the name of the field set to nil.
*/
func CheckNilPtr(value reflect.Value, name string) (bool, string) {
	switch value.Kind() {
	case reflect.Ptr:
		if value.IsNil() {
			return true, name
		}
		return CheckNilPtr(value.Elem(), name)
	case reflect.Struct:
		return CheckStruct(value, name)
	case reflect.Slice:
		return CheckArrayOrSlice(value, name)
	case reflect.Array:
		return CheckArrayOrSlice(value, name)
	case reflect.Map:
		return CheckMap(value, name)
	case reflect.Interface:
		return CheckNilPtr(value.Elem(), name)
	}
	return false, ""
}

func CheckMap(value reflect.Value, name string) (bool, string) {
	for _, key := range value.MapKeys() {
		if ok, fieldName := CheckNilPtr(value.MapIndex(key), fmt.Sprintf("%s[%v]", name, key)); ok {
			return true, fieldName
		}
	}
	return  false, ""
}

func CheckStruct(value reflect.Value, _ string) (bool, string) {
	for i := 0; i < value.NumField(); i += 1 {
		if ok, fieldName := CheckNilPtr(value.Field(i), value.Type().Field(i).Name); ok {
			return true, fieldName
		}
	}
	return false, ""
}

func CheckArrayOrSlice(value reflect.Value, name string) (bool, string) {
	for i := 0; i < value.Len(); i += 1 {
		if ok, fieldName := CheckNilPtr(value.Index(i), fmt.Sprintf("%s[%d]", name, i)); ok {
			return true, fieldName
		}
	}
	return false, ""
}
