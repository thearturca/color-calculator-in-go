package utils

func GetValuesSliceFromPointerSlice[T any](s []*T) []T {
	resSlice := make([]T, 0, len(s))

	for _, v := range s {
		resSlice = append(resSlice, *v)
	}

	return resSlice
}
