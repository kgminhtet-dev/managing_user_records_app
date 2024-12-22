package data

import (
	"fmt"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func newSqlite() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatalf("error connecting to sqlite database: %v", err)
	}
	return db
}

func newPostgres(cfg *config.DatabaseConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.DbName, cfg.Port, cfg.SSLMode, cfg.TimeZone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error connecting to %s database: %v", cfg.Name, err)
	}
	return db
}

func New(cfg *config.DatabaseConfig) *gorm.DB {
	var db *gorm.DB

	switch cfg.Name {
	case "sqlite":
		db = newSqlite()
	default:
		db = newPostgres(cfg)
	}

	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatalf("Migration error %v", err)
	}

	return db
}
