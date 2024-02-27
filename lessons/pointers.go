package lessons

import "fmt"

func updateName(x *string) {
	*x = "luigi"
}

func Pointers() {

	name := "mario"
	fmt.Println("memory address of name is:", &name)

	namePointer := &name
	fmt.Println("memory address:", namePointer)
	fmt.Println("value at memory address:", *namePointer)

	updateName(namePointer)

	fmt.Println(name)

}