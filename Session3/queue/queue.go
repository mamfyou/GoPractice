package queue

import "errors"

type Queue struct{
	max_size int
	items []int
}

var ErrQueueIsFull = errors.New("Queue is Full")
var ErrQueueIsEmpty = errors.New("Queue is Empty")


func NewQueue(max_size int) *Queue {
	return &Queue{max_size: max_size}
}

func (q *Queue) Dequeue() (elem int, err error) {
	if len(q.items) == 0 {
		return elem, ErrQueueIsEmpty
	} else if len(q.items) > 1{
		elem = q.items[0]
		q.items = q.items[1:]
	} else {
		elem = q.items[0]
		q.items = []int{}
	}
	return elem, nil
}

func (q *Queue) Enqueue(number int) (err error) {
	if len(q.items) >= q.max_size {
		return ErrQueueIsFull
	} else {
		q.items = append(q.items, number)
	}
	return nil
}