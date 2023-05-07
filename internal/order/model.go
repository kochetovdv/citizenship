package order

import (
	"fmt"
)

type Orders struct {
	Orders map[string]Order
}

// Create a list of orders
func NewOrders() *Orders {
	return &Orders{
		make(map[string]Order),
	}
}

// Adds a new order to the list of orders
func (o *Orders) Add(filename string, order Order) {
	o.Orders[filename] = order
}

// Delete an order from the list of orders
func (o *Orders) Delete(filename string) {
	delete(o.Orders, filename)
}

// Prints statistics about the orders
func (o *Orders) Statistics() {
	totalOrders := len(o.Orders)
	fmt.Printf("Total orders: %d\n", totalOrders)
}

// TODO filename to key of map
type Order struct {
	Date     string
	Link     string
	Number   string
}

// Creates a new order
func NewOrder(date string, link string, number string) Order {
	o := new(Order)
	o.Date = date
	o.Link = link
	o.Number = number
	return *o
}
