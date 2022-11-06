package api

import (
	"context"
	"encoding/json"
	"fmt"
)

// UFs

// UFsResponse represents the response of the UFs methods.
type UFsResponse []UF

type UFs struct {
	client httpClient
}

// UFs List all the UFs.
// Returns a list of UFs.
// GET https://servicodados.ibge.gov.br/api/v1/localidades/estados
func (ufs *UFs) UFs(ctx context.Context) (*UFsResponse, error) {
	if ufs.client == nil {
		return nil, ErrHTTPClientNotSet
	}

	url := fmt.Sprintf("%s/estados", baseURL)

	b, err := ufs.client.Get(ctx, url)
	if err != nil {
		return nil, err
	}

	var r UFsResponse
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
