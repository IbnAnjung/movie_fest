package utils

type MetaPagination struct {
	Limit     int
	Page      int
	Offset    int
	TotalPage int
	TotalRaw  int64
}

type Pagination interface {
	Init(meta *MetaPagination)
	SetTotalRaw(totalRow int64)
	GetMeta() MetaPagination
}
