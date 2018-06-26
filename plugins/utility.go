package utility

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// SendPublicRequest sends a public request
func SendPublicRequest(method string, url string, body io.Reader) (int, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, body)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("error: %s", err)
	}

	if status := resp.StatusCode; status > 400 {
		err := fmt.Errorf("%v response when requesting user info", status)
		return status, err
	}
	return 0, nil
}
