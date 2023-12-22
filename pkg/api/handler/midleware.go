package handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/athunlal/Note-Taking-Application/pkg/domain"
	"github.com/gin-gonic/gin"
)

func (h *UserHandler) Validate(Accesstoken string) (*domain.User, error) {

	userData := domain.User{}
	ok, _ := h.jwtUseCase.VerifyToken(Accesstoken)
	if !ok {
		return nil, errors.New("Token failed")
	}
	userData, err := h.jwtUseCase.ValidateJwtUser(userData.ID)
	if err != nil {
		return nil, err
	}
	return &userData, nil

}

func (h *UserHandler) AuthRequired(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("Authorization")

	if authorization == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	token := strings.Split(authorization, "Bearer ")
	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	res, err := h.Validate(token[1])
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err,
		})
	}

	str := strconv.FormatInt(int64(res.ID), 10)
	ctx.Set("userId", str)
	ctx.Next()
}
