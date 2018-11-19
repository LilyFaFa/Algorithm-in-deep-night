package main

import (
	"fmt"
	"sync"
)

/*
var l sync.Mutex
var str string

func f() {
	a := "hello world"
	fmt.Println(a)
	l.Unlock()
}

func main() {
	l.Lock()
	go f()
	l.Lock()
	fmt.Println("hello")
	l.Unlock()
}
*/
var str string
var wg sync.WaitGroup

func setup() {
	str = "hello"
}

func doprint() {
	once := sync.Once{}
	once.Do(setup)
	fmt.Println(str)
	wg.Done()
}

func main() {
	wg.Add(3)
	for i := 0; i < 10; i++ {
		go doprint()
	}
	wg.Wait()
}
