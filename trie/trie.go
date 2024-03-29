package trie

type Node struct {
	passCount int
	isEnd     bool
	children  map[any]*Node
}

type Trie struct {
	root *Node
	size int
}

func New() *Trie {
	return &Trie{
		root: &Node{children: map[any]*Node{}},
		size: 0,
	}
}

func (t *Trie) Put(word string) bool {
	node := t.search(word)
	if node != nil && node.isEnd {
		return true
	}

	node = t.root
	for _, e := range word {
		if node.children[e] == nil {
			node.children[e] = &Node{children: make(map[any]*Node)}
		}
		node.children[e].passCount++
		node = node.children[e]
	}
	node.isEnd = true
	t.size++
	return true
}

func (t *Trie) Delete(word string) bool {
	node := t.search(word)
	if node == nil || !node.isEnd {
		return false
	}

	node = t.root
	for _, e := range word {
		node.children[e].passCount--
		if node.children[e].passCount == 0 {
			node.children[e] = nil
			t.size--
			return true
		}
		node = node.children[e]
	}
	t.size--
	node.isEnd = false
	return true
}

func (t *Trie) StartWith(prefix string) bool {
	return t.search(prefix) != nil
}

func (t *Trie) PassCount(prefix string) int {
	node := t.search(prefix)
	if node == nil {
		return 0
	}
	return node.passCount
}

func (t *Trie) search(target string) *Node {
	node := t.root
	for _, e := range target {
		if node.children[e] == nil {
			return nil
		}
		node = node.children[e]
	}
	return node
}
