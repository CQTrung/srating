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
	if err := r.GetDB(c).Preload("Department").First(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetUserByUsername(c context.Context, username string) (*domain.User, error) {
	var user *domain.User
	if err := r.GetDB(c).Preload("Department").Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) UpdateUser(c context.Context, user *domain.User) error {
	return r.GetDB(c).Save(user).Error
}

func (r *userRepository) ChangeStatus(c context.Context, id uint, status domain.Status) error {
	// Construct the raw SQL update query
	query := "UPDATE users SET status = ? WHERE id = ?"

	// Execute the raw SQL query with parameters
	if err := r.GetDB(c).Exec(query, status, id).Error; err != nil {
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

func (r *userRepository) UpdateEmployee(c context.Context, user *domain.User) error {
	// Assuming your table name is "users"
	query := `
        UPDATE users
        SET username = ?, password = ?, phone = ?, email = ?, short_name = ?, full_name = ?,
            field = ?, department_id = ?, role = ?, status = ?
        WHERE id = ?
    `
	args := []interface{}{
		user.Username, user.Password, user.Phone, user.Email, user.ShortName,
		user.FullName, user.Field, user.DepartmentID, user.Role, user.Status, user.ID,
	}

	result := r.GetDB(c).Exec(query, args...)

	if result.Error != nil {
		return result.Error
	}

	// Check the number of rows affected to ensure that the update was successful
	if rowsAffected := result.RowsAffected; rowsAffected == 0 {
		return errors.New("no rows were updated")
	}

	return nil
}

func (r *userRepository) DeleteEmployee(c context.Context, id uint) error {
	return r.GetDB(c).Delete(&domain.User{}, id).Error
}

func (r *userRepository) CreateUser(c context.Context, user *domain.User) error {
	return r.GetDB(c).Create(&user).Error
}
