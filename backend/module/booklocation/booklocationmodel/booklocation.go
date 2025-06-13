package booklocationmodel

import (
	"book-store-management-backend/common"
	"time"
)

type BookLocation struct {
	ID         *string `json:"id" gorm:"column:id"`
	BookID     *string `json:"bookId" gorm:"column:bookId;fk"`
	Location   *string `json:"location" gorm:"column:location"`
	QRCodePath *string `json:"qrCodePath" gorm:"column:qrCodePath"`

	CreatedAt *time.Time `json:"createdAt,omitempty" gorm:"createdAt; column:createdAt;"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" gorm:"updatedAt; column:updatedAt;"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"deletedAt; column:deletedAt;"`
	IsActive  *bool      `json:"isActive,omitempty" gorm:"isActive; column:isActive; default:1"`
}

func (*BookLocation) TableName() string {
	return common.TableBookLocation
}
