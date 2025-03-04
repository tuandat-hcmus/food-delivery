package restaurantstorage

import (
	"context"
	"rest/common"
	"rest/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) SoftDeleteData(
	ctx context.Context, 
	id int, // co the thay bang conditions
) error {
	db := s.db

	if err := db.
		Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status" : 0,
		}).Error; err != nil {
			return common.ErrDB(err)
		}
		return nil
}