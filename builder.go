package orm

type Builder[M any] interface {
	condition[M]
	executor[M]
}

type condition[M any] interface {
	Select(query any, args ...any) Builder[M]
	Where(query any, args ...any) Builder[M]
	WhereIn(column string, values ...any) Builder[M]
	WhereBetween(column string, from any, to any) Builder[M]
	OrWhere(query any, args ...any) Builder[M]
	OrWhereIn(column string, values ...any) Builder[M]
	GroupBy(group string, groups ...string) Builder[M]
	OrderBy(column string, direction string) Builder[M]
	Limit(limit int) Builder[M]
	With(relation string, args ...any) Builder[M]
	Scopes(scopes ...func(Builder[M]) Builder[M]) Builder[M]
	Raw(sql string, values ...any) Builder[M]
}

type executor[M any] interface {
	Create() (M, error)
	Upsert() (M, error)
	First() (M, error)
	Last() (M, error)
	All() ([]M, error)
	Delete() error
	Insert(values ...any) error
	Increment(column string, value int) error
	Decrement(column string, value int) error
	Paginate(page int32, pageSize int32) ([]M, *Paginator, error)
}

//type Callback[M any] func(b Builder[M])
