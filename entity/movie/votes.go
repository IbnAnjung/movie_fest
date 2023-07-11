package movie

type GetVotesInput struct {
	MinVotes int64
	MaxVotes int64
	Sort     string
}
