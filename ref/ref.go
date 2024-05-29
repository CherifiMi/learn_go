package ref

import "reflect"

func walk(x interface{}, fn func(string)) {
	val := reflect.ValueOf(x)
	for i := 0; i < val.NumField(); i++ {
		s := val.Field(i)
		if s.Kind() == reflect.String {
			fn(s.String())
		}
	}
}
