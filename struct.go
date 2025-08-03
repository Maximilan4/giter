package giter

import (
	"fmt"
	"iter"
	"reflect"
	"strings"

	"github.com/Maximilan4/gentypes"
)

type (
	FieldsIterOpts struct {
		WithNonExportedFields bool
		Recursive             bool
		LookupTag             string
	}
)

// WithLookupTag - struct iterator option for customizing lookup struct tag
func WithLookupTag(tag string) gentypes.GenericOption[FieldsIterOpts] {
	return func(opts *FieldsIterOpts) {
		opts.LookupTag = tag
	}
}

// WithRecursive - struct iterator option - allows to iterate over nested structs fields/values
func WithRecursive() gentypes.GenericOption[FieldsIterOpts] {
	return func(opts *FieldsIterOpts) {
		opts.Recursive = true
	}
}

// WithNonExportedFields - struct iterator option - allow to iterate over private fields
func WithNonExportedFields() gentypes.GenericOption[FieldsIterOpts] {
	return func(opts *FieldsIterOpts) {
		opts.WithNonExportedFields = true
	}
}

// MustStructFieldsFor - like StructFieldsFor, but panics on err
func MustStructFieldsFor[T any](options ...gentypes.GenericOption[FieldsIterOpts]) iter.Seq[*reflect.StructField] {
	i, err := StructFieldsFor[T](options...)
	if err != nil {
		panic(err)
	}

	return i
}

// StructFieldsFor - return iterator over *reflect.StructField for given type through generic
func StructFieldsFor[T any](options ...gentypes.GenericOption[FieldsIterOpts]) (iter.Seq[*reflect.StructField], error) {
	return StructFieldOfType(reflect.TypeFor[T](), options...)
}

// MustStructFieldsOf - like StructFieldsOf, but panics on err
func MustStructFieldsOf[T any](s T, options ...gentypes.GenericOption[FieldsIterOpts]) iter.Seq[*reflect.StructField] {
	i, err := StructFieldsOf(s, options...)
	if err != nil {
		panic(err)
	}

	return i
}

// StructFieldsOf - return iterator over *reflect.StructField for given struct
func StructFieldsOf[T any](
		s T,
		options ...gentypes.GenericOption[FieldsIterOpts],
) (iter.Seq[*reflect.StructField], error) {
	v := reflect.Indirect(reflect.ValueOf(s))
	return StructFieldOfType(v.Type(), options...)
}

// StructFieldsValuesOf - return iterator over field values of given struct
func StructFieldsValuesOf[T any](s T, options ...gentypes.GenericOption[FieldsIterOpts]) (iter.Seq[any], error) {
	v := reflect.Indirect(reflect.ValueOf(s))
	return StructFieldsValuesOfType(v, options...)
}

// MustStructFieldOfType - like StructFieldOfType, but panics on err
func MustStructFieldOfType(
		t reflect.Type,
		options ...gentypes.GenericOption[FieldsIterOpts],
) iter.Seq[*reflect.StructField] {
	i, err := StructFieldOfType(t, options...)
	if err != nil {
		panic(err)
	}

	return i
}

// StructFieldOfType - return iterator over *reflect.StructField by given reflect.Type
func StructFieldOfType(
		sType reflect.Type,
		options ...gentypes.GenericOption[FieldsIterOpts],
) (iter.Seq[*reflect.StructField], error) {
	if sType.Kind() == reflect.Pointer {
		sType = sType.Elem()
	}

	if sType.Kind() != reflect.Struct {
		return nil, fmt.Errorf("struct expected, given %T", sType.Name())
	}

	opts := FieldsIterOpts{LookupTag: lookupStructTag}
	for _, opt := range options {
		opt(&opts)
	}

	return fieldOfType(sType, opts), nil
}

// StructFieldsValuesOfType - return iterator over fields values by given reflect.Value
func StructFieldsValuesOfType(
		v reflect.Value,
		options ...gentypes.GenericOption[FieldsIterOpts],
) (iter.Seq[any], error) {
	v = reflect.Indirect(v)
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("struct expected, given %T", v.Type().Name())
	}

	opts := FieldsIterOpts{LookupTag: lookupStructTag}
	for _, opt := range options {
		opt(&opts)
	}

	return valuesOf(v, opts), nil
}

// StructFieldsNamesValuesOfType - return iterator over fields values by given reflect.Value
func StructFieldsNamesValuesOfType(
		v reflect.Value,
		options ...gentypes.GenericOption[FieldsIterOpts],
) (iter.Seq2[string, any], error) {
	v = reflect.Indirect(v)
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("struct expected, given %T", v.Type().Name())
	}

	opts := FieldsIterOpts{LookupTag: lookupStructTag}
	for _, opt := range options {
		opt(&opts)
	}

	return fieldValuesSeq2(v, opts), nil
}

func fieldOfType(sType reflect.Type, opts FieldsIterOpts) iter.Seq[*reflect.StructField] {
	return func(yield func(*reflect.StructField) bool) {
		var (
			kind reflect.Kind
			tag  = opts.LookupTag
		)

		if tag == "" {
			tag = lookupStructTag
		}

		for i := 0; i < sType.NumField(); i++ {
			field := sType.Field(i)
			if !field.IsExported() && !opts.WithNonExportedFields {
				continue
			}

			if tagv, ok := field.Tag.Lookup(tag); ok && tagv == "-" {
				continue
			}

			if !yield(&field) {
				return
			}

			if !opts.Recursive {
				continue
			}

			typ := field.Type
			kind = typ.Kind()
			if kind == reflect.Pointer {
				typ = typ.Elem()
				kind = typ.Kind()
			}

			if kind != reflect.Struct {
				continue
			}

			next, done := iter.Pull(fieldOfType(typ, opts))
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

func valuesOf(sValue reflect.Value, opts FieldsIterOpts) iter.Seq[any] {
	return func(yield func(any) bool) {
		var (
			tag   = opts.LookupTag
			ftype reflect.Type
			kind  reflect.Kind
		)

		if tag == "" {
			tag = lookupStructTag
		}

		for i := 0; i < sValue.NumField(); i++ {
			v := reflect.Indirect(sValue.Field(i))
			field := sValue.Type().Field(i)
			if !field.IsExported() {
				continue
			}

			if tagv, ok := field.Tag.Lookup(tag); ok && tagv == "-" {
				continue
			}

			ftype = field.Type
			kind = ftype.Kind()
			if kind == reflect.Pointer {
				ftype = ftype.Elem()
				kind = ftype.Kind()
			}

			if kind == reflect.Struct && opts.Recursive {
				next, done := iter.Pull(valuesOf(v, opts))
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
				continue
			}

			if !yield(v.Interface()) {
				return
			}
		}
	}
}

func fieldValuesSeq2(sValue reflect.Value, opts FieldsIterOpts) iter.Seq2[string, any] {
	return func(yield func(string, any) bool) {
		var (
			tag   = opts.LookupTag
			ftype reflect.Type
			kind  reflect.Kind
		)

		if tag == "" {
			tag = lookupStructTag
		}

		for i := 0; i < sValue.NumField(); i++ {
			v := reflect.Indirect(sValue.Field(i))
			field := sValue.Type().Field(i)
			if !field.IsExported() {
				continue
			}

			var fName string
			if tagv, ok := field.Tag.Lookup(tag); ok && tagv == "-" {
				continue
			} else if tagv != "" {
				fName = tagv
			} else {
				fName = field.Name
			}

			ftype = field.Type
			kind = ftype.Kind()
			if kind == reflect.Pointer {
				ftype = ftype.Elem()
				kind = ftype.Kind()
			}

			if kind == reflect.Struct && opts.Recursive {
				next, done := iter.Pull2(fieldValuesSeq2(v, opts))
				var buf strings.Builder
				for {
					nfName, nfValue, found := next()
					if !found {
						break
					}

					expLen := len(fName) + len(nfName) + 1
					if bufCap := buf.Cap(); bufCap < expLen {
						buf.Grow(expLen - bufCap)
					}
					buf.WriteString(fName)
					buf.WriteByte('.')
					buf.WriteString(nfName)

					if !yield(buf.String(), nfValue) {
						return
					}
					buf.Reset()
				}

				done()
				continue
			}

			if !yield(fName, v.Interface()) {
				return
			}
		}
	}
}
