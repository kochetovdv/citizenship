package order

import (
	"fmt"
)

// TODO make map[string]Order
type Orders struct {
	Orders []*Order
}

// Create a list of orders
func NewOrders() *Orders {
	return &Orders{
		[]*Order{},
	}
}

// Adds a new order to the list of orders
func (o *Orders) Add(order Order) {
	for _, o := range o.Orders {
		if o.Filename == order.Filename {
			return
		}
	}
	o.Orders = append(o.Orders, &order)
}

// Adds a list of orders to the list of orders
func (o *Orders) AddRange(orders ...Order) {
	for _, order := range orders {
		o.Add(order)
	}
}

// Prints statistics about the orders
func (o *Orders) Statistics() {
	totalOrders := len(o.Orders)
	fmt.Printf("Total orders: %d\n", totalOrders)
}

// TODO filename to key of map
type Order struct {
	Date     string
	Filename string
	Link     string
	Number   string
}

// Creates a new order
func NewOrder(date string, filename string, link string, number string) Order {
	o := new(Order)
	o.Date = date
	o.Filename = filename
	o.Link = link
	o.Number = number
	return *o
}
