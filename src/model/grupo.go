package model

import (
	"database/sql"
	"encoding/json"
)

type Grupo struct {
	ID    int64          `json:"id"`
	Grupo sql.NullString `json:"grupo"`
}

func (e *Grupo) MarshalJSON() ([]byte, error) {

	type GrupoAlias struct {
		ID    int64  `json:"id"`
		Grupo string `json:"grupo"`
	}

	aux := GrupoAlias{
		ID:    e.ID,
		Grupo: e.Grupo.String,
	}

	return json.Marshal(aux)
}

type GrupoModel struct {
	DB *sql.DB
}
