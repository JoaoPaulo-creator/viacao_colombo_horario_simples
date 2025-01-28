package utils

type constraints interface {
	int | uint | uint8 | uint16 | string
}

func Contains[T constraints](s []T, v any) bool {
	for _, i := range s {
		if i == v {
			return true
		}
	}

	return false
}
