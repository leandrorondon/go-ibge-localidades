package api_test

import (
	"context"
	"errors"
	"testing"

	"github.com/leandrorondon/go-ibge-localidades/api"
	"github.com/stretchr/testify/assert"
)

func TestUFs(t *testing.T) {
	tests := []struct {
		name             string
		httpClient       httpClient
		expectedError    string
		expectedResponse *api.UFsResponse
	}{
		{
			name:          "invalid http client",
			expectedError: api.ErrHTTPClientNotSet.Error(),
		},
		{
			name: "get error",
			httpClient: &httpClientMock{
				GetError: errors.New("get error"),
			},
			expectedError: "get error",
		},
		{
			name: "invalid response",
			httpClient: &httpClientMock{
				GetResponse: []byte(`{`),
			},
			expectedError: "unexpected end of JSON input",
		},
		{
			name: "success",
			httpClient: &httpClientMock{
				GetResponse: []byte(`[{"id":1,"sigla":"AA","nome":"Estado A","regiao":{"id":1,"sigla":"A","nome":"Região A"}},{"id":2,"sigla":"BB","nome":"Estado B","regiao":{"id":2,"sigla":"B","nome":"Região B"}}]`),
			},
			expectedResponse: &api.UFsResponse{
				{ID: 1, Nome: "Estado A", Sigla: "AA", Regiao: api.Regiao{ID: 1, Nome: "Região A", Sigla: "A"}},
				{ID: 2, Nome: "Estado B", Sigla: "BB", Regiao: api.Regiao{ID: 2, Nome: "Região B", Sigla: "B"}},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			apiClient := api.New(test.httpClient)
			r, err := apiClient.UFs.UFs(context.Background())
			if test.expectedError != "" {
				assert.ErrorContains(t, err, test.expectedError)
			} else {
				assert.NoError(t, err)
			}
			assert.EqualValues(t, test.expectedResponse, r)
		})
	}
}
