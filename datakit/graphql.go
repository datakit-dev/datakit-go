package datakit

import (
	"net/http"

	"github.com/hasura/go-graphql-client"
)

const GRAPHQL_PATH = "/graphql"

type GQLClient struct {
	Client *graphql.Client
}

func NewGQLClient(apiUrl string, apiToken string) *GQLClient {
	httpClient := &http.Client{
		Transport: NewTransport(apiToken),
	}
	return &GQLClient{
		Client: graphql.NewClient(apiUrl+GRAPHQL_PATH, httpClient),
	}
}
