package booklocationmodel

import (
	"book-store-management-backend/common"
	"errors"
)

var (
	ErrBookLocationCreateNoPermission = common.NewCustomError(
		errors.New("no permission to create book location"),
		"Bạn không có quyền tạo vị trí sách",
		"ErrBookLocationCreateNoPermission",
	)
	ErrBookLocationUpdateNoPermission = common.NewCustomError(
		errors.New("no permission to update book location"),
		"Bạn không có quyền chỉnh sửa vị trí sách",
		"ErrBookLocationUpdateNoPermission",
	)
	ErrBookLocationViewNoPermission = common.NewCustomError(
		errors.New("no permission to view book location"),
		"Bạn không có quyền xem vị trí sách",
		"ErrBookLocationViewNoPermission",
	)
	ErrBookLocationDeleteNoPermission = common.NewCustomError(
		errors.New("no permission to delete book location"),
		"Bạn không có quyền xóa vị trí sách",
		"ErrBookLocationDeleteNoPermission",
	)
	ErrBookLocationIdInvalid = common.NewCustomError(
		errors.New("id of BookLocation is invalid"),
		"Mã vị trí sách không hợp lệ",
		"ErrBookLocationIdInvalid",
	)
	ErrBookLocationQRCodePathInvalid = common.NewCustomError(
		errors.New("QR code path for BookLocation is invalid"),
		"Đường dẫn QR code cho vị trí sách không hợp lệ",
		"ErrBookLocationQRCodePathInvalid",
	)
	ErrBookLocationBookIdInvalid = common.NewCustomError(
		errors.New("id of Book for BookLocation is invalid"),
		"Mã sách cho vị trí sách không hợp lệ",
		"ErrBookLocationBookIdInvalid",
	)
	ErrBookLocationLocationEmpty = common.NewCustomError(
		errors.New("location for BookLocation is empty"),
		"Vị trí sách không thể trống",
		"ErrBookLocationLocationEmpty",
	)
	ErrBookLocationDuplicate = common.ErrDuplicateKey(
		errors.New("BookLocation already exists"),
	)
)
