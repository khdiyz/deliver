package service

import (
	"deliver/internal/repository"
	"deliver/models"
	"deliver/pkg/logger"

	"google.golang.org/grpc/codes"
)

type RoleService struct {
	repo repository.Repository
	log  logger.Logger
}

func NewRoleService(repo repository.Repository, log logger.Logger) *RoleService {
	return &RoleService{
		repo: repo,
		log:  log,
	}
}

func (s *RoleService) GetList(pagination *models.Pagination) ([]models.Role, error) {
	roles, err := s.repo.Role.GetList(pagination)
	if err != nil {
		return nil, serviceError(err, codes.Internal)
	}

	return roles, nil
}
