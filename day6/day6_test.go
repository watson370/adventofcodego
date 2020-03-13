package day6

import (
	"testing"
)

func TestSplitOpCode(t *testing.T) {
	input := `COM)B
	B)C
	C)D
	D)E
	E)F
	B)G
	G)H
	D)I
	E)J
	J)K
	K)L`
	direct, indirect := Run(input)
	if direct+indirect != 42 {
		t.Error("number of total orbits should be 42")
	}

}
