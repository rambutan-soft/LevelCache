# LevelCache
 Same with ETCD, the keys has tree structure, but it is a in memory key-value store


example:
```
import "fmt"
import "github.com/rambutan-soft/LevelCache"
cache := levelcache.New()
cache.Set("/hello/world/cache/", []byte("a"))
cache.Set("/hello/world/cache1/", []byte("b"))
cache.Set("/hello/world/cache2/", []byte("c"))
cache.Set("/hello/world1/cache3/", []byte("d"))
cache.Set("/hello/world1/cache4/", []byte("e"))

nodes := cache.Find("/hello/world/cache/")
//"/hello/world/cache/", []byte("a")

nodes = cache.Find("/hello/world/")

//"/hello/world/cache/", []byte("a")
//"/hello/world/cache1/", []byte("b")
//"/hello/world/cache2/", []byte("c")

nodes = cache.Find("/hello/")
//show all results

```
