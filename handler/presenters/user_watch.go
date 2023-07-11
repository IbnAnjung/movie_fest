package presenters

type StartWatchRequest struct {
	MovieID   int64 `json:"movie_id"`
	StartTime int64 `json:"start_time"`
}

type StartWatchResponse struct {
	PlaybackID string `json:"playback_id"`
}

type PlaybackRequest struct {
	PlaybackID string `json:"playback_id"`
	Endtime    int64  `json:"end_time"`
}
