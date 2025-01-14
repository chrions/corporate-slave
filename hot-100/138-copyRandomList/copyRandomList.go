package fyf

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

var m map[*Node]*Node

func copyRandomList(head *Node) *Node {
	m = map[*Node]*Node{}
	return deepCopy(head)
}

func deepCopy(node *Node) *Node {
	if node == nil {
		return nil
	}

	if v, ok := m[node]; ok {
		return v
	}
	newNode := &Node{
		Val: node.Val,
	}
	m[node] = newNode
	newNode.Next = deepCopy(node.Next)
	newNode.Random = deepCopy(node.Random)
	return newNode

}
