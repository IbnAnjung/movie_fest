package movie

type MovieDetail struct {
	ID           int64
	Filename     string
	Title        string
	Duration     int64
	Artists      string
	Genres       []MovieDetailGenres
	Description  string
	PublicUrl    string
	ViewsCounter int64
	VotesCounter int64
}

type MovieDetailGenres struct {
	ID   int32
	Name string
}
