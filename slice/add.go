package slice

import (
	"github.com/hkyz0/gokit/internal/errs"
)

// Add 在index处添加元素element
// index 范围应为[0, len(src)]
// 如果index == len(src) 则表示往末尾添加元素
func Add[T any](slice []T, element T, index int) ([]T, error) {
	if index < 0 || index > len(slice) {
		return nil, errs.NewErrIndexOutOfRang(index)
	}
	return append(slice[:index], append([]T{element}, slice[index:]...)...), nil
}
