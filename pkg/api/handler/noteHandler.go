package handler

import (
	"net/http"

	"github.com/athunlal/Note-Taking-Application/pkg/domain"
	"github.com/athunlal/Note-Taking-Application/pkg/response"
	"github.com/gin-gonic/gin"
)

//User get note handler
func (h *UserHandler) GetNote(ctx *gin.Context) {
	note := domain.Notes{}
	if err := ctx.BindJSON(&note); err != nil {
		response.JsonInputValidation(ctx)
		return
	}
	res, err := h.usecase.GetNotes(note.Sid)
	if err != nil {
		statusCode := http.StatusUnauthorized
		if err.Error() == "Not found" {
			statusCode = http.StatusNotFound
		}
		response.JSONErrorResponse(ctx, statusCode, false, "Fetching notes failed", err)
		return
	}
	response.JSONResponse(ctx, http.StatusOK, true, "Successfully Fetch notes", res)
}

//User CreateNote handler
func (h *UserHandler) CreateNote(ctx *gin.Context) {
	notes := domain.Notes{}
	if err := ctx.Bind(&notes); err != nil {
		response.JsonInputValidation(ctx)
		return
	}
	if err := h.usecase.CreateNotes(notes); err != nil {
		response.JSONErrorResponse(ctx, http.StatusBadRequest, false, "Creating notes failed", err)
		return
	}
	response.JSONResponse(ctx, http.StatusOK, true, "Successfully create note", nil)
}

//user Delete note handler
func (h *UserHandler) DeleteNote(ctx *gin.Context) {
	var note domain.Notes
	if err := ctx.BindJSON(&note); err != nil {
		response.JsonInputValidation(ctx)
		return
	}
	if err := h.usecase.DeleteNote(note); err != nil {
		response.JSONErrorResponse(ctx, http.StatusUnauthorized, false, "Deleting notes failed", err)
		return
	}
	response.JSONResponse(ctx, http.StatusOK, true, "Successfully delete note", nil)
}
