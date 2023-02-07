package permutation

// 统计 conf 配置的组合情况在 result 里出现的次数
func CountItemNum[T comparable](result [][]T, conf map[T]int) int {
	var count int
	for _, r := range result {
		c := make(map[T]int)
		for _, v := range r {
			c[v]++
		}

		ok := true
		for k, n := range conf {
			if n != c[k] {
				ok = false
			}
		}
		if ok {
			count++
		}
	}
	return count
}
