package utils

func Map[T, V any](ts []T, fn func(T) V) []V {
    result := make([]V, len(ts))
    for i, t := range ts {
        result[i] = fn(t)
    }
    return result
}

func MapRef[T, V any](ts []T, fn func(*T) V) []V {
    result := make([]V, len(ts))
    for i, t := range ts {
        result[i] = fn(&t)
    }
    return result
}

func Any[T any](ts []T, fn func(T) bool) bool {
    for _, t := range ts {
        if fn(t) {
          return true
        }
    }
    return false
}
