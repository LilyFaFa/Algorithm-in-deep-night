package collections

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
	if _, ok := h.hashMap[hashnum]; !ok {
		h.hashMap[hashnum] = []interface{}{item}
	} else {
		h.hashMap[hashnum] = append(h.hashMap[hashnum], item)
	}
}

func (h )