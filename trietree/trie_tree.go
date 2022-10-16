package trietree

type node struct {
	Pass int
	End  int
	Next []*node // 表示下一个节点，初始化时为26个空间，分别小写字母a-z
	// 注：如果多的话，可以用map<byte, node>来表示
}

// NewTrieTree 初始化前缀树
func NewTrieTree() *node {
	return &node{
		Pass: 0,
		End:  0,
		Next: make([]*node, 26),
	}
}

// Insert 插入字符串到字符
func (root *node) Insert(word string) {
	root.Pass++
	curNode := root
	for _, char := range word {
		index := char - 'a'
		if curNode.Next[index] == nil {
			curNode.Next[index] = NewTrieTree()
		}
		curNode.Next[index].Pass++
		curNode = curNode.Next[index]
	}
	curNode.End++
}

func (root *node) Delete(word string) {
	if !root.IsExist(word) { // 检查字符是否存在，避免误删
		return
	}
	root.Pass--
	curNode := root
	for _, char := range word {
		index := char - 'a'
		if curNode.Next[index].Pass == 1 { // 如果当前只剩下1个，表示减后=0，所以直接置为nil
			curNode.Next[index] = nil
			return
		}
		curNode.Next[index].Pass--
		curNode = curNode.Next[index]
	}
	curNode.End--
}

func (root *node) IsExist(word string) bool {
	return root.Search(word) > 0
}

func (root *node) Search(word string) int {
	curNode := root
	for _, char := range word {
		index := char - 'a'
		if curNode.Next[index] == nil {
			return 0
		}
		curNode = curNode.Next[index]
	}
	return curNode.End
}

func (root *node) PrefixMatch(word string) int {
	curNode := root
	for _, char := range word {
		index := char - 'a'
		if curNode.Next[index] == nil {
			return 0
		}
		curNode = curNode.Next[index]
	}
	return curNode.Pass
}
