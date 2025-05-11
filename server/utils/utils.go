package utils

func LazyTernary[T any](condition bool, trueValue, falseValue func() T) T {
	if condition {
		return trueValue()
	}

	return falseValue()
}

func Ternary[T any](condition bool, trueValue, falseValue T) T {
	if condition {
		return trueValue
	}

	return falseValue
}
