package repositories

import (
	"context"
	"errors"

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
	if err := r.GetDB(c).Preload("Avatar").Preload("Department").First(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetUserByUsername(c context.Context, username string) (*domain.User, error) {
	var user *domain.User
	if err := r.GetDB(c).Model(&domain.User{}).Preload("Avatar").Preload("Department").Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) UpdateUser(c context.Context, user *domain.User) error {
	media := &domain.Media{Model: domain.Model{ID: user.MediaID}}
	if err := r.GetDB(c).Model(&user).Association("Avatar").Replace(media); err != nil {
		return err
	}
	return r.GetDB(c).Model(&domain.User{}).Where("id = ?", user.ID).Updates(user).Error
}

func (r *userRepository) ChangeStatus(c context.Context, id uint, status domain.Status) error {
	return r.GetDB(c).Model(&domain.User{}).Where("id = ?", id).Update("status", status).Error
}

func (r *userRepository) GetAllEmployee(c context.Context, input domain.GetAllUserRequest) (int64, int64, []*domain.User, error) {
	query := r.GetDB(c).Model(&domain.User{}).Preload("Avatar").Preload("Department").Where("role = ?", domain.EmployeeRole).Order("updated_at desc")
	total := int64(0)
	if err := query.Count(&total).Error; err != nil {
		return 0, 0, nil, err
	}
	users := []*domain.User{}
	query = query.Scopes(r.Paginate(input.Page, input.Limit)).Order("updated_at desc")
	if err := query.Find(&users).Error; err != nil {
		return 0, 0, nil, err
	}
	totalCount := int64(0)
	if err := query.Count(&totalCount).Error; err != nil {
		return 0, 0, nil, err
	}
	return total, totalCount, users, nil
}

func (r *userRepository) CountUserByRole(c context.Context) ([]*domain.GetUserByRoleResponse, error) {
	result := []*domain.GetUserByRoleResponse{}
	query := r.GetDB(c)
	if err := query.Model(&domain.User{}).Select("role, count(*)").Group("role").Scan(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *userRepository) CountTotalField(c context.Context) (int64, error) {
	var count int64
	if err := r.GetDB(c).Model(&domain.User{}).Select("count(distinct field)").Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *userRepository) DeleteEmployee(c context.Context, id uint) error {
	return r.GetDB(c).Delete(&domain.User{}, id).Error
}

func (r *userRepository) CreateUser(c context.Context, user *domain.User) error {
	return r.GetDB(c).Save(user).Error
}

func (r *userRepository) ChangePassword(c context.Context, id uint, oldPassword, newPassword string) error {
	query := "UPDATE users SET password = ? WHERE id = ? AND password = ?"
	// Execute the raw SQL query with parameters
	result := r.GetDB(c).Exec(query, newPassword, id, oldPassword)
	if err := result.Error; err != nil {
		return err
	}
	if rowsAffected := result.RowsAffected; rowsAffected == 0 {
		return errors.New("no rows were updated")
	}
	return nil
}

func (r *userRepository) ResetPassword(c context.Context, id uint, newPassword string) error {
	query := "UPDATE users SET password = ? WHERE id = ?"
	// Execute the raw SQL query with parameters
	result := r.GetDB(c).Exec(query, newPassword, id)
	if err := result.Error; err != nil {
		return err
	}
	if rowsAffected := result.RowsAffected; rowsAffected == 0 {
		return errors.New("no rows were updated")
	}
	return nil
}
