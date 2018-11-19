package main

import (
	"container/list"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
)

type cacheItem struct {
	Key string
	Val interface{}
}

type LRU struct {
	// 最大存储数量
	maxNum int
	// 当前存储数量
	curNum int
	// 锁，保证数据一致性
	mutex sync.Mutex
	// 链表
	data *list.List
}

//添加数据
func (l *LRU) add(key string, value interface{}) error {
	//判断key是否存在
	if e, _ := l.exist(key); e {
		return errors.New(key + "已存在")
	}
	//判断当前存储数量与最大存储数量
	if l.maxNum == l.curNum {
		//链表已满，则删除链表尾部元素
		l.clear()
	}
	l.mutex.Lock()
	l.curNum++
	//json序列化数据
	data, _ := json.Marshal(cacheItem{key, value})
	//把数据保存到链表头部
	l.data.PushFront(data)
	l.mutex.Unlock()
	return nil
}

//设置数据
func (l *LRU) set(key string, value interface{}) error {
	e, item := l.exist(key)
	if !e {
		return l.add(key, value)
	}
	l.mutex.Lock()
	data, _ := json.Marshal(cacheItem{key, value})
	//设置链表元素数据
	item.Value = data
	l.mutex.Unlock()
	return nil
}

//清理数据
func (l *LRU) clear() interface{} {
	l.mutex.Lock()
	l.curNum--
	//删除链表尾部元素
	v := l.data.Remove(l.data.Back())
	l.mutex.Unlock()
	return v
}

//获取数据
func (l *LRU) get(key string) interface{} {
	e, item := l.exist(key)
	if !e {
		return nil
	}
	l.mutex.Lock()
	//数据被访问，则把元素移动到链表头部
	l.data.MoveToFront(item)
	l.mutex.Unlock()
	var data cacheItem
	json.Unmarshal(item.Value.([]byte), &data)
	return data.Val
}

//删除数据
func (l *LRU) del(key string) error {
	e, item := l.exist(key)
	if !e {
		return errors.New(key + "不存在")
	}
	l.mutex.Lock()
	l.curNum--
	//删除链表元素
	l.data.Remove(item)
	l.mutex.Unlock()
	return nil
}

//判断是否存在
func (l *LRU) exist(key string) (bool, *list.Element) {
	var data cacheItem
	//循环链表，判断key是否存在
	for v := l.data.Front(); v != nil; v = v.Next() {
		json.Unmarshal(v.Value.([]byte), &data)
		if key == data.Key {
			return true, v
		}
	}
	return false, nil
}

//返回长度
func (l *LRU) len() int {
	return l.curNum
}

//打印链表
func (l *LRU) print() {
	var data cacheItem
	for v := l.data.Front(); v != nil; v = v.Next() {
		json.Unmarshal(v.Value.([]byte), &data)
		fmt.Println("key:", data.Key, " value:", data.Val)
	}
}

//创建一个新的LRU
func LRUNew(maxNum int) *LRU {
	return &LRU{
		maxNum: maxNum,
		curNum: 0,
		mutex:  sync.Mutex{},
		data:   list.New(),
	}
}
