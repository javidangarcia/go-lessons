package main

import "fmt"

func variables() {

	// strings
	var nameOne string = "mario"
	var nameTwo = "luigi"
	var nameThree string
	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "peach"
	nameThree = "bowser"
	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "yoshi"
	fmt.Println(nameFour)

	// integers
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40
	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory
	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint16 = 256
	fmt.Println(numOne, numTwo, numThree)

	// floats
	var scoreOne float32 = 25.98
	var scoreTwo float64 = 883759237.76
	scoreThree := 22.5
	fmt.Println(scoreOne, scoreTwo, scoreThree)

}