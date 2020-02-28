package day2

import (
	"fmt"
	"testing"
)

func TestIntcode(t *testing.T) {
	// arrval := strings.Split(stringval, ",")
	goodvals := make(map[string]string)
	goodvals["1,9,10,3,2,3,11,0,99,30,40,50"] = "3500,9,10,70,2,3,11,0,99,30,40,50"
	goodvals["1,0,0,0,99"] = "2,0,0,0,99"
	goodvals["2,3,0,3,99"] = "2,3,0,6,99"
	goodvals["2,4,4,5,99,0"] = "2,4,4,5,99,9801"
	goodvals["1,1,1,4,99,5,6,0,99"] = "30,1,1,4,2,5,6,0,99"

	for key, value := range goodvals {
		res, e := ProcessCommaSeparated(key)
		if e != nil {
			t.Error("failed because ", e)
		}
		if res != value {
			t.Errorf("there was an err, an incode of %s needs to be calculated to %s, but calculated %s", key, value, res)
		} else {
			t.Logf("an intocde of %s needs to be processed to %s, and we calculated %s", key, value, res)

		}
	}
}

func TestConversion(t *testing.T) {
	res, e := StrToIntArr("0, 1, 2, 3, 4, 5, 6, 7, 8, 9")
	if e != nil {
		t.Error("failed with error " + e.Error())
	}
	for i := 0; i < 10; i++ {
		if i != res[i] {
			t.Error("Batman, We have a problem ")
		}
	}

	myslice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	mystring := ArrToStr(myslice)
	expected := "0,1,2,3,4,5,6,7,8,9"
	if mystring != expected {
		t.Error("string " + mystring + "does not equal " + expected)
	}
}

func TestProcess(t *testing.T) {
	input := []int{1, 0, 0, 0, 99}
	expected := []int{2, 0, 0, 0, 99}
	res, e := Process(input)

	if e != nil {
		t.Error("failed with the error of " + e.Error())
	}
	fmt.Println(res)
	t.Log(res)
	t.Log(e)
	for i, v := range expected {
		if res[i] != v {
			t.Error(res)
			t.Error(expected)
		}
	}
}
