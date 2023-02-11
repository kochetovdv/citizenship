package app

import (
	"citizenship/internal/downloader"
	"citizenship/internal/order"
	"citizenship/internal/parser"
)

type expDownloader interface {
	Download() []order.Order
}

type expParser interface {
	Parse(pdf string)
}

type App struct {
	//	logger Logger
	//	ctx    Context
	downloader expDownloader
	parser     expParser
}

// func NewApp(logger Logger, ctx Context) *app {
func NewApp() *App {
	/*	a := new(app)
		a.ctx = ctx
		a.logger = logger
		return a*/
	a := new(App)
	a.downloader = downloader.NewDownloader("http://cetatenie.just.ro/ordine-articolul-11/", "./downloads")

	//TODO
	a.parser = parser.NewParser("", "")

	return a
}

func (a *App) Run() {
	//	a.downloader.Download()
	a.parser.Parse("")
}
