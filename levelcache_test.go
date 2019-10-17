package levelcache

import (
	"testing"
)

func TestFindExisting(t *testing.T) {
	trie := New()
	trie.insert("/hello/world/cache/")
	node, ok := trie.find("/hello/")

	if node.key != "hello" {
		t.Errorf("Expected value to be t, but it was %s instead.", node.key)
	}

	if node.leaf != false {
		t.Errorf("Expected value to be true, but it was %t instead.", node.leaf)
	}

	if !ok {
		t.Errorf("Expected ok to be true, but it was %t instead.", ok)
	}

	node2, ok := trie.find("/hello/world/cache")

	if node2.key != "cache" {
		t.Errorf("Expected value to be t, but it was %s instead.", node2.key)
	}

	if node2.leaf != true {
		t.Errorf("Expected value to be true, but it was %t instead.", node2.leaf)
	}

	if !ok {
		t.Errorf("Expected ok to be true, but it was %t instead.", ok)
	}
}

func TestFindExisting2(t *testing.T) {
	trie := New()
	trie.insert("/hello/world/cache/")
	trie.insert("/hello/world/cache1/")
	trie.insert("/hello/world/cache2/")
	trie.insert("/hello/world1/cache/")
	trie.insert("/hello/world1/cache1/")

	node, _ := trie.find("/hello/")
	if len(node.children) != 2 {
		t.Errorf("Expected value to be 2, but it was %d instead.", len(node.children))
	}

	node, _ = trie.find("/hello/world")
	if len(node.children) != 3 {
		t.Errorf("Expected value to be 2, but it was %d instead.", len(node.children))
	}
}
