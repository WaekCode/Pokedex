package pokecache

import (
	"sync"
	"time"


)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
	
}


type Cache struct{

	data map[string]cacheEntry
	mu sync.Mutex

}

func (c *Cache) Get(key string) ([]byte,bool)  {
	c.mu.Lock()
	defer c.mu.Unlock()

	data,ok := c.data[key]
	if !ok{
		return []byte{},false
	
	}
	return data.val,true

}



func (c *Cache) Add(key string,value []byte)  {
	c.mu.Lock()
	entry := cacheEntry{}
	entry.val = value
	entry.createdAt = time.Now()
	c.data[key] = entry
	 
	c.mu.Unlock()
}


func (c *Cache) reapLoop(interval time.Duration)  {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key,value := range c.data{
		if time.Since(value.createdAt) > interval{
			delete(c.data,key)
		}
	}
}


func NewCache(interval time.Duration) *Cache{
	inti := make(map[string]cacheEntry)
	c := &Cache{
		data : inti,
	}

	ticker := time.NewTicker(interval)

	go func() {
		for range ticker.C {
			c.reapLoop(interval)

		}
	}()



	return c

}


