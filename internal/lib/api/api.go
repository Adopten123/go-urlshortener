package api

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrInvalidStatusCode = errors.New("invalid status code")
)

func GetRedirect(url string) (string, error) {
	const op = "api.GetRedirect"

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	response, err := client.Get(url)
	if err != nil {
		return "", err
	}

	defer func() { _ = response.Body.Close() }()
	if response.StatusCode != http.StatusFound {
		return "", fmt.Errorf("%s: %w: %d", op, ErrInvalidStatusCode, response.StatusCode)
	}
	return response.Header.Get("Location"), nil
}
