package order

import (
	"citizenship/internal/issue"
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

func (o *Orders) AddRange(orders ...*Order) {
	for _, order := range orders {
		o.Add(order)
	}
}

// Prints statistics about the orders
func (o *Orders) Statistics() {
	totalOrders := len(o.Orders)
	fmt.Printf("Total orders: %d\n", totalOrders)
	totalIssues := 0
	for _, order := range o.Orders {
		totalIssues += order.Issues.Count()
	}
	fmt.Printf("Total issues: %d\n", totalIssues)

}

func (o *Orders) AddIssue(issue *issue.Issue) {

}

type Order struct {
	id       string
	Date     string
	Filename string
	Link     string
	Number   string
	Issues   issue.Issues
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