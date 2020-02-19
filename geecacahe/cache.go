package geecache

import (
	"geecache/geecacahe/lru"
	"sync"
)

type cacahe struct {
	mu          sync.Mutex
	lru         *lru.Cache
	cacaheBytes int64
}

func (c *cacahe) add(key string, value ByteView) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		c.lru = lru.New(c.cacaheBytes, nil)
	}
	c.lru.Add(key, value)
}

func (c *cacahe) get(key string) (value ByteView, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		return
	}

	if v, ok := c.lru.Get(key); ok {
		return v.(ByteView), ok
	}
	return
}
