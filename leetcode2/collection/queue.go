package collection

import (
	"container/list"
)

type Queue struct {
	list *list.List
}

func NewQueue() *Queue {
	l := new(list.List)
	return &Queue{l}
}

func (q *Queue) Dequeue() (value interface{}) {
	if elem := q.list.Front(); elem != nil {
		q.list.Remove(elem)
		return elem.Value
	}
	return nil
}

func (q *Queue) Enqueue(value interface{}) {
	q.list.PushFront(value)
}
