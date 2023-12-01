package publisherstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/publisher/publishermodel"
	"context"
)

func (s *sqlStore) CreatePublisher(ctx context.Context, data *publishermodel.Publisher) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		if gormErr := common.GetGormErr(err); gormErr != nil {
			switch key := gormErr.GetDuplicateErrorKey("PRIMARY"); key {
			case "PRIMARY":
				return publishermodel.ErrPublisherIdDuplicate
			}
		}
		return common.ErrDB(err)
	}

	return nil
}
