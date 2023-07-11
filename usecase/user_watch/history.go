package user_watch

import (
	"context"

	enUserWatch "github.com/IbnAnjung/movie_fest/entity/user_watch"
)

type historyValidationObject struct {
	UserID int64 `validate:"required"`
}

func (i *historyValidationObject) set(src enUserWatch.HistoryInput) {
	i.UserID = src.UserID
}

func (uc userWatchUC) History(ctx context.Context, input enUserWatch.HistoryInput) (output []enUserWatch.UserWathHistory, err error) {

	iv := historyValidationObject{}
	iv.set(input)

	if err = uc.validator.ValidateStruct(&iv); err != nil {
		return
	}

	output, err = uc.userWatchRepository.GetUserWathHistory(ctx, input.UserID)
	if err != nil {
		return
	}

	return
}
