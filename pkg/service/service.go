package service

import (
	"AGZ/pkg/repository"
	"AGZ/pkg/structures"
)


type Authorization interface {
	SetReg(user structures.User) (error)
	AuthenticationUser(user structures.User)(structures.Tokens, error)
	ActionUserAccess(token structures.Tokens)(err error)
	ActionUserRefresh(token structures.Tokens)(structures.Tokens, error)
}
type Procedures interface {
	CreateNews(new structures.New) error
	GetEntity() ([]structures.Entity, error)
}
type Service struct {
	Authorization
	Procedures
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Procedures:    NewProceduresService(repos.Procedures),
	}
}