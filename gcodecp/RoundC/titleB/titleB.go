package main

import (
	"bufio"
	//"errors"
	//"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type MyOutPut struct {
	wrr         io.Writer
	initialized bool
	lineChan    chan string
}
type MyInput struct {
	rdr         io.Reader
	lineChan    chan string
	initialized bool
}

func (mo *MyOutPut) start(done chan struct{}) {
	w := bufio.NewWriter(mo.wrr)
	defer func() { close(mo.lineChan) }()

	for {
		if !mo.initialized {
			mo.initialized = true
			done <- struct{}{}
		}
		res, ok := <-mo.lineChan

		if !ok {
			panic("trying to read from a closed channel")
		}
		_, err := w.WriteString(res)
		if err != nil {
			panic(err)
		}
		err = w.Flush()
		if err != nil {
			panic(err)
		}
	}

}

func (mo *MyOutPut) writeLine(s string) {
	// if this is the first call, initialize
	if !mo.initialized {
		mo.lineChan = make(chan string)
		done := make(chan struct{})
		go mo.start(done)
		<-done
	}
	mo.lineChan <- s
}

func (mo *MyOutPut) writeInt(i int) {
	s := strconv.Itoa(i)
	mo.writeLine(s)
}

func (mo *MyOutPut) writeString(s string) {
	mo.writeLine(s)
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
	fi, err := os.Open("B-large-practice.in")
	if err != nil {
		panic(err)
	}
	mi := MyInput{rdr: fi}
	fo, _ := os.OpenFile("B-large-practice.out", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	mo := MyOutPut{wrr: fo}

	// mi := MyInput{rdr: os.Stdin}
	// mo := MyOutPut{wrr: os.Stdout}
	count := mi.readInt()
	for caseNo := 0; caseNo < count; caseNo++ {
		var numbers [][]int
		N := mi.readInt()
		for i := 0; i < N; i++ {
			number := mi.readIntSlice()
			numbers = append(numbers, number)
		}
		result := 0
		tmp := []int{}
		used := make(map[int]bool)

		backtrack(numbers, tmp, 0, used, &result, true)
		c := strconv.Itoa(caseNo + 1)
		mo.writeString("Case #" + c + ": ")
		mo.writeInt(result)
		mo.writeString("\n")
	}
}

func backtrack(numbers [][]int, tmp []int, last int, used map[int]bool, result *int, check bool) {
	if last > len(numbers) {
		return
	}

	if len(tmp) >= 3 && check {
		sort.Ints(tmp)
		sum := 0
		for i := 0; i < len(tmp)-1; i++ {
			sum += tmp[i]
		}
		if tmp[len(tmp)-1] < sum {
			*result++
		}
	}

	if last == len(numbers) {
		return
	}

	_, ok := used[last]
	if ok {
		t := make([]int, len(tmp))
		copy(t, tmp)
		u := make(map[int]bool)
		for k, v := range used {
			u[k] = v
		}
		backtrack(numbers, t, last+1, u, result, false)
		return
	}

	for i := last; i < len(numbers); i++ {
		_, ok := used[i]
		if numbers[last][i] != 0 && !ok {
			t := make([]int, len(tmp))
			copy(t, tmp)
			t = append(t, numbers[last][i])
			u := make(map[int]bool)
			for k, v := range used {
				u[k] = v
			}
			u[i] = true
			u[last] = true
			backtrack(numbers, t, last+1, u, result, true)

		}
	}
	t2 := make([]int, len(tmp))
	copy(t2, tmp)
	u2 := make(map[int]bool)
	for k, v := range used {
		u2[k] = v
	}
	backtrack(numbers, t2, last+1, u2, result, false)
}
