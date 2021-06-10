package trie

import (
	"fmt"
	"strings"
	"testing"
)

var data = map[string]string{
	"cat":   "猫",
	"dog":   "狗",
	"panda": "熊猫",
	"pan":   "平底锅",
}

var data2 = map[string]string{
	"cat":   "猫猫",
	"dog":   "狗狗",
	"panda": "熊猫熊猫",
	"pan":   "平底锅平底锅",
	"pig":   "猪",
}

func TestNew(t *testing.T) {
	tree := New()
	for k, v := range data {
		tree.Add(k, v)
	}

	t.Log(tree)

	if tree.GetSize() != len(data) {
		t.Errorf("GetSize error: len(data)=%d, GetSize()=%d\n", len(data), tree.GetSize())
	}

	for k := range data {
		if !tree.Contains(k) {
			t.Errorf("Contains error: word=%s\n", k)
		}
	}

	for k := range data {
		if !tree.HasPrefix(string(k[0])) {
			t.Errorf("HasPrefix error: prefix=%s\n", string(k[0]))
		}
	}

	res := tree.GetPrefix("dog")
	if len(res) != 1 {
		t.Error("GetPrefix error: prefix=dog")
	}
	t.Logf("GetPrefix(dog) %s\n", res["dog"])
	res = tree.GetPrefix("pan")
	if len(res) != 2 {
		t.Error("GetPrefix error: prefix=pan")
	}
	for k := range res {
		t.Logf("GetPrefix(pan) %s\n", res[k])
		if _, ok := data[k]; !ok {
			t.Errorf("GetPrefix error: prefix=pan, words=%v", res)
		}
	}

	res = tree.GetPrefix("")
	if len(res) != len(data) {
		t.Errorf("GetPrefix error: len(res):%d != len(data):%d\n", len(res), len(data))
	}
	for k, v := range res {
		t.Logf("GetPrefix() %s:%s\n", k, data[k])
		if data[k] != v {
			t.Errorf("GetPrefix error: data[K]=%v, res[k]=%v\n", data[k], v)
		}
	}

	for k := range data {
		prefix := string(k[0])
		count := tree.GetPrefixCount(prefix)
		var num int
		for k := range data {
			if strings.HasPrefix(k, prefix) {
				num++
			}
		}
		t.Logf("GetPrefixCount prefix=%s, count=%d\n", prefix, num)
		if num != count {
			t.Errorf("GetPrefixCount error: prefix=%s, correct=%d result=%d\n", prefix, num, count)
		}
	}

	for k, v := range data2 {
		old := tree.Add(k, v)
		t.Logf("Add(%s) old=%v, now=%v\n", k, old, v)
		if _, ok := data[k]; !ok {
			if old != nil {
				t.Errorf("Add no exist error: word=%s, return=%v\n", k, old)
			}
		} else {
			if old != data[k] {
				t.Errorf("Add exist error: word=%s, return=%v\n", k, old)
			}
		}
	}

	t.Log(tree)

	if tree.Remove("pa") != nil {
		t.Error("Remove error: pa")
	}

	for k, v := range data2 {
		if old := tree.Remove(k); old != v {
			t.Errorf("Remove error: word=%s, correct=%v, return=%v\n", k, v, old)
		}
	}

	if tree.GetSize() != 0 {
		t.Errorf("Remove error: size=%d\n", tree.GetSize())
	}

	t.Log(tree)
}

var chineseWord = map[string]string{
	"李白":  "《望庐山瀑布》",
	"李商隐": "《锦瑟》",
	"杜甫":  "《登高》",
}

func TestNew2(t *testing.T) {
	tree := New()
	for k, v := range chineseWord {
		tree.Add(k, v)
	}

	fmt.Println("--> String\n", tree)

	fmt.Println("--> GetPrefix(李)")
	for k, v := range tree.GetPrefix("李") {
		fmt.Println(k, v)
	}

	fmt.Println("--> GetPrefix(杜)")
	for k, v := range tree.GetPrefix("杜") {
		fmt.Println(k, v)
	}

	fmt.Println("--> GetPrefix()")
	for k, v := range tree.GetPrefix("") {
		fmt.Println(k, v)
	}
}

func TestTrie_Range(t *testing.T) {
	tree := New()
	for k, v := range chineseWord {
		tree.Add(k, v)
	}

	tree.Range(func(word string, value interface{}) bool {
		fmt.Println(word, value)
		return true
	})
}
