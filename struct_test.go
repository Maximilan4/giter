package giter

import (
	"maps"
	"reflect"
	"slices"
	"testing"
)

func TestFieldOfType(t *testing.T) {
	type TestStruct struct {
		A, B int
		C    float32 `giter:"-"`
		d    string
	}

	typ := reflect.TypeOf(TestStruct{})
	iter := fieldOfType(typ, FieldsIterOpts{})
	for field := range iter {
		if field.Name == "C" {
			t.Fatal("C field must be excluded from iteration")
		}

		expField, found := typ.FieldByName(field.Name)
		if !found {
			t.Errorf("field %s not found", field.Name)
		}

		if !reflect.DeepEqual(field, &expField) {
			t.Errorf("field %s mismatch", field.Name)
		}
	}
}

func TestFieldOfTypeWithUnexported(t *testing.T) {
	type TestStruct struct {
		A, B int
		C    float32 `giter:"-"`
		d    string
	}

	typ := reflect.TypeOf(TestStruct{})
	iter := fieldOfType(typ, FieldsIterOpts{WithNonExportedFields: true})
	for field := range iter {
		if field.Name == "C" {
			t.Fatal("C field must be excluded from iteration")
		}

		expField, found := typ.FieldByName(field.Name)
		if !found {
			t.Errorf("field %s not found", field.Name)
		}

		if !reflect.DeepEqual(field, &expField) {
			t.Errorf("field %s mismatch", field.Name)
		}
	}
}

func TestFieldOfTypeWithRecursive(t *testing.T) {
	type InnerStruct struct {
		F, G int
	}

	type TestStruct struct {
		A, B int
		C    float32 `giter:"-"`
		d    string
		E    *InnerStruct
	}

	typ := reflect.TypeOf(TestStruct{})
	iter := fieldOfType(typ, FieldsIterOpts{Recursive: true})

	expNames := map[string]int{
		"A": 0,
		"B": 0,
		"E": 0,
		"F": 0,
		"G": 0,
	}
	for field := range iter {
		if _, found := expNames[field.Name]; !found {
			t.Errorf("field %s unexpected in iteration", field.Name)
		}

		expNames[field.Name]++
	}

	for k, v := range expNames {
		if v == 0 {
			t.Errorf("field %s is required for iteration", k)
		}
	}
}

func TestStructFieldsValuesOfType(t *testing.T) {
	type InnerStruct struct {
		F, G int
	}

	type TestStruct struct {
		A, B int
		C    float32 `giter:"-"`
		d    string
		E    *InnerStruct
	}

	v := reflect.ValueOf(TestStruct{
		A: 1,
		B: 2,
		C: float32(0.64),
		d: "nonexported",
		E: &InnerStruct{
			F: 1,
			G: 2,
		},
	})
	itr := valuesOf(v, FieldsIterOpts{Recursive: true})
	exp := []any{1, 2, 1, 2}
	got := slices.Collect(itr)

	if !reflect.DeepEqual(exp, got) {
		t.Errorf("incorrect iter results - exp %v, got %v", exp, got)
	}
}

func Test_fieldValuesSeq2(t *testing.T) {
	type InnerStruct struct {
		F, G int
	}

	type TestStruct struct {
		A, B int
		C    float32 `giter:"-"`
		d    string
		E    *InnerStruct
	}

	v := reflect.ValueOf(TestStruct{
		A: 1,
		B: 2,
		C: float32(0.64),
		d: "nonexported",
		E: &InnerStruct{
			F: 1,
			G: 2,
		},
	})

	itr := fieldValuesSeq2(v, FieldsIterOpts{Recursive: true})
	exp := map[string]any{
		"A":   1,
		"B":   2,
		"E.F": 1,
		"E.G": 2,
	}
	got := maps.Collect(itr)

	if !reflect.DeepEqual(exp, got) {
		t.Errorf("incorrect iter results - exp %v, got %v", exp, got)
	}
}

func TestStructFieldsNamesValuesOfTypeErr(t *testing.T) {
	type tc struct {
		name   string
		v      any
		expErr bool
	}

	tcs := []tc{
		{"int", 1, true},
		{"str", "123", true},
		{"float", 0.6, true},
		{"complex", complex(0.6, 0.7), true},
		{"chan", make(chan int), true},
		{"slice", []int{1}, true},
		{"map", map[int]int{1: 1}, true},
		{"struct", struct{ A string }{}, false},
		{"struct2", &struct{ A string }{}, false},
	}

	for _, testcase := range tcs {
		t.Run(testcase.name, func(t *testing.T) {
			v := reflect.ValueOf(testcase.v)
			_, err := StructFieldsNamesValuesOfType(v)
			if (err == nil) == testcase.expErr {
				t.Errorf("err expected, got nil")
			}
		})
	}
}

func TestStructFieldsValuesOfTypeErr(t *testing.T) {
	type tc struct {
		name   string
		v      any
		expErr bool
	}

	tcs := []tc{
		{"int", 1, true},
		{"str", "123", true},
		{"float", 0.6, true},
		{"complex", complex(0.6, 0.7), true},
		{"chan", make(chan int), true},
		{"slice", []int{1}, true},
		{"map", map[int]int{1: 1}, true},
		{"struct", struct{ A string }{}, false},
		{"struct2", &struct{ A string }{}, false},
	}

	for _, testcase := range tcs {
		t.Run(testcase.name, func(t *testing.T) {
			v := reflect.ValueOf(testcase.v)
			_, err := StructFieldsValuesOfType(v)
			if (err == nil) == testcase.expErr {
				t.Errorf("err expected, got nil")
			}
		})
	}
}

func TestStructFieldOfTypeErr(t *testing.T) {
	type tc struct {
		name   string
		v      any
		expErr bool
	}

	tcs := []tc{
		{"int", 1, true},
		{"str", "123", true},
		{"float", 0.6, true},
		{"complex", complex(0.6, 0.7), true},
		{"chan", make(chan int), true},
		{"slice", []int{1}, true},
		{"map", map[int]int{1: 1}, true},
		{"struct", struct{ A string }{}, false},
		{"struct2", &struct{ A string }{}, false},
	}

	for _, testcase := range tcs {
		t.Run(testcase.name, func(t *testing.T) {
			v := reflect.ValueOf(testcase.v)
			_, err := StructFieldOfType(v.Type())
			if (err == nil) == testcase.expErr {
				t.Errorf("err expected, got nil")
			}
		})
	}
}

func TestMustStructFieldOfTypeErr(t *testing.T) {
	type tc struct {
		name   string
		v      any
		expErr bool
	}

	tcs := []tc{
		{"int", 1, true},
		{"str", "123", true},
		{"float", 0.6, true},
		{"complex", complex(0.6, 0.7), true},
		{"chan", make(chan int), true},
		{"slice", []int{1}, true},
		{"map", map[int]int{1: 1}, true},
		{"struct", struct{ A string }{}, false},
		{"struct2", &struct{ A string }{}, false},
	}

	for _, testcase := range tcs {
		t.Run(testcase.name, func(t *testing.T) {
			v := reflect.ValueOf(testcase.v)
			defer func() {
				aerr := recover()

				if (aerr == nil) == testcase.expErr {
					t.Errorf("err expected, got nil")
				}
			}()
			_ = MustStructFieldOfType(v.Type())

		})
	}
}
