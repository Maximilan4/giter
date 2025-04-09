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

	for v := range giter.Filter(t, func(v string) bool { return v == "B" }) {
		fmt.Println(v)
	}

	// t2 := giter.Map(t, func(k int, v string) (int, rune) {
	// 	if len(v) != 1 {
	// 		return k, '-'
	// 	}
	//
	// 	return k, rune(v[0])
	// })
	//
	// for k, v := range t2 {
	// 	fmt.Printf("%d - %s\n", k, string(v))
	// }
	//
	// t3 := giter.Reduce(t, 0, func(acc int, _ int, v string) int {
	// 	return acc + int(v[0])
	// })
	// fmt.Println(t3)
}
