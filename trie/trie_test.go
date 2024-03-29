package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	trie := New()

	// Put
	ok := trie.Put("apple")
	assert.True(t, ok)
	assert.Equal(t, 1, trie.size)

	ok = trie.Put("banana")
	assert.True(t, ok)
	assert.Equal(t, 2, trie.size)

	ok = trie.Put("app")
	assert.True(t, ok)
	assert.Equal(t, 3, trie.size)

	ok = trie.Put("bat")
	assert.True(t, ok)
	assert.Equal(t, 4, trie.size)

	// Put 重复值
	ok = trie.Put("apple")
	assert.True(t, ok)
	assert.Equal(t, 4, trie.size)

	// Start With
	val := trie.StartWith("app")
	assert.True(t, val)
	val = trie.StartWith("appg")
	assert.False(t, val)

	// Delete
	del := trie.Delete("a")
	assert.False(t, del)

	del = trie.Delete("apple")
	assert.True(t, del)
	assert.Equal(t, 3, trie.size)

}
