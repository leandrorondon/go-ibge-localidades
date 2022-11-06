package api

import (
	"context"
	"errors"
)

const (
	baseURL = "https://servicodados.ibge.gov.br/api/v1/localidades"
)

var (
	// ErrHTTPClientNotSet is returned when a method is called and the http client is not configured.
	ErrHTTPClientNotSet = errors.New("http client not set")
)

// API is the IBGE Localidades API object.
type API struct {
	UFs        UFs
	Distritos  Distritos
	Municipios Municipios
}

type httpClient interface {
	Get(ctx context.Context, url string) ([]byte, error)
	Post(ctx context.Context, url string, body any) ([]byte, error)
}

func New(client httpClient) *API {
	return &API{
		UFs:        UFs{client: client},
		Distritos:  Distritos{client: client},
		Municipios: Municipios{client: client},
	}
}
