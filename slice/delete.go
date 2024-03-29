package slice

import "github.com/zht-account/gotools"

// Delete 根据下标删除切片中元素
func Delete[T any](index int, s []T) (res []T, t T, err error) {
	if index < 0 || index >= len(s) {
		return s, t, gotools.NewErrIndexOutOfRange(len(s), index)
	}
	t = s[index]
	return FilterDelete(s, func(key int, value T) bool {
		if key == index {
			return true
		}
		return false
	}), t, nil
}

func FilterDelete[T any](s []T, filter func(key int, value T) bool) []T {
	i := 0
	for k, v := range s {
		if filter(k, v) {
			continue
		}
		s[i] = v
		i++
	}
	return s[:i]
}
