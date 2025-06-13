package booklocationstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/booklocation/booklocationmodel"
	"context"
)

// GetAllBookLocation retrieves all book locations from the database
// based on whether you want to fetch only active locations or all locations.
func (s *sqlStore) GetAllBookLocation(
	ctx context.Context,
	justGetAllActiveLocation bool,
	moreKeys ...string) ([]booklocationmodel.BookLocation, error) {

	var result []booklocationmodel.BookLocation
	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	// Query for the book locations, filtering based on the active status if required
	if err := db.
		Where("isActive = ?", justGetAllActiveLocation).
		Order("location").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
