package service

import (
	"deliver/internal/repository"
	"deliver/models"
	"deliver/pkg/logger"

	"google.golang.org/grpc/codes"
)

type ProductService struct {
	repo repository.Repository
	log  logger.Logger
}

func NewProductService(repo repository.Repository, log logger.Logger) *ProductService {
	return &ProductService{
		repo: repo,
		log:  log,
	}
}

func (s *ProductService) Create(product models.ProductCreateRequest) (int64, error) {
	id, err := s.repo.Product.Create(product)
	if err != nil {
		return 0, serviceError(err, codes.Internal)
	}

	return id, nil
}

func (s *ProductService) GetList(pagination *models.Pagination) ([]models.Product, error) {
	products, err := s.repo.Product.GetList(pagination)
	if err != nil {
		return nil, serviceError(err, codes.Internal)
	}

	return products, nil
}

func (s *ProductService) GetById(id int64) (models.Product, error) {
	product, err := s.repo.Product.GetById(id)
	if err != nil {
		return models.Product{}, serviceError(err, codes.Internal)
	}

	return product, nil
}

func (s *ProductService) Update(product models.ProductUpdateRequest) error {
	err := s.repo.Product.Update(product)
	if err != nil {
		return serviceError(err, codes.Internal)
	}

	return nil
}

func (s *ProductService) DeleteById(id int64) error {
	err := s.repo.Product.DeleteById(id)
	if err != nil {
		return serviceError(err, codes.Internal)
	}

	return nil
}
