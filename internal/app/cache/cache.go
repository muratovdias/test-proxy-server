package cache

import (
	"fmt"
	"sync"

	"github.com/muratovdias/test-proxy-server/internal/app/entities"
)

type Cache struct {
	cache map[string]entities.ProxyResponse
	mutex *sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		cache: make(map[string]entities.ProxyResponse),
		mutex: &sync.RWMutex{},
	}
}

func (c *Cache) Get(key string) (entities.ProxyResponse, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	data, ok := c.cache[key]
	return data, ok
}

func (c *Cache) Set(key string, response entities.ProxyResponse) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cache[key] = response
	fmt.Println("response saved in cache")
}
