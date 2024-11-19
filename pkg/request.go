package pkg

import (
	"io"
	"net/http"
)

func DoRequest(method string, url string) []byte {
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil
	}

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return nil
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil
	}

	return body
}
