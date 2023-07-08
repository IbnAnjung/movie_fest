package movie

type UpdateMetaDataInput struct {
	Title       string
	Description string
	Duration    int64
	Artists     string
	Genres      []int32
	MovieID     int64
}
