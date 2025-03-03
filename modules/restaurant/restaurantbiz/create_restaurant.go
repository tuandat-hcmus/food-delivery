package restaurantbiz

import (
	"context"
	"rest/modules/restaurant/restaurantmodel"
) 
// khong he import gi cua storage

type CreateRestaurantStore interface {
	Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
}

type createRestaurantBiz struct {
	store CreateRestaurantStore
}

// dependency injection
func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

func (biz *createRestaurantBiz) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}
	// khong can biet chi tiet tang storage, chi goi va su dung
	err := biz.store.Create(ctx, data)
	return err
}