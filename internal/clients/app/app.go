package app

import(
	"citizenship/internal/downloader"
)

type Downloader interface {
	Download(url string)
}

type Parser interface {
	Parse(pdf string)
}

type app struct {
//	logger Logger
//	ctx    Context
	downloader Downloader
	_ Parser
}

// func NewApp(logger Logger, ctx Context) *app {
func NewApp() *app {
/*	a := new(app)
	a.ctx = ctx
	a.logger = logger
	return a*/
	a:=new(app)
	a.downloader=downloader.NewDownloader()
	return a
}

func (a *app) Run() {
	a.downloader.Download("http://cetatenie.just.ro/ordine-articolul-11/")
}