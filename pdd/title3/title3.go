package main

import (
	"bufio"
	"fmt"

	"io"
	"os"
	"strconv"
	"strings"
)

type MyInput struct {
	rdr         io.Reader
	lineChan    chan string
	initialized bool
}

func (mi *MyInput) start(done chan struct{}) {
	r := bufio.NewReader(mi.rdr)
	defer func() { close(mi.lineChan) }()

	for {
		line, err := r.ReadString('\n')
		if !mi.initialized {
			mi.initialized = true
			done <- struct{}{}
		}
		// 忽略空行
		l := strings.TrimSpace(line)
		if l == "" {
			continue
		}

		mi.lineChan <- l
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
	}
}

func (mi *MyInput) readLine() string {
	// if this is the first call, initialize
	if !mi.initialized {
		mi.lineChan = make(chan string)
		done := make(chan struct{})
		go mi.start(done)
		<-done
	}

	res, ok := <-mi.lineChan
	if !ok {
		panic("trying to read from a closed channel")
	}
	return res
}

func (mi *MyInput) readInt() int {
	line := mi.readLine()
	i, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	return i
}

func (mi *MyInput) readIntSlice() []int {
	line := mi.readLine()
	numbers := []int{}
	numbersStr := strings.Split(line, " ")
	for _, n := range numbersStr {
		i, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, i)
	}
	return numbers
}

func main() {

	mi := MyInput{rdr: os.Stdin}

	in := mi.readIntSlice()
	number := in[0]
	member := in[1]

	var totle [][]int
	for i := 0; i < number; i++ {
		in := mi.readIntSlice()
		n := make([]int, number)
		for _, j := range in {
			n[j] = 1
		}
		totle = append(totle, n)
	}

	max := 0
	friend := 0

	for i := 0; i < number; i++ {
		if i == member {
			continue
		}

		l := 0
		for j := 0; j < number; j++ {
			if totle[i][j] == 1 && totle[member][j] == 1 {
				l++

				if l > max {
					max = l
					friend = i
				}
			}
		}
	}
	fmt.Println(totle)
	if max == 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(friend)
	}

}
