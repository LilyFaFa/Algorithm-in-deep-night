package main

import (
	"fmt"
	"sync"
	"time"
)

var count int
var wg sync.WaitGroup

func main() {
	//wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println(i)
			//wg.Done()
			time.Sleep(2 * time.Second)
		}()

	}
	//wg.Wait()
}
