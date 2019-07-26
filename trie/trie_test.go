package dataStructure

import (
	"testing"
)

func TestTrie(t *testing.T) {
	test := map[string]bool{
		"cat":   true,
		"dog":   true,
		"panda": true,
		"pan":   false,
	}

	trie := NewTrie()
	for k, v := range test {
		if v {
			trie.Add(k)
		}
	}

	for k, v := range test {
		if trie.Contains(k) != v {
			t.Error("字典树错误")
		}
	}

	trie.Add("pan")
	if !trie.Contains("pan") {
		t.Error("字典树错误")
	}

	if !trie.IsPrefix("pa") {
		t.Error("前缀查询错误")
	}
	if trie.IsPrefix("dd") {
		t.Error("前缀查询错误")
	}
	if trie.IsPrefix("catt") {
		t.Error("前缀查询错误")
	}
}

func TestTrie2(t *testing.T) {
	test := map[string]bool{
		"cat":   true,
		"dog":   true,
		"panda": true,
		"pan":   false,
	}

	trie := NewTrie()
	for k, v := range test {
		if v {
			trie.Add2(k)
		}
	}

	for k, v := range test {
		if trie.Contains2(k) != v {
			t.Error("字典树错误")
		}
	}

	trie.Add("pan")
	if !trie.Contains2("pan") {
		t.Error("字典树错误")
	}
}

func TestTrie_Remove(t *testing.T) {
	trie := NewTrie()
	trie.Add("dog")
	trie.Add("deer")
	trie.Add("panda")
	trie.Add("pan")

	if !trie.Contains("dog") {
		t.Error("no dog")
	}
	if !trie.Contains("deer") {
		t.Error("no deer")
	}
	if !trie.Contains("panda") {
		t.Error("no panda")
	}
	if !trie.Contains("pan") {
		t.Error("no pan")
	}

	trie.Remove("dog")
	if trie.Contains("dog") {
		t.Error("no dog")
	}
	if !trie.Contains("deer") {
		t.Error("no deer")
	}

	trie.Remove("panda")
	if trie.Contains("panda") {
		t.Error("no panda")
	}
	if !trie.Contains("pan") {
		t.Error("no pan")
	}

	trie.Add("panda")
	trie.Remove("pan")
	if trie.Contains("pan") {
		t.Error("no pan")
	}
	if !trie.Contains("panda") {
		t.Error("no panda")
	}
}