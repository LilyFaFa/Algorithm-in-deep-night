package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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

func (mi *MyInput) readByteSlice() []byte {
	line := mi.readLine()
	bytes := []byte{}
	bytesStr := strings.Split(line, " ")
	for _, b := range bytesStr {
		bytes = append(bytes, []byte(b)[0])
	}
	return bytes
}

type mm struct {
	min int
	max int
}

type exist struct {
	partners []int
	time     int
}

func main() {

	fi, _ := os.Open("B-large-practice.in")

	mi := MyInput{rdr: fi}
	fo, _ := os.OpenFile("B-large-practice.out", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	mo := MyOutPut{wrr: fo}
	//	mi := MyInput{rdr: os.Stdin}
	//mo := MyOutPut{wrr: os.Stdout}
	count := mi.readInt()

	for caseNo := 0; caseNo < count; caseNo++ {
		result := 0

		rc := mi.readIntSlice()
		ncount := rc[0]
		nnumbers := mi.readIntSlice()
		np := [][]mm{}
		flag := false
		for i := 0; i < ncount; i++ {
			p := mi.readIntSlice()
			// 升序排列
			sort.Ints(p)
			fmt.Println(p)
			m := []mm{}
			for j := 0; j < len(p); j++ {
				//运算不准确，加减了一个很小的值
				minnumber := float64(p[j])/(1.1*float64(nnumbers[i])) - 0.000000000001
				maxnumber := float64(p[j])/(0.9*float64(nnumbers[i])) + 0.0000000000001
				//fmt.Println(minnumber, ":", maxnumber)
				if int(math.Floor(maxnumber)) < int(math.Ceil(minnumber)) {

				} else {
					s := mm{
						min: int(math.Ceil(minnumber)),
						max: int(math.Floor(maxnumber)),
					}
					m = append(m, s)
				}
			}
			if len(m) == 0 {
				flag = true
			} else {
				np = append(np, m)
			}
			if caseNo == 0 {
				fmt.Println(np)
			}
		}

		if flag == true {
			c := strconv.Itoa(caseNo + 1)
			mo.writeString("Case #" + c + ":")
			mo.writeString(" ")
			mo.writeInt(result)
			mo.writeString("\n")
			continue
		}

		//fmt.Println(np)

		ran := make([]int, ncount)
		maxmin := np[0][ran[0]].min
		minmax := np[0][ran[0]].max
		for {
			fg := false

			for i := 0; i < ncount; i++ {
				for np[i][ran[i]].max < maxmin {
					fmt.Println("jijijijijiji", maxmin, np[i][ran[i]].max)
					ran[i]++
					if ran[i] >= len(np[i]) {
						goto HERE
					}
				}

				if np[i][ran[i]].min > minmax {
					maxmin = np[i][ran[i]].min
					minmax = np[i][ran[i]].max
					fmt.Println("jijiji", i, np[i][ran[i]].min, minmax)
					fg = true
					break
				}

				if np[i][ran[i]].min > maxmin {
					maxmin = np[i][ran[i]].min
				}

				if np[i][ran[i]].max < minmax {
					minmax = np[i][ran[i]].max
				}

			}
			if fg == false {
				fmt.Println("hahahh")
				result++
				fmt.Println(ran)
				for k, m := range ran {
					if m == (len(np[k]) - 1) {
						fmt.Println("ooooooooooo", m)
						goto HERE
					} else {
						ran[k]++

					}
				}
				fmt.Println(ran)
				minmax = np[0][ran[0]].max
				maxmin = np[0][ran[0]].min
			}
		}
	HERE:
		c := strconv.Itoa(caseNo + 1)
		mo.writeString("Case #" + c + ":")
		mo.writeString(" ")
		mo.writeInt(result)
		mo.writeString("\n")

	}

}
