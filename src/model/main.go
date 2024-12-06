package model

import "database/sql"

type Model struct {
	Grupo GrupoModel
}

func New(db *sql.DB) *Model {
	return &Model{
		Grupo: GrupoModel{DB: db},
	}
}
