package permutation

import (
	"fmt"
)

func ExampleFullPermutation() {
	list := []string{"a", "b"}
	result := FullPermutation(list, 4)
	for _, v := range result {
		fmt.Println(v)
	}
	// Output:
	// [a a a a]
	// [a a a b]
	// [a a b a]
	// [a a b b]
	// [a b a a]
	// [a b a b]
	// [a b b a]
	// [a b b b]
	// [b a a a]
	// [b a a b]
	// [b a b a]
	// [b a b b]
	// [b b a a]
	// [b b a b]
	// [b b b a]
	// [b b b b]
}

func ExampleCountItemNum() {
	list := []string{"a", "b"}
	result := FullPermutation(list, 4)

	n := CountItemNum(result, map[string]int{
		"a": 3,
		"b": 1,
	})

	fmt.Println(n)
	// Output: 4
}
