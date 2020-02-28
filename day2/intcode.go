package day2

import (
	"errors"
	"strconv"
	"strings"
)

//ProcessCommaSeparated stuff
func ProcessCommaSeparated(s string) (string, error) {
	//initialize "computer memory"
	inputintarr, e := StrToIntArr(s)
	if e != nil {
		return "", e
	}

	output, e := Process(inputintarr)
	if e != nil {
		//kaboom
		return "", errors.New("Process failed with error " + e.Error())
	}
	return ArrToStr(output), nil
}

//Process   1, 2, 99, or error
func Process(ints []int) ([]int, error) {
	set := make(map[int]string)
	set[1] = "addition"
	set[2] = "multiplication"
	set[99] = "terminate"
	instructionpointer := 0 //pointer to current memory address
	instruction := ""       //instruction
	ok := true
	firstoperand := 0
	secondoperand := 0
	result := 0
	resultpos := 0
	params := make([]int, 0, 10)

	for {
		instruction, ok = set[ints[instructionpointer]]
		if !ok {
			return ints, errors.New("unexpected value " + strconv.Itoa(ints[instructionpointer]) + " at position " + strconv.Itoa(instructionpointer))
		}

		//get a slice of the "parameters"
		//zero out
		params = params[:0]
		switch instruction {
		case "addition", "multiplication":
			if len(ints) > instructionpointer+3+1 {
				for l := 1; l < 4; l++ {
					params = append(params, ints[instructionpointer+l])
				}
			}
		}
		//perform instruction
		switch instruction {
		case "addition":
			if len(params) != 3 {
				return []int{}, errors.New("intruction code 1 for addition requires 3 params, but received " + strconv.Itoa(len(params)))
			}
			firstoperand = ints[params[0]]
			secondoperand = ints[params[1]]
			resultpos = params[2]
			result = firstoperand + secondoperand
			// log.Printf("current position %d first operand %d second operand %d  result position %d", instructionpointer, firstoperand, secondoperand, resultpos)

		case "multiplication":
			if len(params) != 3 {
				return []int{}, errors.New("intruction code 2 for multiplicaiton requires 3 params, but received " + strconv.Itoa(len(params)))
			}
			firstoperand = ints[params[0]]
			secondoperand = ints[params[1]]
			resultpos = params[2]
			result = firstoperand * secondoperand
		case "terminate":
			return ints, nil
		default:
			return ints, errors.New("unexpected instruction value " + instruction)
		}

		ints[resultpos] = result
		instructionpointer += 4
		// log.Printf("next instruction at %d", instructionpointer)

	}
}

//ArrToStr  stuff
func ArrToStr(vals []int) string {
	builder := strings.Builder{}
	for i, v := range vals {
		if i != 0 {
			builder.WriteString(",")
		}
		builder.WriteString(strconv.Itoa(v))

	}
	return builder.String()
}

//StrToIntArr takes in a sting with comma separated vals and returns an []int
func StrToIntArr(s string) ([]int, error) {
	input := strings.Split(s, ",")
	// fmt.Println(input)
	inputintarr := make([]int, len(input))
	for i, s := range input {
		tmp, e := strconv.Atoi(strings.Join(strings.Fields(s), ""))
		if e != nil {
			//kaboom
			return []int{}, errors.New("what happened was " + e.Error())
		}
		inputintarr[i] = tmp
	}
	return inputintarr, nil
}
