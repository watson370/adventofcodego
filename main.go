package main

import (
	"fmt"
	"log"

	"github.com/watson370/adventofcodego/day1"
	"github.com/watson370/adventofcodego/day2"
	"github.com/watson370/adventofcodego/day3"
	"github.com/watson370/adventofcodego/day4"
)

func main() {
	// mass := float64(12)
	// res := day1.Calc(mass)
	// fmt.Printf("The mass for day 1 is %f and the calculated fuel is %d\n", mass, res)
	var res int64 = 0
	var totalMass int64 = 0
	for _, v := range day1inputs {
		totalMass += int64(v)
		res += int64(day1.Calc(float64(v)))
	}
	fmt.Printf("The calculated fuel for all of the day1 modules is %d for a total mass of %d\n", res, totalMass)
	var reswadj int64 = 0
	var totalMasswadj int64 = 0
	for _, v := range day1inputs {
		totalMasswadj += int64(v)
		reswadj += int64(day1.CalcWithFuelAdjustment(float64(v)))
	}
	fmt.Printf("The calculated fuel for all of the day1 modules including adjustment for fuel mass is %d for a total mass of %d\n", reswadj, totalMasswadj)

	//day 2
	//before running the program, replace position 1 with the value 12 and replace position 2 with the value 2
	adjinput, _ := day2.StrToIntArr(day2input)
	adjinput[1] = 12
	adjinput[2] = 2
	adjinputstr := day2.ArrToStr(adjinput)
	fmt.Printf("\n\n\n")
	fmt.Println(adjinputstr)
	day2res, e := day2.ProcessCommaSeparated(adjinputstr)
	if e != nil {
		fmt.Println("o no we have an error " + e.Error())
	}
	fmt.Println(day2res)

	resultarr := []int{}
	needle := 19690720
	//address 1 is noun, address 2 is verb, can be values 0-99 inclusive
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			//reset
			adjinput, _ := day2.StrToIntArr(day2input)
			adjinput[1] = noun
			adjinput[2] = verb
			resultarr, e = day2.Process(adjinput)
			if resultarr[0] == needle {
				fmt.Printf("A value of Noun: %d Verb: %d produces a value of %d needle value %d \n", noun, verb, resultarr[0], needle)
				fmt.Printf("100 * %d + %d = %d\n", noun, verb, (100*noun)+verb)
				break
			}

		}
	}

	//day3
	wp1coords := day3.BuildCoordList(day3input1)
	wp2coords := day3.BuildCoordList(day3input2)
	// log.Println(wp1coords)
	// log.Println("\n\n ---------------")
	// log.Println(wp2coords)
	intersections := day3.FindIntersections(wp1coords, wp2coords)
	min, minSteps := day3.MinManhattenDist(intersections)
	log.Println(intersections)
	log.Printf("The minimum manhatten dist is %d and the min steps is %d", min, minSteps)

	//day 4
	totalPassed := 0
	totalFailed := 0
	for subject := day4min; subject <= day4max; subject++ {
		if day4.Test(subject, day4min, day4max, 1) {
			totalPassed++
		} else {
			totalFailed++
		}
	}
	log.Printf("part 1 %d passed, %d failed\n", totalPassed, totalFailed)
	totalPassed = 0
	totalFailed = 0
	for subject := day4min; subject <= day4max; subject++ {
		if day4.Test(subject, day4min, day4max, 2) {
			totalPassed++
		} else {
			totalFailed++
		}
	}
	log.Printf("Part 2 %d passed, %d failed\n", totalPassed, totalFailed)

}

var day1inputs []int = []int{
	120150,
	105328,
	70481,
	86674,
	112434,
	94883,
	147500,
	146007,
	103982,
	65758,
	132357,
	60885,
	97516,
	96977,
	129085,
	80330,
	124081,
	102501,
	102505,
	70029,
	54155,
	69253,
	60120,
	53192,
	89470,
	137125,
	136296,
	104546,
	92859,
	74937,
	135044,
	66238,
	126678,
	86364,
	138559,
	82393,
	96947,
	107749,
	115332,
	117563,
	95431,
	99640,
	107667,
	120427,
	108389,
	51567,
	57493,
	68518,
	114565,
	107248,
	50627,
	122517,
	129687,
	118989,
	52459,
	83726,
	106765,
	75872,
	147111,
	78822,
	65058,
	142460,
	122496,
	148942,
	72753,
	141599,
	105711,
	141860,
	134066,
	116716,
	107455,
	56673,
	148238,
	92318,
	86652,
	69312,
	91352,
	94528,
	73441,
	137814,
	80247,
	101115,
	61773,
	100951,
	77189,
	119083,
	93841,
	109090,
	83370,
	70230,
	144812,
	67647,
	105358,
	135780,
	85225,
	100697,
	100998,
	131151,
	52826,
	114084,
}
var day2input = "1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,13,1,19,1,9,19,23,2,23,13,27,1,27,9,31,2,31,6,35,1,5,35,39,1,10,39,43,2,43,6,47,1,10,47,51,2,6,51,55,1,5,55,59,1,59,9,63,1,13,63,67,2,6,67,71,1,5,71,75,2,6,75,79,2,79,6,83,1,13,83,87,1,9,87,91,1,9,91,95,1,5,95,99,1,5,99,103,2,13,103,107,1,6,107,111,1,9,111,115,2,6,115,119,1,13,119,123,1,123,6,127,1,127,5,131,2,10,131,135,2,135,10,139,1,13,139,143,1,10,143,147,1,2,147,151,1,6,151,0,99,2,14,0,0"

var day3input1 = "R1003,U756,L776,U308,R718,D577,R902,D776,R760,U638,R289,D70,L885,U161,R807,D842,R175,D955,R643,U380,R329,U573,L944,D2,L807,D886,L549,U592,R152,D884,L761,D915,L726,D677,L417,D651,L108,D377,L699,D938,R555,D222,L689,D196,L454,U309,L470,D234,R198,U689,L996,U117,R208,D310,R572,D562,L207,U244,L769,U186,R153,D756,R97,D625,R686,U244,R348,U586,L385,D466,R483,U718,L892,D39,R692,U756,L724,U148,R70,U224,L837,D370,L309,U235,R382,D579,R404,D146,R6,U584,L840,D863,R942,U646,R146,D618,L12,U210,R126,U163,R931,D661,L710,D883,L686,D688,L148,D19,R703,U530,R889,U186,R779,D503,R417,U272,R541,U21,L562,D10,L349,U998,R69,D65,R560,D585,L949,D372,L110,D865,R212,U56,L936,U957,L88,U612,R927,U642,R416,U348,L541,D416,L808,D759,R449,D6,L517,D4,R494,D143,L536,U341,R394,U179,L22,D680,L138,U249,L285,U879,L717,U756,L313,U222,R823,D208,L134,U984,R282,U635,R207,D63,L416,U511,L179,D582,L651,U932,R646,U378,R263,U138,L920,U523,L859,D556,L277,D518,R489,U561,L457,D297,R72,U920,L583,U23,L395,D844,R776,D552,L55,D500,R111,U409,R685,D427,R275,U739,R181,U333,L215,U808,R341,D537,R336,U230,R247,U748,R846,U404,R850,D493,R891,U176,L744,U585,L987,D849,R271,D848,L555,U801,R316,U753,L390,U97,L128,U45,R706,U35,L928,U913,R537,D512,R152,D410,R76,D209,R183,U941,R289,U632,L923,D190,R488,D934,R442,D303,R178,D250,R204,U247,R707,U77,R428,D701,R386,U110,R641,U925,R703,D387,L946,U415,R461,D123,L214,U236,L959,U517,R957,D524,R812,D668,R369,U340,L606,D503,R755,U390,R142,D921,L976,D36,L965,D450,L722,D224,L303,U705,L584"
var day3input2 = "L993,U810,L931,D139,R114,D77,L75,U715,R540,D994,L866,U461,R340,D179,R314,D423,R629,D8,L692,U446,L88,D132,L128,U934,L465,D58,L696,D883,L955,D565,R424,U286,R403,U57,L627,D930,R887,D941,L306,D951,R918,U587,R939,U821,L65,D18,L987,D707,L360,D54,L932,U366,R625,U609,R173,D637,R661,U888,L68,U962,R270,U369,R780,U845,L813,U481,R66,D182,R420,U605,R880,D276,L6,D529,R883,U189,R380,D472,R30,U35,L510,D844,L146,U875,R152,U545,R274,U920,R432,U814,R583,D559,L820,U135,L353,U975,L103,U615,R401,U692,L676,D781,R551,D985,L317,U836,R115,D216,L967,U286,R681,U144,L354,U678,L893,D487,R664,D185,R787,D909,L582,D283,L519,D893,L56,U768,L345,D992,L248,U439,R573,D98,L390,D43,L470,D435,R176,U468,R688,U388,L377,U800,R187,U641,L268,U857,L716,D179,R212,U196,L342,U731,R261,D92,R183,D623,L589,D215,L966,U878,L784,U740,R823,D99,L167,D992,R414,U22,L27,U390,R286,D744,L360,U554,L756,U715,R939,D806,R279,U292,L960,U633,L428,U949,R90,D321,R749,U395,L392,U348,L33,D757,R289,D367,L562,D668,L79,D193,L991,D705,L562,U25,R146,D34,R325,U203,R403,D714,R607,U72,L444,D76,R267,U924,R289,U962,L159,U726,L57,D540,R299,U343,R936,U90,L311,U243,L415,D426,L936,D570,L539,D731,R367,D374,L56,D251,L265,U65,L14,D882,L956,U88,R688,D34,R866,U777,R342,D270,L344,D953,L438,D855,L587,U320,L953,D945,L473,U559,L487,D602,R255,U871,L854,U45,R705,D247,R955,U885,R657,D664,L360,D764,L549,D676,R85,U189,L951,D922,R511,D429,R37,U11,R821,U984,R825,U874,R753,D524,L537,U618,L919,D597,L364,D231,L258,U818,R406,D208,R214,U530,R261"

// 359282-820401
var day4min = 359282
var day4max = 820401
