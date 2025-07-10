package ptr

// Of 는 주어진 값의 포인터를 반환합니다.
func Of[T any](v T) *T {
	return &v
}
