package pagination

type PaginationResponse[T any] struct {
	Data        []*T `json:"data"`
	Page        int  `json:"page"`
	Limit       int  `json:"limit"`
	Total       int  `json:"total"`
	TotoalPages int  `json:"total_pages"`
}

func NewPagination[T any](data []*T, page, limit, total, totalPages int) *PaginationResponse[T] {
	return &PaginationResponse[T]{
		Data:        data,
		Page:        page,
		Limit:       limit,
		Total:       total,
		TotoalPages: totalPages,
	}
}
