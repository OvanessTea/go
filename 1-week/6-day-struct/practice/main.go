package main

import "fmt"

type Product struct {
	ID    int64
	Name  string
	Price float64
}

type User struct {
	ID   int64
	Name string
	Age  int
}

func main() {
	// Task 1
	products := []Product{
		{
			ID:    1,
			Name:  "oreo",
			Price: 1.59,
		},
		{
			ID:    2,
			Name:  "car",
			Price: 20000.50,
		},
	}
	for _, product := range products {
		fmt.Println(product.ID)
		fmt.Println(product.Name)
		fmt.Println(product.Price)
	}

	// Task 2
	fmt.Println(products[0].isExpensive())
	fmt.Println(products[1].isExpensive())

	// Task 3
	fmt.Println(products[1].Price)
	products[1].ApplyDiscount(10)
	fmt.Println(products[1].Price)

	// Task 4
	changeName(products[0], "New")     // keeps its own name
	changeNamePtr(&products[1], "New") // changes to New
	fmt.Println(products)

	// Task 5
	users := []User{
		{ID: 1, Name: "Alice", Age: 20},
		{ID: 2, Name: "Bob", Age: 25},
		{ID: 3, Name: "Charlie", Age: 30},
	}
	makeAdultsOlder(users)
	fmt.Println(users)

	// Task to think 1
	// user.Age == 25 'cause we don't use pointer receiver as props

	// Task to think 2
	// user.Age == 100 'cause we use pointer receiver and address of  original value

	// Task to think 3
	// No for _, user := range users uses copies and won't change the original slice

	// Task to think 4
	// for i := range users will change original slice 'cause users[i] use address of the original struct

}

func (p Product) isExpensive() bool {
	return p.Price >= 1000
}

func (p *Product) ApplyDiscount(percent float64) {
	p.Price = p.Price / 100 * (100 - percent)
}

func changeName(product Product, name string) {
	product.Name = name
}

func changeNamePtr(product *Product, name string) {
	product.Name = name
}

func makeAdultsOlder(users []User) {
	for i := range users {
		users[i].Age++
	}
}
