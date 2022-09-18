package simple

import (
	"github.com/dgraph-io/ristretto"
	"github.com/dollarkillerx/go-proxy-manager/internal/app/backend/pkg/models"
	"github.com/dollarkillerx/go-proxy-manager/internal/utils"
	"github.com/glebarez/sqlite"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"log"
)

type Simple struct {
	cache *ristretto.Cache
	db    *gorm.DB
}

func NewSimple() (*Simple, error) {
	cache, err := utils.NewRCache()
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(
		&models.UserCenter{},
		&models.Tasks{},
	)

	var userCount int64
	err = db.Model(&models.UserCenter{}).Count(&userCount).Error
	if err != nil {
		log.Println(err)
	}
	if err == nil && userCount == 0 {
		sale := uuid.New().String()
		err = db.Model(&models.UserCenter{}).Create(&models.UserCenter{
			BasicModel: models.BasicModel{
				ID: uuid.New().String(),
			},
			Account:  "manager@manager.com",
			Password: utils.GenPassword("admin", sale),
			Sale:     sale,
		}).Error
		if err != nil {
			log.Fatalln(err)
		}
	}

	sim := Simple{
		cache: cache,
		db:    db,
	}

	return &sim, nil
}
