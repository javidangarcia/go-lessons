package main

import (
	"fmt"
	"math"
	"strings"
)

func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Goodbye %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, value := range n {
		f(value)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, value := range names {
		initials = append(initials, value[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func functions() {

	sayGreeting("mario")
	sayGreeting("luigi")
	sayBye("mario")

	cycleNames([]string{"bowser", "koopa", "peach"}, sayGreeting)
	cycleNames([]string{"bowser", "koopa", "peach"}, sayBye)

	a1 := circleArea(10.5)
	a2 := circleArea(15)

	fmt.Printf("circle 1 is %0.3f and circle 2 is %0.3f \n", a1, a2)

	fn1, ln1 := getInitials("javier garcia")
	fmt.Println(fn1, ln1)

	fn2, ln2 := getInitials("jaylen waddle")
	fmt.Println(fn2, ln2)

	fn3, ln3 := getInitials("javier")
	fmt.Println(fn3, ln3)

}