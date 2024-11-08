package cache

type Cache struct {
	db map[string]interface{}
}

func New() Cache {
	db := make(map[string]interface{})
	return Cache{
		db: db,
	}
}

func (c *Cache) Set(key string, value interface{}) {
	c.db[key] = value
}

func (c Cache) Get(key string) interface{} {
	if c.db[key] == nil {
		return "There is no such key in DB."
	}
	return c.db[key]
}

func (c Cache) Delete(key string) interface{} {
	if c.db[key] != nil {
		delete(c.db, key)
		return "Key was deleted."
	}
	return "There is no such key in DB."

}
