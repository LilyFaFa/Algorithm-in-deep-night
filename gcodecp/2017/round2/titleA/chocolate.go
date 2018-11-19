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
	fi, _ := os.Open("A-large-practice.in")
	mi := MyInput{rdr: fi}
	fo, _ := os.OpenFile("A-large-practice.out", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	mo := MyOutPut{wrr: fo}
	//mi := MyInput{rdr: os.Stdin}
	//mo := MyOutPut{wrr: os.Stdout}

	count := mi.readInt()
	for caseNo := 0; caseNo < count; caseNo++ {
		inputs := mi.readIntSlice()
		//		N := inputs[0]
		P := inputs[1]
		result := 0
		if P == 2 {
			result = P2(mi.readIntSlice())
		} else if P == 3 {
			result = P3(mi.readIntSlice())
		} else if P == 4 {
			result = P4(mi.readIntSlice())
		}

		c := strconv.Itoa(caseNo + 1)
		mo.writeString("Case #" + c + ":")
		mo.writeString(" ")
		mo.writeInt(result)
		mo.writeString("\n")

	}
}

func P2(input []int) int {
	middle := make(map[int]int)
	middle[0] = 0
	middle[1] = 0
	for _, item := range input {
		if item%2 == 0 {
			middle[0]++
		} else {
			middle[1]++
		}
	}
	if middle[1]%2 == 0 {
		return middle[0] + middle[1]/2
	} else {
		return middle[0] + middle[1]/2 + 1
	}
}

func P3(input []int) int {
	middle := make(map[int]int)
	middle[0] = 0
	middle[1] = 0
	middle[2] = 0
	for _, item := range input {
		if item%3 == 0 {
			middle[0]++
		} else if item%3 == 1 {
			middle[1]++
		} else if item%3 == 2 {
			middle[2]++
		}
	}

	max := 0
	min := 0
	if middle[1] > middle[2] {
		max = middle[1]
		min = middle[2]
	} else {
		max = middle[2]
		min = middle[1]
	}

	if (max-min)%3 == 0 {
		return middle[0] + min + (max-min)/3
	} else {
		return middle[0] + min + (max-min)/3 + 1
	}

}

func P4(input []int) int {
	middle := make(map[int]int)
	middle[0] = 0
	middle[1] = 0
	middle[2] = 0
	middle[3] = 0
	for _, item := range input {
		if item%4 == 0 {
			middle[0]++
		} else if item%4 == 1 {
			middle[1]++
		} else if item%4 == 2 {
			middle[2]++
		} else if item%4 == 3 {
			middle[3]++
		}
	}

	max := 0
	min := 0
	if middle[1] > middle[3] {
		max = middle[1]
		min = middle[3]
	} else {
		max = middle[3]
		min = middle[1]
	}

	if middle[2]%2 == 0 {
		if (max-min)%4 == 0 {
			return middle[0] + min + (max-min)/4 + middle[2]/2
		} else {
			return middle[0] + min + (max-min)/4 + 1 + middle[2]/2
		}
	} else {
		if max > min+1 {
			if (max-min-2)%4 == 0 {
				return middle[0] + min + 1 + (max-min-2)/4 + middle[2]/2
			} else {
				return middle[0] + min + 1 + (max-min-2)/4 + 1 + middle[2]/2
			}
		} else {
			return middle[0] + middle[2]/2 + min + 1
		}
	}

}
