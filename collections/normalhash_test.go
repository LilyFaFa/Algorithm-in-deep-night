package collections

import (
	"fmt"
	"testing"
)

func hashfunc(item interface{}) int {
	value, ok := item.(int)
	if ok {
		return value
	}
	return 0
}

func TestNormalHash(t *testing.T) {
	item := 6
	h := NewNormalHash(hashfunc, 3)
	h.Add(item)
	exist := h.Check(item)
	if exist {
		fmt.Println("hello,add")
	}

	h.Delete(item)
	exist = h.Check(item)
	if !exist {
		fmt.Println("hello,delete")
	}

}
