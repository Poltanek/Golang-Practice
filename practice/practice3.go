// Arrays, Slices, Maps and Loops

package main

import "fmt"

func printMe(value string) {
	fmt.Println(value) // Print the value passed as an argument
}

func main() {
	printMe("================")

	// Creating an array of int32 with fixed size (3 elements)
	intArr := [...]int32{1, 2, 3}
	fmt.Println(intArr) // Output: [1 2 3]

	// Creating a slice of int32
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Println(intSlice) // Output: [4 5 6]

	// Appending an element (7) to the slice
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice) // Output: [4 5 6 7]

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap2["Adam"])
	var age, ok = myMap2["Jason"]
	if ok {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Println("Invalid Name")
	}

	for name := range myMap2 {
		fmt.Printf("Name: %v, Age:%v \n", name, age)
	}

	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}

	var i int = 0
	for i < 10 {
		fmt.Println(i)
		i = i + 1
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	var myString = []rune("resume")
	var indexed = myString[1]
	fmt.Printf("%v, %T\n", indexed, indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}
	fmt.Printf("\n The length of myString is %v", len(myString))

	var myRune = 'a'
	fmt.Printf("\nmyRune = %v", myRune)
}
