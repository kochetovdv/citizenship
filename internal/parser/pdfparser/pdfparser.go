package pdfparser

import (
	"citizenship/internal/issue"
	"citizenship/internal/order"
	"fmt"
	"io/ioutil"
	"regexp"
	_ "strconv"
	"strings"

	"github.com/ledongthuc/pdf"
	//_ "github.com/jung-kurt/gofpdf" // or use "github.com/hhrutter/pdfcpu"
)

type PDFParser struct {
	path string
}

func NewParser(path string) *PDFParser {
	p := new(PDFParser)
	p.path = path
	return p
}

// Parse the pdf file and return the list of orders with the issues
func (p *PDFParser) Parse(orders *order.Orders) (*order.Orders, error) {
	for _, order := range orders.Orders {
		// Open the PDF file
		file, pdfReader, err := pdf.Open(p.path + "/" + order.Filename)

		if err != nil {
			fmt.Printf("error in openning file: %s\n", err)
			file.Close()
			continue
		}
		defer file.Close()

		// Extract the text from the PDF
		text, err := pdfReader.GetPlainText()
		if err != nil {
			file.Close()
			fmt.Printf("error during text extracting from file %s\n", err)
			continue
		}

		data, _ := ioutil.ReadAll(text)

		// Split the text into lines
		lines := strings.Split(string(data), "\n")
		issues := issue.NewIssues()
		for _, line := range lines {
			// Extract the digits using a regular expression
			re := regexp.MustCompile(`(\d+\/\d{4})`)
			digits := re.FindAllString(line, -1)
			if len(digits) == 0 {
				continue
			}
			for _, digit := range digits {
				issues.Add(&issue.Issue{Number: digit})
			}
			order.Issues = *issues
		}
	}
	return orders, nil
}
