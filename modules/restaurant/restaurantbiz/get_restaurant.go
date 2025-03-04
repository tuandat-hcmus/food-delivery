package restaurantbiz

import (
	"context"
	"rest/common"
	"rest/modules/restaurant/restaurantmodel"
)

type GetRestaurantStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
}

type getRestaurantBiz struct {
	store GetRestaurantStore
}

func NewGetRestaurantBiz (store GetRestaurantStore) *getRestaurantBiz {
	return &getRestaurantBiz{store: store}
}

func (biz *getRestaurantBiz) GetRestaurant(
	ctx context.Context, 
	id int,
) (*restaurantmodel.Restaurant, error) {
	result, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id" : id})
	if err != nil {
		if err == common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
		}
		return nil, common.ErrCannotGetEntity(restaurantmodel.EntityName, err) // truong hop la mot loi cu the, nen tao mot loi moi tuong ung
	}
	if result.Status == 0 {
		return nil, common.ErrEntityDeleted(restaurantmodel.EntityName, nil)
	}
	return result, nil
}