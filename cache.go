package cache

type Cache struct {
	data map[string]int
}

func New() *Cache {
	return &Cache{
		data: make(map[string]int),
	}
}

func (c *Cache) Set(key string, value int) {
	c.data[key] = value
}

func (c *Cache) Get(key string) int {
	return c.data[key]
}

func (c *Cache) Delete(key string) {
	delete(c.data, key)
}
