package day1

//Calc  takes in the mass and returns the fuel required
func Calc(mass float64) int {
	div := mass / float64(3)
	return int(div) - 2
}

//CalcWithFuelAdjustment adjusts for fuel needed for fuel mass
func CalcWithFuelAdjustment(mass float64) int64 {
	var totalFuel int64 = 0
	var round float64 = mass

	for {
		fuel := Calc(round)
		if fuel <= 0 {
			break
		}
		totalFuel += int64(fuel)
		round = float64(fuel)
	}
	return totalFuel
}
