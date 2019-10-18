# LevelCache
 Same with ETCD, the keys has tree structure, but it is a in memory key-value store


example:
```
import "fmt"
import "github.com/rambutan-soft/LevelCache"
cache := levelcache.New()
cache.Set("/hello/world/cache/", []byte("hello"))
cache.Set("/hello/world/cache1/", []byte("hello"))
cache.Set("/hello/world/cache2/", []byte("hello"))
cache.Set("/hello/world1/cache3/", []byte("hello"))
cache.Set("/hello/world1/cache4/", []byte("hello"))

nodes := cache.Find("/hello/world/cache/")
if len(nodes) != 1 {
    fmt.Printf("Expected value to be 1, but it was %d instead.", len(nodes))
}
nodes = cache.Find("/hello/world/")
if len(nodes) != 3 {
    fmt.Printf("Expected value to be 3, but it was %d instead.", len(nodes))
}

nodes = cache.Find("/hello/")
if len(nodes) != 5 {
    fmt.Printf("Expected value to be 5, but it was %d instead.", len(nodes))
}
```
