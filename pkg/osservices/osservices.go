package osservices

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// Check if the path already exists, Return true if it exists, false if it does not
func CheckDir(path string) (bool, error) {
	// Create the directory if it does not exist
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return false, fmt.Errorf("error with creating folder:%s", path)
		}
	}
	return true, nil
}

func SaveToFile(path string, data []byte) error {
	// Create the directory if it does not exist
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error with creating file:%s", path)
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("error with writing file:%s", path)
	}
	return nil
}

func DownloadFile(path, fileName, url string) error {
	// Create the directory if it does not exist
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}

	// Check if the file already exists in the path
	filePath := filepath.Join(path, fileName)
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		log.Printf("File %s already exists\n", fileName)
		return nil
	}

	// Download the file from the URL
	log.Printf("Starting download file from %s\n", url)
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Create a new file and save the response body to it
	log.Printf("Saving file %s\n", fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	log.Printf("File %s downloaded to %s\n", fileName, filePath)
	return nil
}
