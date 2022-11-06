package api_test

import (
	"context"
	"errors"
	"testing"

	"github.com/leandrorondon/go-ibge-localidades/api"
	"github.com/stretchr/testify/assert"
)

func TestMunicipios(t *testing.T) {
	tests := []struct {
		name             string
		httpClient       httpClient
		expectedError    string
		expectedResponse *api.MunicipiosResponse
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
				GetResponse: []byte(`[{"id":5300108,"nome":"Brasília","microrregiao":{"id":53001,"nome":"Brasília","mesorregiao":{"id":5301,"nome":"Distrito Federal","UF":{"id":53,"sigla":"DF","nome":"Distrito Federal","regiao":{"id":5,"sigla":"CO","nome":"Centro-Oeste"}}}},"regiao-imediata":{"id":530001,"nome":"Distrito Federal","regiao-intermediaria":{"id":5301,"nome":"Distrito Federal","UF":{"id":53,"sigla":"DF","nome":"Distrito Federal","regiao":{"id":5,"sigla":"CO","nome":"Centro-Oeste"}}}}}]`),
			},
			expectedResponse: &api.MunicipiosResponse{
				{
					ID:   5300108,
					Nome: "Brasília",
					Microrregiao: api.Microrregiao{
						ID:   53001,
						Nome: "Brasília",
						Mesorregiao: api.Mesorregiao{
							ID:   5301,
							Nome: "Distrito Federal",
							UF: api.UF{
								ID:    53,
								Nome:  "Distrito Federal",
								Sigla: "DF",
								Regiao: api.Regiao{
									ID:    5,
									Sigla: "CO",
									Nome:  "Centro-Oeste",
								},
							},
						},
					},
					RegiaoImediata: api.RegiaoImediata{
						ID:   530001,
						Nome: "Distrito Federal",
						RegiaoIntermediaria: api.RegiaoIntermediaria{
							ID:   5301,
							Nome: "Distrito Federal",
							UF: api.UF{
								ID:    53,
								Nome:  "Distrito Federal",
								Sigla: "DF",
								Regiao: api.Regiao{
									ID:    5,
									Sigla: "CO",
									Nome:  "Centro-Oeste",
								},
							},
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			apiClient := api.New(test.httpClient)
			r, err := apiClient.Municipios.Municipios(context.Background())
			if test.expectedError != "" {
				assert.ErrorContains(t, err, test.expectedError)
			} else {
				assert.NoError(t, err)
			}
			assert.EqualValues(t, test.expectedResponse, r)
		})
	}
}
