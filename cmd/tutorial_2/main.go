package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int = 32767
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	// cant perform arithmetic with mixed type

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	// var result float32 = floatNum32 + intNum32

	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	// Interger division results in integer, and result is rounded down

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2) // to get remainder

	var myString string = "Hello" + " " + "World"
	fmt.Println(myString)

	fmt.Println(len("У"))                    // 2 !- number of bytes , go uses UTF8 encoding
	fmt.Println(utf8.RuneCountInString("У")) // 1

	var myRune rune = 'a' // one char datatype
	fmt.Println(myRune)

	var myBoolean bool = false
	fmt.Println(myBoolean)

	var intNum3 int
	fmt.Println(intNum3)

	// for all int, float and rune defualt value is 0 (uint, uint8, uint16, uint32, uint64, int, int8,int16, int32,int64, float32, float64, rune)
	// for string it's empty string, for boolean it's false

	// var myVar = "text"
	myVar := "text" //shorthand
	fmt.Println(myVar)

	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	const myConst string = "const value"

	fmt.Println(myConst)

}
