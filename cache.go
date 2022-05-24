package cache

import (
	"time"
)

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
	value, ok := a.cache[key]
	if ok == true {
		if value.elapsable {
			if time.Now().Before(value.timestamp) {
				return value.value, true
			} else {
				delete(a.cache, key)
				return "", false
			}
		} else {
			return value.value, true
		}

	} else {
		return value.value, false
	}
}

func (a Cache) Put(key, value string) {
	temp := dataElement{value, time.Now(), false}
	a.cache[key] = temp
}

func (a Cache) Keys() []string {
	var out []string
	for key, v := range a.cache {
		if !v.elapsable {
			out = append(out, key)
		} else {
			if v.timestamp.After(time.Now()) {
				out = append(out, key)
			} else {
				delete(a.cache, key)
			}
		}

	}
	return out
}

func (a Cache) PutTill(key, value string, deadline time.Time) {
	temp := dataElement{value, deadline, true}
	a.cache[key] = temp
}
