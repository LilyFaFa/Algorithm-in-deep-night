package main

import (
	"bufio"
	"fmt"
	//	"fmt"
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

type begin struct {
	end    int
	mother bool
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
		M := inputs[0]
		F := inputs[1]

		mid := make(map[int]begin)
		keys := []int{}

		totleM := 0
		totleF := 0

		avaliableM := []int{}
		avaliableF := []int{}

		for i := 0; i < M; i++ {
			times := mi.readIntSlice()
			B := times[0]
			E := times[1]
			b := begin{end: E, mother: true}
			mid[B] = b
			keys = append(keys, B)
		}

		for i := 0; i < F; i++ {
			times := mi.readIntSlice()
			B := times[0]
			E := times[1]
			b := begin{end: E, mother: false}
			mid[B] = b
			keys = append(keys, B)
		}

		sort.Ints(keys)

		fmt.Println(keys)
		fmt.Println(mid)

		if mid[keys[0]].mother == true {
			totleM = mid[keys[0]].end - keys[0]
		} else {
			totleF = mid[keys[0]].end - keys[0]
		}

		time := 0
		for i := 1; i < len(keys); i++ {

			if mid[keys[i]].mother == mid[keys[i-1]].mother {
				if mid[keys[i]].mother == true {
					avaliableM = append(avaliableM, keys[i]-mid[keys[i-1]].end)
					totleM = totleM + mid[keys[i]].end - mid[keys[i-1]].end
				} else {
					avaliableF = append(avaliableF, keys[i]-mid[keys[i-1]].end)
					totleF = totleF + mid[keys[i]].end - mid[keys[i-1]].end
				}
			} else if mid[keys[i]].mother == true {
				totleM = totleM + mid[keys[i]].end - keys[i]
				fmt.Println(mid[i])
				time++
			} else {
				totleF = totleF + mid[keys[i]].end - keys[i]
				fmt.Println(mid[i])
				time++
			}

			if i == len(keys)-1 {
				if mid[keys[i]].mother == mid[keys[0]].mother {
					if mid[keys[i]].mother == true {
						avaliableM = append(avaliableM, 1440-mid[keys[i]].end+keys[0])
						totleM = totleM + 1440 - mid[keys[i]].end + keys[0]
					} else {
						avaliableF = append(avaliableF, 1440-mid[keys[i]].end+keys[0])
						totleF = totleF + 1440 - mid[keys[i]].end + keys[0]
					}
				} else {
					fmt.Println(mid[i])
					time++
				}
			}

		}

		sort.Ints(avaliableM)
		sort.Ints(avaliableF)

		if totleM > 720 {
			over := totleM - 720
			for i := len(avaliableM) - 1; i > -1; i-- {
				over = over - avaliableM[i]
				fmt.Println(avaliableM[i])
				time = time + 2
				if over <= 0 {
					break
				}
			}
		}

		if totleF > 720 {
			over := totleF - 720
			for i := len(avaliableF) - 1; i > -1; i-- {
				over = over - avaliableF[i]
				time = time + 2
				if over <= 0 {
					break
				}
			}
		}

		if time == 0 {
			time = 2
		}
		c := strconv.Itoa(caseNo + 1)
		mo.writeString("Case #" + c + ":")
		mo.writeString(" ")
		mo.writeInt(time)
		mo.writeString("\n")

	}
}
