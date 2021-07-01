package hashtable

import "fmt"

// capacity 扩容取模用到的素数
var capacity = []int{
	53, 97, 193, 389, 769, 1543, 3079, 6151, 12289, 24593,
	49157, 98317, 196613, 393241, 786433, 1572869, 3145739, 6291469,
	12582917, 25165843, 50331653, 100663319, 201326611, 402653189, 805306457, 1610612741,
}

const (
	upperTol = 10 // 哈希冲突达到10进行扩容
	lowerTol = 2  // 哈希冲突减少到2进行缩容
)

// hash 计算哈希值
func hash(k interface{}, m int) int {
	key := fmt.Sprintf("%s", k)
	hash := 0
	for i := 0; i < len(key); i++ {
		hash = hash*31 + int(key[i])
	}
	return hash % m
}
