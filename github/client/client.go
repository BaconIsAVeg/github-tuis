package client

import (
	"net/http"

	"github.com/cli/go-gh/v2/pkg/api"
	"github.com/google/go-github/v82/github"
)

type BaseClient struct {
	client        *github.Client
	graphqlClient *api.GraphQLClient
}

type ClientOptions struct {
	Token string
	Host  string
}

func NewClient(opts ClientOptions) (*BaseClient, error) {
	token, host, err := detectAuth(opts)
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{}
	if token != "" {
		httpClient.Transport = &authTransport{
			token:   token,
			baseURL: host,
		}
	}

	client := github.NewClient(httpClient)
	if host != "" && host != "github.com" {
		var err error
		client, err = github.NewClient(httpClient).WithEnterpriseURLs(host, host)
		if err != nil {
			return nil, err
		}
	}

	var graphqlClient *api.GraphQLClient
	var gqlErr error
	if token != "" {
		graphqlClient, gqlErr = api.NewGraphQLClient(api.ClientOptions{
			AuthToken:    token,
			LogIgnoreEnv: true,
		})
	} else if host == "github.com" {
		graphqlClient, gqlErr = api.DefaultGraphQLClient()
	} else {
		graphqlClient, gqlErr = api.NewGraphQLClient(api.ClientOptions{
			Host:         host,
			LogIgnoreEnv: true,
		})
	}
	if gqlErr != nil {
		graphqlClient = nil
	}

	return &BaseClient{
		client:        client,
		graphqlClient: graphqlClient,
	}, nil
}

func (c *BaseClient) REST() *github.Client {
	return c.client
}

func (c *BaseClient) GraphQL() *api.GraphQLClient {
	return c.graphqlClient
}

func (c *BaseClient) Close() {}
