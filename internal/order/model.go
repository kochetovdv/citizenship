package order

import (
	_ "fmt"
	_ "regexp"
	_ "strconv"
	_ "strings"

	_ "github.com/pkg/errors"

	//_ "github.com/jung-kurt/gofpdf" // or use "github.com/hhrutter/pdfcpu"
)

type Orders struct {
   Orders []*Order
}

func NewOrders() *Orders {
    return &Orders{
        []*Order{},
    }
}

func (o *Orders) Add(order *Order) {
    o.Orders = append(o.Orders, order)
}

type Order struct {
	Date     string
	Filename string
	Link     string
	Number   string
}