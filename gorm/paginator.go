package gorm

import (
	"github.com/god-jay/generics-orm"
	"gorm.io/gorm"
)

func paginate(page *int32, pageSize *int32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if *page == 0 {
			*page = 1
		}
		if *pageSize <= 0 {
			*pageSize = 10
		}

		offset := (*page - 1) * *pageSize
		return db.Offset(int(offset)).Limit(int(*pageSize))
	}
}

func paginator(find *gorm.DB, loadTo any, page int32, size int32) (*orm.Paginator, error) {
	var total int64
	find.Count(&total)

	err := find.Scopes(paginate(&page, &size)).Find(loadTo).Error

	return &orm.Paginator{
		CurrentPage: page,
		PageSize:    size,
		Total:       int32(total),
	}, err
}
