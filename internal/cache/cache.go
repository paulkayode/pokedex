package cache

import "sync"

type Cache struct{
	V map[string][]byte
	mu sync.RWMutex
}


func (m *  Cache)Get(key string) ([]byte, bool){
	m.mu.RLock()
	defer m.mu.RUnlock()
	val, ok := m.V[key]
	return val, ok
}

func (m * Cache) Put(key string, val []byte){
	m.mu.Lock()
	defer m.mu.Unlock()
	m.V[key] = val
}