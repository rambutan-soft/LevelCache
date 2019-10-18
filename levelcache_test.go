package levelcache

import (
	"testing"
)

func TestKeyFind(t *testing.T) {
	cache := New()
	cache.Set("/hello/world/cache/", []byte("hello"))

	_, err := cache.Get("/hello/")
	if err == nil {
		t.Error("Expected error, but no error")
	}

	v, err := cache.Get("/hello/world/cache/")
	if err != nil {
		t.Error("Expected no error, but get error")
	}

	if string(v) != "hello" {
		t.Errorf("Expected [hello], but got %s", v)
	}
}

func TestKeyTree(t *testing.T) {
	cache := New()
	cache.Set("/hello/world/cache/", []byte("hello"))
	cache.Set("/hello/world/cache1/", []byte("hello"))
	cache.Set("/hello/world/cache2/", []byte("hello"))
	cache.Set("/hello/world1/cache3/", []byte("hello"))
	cache.Set("/hello/world1/cache4/", []byte("hello"))

	nodes := cache.Find("/hello/world/cache/")
	if len(nodes) != 1 {
		t.Errorf("Expected value to be 1, but it was %d instead.", len(nodes))
	}
	nodes = cache.Find("/hello/world/")
	if len(nodes) != 3 {
		t.Errorf("Expected value to be 3, but it was %d instead.", len(nodes))
	}

	nodes = cache.Find("/hello/")
	if len(nodes) != 5 {
		t.Errorf("Expected value to be 5, but it was %d instead.", len(nodes))
	}

}
