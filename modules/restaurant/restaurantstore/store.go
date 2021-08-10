package restaurantstore

import "gorm.io/gorm"

type sqlStrore struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *sqlStrore {
	return &sqlStrore{db: db}
}
