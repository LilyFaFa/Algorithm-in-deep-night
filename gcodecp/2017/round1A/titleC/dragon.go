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
	imp := "IMPOSSIBLE"
	fi, _ := os.Open("test")

	mi := MyInput{rdr: fi}
	fo, _ := os.OpenFile("test2", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	mo := MyOutPut{wrr: fo}
	//	mi := MyInput{rdr: os.Stdin}
	//mo := MyOutPut{wrr: os.Stdout}
	count := mi.readInt()

	for caseNo := 0; caseNo < count; caseNo++ {
		result := -1
		inputs := mi.readIntSlice()
		Battle(inputs[0], inputs[1], inputs[2], inputs[3], inputs[4], inputs[5], 0, &result, inputs[0])

		c := strconv.Itoa(caseNo + 1)
		mo.writeString("Case #" + c + ":")
		mo.writeString(" ")
		if result == -1 {
			mo.writeString(imp)
		} else {
			mo.writeInt(result)
		}

		mo.writeString("\n")
	}
}

func Battle(Hd int, Ad int, Hk int, Ak int, B int, D int, time int, result *int, Hf int) {
	//放弃
	if *result != -1 && time >= *result {
		return
	}

	//失败了
	if Hd <= 0 {
		return
	}

	//可以击败对方，攻击
	if Ad >= Hk {
		if *result < 0 {
			*result = time + 1
		} else if *result > time+1 {
			*result = time + 1
		}
		return
	}

	if Hd <= Ak-D {
		//必须治愈
		Battle(Hf-Ak, Ad, Hk, Ak, B, D, time+1, result, Hf)
		return
	}

	if Hd > Ak-D && Hd <= Ak {
		//消弱对方或者治愈
		Battle(Hf-Ak, Ad, Hk, Ak, B, D, time+1, result, Hf)
		if Ak >= D {
			Battle(Hd-Ak+D, Ad, Hk, Ak-D, B, D, time+1, result, Hf)
		} else if Ak > 0 {
			Battle(Hd, Ad, Hk, 0, B, D, time+1, result, Hf)
		}
		return
	}

	//攻击对方
	if Ad > 0 {
		Battle(Hd-Ak, Ad, Hk-Ad, Ak, B, D, time+1, result, Hf)
	}

	//提高自己的攻击力
	if B > 0 {
		Battle(Hd-Ak, Ad+B, Hk, Ak, B, D, time+1, result, Hf)
	}

	//降低对方的攻击力
	if Ak >= D {
		Battle(Hd-Ak+D, Ad, Hk, Ak-D, B, D, time+1, result, Hf)
	} else if Ak > 0 {
		Battle(Hd, Ad, Hk, 0, B, D, time+1, result, Hf)
	}

}
