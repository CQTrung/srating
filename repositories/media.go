package repositories

import (
	"context"

	"srating/domain"

	"gorm.io/gorm"
)

type mediaRepository struct {
	Repository
}

func NewMediaRepository(db *gorm.DB) domain.MediaRepository {
	return &mediaRepository{
		Repository{
			database: db,
		},
	}
}

func (r *mediaRepository) Upload(c context.Context, medias *domain.Media) error {
	db := r.GetDB(c)
	return db.Save(medias).Error
}

func (r *mediaRepository) GetAll(c context.Context) ([]*domain.Media, error) {
	var medias []*domain.Media
	query := r.GetDB(c)
	if err := query.Order("updated_at desc").Find(&medias).Error; err != nil {
		return nil, err
	}
	return medias, nil
}
