package cache

import "time"

type dataElement struct {
	value     string
	timestamp time.Time
	elapsable bool
}

type Cache struct {
	cache map[string]dataElement
}

func NewCache() Cache {
	var m = make(map[string]dataElement)

	return Cache{m}
}

func (a Cache) Get(key string) (string, bool) {
	return "", false
}

func (a Cache) Put(key, value string) {
	var temp = dataElement{value, time.Now(), false}
	a.cache[key] = temp
}

func (a Cache) Keys() []string {
	var out []string
	for _, v := range a.cache {
		if v.timestamp.Before(time.Now()) {
			out = append(out, v.value)
		}
	}
	return out
}

func (a *Cache) PutTill(key, value string, deadline time.Time) {
	var temp = dataElement{value, deadline, true}
	a.cache[key] = temp
}
