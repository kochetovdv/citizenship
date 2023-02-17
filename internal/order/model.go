package order

import (
	"fmt"
	_ "regexp"
	_ "strconv"
	_ "strings"

	"github.com/google/uuid"
	_ "github.com/pkg/errors"
)

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
func (o *Orders) Add(order *Order) {
	for _, o1 := range o.Orders {
		if o1.id == order.id {
			return
		}
	}
	o.Orders = append(o.Orders, order)
}

// Prints the orders to the console
func (o *Orders) Print() {
	for i, order := range o.Orders {
		fmt.Printf("%d. %s\n", i+1, order)
		//		fmt.Printf("%d. Date:%s\nFilename:%s\nLink:%s\nNumber:%s\n", i+1, order.Date, order.Filename, order.Link, order.Number)
	}
}

// Prints statistics about the orders
func (o *Orders) Statistics() {
	total := len(o.Orders)
	fmt.Printf("Total orders: %d\n", total)
}

type Order struct {
	id       string
	Date     string
	Filename string
	Link     string
	Number   string
}

// Creates a new order
func NewOrder(date string, filename string, link string, number string) *Order {
	o := new(Order)

	id := uuid.New()
	o.id = id.String()

	o.Date = date
	o.Filename = filename
	o.Link = link
	o.Number = number
	return o
}

// Prints the order to the console
func (o *Order) Print() {
	fmt.Printf("Date:%s\tFilename:%s\tLink:%s\tNumber:%s\t", o.Date, o.Filename, o.Link, o.Number)
}
