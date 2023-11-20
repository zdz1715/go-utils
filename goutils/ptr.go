package goutils

// Ptr returns a pointer to a variable
func Ptr[T any](s T) *T {
	return &s
}

// Val returns the value of the pointer
func Val[T any](s *T) T {
	return *s
}
