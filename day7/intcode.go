package day7

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//ProcessCommaSeparated stuff
func ProcessCommaSeparated(s string, inputs []int) (int, error) {
	//initialize "computer memory"
	inputintarr, e := StrToIntArr(s)
	if e != nil {
		return 0, e
	}

	output, e := Process(inputintarr, inputs)
	if e != nil {
		//kaboom
		return 0, errors.New("Process failed with error " + e.Error())
	}
	return output, nil
}

//Process   1-8, 99
func Process(ints []int, inputs []int) (int, error) {
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
			return 0, errors.New("unexpected value " + strconv.Itoa(ints[instructionpointer]) + " at position " + strconv.Itoa(instructionpointer))
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
			result = inputs[0]
			inputs = inputs[1:]

			ints[resultpos] = result
			instructionpointer += 2
		case "output":
			result = GetVal(ints, opcodeandparammodes, instructionpointer, 1)
			return result, nil //in this particular case, we will never keep coputing after an output
			// OutputResult(result)
			// instructionpointer += 2
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
			return 0, nil //this case is a problem now
		default:
			return 0, errors.New("unexpected instruction value " + instruction)
		}
	}
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

// OutputResult oututs the int to standard out
func OutputResult(result int) {
	outputstring := strconv.Itoa(result)
	w := bufio.NewWriter(os.Stdout)
	fmt.Fprint(w, outputstring)
	fmt.Fprint(w, "\n")
	w.Flush()
}

//part two requires the state to be stored between rounds, to keep paralell I would need to have each amplifier keep a map with the key of the phases and a value of the intcode program, just making it serial
//have the amplifier run this, and modify input and output to write to the channels
//need to input the phase, then inputs come from the channel
func ProcessPartTwo(chin chan int, chout chan int, intcodeprogram []int, node int) ([]int, error) {
	ints := make([]int, len(intcodeprogram))
	copy(ints, intcodeprogram)
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
		log.Printf("node %d starting new loop\n", node)
		opcodeandparammodes = SplitOpCode(ints[instructionpointer])
		log.Printf("node %d opcodeandparammodes %v\n", node, opcodeandparammodes)
		instruction, ok = set[opcodeandparammodes[0]]
		if !ok {
			//todo I'm getting an invalid opcode of 93. Is there a way to verify if there is a problem with the intcode vs my program?
			log.Printf("node %d unexpected value "+strconv.Itoa(ints[instructionpointer])+" at position "+strconv.Itoa(instructionpointer)+"\n", node)
			close(chout)
			return ints, errors.New("unexpected value " + strconv.Itoa(ints[instructionpointer]) + " at position " + strconv.Itoa(instructionpointer))
		}
		log.Printf("node %d starting next instruction %v\n", node, instruction)

		//perform instruction
		switch instruction {
		case "addition":
			log.Printf("node %d processing addition\n", node)
			firstoperand = GetVal(ints, opcodeandparammodes, instructionpointer, 1)  //ints[params[0]]
			secondoperand = GetVal(ints, opcodeandparammodes, instructionpointer, 2) //ints[params[1]]
			resultpos = ints[instructionpointer+3]                                   // params[2]
			result = firstoperand + secondoperand
			ints[resultpos] = result
			instructionpointer += 4
		case "multiplication":
			log.Printf("node %d processing multiplication\n", node)
			firstoperand = GetVal(ints, opcodeandparammodes, instructionpointer, 1)  //ints[params[0]]
			secondoperand = GetVal(ints, opcodeandparammodes, instructionpointer, 2) //ints[params[1]]
			resultpos = ints[instructionpointer+3]                                   // params[2]
			result = firstoperand * secondoperand
			ints[resultpos] = result
			instructionpointer += 4
		case "input":
			log.Printf("node %d processing input\n", node)
			resultpos = ints[instructionpointer+1]
			result, ok = <-chin //get input
			if !ok {
				log.Printf("Node %d got an error on getting input, closing chout\n", node)
				close(chout)
				return ints, nil
			}
			log.Printf("node %d got input signal %d\n", node, result)
			ints[resultpos] = result
			instructionpointer += 2
		case "output":
			log.Printf("node %d processing output\n", node)
			result = GetVal(ints, opcodeandparammodes, instructionpointer, 1)
			chout <- result //output result
			instructionpointer += 2
			log.Printf("node %d FINISHED processing output\n", node)
		case "jump-if-true":
			log.Printf("node %d processing jump if true\n", node)
			// Opcode 5 is jump-if-true: if the first parameter is non-zero, it sets the instruction pointer to the value from the second parameter. Otherwise, it does nothing.
			firstoperand = GetVal(ints, opcodeandparammodes, instructionpointer, 1)
			secondoperand = GetVal(ints, opcodeandparammodes, instructionpointer, 2)
			if firstoperand != 0 {
				instructionpointer = secondoperand
			} else {
				instructionpointer += 3
			}
		case "jump-if-false":
			log.Printf("node %d processing jump if false\n", node)
			firstoperand = GetVal(ints, opcodeandparammodes, instructionpointer, 1)
			secondoperand = GetVal(ints, opcodeandparammodes, instructionpointer, 2)
			if firstoperand == 0 {
				instructionpointer = secondoperand
			} else {
				instructionpointer += 3
			}
		case "less-than":
			log.Printf("node %d processing less than\n", node)
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
			log.Printf("node %d processing equal\n", node)
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
			log.Printf("node %d got a terminate signal, stopping\n", node)
			close(chout)
			return ints, nil
		default:
			log.Printf("node %d error, got instruction %v", node, instruction)
			return ints, errors.New("unexpected instruction value " + instruction)
		}
		log.Printf("node %d ending switch\n", node)
	}
}
