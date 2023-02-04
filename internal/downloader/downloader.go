package downloader

import (
	_ "net/http"
	_ "fmt"
	_ "log"
)

type Downloader struct {
	
}

func NewDownloader() *Downloader {
	return &Downloader{}
}


func (d *Downloader) Download(url string) {
	
}