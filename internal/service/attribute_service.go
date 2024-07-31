package service

import (
	"deliver/internal/repository"
	"deliver/models"
	"deliver/pkg/logger"

	"google.golang.org/grpc/codes"
)

type AttributeService struct {
	repo repository.Repository
	log  logger.Logger
}

func NewAttributeService(repo repository.Repository, log logger.Logger) *AttributeService {
	return &AttributeService{
		repo: repo,
		log:  log,
	}
}

func (s *AttributeService) Create(attribute models.AttributeCreateRequest) (int64, error) {
	id, err := s.repo.Attribute.Create(attribute)
	if err != nil {
		return 0, serviceError(err, codes.Internal)
	}

	return id, nil
}

func (s *AttributeService) GetList(pagination *models.Pagination) ([]models.Attribute, error) {
	attributes, err := s.repo.Attribute.GetList(pagination)
	if err != nil {
		return nil, serviceError(err, codes.Internal)
	}

	return attributes, nil
}

func (s *AttributeService) GetById(id int64) (models.Attribute, error) {
	attribute, err := s.repo.Attribute.GetById(id)
	if err != nil {
		return models.Attribute{}, serviceError(err, codes.Internal)
	}

	return attribute, nil
}

func (s *AttributeService) Update(attribute models.AttributeUpdateRequest) error {
	err := s.repo.Attribute.Update(attribute)
	if err != nil {
		return serviceError(err, codes.Internal)
	}

	return nil
}

func (s *AttributeService) DeleteById(id int64) error {
	err := s.repo.Attribute.DeleteById(id)
	if err != nil {
		return serviceError(err, codes.Internal)
	}

	return nil
}
