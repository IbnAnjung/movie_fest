package presenters

type UploadMovie struct {
	ID  int64  `json:"id"`
	Url string `json:"url"`
}

type UpdateMovieMetaDataRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Duration    int64   `json:"duration"`
	Artists     string  `json:"artists"`
	Genres      []int32 `json:"genres"`
	MovieID     int64   `json:"movie_id"`
}

type DetailMovieResponse struct {
	ID           int64                      `json:"id"`
	Title        string                     `json:"title"`
	Duration     int64                      `json:"duration"`
	Artists      string                     `json:"artist"`
	Genres       []MovieDetailGenreResponse `json:"genres"`
	Description  string                     `json:"description"`
	PublicUrl    string                     `json:"public_url"`
	ViewsCounter int64                      `json:"views_counter"`
	VotesCounter int64                      `json:"votes_counter"`
}

type MovieDetailGenreResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type MovieMostViewResponse struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Views int64  `json:"views"`
}

type ListMoviePaginationRequest struct {
	PaginationRequest
	Search string `form:"search"`
	Views  string `form:"views"`
}

type ListMoviePaginationResponse struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Views int64  `json:"views"`
	Votes int64  `json:"votes"`
}
