package repository

import (
	"AGZ/pkg/structures"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ProceduresPostgres struct {
	db *sqlx.DB
}

func NewProceduresPostgres(db *sqlx.DB) *ProceduresPostgres {
	return &ProceduresPostgres{db: db}
}
func (r *ProceduresPostgres) CreateNews(new structures.New) error {
	query := fmt.Sprintf("call news.create_news('%s', '%s', '%s', '%s', '%s')", new.Title, new.Description, new.Picture, new.TypeTags, new.Tags)
	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}

	return err
}
func (r *ProceduresPostgres) GetEntity() ([]structures.Entity, error) {
	query := fmt.Sprintf("select * from purchase.give_entity()")
	var res []structures.Entity
	err := r.db.Select(&res, query)
	if err != nil {
		return []structures.Entity{}, err
	}

	return res, err
}