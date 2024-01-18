package cache

import ("sync"
		"time"
       )
type Cache struct{
	v map[string]cacheEntry
	mu *sync.Mutex
}
type cacheEntry struct {
	val []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) Cache{
	val := Cache{v: map[string]cacheEntry{}, mu: &sync.Mutex{}}
	go val.reapLoop(interval)
	return val

}

func (m *  Cache)reapLoop(interval time.Duration){
	ticker := time.NewTicker(interval)
	for now := range ticker.C {
		cutoff := now.Add(-interval) 
		for key,val := range m.v{
			if val.createdAt.Unix() <= cutoff.Unix(){
				m.mu.Lock()
				delete(m.v, key)
				m.mu.Unlock()
			}
		}
	}
    
}

func (m *  Cache)Get(key string) ([]byte, bool){
	m.mu.Lock()
	defer m.mu.Unlock()
	Val, ok := m.v[key]
	return Val.val, ok
}

func (m * Cache) Put(key string, val []byte){
	m.mu.Lock()
	defer m.mu.Unlock()
	m.v[key] = cacheEntry{val: val, createdAt: time.Now()}
}