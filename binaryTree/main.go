// https://www.geeksforgeeks.org/preorder-traversal-of-binary-tree/
package binaryTree

import (
	"container/list"
	"fmt"
)

type dataType = int

type BinaryTreeNode struct {
	Data dataType

	R *BinaryTreeNode
	L *BinaryTreeNode
}

func (root *BinaryTreeNode) DeleteNode(data dataType) *BinaryTreeNode {
	if root == nil {
		return nil
	}

	q := list.New()

	q.PushBack(root)

	var target *BinaryTreeNode

	for q.Len() != 0 {
		temp := q.Front()

		q.Remove(temp)

		node := temp.Value.(*BinaryTreeNode)

		if node.Data == data {
			target = node
			break
		}

		if node.L != nil {
			q.PushBack(node.L)
		}

		if node.R != nil {
			q.PushBack(node.R)
		}
	}

	if target == nil {
		return root
	}

	var last [2]*BinaryTreeNode

	q = list.New()

	q.PushBack([2]*BinaryTreeNode{root, nil})

	for q.Len() != 0 {
		temp := q.Front()

		q.Remove(temp)

		node := temp.Value.([2]*BinaryTreeNode)
		last = node

		if node[0].L != nil {
			q.PushBack([2]*BinaryTreeNode{node[0].L, node[0]})
		}

		if node[0].R != nil {
			q.PushBack([2]*BinaryTreeNode{node[0].R, node[0]})
		}

	}

	lastNode := last[0]
	lastParent := last[1]

	target.Data = lastNode.Data

	if lastParent != nil {
		if lastParent.L == lastNode {
			lastParent.L = nil
		} else {
			lastParent.R = nil
		}

	} else {
		return nil
	}

	return root
}

func (root *BinaryTreeNode) InsertNode(data dataType) *BinaryTreeNode {
	if root == nil {
		return &BinaryTreeNode{Data: data}
	}

	q := list.New()

	q.PushBack(root)

	for q.Len() != 0 {
		temp := q.Front()

		q.Remove(temp)

		node := temp.Value.(*BinaryTreeNode)

		if node.L == nil {
			node.L = &BinaryTreeNode{Data: data}
			break
		} else {
			q.PushBack(node.L)
		}

		if node.R == nil {
			node.R = &BinaryTreeNode{Data: data}
			break
		} else {
			q.PushBack(node.R)
		}

	}

	return root
}

func (root *BinaryTreeNode) SearchDFS(data dataType) bool {
	if root == nil {
		return false
	}

	if root.Data == data {
		return true
	}

	leftRes := root.L.SearchDFS(data)
	rightRes := root.R.SearchDFS(data)

	return leftRes || rightRes
}

func processNode(node *BinaryTreeNode) {
	fmt.Println(node.Data, " ")
}

func PreorderTraversalCurrLeftRight(root *BinaryTreeNode) {
	if root == nil {
		return
	}

	q := list.New()

	q.PushBack(root)

	for q.Len() != 0 {
		temp := q.Front()

		q.Remove(temp)

		node := temp.Value.(*BinaryTreeNode)

		processNode(node)

		if node.L != nil {
			q.PushBack(node.L)
		}

		if node.R != nil {
			q.PushBack(node.R)
		}
	}

}

func InorderTraversalLeftCurrRight(root *BinaryTreeNode) {
	if root == nil {
		return
	}

	InorderTraversalLeftCurrRight(root.L)

	processNode(root)

	InorderTraversalLeftCurrRight(root.R)
}

func PostorderTraversalLeftRightCurr(root *BinaryTreeNode) {
	if root == nil {
		return
	}

	PostorderTraversalLeftRightCurr(root.L)

	PostorderTraversalLeftRightCurr(root.R)

	processNode(root)

}
