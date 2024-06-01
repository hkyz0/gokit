package slice

import (
	"errors"
	"github.com/hkyz0/gokit"
)

// Sum 求和
func Sum[T gokit.Number](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

// Max 返回最大值
// slice存在为空的情况
func Max[T gokit.RealNumber](slice []T) (T, error) {
	if len(slice) == 0 {
		var max T
		return max, errors.New("slice is empty")
	}
	max := slice[0]
	for _, v := range slice {
		if max < v {
			max = v
		}
	}
	return max, nil
}

// Min 返回最小值
// slice存在为空的情况
func Min[T gokit.RealNumber](slice []T) (T, error) {
	if len(slice) == 0 {
		var min T
		return min, errors.New("slice is empty")
	}
	min := slice[0]
	for _, v := range slice {
		if min > v {
			min = v
		}
	}
	return min, nil
}
