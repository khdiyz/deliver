package service

import (
	"deliver/internal/models"
	"deliver/internal/repository"
	"deliver/pkg/helper"
	"deliver/pkg/logger"
	"errors"

	"google.golang.org/grpc/codes"
)

type UserService struct {
	repo repository.Repository
	log  logger.Logger
}

func NewUserService(repo repository.Repository, log logger.Logger) *UserService {
	return &UserService{
		repo: repo,
		log:  log,
	}
}

func (s *UserService) Create(input models.UserCreateRequest) (int64, error) {
	_, err := s.repo.Role.GetById(input.RoleId)
	if err != nil {
		return 0, serviceError(errors.New("role with this id does not exist"), codes.InvalidArgument)
	}

	hash, err := helper.GenerateHash(input.Password)
	if err != nil {
		return 0, serviceError(err, codes.Internal)
	}
	input.Password = hash

	id, err := s.repo.User.Create(input)
	if err != nil {
		return 0, serviceError(err, codes.Internal)
	}

	return id, nil
}

func (s *UserService) GetByEmail(email string) (models.User, error) {
	user, err := s.repo.User.GetByEmail(email)
	if err != nil {
		return models.User{}, serviceError(err, codes.Internal)
	}

	return user, err
}

func (s *UserService) GetById(id int64) (models.User, error) {
	user, err := s.repo.User.GetById(id)
	if err != nil {
		return models.User{}, serviceError(err, codes.Internal)
	}

	return user, err
}
