package movie

import "github.com/IbnAnjung/movie_fest/entity/utils"

type ListMovieWithPaginationInput struct {
	utils.MetaPagination
	Search   string
	MinViews int64
	MaxViews int64
}

type ListMovieWithPaginationOutput struct {
	Meta   utils.MetaPagination
	Movies []Movie
}
