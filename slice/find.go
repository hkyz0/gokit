package slice

// Find 查找元素
// 如果没有找到，第二个返回值返回 false
func Find[T any](slice []T, match matchFunc[T]) (T, bool) {
	for _, v := range slice {
		if match(v) {
			return v, true
		}
	}
	var t T
	return t, false
}

// FindAll 查找所有符合条件的元素
func FindAll[T any](slice []T, match matchFunc[T]) []T {
	// 符合条件元素应该是少数,所以会除以 8
	// 也就是触发扩容的情况下，最多三次就会和原本的容量一样
	// +1 是为了保证，容量至少为1
	res := make([]T, 0, len(slice)>>3+1)
	for _, v := range slice {
		if match(v) {
			res = append(res, v)
		}
	}
	return res
}
