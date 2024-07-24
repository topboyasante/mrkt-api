package database

import (
	// "fmt"

	"github.com/topboyasante/mrkt-api/api/v1/models"
	"github.com/topboyasante/mrkt-api/api/v1/utils"
	"github.com/topboyasante/mrkt-api/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	DB *gorm.DB
)

func InitDB() error {
	// Use this in test mode
	// dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", config.ENV.DBUser, config.ENV.DBPassword, config.ENV.DBHost, config.ENV.DBPort, config.ENV.DBName)

	// Use this in live mode
	dsn := config.ENV.ConnectionString

	var err error

	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		utils.Logger().Fatalf("failed to open database connection: %v", err)
		return err
	}

	err = DB.AutoMigrate(&models.User{}, &models.Listing{})
	if err != nil {
		utils.Logger().Fatalf("failed to run database migrations: %s", err)
		return err
	}

	utils.Logger().Info("Connected to database and migrations applied")
	return nil
}
