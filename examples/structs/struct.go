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
	for field := range giter.MustFieldsFor[TestStruct](true) {
		fmt.Println(field)
	}
}
