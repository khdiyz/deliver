package service

import (
	"deliver/internal/constants"
	"deliver/internal/repository"
	"deliver/models"
	"deliver/pkg/logger"

	"google.golang.org/grpc/codes"
)

type OptionService struct {
	repo repository.Repository
	log  logger.Logger
}

func NewOptionService(repo repository.Repository, log logger.Logger) *OptionService {
	return &OptionService{
		repo: repo,
		log:  log,
	}
}

func (s *OptionService) Create(option models.OptionCreateRequest) (int64, error) {
	_, err := s.repo.Attribute.GetById(option.AttributeId)
	if err != nil {
		return 0, serviceError(err, codes.NotFound)
	}

	id, err := s.repo.Option.Create(option)
	if err != nil {
		return 0, serviceError(err, codes.Internal)
	}

	return id, nil
}

func (s *OptionService) GetList(pagination *models.Pagination, filters map[string]interface{}) ([]models.Option, error) {
	options, err := s.repo.Option.GetList(pagination, filters)
	if err != nil {
		return nil, serviceError(err, codes.Internal)
	}

	return options, nil
}

func (s *OptionService) Get(attributeId, optionId int64) (models.Option, error) {
	option, err := s.repo.Option.GetById(optionId)
	if err != nil {
		return models.Option{}, serviceError(err, codes.Internal)
	}

	if attributeId != option.AttributeId {
		return models.Option{}, serviceError(constants.ErrorDataIsEmpty, codes.NotFound)
	}

	return option, nil
}

func (s *OptionService) Update(option models.OptionUpdateRequest) error {
	_, err := s.Get(option.AttributeId, option.Id)
	if err != nil {
		return serviceError(constants.ErrorDataIsEmpty, codes.NotFound)
	}

	err = s.repo.Option.Update(option)
	if err != nil {
		return serviceError(err, codes.Internal)
	}

	return nil
}

func (s *OptionService) Delete(attributeId, optionId int64) error {
	_, err := s.Get(attributeId, optionId)
	if err != nil {
		return serviceError(constants.ErrorDataIsEmpty, codes.NotFound)
	}

	err = s.repo.Option.DeleteById(optionId)
	if err != nil {
		return serviceError(err, codes.Internal)
	}

	return nil
}
