package jsonx

func DeepCopy[T any](src T) (T, error) {
	var dst T
	b, err := Marshal(src)
	if err != nil {
		return dst, err
	}
	err = Unmarshal(b, &dst)
	return dst, err
}
