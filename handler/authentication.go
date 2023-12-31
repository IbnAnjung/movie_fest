package handler

import (
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

func (h authenticationHandler) Login(c *gin.Context) {
	req := presenters.LoginRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.GeneralErrorResponse(c, err)
		return
	}

	user, err := h.authUC.Login(c, enAuth.Login{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		utils.GeneralErrorResponse(c, err)
		return
	}

	response := presenters.LoginResponse{
		ID:       user.ID,
		Username: user.Username,
		Token:    user.Token,
	}

	utils.SuccessResponse(c, http.StatusOK, "success", response)
}

func (h authenticationHandler) Logout(c *gin.Context) {
	claims, err := getUserJwt(c)
	if err != nil {
		utils.GeneralErrorResponse(c, err)
		return
	}

	if err := h.authUC.Logout(c, claims.ID); err != nil {
		utils.GeneralErrorResponse(c, err)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "success", nil)
}
