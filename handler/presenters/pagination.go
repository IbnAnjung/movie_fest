package presenters

type PaginationRequest struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}
