package toolbelt

func Ref[T any](value T) *T {
	return &value
}
