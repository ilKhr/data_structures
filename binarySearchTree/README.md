# How use it
```go
	var bst *binarySearchTree.BSTNode

	bst = bst.InsertNode(50)
	bst.InsertNode(30)
	bst.InsertNode(70)
	bst.InsertNode(40)
	bst.InsertNode(60)
	bst.InsertNode(80)

	binarySearchTree.DrawTree(bst)

	bst.DeleteNode(50)
	binarySearchTree.DrawTree(bst)
```
