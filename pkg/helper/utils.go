package helper

// If returns vt if cond is true, vf otherwise.
func If[T any](cond bool, vt, vf T) T {
	if cond {
		return vt
	}
	return vf
}

// Coalesce returns v if not empty, d otherwise
func Coalesce[T string](v, d T) T {
	if len(v) == 0 {
		return d
	}
	return v
}
