package day7

import (
	"log"
	"testing"
)

func TestGivens(t *testing.T) {

	phases := []int{4, 3, 2, 1, 0}
	program, e := StrToIntArr(sampleinput1)
	if e != nil {
		t.Error("BOOM")
	}
	chin, chout := Setup(program)
	result := Run(chin, chout, phases)
	if result != 43210 {
		log.Printf("the output is %d\n", result)
		t.Fail()
	}
}
func TestGiven2(t *testing.T) {

	phases := []int{0, 1, 2, 3, 4}
	program, e := StrToIntArr(sampleinput2)
	if e != nil {
		t.Error("BOOM")
	}
	chin, chout := Setup(program)
	result := Run(chin, chout, phases)
	if result != 54321 {
		log.Printf("the output is %d\n", result)
		t.Fail()
	}
}
func TestGiven3(t *testing.T) {

	phases := []int{1, 0, 4, 3, 2}
	program, e := StrToIntArr(sampleinput3)
	if e != nil {
		t.Error("BOOM")
	}
	chin, chout := Setup(program)
	result := Run(chin, chout, phases)
	if result != 65210 {
		log.Printf("the output is %d\n", result)
		t.Fail()
	}
}
func TestPartTwo(t *testing.T) {
	phases := []int{9, 8, 7, 6, 5}
	program, e := StrToIntArr(sampleinputparttwo)
	if e != nil {
		t.Error("BOOM")
	}
	result := RunPartTwo(program, phases)
	if result != 139629729 {
		log.Printf("part two got %d and expected 139629729\n", result)
		t.Error("test for part two errored out")
	}

}
func TestPartTwoB(t *testing.T) {
	phases := []int{9, 7, 8, 5, 6}
	program, e := StrToIntArr(sampleinputparttwob)
	if e != nil {
		t.Error("BOOM")
	}
	result := RunPartTwo(program, phases)
	if result != 18216 {
		log.Printf("part two got %d and expected 139629729\n", result)
		t.Error("test for part two errored out")
	}

}
func TestSplitopcode(t *testing.T) {
	res := SplitOpCode(993)
	log.Println(res)
	t.Error("BOOM")

}

var sampleinput1 string = "3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0"
var sampleinput2 string = "3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0"
var sampleinput3 string = "3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0"

var sampleinputparttwo string = "3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5"
var sampleinputparttwob string = "3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10"
