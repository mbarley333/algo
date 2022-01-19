package algo

type Node struct {
	Data string
	Next *Node
}

type LinkedList struct {
	List *Node
}

func (l LinkedList) Read(index int) string {
	current_node := l.List
	current_index := 0

	for current_index < index {
		current_node = current_node.Next
		current_index += 1
	}
	return current_node.Data
}
