package configuration

import (
	"fmt"
	"sync"
)

var (
	c    *Config
	once sync.Once
)

//GetConfiguration singleton pattern implementation
func GetConfiguration() *Config {
	once.Do(func() {
		c = &Config{
			items: make(map[string]string),
		}
	})
	return c
}

//Config configuration structure
type Config struct {
	items map[string]string
	mu    sync.RWMutex
}

//Set function
func (c *Config) Set(key, data string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = data
}

//Get method
func (c *Config) Get(key string) (string, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	item, ok := c.items[key]
	if !ok {
		return "", fmt.Errorf("The '%s' is not presented", key)
	}
	return item, nil
}
