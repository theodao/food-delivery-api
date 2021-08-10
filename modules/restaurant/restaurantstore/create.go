package restaurantstore

import (
	"context"
	"go-api-training/modules/restaurant/restaurantmodel"
)

func (s *sqlStrore) Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
