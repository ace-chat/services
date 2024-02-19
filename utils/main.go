package utils

func Deduplicate[T any](arr []T) []T {
	result := make([]T, 0, len(arr))
	temp := map[any]struct{}{}
	for _, item := range arr {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
