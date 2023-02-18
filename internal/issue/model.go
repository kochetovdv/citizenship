package issue

import "fmt"

type Issues struct {
	Issues []*Issue
}

func NewIssues() *Issues {
	return &Issues{
		[]*Issue{},
	}
}

func (i *Issues) Add(issue *Issue) {
	i.Issues = append(i.Issues, issue)
}

// Prints the issues to the console
func (i *Issues) Print() {
	for i, issue := range i.Issues {
		fmt.Printf("%d. %s\n", i+1, issue)
	}
}

// Prints statistics about the issues
func (i *Issues) Statistics() {
	total := len(i.Issues)
	fmt.Printf("Total issues: %d\n", total)
}

// Gets the number of issues
func (i *Issues) Count() int {
	return len(i.Issues)
}

type Issue struct {
	Number string
}

func NewIssue(number string) *Issue {
	return &Issue{
		Number: number,
	}
}

func (i *Issue) Print() {
	fmt.Printf("Number: %s\n", i.Number)
}