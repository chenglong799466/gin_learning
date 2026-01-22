package service

import (
	"gin_learning/internal/repo"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

type CreateUserReq struct {
	Name string `json:"name" binding:"required"`
}

func (s *UserService) CreateUser(req *CreateUserReq) (*repo.User, error) {

	var user *repo.User
	user.Name = req.Name

	userRepo := repo.NewUserRepo()
	if err := userRepo.CreateUser(user); err != nil {
		return nil, err
	}
	return user, nil
}
