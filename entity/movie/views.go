package movie

type GetViewsInput struct {
	MinViews int64
	MaxViews int64
	Sort     string
}
