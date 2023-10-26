package main

import (
	"fmt"
	"strings"
)

func main() {
	// var myString = "Résumé"
	var myString = []rune("Résumé")
	// we can index string just like an array
	var indexed = myString[1]
	fmt.Printf("%v, %T\n", indexed, indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}

	// go uses utf-8 to represent string

	// An easier way to deal with iterating and indexing strings
	// is to cast them to an array of runes rather than dealing with
	// the underlying byte array of a string

	var myRune = 'a'
	fmt.Printf("\nmyRune = %v", myRune)

	var strSlice = []string{"k", "i", "s", "h", "o", "r"}
	var catStr = "" // Instead use strings Builder
	var stringBuilder strings.Builder
	for i := range strSlice {
		// catStr += strSlice[i]
		stringBuilder.WriteString(strSlice[i])
	}

	catStr = stringBuilder.String()

	fmt.Printf("\n%v", catStr)

	// Strings are immutable in go
}
