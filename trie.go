package dataStructure

/*
 字典树，如果存储某个元素时用英文字符串表示，则查询和添加的时间复杂度只和英文字符串的长度有关，和所有元素个数无关
*/

// 字典树/前缀树
type Trie struct {
	size int
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
	}
}

func (t *Trie) GetSize() int {
	return t.size
}

func (t *Trie) Add(word string) {
	cur := t.root
	for i := 0; i < len(word); i++ {
		c := string(word[i])
		if cur.Next[c] == nil {
			cur.Next[c] = NewTrieNode()
		}
		cur = cur.Next[c]
	}
	if !cur.IsWord {
		cur.IsWord = true
		t.size++
	}
}

func (t *Trie) Add2(word string) {
	t.add(t.root, word)
}

func (t *Trie) add(node *TrieNode, word string) {
	if len(word) == 0 {
		if !node.IsWord {
			node.IsWord = true
			t.size++
		}
		return
	}
	c := string(word[0])
	if node.Next[c] == nil {
		trieNode := NewTrieNode()
		node.Next[c] = trieNode
		t.add(trieNode, word[1:])
		return
	}
	t.add(node.Next[c], word[1:])
}

func (t *Trie) Contains(word string) bool {
	cur := t.root
	for i := 0; i < len(word); i++ {
		c := string(word[i])
		if cur.Next[c] == nil {
			return false
		}
		cur = cur.Next[c]
	}
	return cur.IsWord
}

func (t *Trie) Contains2(word string) bool {
	return t.contains(t.root, word)
}

func (t *Trie) contains(node *TrieNode, word string) bool {
	if len(word) == 0 {
		return node.IsWord
	}
	c := string(word[0])
	if node.Next[c] == nil {
		return false
	}
	return t.contains(node.Next[c], word[1:])
}

func (t *Trie) IsPrefix(prefix string) bool {
	cur := t.root
	for i := 0; i < len(prefix); i++ {
		c := string(prefix[i])
		if cur.Next[c] == nil {
			return false
		}
		cur = cur.Next[c]
	}
	return true
}

func (t *Trie) Remove(word string) {
	t.remove(t.root, word, 0)
}

func (t *Trie) remove(node *TrieNode, word string, index int) {
	if index == len(word) {
		node.IsWord = false
		return
	}
	c := string(word[index])
	if node.Next[c] == nil {
		return
	}
	t.remove(node.Next[c], word, index+1)
	// 除了当前元素还有其他元素使用，不能删除
	if len(node.Next[c].Next) == 0 && !node.Next[c].IsWord {
		delete(node.Next, c)
	}
}
