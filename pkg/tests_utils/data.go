package tests_utils

func GetPointer[T any](s T) *T {
	return &s
}
