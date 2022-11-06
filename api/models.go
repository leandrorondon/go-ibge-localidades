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

// RegiaoIntermediaria representa uma Região Intermediária.
type RegiaoIntermediaria struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
	UF   UF     `json:"UF"`
}

// RegiaoImediata representa uma Região Imediata.
type RegiaoImediata struct {
	ID                  int                 `json:"id"`
	Nome                string              `json:"nome"`
	RegiaoIntermediaria RegiaoIntermediaria `json:"regiao-intermediaria"`
}

// Mesorregiao representa uma Mesorregião.
type Mesorregiao struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
	UF   UF     `json:"UF"`
}

// Microrregiao representa uma Microrregião.
type Microrregiao struct {
	ID          int         `json:"id"`
	Nome        string      `json:"nome"`
	Mesorregiao Mesorregiao `json:"mesorregiao"`
}

// Municipio representa um Município.
type Municipio struct {
	ID             int            `json:"id"`
	Nome           string         `json:"nome"`
	Microrregiao   Microrregiao   `json:"microrregiao"`
	RegiaoImediata RegiaoImediata `json:"regiao-imediata"`
}

// Distrito representa um Distrito.
type Distrito struct {
	ID        int       `json:"id"`
	Nome      string    `json:"nome"`
	Municipio Municipio `json:"municipio"`
}
