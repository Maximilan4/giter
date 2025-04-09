package giter

import (
	"reflect"
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
	expNames := map[string]struct{}{
		"A": {},
		"B": {},
		"E": {},
		"F": {},
		"G": {},
	}
	for field := range iter {
		if _, found := expNames[field.Name]; !found {
			t.Errorf("field %s unexpected in iteration", field.Name)
		}
	}
}
