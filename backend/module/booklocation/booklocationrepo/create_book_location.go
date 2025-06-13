package booklocationrepo

import (
	"book-store-management-backend/module/booklocation/booklocationmodel"
	"book-store-management-backend/module/booklocation/booklocationstore"
	"context"
)

type CreateBookLocationStore interface {
	CreateBookLocation(ctx context.Context, bookLocation *booklocationstore.BookLocationDbModel) error
}

type createBookLocationRepo struct {
	bookLocationStore CreateBookLocationStore
}

func NewCreateBookLocationRepo(
	bookLocationStore CreateBookLocationStore,
) *createBookLocationRepo {
	return &createBookLocationRepo{
		bookLocationStore: bookLocationStore,
	}
}

func (repo *createBookLocationRepo) CreateBookLocation(ctx context.Context, data *booklocationmodel.BookLocation) error {
	// Create BookLocation DbModel
	dbData := booklocationstore.BookLocationDbModel{
		ID:         data.ID,
		BookID:     data.BookID,
		Location:   data.Location,
		QRCodePath: data.QRCodePath,
	}

	// Call the store to create the book location
	if err := repo.bookLocationStore.CreateBookLocation(ctx, &dbData); err != nil {
		return err
	}

	return nil
}
