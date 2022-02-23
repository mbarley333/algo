package algo

import "errors"

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

}

func (d *DoublyLinkedList) SearchForward(value string) (bool, error) {

	if d.FirstNode == nil {
		return false, errors.New("list is empty")
	}

	temp := d.FirstNode
	for temp != nil {
		if temp.Data == value {
			return true, nil
		}
		temp = temp.NextNode
	}

	return false, nil

}

func (d *DoublyLinkedList) SearchReverse(value string) (bool, error) {

	if d.LastNode == nil {
		return false, errors.New("list is empty")
	}

	temp := d.LastNode
	for temp != nil {
		if temp.Data == value {
			return true, nil
		}
		temp = temp.PreviousNode
	}

	return false, nil

}
