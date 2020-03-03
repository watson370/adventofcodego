package day4

import "testing"

// 111111 meets these criteria (double 11, never decreases).
// 223450 does not meet these criteria (decreasing pair of digits 50).
// 123789 does not meet these criteria (no double).
func TestIntcode(t *testing.T) {
	res := Test(111111, 0, 999999, 1)
	if !res {
		t.Error("111111 returned false, but should be true")
	}
	res = Test(223450, 0, 999999, 1)
	if res {
		t.Error("223450")
	}
	res = Test(123789, 0, 999999, 1)
	if res {
		t.Error("123789 passed, but should fail")
	}
	// 	112233 meets these criteria because the digits never decrease and all repeated digits are exactly two digits long.
	// 123444 no longer meets the criteria (the repeated 44 is part of a larger group of 444).
	// 111122 meets the criteria (even though 1 is repeated more than twice, it still contains a double 22).
	res = Test(112233, 0, 999999, 2)
	if !res {
		t.Error("112233 failed, but should pass")
	}
	res = Test(123444, 0, 999999, 2)
	if res {
		t.Error("123444 passed, but should fail")
	}
	res = Test(111122, 0, 999999, 2)
	if !res {
		t.Error("111122 should pass, but it failed")
	}
}
