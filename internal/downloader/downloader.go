package downloader

import (

	//"citizenship/pkg/regulars"
	"citizenship/internal/order"
	"citizenship/pkg/osservices"
	"citizenship/pkg/webservices"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/ledongthuc/pdf"
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
// Download downloads a files from a url and returns list of downloaded files
func (d *Downloader) Download(listOfOrders *order.Orders) (*order.Orders, error) {
	// Create the directory if it does not exist
	isExist, err := osservices.CheckDir(d.path)
	if err != nil {
		log.Printf("Error with creating folder:%s\n", d.path)
		return nil, err
	}
	if !isExist {
		log.Printf("Error with checking folder:%s\n", d.path)
		return nil, err
	}

	ordersExist := order.NewOrders()
	ordersToDownload := order.NewOrders()
	// Checking files in the directory that are already downloaded
	for filename, order := range listOfOrders.Orders {
		// Check if the file already exists in the path
		filePath := filepath.Join(d.path, filename)
		if _, err := os.Stat(filePath); !os.IsNotExist(err) {
			//			log.Printf("File %s already exists\n", order.Filename)
			ordersExist.Add(filename, order)
			continue
		}
		ordersToDownload.Add(filename, order)
	}

	downloadedFiles := order.NewOrders()
	filesForRedownload := order.NewOrders()
	// Download the files
	for filename, order := range ordersToDownload.Orders {
		err := d.download(filename, order)
		if err != nil {
			log.Printf("Error with downloading file: %s\n", err)
			continue
		}
		err = d.ckeckingFile(filename)
		if err != nil {
			log.Printf("Error with checking file: %s\n", err)
			filesForRedownload.Add(filename, order)
			continue
		}
		// If file is downloaded, add it to the list of downloaded files
		downloadedFiles.Add(filename, order)
	}

	for filename, order := range downloadedFiles.Orders {
		ordersExist.Add(filename, order)
	}
	return ordersExist, nil
}


//TODO How to check that downloaded file is PDF and not just html response?
func (d *Downloader) download(filename string, order order.Order) error {
	filePath := filepath.Join(d.path, filename)
	// Download the file from the URL
	log.Printf("Starting download file from %s\n", order.Link)
	response, err := webservices.GetResponse(order.Link)
	if err != nil {
		return fmt.Errorf("error with getting response: %s", err)
	}

	// Create a new file and save the response body to it
	log.Printf("Saving file %s\n", filename)
	err = osservices.SaveToFile(filePath, response)
	if err != nil {
		return fmt.Errorf("error with saving file: %s", err)
	}
	log.Printf("File %s downloaded to %s\n", filename, d.path)
	return nil
}

func (d *Downloader) ckeckingFile(filename string) error {
	_, _, err := pdf.Open(d.path + "/" + filename)
	if err != nil {
		fmt.Printf("error in opening file: %s\n", err)
		return err
	}
	return nil
}
