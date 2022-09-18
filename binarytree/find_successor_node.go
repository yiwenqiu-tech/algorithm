package binarytree

//该结构比普通二叉树节点结构多了一个指向父节点的parent指针。
//假设有一棵Node类型的节点组成的二叉树，树中每个节点的parent指针都正确地指向自己的父节点，头节点的parent指向null。只给一个在二叉树中的某个节点node，
//请实现返回node的后继节点的函数。在二叉树的中序遍历的序列中，node的下一个节点叫作node的后继节点。

type NodeWithParent struct {
	Value  int
	Left   *NodeWithParent
	Right  *NodeWithParent
	Parent *NodeWithParent
}

// FindSuccessorNode1
// 暴力解法，如题目所述，在二叉树的中序遍历的序列中，node的下一个节点叫作node的后继节点。
// 	所以，最简单的方法是，中序排序后，再遍历找到某个节点后面的节点
func FindSuccessorNode1(head *NodeWithParent, node *NodeWithParent) *NodeWithParent {
	if head == nil {
		return nil
	}
	var midOrderRes []*NodeWithParent
	processForFindSuccessorNode1(head, &midOrderRes)

	successorNodeIndex := -1
	for index, r := range midOrderRes {
		if r == node {
			successorNodeIndex = index + 1
		}
	}
	if successorNodeIndex >= len(midOrderRes) {
		return nil
	}
	return (midOrderRes)[successorNodeIndex]
}

func processForFindSuccessorNode1(head *NodeWithParent, res *[]*NodeWithParent) {
	if head == nil {
		return
	}
	processForFindSuccessorNode1(head.Left, res)
	*res = append(*res, head)
	processForFindSuccessorNode1(head.Right, res)
}

// FindSuccessorNode2
// 	重新审题，由于该树与常规二叉树不一样，该树的节点具备父节点的指针，所以可以根据中序排序的特点来寻找后继节点
//	整理特点如下：
//		1.若存在右子树，则取右子树的最左节点
//		2.若不存在右子树，则找到它第一个右侧的父节点
func FindSuccessorNode2(head *NodeWithParent, node *NodeWithParent) *NodeWithParent {
	if head == nil {
		return nil
	}
	if node == nil {
		return nil
	}
	if node.Right != nil {
		res := node.Right
		for res.Left != nil {
			res = res.Left
		}
		return res
	}
	cur := node
	for cur.Parent != nil {
		if cur.Parent.Left == cur {
			return cur.Parent
		}
		cur = cur.Parent
	}
	return nil
}
