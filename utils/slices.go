package utils

func SliceInt64Unique(input []int64) []int64 {
	u := make([]int64, 0, len(input))
	m := make(map[int64]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}

	return u
}

func SliceStringContains(a []string, s string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}

	return false
}
