package simple

import (
	"github.com/dgraph-io/ristretto"
	"github.com/dollarkillerx/go-proxy-manager/internal/utils"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
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

	sim := Simple{
		cache: cache,
		db:    db,
	}

	return &sim, nil
}
