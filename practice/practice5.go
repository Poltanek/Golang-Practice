// Pointers

package main

import (
	"fmt"
)

func main() {
	var var1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\nThe memory location of the var1 array is %p", &var1)

	var result [5]float64 = square(var1)
	fmt.Printf("\nThe result is: %v", result)
	fmt.Printf("\nThe value of thing1 is: %v", var1)
}

func square(var2 [5]float64) [5]float64 {
	fmt.Printf("\n The memory location of the var2 array is; %p", &var2)
	for i := range var2 {
		var2[i] = var2[i] * var2[i]
	}

	return var2
}
