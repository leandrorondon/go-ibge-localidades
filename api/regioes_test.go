package api_test

import (
	"context"
	"errors"
	"testing"

	"github.com/leandrorondon/go-ibge-localidades/api"
	"github.com/stretchr/testify/assert"
)

func TestRegioes(t *testing.T) {
	tests := []struct {
		name             string
		httpClient       httpClient
		expectedError    string
		expectedResponse *api.RegioesResponse
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
				GetResponse: []byte(`[{"id":5,"sigla":"CO","nome":"Centro-Oeste"}]`),
			},
			expectedResponse: &api.RegioesResponse{
				{
					ID:    5,
					Sigla: "CO",
					Nome:  "Centro-Oeste",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			apiClient := api.New(test.httpClient)
			r, err := apiClient.Regioes.Regioes(context.Background())
			if test.expectedError != "" {
				assert.ErrorContains(t, err, test.expectedError)
			} else {
				assert.NoError(t, err)
			}
			assert.EqualValues(t, test.expectedResponse, r)
		})
	}
}
