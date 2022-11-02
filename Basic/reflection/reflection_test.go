package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	expected := "Chris"
	age := 18

	cases := []struct {
		Name        string
		Input       interface{}
		ExpectCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
				Age  int
			}{expected, age},
			[]string{expected},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			// walk(test.Input, func(input string) {
			// 	got = append(got, input)
			// })

			walkRefactor(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectCalls)
			}

		})
	}
}

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	numbersOfValues := 0
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		numbersOfValues = val.NumField()
		getField = val.Field
	case reflect.Slice:
		numbersOfValues = val.Len()
		getField = val.Index
	}

	for i := 0; i < numbersOfValues; i++ {
		walk(getField(i).Interface(), fn)

	}

}

func walkOptimized(x interface{}, fn func(input string)) {
	val := getValue(x)
	numberOfValues := 0
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		numberOfValues = val.NumField()
		getField = val.Field
	case reflect.Slice, reflect.Array:
		numberOfValues = val.Len()
		getField = val.Index
	}

	for i := 0; i < numberOfValues; i++ {
		walkOptimized(getField(i).Interface(), fn)
	}

}

func walkWithMap(x interface{}, fn func(input string)) {
	val := getValue(x)

	numberOfValues := 0
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		numberOfValues = val.NumField()
		getField = val.Field
	case reflect.Slice, reflect.Array:
		numberOfValues = val.Len()
		getField = val.Index
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkWithMap(val.MapIndex(key).Interface(), fn)
		}
	}

	for i := 0; i < numberOfValues; i++ {
		walkWithMap(getField(i).Interface(), fn)
	}

}

func walkRefactor(x interface{}, fn func(input string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	}
}

func TestWalkWithStruct(t *testing.T) {
	expected := "Chris"
	cases := []struct {
		Name        string
		Input       interface{}
		ExpectCalls []string
	}{
		{
			"struct with one string field",
			Person{
				expected,
				Profile{
					18,
					"Shanghai",
				},
			},
			[]string{expected},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectCalls)
			}

		})
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
