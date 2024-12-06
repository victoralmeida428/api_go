package repository

import (
	"api/src/model"
	"context"
	"database/sql"
	"time"
)

type Grupo struct {
	db *sql.DB
}

func (g *Grupo) FindAll() ([]model.Grupo, error) {
	query := `
	select id, grupo from ep_dw.grupo
`
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rows, err := g.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []model.Grupo
	for rows.Next() {
		var result model.Grupo
		if err = rows.Scan(&result.ID, &result.Grupo); err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}

func (g *Grupo) FindById(id int) (model.Grupo, error) {
	var grupo model.Grupo
	return grupo, nil
}

func (g *Grupo) FindByField(field string, value interface{}) ([]model.Grupo, error) { return nil, nil }

func (g *Grupo) Insert(grupo model.Grupo) (model.Grupo, error) {
	return grupo, nil
}
func (g *Grupo) Update(grupo model.Grupo) (model.Grupo, error) {
	return grupo, nil
}
func (g *Grupo) Delete(grupo model.Grupo) (model.Grupo, error) {
	return grupo, nil
}
