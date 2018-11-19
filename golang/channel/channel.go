package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	// 1 使用channl通信
	/*
		ch := make(chan int)
		wg.Add(1)

		go func() {
			cur := <-ch
			fmt.Println("receive", cur)
			wg.Done()
		}()

		go func() {
			fmt.Println("send", 42)
			ch <- 42
		}()

		wg.Wait()
	*/
	/*
		ch1 := make(chan int)
		ch2 := make(chan int)

		go func(ch chan int) {
			ch <- 1
		}(ch1)

		go func(ch chan int) {
			ch <- 2
		}(ch2)

		select {
		case <-ch1:
			fmt.Println("first")
		case <-ch2:
			fmt.Println("second")
		}
	*/
	ch := make(chan int)
	inchannel(ch)
	for i := 0; i < 5; i++ {
		k := <-ch
		fmt.Println(k)
	}
}

func inchannel(ch chan int) {

	mem := []int{1, 2, 3, 4, 5}

	for _, i := range mem {
		go func() {
			fmt.Println(i)
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("done")
	for _, i := range mem {
		fmt.Println(i)
		go func() {
			ch <- i
		}()
	}

}
