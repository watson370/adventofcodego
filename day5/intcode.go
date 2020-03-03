package day5

import (
	"bufio"
	"errors"
	"fmt"
	"os"
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
	set[3] = "input"
	set[4] = "output"
	set[5] = "jump-if-true"
	set[6] = "jump-if-false"
	set[7] = "less-than"
	set[8] = "equal"
	set[99] = "terminate"
	instructionpointer := 0 //pointer to current memory address
	instruction := ""       //instruction
	ok := true
	firstoperand := 0
	secondoperand := 0
	result := 0
	resultpos := 0
	var opcodeandparammodes []int

	for {
		opcodeandparammodes = SplitOpCode(ints[instructionpointer])
		instruction, ok = set[opcodeandparammodes[0]]
		if !ok {
			return ints, errors.New("unexpected value " + strconv.Itoa(ints[instructionpointer]) + " at position " + strconv.Itoa(instructionpointer))
		}

		//perform instruction
		switch instruction {
		case "addition":
			firstoperand = GetVal(ints, opcodeandparammodes, instructionpointer, 1)  //ints[params[0]]
			secondoperand = GetVal(ints, opcodeandparammodes, instructionpointer, 2) //ints[params[1]]
			resultpos = ints[instructionpointer+3]                                   // params[2]
			result = firstoperand + secondoperand
			ints[resultpos] = result
			instructionpointer += 4
		case "multiplication":
			firstoperand = GetVal(ints, opcodeandparammodes, instructionpointer, 1)  //ints[params[0]]
			secondoperand = GetVal(ints, opcodeandparammodes, instructionpointer, 2) //ints[params[1]]
			resultpos = ints[instructionpointer+3]                                   // params[2]
			result = firstoperand * secondoperand
			ints[resultpos] = result
			instructionpointer += 4
		case "input":
			resultpos = ints[instructionpointer+1]
			result = GetInput()
			ints[resultpos] = result
			instructionpointer += 2
		case "output":
			result = GetVal(ints, opcodeandparammodes, instructionpointer, 1)
			OutputResult(result)
			instructionpointer += 2
		case "jump-if-true":
			// Opcode 5 is jump-if-true: if the first parameter is non-zero, it sets the instruction pointer to the value from the second parameter. Otherwise, it does nothing.
			firstoperand = GetVal(ints, opcodeandparammodes, instructionpointer, 1)
			secondoperand = GetVal(ints, opcodeandparammodes, instructionpointer, 2)
			if firstoperand != 0 {
				instructionpointer = secondoperand
			} else {
				instructionpointer += 3
			}
		case "jump-if-false":
			firstoperand = GetVal(ints, opcodeandparammodes, instructionpointer, 1)
			secondoperand = GetVal(ints, opcodeandparammodes, instructionpointer, 2)
			if firstoperand == 0 {
				instructionpointer = secondoperand
			} else {
				instructionpointer += 3
			}
		case "less-than":
			firstoperand = GetVal(ints, opcodeandparammodes, instructionpointer, 1)
			secondoperand = GetVal(ints, opcodeandparammodes, instructionpointer, 2)
			resultpos = ints[instructionpointer+3]
			if firstoperand < secondoperand {
				ints[resultpos] = 1
			} else {
				ints[resultpos] = 0
			}
			instructionpointer += 4
		case "equal":
			firstoperand = GetVal(ints, opcodeandparammodes, instructionpointer, 1)
			secondoperand = GetVal(ints, opcodeandparammodes, instructionpointer, 2)
			resultpos = ints[instructionpointer+3]
			if firstoperand == secondoperand {
				ints[resultpos] = 1
			} else {
				ints[resultpos] = 0
			}
			instructionpointer += 4
		case "terminate":
			return ints, nil
		default:
			return ints, errors.New("unexpected instruction value " + instruction)
		}
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

//SplitOpCode takes in the integer opcode and return an int slice with the opcode in position 0, and the parameter modes in the following positions
func SplitOpCode(opcode int) []int {
	strval := strconv.Itoa(opcode)
	toreturn := make([]int, 4)
	if len(strval) == 1 {
		opcode, _ = strconv.Atoi(strval)
	} else {
		opcode, _ = strconv.Atoi(strval[len(strval)-2:])
	}
	toreturn[0] = opcode
	//i index to write to r index to read from
	//i starts at 1, 0 is the opcode, r -1 for zero based, -2 for the opcode
	for i, r := 1, len(strval)-2-1; r >= 0; i, r = i+1, r-1 {
		intval, _ := strconv.Atoi(string(strval[r]))
		toreturn[i] = intval
	}
	return toreturn
}

//GetVal takes in the []int of the program, the []int that holds the parameter input modes, the index of the current instruction, and the position requested, then returns the correct value based on the mode
//the position requested is 1, 2, or 3, not index based, but the number of the param, but turns out to be the index in the opcodeandparammodes
func GetVal(ints []int, opcodeandparammodes []int, instructionpointer int, position int) int {
	mode := opcodeandparammodes[position]
	if mode == 0 { //position mode
		return ints[ints[instructionpointer+position]]
	} else if mode == 1 { //immediate mode
		return ints[instructionpointer+position]
	}
	//todo decide how to handle unsupported modes
	return -1
}

//GetInput asks the user for the ID of the system to test by running an input instruction
func GetInput() int {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter the ID of the system to test: ")
		// Scans a line from Stdin(Console)
		scanner.Scan()
		// Holds the string that scanned
		text := scanner.Text()
		res, e := strconv.Atoi(text)
		if e != nil {
			fmt.Printf("The system id is an int, for example 1 for the air conditioning system. %s is not valid input", text)
		} else {
			//close scanner ?
			return res
		}

	}
}

// OutputResult oututs the int to standard out
func OutputResult(result int) {
	outputstring := strconv.Itoa(result)
	w := bufio.NewWriter(os.Stdout)
	fmt.Fprint(w, outputstring)
	fmt.Fprint(w, "\n")
	w.Flush()
}
