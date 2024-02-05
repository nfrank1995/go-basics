package structs

import "fmt"

type item struct {
	name  string
	price int
}

type price interface {
	calculatePrice() int
}

type basket struct {
	items []item
}

func (b basket) calculatePrice() int {
	endPrice := 0
	for _, i := range b.items {
		endPrice += i.price
	}
	return endPrice
}

func (b *basket) addItems(i ...item) {
	b.items = append(b.items, i...)
}

func (b basket) printItems() {
	if len(b.items) == 0 {
		fmt.Printf("Your basket is empty!\n")
	} else {
		fmt.Printf("Items in basket: %v\n", b.items)
	}
}

func printTotal(p price) {
	fmt.Printf("Total price: %d\n", p.calculatePrice())
}

func Run() {
	basket := basket{}
	fmt.Printf("Hello! Here are the items in your basket.\n")
	basket.printItems()
	basket.addItems(
		item{
			name:  "Banana",
			price: 69,
		},
		item{
			name:  "Tomate",
			price: 299,
		},
		item{
			name:  "Bread",
			price: 189,
		},
		item{
			name:  "Fish",
			price: 1350,
		},
	)
	basket.printItems()
	printTotal(basket)
}
