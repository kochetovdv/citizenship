package issue

type Issues struct {
	Issues []*Issue
}

func NewOrders() *Issues {
	return &Issues{
		[]*Issue{},
	}
}

func (i *Issues) Add(issue *Issue) {
	i.Issues = append(i.Issues, issue)
}

type Issue struct {
	Filename string
	Number   int
	Year     int
}
