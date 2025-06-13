package booklocationstore

import (
	"book-store-management-backend/common"
	"time"
)

type BookLocationDbModel struct {
	ID         *string    `json:"id" gorm:"column:id"`
	BookID     *string    `json:"bookId" gorm:"column:bookid; fk"`
	Location   *string    `json:"location" gorm:"column:location"`
	QRCodePath *string    `json:"qrCodePath" gorm:"column:qrcodepath"`
	CreatedAt  *time.Time `json:"createdAt,omitempty" gorm:"createdAt; column:createdAt;"`
	UpdatedAt  *time.Time `json:"updatedAt,omitempty" gorm:"updatedAt; column:updatedAt;"`
	DeletedAt  *time.Time `json:"deletedAt,omitempty" gorm:"deletedAt; column:deletedAt;"`
	IsActive   *bool      `json:"isActive,omitempty" gorm:"isActive; column:isActive; default:1"`
}

func (*BookLocationDbModel) TableName() string {
	return common.TableBookLocation
}
