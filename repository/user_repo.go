package repository

import (
	"github.com/vanya-egorov/user_api/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.User) error
	GetAll(offset, limit int) ([]model.User, error)
	GetByID(id uint) (*model.User, error)
	Update(user *model.User) error
	Delete(id uint) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *userRepo) GetAll(offset, limit int) ([]model.User, error) {
	var users []model.User
	result := r.db.Offset(offset).Limit(limit).Find(&users)
	return users, result.Error
}

func (r *userRepo) GetByID(id uint) (*model.User, error) {
	var user model.User
	result := r.db.First(&user, id)
	return &user, result.Error
}

func (r *userRepo) Update(user *model.User) error {
	return r.db.Save(user).Error
}

func (r *userRepo) Delete(id uint) error {
	return r.db.Delete(&model.User{}, id).Error
}
