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

func ForEach[T any](arr []T, callback func(T)) {
	for _, v := range arr {
		callback(v)
	}
}

func Pop[T any](arr []T) (T, []T) {
	return arr[len(arr)-1], arr[:len(arr)-1]
}

func Shift[T any](arr []T) (T, []T) {
	return arr[0], arr[1:]
}

func Unshift[T any](arr []T, v T) []T {
	return append([]T{v}, arr...)
}

func Join[T any](arr []T, sep string) string {
	var r string
	for i := range arr {
		if i > 0 {
			r += sep
		}
	}
	return r
}

func Concat[T any](arr ...[]T) []T {
	var r []T
	for _, v := range arr {
		r = append(r, v...)
	}
	return r
}

func Includes[T comparable](arr []T, v T) bool {
	for _, vv := range arr {
		if vv == v {
			return true
		}
	}
	return false
}
