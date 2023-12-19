package repositories

import (
	"context"

	"srating/domain"

	"gorm.io/gorm"
)

type userRepository struct {
	Repository
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		Repository{
			database: db,
		},
	}
}

func (r *userRepository) GetUserByID(c context.Context, id uint) (*domain.User, error) {
	var user *domain.User
	if err := r.GetDB(c).First(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetUserByUsername(c context.Context, username string) (*domain.User, error) {
	var user *domain.User
	if err := r.GetDB(c).Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) UpdateUser(c context.Context, user *domain.User) error {
	return r.GetDB(c).Save(&user).Error
}

func (r *userRepository) ChangeStatus(c context.Context, id uint, status domain.Status) error {
	if err := r.GetDB(c).Model(&domain.User{}).Where("id = ?", id).Update("status", status).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetAllEmployee(c context.Context) ([]*domain.User, error) {
	var users []*domain.User
	if err := r.GetDB(c).Where("role = ?", domain.EmployeeRole).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
