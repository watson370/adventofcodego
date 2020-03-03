package day4

import (
	"log"
	"strconv"
)

//Test tests if an int meets the criteria  pass 1 for part1 and 2 for part2
func Test(subject int, min int, max int, part int) bool {
	//The value is within the range given in your puzzle input.
	if subject < min || subject > max {
		log.Printf("%d is not in range %d - %d", subject, min, max)
		return false
	}
	strRep := strconv.Itoa(subject)
	//It is a six-digit number.
	if len(strRep) != 6 {
		log.Printf("%d has the wrong number of digits, contains %d", subject, len(strRep))
		return false
	}
	//ascii 48-57 is 0-9, we don't need actual int value
	minsize := byte(0) // any num larger than 57
	repeatcount := 0
	containsRepeat := false
	containsOnlyTwoRepeat := false
	for i := 0; i < len(strRep); i++ {
		place := strRep[i]
		// log.Printf("place %d minsize %d", place, minsize)
		//Going from left to right, the digits never decrease; they only ever increase or stay the same (like 111123 or 135679).
		if place < minsize {
			// log.Printf("%d is less than %d, violated no decreasing", place, minsize)
			return false
		}
		if place == minsize {
			containsRepeat = true
			if repeatcount == 0 {
				repeatcount = 2
			} else {
				repeatcount++
			}
		} else {
			//make sure it was even
			if part == 2 && repeatcount == 2 {
				// log.Printf("%s failed with an odd number of repeats ", strRep)
				containsOnlyTwoRepeat = true
			}
			//zero it out
			repeatcount = 0
		}
		if place > minsize {
			minsize = place
		}
	}
	//check again to cover last digit repeats
	if part == 2 && repeatcount == 2 {
		// log.Printf("%s failed with an odd number of repeats on last digit", strRep)

		containsOnlyTwoRepeat = true
	}

	//Two adjacent digits are the same (like 22 in 122345).
	if !containsRepeat {
		// log.Println("does not contain a consecutive double")
		return false
	}
	if part == 2 && !containsOnlyTwoRepeat {
		return false
	}
	return true
}
