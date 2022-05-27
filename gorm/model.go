package gorm

import (
	"context"
	"github.com/god-jay/generics-orm"
	"gorm.io/gorm"
)

type Model[M any] struct {
	ds    *gorm.DB //data source
	model M
}

func NewModel[M any](ctx context.Context, db *gorm.DB, m M) orm.Builder[M] {
	return &Model[M]{
		ds:    db.Model(m).WithContext(ctx),
		model: m,
	}
}
