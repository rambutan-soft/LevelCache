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

func BenchmarkInsert(b *testing.B) {
	// run the Fib function b.N times
	cache := New()
	//b.ResetTimer()
	for n := 0; n < b.N; n++ {
		cache.Set(string(n), []byte("123.456788"))
	}
}

func BenchmarkGet(b *testing.B) {
	// run the Fib function b.N times
	cache := New()
	//b.ResetTimer()
	for n := 0; n < b.N; n++ {
		cache.Set(string(n), []byte("123.456788"))
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		cache.Get(string(n))
	}
}

func BenchmarkParallelInsert(b *testing.B) {
	cache := New()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			cache.Set(string(b.N), []byte("123.456788"))
		}
	})
}

func BenchmarkParallelGet(b *testing.B) {
	cache := New()
	//b.ResetTimer()
	for n := 0; n < b.N; n++ {
		cache.Set(string(n), []byte("123.456788"))
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			cache.Set(string(b.N), []byte("123.456788"))
		}
	})
}

//BenchmarkInsert-4   	 1000000	      2138 ns/op	     456 B/op	       5 allocs/op
//BenchmarkParallelInsert-4   	 1000000	      1060 ns/op	     137 B/op	       3 allocs/op

// allocs/op means how many distinct memory allocations occurred per op (single iteration).
// B/op is how many bytes were allocated per op.
//go test -bench=. -benchmem -cpuprofile profile.out

//ogo tool pprof profile.out
