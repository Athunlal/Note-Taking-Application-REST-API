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
	usecase    interfaces.UserUseCase
	jwtUseCase interfaces.JwtUseCase
}

func NewUserHandler(useCase interfaces.UserUseCase, jwtUseCase interfaces.JwtUseCase) *UserHandler {
	return &UserHandler{
		usecase:    useCase,
		jwtUseCase: jwtUseCase,
	}
}

func (h *UserHandler) Register(ctx *gin.Context) {
	var user domain.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		utils.JsonInputValidation(ctx)
		return
	}

	if err := h.usecase.RegisterUser(user); err != nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"Success": false,
			"Message": "Registration failed",
			"Error":   err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "User registered successfully",
		})
	}
}

func (h *UserHandler) Login(ctx *gin.Context) {
	user := domain.User{}
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Success": false,
			"Message": "Login credentials failed",
			"err":     err,
		})
		return
	}

	res, err := h.usecase.UserLogin(user)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Success": false,
			"Message": "Login credentials failed",
			"err":     err.Error(),
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
		"data":        res,
		"Accesstoken": accessToken,
	})
}

//fetching notes by user
func (h *UserHandler) GetNotes(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.GetString("userId"))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Success": false,
			"Message": "Login credentials failed",
			"err":     err,
		})
		return
	}

	res, err := h.usecase.GetNotes(id)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Success": false,
			"Message": "Login credentials failed",
			"err":     err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Message": "User successfully logged in",
		"data":    res,
	})
}

//Creating notes by user
func (h *UserHandler) CreateNote(ctx *gin.Context) {
	notes := domain.Notes{}
	if err := ctx.Bind(&notes); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Success": false,
			"Message": "Creating notes faild",
			"err":     err,
		})
		return
	}

	if err := h.usecase.CreateNotes(notes); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Success": false,
			"Message": "Creating notes faild",
			"err":     err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Message": "User successfully logged in",
	})
}
