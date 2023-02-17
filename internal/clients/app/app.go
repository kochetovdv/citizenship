package app

import (
	//"citizenship/internal/downloader"
	//	"citizenship/internal/order"
	"citizenship/internal/downloader"
	"citizenship/internal/order"
	"citizenship/internal/parser/pdfparser"
	"citizenship/internal/parser/siteparser"
)

type expSiteParser interface {
	Parse(url string) *order.Orders
}

type expFileDownloader interface {
	Download(*order.Orders) error
}

// TODO path, filename. Possible return data
type expPDFParser interface {
	Parse(filePDF string)
}

type App struct {
	//	logger Logger
	//	ctx    Context
	siteParser     expSiteParser
	fileDownloader expFileDownloader
	pdfParser      expPDFParser
}

// func NewApp(logger Logger, ctx Context) *app {
func NewApp() *App {
	/*	a := new(app)
		a.ctx = ctx
		a.logger = logger
		return a*/
	a := new(App)
	a.siteParser = siteparser.NewSiteParser("http://cetatenie.just.ro/ordine-articolul-11/")
	a.fileDownloader = downloader.NewDownloader("./downloads")

	//TODO
	a.pdfParser = pdfparser.NewParser("./downloads")

	return a
}

func (a *App) Run() {
	parsedListOfOrders:=a.siteParser.Parse("http://cetatenie.just.ro/ordine-articolul-11/")
	//TODO Add file downloader
	a.fileDownloader.Download(parsedListOfOrders)
	a.pdfParser.Parse("")
}
