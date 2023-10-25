package main
import "fmt"

func main(){
	var p *int32 = new(int32)
	var i int32
	fmt.Printf("The value p ponints to is: %v \n",*p)
	fmt.Printf("\nThe value if i is: %v \n",i)
	p = &i
	*p = 1
	fmt.Printf("The value p ponints to is: %v \n",*p)
	fmt.Printf("\nThe value if i is: %v \n",i)

	// Slices contain pointers to an underlying array

	var slice = []int32{1,2,3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Println(slice)
	fmt.Println(sliceCopy)


	var thing1 = [5]float64{1,2,3,4,5}
	fmt.Printf("\n The memory location of the thing1 array is:%p", &thing1)

	var result [5]float64 = square(thing1)
	var result2 [5]float64 = squareWithPointers(&thing1)
	fmt.Printf("\nThe result is: %v", result)
	fmt.Printf("\nThe result2 is: %v", result2)
}

func square(thing2 [5]float64) [5]float64 {
	fmt.Printf("\n The memory location of the thing2 array is:%p", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i]*thing2[i];
	}
	return thing2
}

func squareWithPointers(thingP *[5]float64) [5]float64 {
	fmt.Printf("\n The memory location of the thingP array is:%p", thingP)
	for i := range thingP {
		thingP[i] = thingP[i]*thingP[i];
	}
	return *thingP
}

