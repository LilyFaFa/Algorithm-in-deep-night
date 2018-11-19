package main

type MyCalendar struct {
	events map[int]int
}

func Constructor() MyCalendar {
	events := make(map[int]int)
	return MyCalendar{events}
}

func (this *MyCalendar) Book(start int, end int) bool {
	flag := true
	// 循环检测是否时间与所有的events不冲突
	for s, e := range this.events {
		if s >= end || start >= e {
			continue
		} else {
			flag = false
			break
		}
	}
	if flag {
		this.events[start] = end
	}
	return flag

}
