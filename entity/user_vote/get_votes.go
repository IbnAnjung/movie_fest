package user_vote

type GetViewsInput struct {
	MinVotes int64
	MaxVotes int64
	Sort     string
}

type GetViewsOutput struct {
	ID       int64
	Title    string
	Duration int64
	Artists  string
	Votes    int64
}
