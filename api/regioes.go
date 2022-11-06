package api

import (
	"context"
	"encoding/json"
	"fmt"
)

// Regioes

// RegioesResponse represents the response of the Regioes methods.
type RegioesResponse []Regiao

type Regioes struct {
	client httpClient
}

// Regioes List all the Municípios.
// Returns a list of Regioes.
// GET https://servicodados.ibge.gov.br/api/v1/localidades/regioes
func (d *Regioes) Regioes(ctx context.Context) (*RegioesResponse, error) {
	if d.client == nil {
		return nil, ErrHTTPClientNotSet
	}

	url := fmt.Sprintf("%s/regioes", baseURL)

	b, err := d.client.Get(ctx, url)
	if err != nil {
		return nil, err
	}

	var r RegioesResponse
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
