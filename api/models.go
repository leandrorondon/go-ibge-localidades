package api

// Regiao representa uma Região do país.
type Regiao struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Sigla string `json:"sigla"`
}

// UF representa uma Unidade da Federação.
type UF struct {
	ID     int    `json:"id"`
	Nome   string `json:"nome"`
	Sigla  string `json:"sigla"`
	Regiao Regiao `json:"regiao"`
}
