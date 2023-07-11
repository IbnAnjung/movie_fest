package movie

import "github.com/IbnAnjung/movie_fest/entity/utils"

type ListMovieWithPaginationInput struct {
	utils.MetaPagination
	Search string
}

type ListMovieWithPaginationOutput struct {
	Meta   utils.MetaPagination
	Movies []Movie
}
