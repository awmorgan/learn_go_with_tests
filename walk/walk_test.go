package walk

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	type Profile struct {
		Age  int
		City string
	}
	type Person struct {
		Name    string
		Profile Profile
	}
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
		{
			"Nested fields",
			Person{
				"Bill", Profile{33, "London"}},
			[]string{"Bill", "London"},
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
