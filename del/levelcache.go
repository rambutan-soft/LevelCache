package levelcache

import (
	mr "math/rand"
	"time"
)

// this package is designed to create in-memory cache
// for KV pairs
// Value is a string but can store anything in the string
// save a JSON object is valid

// Interface is an interface for creating a cache list
type Interface interface {
	// Len returns the length of the Cache
	Len()
	// GetValueByKey returns the value given a key
	GetValueByKey(key string)
	// GetKeyByIndex returns the key at given index
	GetKeyByIndex(index int)
	// GetValueByIndex returns the value at given index
	GetValueByIndex(index int)
	// GetRandomKey returns a random key pair
	GetRandomKey()
	// GetRandomKV returns a random KV pair
	GetRandomKV()
	// Add
	Add(key string, value string)
	// create a New Cache object
	NewCache()
}

// Cache will be a generic cache
type Cache struct {
	KeyValue KV
	Key      Keys
	Size     int
}

// KV is the KV cache
type KV map[string]string

// Keys will hold the list of the ID's (or keys) that map to a value in KV Cache
type Keys []string

// Add will upsert a new key value pair in KV cache
func (c *Cache) Add(key string, value string) {
	// Only add a new key and increase size if key doesn't exists
	if _, exists := c.KeyValue[key]; !exists {
		c.Key = append(c.Key, key)
		c.Size++
	}
	c.KeyValue[key] = value
}

// Delete will remove key value from the cache
func (c *Cache) Delete(key string) {
	if _, exists := c.KeyValue[key]; !exists {
		return
	}
	delete(c.KeyValue, key)
	for i, k := range c.Key {
		if k == key {
			c.Key = append(c.Key[:i], c.Key[i+1:]...)
		}
	}
	c.Size--
}

// GetValueByKey will return the value mapped to KV cache
func (c *Cache) GetValueByKey(key string) string {
	return c.KeyValue[key]
}

// Len will return the size of the cache
func (c *Cache) Len() int {
	return c.Size
}

// GetKeyByIndex will return the key at given index position
func (c *Cache) GetKeyByIndex(index int) string {
	return c.Key[index]
}

// GetValueByIndex will return the id given an index position
func (c *Cache) GetValueByIndex(index int) string {
	return c.GetValueByKey(c.Key[index])
}

// GetRandomKey returns a random key indexed with a randomly generated position.
//
// seed values (default Source) that do not change will generate the same
// pseudo-random sequence, which means Intn will return the same number.
func (c *Cache) GetRandomKey() string {
	var randomInt int
	// Seed uses the Unix time in nanoseconds to initialize the default Source.
	// Providing a unique default Source (seed value) produces a pseudo-random number.
	mr.Seed(time.Now().UTC().UnixNano())
	lenIndex := c.Len()

	// Making sure we dont divide by 0
	// Intn returns a pseudo-random number in [0, lenIndex) from the default seed Source,
	// which represents a pseudo-random sequence.
	if randomInt = mr.Intn(lenIndex); randomInt == 0 {
		randomInt = 1
	}
	// pos is assigned to the remainder of lenIndex / randomInt, which is then used
	// to index Key in cache.
	pos := lenIndex % randomInt
	return c.Key[pos]
}

// GetRandomKV returns a random key from Keys using a randomly generated pos,
// and a random value from KV using the randomly selected key.
//
// seed values (default Source) that do not change will generate the same
// pseudo-random sequence, which means Intn will return the same number.
func (c *Cache) GetRandomKV() (string, string) {
	var randomInt int
	// Seed initializes the default Source using the Unix time in nanoseconds.
	mr.Seed(time.Now().UTC().UnixNano())
	// Len returns length of cache
	lenIndex := c.Len()

	// Making sure we dont divide by 0
	// Intn returns a pseudo-random number in [0, lenIndex) from the default seed Source,
	// which represents a pseudo-random sequence.
	if randomInt = mr.Intn(lenIndex); randomInt == 0 {
		randomInt = 1
	}
	// pos is assigned to the remainder of lenIndex / randomInt, which is then used
	// to index and return a random Key and value in cache.
	pos := lenIndex % randomInt
	return c.Key[pos], c.GetValueByKey(c.Key[pos])
}

// NewCache creates a new Cache and return the pointer.
func NewCache() *Cache {
	var a Cache
	// make returns the same type as KV, which in this case is an empty map.
	a.KeyValue = make(KV)
	return &a
}