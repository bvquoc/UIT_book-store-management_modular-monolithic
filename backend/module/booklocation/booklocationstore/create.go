package booklocationstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/booklocation/booklocationmodel"
	"context"
)

func (s *sqlStore) CreateBookLocation(ctx context.Context, data *BookLocationDbModel) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		if gormErr := common.GetGormErr(err); gormErr != nil {
			switch key := gormErr.GetDuplicateErrorKey("PRIMARY"); key {
			case "PRIMARY":
				return booklocationmodel.ErrBookLocationDuplicate
			}
		}
		return common.ErrDB(err)
	}

	return nil
}
