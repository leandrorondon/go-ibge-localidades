package ibgelocalidades

import (
	"net/http"
	"time"

	"github.com/leandrorondon/go-ibge-localidades/api"
	"github.com/leandrorondon/go-ibge-localidades/internal/httpclient"
)

type options []httpclient.ClientOption

// Option describes the type for functional options.
type Option func(*options)

// WithTransport is the option to set a custom transport layer to the API client.
func WithTransport(rt http.RoundTripper) Option {
	return func(o *options) {
		*o = append(*o, httpclient.WithTransport(rt))
	}
}

// WithTimeout is the option to set a timeout to API requests.
func WithTimeout(timeout time.Duration) Option {
	return func(o *options) {
		*o = append(*o, httpclient.WithTimeout(timeout))
	}
}

// New creates a new API client.
func New(opts ...Option) *api.API {
	var clientOpts options
	for _, opt := range opts {
		opt(&clientOpts)
	}
	client := httpclient.New(clientOpts...)

	return api.New(client)
}
