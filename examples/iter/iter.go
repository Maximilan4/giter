package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/Maximilan4/giter"
)

func main() {
	t := slices.Values([]string{"pfA", "pfB", "sufC"})

	t2 := giter.Iterate(t).
			Filter(func(v string) bool {
				return strings.HasPrefix(v, "pf")
			}).
			Map(func(v string) string {
				return strings.ToUpper(v)
			}).
			Reduce("test:", func(start, v string) string {
				start += v
				return start
			})
	fmt.Println(t2)
}
