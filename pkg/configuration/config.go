package configuration

import (
	"fmt"
	"sync"
)

var (
	c    *configuration
	once sync.Once
)

//GetConfiguration singleton pattern implementation
func GetConfiguration() *configuration {
	once.Do(func() {
		c = &configuration{
			items: make(map[string]string),
		}
	})
	return c
}

type configuration struct {
	items map[string]string
	mu    sync.RWMutex
}

//Set function
func (c *configuration) Set(key, data string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = data
}

//Get method
func (c *configuration) Get(key string) (string, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	item, ok := c.items[key]
	if !ok {
		return "", fmt.Errorf("The '%s' is not presented", key)
	}
	return item, nil
}
