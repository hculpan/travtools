package util

import (
	"io"
	"log"
	"net/http"
)

func FetchData(dest string) ([]byte, error) {
	log.Printf("Fetching data from %s\n", dest)
	resp, err := http.Get(dest)
	if err != nil {
		return []byte{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}
