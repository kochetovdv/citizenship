package parser

import (
	"fmt"
	"io/ioutil"
	"regexp"
	_ "strconv"
	"strings"
	_ "strings"

	"github.com/ledongthuc/pdf"
)

type Parser struct {
	path    string
	filePDF string
}

func NewParser(path, filePDF string) *Parser {
	p := new(Parser)
	p.path = path
	p.filePDF = filePDF
	return p
}

// TODO add error handling, change return data type to Orders.
func (p *Parser) Parse(fileName string) { //([]issue.Issue, error) {
	// Open the PDF file
	//TODO
	file, pdfReader, err := pdf.Open("downloads/O-15-din-07.01.2022-NPE.pdf")
	if err != nil {
		//TODO
		fmt.Println("error in file opening")
		//	return nil, fmt.Errorf("error during open file %s", err)
	}
	defer file.Close()

	// Extract the text from the PDF
	text, err := pdfReader.GetPlainText()
	if err != nil {
		fmt.Println("error in extracting plain text")
		// TODO
		//		return nil, fmt.Errorf("error during text extracting from file %s", err)
	}

	data, _ := ioutil.ReadAll(text)

	// Split the text into lines
	lines := strings.Split(string(data), "\n")

	// Create a slice to store the parsed DigitYear objects
	//	var digitYears []issue.Issue

	// Iterate over the lines and extract the digits and year
	for _, line := range lines {
		// Extract the digits using a regular expression
		re := regexp.MustCompile(`(\d+\/\d{4})`)
		digits := re.FindAllString(line, -1)
		/*
			// Extract the year using another regular expression
			yearRe := regexp.MustCompile("year: [0-9]+")
			yearStr := yearRe.FindString(line)
			year, _ := strconv.Atoi(yearStr[6:])

			if len(digits) == 0 || year == 0 {
				continue
			}

			digit, _ := strconv.Atoi(digits[0])
			digitYears = append(digitYears, issue.Issue{Number: digit, Year: year}) */
		fmt.Println(digits)
	}

	//	fmt.Println(text)

	/*	for _, line := range digitYears {
		fmt.Println(line)
	}*/

	//	return nil, nil
}
