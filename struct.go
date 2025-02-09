package giter

import (
	"fmt"
	"iter"
	"reflect"
)

func MustFieldsFor[T any](withNonExported bool) iter.Seq[reflect.StructField] {
	i, err := FieldsFor[T](withNonExported)
	if err != nil {
		panic(err)
	}

	return i
}

func FieldsFor[T any](withNonExported bool) (iter.Seq[reflect.StructField], error) {
	return FieldOfType(reflect.TypeFor[T](), withNonExported)
}

func MustFieldsOf[T any](s T, withNonExported bool) iter.Seq[reflect.StructField] {
	i, err := FieldsOf(s, withNonExported)
	if err != nil {
		panic(err)
	}

	return i
}

func FieldsOf[T any](s T, withNonExported bool) (iter.Seq[reflect.StructField], error) {
	v := reflect.Indirect(reflect.ValueOf(s))
	return FieldOfType(v.Type(), withNonExported)
}

func MustFieldOfType(t reflect.Type, withNonExported bool) iter.Seq[reflect.StructField] {
	i, err := FieldOfType(t, withNonExported)
	if err != nil {
		panic(err)
	}

	return i
}

func FieldOfType(sType reflect.Type, withNonExported bool) (iter.Seq[reflect.StructField], error) {
	if sType.Kind() != reflect.Struct {
		return nil, fmt.Errorf("struct expected, given %T", sType.Name())
	}
	return func(yield func(reflect.StructField) bool) {
		var (
			field reflect.StructField
		)

		for i := 0; i < sType.NumField(); i++ {
			field = sType.Field(i)
			if !field.IsExported() && !withNonExported {
				continue
			}

			if tag, ok := field.Tag.Lookup("giter"); ok && tag == "-" {
				continue
			}

			if !yield(field) {
				return
			}
		}
	}, nil
}
