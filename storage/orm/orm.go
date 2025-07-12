package orm

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"github.com/resetcentral/media_transfer/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GormStorage struct {
	db *gorm.DB
}

type DBConfig struct {
	Host     string
	Name     string
	User     string
	Password string
}

func NewGormStorage() (GormStorage, error) {
	var storage GormStorage
	var dbConfig DBConfig

	err := envconfig.Process("db", &dbConfig)
	if err != nil {
		return storage, err
	}
	dsnFmt := "%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(dsnFmt, dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return storage, err
	}

	err = db.AutoMigrate(&models.Task{})

	if err != nil {
		return storage, err
	}

	storage.db = db
	return storage, nil
}
