package algo

import (
	"fmt"
	"io"
)

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

func (l *LinkedList) Delete(index int) error {

	if index == 0 {
		l.List = l.List.Next
		return nil
	}

	current_node := l.List
	current_index := 0

	for current_index < index-2 {
		current_node = current_node.Next
		current_index += 1

	}

	node_after_delete := current_node.Next.Next

	current_node.Next = node_after_delete

	return nil
}

func (l *LinkedList) PrintLinkedList(w io.Writer) {

	for l.List != nil {
		fmt.Fprintf(w, l.List.Data)
		l.List = l.List.Next

	}

}

func (l *LinkedList) PrintLinkedListLast(w io.Writer) {

	for l.List != nil {

		if l.List.Next == nil {
			fmt.Fprintf(w, l.List.Data)
		}
		l.List = l.List.Next

	}

}

func (l *LinkedList) ReverseLinkedList() {

	// track 3 vars
	// currentNode is set to starting node
	// previousNode is set to empty
	currentNode := l.List
	previousNode := &Node{}
	nextNode := &Node{}

	// loop until you reach the end
	for currentNode != nil {

		// next becomes next next
		// A -> B
		// B -> C
		nextNode = currentNode.Next

		// reverse direction Next points to prev
		// A Next = nil
		// B Next = A
		// C Next = B
		currentNode.Next = previousNode

		// shift right for next pass -- previous becomes current
		// nil -> A
		// A -> B
		previousNode = currentNode

		// shift right for next pass -- current becomes next
		// A -> B
		// B -> C
		// C -> nil
		currentNode = nextNode

	}
	// repoint final element since loop breaks on nil
	l.List = previousNode

}

func RemoveFromMiddle(node *Node) {

	node.Data = node.Next.Data
	node.Next = node.Next.Next

}
