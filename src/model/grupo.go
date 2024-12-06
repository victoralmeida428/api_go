package model

import "database/sql"

type Grupo struct {
	ID    int64          `json:"id"`
	Grupo sql.NullString `json:"grupo"`
}

type GrupoModel struct {
	DB *sql.DB
}
