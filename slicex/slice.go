package slicex

func InSlice[T comparable](list []T, target T) bool {
	for _, v := range list {
		if v == target {
			return true
		}
	}
	return false
}
