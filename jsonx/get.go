package jsonx

func Get(data []byte, key string) (any, error) {
	var m map[string]any
	if err := Unmarshal(data, &m); err != nil {
		return nil, err
	}
	return m[key], nil
}
