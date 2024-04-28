//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen -package brec -generate types,client -o client.gen.go openapi.json

package blrec

import (
	"encoding/json"
	"net/http"
)

type BlrecClient struct {
	*Client
}

func NewBlrecClient(endpoint string) (*BlrecClient, error) {
	c, err := NewClient(endpoint)
	if err != nil {
		return nil, err
	}

	return &BlrecClient{
		Client: c,
	}, nil
}

func NewBlrecClientWithAuth(endpoint, apiKey string) (*BlrecClient, error) {
	doer := &httpDoerWithAuth{
		Client: http.DefaultClient,
		apiKey: apiKey,
	}

	c, err := NewClient(endpoint, WithHTTPClient(doer))
	if err != nil {
		return nil, err
	}

	return &BlrecClient{
		Client: c,
	}, nil
}

type httpDoerWithAuth struct {
	*http.Client
	apiKey string
}

func (d *httpDoerWithAuth) Do(req *http.Request) (*http.Response, error) {
	req.Header.Set("X-Api-Key", d.apiKey)
	return d.Client.Do(req)
}

func DecodeResponse(resp *http.Response, v interface{}) error {
	return json.NewDecoder(resp.Body).Decode(v)
}
