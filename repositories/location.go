package repositories

import (
	"context"

	"srating/domain"

	"gorm.io/gorm"
)

type locationRepository struct {
	Repository
}

func NewLocationRepository(db *gorm.DB) domain.LocationRepository {
	return &locationRepository{
		Repository{
			database: db,
		},
	}
}

func (r *locationRepository) CreateLocation(c context.Context, location *domain.Location) error {
	db := r.GetDB(c)
	return db.Save(location).Error
}

func (r *locationRepository) GetAllLocation(c context.Context, input domain.GetAllLocationRequest) (int64, int64, []*domain.Location, error) {
	locations := []*domain.Location{}
	total := int64(0)

	query := r.GetDB(c).Model(&domain.Location{})

	if input.LocationID != 0 {
		query = query.Where("id = ?", input.LocationID)
	}

	// Count the total number of locations matching the conditions
	if err := query.Count(&total).Error; err != nil {
		return 0, 0, nil, err
	}

	// Apply pagination only if input.Limit is greater than 0
	if input.Limit > 0 {
		query = query.Scopes(r.Paginate(input.Page, input.Limit))
	}

	query = query.Order("updated_at desc")

	// Find the locations with the current conditions and pagination
	if err := query.Find(&locations).Error; err != nil {
		return 0, 0, nil, err
	}

	// The totalCount should be equal to the total because we want all records if no limit is applied
	totalCount := total

	return total, totalCount, locations, nil
}


func (r *locationRepository) GetLocationDetail(c context.Context, id uint) (*domain.Location, error) {
	var location domain.Location
	query := r.GetDB(c)
	if err := query.Where("id = ?", id).First(&location).Error; err != nil {
		return nil, err
	}
	return &location, nil
}

func (r *locationRepository) UpdateLocation(c context.Context, location *domain.Location) error {
	db := r.GetDB(c)
	return db.Save(location).Error
}

func (r *locationRepository) DeleteLocation(c context.Context, id uint) error {
	db := r.GetDB(c)
	return db.Delete(&domain.Location{}, id).Error
}
