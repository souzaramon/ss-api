package util

import "reflect"

func IsEmptyStruct(dto any) bool {
	v := reflect.ValueOf(dto)

	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		if field.Kind() == reflect.Pointer && !field.IsNil() {
			return false
		}
	}

	return true
}
