package presenters

type MovieGenresMostViewResponse struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Views int64  `json:"views"`
}
