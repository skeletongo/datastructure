package tree234

type Tree234 struct {
	root    *node
	Compare func(a, b interface{}) int
}

// New 创建红黑树(2-3-4树)
// 参数 compare 为自定义元素大小比较函数
// 大小比较函数 返回值：
// 负数	表示	a<b
// 0	表示	a=b
// 正数	表示	a>b
func New(compare func(a, b interface{}) int) *Tree234 {
	return &Tree234{Compare: compare}
}

func (r *Tree234) GetSize() int {
	if r.root == nil {
		return 0
	}
	return r.root.n
}

func (r *Tree234) IsEmpty() bool {
	return r.root == nil
}

func inOrder(n *node, list *[]interface{}) {
	if n == nil {
		return
	}
	inOrder(n.left, list)
	*list = append(*list, n.key)
	inOrder(n.right, list)
}

// isBST 判断是不是二分搜索树
func (r *Tree234) isBST() bool {
	list := new([]interface{})
	inOrder(r.root, list)
	for i := 1; i < len(*list); i++ {
		if r.Compare((*list)[i-1], (*list)[i]) > 0 {
			return false
		}
	}
	return true
}

// isBalanced 判断是不是平衡二叉树(黑平衡)
func (r *Tree234) isBalanced() bool {
	return true
}

func size(n *node) int {
	if n == nil {
		return 0
	}
	return n.n
}

// 隐含条件，n不为nil
func getSize(n *node) int {
	return size(n.left) + size(n.right) + 1
}

// isRed 是否为红节点
func isRed(n *node) bool {
	if n == nil {
		return Black
	}
	return n.color // n.color == red 不用等号判断需要定义Red为true
}

// leftRotate 左旋转
//   node                     x
//  /   \     左旋转         /  \
// T1   x   --------->   node   T3
//     / \              /   \
//    T2 T3            T1   T2
func leftRotate(n *node) *node {
	x := n.right
	n.right = x.left
	x.left = n
	x.color = n.color
	n.color = Red
	x.n = n.n
	n.n = getSize(n)
	return x
}

// rightRotate 右旋转
//     node                   x
//    /   \     右旋转       /  \
//   x    T2   ------->   y   node
//  / \                       /  \
// y  T1                     T1  T2
func rightRotate(n *node) *node {
	x := n.left
	n.left = x.right
	x.right = n
	x.color = n.color
	n.color = Red
	x.n = n.n
	n.n = getSize(n)
	return x
}

// flipColors 颜色翻转
func flipColors(n *node) {
	n.color = Red
	n.left.color = Black
	n.right.color = Black
}

func isFlipColors(n *node) {
	if isRed(n.left) && isRed(n.right) {
		if isRed(n.left.left) || isRed(n.left.right) || isRed(n.right.left) || isRed(n.right.right) {
			flipColors(n)
		}
	}
}

// balance 维护红节点的位置，
// 1.将不正确的4节点变幻成正确的4节点
// 2.将5节点分解成一个2节点和一个4节点，并将剩余的节点向它的父节点融合
// 隐含条件，n不为nil
func balance(n *node) *node {
	if isRed(n.left) {
		if isRed(n.left.right) {
			n.left = leftRotate(n.left)
		}
		if isRed(n.left.left) {
			n = rightRotate(n)
		}
		isFlipColors(n)
	} else if isRed(n.right) {
		if isRed(n.right.left) {
			n.right = rightRotate(n.right)
		}
		if isRed(n.right.right) {
			n = leftRotate(n)
		}
		isFlipColors(n)
	}
	n.n = getSize(n)
	return n
}

func (r *Tree234) put(n *node, key, value interface{}) *node {
	if n == nil {
		return newNode(key, value)
	}

	res := r.Compare(n.key, key)
	if res > 0 {
		n.left = r.put(n.left, key, value)
	} else if res < 0 {
		n.right = r.put(n.right, key, value)
	} else {
		n.value = value
		return n
	}

	return balance(n)
}

func (r *Tree234) Put(key, value interface{}) {
	r.root = r.put(r.root, key, value)
	r.root.color = Black
}

// put 方式定义允许存在的节点有四种，分别是2节点，左倾3节点，右倾3节点，4节点
// put2 方式定义允许存在的节点有三种，分别是2节点，左倾3节点，4节点
// 区别：
// put2 方式代码简化，减少了很多判断，去掉了判断右倾3节点的情况，但是增加了左旋转的次数，相对于 put 方式
// put 方式判断的情况比较多但是减少了旋转次数
func (r *Tree234) put2(n *node, key, value interface{}) *node {
	if n == nil {
		return newNode(key, value)
	}

	if isRed(n.left) && isRed(n.right) {
		flipColors(n)
	}

	res := r.Compare(n.key, key)
	if res > 0 {
		n.left = r.put(n.left, key, value)
	} else if res < 0 {
		n.right = r.put(n.right, key, value)
	} else {
		n.value = value
		return n
	}

	if isRed(n.right) && !isRed(n.left) {
		n = leftRotate(n)
	}
	if isRed(n.left) && isRed(n.left.left) {
		n = rightRotate(n)
	}

	n.n = getSize(n)
	return n
}

func (r *Tree234) contains(n *node, key interface{}) bool {
	if n == nil {
		return false
	}

	res := r.Compare(n.key, key)
	if res > 0 {
		return r.contains(n.left, key)
	}
	if res < 0 {
		return r.contains(n.right, key)
	}
	return true
}

func (r *Tree234) Contains(key interface{}) bool {
	return r.contains(r.root, key)
}

func (r *Tree234) get(n *node, key interface{}) interface{} {
	if n == nil {
		return nil
	}

	res := r.Compare(n.key, key)
	if res > 0 {
		return r.get(n.left, key)
	}
	if res < 0 {
		return r.get(n.right, key)
	}
	return n.value
}

func (r *Tree234) Get(key interface{}) interface{} {
	return r.get(r.root, key)
}
