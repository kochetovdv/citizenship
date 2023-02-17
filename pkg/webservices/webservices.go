package webservices

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetResponse(url string) ([]byte, error) {
	// TODO after refactoring, change to use pkg/web-services
	// response, err:= connect(s.url)
	log.Printf("Connecting: %s", url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error with connecting: %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()

	if resp.StatusCode > 299 {
		log.Printf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
	}

	if err != nil {
		return nil, fmt.Errorf("error with connecting: %s", err)
	}

	return body, nil
}
