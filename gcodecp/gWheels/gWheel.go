package main

import (
	"bufio"
	"errors"
	//"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type MyInput struct {
	// 抽象接口
	rdr         io.Reader
	lineChan    chan string
	initialized bool
}

type MyOutPut struct {
	wrr         io.Writer
	initialized bool
	lineChan    chan string
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
	fi, _ := os.Open("A-large-practice.in")
	mi := MyInput{rdr: fi}
	//mi := MyInput{rdr: os.Stdin}
	//mo := MyOutPut{wrr: os.Stdout}
	fo, _ := os.OpenFile("A-large-practice.out", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	mo := MyOutPut{wrr: fo}
	count := mi.readInt()
	for caseNo := 0; caseNo < count; caseNo++ {
		//fmt.Printf("Case #%d:\n", caseNo+1)
		c := strconv.Itoa(caseNo + 1)
		mo.writeString("Case #" + c + ":" + " \n")
		petS := mi.readIntSlice()
		if len(petS) != 3 {
			panic(errors.New("error tmplate"))
		}
		Np := mi.readIntSlice()
		if len(Np) != petS[0] {
			panic(errors.New("error tmplate p"))
		}

		Ne := mi.readIntSlice()
		if len(Ne) != petS[1] {
			panic(errors.New("error tmplate e"))
		}

		Nt := mi.readIntSlice()
		if len(Nt) != petS[2] {
			panic(errors.New("error tmplate t"))
		}

		num := mi.readInt()
		tests := [][]int{}
		for j := 0; j < num; j++ {
			curr := mi.readIntSlice()
			if len(curr) != 2 {
				panic(errors.New("error tmplate"))
			}
			tests = append(tests, curr)
		}

		for po, test := range tests {
			flag := 0
			yes := "Yes"
			no := "No"
			var ettt [][]int
			var eptt [][]int
			for _, e := range Ne {
				var et []int
				for _, t := range Nt {
					et = append(et, test[0]*e*t)
				}
				ettt = append(ettt, et)

				var pt []int
				for _, p := range Np {
					pt = append(pt, test[1]*e*p)
				}

				eptt = append(eptt, pt)
			}
			/*
				for i, e := range Ne {
					if i == j {
						continue
					}
				}
				for _, t := range tt {
					for _, a := range Np {
						if (a*e*test[1])%t == 0 {
							flag = 1
							mo.writeString(yes + "\n")
							break Loop
						}
					}
				}

				if flag == 0 {
					mo.writeString(no + "\n")
				}
			*/
		}
		mo.writeString("\n")
	}

}
