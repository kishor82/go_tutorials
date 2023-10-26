package main

import (
	"fmt"
	"time"
)

func main() {
	// Arrays:
	// Fixed length,Same Type, Indexable, Contiguous in Memory

	// var intArr [3]int32
	// var intArr [3]int32 = [3]int32{1,2,3}
	intArr := [...]int32{1, 2, 3}

	intArr[1] = 123

	fmt.Println(intArr)
	fmt.Println(intArr[1:3]) // print 1st and 3rd element value

	// Print memory location

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	// Slices Wrap arrays to give a more general, powerful and convenient interface to sequences of data.

	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Println(intSlice)

	fmt.Printf("The length is %v with capacity %v \n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("The length is %v with capacity %v \n", len(intSlice), cap(intSlice))
	// fmt.Println(intSlice[4]) - index out of range error

	// append multiple values

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)

	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 3, 8) // create slice with length 3 and capacity 8

	fmt.Println(intSlice3)

	// MAP

	// create map
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	// create and assign map
	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2["Adam"])

	// if key doesn't exist it will return default value of type
	fmt.Println(myMap2["Something"]) // 0

	var age, ok = myMap2["Jason"]
	if ok {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Println("Invalid Name")
	}

	// Iterate through map : No order preserved in map

	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v \n", name, age)
	}

	delete(myMap2, "Adam") // delete by reference

	// Iterate Array
	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}

	// go does not have while loop but

	var i int = 0
	for {
		if i >= 10 {
			break
		}
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println("-------------------------------")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("-------------------------------")

	// Speed test : showing the benefits of setting the capacity of your slice ahead of time if possible

	var n int = 1000000
	var testSlice = []int{}
	var testSlice2 = make([]int, n)

	fmt.Printf("Total time without preallocation: %v\n", timeLoop(testSlice, n))
	fmt.Printf("Total time with preallocation: %v", timeLoop(testSlice2, n))

}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}

	return time.Since(t0)
}
