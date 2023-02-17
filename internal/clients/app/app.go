package app

import (
	//"citizenship/internal/downloader"
	//	"citizenship/internal/order"
	"citizenship/internal/order"
	"citizenship/internal/parser/pdfparser"
	"citizenship/internal/parser/siteparser"
)

type expSiteParser interface {
	Parse(url string) *order.Orders
}

type expFileDownloader interface {
	Download(url string) error
}

type expPDFParser interface {
	Parse(pdf string)
}

type App struct {
	//	logger Logger
	//	ctx    Context
	siteParser expSiteParser
	//	fileDownloader expFileDownloader
	pdfParser expPDFParser
}

// func NewApp(logger Logger, ctx Context) *app {
func NewApp() *App {
	/*	a := new(app)
		a.ctx = ctx
		a.logger = logger
		return a*/
	a := new(App)
	a.siteParser = siteparser.NewSiteParser("http://cetatenie.just.ro/ordine-articolul-11/", "./downloads")

	// TODO file downloader
	// a.fileDownloader = downloader.NewFileDownloader("./downloads")

	//TODO
	a.pdfParser = pdfparser.NewParser("", "")

	return a
}

func (a *App) Run() {
	a.siteParser.Parse("http://cetatenie.just.ro/ordine-articolul-11/")

	a.pdfParser.Parse("")
}
