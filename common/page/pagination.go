package page

import (
	"gorm.io/gorm"
)

type Page struct {
	PageNum  int `form:"PageNum" json:"pageNum"`   //页码   如果值为-1，返回所有数据
	PageSize int `form:"pageSize" json:"pageSize"` //每页大小 如果值为-1，返回所有数据
}

func Paginate(pageSize, pageIndex int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (pageIndex - 1) * pageSize
		if offset < 0 {
			offset = 0
		}
		return db.Offset(offset).Limit(pageSize)
	}
}

func (m *Page) GetPageNum() int {
	if m.PageNum <= 0 {
		m.PageNum = 1
	}
	return m.PageNum
}

func (m *Page) GetPageSize() int {
	if m.PageSize <= 0 {
		m.PageSize = 10
	}
	return m.PageSize
}
