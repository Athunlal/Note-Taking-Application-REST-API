package handler

import (
	"net/http"

	"github.com/athunlal/Note-Taking-Application/pkg/domain"
	"github.com/athunlal/Note-Taking-Application/pkg/response"
	"github.com/athunlal/Note-Taking-Application/pkg/usecase/interfaces"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	usecase interfaces.UserUseCase
}

func NewUserHandler(useCase interfaces.UserUseCase) *UserHandler {
	return &UserHandler{
		usecase: useCase,
	}
}

//User Register handler
func (h *UserHandler) Register(ctx *gin.Context) {
	var user domain.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		response.JsonInputValidation(ctx)
		return
	}
	if err := h.usecase.RegisterUser(user); err != nil {
		response.JSONErrorResponse(ctx, http.StatusConflict, false, "Registration failed", err)
		return
	}
	response.JSONResponse(ctx, http.StatusOK, true, "User registered successfully", nil)

}

//User login handler
func (h *UserHandler) Login(ctx *gin.Context) {
	user := domain.User{}
	if err := ctx.BindJSON(&user); err != nil {
		response.JsonInputValidation(ctx)
		return
	}
	res, err := h.usecase.UserLogin(user)
	if err != nil {
		response.JSONErrorResponse(ctx, http.StatusUnauthorized, false, "Login credentials failed", err)
		return
	}
	response.JSONResponse(ctx, http.StatusOK, true, "User successfully logged in", res)
}
