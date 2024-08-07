package database

import (
	"fmt"

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
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?pgbouncer=true", config.ENV.DBUser, config.ENV.DBPassword, config.ENV.DBHost, config.ENV.DBPort, config.ENV.DBName)

	var err error

	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
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
