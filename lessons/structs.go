package lessons

import "fmt"

type bill struct {
	name string
	items map[string]float64
	tip float64
}

func newBill(name string) bill {
	b := bill{
		name: name,
		items: map[string]float64{},
		tip: 0,
	}

	return b
}

func (b *bill) format() string {
	fs := "Bill Breakdown: \n"
	var total float64 = 0

	for key, value := range b.items {
		fs += fmt.Sprintf("%-25v ...$%0.2f \n", key + ":", value)
		total += value
	}

	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "tip:", b.tip)

	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total + b.tip)

	return fs
}

func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

func Structs() {

	myBill := newBill("javier's bill")

	myBill.addItem("soup", 4.50)
	myBill.addItem("chicken", 8.99)
	myBill.addItem("fries", 3.50)
	myBill.addItem("coffee", 3.25)

	myBill.updateTip(10)

	fmt.Println(myBill.format())

}