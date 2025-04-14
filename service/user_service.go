package service

import (
	"context"
	"github.com/vanya-egorov/user_api/model"
	"github.com/vanya-egorov/user_api/repository"
	"github.com/vanya-egorov/user_api/usecase"
)

type UserService interface {
	CreateUser(ctx context.Context, input model.UserInput) (*model.User, error)
	GetUsers(ctx context.Context, page, limit int) ([]model.User, error)
	GetUserByID(ctx context.Context, id int) (*model.User, error)
	UpdateUser(ctx context.Context, id int, input model.UserInput) (*model.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	enriched, _ := usecase.EnrichName(input.Name)
	user := &model.User{
		Name:        input.Name,
		Surname:     input.Surname,
		Patronymic:  input.Patronymic,
		Age:         enriched.Age,
		Gender:      enriched.Gender,
		Nationality: enriched.Nationality,
	}
	err := s.repo.Create(user)
	return user, err
}

func (s *userService) GetUsers(ctx context.Context, page, limit int) ([]model.User, error) {
	offset := (page - 1) * limit
	return s.repo.GetAll(offset, limit)
}

func (s *userService) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	return s.repo.GetByID(uint(id))
}

func (s *userService) UpdateUser(ctx context.Context, id int, input model.UserInput) (*model.User, error) {
	user, err := s.repo.GetByID(uint(id))
	if err != nil {
		return nil, err
	}
	user.Name = input.Name
	user.Surname = input.Surname
	user.Patronymic = input.Patronymic
	return user, s.repo.Update(user)
}

func (s *userService) DeleteUser(ctx context.Context, id int) error {
	return s.repo.Delete(uint(id))
}
