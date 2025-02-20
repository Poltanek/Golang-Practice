// Including external packages
package main

// Including external libraries
// Allows to fcall functions like
// fmt.Println(), fmt.Printf(), fmt.Scan().
import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// 2,147,483,647 is how much a int32 can hold
	// Maximum it can hold within a int16
	var Num int = 32767
	Num = Num + 1
	fmt.Println(Num)
	var floatNum float32 = 12345678.9
	fmt.Println(floatNum)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)

	var myString string = "Hello" + " " + "World"
	fmt.Println(myString)

	// Counts the number of runes "Y"
	// utf8 package handles UTF-8 encoding which is the standard encoding for text in Go
	// A rune in Go is an alias for int32, representing a Unicode character
	fmt.Println(utf8.RuneCountInString("Y"))

	var myRune rune = 'a'
	fmt.Println(myRune)

	var myBoolean bool = false
	fmt.Println(myBoolean)

	var intNum3 rune
	fmt.Println(intNum3)

	var1, var2 := 1, 2
	fmt.Println(var1, var2)

}
