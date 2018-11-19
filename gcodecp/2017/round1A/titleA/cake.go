package main

import (
	"bufio"
	"fmt"
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

type exist struct {
	partners []int
	time     int
}

func main() {

	fi, _ := os.Open("A-large-practice.in")

	mi := MyInput{rdr: fi}
	fo, _ := os.OpenFile("A-large-practice.out", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	mo := MyOutPut{wrr: fo}
	//mi := MyInput{rdr: os.Stdin}
	//mo := MyOutPut{wrr: os.Stdout}
	count := mi.readInt()
	for caseNo := 0; caseNo < count; caseNo++ {
		var inputs [][]byte
		rc := mi.readIntSlice()
		rcount := rc[0]
		ccount := rc[1]
		split := []int{-1}
		kv := make(map[int][]int)
		for i := 0; i < rcount; i++ {
			count := 0
			in := mi.readLine()
			input := []byte(in)
			for j := 0; j < len(input); j++ {
				if input[j] != '?' {
					count++
					if count == 1 {
						kv[i] = []int{-1, j}
					} else {
						kv[i] = append(kv[i], j)
					}
				}
			}
			if count != 0 {
				split = append(split, i)
			}
			inputs = append(inputs, input)
		}
		fmt.Println(inputs)
		fmt.Println(split)
		fmt.Println(kv)
		//按行切
		i := 1
		for i = 1; i < len(split); i++ {
			j := 1
			for j = 1; j < len(kv[split[i]]); j++ {
				for a := split[i-1] + 1; a <= split[i]; a++ {
					for b := kv[split[i]][j-1] + 1; b <= kv[split[i]][j]; b++ {
						inputs[a][b] = inputs[split[i]][kv[split[i]][j]]
					}
				}
			}
			if kv[split[i]][j-1] != ccount-1 {
				for k := kv[split[i]][j-1] + 1; k < ccount; k++ {

					for a := split[i-1] + 1; a <= split[i]; a++ {
						inputs[a][k] = inputs[a][k-1]
					}
				}
			}
		}
		if split[i-1] != rcount-1 {
			for k := split[i-1] + 1; k < rcount; k++ {
				for a := 0; a < ccount; a++ {
					inputs[k][a] = inputs[k-1][a]
				}
			}
		}
		fmt.Println(inputs)

		c := strconv.Itoa(caseNo + 1)
		mo.writeString("Case #" + c + ":")
		mo.writeString("\n")
		for i := 0; i < rcount; i++ {
			result := ""
			for j := 0; j < ccount; j++ {
				result = result + string(inputs[i][j])
			}
			mo.writeString(result)
			mo.writeString("\n")
		}
		//mo.writeString("\n")
	}

}
