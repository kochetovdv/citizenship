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
func (p *PDFParser) Parse(orders *order.Orders) (issue.Issues, error) {
	issueList := issue.NewIssues()
	var ordersToDelete []string
	for filename := range orders.Orders {
		// Open the PDF file
		file, pdfReader, err := pdf.Open(p.path + "/" + filename)
		if err != nil {
			fmt.Printf("error in openning file: %s\n", err)
			fmt.Println(filename)
			ordersToDelete = append(ordersToDelete, filename)
			file.Close()
			continue
		}
		defer file.Close()

		// Extract the text from the PDF
		text, err := pdfReader.GetPlainText()
		if err != nil {
			fmt.Printf("error during text extracting from file %s\n", err)
			fmt.Println(filename)
			ordersToDelete = append(ordersToDelete, filename)
			file.Close()
			continue
		}

		data, _ := ioutil.ReadAll(text)

		// Split the text into lines
		lines := strings.Split(string(data), "\n")
		var parsedIssues []string
		for _, line := range lines {
			// Extract the digits using a regular expression
			re := regexp.MustCompile(`(\d+\/\d{4})`)
			digits := re.FindAllString(line, -1)
			if len(digits) == 0 || digits == nil {
				continue
			}
			parsedIssues = append(parsedIssues, digits...)
		}
		issueList.Add(filename, parsedIssues)
	}
	for _, filename := range ordersToDelete {
		orders.Delete(filename)
	}
	return *issueList, nil
}


// TODO or remove. Possible delete and redownload
// Delete files from folder and that are not in the list of orders
func (p *PDFParser) Delete(orders *order.Orders) {
	
}