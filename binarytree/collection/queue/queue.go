package queue

import (
	"errors"
)

type Queue struct {
	Item []interface{}
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) DeQueue() (interface{}, error) {
	if len(q.Item) == 0 {
		return nil, errors.New("The queue is empty")
	}
	item := q.Item[0]
	q.Item = q.Item[1:]
	return item, nil
}

func (q *Queue) EnQueue(i interface{}) {
	q.Item = append(q.Item, i)

}
