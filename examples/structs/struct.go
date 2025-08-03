package main

import (
	"fmt"

	"github.com/Maximilan4/giter"
)

type TestStruct struct {
	A int
	B string
	C bool
	D struct {
		A, B string
	}
	e float64
	F complex64 `giter:"-"`
}

func main() {
	for field := range giter.MustStructFieldsFor[TestStruct](giter.WithNonExportedFields(), giter.WithRecursive()) {
		fmt.Println(field.Name)
	}
}
