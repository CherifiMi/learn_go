package ref

import "reflect"

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func walk(x interface{}, fn func(string)) {
	val := reflect.ValueOf(x)
	for i := 0; i < val.NumField(); i++ {
		s := val.Field(i)

		switch s.Kind() {
		case reflect.String:
			fn(s.String())
		case reflect.Struct:
			walk(s.Interface(), fn)
		}
	}
}
