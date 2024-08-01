package service

import (
	"database/sql"
	"deliver/config"
	"deliver/internal/constants"
	"deliver/internal/models"
	"deliver/internal/repository"
	"deliver/internal/storage"
	"deliver/pkg/logger"
	"errors"

	"google.golang.org/grpc/codes"
)

type ProductService struct {
	repo repository.Repository
	log  logger.Logger
	cfg  *config.Config
}

func NewProductService(repo repository.Repository, log logger.Logger, cfg config.Config) *ProductService {
	return &ProductService{
		repo: repo,
		log:  log,
		cfg:  &cfg,
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

	for i := range products {
		products[i].Photo = storage.GenerateLink(s.cfg, products[i].Photo)
	}

	return products, nil
}

func (s *ProductService) GetById(id int64) (models.Product, error) {
	product, err := s.repo.Product.GetById(id)
	if err != nil {
		return models.Product{}, serviceError(err, codes.Internal)
	}

	product.Photo = storage.GenerateLink(s.cfg, product.Photo)

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

func (s *ProductService) AddAttributeToProduct(productId, attributeId int64) error {
	_, err := s.repo.Product.GetById(productId)
	if err != nil {
		return serviceError(err, codes.NotFound)
	}

	_, err = s.repo.Attribute.GetById(attributeId)
	if err != nil {
		return serviceError(err, codes.NotFound)
	}

	_, err = s.repo.ProductAttribute.GetByProductIdAndAttributeId(productId, attributeId)
	if err == nil {
		return serviceError(errors.New("already exists"), codes.InvalidArgument)
	} else if err != sql.ErrNoRows {
		return serviceError(err, codes.Internal)
	}

	_, err = s.repo.ProductAttribute.Create(models.AddAttributeToProduct{
		ProductId:   productId,
		AttributeId: attributeId,
	})

	if err != nil {
		return serviceError(err, codes.Internal)
	}

	return nil
}

func (s *ProductService) RemoveAttributeFromProduct(productId, attributeId int64) error {
	_, err := s.repo.ProductAttribute.GetByProductIdAndAttributeId(productId, attributeId)
	if err != nil {
		return serviceError(constants.ErrDataIsEmpty, codes.InvalidArgument)
	}

	err = s.repo.ProductAttribute.DeleteByProductIdAndAttributeId(productId, attributeId)
	if err != nil {
		return serviceError(err, codes.Internal)
	}

	return nil
}

func (s *ProductService) AddToCart(userId int64, request models.CartProductCreateRequest) error {
	var err error

	request.CartId, err = s.repo.Cart.GetCartIdByUserId(userId)
	if err != nil {
		return serviceError(err, codes.Internal)
	}

	_, err = s.repo.Product.GetById(request.ProductId)
	if err != nil {
		return serviceError(err, codes.InvalidArgument)
	}

	_, err = s.repo.Cart.CreateCartProduct(request)
	if err != nil {
		return serviceError(err, codes.Internal)
	}

	// for i := range options {
	// 	attribute, err := s.repo.Attribute.GetById(options[i].AttributeId)
	// 	if err != nil {
	// 		return serviceError(err, codes.InvalidArgument)
	// 	}

	// 	attributeOptionIds := []int64{}
	// 	for j := range attribute.Options {
	// 		attributeOptionIds = append(attributeOptionIds, attribute.Options[j].Id)
	// 	}

	// 	if !helper.IsArrayContainsInt64(attributeOptionIds, options[i].OptionId) {
	// 		return serviceError(errors.New("invalid option id"), codes.InvalidArgument)
	// 	}
	// }

	// err = s.repo.Cart.CreateCartProduct()
	// if err != nil {
	// 	return serviceError(err, codes.Internal)
	// }

	return nil
}
