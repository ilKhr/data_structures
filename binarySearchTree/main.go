// https://www.geeksforgeeks.org/binary-search-tree-data-structure/
package binarySearchTree

import (
	"container/list"
	"fmt"
	"strings"
)

type dataType = int

type BSTNode struct {
	Data dataType
	L    *BSTNode
	R    *BSTNode
}

func (root *BSTNode) InsertNode(data dataType) *BSTNode {
	if root == nil {
		return &BSTNode{Data: data}
	}

	q := list.New()

	q.PushBack(root)

	for q.Len() != 0 {
		tmp := q.Front()

		q.Remove(tmp)

		node := tmp.Value.(*BSTNode)

		if data == node.Data {
			panic("Duplicate value")
		}

		if data < node.Data {

			if node.L == nil {
				node.L = &BSTNode{Data: data}
				break
			} else {
				if node.L != nil {
					q.PushBack(node.L)
				}
			}
		}

		if data > node.Data {
			if node.R == nil {
				node.R = &BSTNode{Data: data}
				break
			} else {
				if node.R != nil {
					q.PushBack(node.R)
				}
			}
		}

	}

	return root
}

func (root *BSTNode) DeleteNode(data dataType) *BSTNode {
	if root == nil {
		return root
	}

	if data > root.Data {
		root.R = root.R.DeleteNode(data)
	} else if data < root.Data {
		root.L = root.L.DeleteNode(data)
	} else {

		if root.L == nil {
			return root.R
		}

		if root.R == nil {
			return root.L
		}

		suc := root.GetSuccessorForDelete()

		root.Data = suc.Data
		root.R = root.R.DeleteNode(suc.Data)
	}

	return root
}

func (root *BSTNode) GetSuccessorForDelete() *BSTNode {

	if root == nil {
		return root
	}

	cur := root.R

	for cur.L != nil {
		cur = cur.L
	}

	return cur
}

func (root *BSTNode) GetMinimalNode() *BSTNode {

	if root == nil {
		return root
	}

	q := list.New()

	q.PushBack(root)

	for q.Len() != 0 {
		tmp := q.Front()

		q.Remove(tmp)

		node := tmp.Value.(*BSTNode)

		if node.L == nil && node.R == nil {
			return node
		}

		if node.L != nil {
			q.PushBack(node.L)
		}
	}

	return root
}

func (root *BSTNode) Inorder() {
	if root != nil {
		root.L.Inorder()
		fmt.Println(root.Data)
		root.R.Inorder()
	}
}

func (root *BSTNode) SearchDFS(data dataType) bool {
	if root == nil {
		return false
	}

	q := list.New()

	q.PushBack(root)

	for q.Len() != 0 {
		tmp := q.Front()

		q.Remove(tmp)

		node := tmp.Value.(*BSTNode)

		if data == node.Data {
			return true
		}

		if data < node.Data && node.L != nil {
			q.PushBack(node.L)
		}

		if data > node.Data && node.R != nil {
			q.PushBack(node.R)
		}

	}

	return false
}

func printTree(root *BSTNode, prefix string, isLeft bool) []string {
	if root == nil {
		return []string{}
	}

	var result []string
	if isLeft {
		result = append(result, fmt.Sprintf("%s├── %d", prefix, root.Data))
	} else {
		result = append(result, fmt.Sprintf("%s└── %d", prefix, root.Data))
	}

	childPrefix := prefix
	if isLeft {
		childPrefix += "│   "
	} else {
		childPrefix += "    "
	}

	if root.L != nil || root.R != nil {
		if root.L != nil {
			result = append(result, printTree(root.L, childPrefix, true)...)
		} else {
			result = append(result, fmt.Sprintf("%s├── nil", childPrefix))
		}

		if root.R != nil {
			result = append(result, printTree(root.R, childPrefix, false)...)
		} else {
			result = append(result, fmt.Sprintf("%s└── nil", childPrefix))
		}
	}

	return result
}

func DrawTree(root *BSTNode) {
	lines := printTree(root, "", false)
	fmt.Println(strings.Join(lines, "\n"))
}
