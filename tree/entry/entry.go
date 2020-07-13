package main

import (
	"fmt"
	"learngo/tree"
)

type myTreeNode struct {
	*tree.Node // Embedding
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.Node == nil {
		return
	}
	left := myTreeNode{myNode.Left}
	left.postOrder()
	right := myTreeNode{myNode.Right}
	right.postOrder()
	myNode.Print()
}

func (myNode *myTreeNode) Traverse() {		// 重载
	fmt.Println("This method is shadowed.")
}

func main() {
	root := myTreeNode{&tree.Node{Value: 3}}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	root.Traverse()			// 重载的方法
	root.Node.Traverse()	// 基类的方法
	fmt.Println()
	root.postOrder()
	fmt.Println()

	root.Traverse()

	//var baseRoot *tree.Node
	//baseRoot := &root		不能把子类的指针赋值给父类

	c := root.TraverseWithChannel()
	maxNode := 0
	for node := range c {
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}
	fmt.Println("Max node value:", maxNode)
}
