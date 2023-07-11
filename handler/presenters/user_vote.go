package presenters

type UserVoteUnvoteRequest struct {
	MovieID int64 `json:"movie_id"`
}

type VotedMoviesResponse struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Duration int64  `json:"duration"`
	Artist   string `json:"artist"`
}
