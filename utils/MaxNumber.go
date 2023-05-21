package utils

func MaxNumber[T Comparable](a, b T) T {
	if a > b {
		return a
	}

	return b
}

type Comparable interface {
	int | float64 | float32 | string | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}
