package app

import (
	"citizenship/internal/downloader"
	"citizenship/internal/order"
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
	_          expParser
}

// func NewApp(logger Logger, ctx Context) *app {
func NewApp() *App {
	/*	a := new(app)
		a.ctx = ctx
		a.logger = logger
		return a*/
	a := new(App)
	a.downloader = downloader.NewDownloader("http://cetatenie.just.ro/ordine-articolul-11/", "./downloads")

	return a
}

func (a *App) Run() {
	a.downloader.Download()
}
