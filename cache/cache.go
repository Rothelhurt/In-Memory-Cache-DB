package cache

import (
	"context"
	"fmt"
	"time"
)

type Cache struct {
	db map[string]context.Context
}

func New() Cache {
	db := make(map[string]context.Context)
	return Cache{
		db: db,
	}
}

func (c *Cache) Set(key string, value interface{}, ttl int) {

	c.db[key], _ = context.WithTimeout(context.Background(), time.Duration(ttl)*time.Second)
	c.db[key] = context.WithValue(c.db[key], key, value)
}

func (c Cache) Get(key string) interface{} {
	if c.db[key] == nil {
		return "There is no such key in DB."
	}
	select {
	case <-c.db[key].Done():
		delete(c.db, key)
		return "Key was removed by timeout."
	default:
		return c.db[key].Value(key)
	}
}

func (c Cache) Delete(key string) interface{} {
	if c.db[key] != nil {
		delete(c.db, key)
		return "Key was deleted."
	}
	return "There is no such key in DB."

}

func (c Cache) ShowDB() {
	for key := range c.db {
		fmt.Printf("Key: %s, Value: %d\n", key, c.db[key].Value(key))
	}
}
