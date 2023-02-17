package downloader

import (

	//"citizenship/pkg/regulars"
	"citizenship/internal/order"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// Downloader is an structure for downloading files from a links
type Downloader struct {
	path string
}

// NewDownloader creates a new Downloader
func NewDownloader(path string) *Downloader {
	d := new(Downloader)
	d.path = path
	return d
}

// TODO Error handling
func (d *Downloader) Download(listOfOrders *order.Orders) error {
	//todo connect
	// todo download file

	for _, order := range listOfOrders.Orders {
		// Create the directory if it does not exist
		if _, err := os.Stat(d.path); os.IsNotExist(err) {
			err = os.MkdirAll(d.path, 0755)
			if err != nil {
				continue
				//	return err
			}
		}

		// Check if the file already exists in the path
		filePath := filepath.Join(d.path, order.Filename)
		if _, err := os.Stat(filePath); !os.IsNotExist(err) {
			log.Printf("File %s already exists\n", order.Filename)
			continue
			//	return nil
		}

		// Download the file from the URL
		log.Printf("Starting download file from %s\n", order.Link)
		response, err := http.Get(order.Link)
		if err != nil {
			continue
			//return err
		}
		defer response.Body.Close()

		// Create a new file and save the response body to it
		log.Printf("Saving file %s\n", order.Filename)
		file, err := os.Create(filePath)
		if err != nil {
			continue
			//			return err
		}
		defer file.Close()

		_, err = io.Copy(file, response.Body)
		if err != nil {
			continue
			//			return err
		}
		log.Printf("File %s downloaded to %s\n", order.Filename, filePath)
	}
	return nil
}
