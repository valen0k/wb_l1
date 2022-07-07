package ex07

import "sync"

type ConcurrencyMap struct {
	sync.RWMutex
	data map[int]string
}

func NewConcurrencyMap() *ConcurrencyMap {
	return &ConcurrencyMap{data: make(map[int]string)}
}

func (m *ConcurrencyMap) Set(key int, value string) {
	m.Lock()
	defer m.Unlock()
	m.data[key] = value
}
