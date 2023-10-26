package main

import (
	"errors"
	"fmt"
)

func main() {
	var printvalue string = "hello"
	printMe(printvalue)

	var numerator int = 11
	var denominator int = 0
	var divisionResult, remainder, err = intDivision(numerator, denominator)
	// if err!=nil{
	// 	fmt.Printf(err.Error())
	// }else if remainder == 0 {
	// 	fmt.Printf("The result of the interger division is %v \n", divisionResult)
	// } else {
	// 	fmt.Printf("The result of the interger division is %v with remainder %v \n", divisionResult, remainder)
	// }

	switch {
	case err != nil:
		fmt.Printf(err.Error() + "\n")
	case remainder == 0:
		fmt.Printf("The result of the interger division is %v \n", divisionResult)
	default:
		fmt.Printf("The result of the interger division is %v with remainder %v \n", divisionResult, remainder)
	}

	switch remainder {
	case 0:
		fmt.Printf("The division was exact")
	case 1, 2:
		fmt.Printf("The division was close")
	default:
		fmt.Printf("The division was not close")
	}
}

func printMe(printVal string) {
	fmt.Println(printVal)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot Divide by Zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
