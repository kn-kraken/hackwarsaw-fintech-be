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

func MapRef2[T1, T2, V any](t1s []T1, t2s []T2, fn func(*T1, *T2) V) []V {
    result := make([]V, len(t1s))
    for i, t := range t1s {
        result[i] = fn(&t, &t2s[i])
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
