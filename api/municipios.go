package api

import (
	"context"
	"encoding/json"
	"fmt"
)

// Municipios

// MunicipiosResponse represents the response of the Municipios methods.
type MunicipiosResponse []Municipio

type Municipios struct {
	client httpClient
}

// Municipios List all the Municípios.
// Returns a list of Municipios.
// GET https://servicodados.ibge.gov.br/api/v1/localidades/municipios
func (d *Municipios) Municipios(ctx context.Context) (MunicipiosResponse, error) {
	if d.client == nil {
		return nil, ErrHTTPClientNotSet
	}

	url := fmt.Sprintf("%s/municipios", baseURL)

	b, err := d.client.Get(ctx, url)
	if err != nil {
		return nil, err
	}

	var r MunicipiosResponse
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
