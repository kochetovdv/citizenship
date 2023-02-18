package downloader

import (

	//"citizenship/pkg/regulars"
	"citizenship/internal/order"
	"citizenship/pkg/osservices"
	"citizenship/pkg/webservices"
	"log"
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
	for _, order := range listOfOrders.Orders {
		// Check if the file already exists in the path
		filePath := filepath.Join(d.path, order.Filename)
		if _, err := os.Stat(filePath); !os.IsNotExist(err) {
			log.Printf("File %s already exists\n", order.Filename)
			ordersExist.Add(order)
			continue
			//	return nil
		}
		ordersToDownload.Add(order)
	}

	downloadedFiles := order.NewOrders()
	// Download the files
	for _, order := range ordersToDownload.Orders {
		filePath := filepath.Join(d.path, order.Filename)

		// Download the file from the URL
		log.Printf("Starting download file from %s\n", order.Link)
		response, err := webservices.GetResponse(order.Link)
		if err != nil {
			continue
			//return err
		}

		// Create a new file and save the response body to it
		log.Printf("Saving file %s\n", order.Filename)
		osservices.SaveToFile(filePath, response)
		if err != nil {
			continue
			//			return err
		}
		log.Printf("File %s downloaded to %s\n", order.Filename, filePath)
		// If file is downloaded, add it to the list of downloaded files
		downloadedFiles.Add(order)
	}
	downloadedFiles.Statistics()
	
	ordersExist.AddRange(downloadedFiles.Orders...)
	return ordersExist, nil
}
