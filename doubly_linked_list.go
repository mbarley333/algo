package algo

import "fmt"

type DNode struct {
	Data         string
	NextNode     *DNode
	PreviousNode *DNode
}

type DoublyLinkedList struct {
	FirstNode *DNode
	LastNode  *DNode
}

func (d *DoublyLinkedList) InsertAtEnd(data string) {

	newNode := &DNode{
		Data: data,
	}

	if d.FirstNode == nil {
		d.FirstNode = newNode
		d.LastNode = newNode

	} else {
		newNode.PreviousNode = d.LastNode
		d.LastNode.NextNode = newNode
		d.LastNode = newNode

	}

	fmt.Println(d)

}
