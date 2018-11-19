package main

import (
	"bufio"
	"fmt"
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

var (
	N int
	O int
	D int

	A int
	B int
	C int

	M int
	L int
)

type mid struct {
	o     int
	sweet int
}

func maxSweet(nums []int) (int, bool) {

	minNum := make([]int, len(nums))
	maxNum := make([]int, len(nums))
	minNum[len(nums)-1] = 0
	maxNum[len(nums)-1] = 0

	for i := len(nums) - 2; i >= 0; i-- {
		minNum[i] = 0
		maxNum[i] = 0

		if nums[i+1] > maxNum[i] {
			maxNum[i] = nums[i+1]
		}

		if nums[i+1] < minNum[i] {
			minNum[i] = nums[i+1]
		}

		if maxNum[i+1]+nums[i+1] > maxNum[i] {
			maxNum[i] = maxNum[i+1] + nums[i+1]
		}
		if minNum[i+1]+nums[i+1] < minNum[i] {
			minNum[i] = minNum[i+1] + nums[i+1]
		}
	}

	sweets := make([][]mid, len(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 1 && O > 0 {
			if nums[i]+minNum[i] <= D && D <= nums[i]+maxNum[i] {
				m := mid{
					o:     1,
					sweet: nums[i],
				}
				sweets[i] = append(sweets[i], m)
			}
		} else if nums[i]%2 == 0 {
			if nums[i]+minNum[i] <= D && D <= nums[i]+maxNum[i] {
				m := mid{
					o:     0,
					sweet: nums[i],
				}
				sweets[i] = append(sweets[i], m)
			}
		}
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < len(sweets[i-1]); j++ {

			if nums[i]%2 == 1 {

				if sweets[i-1][j].o+1 <= O {
					if nums[i]+sweets[i-1][j].sweet+minNum[i] <= D && D <= nums[i]+sweets[i-1][j].sweet+maxNum[i] {
						m := mid{
							o:     sweets[i-1][j].o + 1,
							sweet: nums[i] + sweets[i-1][j].sweet,
						}
						sweets[i] = append(sweets[i], m)
					}
				}

			} else {
				if nums[i]+sweets[i-1][j].sweet+minNum[i] <= D && D <= nums[i]+sweets[i-1][j].sweet+maxNum[i] {
					m := mid{
						o:     sweets[i-1][j].o,
						sweet: nums[i] + sweets[i-1][j].sweet,
					}
					sweets[i] = append(sweets[i], m)
				}
			}
		}
	}

	possiable := false
	var max int
	for i := 0; i < len(sweets); i++ {
		for j := 0; j < len(sweets[i]); j++ {
			ss := sweets[i][j]
			if ss.sweet <= D && !possiable {
				possiable = true
				max = ss.sweet
			} else if ss.sweet <= D && possiable && ss.sweet > max {
				max = ss.sweet
			}

		}
	}
	return max, possiable
}

func main() {
	fi, _ := os.Open("A-small-attempt0.in")
	mi := MyInput{rdr: fi}
	fo, _ := os.OpenFile("A-small-attempt0.out", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	mo := MyOutPut{wrr: fo}
	//mi := MyInput{rdr: os.Stdin}
	//mo := MyOutPut{wrr: os.Stdout}

	count := mi.readInt()
	for caseNo := 0; caseNo < count; caseNo++ {
		inputs := mi.readIntSlice()
		N = inputs[0]
		O = inputs[1]
		D = inputs[2]

		inputs = mi.readIntSlice()
		X1 := inputs[0]
		X2 := inputs[1]

		A = inputs[2]
		B = inputs[3]
		C = inputs[4]
		M = inputs[5]
		L = inputs[6]
		var nums []int
		nums = append(nums, inputs[0]+L)
		nums = append(nums, inputs[1]+L)

		for i := 2; i < N; i++ {
			num := ((A%M)*(X1%M)%M + ((B%M)*(X2%M))%M + C%M) % M
			nums = append(nums, num+L)
			if i%2 == 0 {
				X1 = num
			} else {
				X2 = num
			}
		}

		fmt.Print("start")
		fmt.Print(nums)

		result, poss := maxSweet(nums)
		if poss {
			c := strconv.Itoa(caseNo + 1)
			mo.writeString("Case #" + c + ":")
			mo.writeString(" ")
			mo.writeInt(result)
			mo.writeString("\n")
		} else {
			c := strconv.Itoa(caseNo + 1)
			mo.writeString("Case #" + c + ":")
			mo.writeString(" ")
			mo.writeString("IMPOSSIBLE")
			mo.writeString("\n")
		}

	}
}
