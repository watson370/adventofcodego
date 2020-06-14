package main

import "testing"

//check mark for success, x for fail
const (
	succeed = "\u2713"
	failed  = "\u2717"
)

func TestCopyProgram(t *testing.T) {

	tests := []struct {
		left     int
		right    int
		expected int
	}{
		{5, 25, 5},
		{6, 9, 3},
		{17, 11, 1},
	}
	t.Log("Given the need to test calculating the greatest common denominator of two integers.")
	{
		for i, tt := range tests {
			t.Logf("\tTest: %d\tWhen checking %d, %d  for gcd of  %d", i, tt.left, tt.right, tt.expected)
			{
				res := greatestCommonDenominator(tt.left, tt.right)
				if res == tt.expected {
					t.Logf("\t%s\tShould get gcd of %d", succeed, tt.expected)
				} else {
					t.Errorf("\t%s\tShould get a gcd of %d but got : %d", failed, tt.expected, res)
				}
			}
		}
	}
	left := 5
	right := 25
	expected := 5
	res := greatestCommonDenominator(left, right)
	if res != expected {
		t.Errorf("expected %d got %d\n", expected, res)
	}
}

func TestLeastCommonMultiple(t *testing.T) {
	tests := []struct {
		vals     []int
		expected int
	}{
		{[]int{5, 25}, 25},
		{[]int{6, 9}, 18},
		{[]int{11, 11}, 11},
	}
	t.Log("Given the need to test calculating the least common multiple of two integers.")
	{
		for i, tt := range tests {
			t.Logf("\tTest: %d\tWhen checking %v,  for lcm of  %d", i, tt.vals, tt.expected)
			{
				res := leastCommonMultiple(tt.vals)
				if res == tt.expected {
					t.Logf("\t%s\tShould get lcm of %d", succeed, tt.expected)
				} else {
					t.Errorf("\t%s\tShould get a lcm of %d but got : %d", failed, tt.expected, res)
				}
			}
		}
	}

}
