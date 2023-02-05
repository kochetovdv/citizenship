package order

import (
	_ "fmt"
	_ "regexp"
	_ "strconv"
	_ "strings"

	_ "github.com/pkg/errors"

	_ "github.com/jung-kurt/gofpdf" // or use "github.com/hhrutter/pdfcpu"
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

func Main() {
    // Open the PDF file
 /*   pdf, err := gofpdf.Open("input.pdf")
    if err != nil {
        panic(errors.Wrap(err, "error opening PDF file"))
    }
    defer pdf.Close()

    // Extract the text from the PDF
    text, err := pdf.ExtractText()
    if err != nil {
        panic(errors.Wrap(err, "error extracting text from PDF"))
    }

    // Split the text into lines
    lines := strings.Split(text, "\n")

    // Create a slice to store the parsed DigitYear objects
    var digitYears []DigitYear

    // Iterate over the lines and extract the digits and year
    for _, line := range lines {
        // Extract the digits using a regular expression
        re := regexp.MustCompile("[0-9]+")
        digits := re.FindAllString(line, -1)

        // Extract the year using another regular expression
        yearRe := regexp.MustCompile("year: [0-9]+")
        yearStr := yearRe.FindString(line)
        year, _ := strconv.Atoi(yearStr[6:])

        if len(digits) == 0 || year == 0 {
            continue
        }

        digit, _ := strconv.Atoi(digits[0])
        digitYears = append(digitYears, DigitYear{Digit: digit, Year: year})
    }

    fmt.Println(digitYears)*/
}