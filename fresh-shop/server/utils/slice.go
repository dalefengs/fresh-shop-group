package utils

// SliceDifference 求差集
// s1 中有的但 s2 中没有的
func SliceDifference[T comparable](s1, s2 []T) []T {
	// 创建一个 map 用于存储 s2 中的元素。
	m := make(map[T]bool, len(s2))
	for _, v := range s2 {
		m[v] = true
	}
	var result []T
	for _, v := range s1 {
		// 如果没有
		if !m[v] {
			result = append(result, v)
		}
	}
	return result
}

// SliceUnionSet 求并集
// s1 和 s2 都有的元素
func SliceUnionSet[T comparable](s1, s2 []T) []T {
	m := make(map[T]bool)
	for _, v := range s1 {
		m[v] = true
	}
	var result []T
	for _, v := range s2 {
		if m[v] {
			result = append(result, v)
		}
	}
	return result
}
