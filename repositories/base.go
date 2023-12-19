package repositories

import (
	"context"

	"gorm.io/gorm"
)

type key struct{}

type Repository struct {
	database *gorm.DB
}

func (r *Repository) GetDB(ctx context.Context) *gorm.DB {
	val, ok := ctx.Value(key{}).(*gorm.DB)
	if !ok {
		return r.database.WithContext(ctx)
	}
	return val.WithContext(ctx)
}

func (r *Repository) Transaction(ctx context.Context, callback func(ctx context.Context) error) error {
	return r.GetDB(ctx).Transaction(func(tx *gorm.DB) error {
		txCtx := context.WithValue(ctx, key{}, tx)
		return callback(txCtx)
	})
}

func (r *Repository) Paginate(page, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 && limit == 0 {
			return db
		}
		if page <= 0 {
			page = 1
		}
		if limit <= 0 {
			limit = 10
		}
		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}
}
