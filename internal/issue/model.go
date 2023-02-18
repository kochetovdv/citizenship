package issue

import "fmt"

type Issues struct {
	issues map[string][]string
}

func NewIssues() *Issues {
	return &Issues{
		make(map[string][]string),
	}
}

func (i *Issues) Add(filename string, issues []string) {
	i.issues[filename] = issues
}

// Prints the issues to the console
func (i *Issues) Print() {
	for filename, issues := range i.issues {
		for _, issue := range issues {
			fmt.Printf("Filename:%s\tissue:%s\n", filename, issue)
		}
	}
}

// Prints statistics about the issues
func (i *Issues) Statistics() {
	totalIssues := 0
	for _, issues := range i.issues {
		totalIssues += len(issues)
	}
	fmt.Printf("Total issues: %d\n", totalIssues)
}

// Gets the number of issues
func (i *Issues) Count() int {
	return len(i.issues)
}
