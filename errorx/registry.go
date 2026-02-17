package errorx

import (
	"fmt"
	"sort"
	"sync"
)

type CodeMeta struct {
	Code int
	Msg  string
}

type registry struct {
	mu    sync.RWMutex
	codes map[int]*CodeMeta
}

var globalRegistry = &registry{
	codes: make(map[int]*CodeMeta),
}

func Register(code int, key, defaultMsg string) error {
	globalRegistry.mu.Lock()
	defer globalRegistry.mu.Unlock()

	if _, exists := globalRegistry.codes[code]; exists {
		return fmt.Errorf("errorx: duplicate code %d", code)
	}

	globalRegistry.codes[code] = &CodeMeta{
		Code: code,
		Msg:  defaultMsg,
	}

	return nil
}

func MustRegister(code int, key, defaultMsg string) {
	if err := Register(code, key, defaultMsg); err != nil {
		panic(err)
	}
}

func GetMeta(code int) (*CodeMeta, bool) {
	globalRegistry.mu.RLock()
	defer globalRegistry.mu.RUnlock()

	meta, ok := globalRegistry.codes[code]
	return meta, ok
}

func ListCodes() []*CodeMeta {
	globalRegistry.mu.RLock()
	defer globalRegistry.mu.RUnlock()

	list := make([]*CodeMeta, 0, len(globalRegistry.codes))
	for _, v := range globalRegistry.codes {
		list = append(list, v)
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].Code < list[j].Code
	})

	return list
}
