package ex11

func IntersectionSets[T comparable](left, right []T) []T {
	res := make([]T, 0)
	buf := make(map[T]struct{})

	for _, val := range left {
		buf[val] = struct{}{}
	}

	for _, val := range right {
		if _, ok := buf[val]; ok {
			res = append(res, val)
		}
	}

	return res
}
