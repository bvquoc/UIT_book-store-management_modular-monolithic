package booklocationbiz

import (
	"book-store-management-backend/module/booklocation/booklocationmodel"
	"context"
)

type GetAllBookLocationRepo interface {
	// GetAllBookLocation fetches the list of book locations from the database
	GetAllBookLocation(
		ctx context.Context,
		justGetAllActiveLocation bool,
	) ([]booklocationmodel.BookLocation, error)
}

type getAllBookLocationBiz struct {
	repo GetAllBookLocationRepo
}

func NewGetAllBookLocationBiz(
	repo GetAllBookLocationRepo) *getAllBookLocationBiz {
	return &getAllBookLocationBiz{
		repo: repo,
	}
}

func (biz *getAllBookLocationBiz) GetAllBookLocation(
	ctx context.Context,
	justGetAllActiveLocation bool,
) ([]booklocationmodel.BookLocation, error) {
	// Fetch book locations using the repository
	bookLocations, err := biz.repo.GetAllBookLocation(ctx, justGetAllActiveLocation)
	if err != nil {
		return nil, err
	}

	// If necessary, you can manipulate the data here (e.g., set default values)

	return bookLocations, nil
}
