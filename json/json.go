package json

import (
	"bytes"
	"encoding/json"
)

// 格式化
func Format(input string) (string, error) {
	if input == "" {
		return "", nil
	}

	var buf bytes.Buffer
	buf.Grow(len(input) + len(input)/2)
	if err := json.Indent(&buf, []byte(input), "", "  "); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// 压缩
func Compress(input string) (string, error) {
	var obj interface{}
	err := json.Unmarshal([]byte(input), &obj)
	if err != nil {
		return "", err
	}

	compressed, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}

	return string(compressed), nil
}
