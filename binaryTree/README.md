# How use it
```go
var bt *binaryTree.BinaryTreeNode

	for i := range 12 {

		if bt == nil {
			bt = bt.InsertNode(i)
			continue
		}

		bt.InsertNode(i)
	}

	binaryTree.PreorderTraversalCurrLeftRight(bt)

	fmt.Println("DeleteNode 2", bt.DeleteNode(2))

	binaryTree.PreorderTraversalCurrLeftRight(bt)
```
