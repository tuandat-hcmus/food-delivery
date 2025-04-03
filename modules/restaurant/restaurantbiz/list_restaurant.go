package restaurantbiz

import (
	"context"
	"fmt"
	"rest/common"
	"rest/modules/restaurant/restaurantmodel"
)

// khong he import gi cua storage

type ListRestaurantStore interface {
	ListDataByCondition(ctx context.Context,
		conditions map[string]interface{},
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,	 
	) ([]restaurantmodel.Restaurant, error)
}

type LikeStore interface {
	GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error)
}

type listRestaurantBiz struct {
	store ListRestaurantStore
	likeStore LikeStore
}

// dependency injection
func NewListRestaurantBiz(store ListRestaurantStore, likeStore LikeStore) *listRestaurantBiz {
	return &listRestaurantBiz{store: store, likeStore: likeStore}
}

func (biz *listRestaurantBiz) ListRestaurant(
	ctx context.Context, 
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.store.ListDataByCondition(ctx, nil, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListtEntity(restaurantmodel.EntityName, err)
	}
	ids := make([]int, len(result))
	for i := range result {
		ids[i] = result[i].Id
	}

	mapResLike, err := biz.likeStore.GetRestaurantLikes(ctx, ids)
	// Neu loi thi bo qua vi loi nay khong quan trong
	if err != nil {
		fmt.Print("cannot get restaurant likes")
	}

	if v := mapResLike; v != nil {
		for i, item := range result {
			result[i].LikeCount = mapResLike[item.Id]
		}
	}
	return result, nil
}