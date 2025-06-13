package booklocationmodel

import (
	"book-store-management-backend/common"
	"errors"
)

type ReqCreateBookLocation struct {
	ID         *string `json:"id" gorm:"column:id"`
	BookID     *string `json:"bookId" gorm:"column:bookId,fk"`
	Location   *string `json:"location" gorm:"column:location"`
	QRCodePath *string `json:"qrCodePath" gorm:"column:qrCodePath"`
}

func (*ReqCreateBookLocation) TableName() string {
	return common.TableBookLocation
}

func (b *ReqCreateBookLocation) Validate() error {
	if b.BookID == nil || common.ValidateEmptyString(*b.BookID) {
		return errors.New("book ID is required")
	}
	if b.Location == nil || common.ValidateEmptyString(*b.Location) {
		return errors.New("location is required")
	}
	if b.QRCodePath == nil || common.ValidateEmptyString(*b.QRCodePath) {
		return errors.New("QRCodePath is required")
	}
	return nil
}
