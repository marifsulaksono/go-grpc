package helpers

import (
	"go-grpc/protobuf/pagination"
	"math"

	"gorm.io/gorm"
)

func PaginationBuilder(sql *gorm.DB, page int64, pagination *pagination.Pagination) (int64, int64) {
	var total int64
	var limit int64 = 10
	var offset int64

	sql.Count(&total)
	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * limit
	}

	pagination.Total = uint64(total)
	pagination.Limit = uint32(limit)
	pagination.CurrentPage = uint32(page)
	pagination.TotalPage = uint32(math.Ceil(float64(total) / float64(limit)))

	return offset, limit
}
