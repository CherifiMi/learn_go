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
	val := getVlue(x)

	numOfValue := 0
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Slice:
		numOfValue = val.Len()
		getField = val.Index
	case reflect.Struct:
		numOfValue = val.NumField()
		getField = val.Field
	}

	for i := 0; i < numOfValue; i++ {
		walk(getField(i).Interface(), fn)
	}

}

func getVlue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	return val
}
