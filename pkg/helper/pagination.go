package helper

import (
	"deliver/internal/models"
	"deliver/pkg/logger"
	"math"

	"github.com/jmoiron/sqlx"
)

func GetListCount(db *sqlx.DB, loggers *logger.Logger, pagination *models.Pagination, countFilterQuery string, filtersArray []interface{}) error {
	count, err := fetchItemCount(db, countFilterQuery, filtersArray)
	if err != nil {
		loggers.Error(err)
		return err
	}

	updatePaginationLimits(pagination, count)
	updatePaginationTotals(pagination, count)

	return nil
}

func fetchItemCount(db *sqlx.DB, query string, args []interface{}) (int64, error) {
	var count int64
	err := db.Get(&count, query, args...)
	return count, err
}

func updatePaginationLimits(pagination *models.Pagination, count int64) {
	if count < pagination.Offset+pagination.Limit {
		pageCount := calculatePageCount(count, pagination.PageSize)
		var offset int64
		if pageCount > 1 {
			offset = (pageCount - 1) * pagination.PageSize
		}
		pagination.Limit = count - offset
		pagination.Offset = offset
	}
}

func updatePaginationTotals(pagination *models.Pagination, count int64) {
	pagination.ItemTotal = count
	pagination.PageTotal = calculatePageCount(count, pagination.PageSize)

	if pagination.Page > pagination.PageTotal {
		if pagination.PageTotal == 0 {
			pagination.Page = 1
		} else {
			pagination.Page = pagination.PageTotal
		}
	}
}

func calculatePageCount(count int64, pageSize int64) int64 {
	return int64(math.Ceil(float64(count) / float64(pageSize)))
}
