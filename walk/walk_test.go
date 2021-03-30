package walk

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct{ Name string }{"Bill"},
			[]string{"Bill"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Bill", "London"},
			[]string{"Bill", "London"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Bill", 51},
			[]string{"Bill"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
