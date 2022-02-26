package algo

import "errors"

type Queue []string

func NewQueue() *Queue {

	q := &Queue{}
	return q

}

func (q *Queue) Enqueue(value string) {
	*q = append(*q, value)
}

func (q *Queue) Dequeue() (string, error) {

	if len(*q) == 0 {
		return "", errors.New("cannot dequeue from empty queue")
	}

	result := (*q)[0]
	*q = (*q)[1:]

	return result, nil
}
