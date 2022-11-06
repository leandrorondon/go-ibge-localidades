package api

import (
	"context"
	"encoding/json"
	"fmt"
)

// Distritos

// DistritosResponse represents the response of the Distritos methods.
type DistritosResponse []Distrito

type Distritos struct {
	client httpClient
}

// Distritos List all the Distritos.
// Returns a list of Distritos.
// GET https://servicodados.ibge.gov.br/api/v1/localidades/distritos
func (d *Distritos) Distritos(ctx context.Context) (*DistritosResponse, error) {
	if d.client == nil {
		return nil, ErrHTTPClientNotSet
	}

	url := fmt.Sprintf("%s/distritos", baseURL)

	b, err := d.client.Get(ctx, url)
	if err != nil {
		return nil, err
	}

	var r DistritosResponse
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
