package algo

type Node struct {
	Data string
	Next *Node
}

// holds the first node
type LinkedList struct {
	List *Node
}

func (l LinkedList) Read(index int) string {
	current_node := l.List
	current_index := 0
	value := ""

	for current_index < index && current_node.Next != nil {
		current_node = current_node.Next
		current_index += 1
		value = current_node.Data

	}

	if index > current_index {
		value = "NotFound"
	}
	return value

}

func (l LinkedList) Search(value string) int {

	current_node := l.List
	current_index := 0
	isFound := false
	index := -1

	for !isFound {

		if current_node.Next == nil {
			break
		}

		if current_node.Data == value {
			isFound = true
			index = current_index
		}

		current_index += 1
		current_node = current_node.Next

	}

	return index
}

func (l *LinkedList) Insert(index int, data string) error {
	newNode := Node{
		Data: data,
	}

	if index == 0 {
		newNode.Next = l.List
		l.List = &newNode
		return nil
	}

	current_node := l.List
	current_index := 0

	for current_index < index-2 {
		current_node = current_node.Next
		current_index += 1

	}

	newNode.Next = current_node.Next
	current_node.Next = &newNode

	return nil
}
