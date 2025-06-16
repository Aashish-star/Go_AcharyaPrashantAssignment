package repository

import (
"api/internal/repo"

"gorm.io/gorm"
)

type UserRepository interface {
FindByEmail(email string) (*entity.UserDetail, error)
FindByUserName(userName string) (*entity.UserDetail, error)
Save(user *entity.UserDetail) error
}

type userRepository struct {
db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
return &userRepository{db: db}
}

func (r *userRepository) FindByEmail(email string) (*entity.UserDetail, error) {
var user entity.UserDetail
if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
return nil, err
}
return &user, nil
}

func (r *userRepository) FindByUserName(userName string) (*entity.UserDetail, error) {
var user entity.UserDetail
if err := r.db.Where("user_name = ?", userName).First(&user).Error; err != nil {
return nil, err
}
return &user, nil
}

func (r *userRepository) Save(user *entity.UserDetail) error {
return r.db.Create(user).Error
}
