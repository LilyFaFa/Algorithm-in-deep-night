package main

type ItemsList struct {
	val  string
	next *ItemsList
}

type FreList struct {
	fre  int
	next *FreList
	item *ItemsList
}

type AllOne struct {
	items   map[string]int
	allList *FreList
}

/** Initialize your data structure here. */
func Constructor() AllOne {

}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string) {
	value, ok := this.items[key]
	if ok {
		this.items[key] = value + 1
		find := this.allList
		for find != nil {
			// 删除操作
			if find.fre == value {
				i := find.item
				if i.val == key {
					find.item = i.next
				} else {
					for i.next != nil {
						if i.next.val == key {
							i.next = i.next.next
							break
						} else {
							i = i.next
						}
					}
				}
			} else {
				find = find.next
			}
			// 添加操作
			if find.next != nil && find.next.fre == value+1 {
				i := find.next.item
				if i.next != nil {
					i = i.next
				}
				item := ItemsList{
					val:  key,
					next: nil,
				}
				i.next = &item
			} else if find.next == nil {
				item := ItemsList{
					val:  key,
					next: nil,
				}
				fre := FreList{
					fre:  value + 1,
					next: nil,
					item: &item,
				}
				find.next = &fre
			}
		}
	} else {
		find := this.allList
		if find == nil {
			item := ItemsList{
				val:  key,
				next: nil,
			}
			fre := FreList{
				fre:  1,
				next: nil,
				item: &item,
			}
			find = &fre
		} else if find.fre != 1 {
			item := ItemsList{
				val:  key,
				next: nil,
			}
			fre := FreList{
				fre:  1,
				next: &find,
				item: &item,
			}
			this.allList = &fre
		}

	}
}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string) {
	value, ok := this.items[key]
	if ok {
		if value > 1 {
			this.items[key] = value - 1
		} else {
			delete(this.items, key)
		}
		find := this.allList
		for find != nil {
			//添加操作
			if value > 1 {
				if find.fre == value-1 {
					i := find.item
					if i.val == key {
						find.item = i.next
					} else {
						for i.next != nil {
							if i.next.val == key {
								i.next = i.next.next
								break
							} else {
								i = i.next
							}
						}
					}
				} else {
					find = find.next
				}
			}
			// 删除操作
			if find.fre == value {
				i := find.item
				if i.val == key {
					find.item = i.next
				} else {
					for i.next != nil {
						if i.next.val == key {
							i.next = i.next.next
							break
						} else {
							i = i.next
						}
					}
				}
			} else {
				find = find.next
			}
		}
	}
}

/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {

}

/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {

}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
