package service
import
(
	"AGZ/pkg/repository"
	"AGZ/pkg/structures"
)

type AuthService struct {
	repo repository.Authorization
}
func (s *AuthService) SetReg(user structures.User) error {
	return s.repo.SetReg(user)
}
func (s *AuthService) AuthenticationUser(user structures.User) (structures.Tokens, error) {
	return s.repo.AuthenticationUser(user)
}
func (s *AuthService) ActionUserAccess(token structures.Tokens) error{
	return s.repo.ActionUserAccess(token)
}
func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}
func (s *AuthService) ActionUserRefresh(token structures.Tokens) (structures.Tokens, error) {
	return s.repo.ActionUserRefresh(token)
}