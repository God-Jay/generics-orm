package orm

type Paginator struct {
	CurrentPage int32 `json:"current_page"`
	PageSize    int32 `json:"page_size"`
	Total       int32 `json:"total"`
}
