package main

import (
	"bufio"
	//"fmt"
	//	"fmt"
	//"errors"
	//"fmt"
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

type exist struct {
	partners []int
	time     int
}

func main() {

	fi, err := os.Open("A-large.in")
	if err != nil {
		panic(err)
	}

	//fi, _ := os.Open("small.in")
	mi := MyInput{rdr: fi}
	fo, _ := os.OpenFile("A-large.out", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	//fo, _ := os.OpenFile("small.out", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	mo := MyOutPut{wrr: fo}
	///mi := MyInput{rdr: os.Stdin}
	//mo := MyOutPut{wrr: os.Stdout}
	count := mi.readInt()
	for caseNo := 0; caseNo < count; caseNo++ {
		//		fmt.Println("hello")
		//number := 0
		result := make(map[int]int)
		tmp := make(map[int]int)

		countN := mi.readInt()
		m := make(map[int]*exist)

		for i := 0; i < countN; i++ {
			xy := mi.readIntSlice()
			g, ok := m[xy[0]]
			if ok {
				g.partners = append(g.partners, xy[1])
				g.time = g.time + 1
			} else {
				partner := []int{xy[1]}
				e := exist{
					partners: partner,
					time:     1,
				}
				m[xy[0]] = &e
			}

			g, ok = m[xy[1]]
			if ok {
				g.partners = append(g.partners, xy[0])
				g.time = g.time + 1
			} else {
				partner := []int{xy[0]}
				e := exist{
					partners: partner,
					time:     1,
				}
				m[xy[1]] = &e
			}
		}

		flag := 0
		for flag == 0 {
			flag = 1
			//number++
			for key, value := range m {
				if value.time == 1 {
					partner := value.partners[0]
					delete(m, key)

					tmp[key] = partner
					for i, p := range m[partner].partners {
						if p == key {
							m[partner].partners = append(m[partner].partners[:i], m[partner].partners[i+1:]...)
							m[partner].time = m[partner].time - 1
						}
					}
					flag = 0
				}
			}

		}

		circle := make(map[int]bool)
		for i := 0; i <= countN; i++ {
			_, ok := tmp[i]
			if !ok {
				circle[i] = true
			}
		}

		number := 0
		for len(tmp) != 0 {
			number++
			tt := []int{}
			for k, v := range tmp {
				_, ok := circle[v]
				if ok {
					tt = append(tt, k)
					result[k] = number
					//circle[k] = true
					delete(tmp, k)
				}
			}
			for _, t := range tt {
				circle[t] = true
			}
		}

		c := strconv.Itoa(caseNo + 1)
		mo.writeString("Case #" + c + ":")
		for i := 1; i <= countN; i++ {
			mo.writeString(" ")
			g, ok := result[i]
			if ok {
				mo.writeInt(g)
			} else {
				mo.writeInt(0)
			}

		}
		mo.writeString("\n")
	}
}
