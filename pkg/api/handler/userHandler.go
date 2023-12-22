package handler

import (
	"net/http"
	"strconv"

	"github.com/athunlal/Note-Taking-Application/pkg/domain"
	"github.com/athunlal/Note-Taking-Application/pkg/usecase/interfaces"
	"github.com/athunlal/Note-Taking-Application/pkg/utils"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UseCase    interfaces.UserUseCase
	jwtUseCase interfaces.JwtUseCase
}

func NewUserHandler(useCase interfaces.UserUseCase) *UserHandler {
	return &UserHandler{
		UseCase: useCase,
	}
}

func (h *UserHandler) Register(ctx *gin.Context) {
	body := domain.User{}
	err := ctx.BindJSON(&body)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}

	err = h.UseCase.RegisterUser(body)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"Success": false,
			"Message": "Registering the User Failed",
			"Error":   err,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Successfully sent the OTP",
		})
	}
}

func (h *UserHandler) Login(ctx *gin.Context) {
	user := domain.User{}
	err := ctx.Bind(&user)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}

	res, err := h.UseCase.UserLogin(user)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Success": false,
			"Message": "Login credentials failed",
			"err":     err,
		})
		return
	}
	accessToken, err := h.jwtUseCase.GenerateAccessToken(int(user.ID), user.Email, "user")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Success": false,
			"Message": "Login credentials failed",
			"err":     err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Success":     true,
		"Message":     "User successfully logged in",
		"Accesstoken": accessToken,
		"data":        res,
	})

}

func (h *UserHandler) GetNotes(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.GetString("userId"))
	res, err := h.UseCase.GetNotes(id)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Success": false,
			"Message": "Login credentials failed",
			"err":     err,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "User successfully logged in",
			"data":    res,
		})
	}
}
