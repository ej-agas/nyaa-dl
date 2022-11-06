package http

import (
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/ej-agas/nyaa-dl/domain"
)

type Nyaa struct {
	baseUrl string
}

func CreateNyaa(baseUrl string) *Nyaa {
	return &Nyaa{baseUrl: baseUrl}
}

func (n *Nyaa) Search(q domain.Query) ([]byte, error) {
	baseUrl, err := url.Parse(n.baseUrl)

	if err != nil {
		return nil, err
	}

	baseUrl.RawQuery = q.UrlEncode()
	resp, err := http.Get(baseUrl.String())

	if err != nil {
		return nil, err
	}

	return handleResponse(resp)
}

func handleResponse(r *http.Response) ([]byte, error) {
	bodyBytes, err := io.ReadAll(r.Body)

	defer r.Body.Close()

	if err != nil {
		return bodyBytes, err
	}

	if r.StatusCode != http.StatusOK {
		err := errors.New(string(bodyBytes))
		return bodyBytes, err
	}

	return bodyBytes, nil
}
