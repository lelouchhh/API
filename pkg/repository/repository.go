package repository

import (
	"AGZ/pkg/structures"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	SetReg(user structures.User) (error)
	AuthenticationUser(user structures.User) (structures.Tokens, error)
}
type Procedures interface {
	CreateNews(new structures.New) error
	GetEntity() ([]structures.Entity, error)
}
type Repository struct {
	Authorization
	Procedures
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Procedures: NewProceduresPostgres(db),
	}
}