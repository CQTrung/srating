package repositories

import (
	"context"

	"srating/domain"

	"gorm.io/gorm"
)

type departmentRepository struct {
	Repository
}

func NewDepartmentRepository(db *gorm.DB) domain.DepartmentRepository {
	return &departmentRepository{
		Repository{
			database: db,
		},
	}
}

func (r *departmentRepository) CreateDepartment(c context.Context, department *domain.Department) error {
	db := r.GetDB(c)
	return db.Save(department).Error
}

func (r *departmentRepository) GetAllDepartment(c context.Context, input domain.GetAllDepartmentRequest) (int64, int64, []*domain.Department, error) {
	departments := []*domain.Department{}
	total := int64(0)

	query := r.GetDB(c).Model(&domain.Department{})

	if err := query.Count(&total).Error; err != nil {
		return 0, 0, nil, err
	}
	query = query.Scopes(r.Paginate(input.Page, input.Limit)).Order("updated_at desc")
	if err := query.Find(&departments).Error; err != nil {
		return 0, 0, nil, err
	}
	totalCount := int64(0)
	if err := query.Count(&totalCount).Error; err != nil {
		return 0, 0, nil, err
	}

	return total, totalCount, departments, nil
}

func (r *departmentRepository) GetDepartmentDetail(c context.Context, id uint) (*domain.Department, error) {
	var department domain.Department
	query := r.GetDB(c)
	if err := query.Where("id = ?", id).First(&department).Error; err != nil {
		return nil, err
	}
	return &department, nil
}

func (r *departmentRepository) UpdateDepartment(c context.Context, department *domain.Department) error {
	db := r.GetDB(c)
	return db.Save(department).Error
}

func (r *departmentRepository) DeleteDepartment(c context.Context, id uint) error {
	db := r.GetDB(c)
	return db.Delete(&domain.Department{}, id).Error
}
