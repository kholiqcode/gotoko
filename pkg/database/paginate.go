package database

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"math"
	"strconv"
)

type Pagination struct {
	Limit      int         `json:"limit,omitempty;query:limit"`
	Page       int         `json:"page,omitempty;query:page"`
	Sort       string      `json:"sort,omitempty;query:sort"`
	TotalRows  int64       `json:"total_rows"`
	TotalPages int         `json:"total_pages"`
	Rows       interface{} `json:"rows"`
}

func NewPagination(ctx echo.Context) *Pagination {
	page, err := strconv.Atoi(ctx.QueryParam("page"))
	if err != nil {
		page = 1
	}
	pageSize, err := strconv.Atoi(ctx.QueryParam("pageSize"))
	if err != nil {
		pageSize = 10
	}
	sort := ctx.QueryParam("sortBy")
	var sortWithDirection string
	if sort != "" {
		direction := ctx.QueryParam("sortDesc")
		if direction != "" {
			if direction == "true" || direction == "1" {
				sortWithDirection = sort + " desc"
			} else if direction == "false" || direction == "0" {
				sortWithDirection = sort + " asc"
			}
		}

	}
	return &Pagination{
		Page:  page,
		Limit: pageSize,
		Sort:  sortWithDirection,
	}
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetTotalPage() int {
	return p.TotalPages
}

func (p *Pagination) GetTotalRows() int64 {
	return p.TotalRows
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "Id desc"
	}
	return p.Sort
}

func Paginate(value interface{}, pagination *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)

	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}
