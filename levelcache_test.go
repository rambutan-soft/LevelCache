package levelcache

import (
	"testing"
)

func TestFindExisting(t *testing.T) {
	trie := New()
	trie.insert("/hello/world/cache/")
	node, ok := trie.find("/hello/")

	if node.value != "hello" {
		t.Errorf("Expected value to be t, but it was %s instead.", node.value)
	}

	if node.leaf != false {
		t.Errorf("Expected value to be true, but it was %t instead.", node.leaf)
	}

	if !ok {
		t.Errorf("Expected ok to be true, but it was %t instead.", ok)
	}
}

func TestFindExisting2(t *testing.T) {
	trie := New()
	trie.insert("/hello/world/cache/")
	node, ok := trie.find("/hello/world/cache")

	if node.value != "cache" {
		t.Errorf("Expected value to be t, but it was %s instead.", node.value)
	}

	if node.leaf != true {
		t.Errorf("Expected value to be true, but it was %t instead.", node.leaf)
	}

	if !ok {
		t.Errorf("Expected ok to be true, but it was %t instead.", ok)
	}
}
