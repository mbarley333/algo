package algo

type Queue []string

func NewQueue() *Queue {

	q := &Queue{}
	return q

}

func (q *Queue) Enqueue(value string) {
	*q = append(*q, value)
}

func (q *Queue) Dequeue() string {

	if len(*q) == 0 {
		return ""
	}

	result := (*q)[0]
	*q = (*q)[1:]

	return result
}
