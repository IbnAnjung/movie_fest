package handler

import (
	"fmt"
	"net/http"

	enAuth "github.com/IbnAnjung/movie_fest/entity/authentication"
	"github.com/IbnAnjung/movie_fest/handler/presenters"
	"github.com/IbnAnjung/movie_fest/utils"
	"github.com/gin-gonic/gin"
)

type authenticationHandler struct {
	authUC enAuth.AuthenticationUsecase
}

func NewAuthenticationHandler(authUC enAuth.AuthenticationUsecase) authenticationHandler {
	return authenticationHandler{
		authUC: authUC,
	}
}

func (h authenticationHandler) RegisterUser(c *gin.Context) {
	req := presenters.RegisterUserRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.GeneralErrorResponse(c, err)
		return
	}
	fmt.Println(req)
	registerdUser, err := h.authUC.RegisterUser(c, enAuth.Register{
		Username:        req.Username,
		Password:        req.Password,
		ConfirmPassword: req.ConfirmPassword,
	})
	if err != nil {
		utils.GeneralErrorResponse(c, err)
		return
	}

	response := presenters.RegisterUserResponse{
		ID:       registerdUser.ID,
		Username: registerdUser.Username,
		Token:    registerdUser.Token,
	}

	utils.SuccessResponse(c, http.StatusOK, "success", response)
}
