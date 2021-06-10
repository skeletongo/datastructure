// 字典树,主要用于存储字符串，是存储字符串的集合，有较高的搜索性能
// 如果存储某个元素时用英文字符串表示，则查询和添加的时间复杂度只和英文字符串的长度有关，和所有元素个数无关
package trie

import (
	"bytes"
	"fmt"
)

type node struct {
	isWord bool
	next   map[rune]*node // 使用rune类型可以支持多种语言而非只支持英文
	value  interface{}
}

func newNode() *node {
	return &node{
		isWord: false,
		next:   make(map[rune]*node),
	}
}

type Trie struct {
	size int
	root *node
}

// New 创建字典树
func New() *Trie {
	return &Trie{
		root: newNode(),
	}
}

// GetSize 单词数量
func (t *Trie) GetSize() int {
	return t.size
}

// Add 添加单词，非递归方式
func (t *Trie) Add(word string, v interface{}) interface{} {
	if word == "" {
		return nil
	}
	cur := t.root
	for _, v := range word {
		if cur.next[v] == nil {
			cur.next[v] = newNode()
		}
		cur = cur.next[v]
	}
	temp := cur.value
	cur.value = v
	if !cur.isWord {
		cur.isWord = true
		t.size++
	}
	return temp
}

// AddR 添加单词，递归方式
func (t *Trie) AddR(word string, v interface{}) interface{} {
	if word == "" {
		return nil
	}
	return t.add(t.root, []rune(word), v)
}

func (t *Trie) add(node *node, word []rune, v interface{}) interface{} {
	if len(word) == 0 {
		temp := node.value
		node.value = v
		if !node.isWord {
			node.isWord = true
			t.size++
		}
		return temp
	}
	c := word[0]
	if node.next[c] == nil {
		trieNode := newNode()
		node.next[c] = trieNode
		t.add(trieNode, word[1:], v)
		return nil
	}
	return t.add(node.next[c], word[1:], v)
}

// Contains 是否包含某个单词，非递归方式
func (t *Trie) Contains(word string) bool {
	cur := t.root
	for _, v := range word {
		if cur.next[v] == nil {
			return false
		}
		cur = cur.next[v]
	}
	return cur.isWord
}

// ContainsR 是否包含某个单词，递归方式
func (t *Trie) ContainsR(word string) bool {
	return t.contains(t.root, []rune(word))
}

func (t *Trie) contains(node *node, word []rune) bool {
	if len(word) == 0 {
		return node.isWord
	}
	c := word[0]
	if node.next[c] == nil {
		return false
	}
	return t.contains(node.next[c], word[1:])
}

// HasPrefix 是否包含某个前缀的单词
func (t *Trie) HasPrefix(prefix string) bool {
	if prefix == "" {
		return true
	}
	cur := t.root
	for _, v := range prefix {
		if cur.next[v] == nil {
			return false
		}
		cur = cur.next[v]
	}
	return true
}

// Remove 删除某个单词
func (t *Trie) Remove(word string) interface{} {
	if word == "" {
		return nil
	}
	return t.remove(t.root, []rune(word), 0)
}

func (t *Trie) remove(node *node, word []rune, index int) interface{} {
	if index == len(word) {
		if node.isWord {
			t.size--
			node.isWord = false
			return node.value
		}
		return nil
	}
	c := word[index]
	if node.next[c] == nil {
		return nil
	}
	temp := t.remove(node.next[c], word, index+1)
	// 除了当前元素还有其他元素使用，不能删除
	if len(node.next[c].next) == 0 && !node.next[c].isWord {
		delete(node.next, c)
	}
	return temp
}

// GetPrefix 获取所有包含某前缀的数据，包含前缀本身
func (t *Trie) GetPrefix(word string) map[string]interface{} {
	ret := make(map[string]interface{})

	cur := t.root
	for _, v := range word {
		if cur.next[v] == nil {
			return ret
		}
		cur = cur.next[v]
	}

	// 前缀本身表示的数据
	if word != "" && cur.isWord {
		ret[word] = cur.value
	}
	// 包含此前缀的所有数据
	t.findPrefix(ret, []rune(word), cur)
	return ret
}

func (t *Trie) findPrefix(data map[string]interface{}, s []rune, n *node) {
	l := len(s)
	for k, v := range n.next {
		s = append(s, k)
		if v.isWord {
			data[string(s)] = v.value
		}
		t.findPrefix(data, s, v)
		s = s[:l]
	}
}

// GetPrefixCount 查询包含某前缀的数据数量，包含前缀本身
func (t *Trie) GetPrefixCount(word string) int {
	var count int
	cur := t.root
	for _, v := range word {
		if cur.next[v] == nil {
			return count
		}
		cur = cur.next[v]
	}
	if cur.isWord {
		count++
	}
	t.findPrefixCount(&count, cur)
	return count
}

func (t *Trie) findPrefixCount(count *int, n *node) {
	for _, v := range n.next {
		if v.isWord {
			*count++
		}
		t.findPrefixCount(count, v)
	}
}

// Range 遍历字典树
// f函数可以对遍历的数据进行操作，返回值为false时停止遍历
func (t *Trie) Range(f func(word string, value interface{}) bool) {
	t.toRange([]rune{}, t.root, f)
}

func (t *Trie) toRange(s []rune, n *node, f func(word string, value interface{}) bool) bool {
	if n.isWord {
		if !f(string(s), n.value) {
			return false
		}
	}
	if len(n.next) == 0 {
		return true // 没有子节点了
	}
	l := len(s)
	for k, v := range n.next {
		s = append(s, k)
		if !t.toRange(s, v, f) {
			return false
		}
		s = s[:l]
	}
	return true
}

func (t *Trie) String() string {
	buf := bytes.Buffer{}
	buf.WriteString(fmt.Sprintf("size: %d\n", t.GetSize()))
	for k, v := range t.GetPrefix("") {
		buf.WriteString(fmt.Sprintf("%v: %v\n", k, v))
	}
	return buf.String()
}
