package client

import (
	"fmt"
	"net/http"
	"os"

	"github.com/cli/go-gh/v2/pkg/auth"
)

func detectAuth(opts ClientOptions) (token, host string, err error) {
	if opts.Token != "" {
		token = opts.Token
	} else {
		if token = os.Getenv("GH_TOKEN"); token == "" {
			token = os.Getenv("GITHUB_TOKEN")
		}
	}

	if opts.Host != "" {
		host = opts.Host
	} else {
		if host = os.Getenv("GH_HOST"); host == "" {
			host, _ = auth.DefaultHost()
		}
	}

	if token == "" {
		token, _ = auth.TokenForHost(host)
		if token == "" {
			return "", "", fmt.Errorf("no authentication token found")
		}
	}

	return token, host, nil
}

type authTransport struct {
	token   string
	baseURL string
	base    http.RoundTripper
}

func (t *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.base == nil {
		t.base = http.DefaultTransport
	}
	req = req.Clone(req.Context())
	req.Header.Set("Authorization", "token "+t.token)
	return t.base.RoundTrip(req)
}
