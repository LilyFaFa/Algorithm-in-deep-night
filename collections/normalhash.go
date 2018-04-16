package collections

import (
	"errors"
)

const DefaultZone = 10

type Hashfunc func(interface{}) int

type NormalHash struct {
	hashfunc Hashfunc
	zone     int
	hashMap  map[int]([]interface{})
}

func NewNormalHash(hashfunc Hashfunc, zone int) *NormalHash {
	hashmap := make(map[int][]interface{})
	if zone <= 0 {
		zone = DefaultZone
	}
	return &NormalHash{
		hashfunc: hashfunc,
		zone:     zone,
		hashMap:  hashmap,
	}
}

func (h *NormalHash) Add(item interface{}) {
	hashnum := h.hashfunc(item)
	if _, ok := h.hashMap[hashnum%h.zone]; !ok {
		h.hashMap[hashnum%h.zone] = []interface{}{item}
	} else {
		h.hashMap[hashnum%h.zone] = append(h.hashMap[hashnum], item)
	}
}

func (h *NormalHash) Delete(item interface{}) error {
	hashnum := h.hashfunc(item)
	if _, ok := h.hashMap[hashnum%h.zone]; !ok {
		return errors.New("can not find item")
	} else {
		items := h.hashMap[hashnum%h.zone]
		for j, i := range items {
			if i == item {
				items = append(items[:j], items[j+1:]...)
			}
		}
		h.hashMap[hashnum%h.zone] = items
	}
	return nil
}

func (h *NormalHash) Check(item interface{}) bool {
	hashnum := h.hashfunc(item)
	items, ok := h.hashMap[hashnum%h.zone]
	if !ok || len(items) == 0 {
		return false
	}
	return true
}
