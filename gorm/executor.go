package gorm

import (
	"fmt"
	"github.com/god-jay/generics-orm"
	"gorm.io/gorm"
)

func (m *Model[M]) Create() (M, error) {
	err := m.ds.Create(&m.model).Error
	return m.model, err
}

func (m *Model[M]) Upsert() (M, error) {
	err := m.ds.Save(&m.model).Error
	return m.model, err
}

func (m *Model[M]) First() (M, error) {
	err := m.ds.First(&m.model).Error
	return m.model, err
}

func (m *Model[M]) Last() (M, error) {
	err := m.ds.Last(&m.model).Error
	return m.model, err
}

func (m *Model[M]) All() ([]M, error) {
	var ms []M
	err := m.ds.Find(&ms).Error
	return ms, err
}

func (m *Model[M]) Delete() error {
	return m.ds.Delete(m.model).Error
}

func (m *Model[M]) Insert(values ...any) error {
	//TODO
	return nil
}

func (m *Model[M]) Increment(column string, value int) error {
	return m.ds.Updates(map[string]any{
		column: gorm.Expr(fmt.Sprintf("%s + %d", column, value)),
	}).Error
}

func (m *Model[M]) Decrement(column string, value int) error {
	return m.ds.Updates(map[string]any{
		column: gorm.Expr(fmt.Sprintf("%s - %d", column, value)),
	}).Error
}

func (m *Model[M]) Paginate(page int32, pageSize int32) ([]M, *orm.Paginator, error) {
	var ms []M
	paginator, err := paginator(m.ds, ms, page, pageSize)
	return ms, paginator, err
}
