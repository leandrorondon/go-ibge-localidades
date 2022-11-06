package api_test

import "context"

type httpClient interface {
	Get(ctx context.Context, url string) ([]byte, error)
	Post(ctx context.Context, url string, body any) ([]byte, error)
}

type httpClientMock struct {
	GetResponse []byte
	GetError    error
	SpyGetURL   string

	PostResponse []byte
	PostError    error
	SpyPostURL   string
	SpyPostBody  any
}

func (m *httpClientMock) Get(_ context.Context, url string) ([]byte, error) {
	m.SpyGetURL = url

	return m.GetResponse, m.GetError
}

func (m *httpClientMock) Post(_ context.Context, url string, body any) ([]byte, error) {
	m.SpyPostURL = url
	m.SpyPostBody = body

	return nil, nil
}
