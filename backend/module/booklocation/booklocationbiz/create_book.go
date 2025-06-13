package booklocationbiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/generator"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/booklocation/booklocationmodel"
	"context"
)

type CreateBookLocationRepo interface {
	CreateBookLocation(ctx context.Context, data *booklocationmodel.BookLocation) error
}

type createBookLocationBiz struct {
	gen       generator.IdGenerator
	repo      CreateBookLocationRepo
	requester middleware.Requester
}

func NewCreateBookLocationBiz(
	gen generator.IdGenerator,
	repo CreateBookLocationRepo,
	requester middleware.Requester) *createBookLocationBiz {
	return &createBookLocationBiz{
		gen:       gen,
		repo:      repo,
		requester: requester,
	}
}

func (biz *createBookLocationBiz) CreateBookLocation(ctx context.Context, reqData *booklocationmodel.ReqCreateBookLocation, resData *booklocationmodel.ResCreateBookLocation) error {
	// Check if the user has permission to create a book location
	if false && !biz.requester.IsHasFeature(common.BookLocationCreateFeatureCode) {
		return booklocationmodel.ErrBookLocationCreateNoPermission
	}

	// Validate the request data
	if err := reqData.Validate(); err != nil {
		return err
	}

	// Prepare the data to be inserted
	data := &booklocationmodel.BookLocation{
		ID:         reqData.ID,
		BookID:     reqData.BookID,
		Location:   reqData.Location,
		QRCodePath: reqData.QRCodePath,
	}

	// Handle ID generation
	if err := handleBookLocationId(biz.gen, data); err != nil {
		return booklocationmodel.ErrBookLocationIdInvalid
	}

	// Call the repository method to create the book location
	if err := biz.repo.CreateBookLocation(ctx, data); err != nil {
		return err
	}

	// Set the response data
	resData.Id = *data.ID
	return nil
}

func handleBookLocationId(gen generator.IdGenerator, data *booklocationmodel.BookLocation) error {
	if data.ID != nil {
		return nil
	}

	id, err := gen.IdProcess(data.ID)
	if err != nil {
		return err
	}
	data.ID = id
	return nil
}
