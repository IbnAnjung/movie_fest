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
