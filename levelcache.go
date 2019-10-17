package levelcache

import "strings"

type Node struct {
	value    string
	leaf bool
	children map[string]*Node
}

type Trie struct {
	root *Node
}

// /hello/world/
func (t *Trie) insert(value string) {
	node := t.root
	for pos, str := range ParseStr(value) {
		if _, ok := node.children[str]; !ok {
			node.children[str] = &Node{
				value:    str,
				children: make(map[string]*Node),
			}
		}

		node = node.children[str]
		if pos == len(ParseStr(value))-1 {
			node.leaf = true
		}
	}
}

func ParseStr(value string) []string {
	s := strings.TrimSuffix(strings.TrimPrefix(value, "/"), "/")
	return strings.Split(s, "/")
}

func (t *Trie) find(value string) (*Node, bool) {
	node := t.root
	for _, str := range ParseStr(value) {
		if _, ok := node.children[str]; !ok {
			return nil, false
		}

		node = node.children[str]
	}

	return node, true
}

func New() *Trie {
	return &Trie{
		root: &Node{
			children: make(map[string]*Node),
		},
	}
}
