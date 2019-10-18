package levelcache

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

type Node struct {
	Key      string
	NextNode map[string]*Node
}

type LevelCache struct {
	Index *Node
	Data  map[string][]byte
	lock  sync.RWMutex
}

func (c *LevelCache) Set(key string, value []byte) {
	c.lock.Lock()
	defer c.lock.Unlock()
	//build index
	node := c.Index
	for _, str := range ParseStr(key) {
		if _, ok := node.NextNode[str]; !ok {
			node.NextNode[str] = &Node{
				Key:      str,
				NextNode: make(map[string]*Node),
			}
		}

		node = node.NextNode[str]
	}
	c.Data[key] = value
}

func (c *LevelCache) Get(key string) ([]byte, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	if v, ok := c.Data[key]; ok {
		return v, nil
	}
	return nil, errors.New("Key not found")
}

func (c *LevelCache) Find(value string) (res map[string][]byte) {
	res = make(map[string][]byte)
	c.lock.RLock()
	defer c.lock.RUnlock()
	node := c.Index
	var ss []string
	for _, str := range ParseStr(value) {
		if _, ok := node.NextNode[str]; !ok {
			return
		}
		node = node.NextNode[str]
		ss = append(ss, str)
	}
	if len(node.NextNode) == 0 {
		k := fmt.Sprintf("/%s/", strings.Join(ss, "/"))
		v, ok := c.Data[k]
		if ok {
			res[k] = v
		}
	} else {
		var keys []string
		getNodeKeys(strings.Join(ss, "/"), node, &keys)
		for _, k := range keys {
			v, ok := c.Data[k]
			if ok {
				res[k] = v
			}
		}
	}
	return
}

func getNodeKeys(key string, node *Node, res *[]string) {
	for k, n := range node.NextNode {
		if len(n.NextNode) == 0 {
			newKey := fmt.Sprintf("/%s/%s/", key, k)
			*res = append(*res, newKey)
		} else {
			newKey := fmt.Sprintf("%s/%s", key, k)
			getNodeKeys(newKey, n, res)
		}
	}
}

func New() *LevelCache {
	return &LevelCache{
		Index: &Node{
			NextNode: make(map[string]*Node),
		},
		Data: make(map[string][]byte),
	}
}

//ParseStr ... common func to parse string by char
func ParseStr(value string) []string {
	s := strings.TrimSuffix(strings.TrimPrefix(value, "/"), "/")
	return strings.Split(s, "/")
}
