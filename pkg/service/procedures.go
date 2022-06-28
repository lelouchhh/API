package service

import (
	"AGZ/pkg/repository"
	"AGZ/pkg/structures"
)

type ProceduresService struct {
	repo repository.Procedures
}
func (s *ProceduresService) CreateNews(new structures.New) (error) {
	return s.repo.CreateNews(new)
}
func (s *ProceduresService) GetEntity() ([]structures.Entity, error) {
	return s.repo.GetEntity()
}


func NewProceduresService(repo repository.Procedures) *ProceduresService {
	return &ProceduresService{repo: repo}
}