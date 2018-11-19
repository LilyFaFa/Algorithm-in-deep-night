package main

import (
	"sync"
)

var count int

func Add(lock *sync.Mutex, c chan int) {
	lock.Lock()
	lock.Unlock()
}

func main() {
	c := make(chan int, 3)
	for i := 0; i < 10; i++ {

	}
}
