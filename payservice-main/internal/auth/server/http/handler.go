package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"payservice/internal/auth/auth"
	"payservice/internal/auth/entity"
)

type EndpointHandler struct {
	authService auth.UseCase
	logger      *zap.SugaredLogger
}

func NewEndpointHandler(authService auth.UseCase, logger *zap.SugaredLogger) *EndpointHandler {
	return &EndpointHandler{
		authService: authService,
		logger:      logger,
	}
}

// Login		godoc
// @Summary			Gets data to login in to the system
// @Description		Get login and password to get jwt token for further work in the system
// @Param			request body ReqLogin true "Login"
// @Produce			application/json
// @Success			200 {object} RespLogin
// @Router			/login [post]
func (h EndpointHandler) Login(ctx *gin.Context) {
	request := struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}{}

	err := ctx.ShouldBindJSON(&request)

	if err != nil {
		h.logger.Error("Bad Request: JSON request is BAD!!!")
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Bad request"})
		return
	}

	tokenRequest := auth.GenerateTokenRequest{
		Login:    request.Login,
		Password: request.Password,
	}

	userToken, err := h.authService.GenerateToken(tokenRequest, ctx)

	if err != nil {
		h.logger.Error("Wrong Credentials")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Login Failed": "Wrong credentials",
		})

		return
	}

	response := struct {
		AccessToken  string
		RefreshToken string
	}{
		AccessToken:  userToken.AccessToken,
		RefreshToken: userToken.RefreshToken,
	}

	h.logger.Info("User is authorized")
	ctx.JSON(http.StatusOK, gin.H{
		"accessToken":  response.AccessToken,
		"refreshToken": response.RefreshToken,
	})
}

// Refresh		godoc
// @Summary			Gets refreshToken and gives new Token in to work in system
// @Description		request with refreshToken, obtain new token for continued system access and operation
// @Produce			application/json
// @Success			200 {object} RespLogin
// @Router			/refresh [put]
func (h EndpointHandler) Refresh(ctx *gin.Context) {
	refreshToken := ctx.Request.Header.Get("Authorization")

	result, err := h.authService.RenewToken(refreshToken)
	if err != nil {
		h.logger.Error("Wrong token")
		ctx.JSON(http.StatusOK, "Wrong token")
		return
	}

	response := struct {
		AccessToken  string
		RefreshToken string
	}{
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
	}

	h.logger.Info("Token is renewed")
	ctx.JSON(http.StatusOK, gin.H{
		"accessToken":  response.AccessToken,
		"refreshToken": response.RefreshToken,
	})

}

// Logout		godoc
// @Summary			Delete JWT token from database
// @Description		deletes the token from the database, thus ceasing to serve it
// @Produce			application/json
// @Success			200 {object} string
// @Router			/logout [delete]
func (h EndpointHandler) Logout(ctx *gin.Context) {
	accessToken := ctx.Request.Header.Get("Authorization")

	err := h.authService.DeleteToken(accessToken)

	if err != nil {
		h.logger.Error("Wrong token")
		ctx.JSON(http.StatusOK, "Wrong token")
		return
	}

	h.logger.Info("User is logout")
	ctx.JSON(http.StatusOK, "User is logout")
}

// Register		godoc
// @Summary			UserRegister
// @Description		Registration a new User with Confirmation
// @Param			request body UserRegister true "Register"
// @Produce			application/json
// @Success			200
// @Router			/register [post]
func (h EndpointHandler) Register(ctx *gin.Context) {
	var rq entity.UserRegister

	err := ctx.ShouldBindJSON(&rq)

	if err != nil {
		h.logger.Error("Bad Request: JSON request is BAD!!!")
		ctx.JSON(http.StatusBadRequest, "Bad request")
		return
	}

	err = h.authService.RegisterProcess(rq, ctx)
	if err != nil {
		h.logger.Error("Register process is crashed")
		ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	h.logger.Info("User is Created")
	ctx.JSON(http.StatusOK, "User is Created. And now need to confirm account!")
}

// ConfirmUser		godoc
// @Summary			Confirm a registered user
// @Description		Confirmation of the registered user, for further use in the system
// @Produce			application/json
// @Success			200
// @Router			/confirm [post]
func (h EndpointHandler) ConfirmUser(ctx *gin.Context) {
	var code auth.Code

	err := ctx.ShouldBindJSON(&code)
	if err != nil {
		h.logger.Error("Bad Request: JSON request is BAD!!!")
		ctx.JSON(http.StatusBadRequest, "Bad request")
		return
	}

	err = h.authService.ConfirmUser(code.UCode)
	if err != nil {
		h.logger.Error("Bad User Code")
		ctx.JSON(http.StatusBadRequest, "Bad UserCode")
		return
	}

	h.logger.Info("Confirm process is completed")
	ctx.JSON(http.StatusOK, "User is Confirmed")
}
