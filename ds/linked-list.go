package ds

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (list *LinkedList) Prepend(data int) {
	node := &Node{Data: data}

	node.Next = list.Head
	list.Head = node
}

func (list *LinkedList) Append(data int) {
	node := &Node{Data: data}

	if list.Head == nil {
		list.Head = node
		return
	}

	current := list.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = node
}
