package jsonx

import (
	"encoding/json"
)

func Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

func MustMarshal(v any) []byte {
	b, err := Marshal(v)
	if err != nil {
		panic(err)
	}
	return b
}

func MustUnmarshal(data []byte, v any) {
	if err := Unmarshal(data, v); err != nil {
		panic(err)
	}
}

func MustMarshalToString(v any) string {
	return string(MustMarshal(v))
}

func MarshalToString(v any) (string, error) {
	b, err := Marshal(v)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func ToMap(v any) (map[string]any, error) {
	b, err := Marshal(v)
	if err != nil {
		return nil, err
	}
	var m map[string]any
	err = Unmarshal(b, &m)
	return m, err
}
