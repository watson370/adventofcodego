package day1

import (
	"testing"
)

func TestCalc(t *testing.T) {
	goodvals := make(map[float64]int)
	goodvals[float64(12)] = 2
	goodvals[float64(14)] = 2
	goodvals[float64(1969)] = 654
	goodvals[float64(100756)] = 33583

	for key, value := range goodvals {
		res := Calc(key)
		if res != value {
			t.Errorf("there was an err, a mass of %f needs fuel %d, but calculated %d", key, value, res)
		} else {
			t.Logf("a mass of %f needs fuel %d, and we calculated %d", key, value, res)

		}
	}
}

func TestCalcWithAdj(t *testing.T) {
	goodvals := make(map[float64]int64)
	goodvals[float64(14)] = 2
	goodvals[float64(1969)] = 966
	goodvals[float64(100756)] = 50346

	for key, value := range goodvals {
		res := CalcWithFuelAdjustment(key)
		if res != value {
			t.Errorf("there was an err, a mass of %f needs fuel %d, but calculated %d", key, value, res)
		} else {
			t.Logf("a mass of %f needs fuel %d, and we calculated %d", key, value, res)

		}
	}
}
