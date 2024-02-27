package lessons

import "fmt"

func Maps() {

	menu := map[string]float64{
		"soup": 4.99,
		"pie": 7.99,
		"salad": 6.99,
		"chicken": 8.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	for key, value := range menu {
		fmt.Println(key, "-", value)
	}

	phonebook := map[int]string{
		3052446789: "mario",
		7863372036: "luigi",
		3054994532: "peach",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[3052446789])

	phonebook[7863372036] = "bowser"

	fmt.Println(phonebook)

}