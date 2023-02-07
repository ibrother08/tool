package permutation

import (
	"math"
)

// 列出用 list 中所有项，组成 deep 长度的数组的所有可能。
// 在 list 中的每一项都可以重复出现。
func FullPermutation[T comparable](list []T, deep int) [][]T {
	length := float64(len(list))
	max := int(math.Pow(length, float64(deep)))
	result := make([][]T, max)

	for n := 0; n < max; n++ {
		result[n] = make([]T, deep)
		for lv := 0; lv < deep; lv++ {
			// A % B = (A + B * K) % B
			i := n / int(math.Pow(length, float64(deep-1-lv))) % int(length)
			result[n][lv] = list[i]
		}
	}

	return result
}
