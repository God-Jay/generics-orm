package gorm

import (
	"fmt"
	"github.com/god-jay/generics-orm"
)

func (m *Model[M]) Select(query any, args ...any) orm.Builder[M] {
	m.ds = m.ds.Select(query, args...)
	return m
}

func (m *Model[M]) Where(query any, args ...any) orm.Builder[M] {
	m.ds = m.ds.Where(query, args...)
	return m
}

func (m *Model[M]) WhereIn(column string, values ...any) orm.Builder[M] {
	m.ds = m.ds.Where(fmt.Sprintf("%s IN ?", column), values)
	return m
}

func (m *Model[M]) WhereBetween(column string, from any, to any) orm.Builder[M] {
	m.ds = m.ds.Where(fmt.Sprintf("%s BETWEEN ? AND ?", column), from, to)
	return m
}

func (m *Model[M]) OrWhere(query any, args ...any) orm.Builder[M] {
	m.ds = m.ds.Or(query, args...)
	return m
}

func (m *Model[M]) OrWhereIn(column string, values ...any) orm.Builder[M] {
	m.ds = m.ds.Or(fmt.Sprintf("%s IN ?", column), values)
	return m
}

func (m *Model[M]) GroupBy(group string, groups ...string) orm.Builder[M] {
	tx := m.ds.Group(group)
	for _, group := range groups {
		tx = tx.Group(group)
	}
	m.ds = tx
	return m
}

func (m *Model[M]) OrderBy(column string, direction string) orm.Builder[M] {
	m.ds = m.ds.Order(fmt.Sprintf("%s %s", column, direction))
	return m
}

func (m *Model[M]) Limit(limit int) orm.Builder[M] {
	m.ds = m.ds.Limit(limit)
	return m
}

func (m *Model[M]) With(relation string, args ...any) orm.Builder[M] {
	m.ds = m.ds.Preload(relation, args...)
	return m
}

func (m *Model[M]) Scopes(scopes ...func(orm.Builder[M]) orm.Builder[M]) orm.Builder[M] {
	//TODO
	//for _, scope := range scopes {
	//}
	//m.ds = m.ds.Scopes(scopes...)
	return m
}

func (m *Model[M]) Raw(sql string, values ...any) orm.Builder[M] {
	m.ds = m.ds.Raw(sql, values...)
	return m
}
