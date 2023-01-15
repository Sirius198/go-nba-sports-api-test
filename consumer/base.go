package consumer

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

// Get JSON data
func Get(apiUrl string) ([]byte, error) {

	baseUrl := os.Getenv("BASE_URL")
	apiKey := os.Getenv("API_SUBSCRIPTION_KEY")

	apiUrl, _ = url.JoinPath(baseUrl, apiUrl)
	apiUrl = fmt.Sprintf("%s?key=%s", apiUrl, apiKey)

	resp, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func Post() {}
