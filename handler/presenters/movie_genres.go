package presenters

type MovieGenresMostViewResponse struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Views int64  `json:"views"`
}

type MovieGenresMostVoteResponse struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Votes int64  `json:"votes"`
}
