package service

import (
	"deliver/internal/repository"
	"deliver/models"
	"deliver/pkg/logger"

	"google.golang.org/grpc/codes"
)

type CategoryService struct {
	repo repository.Repository
	log  logger.Logger
}

func NewCategoryService(repo repository.Repository, log logger.Logger) *CategoryService {
	return &CategoryService{
		repo: repo,
		log:  log,
	}
}

func (s *CategoryService) Create(category models.CategoryCreateRequest) (int64, error) {
	id, err := s.repo.Category.Create(category)
	if err != nil {
		return 0, serviceError(err, codes.Internal)
	}

	return id, nil
}

func (s *CategoryService) GetList(pagination *models.Pagination) ([]models.Category, error) {
	categories, err := s.repo.Category.GetList(pagination)
	if err != nil {
		return nil, serviceError(err, codes.Internal)
	}

	return categories, nil
}

func (s *CategoryService) GetById(id int64) (models.Category, error) {
	category, err := s.repo.Category.GetById(id)
	if err != nil {
		return models.Category{}, serviceError(err, codes.Internal)
	}

	return category, nil
}

func (s *CategoryService) Update(category models.CategoryUpdateRequest) error {
	err := s.repo.Category.Update(category)
	if err != nil {
		return serviceError(err, codes.Internal)
	}

	return nil
}

func (s *CategoryService) DeleteById(id int64) error {
	err := s.repo.Category.DeleteById(id)
	if err != nil {
		return serviceError(err, codes.Internal)
	}

	return nil
}
