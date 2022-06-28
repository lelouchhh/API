package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"AGZ/pkg/structures"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}
func (r *AuthPostgres) SetReg(user structures.User) (error) {
	query := fmt.Sprintf("call auth.set_reg('%s', '%s')", user.Name, user.Password)
	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}

	return err
}
func (r *AuthPostgres) AuthenticationUser(user structures.User) (structures.Tokens, error) {
	query := fmt.Sprintf("select * from auth.authentication_user('%s', '%s', '%s')", user.Name, user.Password, user.DeviceId)
	var rows structures.Tokens
	fmt.Println(query)
	err := r.db.Get(&rows, query)
	if err != nil {
		return rows, err
	}

	return rows, err
}