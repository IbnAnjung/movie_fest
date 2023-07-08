package utils

import (
	"math"

	enUtil "github.com/IbnAnjung/movie_fest/entity/utils"
)

type pagination struct {
	meta *enUtil.MetaPagination
}

func NewPagination() enUtil.Pagination {
	return &pagination{}
}

func (p *pagination) Init(meta *enUtil.MetaPagination) {

	if meta.Page == 0 {
		meta.Page = 1
	}

	if meta.Limit == 0 {
		meta.Limit = 10
	}

	meta.Offset = (meta.Page - 1) * meta.Limit

	p.meta = meta

}

func (p *pagination) SetTotalRaw(totalRow int64) {
	p.meta.TotalRaw = totalRow
	p.meta.TotalPage = int(math.Ceil(float64(totalRow) / float64(p.meta.Limit)))
}

func (p *pagination) GetMeta() enUtil.MetaPagination {
	return *p.meta
}
