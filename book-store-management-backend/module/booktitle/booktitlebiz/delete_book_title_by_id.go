package booktitlebiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/booktitle/booktitlemodel"
	"context"
)

type DeleteBookRepo interface {
	DeleteBookTitle(ctx context.Context, id string) error
}

type DeleteRepo interface {
	DeleteBookTitle(ctx context.Context, id string) error
}

type deleteBookBiz struct {
	requester middleware.Requester
	repo      DeleteRepo
}

func NewDeleteBookBiz(requester middleware.Requester, repo DeleteRepo) *deleteBookBiz {
	return &deleteBookBiz{
		requester: requester,
		repo:      repo,
	}
}

func (biz *deleteBookBiz) DeleteBookTitle(ctx context.Context, id string) error {
	if !biz.requester.IsHasFeature(common.BookTitleDeleteFeatureCode) {
		return booktitlemodel.ErrBookTitleDeleteNoPermission
	}

	return biz.repo.DeleteBookTitle(ctx, id)
}
