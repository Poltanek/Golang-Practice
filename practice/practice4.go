// Struct and Interfaces

package main

import "fmt"

// Struct for gas engine
type gasEngine struct {
	mpg     uint8
	gallons uint8
}

// Struct for electric engine
type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

// Method for gasEngine: calculates miles left based on fuel
func (e gasEngine) milesLeft() uint8 {
	return e.mpg * e.gallons
}

// Method for electricEngine: calculates miles left based on battery
func (e electricEngine) milesLeft() uint8 {
	return e.mpkwh * e.kwh
}

// Interface requiring milesLeft() method
type engine interface {
	milesLeft() uint8
}

// Function to check if the vehicle can reach the destination
func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Head to fuel up first!")
	}
}

func main() {
	// Creating an electric engine with 25 miles per kWh and 15 kWh battery
	var myEngine electricEngine = electricEngine{25, 15}
	canMakeIt(myEngine, 50) // Checking if it can travel 50 miles

	// Creating a gas engine with 30 mpg and 2 gallons
	var myGasEngine gasEngine = gasEngine{30, 2}
	canMakeIt(myGasEngine, 50) // Checking if it can travel 50 miles
}
