package main

import (
	"fmt"
	"os"
	"strconv"
)

const MAXIMUM_SIZE = 3
const MAXIMUM_ARGRUMENTS = MAXIMUM_SIZE + 1
const MAXIMUM_VALUE = 200
const MINNIMUM_VALUE = 0

func main() {
	if ShowHelpIfArgsMissMatch() {
		return
	}

	var sides [MAXIMUM_SIZE]int
	var isInteger bool
	if isInteger, sides = IsAllArgsAreInt(); !isInteger {
		fmt.Println("[!] Please enter all 3 parameters with integer value.")
		return
	}

	if IsAnyNegativeValue(sides) {
		fmt.Println("[!] All 3 argruments must be positive value.")
		return
	}

	if IsAnyValueMoreThanMaximumRange(sides, MAXIMUM_VALUE) {
		fmt.Println("[!] There is some value is out of range.")
		return
	}

	if SumOfTwoSidesOfTriangleMustBeMoreThanTheLeftSide(sides) {
		fmt.Println("[!] Not a triangle.")
		return
	}

	if IsAnyRightTriangle(sides) {
		fmt.Println("[A] Right Triangle.")
		return
	}

	if IsEquilateral(sides) {
		fmt.Println("[A] Equilateral.")
		return
	}

	if IsAnyIsosceles(sides) {
		fmt.Println("[A] Isosceles.")
		return
	}

	fmt.Println("[!] Not a truangle.")
}

func ShowHelpIfArgsMissMatch() bool {
	if len(os.Args) != MAXIMUM_ARGRUMENTS {
		fmt.Println("[I] Please enter 3 parameters for 3 sides of triangle.")
		fmt.Println("    eg. " + os.Args[0] + " 3 4 5")
		return true
	}
	return false
}

func IsAllArgsAreInt() (isTriangle bool, sides [MAXIMUM_SIZE]int) {
	for i := 1; i < MAXIMUM_SIZE+1; i++ {
		side, err := strconv.Atoi(os.Args[i])
		if err != nil {
			isTriangle = false
			return
		} else {
			sides[i-1] = side
		}
	}
	isTriangle = true
	return
}

func IsAnyNegativeValue(sides [MAXIMUM_SIZE]int) bool {
	for _, sideWidth := range sides {
		if sideWidth < MINNIMUM_VALUE {
			return true
		}
	}
	return false
}

func IsAnyValueMoreThanMaximumRange(sides [MAXIMUM_SIZE]int, maximumValueValidable int) bool {
	for _, sideWidth := range sides {
		if sideWidth > maximumValueValidable {
			return true
		}
	}
	return false
}

func SumOfTwoSidesOfTriangleMustBeMoreThanTheLeftSide(sides [MAXIMUM_SIZE]int) bool {
	if IsSumOfABIsLessThanC(sides[0], sides[1], sides[2]) {
		return true
	} else if IsSumOfABIsLessThanC(sides[1], sides[2], sides[0]) {
		return true
	}
	return false
}

func IsSumOfABIsLessThanC(a int, b int, c int) bool {
	if a+b < c {
		return true
	}
	return false
}

func IsAnyRightTriangle(sides [MAXIMUM_SIZE]int) bool {
	if IsRightTriangle(sides[0], sides[1], sides[2]) ||
		IsRightTriangle(sides[0], sides[2], sides[1]) ||
		IsRightTriangle(sides[1], sides[2], sides[0]) {
		return true
	}
	return false
}

func IsRightTriangle(a int, b int, c int) bool {
	return PowerTwo(a)+PowerTwo(b) == c
}

func PowerTwo(number int) int {
	return number * number
}

func IsEquilateral(sides [MAXIMUM_SIZE]int) bool {
	for i := 1; i < len(sides); i++ {
		if sides[0] != sides[i] {
			return false
		}
	}
	return true
}

func IsAnyIsosceles(sides [MAXIMUM_SIZE]int) bool {
	if IsIsosceles(sides[0], sides[1], sides[2]) ||
		IsIsosceles(sides[0], sides[2], sides[1]) ||
		IsIsosceles(sides[1], sides[2], sides[0]) {
		return true
	}
	return false
}

func IsIsosceles(a int, b int, c int) bool {
	return a == b && a+b > c
}
