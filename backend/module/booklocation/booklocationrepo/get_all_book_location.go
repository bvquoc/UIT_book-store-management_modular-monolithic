package booklocationrepo

import (
	"book-store-management-backend/module/booklocation/booklocationmodel"
	"context"
)

// GetBookLocationStore interface remains unchanged, as it's already correct
type GetBookLocationStore interface {
	// GetAllBookLocation retrieves all book locations from the database
	// based on whether you want to fetch only active locations or all locations.
	GetAllBookLocation(
		ctx context.Context,
		justGetAllActiveLocation bool,
		moreKeys ...string,
	) ([]booklocationmodel.BookLocation, error)
}

type GetAllBookLocationRepo interface {
	// GetAllBookLocation fetches all book locations based on active status
	GetAllBookLocation(
		ctx context.Context,
		justGetAllActiveLocation bool,
	) ([]booklocationmodel.BookLocation, error)
}

type getAllBookLocationRepo struct {
	bookLocationStore GetBookLocationStore
}

func NewGetAllBookLocationRepo(
	bookLocationStore GetBookLocationStore,
) *getAllBookLocationRepo {
	return &getAllBookLocationRepo{
		bookLocationStore: bookLocationStore,
	}
}

// GetAllBookLocation implementation with updated signature
func (repo *getAllBookLocationRepo) GetAllBookLocation(
	ctx context.Context,
	justGetAllActiveLocation bool,
) ([]booklocationmodel.BookLocation, error) {
	// Call the store to get all book locations
	bookLocations, err := repo.bookLocationStore.GetAllBookLocation(
		ctx, justGetAllActiveLocation)
	if err != nil {
		return nil, err
	}

	return bookLocations, nil
}
