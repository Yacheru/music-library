package handlers

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"music-library/internal/entities"
	"music-library/internal/service"
	"music-library/pkg/constants"
	"net/http"
	"strconv"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) EditSong(ctx *gin.Context) {
	var entitySong = new(entities.Song)
	if err := ctx.ShouldBindJSON(entitySong); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	title := ctx.Query("title")
	link := ctx.Query("link")

	newSong, err := h.service.Music.EditSong(ctx, title, link, entitySong)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			NewErrorResponse(ctx, http.StatusNotFound, "Song not found")
			return
		}

		NewErrorResponse(ctx, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	NewSuccessResponse(ctx, http.StatusOK, "Song edited successfully", newSong)
}

func (h *Handler) DeleteSong(ctx *gin.Context) {
	title := ctx.Query("title")
	link := ctx.Query("link")

	if err := h.service.Music.DeleteSong(ctx, title, link); err != nil {
		if errors.Is(err, constants.SongNotFoundError) {
			NewErrorResponse(ctx, http.StatusNotFound, "song not found")
			return
		}

		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessResponse(ctx, http.StatusOK, "Song successfully deleted", nil)
	return
}

func (h *Handler) GetVerse(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))

	title := ctx.Query("title")
	link := ctx.Query("title")

	verses, err := h.service.Music.GetVerse(ctx, title, link, limit, offset)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			NewErrorResponse(ctx, http.StatusNotFound, "song not found")
			return
		}

		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	NewSuccessResponse(ctx, http.StatusOK, "Verses", verses)
}

func (h *Handler) GetAllSongs(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))

	var (
		songs []*entities.Song
		err   error
	)

	filter, exists := ctx.Get("filter")
	if exists {
		songs, err = h.service.Music.GetAllSongs(ctx, limit, offset, filter.(*entities.Filter))
	} else {
		songs, err = h.service.Music.GetAllSongs(ctx, limit, offset, nil)
	}

	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	if songs == nil {
		NewErrorResponse(ctx, http.StatusNotFound, "No songs found")
		return
	}

	NewSuccessResponse(ctx, http.StatusOK, "All songs here!", songs)
	return
}

func (h *Handler) StorageNewSong(ctx *gin.Context) {
	var newSongEntity = new(entities.Song)
	if err := ctx.ShouldBindJSON(newSongEntity); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err := h.service.Music.StorageNewSong(ctx, newSongEntity)
	if err != nil {
		if errors.Is(err, constants.NoLyricsFoundError) {
			NewErrorResponse(ctx, http.StatusNotFound, err.Error())
			return
		}
		if errors.Is(err, constants.TimeOutError) {
			NewErrorResponse(ctx, http.StatusRequestTimeout, err.Error())
			return
		}
		if errors.Is(err, constants.SongAlreadyExistsError) {
			NewErrorResponse(ctx, http.StatusConflict, err.Error())
			return
		}

		NewErrorResponse(ctx, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	NewSuccessResponse(ctx, http.StatusOK, "message saved successfully", nil)
	return
}
