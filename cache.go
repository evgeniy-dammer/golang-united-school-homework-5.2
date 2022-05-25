package cache

import (
	"time"
)

type Cache struct {
	cacheItems map[string]Item
}

type Item struct {
	itemValue  string
	itemExpire time.Time
}

func NewCache() Cache {
	return Cache{
		cacheItems: make(map[string]Item),
	}
}

func (cache *Cache) Get(key string) (string, bool) {
	if item, ok := cache.cacheItems[key]; ok {
		if item.itemExpire.IsZero() || time.Now().Before(item.itemExpire) {
			return item.itemValue, true
		}
	}

	return "", false
}

func (cache *Cache) Put(key, value string) {
	cache.cacheItems[key] = Item{itemValue: value}
}

func (cache Cache) Keys() []string {
	var result []string

	for key := range cache.cacheItems {
		if _, ok := cache.Get(key); ok {
			result = append(result, key)
		}
	}

	return result
}

func (cache *Cache) PutTill(key, value string, deadline time.Time) {
	cache.cacheItems[key] = Item{itemValue: value, itemExpire: deadline}
}
