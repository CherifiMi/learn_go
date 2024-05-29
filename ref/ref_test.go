package ref

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	cases := []struct {
		Name         string
		Input        interface{}
		ExpectedCall []string
	}{
		{
			"",
			struct {
				Name string
			}{"cheris"},
			[]string{"cheris"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		{
			"nested fields",
			Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(test.ExpectedCall, got) {
				t.Errorf("got %v, want %v", got, test.ExpectedCall)
			}
		})
	}

}