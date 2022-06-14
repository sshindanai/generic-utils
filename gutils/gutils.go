package gutils

func Map[T, RETURN any](arr []T, callback func(T) RETURN) []RETURN {
	var r []RETURN
	for _, v := range arr {
		r = append(r, callback(v))
	}
	return r
}

func Reduce[T, RETURN any](arr []T, callback func(RETURN, T) RETURN, init RETURN) RETURN {
	var r RETURN
	for _, v := range arr {
		r = callback(r, v)
	}
	return r
}

func Filter[T comparable](arr []T, callback func(T) bool) []T {
	var r []T
	for _, v := range arr {
		if callback(v) {
			r = append(r, v)
		}
	}
	return r
}

func Find[T comparable](arr []T, callback func(T) bool) *T {
	for _, v := range arr {
		if callback(v) {
			return &v
		}
	}
	return nil
}

func FindIndex[T comparable](arr []T, callback func(T) bool) int {
	for i, v := range arr {
		if callback(v) {
			return i
		}
	}
	return -1
}

func Every[T comparable](arr []T, callback func(T) bool) bool {
	for _, v := range arr {
		if !callback(v) {
			return false
		}
	}
	return true
}

func Some[T comparable](arr []T, callback func(T) bool) bool {
	for _, v := range arr {
		if callback(v) {
			return true
		}
	}
	return false
}

func FlatMap[T, RETURN any](arr []T, callback func(T) []RETURN) []RETURN {
	var r []RETURN
	for _, v := range arr {
		r = append(r, callback(v)...)
	}
	return r
}
