//package main

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
				//delete(a.cache, key)
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
			}
		}

	}
	return out
}

func (a Cache) PutTill(key, value string, deadline time.Time) {
	temp := dataElement{value, deadline, true}
	a.cache[key] = temp
}

// type kvPair struct {
// 	key          string
// 	value        string
// 	expValue     string
// 	deadline     time.Time
// 	shouldExpire bool
// }

// func main() {
// 	cases := map[string]struct {
// 		kvPairs []kvPair
// 		expKeys []string
// 		waitFor time.Duration
// 	}{
// 		"empty": {
// 			kvPairs: []kvPair{},
// 			expKeys: []string{},
// 		},
// 		"one value": {
// 			kvPairs: []kvPair{{
// 				key:      "key1",
// 				value:    "value1",
// 				expValue: "value1",
// 			}},
// 			expKeys: []string{"key1"},
// 		},
// 		"several values": {
// 			kvPairs: []kvPair{
// 				{
// 					key:      "key1",
// 					value:    "value1",
// 					expValue: "value1",
// 				},
// 				{
// 					key:      "key2",
// 					value:    "value2",
// 					expValue: "value2",
// 				},
// 				{
// 					key:      "key3",
// 					value:    "value3",
// 					expValue: "value3",
// 				},
// 			},
// 			expKeys: []string{"key1", "key2", "key3"},
// 		},
// 		"overwrite a value": {
// 			kvPairs: []kvPair{
// 				{
// 					key:      "key1",
// 					value:    "value1",
// 					expValue: "anotherValue1",
// 				},
// 				{
// 					key:      "key2",
// 					value:    "value2",
// 					expValue: "value2",
// 				},
// 				{
// 					key:      "key3",
// 					value:    "value3",
// 					expValue: "value3",
// 				},
// 				{
// 					key:      "key1",
// 					value:    "anotherValue1",
// 					expValue: "anotherValue1",
// 				},
// 			},
// 			expKeys: []string{"key1", "key2", "key3"},
// 		},
// 		"expired values": {
// 			kvPairs: []kvPair{
// 				{
// 					key:      "key1",
// 					value:    "value1",
// 					expValue: "value1",
// 				},
// 				{
// 					key:          "key2",
// 					value:        "value2",
// 					expValue:     "value2",
// 					shouldExpire: true,
// 					deadline:     time.Now().Add(time.Second * 2),
// 				},
// 				{
// 					key:      "key3",
// 					value:    "value3",
// 					expValue: "value3",
// 					deadline: time.Now().Add(time.Minute * 2),
// 				},
// 			},
// 			expKeys: []string{"key1", "key3"},
// 			waitFor: time.Second * 3,
// 		},
// 	}

// 	for _, tt := range cases {
// 		c := NewCache()
// 		for _, p := range tt.kvPairs {
// 			if p.deadline.IsZero() {
// 				c.Put(p.key, p.value)
// 			} else {
// 				c.PutTill(p.key, p.value, p.deadline)
// 			}
// 		}
// 		time.Sleep(tt.waitFor)
// 		for _, p := range tt.kvPairs {
// 			v, ok := c.Get(p.key)
// 			if !p.shouldExpire {
// 				if !ok {
// 					fmt.Println("Get: value is not present, while it should")
// 				}

// 				if v == "" {
// 					fmt.Println("Get: returned value is empty, while should be set")
// 				}

// 				if v != p.expValue {
// 					fmt.Println("Get: returned value incorrect: want \"%s\", got \"%s\"", p.expValue, v)
// 				}
// 			} else {
// 				if ok {
// 					fmt.Println("Get: and expired value is present in the cache(returned ok==true), while it should't")
// 				}

// 				if v != "" {
// 					fmt.Println("Get: and expired value is not an empty value: \"%s\"", v)
// 				}
// 			}
// 		}

// 		v, ok := c.Get("notExistingKey")
// 		if ok {
// 			fmt.Println("Get: random key is present in the cache(returned ok==true), while it should't")
// 		}

// 		if v != "" {
// 			fmt.Println("Get: random key is not an empty value: \"%s\"", v)
// 		}

// 		keys := c.Keys()
// 		if len(tt.expKeys) != len(keys) {
// 			fmt.Println("Keys: number of returned keys is incorrect: exp: %d, got %d", len(tt.expKeys), len(keys))
// 		}
// 		for _, expKey := range tt.expKeys {
// 			exists := false
// 			for _, key := range keys {
// 				if expKey == key {
// 					exists = true
// 				}
// 			}
// 			if !exists {
// 				fmt.Println("Keys: a key \"%s\" is not present in the Keys() method output", expKey)
// 			}
// 		}
// 	}

// }
