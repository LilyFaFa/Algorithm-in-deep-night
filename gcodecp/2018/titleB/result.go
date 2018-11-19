package main

import (
	"bufio"
	//	"fmt"
	"io"
	"os"
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

func (mo *MyOutPut) writeFloat(i float64) {
	s := strconv.FormatFloat(i, 'f', -1, 64)
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

func (mi *MyInput) readByteSlice() []byte {
	line := mi.readLine()
	bytes := []byte{}
	bytesStr := strings.Split(line, " ")
	for _, b := range bytesStr {
		bytes = append(bytes, []byte(b)[0])
	}
	return bytes
}

type rh struct {
	r int
	h int
}

func main() {
	fi, _ := os.Open("B-large-practice.in")
	mi := MyInput{rdr: fi}
	fo, _ := os.OpenFile("B-large-practice.out", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	mo := MyOutPut{wrr: fo}
	//mi := MyInput{rdr: os.Stdin}
	//mo := MyOutPut{wrr: os.Stdout}

	count := mi.readInt()
	for caseNo := 0; caseNo < count; caseNo++ {
		inputs := mi.readIntSlice()
		N := inputs[0]
		C := inputs[1]
		M := inputs[2]

		user := make([]int, C+1)
		sit := make([]int, N+1)
		for i := 0; i < M; i++ {
			inp := mi.readIntSlice()
			sit[inp[0]]++
			user[inp[1]]++
		}
		result := 0
		for i := 1; i < C+1; i++ {
			if user[i] > result {
				result = user[i]
			}
		}

		totle := make([]int, N+1)
		for i := 1; i < N+1; i++ {
			totle[i] = totle[i-1] + sit[i]
		}

		for i := 1; i < N+1; i++ {
			if totle[i] > result*i {
				if totle[i]%i == 0 {
					result = totle[i] / i
				} else {
					result = totle[i]/i + 1
				}
			}
		}

		exchange := 0
		for i := 1; i < N+1; i++ {
			if sit[i] > result {
				exchange = exchange + sit[i] - result
			}
		}

		c := strconv.Itoa(caseNo + 1)
		mo.writeString("Case #" + c + ":")
		mo.writeString(" ")
		mo.writeInt(result)
		mo.writeString(" ")
		mo.writeInt(exchange)
		mo.writeString("\n")

	}
}
