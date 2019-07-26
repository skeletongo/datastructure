package dataStructure

/*
 字典树，如果存储某个元素时用英文字符串表示，则查询和添加的时间复杂度只和英文字符串的长度有关，和所有元素个数无关
*/

// 字典树/前缀树
type Trie struct {
	size int
	root *node
}

func NewTrie() *Trie {
	return &Trie{
		root: newNode(),
	}
}

func (t *Trie) GetSize() int {
	return t.size
}

func (t *Trie) Add(word string) {
	if word == "" {
		return
	}
	cur := t.root
	for i := 0; i < len(word); i++ {
		c := string(word[i])
		if cur.next[c] == nil {
			cur.next[c] = newNode()
		}
		cur = cur.next[c]
	}
	if !cur.isWord {
		cur.isWord = true
		t.size++
	}
}

func (t *Trie) Add2(word string) {
	if word == "" {
		return
	}
	t.add(t.root, word)
}

func (t *Trie) add(node *node, word string) {
	if len(word) == 0 {
		if !node.isWord {
			node.isWord = true
			t.size++
		}
		return
	}
	c := string(word[0])
	if node.next[c] == nil {
		trieNode := newNode()
		node.next[c] = trieNode
		t.add(trieNode, word[1:])
		return
	}
	t.add(node.next[c], word[1:])
}

func (t *Trie) Contains(word string) bool {
	cur := t.root
	for i := 0; i < len(word); i++ {
		c := string(word[i])
		if cur.next[c] == nil {
			return false
		}
		cur = cur.next[c]
	}
	return cur.isWord
}

func (t *Trie) Contains2(word string) bool {
	return t.contains(t.root, word)
}

func (t *Trie) contains(node *node, word string) bool {
	if len(word) == 0 {
		return node.isWord
	}
	c := string(word[0])
	if node.next[c] == nil {
		return false
	}
	return t.contains(node.next[c], word[1:])
}

func (t *Trie) IsPrefix(prefix string) bool {
	if prefix == "" {
		return false
	}
	cur := t.root
	for i := 0; i < len(prefix); i++ {
		c := string(prefix[i])
		if cur.next[c] == nil {
			return false
		}
		cur = cur.next[c]
	}
	return true
}

func (t *Trie) Remove(word string) {
	if word == "" {
		return
	}
	t.remove(t.root, word, 0)
}

func (t *Trie) remove(node *node, word string, index int) {
	if index == len(word) {
		node.isWord = false
		return
	}
	c := string(word[index])
	if node.next[c] == nil {
		return
	}
	t.remove(node.next[c], word, index+1)
	// 除了当前元素还有其他元素使用，不能删除
	if len(node.next[c].next) == 0 && !node.next[c].isWord {
		delete(node.next, c)
	}
}
