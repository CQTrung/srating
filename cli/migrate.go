package cli

import (
	"srating/domain"

	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) *cobra.Command {
	return &cobra.Command{
		Use: "migrate",
		RunE: func(cmd *cobra.Command, args []string) error {
			return migrateDatabase(db)
		},
	}
}

func migrateDatabase(db *gorm.DB) error {
	models := []interface{}{
		&domain.Media{},
		&domain.Department{},
		&domain.User{},
		&domain.Feedback{},
		&domain.Category{},
		&domain.FeedbackCategory{},
	}
	return db.AutoMigrate(models...)
}
