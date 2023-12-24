package handler

import (
	"net/http"

	"github.com/athunlal/Note-Taking-Application/pkg/domain"
	"github.com/athunlal/Note-Taking-Application/pkg/usecase/interfaces"
	"github.com/athunlal/Note-Taking-Application/pkg/utils"
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

//user register
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

//user login
func (h *UserHandler) Login(ctx *gin.Context) {
	user := domain.User{}
	if err := ctx.BindJSON(&user); err != nil {
		utils.JsonInputValidation(ctx)
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

	ctx.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Message": "User successfully logged in",
		"data":    res,
	})
}

//fetching notes by user
func (h *UserHandler) GetNote(ctx *gin.Context) {
	note := domain.Notes{}
	if err := ctx.BindJSON(&note); err != nil {
		utils.JsonInputValidation(ctx)
		return
	}

	res, err := h.usecase.GetNotes(note.Sid)
	if err != nil {
		if err.Error() == "Not found" {
			ctx.JSON(http.StatusNotFound, gin.H{
				"Success": false,
				"Message": "Fetching notes failed",
				"err":     err.Error(),
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"Success": false,
				"Message": "Fetching notes failed",
				"err":     err.Error(),
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Message": "Successfully Fetch notes",
		"data":    res,
	})
}

//Creating notes by user
func (h *UserHandler) CreateNote(ctx *gin.Context) {
	notes := domain.Notes{}
	if err := ctx.Bind(&notes); err != nil {
		utils.JsonInputValidation(ctx)
		return
	}

	if err := h.usecase.CreateNotes(notes); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Creating notes faild",
			"err":     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Message": "Successfully create note",
	})
}

//Deleting not by note id
func (h *UserHandler) DeleteNote(ctx *gin.Context) {
	var note domain.Notes
	if err := ctx.BindJSON(&note); err != nil {
		utils.JsonInputValidation(ctx)
		return
	}

	if err := h.usecase.DeleteNote(note); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Success": false,
			"Message": "Deleting notes faild",
			"err":     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Message": "Successfully delete note",
	})
}
