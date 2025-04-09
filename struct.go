package giter

import (
	"fmt"
	"iter"
	"reflect"
)

type (
	FieldsIterOpts struct {
		WithNonExportedFields bool
		Recursive             bool
		LookupTag             string
	}
)

func WithLookupTag(tag string) Option[*FieldsIterOpts] {
	return func(opts *FieldsIterOpts) {
		opts.LookupTag = tag
	}
}

func WithRecursive() Option[*FieldsIterOpts] {
	return func(opts *FieldsIterOpts) {
		opts.Recursive = true
	}
}

func WithNonExportedFields() Option[*FieldsIterOpts] {
	return func(opts *FieldsIterOpts) {
		opts.WithNonExportedFields = true
	}
}

func MustFieldsFor[T any](options ...Option[*FieldsIterOpts]) iter.Seq[*reflect.StructField] {
	i, err := FieldsFor[T](options...)
	if err != nil {
		panic(err)
	}

	return i
}

func FieldsFor[T any](options ...Option[*FieldsIterOpts]) (iter.Seq[*reflect.StructField], error) {
	return FieldOfType(reflect.TypeFor[T](), options...)
}

func MustFieldsOf[T any](s T, options ...Option[*FieldsIterOpts]) iter.Seq[*reflect.StructField] {
	i, err := FieldsOf(s, options...)
	if err != nil {
		panic(err)
	}

	return i
}

func FieldsOf[T any](s T, options ...Option[*FieldsIterOpts]) (iter.Seq[*reflect.StructField], error) {
	v := reflect.Indirect(reflect.ValueOf(s))
	return FieldOfType(v.Type(), options...)
}

func MustFieldOfType(t reflect.Type, options ...Option[*FieldsIterOpts]) iter.Seq[*reflect.StructField] {
	i, err := FieldOfType(t, options...)
	if err != nil {
		panic(err)
	}

	return i
}

func FieldOfType(sType reflect.Type, options ...Option[*FieldsIterOpts]) (iter.Seq[*reflect.StructField], error) {
	if sType.Kind() != reflect.Struct {
		return nil, fmt.Errorf("struct expected, given %T", sType.Name())
	}

	opts := FieldsIterOpts{LookupTag: lookupStructTag}
	for _, opt := range options {
		opt(&opts)
	}

	return fieldOfType(sType, opts), nil
}

func fieldOfType(sType reflect.Type, opts FieldsIterOpts) iter.Seq[*reflect.StructField] {
	return func(yield func(*reflect.StructField) bool) {
		var (
			field reflect.StructField
			kind  reflect.Kind
		)

		for i := 0; i < sType.NumField(); i++ {
			field = sType.Field(i)
			if !field.IsExported() && !opts.WithNonExportedFields {
				continue
			}

			if tag, ok := field.Tag.Lookup("giter"); ok && tag == "-" {
				continue
			}

			if !yield(&field) {
				return
			}

			if !opts.Recursive {
				continue
			}

			kind = field.Type.Kind()
			if kind == reflect.Pointer {
				kind = field.Type.Elem().Kind()
			}

			if kind != reflect.Struct {
				continue
			}

			next, done := iter.Pull(fieldOfType(field.Type.Elem(), opts))
			for {
				f, found := next()
				if !found {
					break
				}

				if !yield(f) {
					return
				}
			}

			done()
		}
	}
}
