package service

import (
	"github.com/Wenev/SheetViz/backend/galactus-service/dtos"
	"github.com/Wenev/SheetViz/backend/galactus-service/model"
	"github.com/Wenev/SheetViz/backend/galactus-service/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(req dtos.CreateUserRequest) (*dtos.UserResponse, error) {
	// if len(req.Username) > 20 || len(req.Username) < 6 {
	// 	return nil, error{}
	// }
	user := &model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	err := s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return &dtos.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (s *UserService) UpdateUser(id uint, req dtos.UpdateUserRequest) (*dtos.UserResponse, error) {

	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	err = s.userRepo.Update(user)
	if err != nil {
		return nil, err
	}

	return &dtos.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (s *UserService) DeleteUser(id uint) error {
	return s.userRepo.Delete(id)
}

func (s *UserService) GetUserById(id uint) (*dtos.UserResponse, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return &dtos.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (s *UserService) GetUsers() (*dtos.UserListResponse, error) {
	user, err := s.userRepo.FindAll()
	if err != nil {
		return nil, err
	}
	result := make([]dtos.UserResponse, 0, len(user))
	for _, u := range user {
		result = append(result, dtos.UserResponse{
			ID:       u.ID,
			Username: u.Username,
			Email:    u.Email,
		})
	}

	return &dtos.UserListResponse{Data: result}, nil
}
