package util

import (
	"math"

	"gorm.io/gorm"
)

type Pagination struct {
	Pagination PaginationInfo `json:"pagination"`
	Data       interface{}    `json:"data"`
}

type PaginationInfo struct {
	Limit      int   `json:"limit"`
	Page       int   `json:"page"`
	TotalRows  int64 `json:"total_rows"`
	TotalPages int   `json:"total_pages"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Pagination.Limit == 0 {
		p.Pagination.Limit = 10
	}
	return p.Pagination.Limit
}

func (p *Pagination) GetPage() int {
	if p.Pagination.Page == 0 {
		p.Pagination.Page = 1
	}
	return p.Pagination.Page
}

func Paginate(totalRows int64, sort string, pagination *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	pagination.Pagination.TotalRows = totalRows
	var totalPages int
	pagination.Pagination.Limit = pagination.GetLimit()
	totalPages = int(math.Ceil(float64(totalRows) / float64(pagination.Pagination.Limit)))

	pagination.Pagination.TotalPages = totalPages

	if pagination.Pagination.Page > pagination.Pagination.TotalPages {
		pagination.Pagination.Page = pagination.Pagination.TotalPages
	}

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(sort)
	}
}
