package day5

import (
	"testing"
)

func TestSplitOpCode(t *testing.T) {
	test := 1002
	expected := []int{2, 0, 1, 0}
	res := SplitOpCode(test)
	for i, v := range expected {
		if v != res[i] {
			t.Log(expected)
			t.Log(res)
			t.Errorf("bad result")
		}
	}

}
func TestGetVal(t *testing.T) {
	ints := []int{1002, 4, 3, 4, 33}
	opcodeandparammodes := []int{2, 0, 1, 0}

	//I need to know if this is the first, second, or third parameter, and the position of the first parameter in the ints []int
	res := GetVal(ints, opcodeandparammodes, 0, 1)
	if res != 33 {
		t.Error("did not get 33")
	}
	res = GetVal(ints, opcodeandparammodes, 0, 2)
	if res != 3 {
		t.Error("did not get 3")
	}
	t.Fail()

}
