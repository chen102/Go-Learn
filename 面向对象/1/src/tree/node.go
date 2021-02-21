package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

//工厂函数
func CreateNode(Value int) *Node {
	return &Node{Value: Value}
}
func (node Node) Print() {
	fmt.Println(node.Value)
}
func (node *Node) Setvalue(Value int) {
	node.Value = Value
}

//nil也可以调用方法
func (node *Node) Setvalue1(Value int) {
	if node == nil {
		fmt.Println("空的")
		return //防止panic
	}
	node.Value = Value
}
func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}
