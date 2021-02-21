package main

import (
	"fmt"
	"queue"
	"tree"
)

//扩展已有类型 组合方式
type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

//duck typing
//描述事物的外部行为而非内部结构
func main() {
	fmt.Println("vim-go")
	root := tree.Node{Value: 3}
	root.Print()
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.Setvalue(4)
	root.Traverse()
	fmt.Println()
	myroot := myTreeNode{&root}
	myroot.postOrder()

	q := queue.Queue{1}
	fmt.Println(q.IsEmpty())
	q.Push(5)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	//var root Node
	//root = Node{value: 3}
	//root.left = &Node{}
	//root.right = &Node{5, nil, nil}
	//root.right.left = new(Node)
	//nodes := []Node{
	//{value: 3},
	//{},
	//{6, nil, &root},
	//}
	//fmt.Println(nodes)
	//root.left.right = CreateNode(2)
	//fmt.Println(root.left.right)
	//root.Print()
	//root.right.Setvalue(2)
	//root.right.Print()
	//var root1 *Node
	//root1.Setvalue1(22) //把nil传进去
	//root1 = &root
	//root1.Setvalue1(24)
	//root1.Print()
	//fmt.Println("-----------------")
	//root.Traverse()
}
