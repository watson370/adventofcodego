package intcode_test

import (
	"testing"

	"github.com/watson370/adventofcodego/day11/intcode"
)

func TestSplitOpCode(t *testing.T) {
	test := int64(1002)
	expected := []int64{2, 0, 1, 0}
	res := intcode.SplitOpCode(test, make([]int64, 4))
	for i, v := range expected {
		if v != res[i] {
			t.Log(expected)
			t.Log(res)
			t.Errorf("bad result")
		}
	}

}
func TestSplitOpCode2(t *testing.T) {
	test := int64(21002)
	expected := []int64{2, 0, 1, 2}
	res := intcode.SplitOpCode(test, make([]int64, 4))
	for i, v := range expected {
		if v != res[i] {
			t.Log(expected)
			t.Log(res)
			t.Errorf("bad result")
		}
	}

}
func TestGetVal(t *testing.T) {
	ints := []int64{1002, 4, 3, 4, 33}
	opcodeandparammodes := []int64{2, 0, 1, 0}

	//I need to know if this is the first, second, or third parameter, and the position of the first parameter in the ints []int
	res := intcode.GetVal(ints, opcodeandparammodes, 0, 1, 0)
	if res != 33 {
		t.Error("did not get 33")
	}
	res = intcode.GetVal(ints, opcodeandparammodes, 0, 2, 0)
	if res != 3 {
		t.Error("did not get 3")
	}
}

func TestCopyProgram(t *testing.T) {
	//when given output
	t.Log("Given the need to test an intcode program that outputs itself as output.")
	ints := []int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}
	cp := make([]int64, len(ints), len(ints))
	copy(cp, ints)
	hp := intcode.NewHullPainter(cp)
	e := hp.Process()
	if e != nil {
		t.Log(e)
	}
	output := hp.GetOutput()
	for i := range ints {
		if ints[i] != output[i] {
			t.Errorf("Output not the same as the intcode Program, %d, should be %d\n", ints[i], output[i])
		}
	}

}
