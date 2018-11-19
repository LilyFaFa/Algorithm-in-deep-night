package main

import (
	"bufio"
	"fmt"
	"sort"
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
	var PI float64
	PI = 3.14159265358979323846
	fi, _ := os.Open("A-large-practice.in")
	mi := MyInput{rdr: fi}
	fo, _ := os.OpenFile("A--practice.out", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	mo := MyOutPut{wrr: fo}
	//mi := MyInput{rdr: os.Stdin}
	//mo := MyOutPut{wrr: os.Stdout}

	count := mi.readInt()
	for caseNo := 0; caseNo < count; caseNo++ {
		inputs := mi.readIntSlice()
		N := inputs[0]
		K := inputs[1]

		mid := make(map[int][]int)
		keys := []int{}
		items := []rh{}

		var result [][]int
		for j := 0; j < N; j++ {
			l := make([]int, K)
			result = append(result, l)
		}

		for i := 0; i < N; i++ {
			item := mi.readIntSlice()
			r := item[0]
			h := item[1]
			_, ok := mid[r]
			if ok {
				mid[r] = append(mid[r], h)
			} else {
				keys = append(keys, r)
				l := []int{h}
				mid[r] = l
			}
		}
		sort.Ints(keys)
		for i := 0; i < len(keys)/2; i++ {
			tmp := keys[i]
			keys[i] = keys[len(keys)-i-1]
			keys[len(keys)-i-1] = tmp
		}
		for _, i := range keys {
			for _, j := range mid[i] {
				items = append(items, rh{r: i, h: j})
			}
		}
		fmt.Println(N)
		fmt.Println(items)
		for index, item := range items {
			result[index][0] = item.r*item.r + 2*item.r*item.h
			if index == 0 {
				continue
			} else if result[index-1][0] > result[index][0] {
				result[index][0] = result[index-1][0]
			}
			for i := 1; i <= index; i++ {
				max := K - 1
				if i < K-1 {
					max = i
				}
				for k := 1; k <= max; k++ {
					if result[index-1][k-1]+2*item.r*item.h > result[index-1][k] {
						result[index][k] = result[index-1][k-1] + 2*item.r*item.h
					} else {
						result[index][k] = result[index-1][k]
					}
				}

			}
		}

		var res float64
		var m int
		for i := 0; i < N; i++ {
			if result[i][K-1] > m {
				m = result[i][K-1]
			}
		}
		res = float64(m) * PI

		c := strconv.Itoa(caseNo + 1)
		mo.writeString("Case #" + c + ":")
		mo.writeString(" ")
		mo.writeFloat(res)
		mo.writeString("\n")

	}
}
