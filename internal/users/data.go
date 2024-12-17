package users

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func newDatabase(cfg *DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.DbName, cfg.Port, cfg.SSLMode, cfg.TimeZone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("error connecting to database %s: %v", cfg.DbName, err))
	}

	db.AutoMigrate(&User{})
	return db, nil
}
