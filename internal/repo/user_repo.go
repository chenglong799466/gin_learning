package repo

import "gorm.io/gorm"

type UserRepo interface {
	QueryUsers() ([]*User, error)
	CreateUser(user *User) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo() UserRepo {
	return &userRepo{}
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (r *userRepo) QueryUsers() ([]*User, error) {

	users := make([]*User, 0)
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil

}

func (r *userRepo) CreateUser(user *User) error {
	return r.db.Create(user).Error
}
