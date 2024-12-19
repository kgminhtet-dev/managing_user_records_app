package testutil

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/config"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users/data"
	"gorm.io/gorm"
	"log"
	"os"
)

func setupEnvironment() {
	if err := os.Setenv("ENV", "testing"); err != nil {
		log.Fatal("Error setting environment variables:", err)
	}
}

func loadConfiguration() *config.Config {
	cfg, err := config.Load("../config.yaml")
	if err != nil {
		log.Fatal("Error loading configuration:", err)
	}
	return cfg
}

func setupDatabase(cfg *config.Config) *gorm.DB {
	db := data.New(cfg.Database)
	return db
}

func SeedDatabase(db *gorm.DB, users []*data.User) {
	db.CreateInBatches(users, len(users))
}

func Setup() *gorm.DB {
	setupEnvironment()
	cfg := loadConfiguration()
	db := setupDatabase(cfg)
	return db
}

func GenerateRandomUsers(size int) []*data.User {
	users := make([]*data.User, size)

	for i, _ := range users {
		user := data.User{}
		user.ID = uuid.New().String()
		user.Name = fmt.Sprintf("user%d", i+1)
		user.Email = fmt.Sprintf("user%d@example.com", i+1)
		users[i] = &user
	}

	return users
}
