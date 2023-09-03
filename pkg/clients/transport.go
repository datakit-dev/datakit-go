package clients

import (
	"net/http"
)

type Transport struct {
	T        http.RoundTripper
	apiToken string
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("auth-provider", "firebase")
	req.Header.Set("authorization", "Bearer "+t.apiToken)
	return t.T.RoundTrip(req)
}

func NewTransport(apiToken string) *Transport {
	return &Transport{
		T:        http.DefaultTransport,
		apiToken: apiToken,
	}
}
