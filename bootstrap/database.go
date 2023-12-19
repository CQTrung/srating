package bootstrap

import (
	"fmt"
	"time"

	"srating/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

func NewPostgresDatabase(env *Env) *gorm.DB {
	username := env.DBUser
	password := env.DBPassword
	dbName := env.DBName
	dbHost := env.DBHost // Change this to your PostgreSQL host
	dbPort := env.DBPort // Change this to your PostgreSQL port
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d TimeZone=UTC", username, password, dbName, dbHost, dbPort)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{
		// SkipDefaultTransaction: true,
		// PrepareStmt:            true,
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		utils.LogFatal(err, "Failed to connect to database")
		return nil
	}

	// Set database connection pool configurations here, if required.
	// For example, you can set MaxIdleConns, MaxOpenConns, etc.
	err = db.Use(
		dbresolver.Register(dbresolver.Config{}).
			SetConnMaxIdleTime(time.Hour).
			SetConnMaxLifetime(2 * time.Minute).
			SetMaxIdleConns(4).
			SetMaxOpenConns(90),
	)
	if err != nil {
		utils.LogFatal(err, "Failed to connect to database")
		return nil
	}
	return db
}

func ClosePostgreConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		utils.LogFatal(err, "Failed to close database")
		return
	}
	dbSQL.Close()
}
