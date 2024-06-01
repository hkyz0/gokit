package slice

import (
	"github.com/hkyz0/gokit/internal/errs"
)

// Delete 删除 index 处的元素
func Delete[T any](slice []T, index int) ([]T, error) {
	if index < 0 || len(slice) <= index {
		return slice, errs.NewErrIndexOutOfRang(index)
	}
	return append(slice[:index], slice[index+1:]...), nil
}

// FilterDelete 删除符合条件的元素
// 考虑到性能问题，所有操作都会在原切片上进行
// 被删除元素之后的元素会往前移动，有且只会移动一次
func FilterDelete[T any](slice []T, filter func(idx int, t T) bool) []T {
	// 记录被删除的元素位置，也称空缺的位置
	emptyIndex := 0
	for idx := range slice {
		// 判断是否满足删除的条件
		if filter(idx, slice[idx]) {
			continue
		}
		// 移动元素
		slice[emptyIndex] = slice[idx]
		emptyIndex++
	}
	return slice[:emptyIndex]
}
