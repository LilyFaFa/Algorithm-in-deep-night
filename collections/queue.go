package collections

import (
	"errors"
)

type Queue struct {
	Items []interface{}
}

func (q *Queue) Dequeue() (interface{}, error) {
	//current queue is empty
	if len(q.Items) == 0 {
		return nil, errors.New("queue is empty")
	}

	item := q.Items[0]
	q.Items = q.Items[1:]
	return item, nil
}

func (q *Queue) Enqueue(item interface{}) {
	q.Items = append(q.Items, item)
}

func (q *Queue) Size() int {
	return len(q.Items)
}

func (q *Queue) IsEmpty() bool {
	return len(q.Items) == 0
}
