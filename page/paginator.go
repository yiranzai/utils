package page

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	defaultPage     = 1
	defaultPageSize = 10
)

// Paginator 是分页程序计算的结果
type Paginator struct {
	Total     int64       `json:"total"`        // 总条数
	Page      int64       `json:"current_page"` // 当前页码
	PageSize  int64       `json:"page_size"`    // 页面大小
	TotalPage int64       `json:"total_page"`   // 总页数
	Rows      interface{} `json:"rows"`         // 详细数据

	ctx *gin.Context
}

// New 创建分页
func New(c *gin.Context) *Paginator {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("per_page"))
	if page <= 0 {
		page = defaultPage
	}
	if pageSize == 0 {
		pageSize = defaultPageSize

	}
	return &Paginator{Page: int64(page), PageSize: int64(pageSize)}
}

// TotalPages 返回最后一页页码
func (p *Paginator) getTotalPage() int64 {
	if p.Total == 0 || p.PageSize == 0 {
		return 1
	}
	if p.Total%p.PageSize == 0 {
		return p.Total / p.PageSize
	}
	return p.Total/p.PageSize + 1
}

// 分页填充结果
func (p *Paginator) Pagination(rows interface{}) *Paginator {
	p.TotalPage = p.getTotalPage()
	p.Rows = rows
	return p
}

// SQL分页
func (p *Paginator) ApplySQL(db *gorm.DB) *gorm.DB {
	page, pageSize := int(p.Page), int(p.PageSize)
	return db.Limit(pageSize).Offset((page - 1) * pageSize)
}

// TotalPages 返回最后一页页码
func (p *Paginator) SetPageSize(pageSize int64) *Paginator {
	p.PageSize = pageSize
	return p
}
